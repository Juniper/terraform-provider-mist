// Package validators provides utilities for tracking test coverage of Terraform schema fields.
// See TESTING_GUIDE.md for detailed usage patterns.
package validators

import (
	"encoding/json"
	"os"
	"reflect"
	"sort"
	"strings"
	"testing"
	"unicode"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

const keyFieldPlaceholder = "{key}"

// FieldCoverageTracker tracks schema fields and their test coverage
type FieldCoverageTracker struct {
	ResourceName             string
	SchemaFields             map[string]*FieldInfo
	NestedMapAttributePaths  map[string]bool
	UnknownFields            map[string]bool     // Deduplicated test paths that don't match schema
	NormalizedFields         map[string]struct{} // Unique normalized field paths that were tested
	SchemaExtractionFailures []string            // Tracks paths where schema extraction failed via reflection
}

// FieldInfo contains metadata about a schema field
type FieldInfo struct {
	Path       string           // Field path in dot notation (e.g., "ldap_server_hosts", "auth.ldap.bind_dn")
	Field      string           // Field name only (e.g., "bind_dn")
	Parent     string           // Full parent path ("" for root, "auth.ldap" for nested)
	Required   bool             // Field is required
	Optional   bool             // Field is optional
	Computed   bool             // Field is computed (auto-populated by provider, intentionally excluded from coverage when Computed-only)
	AttrType   string           // Semantic type: "string", "bool", "int64", "float64", "list", "map", "nested"
	SchemaAttr schema.Attribute // The actual schema attribute for future inspection
	IsTested   bool             // Marked true when test validates this field
}

// NewFieldCoverageTracker creates a new tracker for the given resource
func NewFieldCoverageTracker(resourceName string) *FieldCoverageTracker {
	return &FieldCoverageTracker{
		ResourceName:             resourceName,
		NormalizedFields:         make(map[string]struct{}),
		SchemaFields:             make(map[string]*FieldInfo),
		NestedMapAttributePaths:  make(map[string]bool),
		UnknownFields:            make(map[string]bool),
		SchemaExtractionFailures: make([]string, 0),
	}
}

// FieldCoverageTrackerWithSchema creates a new tracker and extracts fields from the provided schema attributes
func FieldCoverageTrackerWithSchema(resourceName string, attributes map[string]schema.Attribute) *FieldCoverageTracker {
	if os.Getenv("DISABLE_MIST_FIELD_COVERAGE_TRACKER") != "" {
		return nil
	}

	tracker := NewFieldCoverageTracker(resourceName)
	tracker.extractFields("", attributes)
	return tracker
}

// MarkFieldAsTested normalizes the field path and marks it as tested.
func (t *FieldCoverageTracker) MarkFieldAsTested(fieldPath string) {
	normalized := t.normalizeFieldPath(fieldPath)
	field, exists := t.SchemaFields[normalized]
	if exists {
		field.IsTested = true
	}
	t.NormalizedFields[normalized] = struct{}{}
}

// normalizeFieldPath converts test paths to schema paths, using dot notation.
// Uses schema knowledge to distinguish between indices, map keys, and field names.
func (t *FieldCoverageTracker) normalizeFieldPath(fieldPath string) string {
	parts := strings.Split(fieldPath, ".")
	normalized := make([]string, 0, len(parts))

	for i := 0; i < len(parts); i++ {
		part := parts[i]
		parentPath := strings.Join(normalized, ".")

		// Skip array indices (#, or pure digits/punctuation when not in map context)
		if part == "#" || (isNumericOrPunctuation(part) && !t.NestedMapAttributePaths[parentPath]) {
			continue
		}

		// Check if this completes a known schema path
		testPath := parentPath
		if testPath != "" {
			testPath += "."
		}
		testPath += part

		_, exists := t.SchemaFields[testPath]
		if exists {
			normalized = append(normalized, part)
			continue
		}

		// Replace with {key} if parent is a map
		if t.NestedMapAttributePaths[parentPath] {
			normalized = append(normalized, keyFieldPlaceholder)
			continue
		}

		// Unknown field - keep as-is
		normalized = append(normalized, part)
		t.UnknownFields[fieldPath] = true
	}

	return strings.Join(normalized, ".")
}

// isNumericOrPunctuation checks if a string contains only numeric digits and punctuation.
func isNumericOrPunctuation(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, ch := range s {
		if !unicode.IsDigit(ch) && !unicode.IsPunct(ch) {
			return false
		}
	}
	return true
}

// extractFields recursively walks the schema tree and extracts field metadata.
func (t *FieldCoverageTracker) extractFields(path string, attributes map[string]schema.Attribute) {
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
			SchemaAttr: attr,
		}

		// Extract metadata using reflection
		extractFieldMetadata(attr, fieldInfo)

		// Store in map using path as key
		t.SchemaFields[currentPath] = fieldInfo

		// Handle nested attributes recursively
		switch v := attr.(type) {
		case schema.SingleNestedAttribute:
			nestedAttrs := getNestedAttributes(v)
			if nestedAttrs == nil {
				t.SchemaExtractionFailures = append(t.SchemaExtractionFailures, currentPath+" (SingleNestedAttribute)")
				break
			}
			t.extractFields(currentPath, nestedAttrs)
		case schema.ListNestedAttribute:
			nestedAttrs := getListNestedAttributes(v)
			if nestedAttrs == nil {
				t.SchemaExtractionFailures = append(t.SchemaExtractionFailures, currentPath+" (ListNestedAttribute)")
				break
			}
			t.extractFields(currentPath, nestedAttrs)
		case schema.SetNestedAttribute:
			nestedAttrs := getSetNestedAttributes(v)
			if nestedAttrs == nil {
				t.SchemaExtractionFailures = append(t.SchemaExtractionFailures, currentPath+" (SetNestedAttribute)")
				break
			}
			t.extractFields(currentPath, nestedAttrs)
		case schema.MapNestedAttribute:
			nestedAttrs := getMapNestedAttributes(v)
			if nestedAttrs == nil {
				t.SchemaExtractionFailures = append(t.SchemaExtractionFailures, currentPath+" (MapNestedAttribute)")
				break
			}
			// Map uses {key} notation in path
			t.NestedMapAttributePaths[currentPath] = true
			keyPath := currentPath + "." + keyFieldPlaceholder
			t.extractFields(keyPath, nestedAttrs)
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
	requiredField := v.FieldByName("Required")
	if requiredField.IsValid() && requiredField.Kind() == reflect.Bool {
		fieldInfo.Required = requiredField.Bool()
	}
	optionalField := v.FieldByName("Optional")
	if optionalField.IsValid() && optionalField.Kind() == reflect.Bool {
		fieldInfo.Optional = optionalField.Bool()
	}
	computedField := v.FieldByName("Computed")
	if computedField.IsValid() && computedField.Kind() == reflect.Bool {
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
	field := v.FieldByName("Attributes")
	if field.IsValid() && field.CanInterface() {
		attrs, ok := field.Interface().(map[string]schema.Attribute)
		if ok {
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

// getSetNestedAttributes extracts attributes from a SetNestedAttribute using reflection
func getSetNestedAttributes(attr schema.SetNestedAttribute) map[string]schema.Attribute {
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

// FieldCoverageReport writes the current state of the FieldCoverageTracker to Stdout.
func (tracker *FieldCoverageTracker) FieldCoverageReport(t testing.TB) {
	t.Helper()

	type CoverageReport struct {
		ResourceName                string   `json:"resource_name"`
		TestedFieldsCnt             int      `json:"tested_fields_count"`
		UntestedFieldsCnt           int      `json:"untested_fields_count"`
		UntestedFields              []string `json:"untested_fields"`
		UnknownFieldsCnt            int      `json:"unknown_fields_count"`
		UnknownFields               []string `json:"unknown_fields"`
		SchemaExtractionFailuresCnt int      `json:"schema_extraction_failures_count"`
		SchemaExtractionFailures    []string `json:"schema_extraction_failures"`
	}

	// Build report
	untestedFields := make([]string, 0)
	for path, field := range tracker.SchemaFields {
		if !field.IsTested && isComputedOnlyField(field) && !isContainerType(field.SchemaAttr) {
			untestedFields = append(untestedFields, path)
		}
	}

	// Convert unknown fields map to sorted slice
	unknownFields := make([]string, 0, len(tracker.UnknownFields))
	for path := range tracker.UnknownFields {
		unknownFields = append(unknownFields, path)
	}

	sort.Strings(unknownFields)
	sort.Strings(untestedFields)

	// Capture test execution status from testing.TB
	report := CoverageReport{
		ResourceName:                tracker.ResourceName,
		TestedFieldsCnt:             len(tracker.NormalizedFields),
		UntestedFieldsCnt:           len(untestedFields),
		UntestedFields:              untestedFields,
		UnknownFieldsCnt:            len(unknownFields),
		UnknownFields:               unknownFields,
		SchemaExtractionFailuresCnt: len(tracker.SchemaExtractionFailures),
		SchemaExtractionFailures:    tracker.SchemaExtractionFailures,
	}

	// Write JSON report to Stdout
	err := writeToStdout(report)
	if err != nil {
		t.Errorf("failed to write field coverage report: %v", err)
	}
}

// isContainerType checks if an attribute is a container type.
// Container types cannot be tested by themselves and are thus excluded from test coverage counts.
func isContainerType(attr schema.Attribute) bool {
	_, isSingleNested := attr.(schema.SingleNestedAttribute)
	_, isMapNested := attr.(schema.MapNestedAttribute)
	return isSingleNested || isMapNested
}

// isComputedOnlyField checks if a field is Computed-only (Computed=true, Optional=false).
func isComputedOnlyField(field *FieldInfo) bool {
	return field.Computed && !field.Optional
}

// writeToStdout writes a JSON report to Stdout
func writeToStdout(data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	jsonData = append(jsonData, '\n')
	_, err = os.Stdout.Write(jsonData)
	return err
}
