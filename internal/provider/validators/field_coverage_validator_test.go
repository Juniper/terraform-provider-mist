package validators

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNormalizeFieldPath(t *testing.T) {
	tests := []struct {
		name              string
		inputPath         string
		schemaFields      map[string]*FieldInfo
		mapAttributePaths map[string]bool
		expected          string
		description       string
	}{
		{
			name:      "simple_list_array_index",
			inputPath: "privileges.0.role",
			schemaFields: map[string]*FieldInfo{
				"privileges": {
					Path:     "privileges",
					AttrType: "list_nested",
				},
				"privileges.role": {
					Path:     "privileges.role",
					AttrType: "string",
				},
			},
			mapAttributePaths: map[string]bool{},
			expected:          "privileges.role",
			description:       "Array index in list should be removed",
		},
		{
			name:      "map_with_string_key",
			inputPath: "networks.guest.vlan_id",
			schemaFields: map[string]*FieldInfo{
				"networks": {
					Path:     "networks",
					AttrType: "map_nested",
				},
				"networks.{key}.vlan_id": {
					Path:     "networks.{key}.vlan_id",
					AttrType: "int64",
				},
			},
			mapAttributePaths: map[string]bool{
				"networks": true,
			},
			expected:    "networks.{key}.vlan_id",
			description: "String key in map should be replaced with {key}",
		},
		{
			name:      "map_with_ip_address_key",
			inputPath: "extra_routes.10.0.0.0/8.via",
			schemaFields: map[string]*FieldInfo{
				"extra_routes": {
					Path:     "extra_routes",
					AttrType: "map_nested",
				},
				"extra_routes.{key}.via": {
					Path:     "extra_routes.{key}.via",
					AttrType: "string",
				},
			},
			mapAttributePaths: map[string]bool{
				"extra_routes": true,
			},
			expected:    "extra_routes.{key}.via",
			description: "IP address CIDR in map key should be replaced with {key}",
		},
		{
			name:      "nested_maps",
			inputPath: "vrf_instances.default.extra_routes6.2001:db8::/32.via",
			schemaFields: map[string]*FieldInfo{
				"vrf_instances": {
					Path:     "vrf_instances",
					AttrType: "map_nested",
				},
				"vrf_instances.{key}.extra_routes6": {
					Path:     "vrf_instances.{key}.extra_routes6",
					AttrType: "map_nested",
				},
				"vrf_instances.{key}.extra_routes6.{key}.via": {
					Path:     "vrf_instances.{key}.extra_routes6.{key}.via",
					AttrType: "string",
				},
			},
			mapAttributePaths: map[string]bool{
				"vrf_instances":                     true,
				"vrf_instances.{key}.extra_routes6": true,
			},
			expected:    "vrf_instances.{key}.extra_routes6.{key}.via",
			description: "Nested maps should replace both keys with {key}",
		},
		{
			name:      "hash_symbol_in_list_count",
			inputPath: "dns_servers.#",
			schemaFields: map[string]*FieldInfo{
				"dns_servers": {
					Path:     "dns_servers",
					AttrType: "list",
				},
			},
			mapAttributePaths: map[string]bool{},
			expected:          "dns_servers",
			description:       "Hash symbol used for list count should be removed",
		},
		{
			name:      "multiple_array_indices",
			inputPath: "acl_policies.0.actions.1.dst_tag",
			schemaFields: map[string]*FieldInfo{
				"acl_policies": {
					Path:     "acl_policies",
					AttrType: "list_nested",
				},
				"acl_policies.actions": {
					Path:     "acl_policies.actions",
					AttrType: "list_nested",
				},
				"acl_policies.actions.dst_tag": {
					Path:     "acl_policies.actions.dst_tag",
					AttrType: "string",
				},
			},
			mapAttributePaths: map[string]bool{},
			expected:          "acl_policies.actions.dst_tag",
			description:       "Multiple array indices should all be removed",
		},
		{
			name:      "deeply_nested_map_in_list",
			inputPath: "ospf_areas.0.ospf_networks.192.168.1.0/24.passive",
			schemaFields: map[string]*FieldInfo{
				"ospf_areas": {
					Path:     "ospf_areas",
					AttrType: "list_nested",
				},
				"ospf_areas.ospf_networks": {
					Path:     "ospf_areas.ospf_networks",
					AttrType: "map_nested",
				},
				"ospf_areas.ospf_networks.{key}.passive": {
					Path:     "ospf_areas.ospf_networks.{key}.passive",
					AttrType: "bool",
				},
			},
			mapAttributePaths: map[string]bool{
				"ospf_areas.ospf_networks": true,
			},
			expected:    "ospf_areas.ospf_networks.{key}.passive",
			description: "Map inside list should normalize array index and replace map key",
		},
		{
			name:      "pure_numeric_string_in_map",
			inputPath: "port_usages.100.description",
			schemaFields: map[string]*FieldInfo{
				"port_usages": {
					Path:     "port_usages",
					AttrType: "map_nested",
				},
				"port_usages.{key}.description": {
					Path:     "port_usages.{key}.description",
					AttrType: "string",
				},
			},
			mapAttributePaths: map[string]bool{
				"port_usages": true,
			},
			expected:    "port_usages.{key}.description",
			description: "Numeric string as map key should be replaced with {key}, not treated as array index",
		},
		{
			name:      "root_level_field",
			inputPath: "org_id",
			schemaFields: map[string]*FieldInfo{
				"org_id": {
					Path:     "org_id",
					AttrType: "string",
				},
			},
			mapAttributePaths: map[string]bool{},
			expected:          "org_id",
			description:       "Simple root-level field should remain unchanged",
		},
		{
			name:      "nested_object_field",
			inputPath: "radius_config.acct_interim_interval",
			schemaFields: map[string]*FieldInfo{
				"radius_config": {
					Path:     "radius_config",
					AttrType: "nested",
				},
				"radius_config.acct_interim_interval": {
					Path:     "radius_config.acct_interim_interval",
					AttrType: "int64",
				},
			},
			mapAttributePaths: map[string]bool{},
			expected:          "radius_config.acct_interim_interval",
			description:       "Nested object field should remain unchanged",
		},
		{
			name:      "map_followed_by_list",
			inputPath: "acl_tags.management.specs.0.protocol",
			schemaFields: map[string]*FieldInfo{
				"acl_tags": {
					Path:     "acl_tags",
					AttrType: "map_nested",
				},
				"acl_tags.{key}.specs": {
					Path:     "acl_tags.{key}.specs",
					AttrType: "list_nested",
				},
				"acl_tags.{key}.specs.protocol": {
					Path:     "acl_tags.{key}.specs.protocol",
					AttrType: "string",
				},
			},
			mapAttributePaths: map[string]bool{
				"acl_tags": true,
			},
			expected:    "acl_tags.{key}.specs.protocol",
			description: "Map key followed by list index should normalize both",
		},
		{
			name:      "ipv6_address_in_map_key",
			inputPath: "extra_routes6.2001:db8::/32.via",
			schemaFields: map[string]*FieldInfo{
				"extra_routes6": {
					Path:     "extra_routes6",
					AttrType: "map_nested",
				},
				"extra_routes6.{key}.via": {
					Path:     "extra_routes6.{key}.via",
					AttrType: "string",
				},
			},
			mapAttributePaths: map[string]bool{
				"extra_routes6": true,
			},
			expected:    "extra_routes6.{key}.via",
			description: "IPv6 address with CIDR in map key should be replaced with {key}",
		},
		{
			name:      "map_attribute_with_string_key",
			inputPath: "vars.my_var",
			schemaFields: map[string]*FieldInfo{
				"vars": {
					Path:     "vars",
					AttrType: "map",
				},
			},
			mapAttributePaths: map[string]bool{
				"vars": true,
			},
			expected:    "vars.{key}",
			description: "schema.MapAttribute with string key should be replaced with {key}",
		},
		{
			name:      "map_attribute_with_numeric_key",
			inputPath: "env.123",
			schemaFields: map[string]*FieldInfo{
				"env": {
					Path:     "env",
					AttrType: "map",
				},
			},
			mapAttributePaths: map[string]bool{
				"env": true,
			},
			expected:    "env.{key}",
			description: "schema.MapAttribute with numeric key should be replaced with {key}",
		},
		{
			name:      "list_in_nested_object_in_list",
			inputPath: "switch_matching.rules.0.ip_config.network",
			schemaFields: map[string]*FieldInfo{
				"switch_matching": {
					Path:     "switch_matching",
					AttrType: "nested",
				},
				"switch_matching.rules": {
					Path:     "switch_matching.rules",
					AttrType: "list_nested",
				},
				"switch_matching.rules.ip_config": {
					Path:     "switch_matching.rules.ip_config",
					AttrType: "nested",
				},
				"switch_matching.rules.ip_config.network": {
					Path:     "switch_matching.rules.ip_config.network",
					AttrType: "string",
				},
			},
			mapAttributePaths: map[string]bool{},
			expected:          "switch_matching.rules.ip_config.network",
			description:       "List containing nested object containing field should remove array index",
		},
		{
			name:              "empty_path",
			inputPath:         "",
			schemaFields:      map[string]*FieldInfo{},
			mapAttributePaths: map[string]bool{},
			expected:          "",
			description:       "Empty path should return empty string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tracker := &FieldCoverageTracker{
				ResourceName:                "test_resource",
				SchemaFields:                tt.schemaFields,
				MapNormalizationPaths:       tt.mapAttributePaths,
				NestedMapNormalizationPaths: tt.mapAttributePaths,
				UnknownFields:               make(map[string]bool),
			}

			result := tracker.normalizeFieldPath(tt.inputPath)

			assert.Equal(t, tt.expected, result, tt.description)
		})
	}
}

func TestMarkFieldAsTested(t *testing.T) {
	tests := []struct {
		testPath     string
		expectedPath string
		shouldMark   bool
	}{
		{"name", "name", true},
		{"privileges.0.role", "privileges.role", true},
		{"networks.guest.vlan_id", "networks.{key}.vlan_id", true},
		{"vars.my_var", "vars.{key}", true},
		{"vars.123", "vars.{key}", true},
		{"nonexistent.field", "nonexistent.field", false},
	}

	for _, tt := range tests {
		t.Run(tt.testPath, func(t *testing.T) {
			// Create fresh tracker for each test to avoid state mutation
			tracker := &FieldCoverageTracker{
				ResourceName: "test_resource",
				SchemaFields: map[string]*FieldInfo{
					"name": {
						Path:     "name",
						IsTested: false,
					},
					"privileges": {
						Path:     "privileges",
						IsTested: false,
					},
					"privileges.role": {
						Path:     "privileges.role",
						IsTested: false,
					},
					"networks": {
						Path:     "networks",
						IsTested: false,
					},
					"networks.{key}.vlan_id": {
						Path:     "networks.{key}.vlan_id",
						IsTested: false,
					},
					"vars": {
						Path:     "vars",
						IsTested: false,
					},
					"vars.{key}": {
						Path:     "vars.{key}",
						IsTested: false,
					},
				},
				MapNormalizationPaths: map[string]bool{
					"vars": true,
				},
				NestedMapNormalizationPaths: map[string]bool{
					"networks": true,
				},
				UnknownFields:    make(map[string]bool),
				NormalizedFields: make(map[string]struct{}),
			}

			tracker.MarkFieldAsTested(tt.testPath)

			if !tt.shouldMark {
				// Verify the field doesn't exist in schema
				_, exists := tracker.SchemaFields[tt.expectedPath]
				assert.False(t, exists, "Field %q should not exist in schema", tt.expectedPath)
				return
			}

			field, exists := tracker.SchemaFields[tt.expectedPath]
			require.True(t, exists, "Expected field %q to exist in schema", tt.expectedPath)
			assert.True(t, field.IsTested, "Field %q should be marked as tested", tt.expectedPath)
		})
	}
}

func TestIsNumericOrPunctuation(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"0", true},
		{"123", true},
		{"0123456789", true},
		{"1.2.3.4", true},
		{"10.0.0.0/8", true},
		{"#", true},
		{"-5", true},
		{"", false},
		{"abc", false},
		{"12a", false},
		{"guest", false},
		{"mirror-to-analyzer", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := isNumericOrPunctuation(tt.input)
			assert.Equal(t, tt.expected, result, "IsNumericOrPunctuation(%q)", tt.input)
		})
	}
}

func TestExtractAllSchemaFields(t *testing.T) {
	// Schema covering all attribute types to test reflection-based extraction
	testSchema := map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true, // Computed-only field
		},
		"name": schema.StringAttribute{
			Required: true,
		},
		"enabled": schema.BoolAttribute{
			Optional: true,
		},
		"config": schema.SingleNestedAttribute{
			Optional: true,
			Attributes: map[string]schema.Attribute{
				"timeout": schema.Int64Attribute{
					Optional: true,
				},
				"deep": schema.SingleNestedAttribute{
					Optional: true,
					Attributes: map[string]schema.Attribute{
						"deep_field": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
		},
		"tags": schema.ListAttribute{
			Optional:    true,
			ElementType: types.StringType,
		},
		"vars": schema.MapAttribute{
			Optional:    true,
			ElementType: types.StringType,
		},
		"servers": schema.ListNestedAttribute{
			Optional: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"host": schema.StringAttribute{
						Required: true,
					},
					"port": schema.Int64Attribute{
						Optional: true,
					},
				},
			},
		},
		"unique_items": schema.SetNestedAttribute{
			Optional: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"item_id": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
		"metadata": schema.MapNestedAttribute{
			Optional: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"value": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
	}

	// Create tracker directly to bypass env var check
	tracker := NewFieldCoverageTracker("test_resource")
	tracker.extractFields("", testSchema)

	// Verify all nested attribute types are extracted
	expectedFields := map[string]struct {
		attrType string
		parent   string
		required bool
		optional bool
		computed bool
	}{
		"id":                     {attrType: "string", parent: "", required: false, optional: false, computed: true},
		"name":                   {attrType: "string", parent: "", required: true, optional: false, computed: false},
		"enabled":                {attrType: "bool", parent: "", required: false, optional: true, computed: false},
		"config":                 {attrType: "nested", parent: "", required: false, optional: true, computed: false},
		"config.timeout":         {attrType: "int64", parent: "config", required: false, optional: true, computed: false},
		"config.deep":            {attrType: "nested", parent: "config", required: false, optional: true, computed: false},
		"config.deep.deep_field": {attrType: "string", parent: "config.deep", required: true, optional: false, computed: false},
		"tags":                   {attrType: "list", parent: "", required: false, optional: true, computed: false},
		"vars":                   {attrType: "map", parent: "", required: false, optional: true, computed: false},
		"vars.{key}":             {attrType: "map_key", parent: "vars", required: false, optional: true, computed: false},
		"servers":                {attrType: "list_nested", parent: "", required: false, optional: true, computed: false},
		"servers.host":           {attrType: "string", parent: "servers", required: true, optional: false, computed: false},
		"servers.port":           {attrType: "int64", parent: "servers", required: false, optional: true, computed: false},
		"unique_items":           {attrType: "set_nested", parent: "", required: false, optional: true, computed: false},
		"unique_items.item_id":   {attrType: "string", parent: "unique_items", required: true, optional: false, computed: false},
		"metadata":               {attrType: "map_nested", parent: "", required: false, optional: true, computed: false},
		"metadata.{key}.value":   {attrType: "string", parent: "metadata.{key}", required: true, optional: false, computed: false},
	}

	assert.Len(t, tracker.SchemaFields, len(expectedFields), "Should extract all expected fields")

	// Verify each field and its metadata
	for path, expected := range expectedFields {
		field, exists := tracker.SchemaFields[path]
		require.True(t, exists, "Field %q should be extracted", path)
		assert.Equal(t, expected.attrType, field.AttrType, "Field %q should have correct type", path)
		assert.Equal(t, expected.parent, field.Parent, "Field %q should have correct parent", path)
		assert.Equal(t, expected.required, field.Required, "Field %q Required flag mismatch", path)
		assert.Equal(t, expected.optional, field.Optional, "Field %q Optional flag mismatch", path)
		assert.Equal(t, expected.computed, field.Computed, "Field %q Computed flag mismatch", path)
	}

	// Verify map attribute tracking
	assert.True(t, tracker.MapNormalizationPaths["vars"], "'vars' should be marked as MapAttribute path")
	assert.True(t, tracker.NestedMapNormalizationPaths["metadata"], "'metadata' should be marked as MapNestedAttribute path")
	assert.False(t, tracker.NestedMapNormalizationPaths["servers"], "'servers' should not be marked as map attribute path")
}
