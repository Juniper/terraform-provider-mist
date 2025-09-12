# Terraform Provider Testing Checklist

## Quick Implementation Guide

### 1. Files to Create

- [ ] `{resource}_test_structs.go` - Test struct definitions with HCL tags
- [ ] `{resource}_resource_test.go` - Main test implementation
- [ ] `fixtu## Expected Field Count Discrepancies

### Total Validations = Schema Fields + Array Index/Map Key Validations

**Common Scenarios:**

- **More validations than schema fields**: Tests validate array indices (`portal.amazon_email_domains.0`) and map keys that aren't direct schema fields
- **Fewer validations than schema fields**: Missing field coverage - use analysis to identify gaps

**Example Analysis:**

- org_wlan: 283 schema fields, 313 validations = 30 array/map validations + 100% coverage ✅  
- Missing coverage: 283 schema fields, 166 validations = 117 missing fields (59% coverage) ❌

**Key Insight:** Focus on achieving **zero missing fields** rather than exact count matching.resource/{resource}_config.tf` - Test data

### 2. Implementation Pattern (Follow org_wlantemplate_resource_test.go)

- [ ] **Dual test case strategy**: simple_case + fixture_case_N
- [ ] **Automated HCL generation** with `gohcl.EncodeIntoBody()`
- [ ] **Use `GetTestOrgId()`** - never hardcode org IDs
- [ ] **Comprehensive fixture data** with boolean fields set to `true`

### 3. Test Struct Guidelines

- [ ] **HCL tags on fields with CTY tags OR no tags** (see org_wlan_test_structs.go)
- [ ] **Pointer types for optional fields** (`*string`, `*bool`, `*int64`)
- [ ] **Concrete types for required fields** (`string`, `[]string`)

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

- Simple resources: ~10-50 fields  
- Complex nested resources: 200-400+ fields (e.g., org_wlan: 283 fields, org_networktemplate: 375 fields)
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

### Step 3: Find Missing Fields (Corrected Analysis)

```bash
# IMPORTANT: Account for proper list/map container testing
# Containers are properly tested via .# (length) and indexed access (.0, .1, etc.)

# 1. Extract containers that are properly tested via length/indexing
grep -E '\.(#|[0-9]+)$' current_tested_fields.txt | \
  sed 's/\.[#0-9].*$//' | sort -u > properly_tested_containers.txt

# 2. Combine current tested fields with properly tested containers
(cat current_tested_fields.txt; cat properly_tested_containers.txt) | \
  sort -u > corrected_tested_fields.txt

# 3. Find truly missing fields
comm -23 all_schema_fields.txt corrected_tested_fields.txt > truly_missing_fields.txt

# 4. View corrected analysis
echo "CORRECTED FIELD COVERAGE ANALYSIS:"
echo "==================================="
echo "Total schema fields: $(wc -l < all_schema_fields.txt)"
echo "Properly tested fields: $(wc -l < corrected_tested_fields.txt)" 
echo "Actually missing fields: $(wc -l < truly_missing_fields.txt)"
echo ""
echo "Actually missing fields:"
cat truly_missing_fields.txt
```

**Key Insights:**

- **Container objects** (like `coa_servers`, `auth_servers`) are **properly tested** via `.#` and indexed access
- **This correction typically reduces missing fields by 20-30%**
- **Focus on the `truly_missing_fields.txt` output** for actual gaps
- **Examples of proper container testing**:
  - `coa_servers` ✅ via `coa_servers.#` + `coa_servers.0.ip`
  - `auth_servers` ✅ via `auth_servers.#` + `auth_servers.0.host`

### Step 4: Achieve 100% Coverage

- [ ] **Zero missing fields required** - `truly_missing_fields.txt` must be empty
- [ ] **Add missing fields** to test struct and fixture data systematically
- [ ] **Add validation checks** for all missing fields using `TestCheckResourceAttr`
- [ ] **Re-run corrected analysis** until `truly_missing_fields.txt` is empty
- [ ] **Clean up temporary files**: `rm all_schema_fields.txt current_tested_fields.txt truly_missing_fields.txt corrected_tested_fields.txt properly_tested_containers.txt`
- [ ] **Remove helper functions** from test file after verification complete

**Systematic Addition Process:**

1. **Group missing fields by category** (auth.*, portal.*, etc.)
2. **Add fields to test struct** with appropriate Go types and HCL tags
3. **Add fields to fixture data** with realistic test values  
4. **Add TestCheckResourceAttr validations** for each field
5. **Verify with corrected re-analysis** until coverage complete

**Example org_wlan Results:**

- Initial (naive): 283 schema fields, 166 tested = 117 missing (59% coverage)
- Corrected: 283 schema fields, 190 truly tested = 93 missing (67% coverage)
- Target: 283 schema fields, 283 tested = 0 missing (100% coverage)

### Field Type Validation Rules

- [ ] **Computed-only fields** (like `id`): Use `TestCheckResourceAttrSet`
- [ ] **All other fields**: Use `TestCheckResourceAttr` with expected values
- [ ] **Never validate parent nested objects** - test child attributes only

### Manual Test Quality Verification

- [ ] **Boolean field values**: Search test output for `"false"` - should be minimal

```bash
go test -v -run "Test{ResourceName}Model" ./internal/provider/ 2>&1 | grep 'TestCheckResourceAttr.*"false"'
```

- [ ] **Null value checks**: Search for `= null` in fixture data - boolean fields should use `true`

```bash
grep "= null" fixtures/{resource_name}_resource/{resource_name}_config.tf
```

- [ ] **Missing validation checks**: Count validations vs expected fields

```bash
go test -v -run "Test{ResourceName}Model" ./internal/provider/ 2>&1 | grep -c "TestCheckResourceAttr.*{resource_name}\."
```

## Expected Field Count Discrepancies

### Understanding Validation Counts

Total Validations = Schema Fields + Array Index/Map Key Validations

**Common Scenarios:**

- **More validations than schema fields**: Tests validate array indices (`portal.amazon_email_domains.0`) and map keys that aren't direct schema fields
- **Fewer validations than schema fields**: Missing field coverage - use corrected analysis to identify gaps

**Corrected Analysis Results:**

- org_wlan: 283 schema fields, 339 validations = 56 array/map validations + 100% coverage ✅  
- Missing coverage: 283 schema fields, 190 validations = 93 missing fields (67% coverage) ❌

**Key Insight:** Use **corrected analysis** that accounts for proper container testing via `.#` and indexed access.

## Field Implementation Strategy

### One-by-One Field Addition (Recommended for Complex Resources)

When working with complex resources with many missing fields, implement a systematic one-by-one approach:

**Benefits:**

- **Discover schema validation constraints** (e.g., `anticlog_threshold` must be 16-32)
- **Identify field dependencies** (e.g., `dynamic_vlan` requires `auth.enable_mac_auth=true`)
- **Catch provider bugs early** (e.g., memory address issues in type conversion)
- **Ensure proper test validation** for each field before moving to next

**Process:**

1. **Select next field** from `truly_missing_fields.txt`
2. **Add to fixture** with reasonable test value
3. **Add validation** to test file
4. **Run test** and handle any validation errors
5. **Update missing fields list** when successful
6. **Repeat** until zero missing fields

**Example Field Dependencies Discovered:**

- `dynamic_vlan.*` fields require `vlan_enabled=true` AND (`auth.enable_mac_auth=true` OR `auth.type="eap"`)
- `auth.anticlog_threshold` has validation range 16-32 (SAE anti-clogging security)

### Provider Bug Detection Through Testing

**Type Conversion Issues:**

- Watch for fields returning memory addresses instead of values
- **Symptom**: Error like `"was cty.StringVal("10"), but now cty.StringVal("0x14000397000")"`
- **Solution**: Check SDK-to-Terraform conversion functions for improper `.String()` calls
- **Fix Pattern**: Use proper type helpers like `mistutils.VlanAsString()` instead of generic `.String()`

**Example Fix:**

```go
// WRONG: Returns memory addresses
types.StringValue(v.String())

// CORRECT: Proper type conversion  
mistutils.VlanAsString(v)
```

### Schema Validation Discovery

**Field Constraints Found Through Testing:**

- `auth.anticlog_threshold`: Must be Number between 16-32
- Field validation errors provide exact constraints to use in fixtures

**Testing Strategy:**

1. Start with reasonable default values
2. If validation fails, read error message for exact constraints
3. Update fixture with valid value
4. Document discovered constraints for future reference

## Success Criteria

- [ ] **100% schema field coverage** (zero missing fields)
- [ ] **Both test cases pass** (simple_case + fixture_case_N)
- [ ] **Clean test execution** (under 30s typical)
- [ ] **Comprehensive validation checks** (all testable fields covered)

## Reference Implementations

**Primary Pattern**: `org_wlantemplate_resource_test.go` - Complete dual test case implementation

**Field Coverage**: `org_wlan_portal_template_resource_test.go` - 100% coverage methodology (225 fields)

**Nested Objects**: `org_wxtag_resource_test.go` - Complex nested array validation

## Quick Troubleshooting

| Issue | Solution |
|-------|----------|
| Missing test cases | Implement both simple_case AND fixture_case patterns |
| HCL generation errors | Only add HCL tags to fields with CTY tags |
| Nested object errors | Validate child attributes only (e.g., `applies.org_id` not `applies`) |
| Field coverage gaps | Use the 4-step verification process above |

**Clean up:**

```bash
# Remove temporary analysis files
rm all_schema_fields.txt current_tested_fields.txt missing_fields.txt

# Remove helper functions from test file after verification
# Delete the TestExtractSchema() function and all extract* helper functions
```
