package validators

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"unicode"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// FieldCoverageTracker tracks schema fields and their test coverage
type FieldCoverageTracker struct {
	ResourceName      string
	SchemaFields      map[string]*FieldInfo
	MapAttributePaths map[string]bool
}

// FieldInfo contains metadata about a schema field
type FieldInfo struct {
	Path       string           // Field path in dot notation (e.g., "ldap_server_hosts", "auth.ldap.bind_dn")
	Field      string           // Field name only (e.g., "bind_dn")
	Parent     string           // Full parent path ("" for root, "auth.ldap" for nested)
	HasKey     bool             // Indicates if field has a {key}.
	Required   bool             // Field is required
	Optional   bool             // Field is optional
	Computed   bool             // Field is computed
	AttrType   string           // Semantic type: "string", "bool", "int64", "float64", "list", "map", "nested"
	SchemaAttr schema.Attribute // The actual schema attribute for future inspection
	IsTested   bool             // Marked true when test validates this field
}

// NewFieldCoverageTracker creates a new tracker for the given resource
func NewFieldCoverageTracker(resourceName string) *FieldCoverageTracker {
	return &FieldCoverageTracker{
		ResourceName:      resourceName,
		SchemaFields:      make(map[string]*FieldInfo),
		MapAttributePaths: make(map[string]bool),
	}
}

// ExtractAllSchemaFields extracts all field paths from a Terraform schema
// and returns a populated FieldCoverageTracker
func ExtractAllSchemaFields(resourceName string, schemaAttrs map[string]schema.Attribute) *FieldCoverageTracker {
	tracker := NewFieldCoverageTracker(resourceName)
	tracker.extractFields("", false, schemaAttrs)
	return tracker
}

// MarkFieldAsTested marks a field as tested, normalizing the field path to remove array indices
func (t *FieldCoverageTracker) MarkFieldAsTested(fieldPath string) {
	normalized := t.normalizeFieldPath(fieldPath)
	if field, exists := t.SchemaFields[normalized]; exists {
		field.IsTested = true
	}

	// Write normalized field to debug file
	filename := fmt.Sprintf("%s_normalized_fields.txt", t.ResourceName)
	f, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if f != nil {
		fmt.Fprintf(f, "%s\n", normalized)
		f.Close()
	}
}

// normalizeFieldPath removes array indices and uses schema knowledge to replace map keys with {key}
// Examples:
//   - "privileges.0.role" -> "privileges.role"
//   - "switch_mgmt.local_accounts.readonly.password" -> "switch_mgmt.local_accounts.{key}.password"
//   - "extra_routes.10.0.0.0/8.via" -> "extra_routes.{key}.via"
//   - "networks.guest.vlan_id" -> "networks.{key}.vlan_id"
func (t *FieldCoverageTracker) normalizeFieldPath(fieldPath string) string {
	fmt.Println("Starting normilisation process for field path:", fieldPath)

	parts := strings.Split(fieldPath, ".")
	normalized := make([]string, 0, len(parts))

	for i, part := range parts {
		fmt.Printf("Part %d: %s\n", i, part)

		// Build the parent path to check context
		path := strings.Join(normalized, ".")
		if part == "#" || isAllDigits(part) && !t.MapAttributePaths[path] { // Skip array indices (pure numbers or #), but NOT if we're in a map context
			fmt.Printf("Part is an array index, skipping: %s\n", part)
			continue
		}

		// Check if full path exists in schema
		pathCheck := strings.Join(append(normalized, part), ".")
		fmt.Printf("Path check: %s\n", pathCheck)
		if _, exists := t.SchemaFields[pathCheck]; exists {
			normalized = append(normalized, part)
			continue
		}

		// Check if path with {key} exists in schema (for map attributes)
		fmt.Printf("Key path check: %s\n", path)
		if t.MapAttributePaths[path] {
			normalized = append(normalized, "{key}")
			continue
		}

		// Default: keep the part as is (could be a new field not in schema)
		normalized = append(normalized, part)
		fmt.Printf("ERROR: %s\n", part)
	}

	return strings.Join(normalized, ".")
}

// isAllDigits checks if a string contains only numeric digits
func isAllDigits(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, ch := range s {
		if !unicode.IsDigit(ch) {
			return false
		}
	}
	return true
}

// extractFields recursively extracts all fields from schema attributes with metadata
func (t *FieldCoverageTracker) extractFields(path string, hasKey bool, attributes map[string]schema.Attribute) {
	for name, attr := range attributes {
		currentPath := name
		if path != "" {
			currentPath = path + "." + name
		}

		// Create FieldInfo with metadata
		fieldInfo := &FieldInfo{
			Path:       currentPath,
			Field:      name,
			Parent:     path,
			HasKey:     hasKey,
			SchemaAttr: attr,
		}

		// Extract metadata using reflection
		extractFieldMetadata(attr, fieldInfo)

		// Store in map using path as key
		t.SchemaFields[currentPath] = fieldInfo

		// Handle nested attributes recursively
		switch v := attr.(type) {
		case schema.SingleNestedAttribute:
			if nestedAttrs := getNestedAttributes(v); nestedAttrs != nil {
				t.extractFields(currentPath, false, nestedAttrs)
			}
		case schema.ListNestedAttribute:
			if nestedAttrs := getListNestedAttributes(v); nestedAttrs != nil {
				t.extractFields(currentPath, false, nestedAttrs)
			}
		case schema.MapNestedAttribute:
			if nestedAttrs := getMapNestedAttributes(v); nestedAttrs != nil {
				// Map uses {key} notation in path
				t.MapAttributePaths[currentPath] = true
				keyPath := currentPath + ".{key}"
				t.extractFields(keyPath, true, nestedAttrs)
			}
		}
	}
}

// extractFieldMetadata extracts Required, Optional, Computed, and AttrType from schema attribute
func extractFieldMetadata(attr schema.Attribute, fieldInfo *FieldInfo) {
	v := reflect.ValueOf(attr)
	if !v.IsValid() {
		return
	}

	// Extract Required, Optional, Computed using reflection
	if requiredField := v.FieldByName("Required"); requiredField.IsValid() && requiredField.Kind() == reflect.Bool {
		fieldInfo.Required = requiredField.Bool()
	}
	if optionalField := v.FieldByName("Optional"); optionalField.IsValid() && optionalField.Kind() == reflect.Bool {
		fieldInfo.Optional = optionalField.Bool()
	}
	if computedField := v.FieldByName("Computed"); computedField.IsValid() && computedField.Kind() == reflect.Bool {
		fieldInfo.Computed = computedField.Bool()
	}

	// Determine semantic type from attribute type
	fieldInfo.AttrType = getSemanticType(attr)
}

// getSemanticType returns the semantic type string for an attribute
func getSemanticType(attr schema.Attribute) string {
	switch attr.(type) {
	case schema.StringAttribute:
		return "string"
	case schema.BoolAttribute:
		return "bool"
	case schema.Int64Attribute:
		return "int64"
	case schema.Float64Attribute:
		return "float64"
	case schema.NumberAttribute:
		return "number"
	case schema.ListAttribute:
		return "list"
	case schema.SetAttribute:
		return "set"
	case schema.MapAttribute:
		return "map"
	case schema.SingleNestedAttribute:
		return "nested"
	case schema.ListNestedAttribute:
		return "list_nested"
	case schema.SetNestedAttribute:
		return "set_nested"
	case schema.MapNestedAttribute:
		return "map_nested"
	default:
		return "unknown"
	}
}

// getNestedAttributes extracts attributes from a SingleNestedAttribute using reflection
func getNestedAttributes(attr schema.SingleNestedAttribute) map[string]schema.Attribute {
	v := reflect.ValueOf(attr)
	if !v.IsValid() {
		return nil
	}

	// Look for Attributes field directly on the SingleNestedAttribute
	if field := v.FieldByName("Attributes"); field.IsValid() && field.CanInterface() {
		if attrs, ok := field.Interface().(map[string]schema.Attribute); ok {
			return attrs
		}
	}

	return nil
}

// getListNestedAttributes extracts attributes from a ListNestedAttribute using reflection
func getListNestedAttributes(attr schema.ListNestedAttribute) map[string]schema.Attribute {
	v := reflect.ValueOf(attr)
	if !v.IsValid() {
		return nil
	}

	// Look for NestedObject field first
	if nestedObjField := v.FieldByName("NestedObject"); nestedObjField.IsValid() && nestedObjField.CanInterface() {
		nestedObj := nestedObjField.Interface()

		// Get the nested object and look for its Attributes
		nestedV := reflect.ValueOf(nestedObj)
		if nestedV.IsValid() && nestedV.Kind() == reflect.Struct {
			if attributesField := nestedV.FieldByName("Attributes"); attributesField.IsValid() && attributesField.CanInterface() {
				if attrs, ok := attributesField.Interface().(map[string]schema.Attribute); ok {
					return attrs
				}
			}
		}
	}

	return nil
}

// getMapNestedAttributes extracts attributes from a MapNestedAttribute using reflection
func getMapNestedAttributes(attr schema.MapNestedAttribute) map[string]schema.Attribute {
	v := reflect.ValueOf(attr)
	if !v.IsValid() {
		return nil
	}

	// Look for NestedObject field first
	if nestedObjField := v.FieldByName("NestedObject"); nestedObjField.IsValid() && nestedObjField.CanInterface() {
		nestedObj := nestedObjField.Interface()

		// Get the nested object and look for its Attributes
		nestedV := reflect.ValueOf(nestedObj)
		if nestedV.IsValid() && nestedV.Kind() == reflect.Struct {
			if attributesField := nestedV.FieldByName("Attributes"); attributesField.IsValid() && attributesField.CanInterface() {
				if attrs, ok := attributesField.Interface().(map[string]schema.Attribute); ok {
					return attrs
				}
			}
		}
	}

	return nil
}
