# Terraform Provider Testing Checklist

## Quick Implementation Guide

### 1. Files to Create/Verify

- [ ] `{resource}_test_structs.go` - **Check if exists first**; test struct definitions with HCL tags
- [ ] `{resource}_resource_test.go` - **Check if exists first**; main test implementation
- [ ] `fixtures/{resource}_resource/{resource}_config.tf` - Test data

### 2. Implementation Pattern (Follow org_wlantemplate_resource_test.go)

- [ ] Dual test case strategy: simple_case + fixture_case_N
- [ ] Automated HCL generation with `gohcl.EncodeIntoBody()`
- [ ] Use `GetTestOrgId()` - never hardcode org IDs
- [ ] Comprehensive fixture data with boolean fields set to `true`

### 3. Test Struct Guidelines

- [ ] HCL tags on fields with CTY tags OR no tags (see org_wlan_test_structs.go)
- [ ] Pointer types for optional fields (`*string`, `*bool`, `*int64`)
- [ ] Concrete types for required fields (`string`, `[]string`)

### 4. Fixture Data Best Practices

- [ ] Comprehensive field coverage (aim for 100%)
- [ ] Set boolean fields to `true` for maximum test coverage
- [ ] Include optional fields with realistic test values
- [ ] Use `␞` separator for multiple fixtures when:
  - Fields are mutually exclusive or create conflicts
  - Testing different behavioral scenarios
  - Single fixture becomes unwieldy (>100 lines)
- [ ] Prefer single comprehensive fixture when fields can coexist

### 5. Schema Validation Rules

Some schema attributes validate their keys/values. Check resource documentation.

**Map Key Patterns:**
- NAT rules: IP addresses, IP:Port, CIDRs, ports, or variables
- Route tables: CIDR notation
- Multicast groups: Multicast IPs with CIDR

## Field Coverage Verification

### Field Coverage Tracker Overview

The automated field coverage tracker (`internal/provider/validators/field_coverage_validator.go`) identifies untested schema fields:

1. **Automatically extracts** all fields from resource schema using reflection
2. **Intercepts test assertions** to mark fields as tested
3. **Normalizes paths** to handle arrays, maps, and nested structures
4. **Generates JSON reports** of untested fields

**Disable tracking**: Set `DISABLE_MIST_FIELD_COVERAGE_TRACKER` environment variable

**Usage Pattern:**

```go
tracker := validators.FieldCoverageTrackerWithSchema(
    "org_networktemplate",
    resource.OrgNetworktemplateResourceSchema(t.Context()).Attributes,
)
// Pass tracker to testChecks(), run tests, then call tracker.FieldCoverageReport(t)
```

**Tracker Notification:**

- `checks.append()` method → Notifies tracker for single field
- `checks.appendSetNestedCheck()` method → Notifies tracker for nested map fields
- Built-in `append()` function → Does NOT notify tracker (internal slice operations)

**Path Normalization:**

Tracker normalizes test paths to schema paths:
- Array indices removed: `privileges.0.role` → `privileges.role`
- Map keys replaced: `networks.guest.vlan_id` → `networks.{key}.vlan_id`
- Hash symbols removed: `dns_servers.#` → `dns_servers`

**Container exclusions:** Container types (SingleNestedAttribute, MapNestedAttribute) aren't tested directly - only their children.

### Step 1: Run Tests and Generate Coverage Report

```bash
# Run tests to generate coverage report
go test -v -run "Test{ResourceName}Model" ./internal/provider/ 2>&1 | tee test_output.txt

# The tracker automatically outputs JSON report with:
# - tested_fields_count: Number of unique fields validated
# - untested_fields: List of schema fields not validated in tests
# - unknown_fields: Test paths that don't match schema (potential typos)
```

### Step 2: Analyze Coverage Report

Focus on `untested_fields` (missing validation) and `unknown_fields` (typos - should be zero).

### Step 3: Address Untested Fields

1. Group missing fields by category from `untested_fields` list
2. Add fields to test struct with appropriate Go types and HCL tags
3. Add fields to fixture data with realistic test values  
4. Add TestCheckResourceAttr validations for each field
5. Re-run tests until `untested_fields_count` is zero

### Step 4: Achieve 100% Coverage

- [ ] Zero untested fields required - `untested_fields_count` must be 0
- [ ] Zero unknown fields required - `unknown_fields_count` must be 0 (indicates typos)
- [ ] Verify all test assertions pass
- [ ] Check for schema extraction failures - should be 0

**Maximum Achievable Coverage:** Some fields may be untestable due to provider bugs. Document these and aim for 99%+ coverage.

**Large Resources (100+ fields):** Work incrementally in small batches (5-15 fields). Test after each batch for easier debugging. Watch for mutually exclusive fields.

### Field Type Validation Rules

- **Computed-only fields** (like `id`): Use `TestCheckResourceAttrSet`
- **All other fields**: Use `TestCheckResourceAttr` with expected values
- **Container types**: Test child attributes only, not the container itself

## Quick Test Update Guide

### Scenario 1: Improving Coverage (Most Common)

If `testChecks()` has validation code, just add missing fields to fixture file `fixtures/{resource}_resource/{resource}_config.tf`.

### Scenario 2: New Fields After Schema Changes

Find `testChecks()` in `internal/provider/<resource_name>_resource_test.go`, add validation with nil-checks, then add fixture data.

### Step 3: Update Fixture Files

**Location:** `internal/provider/fixtures/<resource_name>/<resource_name>_config.tf`

**Strategy:** Add to main fixture first. If conflicts occur, create separate fixture with `␞` delimiter.

## Best Practices

### Coverage Targets

Aim for 100% coverage. If provider bugs block fields, 99%+ is excellent. Below 95% needs work.

### 1. Always Use Nil-Checks

Optional fields must be checked for nil before access:

```go
// Good
if s.BleConfig != nil {
    if s.BleConfig.IbeaconMajor != nil {
        checks.append(t, "TestCheckResourceAttr", "ble_config.ibeacon_major", fmt.Sprintf("%d", *s.BleConfig.IbeaconMajor))
    }
}

// Bad - will panic if nil
checks.append(t, "TestCheckResourceAttr", "ble_config.ibeacon_major", fmt.Sprintf("%d", *s.BleConfig.IbeaconMajor))
```

### 2. Use Descriptive Fixture Names

Use descriptive names like `ble_ibeacon_full_config`, not `test1`.

### 3. Prefer Adding Over Removing

Add fields to existing fixtures when possible. Only remove for duplicates, invalid data, or removed schema fields.

### 4. Format Checks Appropriately by Type

```go
// Boolean
checks.append(t, "TestCheckResourceAttr", "enabled", fmt.Sprintf("%t", *s.Enabled))

// Integer
checks.append(t, "TestCheckResourceAttr", "port", fmt.Sprintf("%d", *s.Port))

// String
checks.append(t, "TestCheckResourceAttr", "host", *s.Host)
```

### 5. Fixture Organization

Place most comprehensive fixture first, then specific features, then edge cases. Separate with `␞`.

### 6. Run Tests After Changes

```bash
# Run specific resource test
go test -v -run TestOrgDeviceprofileApModel ./internal/provider/
```

### 7. Use Valid API Values

Test configurations should reflect valid API scenarios (correct types, valid ranges, proper formats).

## Troubleshooting

### Terraform Schema Validation Errors

**Cause:** Fixture data violates schema validation rules.

**Fix:** Check error message for attribute path and validation rule. Verify map keys match expected patterns (IPs, CIDRs, etc.).

### Go Compilation Errors

**Fix:** Verify test struct types match schema. Check HCL/CTY tags. Use pointers for optional fields.

### Unknown Fields

**Symptom:** `unknown_fields_count > 0`

**Fix:** Review `unknown_fields` list for typos in `testChecks()` method.

### Type Formatting Errors

**Cause:** `TestCheckResourceAttr` requires string values.

**Fix:** Use `fmt.Sprintf("%d", *intValue)` for integers, `fmt.Sprintf("%t", *boolValue)` for booleans. Always nil-check pointers first.

### Mutually Exclusive Fields

**Symptom:** Schema validation errors about conflicting fields.

**Solution:** Create separate fixtures with `␞` delimiter for mutually exclusive options.

### Provider Bugs Blocking Coverage

**Symptom:** Field accepts values but returns `null` after apply.

**Solution:** Remove field from fixture, document in test file comment, accept 99.x% coverage.

### Reference Implementations

**Good Examples to Follow:**

- **Primary Pattern**: `org_wlantemplate_resource_test.go` - Complete dual test case implementation
- **Field Coverage**: `org_wlan_portal_template_resource_test.go` - 100% coverage methodology (225 fields)
- **Nested Objects**: `org_wxtag_resource_test.go` - Complex nested array validation
- **Test Structs**: `org_wlan_test_structs.go` - Proper HCL tag usage
