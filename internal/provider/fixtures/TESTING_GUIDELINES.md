# Terraform Provider Testing Checklist

## Quick Implementation Guide

### 1. Files to Create

- [ ] `{resource}_test_structs.go` - Test struct definitions with HCL tags
- [ ] `{resource}_resource_test.go` - Main test implementation
- [ ] `fixtures/{resource}_resource/{resource}_config.tf` - Test data

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

### Step 1: Extract Schema Fields with Dot Notation

Create a comprehensive reference file for complete field extraction:

```bash
# 1. Use grep to identify all actual field definitions (not containers)
grep -E '"[a-zA-Z_][a-zA-Z0-9_]*":\s*schema\.(StringAttribute|BoolAttribute|Int64Attribute)' \
  internal/resource_{resource_name}/{resource_name}_resource_gen.go

# 2. Analyze schema structure to understand nesting levels
grep -E '(SingleNestedAttribute|MapNestedAttribute)' \
  internal/resource_{resource_name}/{resource_name}_resource_gen.go

# 3. Create structured reference file with:
#    - Root level fields
#    - Direct nested fields (portal_template.field)  
#    - Map nested fields (portal_template.locales.{key}.field)
#    - All with proper dot notation

# 4. Extract final list from reference file
grep -E '^- `[^`]+`' schema_field_reference.md | \
  sed 's/^- `\([^`]*\)`.*/\1/' | sort -u > schema_fields.txt
```

**Example Results:**

- Simple resources: ~10-50 fields
- Complex nested resources: 200+ fields (e.g., org_wlan_portal_template: 225 fields)
- Reference file provides complete field inventory with proper categorization

### Step 2: Extract Tested Fields

```bash
# Extract field paths from actual test execution (WORKING METHOD)
# Include both TestCheckResourceAttr AND TestCheckResourceAttrSet
# Exclude map length validations (.% fields)
go test -v -run "Test{ResourceName}Model" ./internal/provider/ 2>&1 | grep "TestCheckResourceAttr.*{resource_name}\." | sed -n 's/.*TestCheckResourceAttr[^(]*([^,]*, "\([^"]*\)".*/\1/p' | grep -v '\.%$' | sort -u > tested_fields.txt

# Example for org_wlan_portal_template:
# go test -v -run "TestOrgWlanPortalTemplateModel" ./internal/provider/ 2>&1 | grep "TestCheckResourceAttr.*org_wlan_portal_template\." | sed -n 's/.*TestCheckResourceAttr[^(]*([^,]*, "\([^"]*\)".*/\1/p' | grep -v '\.%$' | sort -u > tested_fields.txt
```

### Step 3: Find Missing Fields

```bash
# For simple resources
comm -23 schema_fields.txt tested_fields.txt

# For complex nested resources, compare against reference file
comm -23 {resource_name}_schema_fields_reference.txt tested_fields.txt
```

### Step 4: Achieve 100% Coverage

- [ ] **Zero missing fields required** - the above command must return empty
- [ ] **For complex nested resources**: Create comprehensive reference file first
- [ ] **Manual verification**: Check test output for proper boolean field values (`true` not `false`)
- [ ] **Add missing fields** to test struct and fixture data
- [ ] **Add validation checks** for all fields
- [ ] **Re-verify until complete**

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

**Total Validations = Schema Fields + Map Length Validations (`.%` fields)**

Example: org_wlan_portal_template shows 227 validations (225 schema fields + 2 map validations) = 100% coverage ✅

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

**Clean up**: `rm schema_fields.txt tested_fields.txt` after verification
