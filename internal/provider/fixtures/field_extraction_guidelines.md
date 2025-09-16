# Terraform Provider Field Finder

## Quick Implementation Guide

### 1. Files to Create

- [ ] `{resource}_test_structs.go` - Test struct definitions with HCL tags
- [ ] `{resource}_resource_test.go` - Main test implementation
- [ ] `fixtu## Expected Field Count Discrepancies

### Total Validations = Schema Fields + Array Index/Map Key Validations

**Common Scenarios:**

- **More validations than schema fields**: Tests validate array indices (`portal.amazon_email_domains.0`) and map keys that aren't direct schema fields
- **Fewer validations than schema fields**: Missing field coverage - use analysis to identify gaps

### 2. Implementation Pattern (Follow org_wlantemplate_resource_test.go)

- [ ] **Dual test case strategy**: simple_case + fixture_case_N
- [ ] **Automated HCL generation** with `gohcl.EncodeIntoBody()`
- [ ] **Use `GetTestOrgId()`** - never hardcode org IDs
- [ ] **Comprehensive fixture data** with boolean fields set to `true`

### 3. Test Struct Guidelines

- **Found in file {resource}_test_structs.go**

- [ ] **HCL tags on fields with CTY tags OR no tags** (see org_wlan_test_structs.go)
- [ ] **Pointer types for optional fields** (`*string`, `*bool`, `*int64`)
- [ ] **Concrete types for required fields** (`string`, `[]string`)
- **Note that this is a generated file and you should not add any fields to structs already present in the file.**
- **The above mentioned file generator is not 100% though and it misses some nested structs. If, and only if, a whole entire nested struct is missing, may you add it to the file.**

### 4. Fixture Data Best Practices

- [ ] **Comprehensive field coverage** (aim for 100%)
- [ ] **Set boolean fields to `true`** for maximum test coverage
- [ ] **Include optional fields** with realistic test values
- [ ] **Use `␞` separator** for multiple fixtures

## Field Coverage Verification (MANDATORY)

### Step 1: Extract Complete Schema Fields (Automated Method)

Use reflection-based extraction for 100% accuracy with nested fields:

```go
// Add temporary helper functions to your *_resource_test.go file
func TestExtractSchema(t *testing.T) {
    ctx := context.Background()
    
    // Get the schema (replace with your resource schema function)
    schemaObj := resource_{resource_name}.{ResourceName}ResourceSchema(ctx)
    
    // Extract all field paths with proper dot notation
    allFields := extractAllFieldPaths("", schemaObj.Attributes)
    sort.Strings(allFields)
    
    // Write to file
    file, _ := os.Create("all_schema_fields.txt")
    defer file.Close()
    
    for _, field := range allFields {
        fmt.Fprintln(file, field)
    }
    
    t.Logf("Generated all_schema_fields.txt with %d fields", len(allFields))
}

func extractAllFieldPaths(prefix string, attributes map[string]schema.Attribute) []string {
    var fields []string
    
    for name, attr := range attributes {
        currentPath := name
        if prefix != "" {
            currentPath = prefix + "." + name
        }
        
        fields = append(fields, currentPath)
        
        // Handle nested attributes
        switch v := attr.(type) {
        case schema.SingleNestedAttribute:
            if nestedAttrs := getNestedAttributes(v); nestedAttrs != nil {
                nestedFields := extractAllFieldPaths(currentPath, nestedAttrs)
                fields = append(fields, nestedFields...)
            }
        case schema.ListNestedAttribute:
            if nestedAttrs := getListNestedAttributes(v); nestedAttrs != nil {
                nestedFields := extractAllFieldPaths(currentPath, nestedAttrs)
                fields = append(fields, nestedFields...)
            }
        case schema.MapNestedAttribute:
            if nestedAttrs := getMapNestedAttributes(v); nestedAttrs != nil {
                mapPath := currentPath + ".{key}"
                nestedFields := extractAllFieldPaths(mapPath, nestedAttrs)
                fields = append(fields, nestedFields...)
            }
        }
    }
    
    return fields
}

// Helper functions for reflection-based nested attribute extraction
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
```

**Required Imports:**

```go
import (
    "context"
    "fmt"
    "os"
    "reflect"
    "sort"
    "testing"
    
    "github.com/hashicorp/terraform-plugin-framework/resource/schema"
    resource_{resource_name} "github.com/Juniper/terraform-provider-mist/internal/resource_{resource_name}"
)
```

**Run the extraction:**

```bash
# This generates all_schema_fields.txt with complete field list
go test -v -run "TestExtractSchema" ./internal/provider/ >/dev/null 2>&1
```

**Example Results:**

- **100% accuracy**: Captures all nested fields including deeply nested structures like `snmp_config.v3_config.vacm.access.prefix_list.*`
- **Proper dot notation**: All fields in testable format with infinite nesting depth support
- **Complete reflection**: Handles SingleNestedAttribute, ListNestedAttribute, and MapNestedAttribute recursively

### Step 2: Extract Currently Tested Fields

```bash
# Extract field paths from test execution output
# Include both TestCheckResourceAttr AND TestCheckResourceAttrSet
# Exclude map length validations (.% fields)
go test -v -run "Test{ResourceName}Model" ./internal/provider/ 2>&1 | \
  grep "TestCheckResourceAttr.*{resource_name}\." | \
  sed -n 's/.*TestCheckResourceAttr[^(]*([^,]*, "\([^"]*\)".*/\1/p' | \
  grep -v '\.%$' | sort -u > current_tested_fields.txt

# Example for org_wlan:
# go test -v -run "TestOrgWlanModel" ./internal/provider/ 2>&1 | \
#   grep "TestCheckResourceAttr.*org_wlan\." | \
#   sed -n 's/.*TestCheckResourceAttr[^(]*([^,]*, "\([^"]*\)".*/\1/p' | \
#   grep -v '\.%$' | sort -u > current_tested_fields.txt
```
