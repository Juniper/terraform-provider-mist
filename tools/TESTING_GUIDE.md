# Terraform Provider Mist - Testing Guide

This comprehensive guide explains how to create and maintain tests for resources in the Terraform Provider Mist project.

## Table of Contents

1. [Overview](#overview)
2. [Quick Implementation Guide](#quick-implementation-guide)
3. [Test Structure](#test-structure)
4. [Testing Workflow](#testing-workflow)
5. [Step-by-Step Instructions](#step-by-step-instructions)
6. [Field Coverage Verification](#field-coverage-verification)
7. [Test Patterns](#test-patterns)
8. [Common Scenarios](#common-scenarios)
9. [Best Practices](#best-practices)
10. [Troubleshooting](#troubleshooting)
11. [Quick Reference](#quick-reference)

---

## Overview

Tests in this project use a **fixture-based approach** combined with **hardcoded simple cases** to validate resource behavior. The testing framework leverages Terraform's plugin testing SDK to perform end-to-end tests.

### Key Components

- **Test Structs**: Go structs representing resource configurations (generated from schemas)
- **Test Checks**: Validation functions that verify Terraform state matches expected values
- **Fixture Files**: HCL configuration files containing comprehensive test scenarios
- **Test Functions**: Table-driven test functions that orchestrate the test execution

### Success Criteria

- [ ] **100% schema field coverage** (zero missing fields)
- [ ] **Both test cases pass** (simple_case + fixture_case_N)
- [ ] **Clean test execution** (under 30s typical)
- [ ] **Comprehensive validation checks** (all testable fields covered)

---

## Quick Implementation Guide

### Files to Create

- [ ] `{resource}_test_structs.go` - Test struct definitions with HCL tags
- [ ] `{resource}_resource_test.go` - Main test implementation
- [ ] `fixtures/{resource}_resource/{resource}_config.tf` - Test data

### Implementation Pattern

Follow `org_wlantemplate_resource_test.go` as the primary reference:

- [ ] **Dual test case strategy**: simple_case + fixture_case_N
- [ ] **Automated HCL generation** with `gohcl.EncodeIntoBody()`
- [ ] **Use `GetTestOrgId()`** - never hardcode org IDs
- [ ] **Comprehensive fixture data** with boolean fields set to `true`

### Test Struct Guidelines

- [ ] **HCL tags on fields with CTY tags OR no tags** (see org_wlan_test_structs.go)
- [ ] **Pointer types for optional fields** (`*string`, `*bool`, `*int64`)
- [ ] **Concrete types for required fields** (`string`, `[]string`)

### Fixture Data Best Practices

- [ ] **Comprehensive field coverage** (aim for 100%)
- [ ] **Set boolean fields to `true`** for maximum test coverage
- [ ] **Include optional fields** with realistic test values
- [ ] **Use `␞` separator** for multiple fixtures

---

## Test Structure

### 1. Test Function Pattern

```go
func TestOrgDeviceprofileApModel(t *testing.T) {
    type testStep struct {
        config OrgDeviceprofileApModel
    }
    
    type testCase struct {
        steps []testStep
    }
    
    testCases := map[string]testCase{
        "simple_case": {
            steps: []testStep{
                {
                    config: OrgDeviceprofileApModel{
                        OrgId: GetTestOrgId(),
                        Name:  "test_ap",
                    },
                },
            },
        },
    }
    
    // Load fixture files...
    // Execute tests...
}
```

### 2. Test Checks Method

Each resource model implements a `testChecks()` method that returns validation checks:

```go
func (s *OrgDeviceprofileApModel) testChecks(t testing.TB, rType, rName string) testChecks {
    checks := newTestChecks(PrefixProviderName(rType) + "." + rName)
    
    // Required fields
    checks.append(t, "TestCheckResourceAttrSet", "org_id")
    checks.append(t, "TestCheckResourceAttr", "name", s.Name)
    
    // Optional fields (conditional)
    if s.BleConfig != nil {
        if s.BleConfig.IbeaconMajor != nil {
            checks.append(t, "TestCheckResourceAttr", "ble_config.ibeacon_major", fmt.Sprintf("%d", *s.BleConfig.IbeaconMajor))
        }
    }
    
    return checks
}
```

### 3. Fixture Files

Located in `internal/provider/fixtures/<resource_name>/` directory:

```hcl
# fixtures/org_deviceprofile_ap_resource/org_deviceprofile_ap_config.tf

# First fixture
name = "comprehensive_test"
ble_config = {
  ibeacon_enabled = true
  ibeacon_major   = 100
  ibeacon_minor   = 200
}
port_config = {
  eth0 = {
    vlan_ids = "10,20,30"
  }
}
␞
# Second fixture (separated by U+241E delimiter)
name = "minimal_config"
␞
# Additional fixtures...
```

---

## Testing Workflow

### When to Update Tests

Tests must be updated when:
1. **Adding new fields** to a resource schema
2. **Modifying field types** (e.g., `list(number)` → `string`)
3. **Changing field behavior** or validation rules
4. **SDK updates** that affect field structure

### The Three-Step Process

```
┌─────────────────────────────────────┐
│ 1. Regenerate Test Structs          │
│    └─> Run gen_test_structs.go      │
└────────────┬────────────────────────┘
             │
             ▼
┌─────────────────────────────────────┐
│ 2. Update Test Checks                │
│    └─> Add/modify testChecks()      │
└────────────┬────────────────────────┘
             │
             ▼
┌─────────────────────────────────────┐
│ 3. Update Fixture Files              │
│    └─> Add/modify HCL configs        │
└─────────────────────────────────────┘
```

---

## Step-by-Step Instructions

### Step 1: Regenerate Test Structs

**When:** After modifying a resource schema or when SDK types change.

**How:**

```bash
cd tools
go run gen_test_structs.go
```

**What it does:**
- Reads resource schemas from `internal/resource_*/terraform_to_sdk*.go`
- Generates Go structs in `internal/provider/*_resource_test_structs.go`
- Creates struct definitions matching Terraform schema types

**Example Output:**
```go
// internal/provider/org_deviceprofile_ap_resource_test_structs.go
type OrgDeviceprofileApModel struct {
    Name      string
    OrgId     string
    BleConfig *BleConfigModel
    // ... other fields
}
```

### Step 2: Add/Update Test Checks

**When:** After regenerating structs or when validation logic needs changes.

**How:**

1. **Locate the test file:** `internal/provider/<resource_name>_resource_test.go`
2. **Find the `testChecks()` method**
3. **Add validation for new/modified fields**

**Example - Adding Check for New Field:**

```go
// Before
if s.BleConfig != nil {
    checks.append(t, "TestCheckResourceAttrSet", "ble_config.%")
    if s.BleConfig.IbeaconEnabled != nil {
        checks.append(t, "TestCheckResourceAttr", "ble_config.ibeacon_enabled", fmt.Sprintf("%t", *s.BleConfig.IbeaconEnabled))
    }
}

// After - Added ibeacon_major and ibeacon_minor
if s.BleConfig != nil {
    checks.append(t, "TestCheckResourceAttrSet", "ble_config.%")
    if s.BleConfig.IbeaconEnabled != nil {
        checks.append(t, "TestCheckResourceAttr", "ble_config.ibeacon_enabled", fmt.Sprintf("%t", *s.BleConfig.IbeaconEnabled))
    }
    if s.BleConfig.IbeaconMajor != nil {
        checks.append(t, "TestCheckResourceAttr", "ble_config.ibeacon_major", fmt.Sprintf("%d", *s.BleConfig.IbeaconMajor))
    }
    if s.BleConfig.IbeaconMinor != nil {
        checks.append(t, "TestCheckResourceAttr", "ble_config.ibeacon_minor", fmt.Sprintf("%d", *s.BleConfig.IbeaconMinor))
    }
}
```

**Example - Updating Check for Type Change:**

```go
// Before (when vlan_ids was list(number))
if len(portCfg.VlanIds) > 0 {
    checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.vlan_ids.#", portPrefix), fmt.Sprintf("%d", len(portCfg.VlanIds)))
    for i, vlanId := range portCfg.VlanIds {
        checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.vlan_ids.%d", portPrefix, i), fmt.Sprintf("%d", vlanId))
    }
}

// After (when vlan_ids is string)
if portCfg.VlanIds != nil {
    checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.vlan_ids", portPrefix), *portCfg.VlanIds)
}
```

---

## Field Coverage Verification

This section covers the mandatory process for achieving 100% field coverage in your tests.

### Understanding Field Count Discrepancies

**Total Validations = Schema Fields + Array Index/Map Key Validations**

**Common Scenarios:**

- **More validations than schema fields**: Tests validate array indices (`portal.amazon_email_domains.0`) and map keys that aren't direct schema fields
- **Fewer validations than schema fields**: Missing field coverage - use corrected analysis to identify gaps

**Example Analysis:**

- org_wlan: 283 schema fields, 339 validations = 56 array/map validations + 100% coverage ✅  
- Missing coverage: 283 schema fields, 190 validations = 93 missing fields (67% coverage) ❌

**Key Insight:** Focus on achieving **zero missing fields** rather than exact count matching.

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
- **100% accuracy**: Captures all nested fields including deeply nested structures
- **Proper dot notation**: All fields in testable format with infinite nesting depth support

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

### Field Type Validation Rules

- [ ] **Computed-only fields** (like `id`): Use `TestCheckResourceAttrSet`
- [ ] **All other fields**: Use `TestCheckResourceAttr` with expected values
- [ ] **Never validate parent nested objects** - test child attributes only

### Nested Attribute Path Structure (CRITICAL)

**Problem**: `SingleNestedAttribute` in schema doesn't always require `.0` indexing in test paths

**Determination Method**: Check existing working tests for similar nested structures in your provider

**Common Patterns:**

1. **Direct nested access** (most common in this provider):

   ```go
   // Schema: band_24 as SingleNestedAttribute
   // Test path: "band_24.allow_rrm_disable" ✅
   // NOT: "band_24.0.allow_rrm_disable" ❌
   checks.append(t, "TestCheckResourceAttr", "band_24.allow_rrm_disable", "true")
   ```

2. **Indexed nested access** (less common):

   ```go
   // Schema: some_config as SingleNestedAttribute  
   // Test path: "some_config.0.field_name" ✅
   checks.append(t, "TestCheckResourceAttr", "some_config.0.field_name", "value")
   ```

**How to Determine Correct Pattern:**

1. **Look at working tests** in same provider (e.g., `org_deviceprofile_ap_resource_test.go`, `device_ap_resource_test.go`)
2. **Check similar nested structures** - if they use direct paths, use direct paths
3. **Test error messages** - if test fails with "Attribute 'field.0.subfield' not found", try without `.0`

**Examples from this provider:**

- ✅ `"radio_config.band_24.allow_rrm_disable"` (direct - works)
- ✅ `"band_24.ant_gain"` (direct - works)  
- ❌ `"band_24.0.ant_gain"` (indexed - fails)

**Common Error Symptom:**

```text
Check failed: mist_resource.test: Attribute 'nested_field.0.subfield' not found
```

**Solution**: Remove `.0` and use direct nested path: `nested_field.subfield`

### Manual Test Quality Verification

- [ ] **Boolean field values**: Search test output for `"false"` - should be minimal

```bash
go test -v -run "Test{ResourceName}Model" ./internal/provider/ 2>&1 | grep 'TestCheckResourceAttr.*"false"'
```

- [ ] **Null value checks**: Search for `= null` in fixture data - boolean fields should use `true`

```bash
grep "= null" internal/provider/fixtures/{resource_name}_resource/{resource_name}_config.tf
```

- [ ] **Missing validation checks**: Count validations vs expected fields

```bash
go test -v -run "Test{ResourceName}Model" ./internal/provider/ 2>&1 | grep -c "TestCheckResourceAttr.*{resource_name}\."
```

### Field Implementation Strategy

#### One-by-One Field Addition (Recommended for Complex Resources)

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

#### Provider Bug Detection Through Testing

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

#### Schema Validation Discovery

**Field Constraints Found Through Testing:**

- `auth.anticlog_threshold`: Must be Number between 16-32
- Field validation errors provide exact constraints to use in fixtures

**Testing Strategy:**

1. Start with reasonable default values
2. If validation fails, read error message for exact constraints
3. Update fixture with valid value
4. Document discovered constraints for future reference

---

### Step 3: Update Fixture Files

**When:** After adding test checks to validate new/modified fields.

**Location:** `internal/provider/fixtures/<resource_name>/<resource_name>_config.tf`

**Strategy:**

1. **Try adding to the main fixture first** (first configuration before any `␞` delimiter)
2. **If conflicts occur**, create a separate fixture for that specific case

**Main Fixture Approach (Preferred):**

```hcl
# Add new fields to the comprehensive fixture
name = "comprehensive_test"
ble_config = {
  ibeacon_enabled = true
  ibeacon_major   = 100  # Added
  ibeacon_minor   = 200  # Added
  ibeacon_uuid    = "f3f51b3e-b3c4-4c3e-b3c4-4c3e4c3e4c3e"
}
port_config = {
  eth0 = {
    vlan_ids = "10,20,30"  # Updated type
  }
}
```

**Separate Fixture Approach (When Conflicts Arise):**

```hcl
# Main fixture
name = "comprehensive_test"
ble_config = {
  ibeacon_enabled = true
}
␞
# Separate fixture for conflicting scenario
name = "ble_only_test"
ble_config = {
  ibeacon_enabled = true
  ibeacon_major   = 100
  ibeacon_minor   = 200
  ibeacon_uuid    = "f3f51b3e-b3c4-4c3e-b3c4-4c3e4c3e4c3e"
}
␞
# Another separate fixture
name = "port_config_test"
port_config = {
  eth0 = {
    vlan_ids = "10,20,30"
  }
}
```

---

## Test Patterns

### Pattern 1: Required Fields

Always test required fields in both simple and fixture tests:

```go
// Simple case
testCases := map[string]testCase{
    "simple_case": {
        steps: []testStep{
            {
                config: OrgDeviceprofileApModel{
                    OrgId: GetTestOrgId(),  // Required
                    Name:  "test_ap",        // Required
                },
            },
        },
    },
}

// Test checks
checks.append(t, "TestCheckResourceAttrSet", "org_id")
checks.append(t, "TestCheckResourceAttr", "name", s.Name)
```

### Pattern 2: Optional Nested Objects

Use nil-checks before accessing nested fields:

```go
if s.BleConfig != nil {
    checks.append(t, "TestCheckResourceAttrSet", "ble_config.%")
    if s.BleConfig.IbeaconMajor != nil {
        checks.append(t, "TestCheckResourceAttr", "ble_config.ibeacon_major", fmt.Sprintf("%d", *s.BleConfig.IbeaconMajor))
    }
}
```

### Pattern 3: Lists/Arrays

Check both count and individual elements:

```go
if len(s.NtpServers) > 0 {
    checks.append(t, "TestCheckResourceAttr", "ntp_servers.#", fmt.Sprintf("%d", len(s.NtpServers)))
    for i, server := range s.NtpServers {
        checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ntp_servers.%d", i), server)
    }
}
```

### Pattern 4: Maps/Objects

Check both count and key-value pairs:

```go
if len(s.Vars) > 0 {
    checks.append(t, "TestCheckResourceAttr", "vars.%", fmt.Sprintf("%d", len(s.Vars)))
    for key, value := range s.Vars {
        checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vars.%s", key), value)
    }
}
```

---

## Common Scenarios

### Scenario 1: Adding a New Optional Field

**Example:** Adding `ibeacon_major` to BLE config

1. **Regenerate structs:**
   ```bash
   cd tools && go run gen_test_structs.go
   ```

2. **Update test checks:**
   ```go
   if s.BleConfig.IbeaconMajor != nil {
       checks.append(t, "TestCheckResourceAttr", "ble_config.ibeacon_major", fmt.Sprintf("%d", *s.BleConfig.IbeaconMajor))
   }
   ```

3. **Update fixture:**
   ```hcl
   ble_config = {
     ibeacon_enabled = true
     ibeacon_major   = 100  # Add this
   }
   ```

### Scenario 2: Changing Field Type

**Example:** Changing `vlan_ids` from `list(number)` to `string`

1. **Regenerate structs:**
   ```bash
   cd tools && go run gen_test_structs.go
   ```
   
   Result: `VlanIds []int` → `VlanIds *string`

2. **Update test checks:**
   ```go
   // Old
   if len(portCfg.VlanIds) > 0 {
       checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.vlan_ids.#", portPrefix), fmt.Sprintf("%d", len(portCfg.VlanIds)))
   }
   
   // New
   if portCfg.VlanIds != nil {
       checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.vlan_ids", portPrefix), *portCfg.VlanIds)
   }
   ```

3. **Update fixture:**
   ```hcl
   # Old
   port_config = {
     eth0 = {
       vlan_ids = [10, 20, 30]
     }
   }
   
   # New
   port_config = {
     eth0 = {
       vlan_ids = "10,20,30"
     }
   }
   ```

### Scenario 3: Adding Nested Object

**Example:** Adding RADIUS config to port config

1. **Regenerate structs** (captures new nested structure)

2. **Add test checks for all nested fields:**
   ```go
   if portCfg.RadiusConfig != nil {
       radiusPrefix := fmt.Sprintf("%s.radius_config", portPrefix)
       checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("%s.%%", radiusPrefix))
       
       if portCfg.RadiusConfig.Network != nil {
           checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.network", radiusPrefix), *portCfg.RadiusConfig.Network)
       }
       
       if len(portCfg.RadiusConfig.AuthServers) > 0 {
           // ... check auth servers
       }
   }
   ```

3. **Create separate fixture** (if complex/conflicts):
   ```hcl
   ␞
   name = "radius_config_test"
   port_config = {
     eth0 = {
       radius_config = {
         network = "default"
         auth_servers = [
           {
             host   = "192.168.1.1"
             secret = "secret123"
           }
         ]
       }
     }
   }
   ```

---

## Best Practices

### 1. Always Use Nil-Checks

Optional fields must be checked for nil before access:

```go
// ✅ Good
if s.BleConfig != nil {
    if s.BleConfig.IbeaconMajor != nil {
        checks.append(t, "TestCheckResourceAttr", "ble_config.ibeacon_major", fmt.Sprintf("%d", *s.BleConfig.IbeaconMajor))
    }
}

// ❌ Bad - will panic if nil
checks.append(t, "TestCheckResourceAttr", "ble_config.ibeacon_major", fmt.Sprintf("%d", *s.BleConfig.IbeaconMajor))
```

### 2. Use Descriptive Fixture Names

```hcl
# ✅ Good - descriptive
name = "ble_ibeacon_full_config"
name = "port_config_with_radius"

# ❌ Bad - unclear
name = "test1"
name = "config"
```

### 3. Format Checks Appropriately by Type

```go
// Boolean
checks.append(t, "TestCheckResourceAttr", "enabled", fmt.Sprintf("%t", *s.Enabled))

// Integer
checks.append(t, "TestCheckResourceAttr", "port", fmt.Sprintf("%d", *s.Port))

// String
checks.append(t, "TestCheckResourceAttr", "host", *s.Host)
```

### 4. Fixture File Organization

```hcl
# 1. Most comprehensive fixture first
name = "comprehensive_test"
# ... all possible fields

␞
# 2. Specific feature fixtures
name = "ble_config_only"
# ... BLE specific

␞
# 3. Edge case fixtures
name = "minimal_required_fields"
# ... only required fields
```

### 5. Run Tests After Changes

```bash
# Run specific resource test
go test -v -run TestOrgDeviceprofileApModel ./internal/provider/

# Run all tests
go test -v ./internal/provider/

# Run with coverage
go test -cover ./internal/provider/
```

### 6. Validate Against Real API

When possible, test configurations should reflect valid API scenarios:

```hcl
# ✅ Valid iBeacon config
ble_config = {
  ibeacon_enabled = true
  ibeacon_uuid    = "f3f51b3e-b3c4-4c3e-b3c4-4c3e4c3e4c3e"  # Valid UUID
  ibeacon_major   = 100    # Valid range: 0-65535
  ibeacon_minor   = 200    # Valid range: 0-65535
}

# ❌ Invalid config
ble_config = {
  ibeacon_enabled = true
  ibeacon_major   = 70000  # Out of range!
}
```

---

## Troubleshooting

### Common Issues and Solutions

| Issue | Error Message | Solution |
|-------|--------------|----------|
| **Fixture parsing errors** | `error decoding hcl: ...` | - Ensure fixture file has valid HCL syntax<br>- Check delimiter is exactly `␞` (U+241E)<br>- Verify all required fields are present |
| **Test check failures** | `TestCheckResourceAttr failed: expected X, got Y` | - Verify fixture data matches expected values<br>- Check type formatting in checks (use `fmt.Sprintf` correctly)<br>- Ensure struct regeneration captured latest schema |
| **Nil pointer panics** | `panic: runtime error: invalid memory address` | - Add nil-checks for all optional fields<br>- Use pattern: `if field != nil { ... }`<br>- Never dereference pointers without checking |
| **Nested attribute errors** | `Attribute 'field.0.subfield' not found` | - Remove `.0` - use direct nested paths (see Nested Attribute Path Structure)<br>- Check similar working tests for correct pattern |
| **Missing test cases** | Only one test case runs | - Implement both simple_case AND fixture_case patterns |
| **HCL generation errors** | HCL encoding fails | - Only add HCL tags to fields with CTY tags<br>- Check tag format matches schema definitions |
| **Field coverage gaps** | Tests don't validate all fields | - Use the 4-step field coverage verification process<br>- Run corrected analysis to find truly missing fields |

---

## Quick Reference

### Essential Commands

```bash
# Regenerate test structs
cd tools && go run gen_test_structs.go

# Run specific resource test
go test -v -run TestOrgDeviceprofileApModel ./internal/provider/

# Run all provider tests
go test -v ./internal/provider/

# Run with coverage
go test -cover ./internal/provider/

# Run with race detector
go test -race ./internal/provider/

# Extract schema fields for coverage analysis
go test -v -run "TestExtractSchema" ./internal/provider/ >/dev/null 2>&1

# Extract currently tested fields
go test -v -run "Test{ResourceName}Model" ./internal/provider/ 2>&1 | \
  grep "TestCheckResourceAttr.*{resource_name}\." | \
  sed -n 's/.*TestCheckResourceAttr[^(]*([^,]*, "\([^"]*\)".*/\1/p' | \
  grep -v '\.%$' | sort -u > current_tested_fields.txt

# Find missing fields (corrected analysis)
grep -E '\.(#|[0-9]+)$' current_tested_fields.txt | \
  sed 's/\.[#0-9].*$//' | sort -u > properly_tested_containers.txt
(cat current_tested_fields.txt; cat properly_tested_containers.txt) | \
  sort -u > corrected_tested_fields.txt
comm -23 all_schema_fields.txt corrected_tested_fields.txt > truly_missing_fields.txt

# Clean up analysis files
rm all_schema_fields.txt current_tested_fields.txt truly_missing_fields.txt \
   corrected_tested_fields.txt properly_tested_containers.txt
```

### File Structure

```
internal/provider/
├── <resource>_resource_test.go           # Test function & testChecks
├── <resource>_resource_test_structs.go   # Generated structs (auto-generated)
└── fixtures/
    ├── README.md
    ├── TESTING_GUIDELINES.md             # Detailed testing guidelines
    └── <resource>_resource/
        └── <resource>_config.tf          # Fixture data
```

### Fixture Delimiter

**Character:** `␞` (U+241E - Symbol for Record Separator)

**Usage:**
```hcl
# Fixture 1
name = "test1"
ble_config = {
  ibeacon_enabled = true
}
␞
# Fixture 2
name = "test2"
port_config = {
  eth0 = {
    vlan_ids = "10,20,30"
  }
}
```

### Reference Implementations

**Best Examples to Follow:**

- **Primary Pattern**: `org_wlantemplate_resource_test.go` - Complete dual test case implementation
- **Field Coverage**: `org_wlan_portal_template_resource_test.go` - 100% coverage methodology (225 fields)
- **Nested Objects**: `org_wxtag_resource_test.go` - Complex nested array validation
- **Test Structs**: `org_wlan_test_structs.go` - Proper HCL tag usage

---

## Complete Testing Checklist

### Initial Setup

- [ ] Run `gen_test_structs.go` to generate test structs
- [ ] Create `{resource}_resource_test.go` with test function
- [ ] Create fixture file in `fixtures/{resource}_resource/{resource}_config.tf`
- [ ] Implement both `simple_case` and `fixture_case` patterns

### Test Implementation

- [ ] Add HCL tags only to fields with CTY tags or no tags
- [ ] Use pointer types for optional fields
- [ ] Use concrete types for required fields
- [ ] Implement `testChecks()` method
- [ ] Add nil-checks for all optional fields
- [ ] Use `GetTestOrgId()` instead of hardcoding

### Field Coverage Verification (MANDATORY)

- [ ] Add `TestExtractSchema()` helper to extract all schema fields
- [ ] Run schema extraction: `go test -v -run "TestExtractSchema"`
- [ ] Extract currently tested fields from test output
- [ ] Run corrected analysis to find truly missing fields
- [ ] Achieve zero missing fields (`truly_missing_fields.txt` empty)
- [ ] Clean up temporary analysis files
- [ ] Remove helper functions from test file

### Fixture Data Quality

- [ ] Set boolean fields to `true` for maximum coverage
- [ ] Include all optional fields with realistic values
- [ ] Use valid values that respect schema constraints
- [ ] Separate conflicting scenarios with `␞` delimiter
- [ ] Verify no `= null` for testable fields

### Test Execution

- [ ] Run tests: `go test -v -run Test{ResourceName}Model`
- [ ] Verify both test cases pass (simple + fixture)
- [ ] Check for nil pointer panics
- [ ] Verify test execution time is reasonable (<30s)
- [ ] Run with race detector: `go test -race`

### Final Verification

- [ ] Verify 100% field coverage (zero missing fields)
- [ ] Check for minimal `"false"` values in test output
- [ ] Ensure proper nested path structure (no incorrect `.0` usage)
- [ ] Commit test files AND fixture files together
- [ ] Update documentation if new patterns discovered

---

## Summary

This guide provides comprehensive instructions for creating and maintaining high-quality tests in the Terraform Provider Mist project. Key principles:

1. **Always aim for 100% field coverage** using the corrected analysis method
2. **Follow established patterns** from reference implementations
3. **Use systematic approaches** for complex resources (one-by-one field addition)
4. **Test quality over speed** - discover constraints and dependencies through testing
5. **Clean up after verification** - remove helper functions and temporary files

For detailed examples and advanced scenarios, refer to the sections above and the reference implementations listed.

---

**Last Updated:** December 16, 2025  
**Primary Reference:** `org_wlantemplate_resource_test.go`  
**Coverage Reference:** `org_wlan_portal_template_resource_test.go`
