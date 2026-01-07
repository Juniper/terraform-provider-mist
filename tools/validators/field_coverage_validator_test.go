package validators

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
			name:      "map_with_numeric_key",
			inputPath: "switch_mgmt.local_accounts.readonly.password",
			schemaFields: map[string]*FieldInfo{
				"switch_mgmt": {
					Path:     "switch_mgmt",
					AttrType: "nested",
				},
				"switch_mgmt.local_accounts": {
					Path:     "switch_mgmt.local_accounts",
					AttrType: "map_nested",
				},
				"switch_mgmt.local_accounts.{key}.password": {
					Path:     "switch_mgmt.local_accounts.{key}.password",
					AttrType: "string",
				},
			},
			mapAttributePaths: map[string]bool{
				"switch_mgmt.local_accounts": true,
			},
			expected:    "switch_mgmt.local_accounts.{key}.password",
			description: "String key 'readonly' in map should be replaced with {key}",
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
			name:      "special_characters_in_map_key",
			inputPath: "port_mirroring.mirror-to-analyzer.input_port_ids_networks.0",
			schemaFields: map[string]*FieldInfo{
				"port_mirroring": {
					Path:     "port_mirroring",
					AttrType: "map_nested",
				},
				"port_mirroring.{key}.input_port_ids_networks": {
					Path:     "port_mirroring.{key}.input_port_ids_networks",
					AttrType: "list",
				},
			},
			mapAttributePaths: map[string]bool{
				"port_mirroring": true,
			},
			expected:    "port_mirroring.{key}.input_port_ids_networks",
			description: "Map key with hyphens followed by list should normalize both",
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
				ResourceName:            "test_resource",
				SchemaFields:            tt.schemaFields,
				NestedMapAttributePaths: tt.mapAttributePaths,
			}

			result := tracker.normalizeFieldPath(tt.inputPath)

			if result != tt.expected {
				t.Errorf("%s\nInput:    %q\nExpected: %q\nGot:      %q",
					tt.description, tt.inputPath, tt.expected, result)
			}
		})
	}
}

func TestMarkFieldAsTested(t *testing.T) {
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
		},
		NestedMapAttributePaths: map[string]bool{
			"networks": true,
		},
	}

	tests := []struct {
		testPath     string
		expectedPath string
		shouldMark   bool
	}{
		{"name", "name", true},
		{"privileges.0.role", "privileges.role", true},
		{"networks.guest.vlan_id", "networks.{key}.vlan_id", true},
		{"nonexistent.field", "nonexistent.field", false},
	}

	for _, tt := range tests {
		t.Run(tt.testPath, func(t *testing.T) {
			tracker.MarkFieldAsTested(tt.testPath)

			if field, exists := tracker.SchemaFields[tt.expectedPath]; exists {
				if tt.shouldMark && !field.IsTested {
					t.Errorf("Field %q should be marked as tested but wasn't", tt.expectedPath)
				}
			} else if tt.shouldMark {
				t.Errorf("Expected field %q to exist in schema", tt.expectedPath)
			}
		})
	}
}

func TestNonAlphabetCharacters(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"0", true},
		{"123", true},
		{"0123456789", true},
		{"", false},
		{"abc", false},
		{"12a", false},
		{"1.2", false},
		{"10.0.0.0", false},
		{"-5", false},
		{"#", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := nonAlphabetCharacters(tt.input)
			if result != tt.expected {
				t.Errorf("isAllDigits(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestExtractAllSchemaFields(t *testing.T) {
	// Create a simple schema for testing
	testSchema := map[string]schema.Attribute{
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
			},
		},
		"tags": schema.ListAttribute{
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

	tracker := ExtractAllSchemaFields("test_resource", testSchema)

	expectedFields := []string{
		"name",
		"enabled",
		"config",
		"config.timeout",
		"tags",
		"servers",
		"servers.host",
		"servers.port",
		"metadata",
		"metadata.{key}.value",
	}

	if len(tracker.SchemaFields) != len(expectedFields) {
		t.Errorf("Expected %d fields, got %d", len(expectedFields), len(tracker.SchemaFields))
	}

	for _, expectedPath := range expectedFields {
		if _, exists := tracker.SchemaFields[expectedPath]; !exists {
			t.Errorf("Expected field %q not found in extracted schema", expectedPath)
		}
	}

	// Verify MapAttributePaths
	if !tracker.NestedMapAttributePaths["metadata"] {
		t.Error("Expected 'metadata' to be marked as map attribute path")
	}

	// Verify field metadata
	if field, exists := tracker.SchemaFields["name"]; exists {
		if !field.Required {
			t.Error("Field 'name' should be marked as required")
		}
		if field.AttrType != "string" {
			t.Errorf("Field 'name' should have type 'string', got %q", field.AttrType)
		}
	}

	if field, exists := tracker.SchemaFields["servers.host"]; exists {
		if field.Parent != "servers" {
			t.Errorf("Field 'servers.host' should have parent 'servers', got %q", field.Parent)
		}
	}

	if field, exists := tracker.SchemaFields["metadata.{key}.value"]; exists {
		if !field.HasKey {
			t.Error("Field 'metadata.{key}.value' should have HasKey=true")
		}
	}
}

func TestGetSemanticType(t *testing.T) {
	tests := []struct {
		name     string
		attr     schema.Attribute
		expected string
	}{
		{"string", schema.StringAttribute{}, "string"},
		{"bool", schema.BoolAttribute{}, "bool"},
		{"int64", schema.Int64Attribute{}, "int64"},
		{"float64", schema.Float64Attribute{}, "float64"},
		{"number", schema.NumberAttribute{}, "number"},
		{"list", schema.ListAttribute{ElementType: types.StringType}, "list"},
		{"set", schema.SetAttribute{ElementType: types.StringType}, "set"},
		{"map", schema.MapAttribute{ElementType: types.StringType}, "map"},
		{"nested", schema.SingleNestedAttribute{Attributes: map[string]schema.Attribute{}}, "nested"},
		{"list_nested", schema.ListNestedAttribute{NestedObject: schema.NestedAttributeObject{}}, "list_nested"},
		{"set_nested", schema.SetNestedAttribute{NestedObject: schema.NestedAttributeObject{}}, "set_nested"},
		{"map_nested", schema.MapNestedAttribute{NestedObject: schema.NestedAttributeObject{}}, "map_nested"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getSemanticType(tt.attr)
			if result != tt.expected {
				t.Errorf("getSemanticType(%T) = %q, want %q", tt.attr, result, tt.expected)
			}
		})
	}
}
