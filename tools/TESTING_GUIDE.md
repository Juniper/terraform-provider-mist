# Terraform Provider Testing Checklist

## Quick Implementation Guide

### 1. Files to Create

- [ ] `{resource}_test_structs.go` - Test struct definitions with HCL tags
- [ ] `{resource}_resource_test.go` - Main test implementation
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
- [ ] Use `␞` separator for multiple fixtures

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
func TestOrgNetworktemplateModel(t *testing.T) {
    resourceType := "org_networktemplate"
    tracker := validators.FieldCoverageTrackerWithSchema(
        resourceType,
        resource_org_networktemplate.OrgNetworktemplateResourceSchema(context.Background()).Attributes,
    )

    for tName, tCase := range testCases {
        t.Run(tName, func(t *testing.T) {
            checks := tCase.testChecks(t, resourceType, tName, tracker)
            // ... run resource tests ...
        })
    }

    if tracker != nil {
        tracker.FieldCoverageReport(t)
    }
}

func (o *OrgNetworktemplateModel) testChecks(..., tracker *validators.FieldCoverageTracker) testChecks {
    checks := newTestChecks(PrefixProviderName(rType) + "." + tName, tracker) // Tracker marks fields as tested when it is appended to the checks
    checks.append(t, "TestCheckResourceAttr", "name", o.Name)
    // ... more checks ...
    return checks
}
```

**Path Normalization Examples:**

- `privileges.0.role` → `privileges.role` (array index removed)
- `networks.guest.vlan_id` → `networks.{key}.vlan_id` (MapNestedAttribute key replaced)
- `vars.my_var` → `vars.{key}` (MapAttribute key replaced)
- `extra_routes.10.0.0.0/8.via` → `extra_routes.{key}.via` (CIDR treated as map key)
- `dns_servers.#` → `dns_servers` (hash symbol removed)
- `vrf_instances.default.extra_routes6.2001:db8::/32.via` → `vrf_instances.{key}.extra_routes6.{key}.via` (nested maps)

**Container Type Exclusion:**

Container types cannot be tested directly - only their children:

- `radius_config` (SingleNestedAttribute) - untestable container
- `radius_config.timeout` (Int64Attribute) - testable field
- `networks` (MapNestedAttribute) - untestable container
- `networks.{key}.vlan_id` (Int64Attribute) - testable field
- `vars` (MapAttribute) - testable as `vars.{key}` (keys are dynamic)

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

The tracker's JSON output shows:

```json
{
  "resource_name": "org_networktemplate",
  "tested_fields_count": 190,
  "untested_fields_count": 93,
  "untested_fields": [
    "auth.anticlog_threshold",
    "auth.enable_mac_auth",
    "dynamic_vlan.enabled"
  ],
  "unknown_fields_count": 0,
  "unknown_fields": [],
  "schema_extraction_failures_count": 0,
  "schema_extraction_failures": []
}
```

**Focus on:**

- `untested_fields`: Schema fields missing test validation
- `unknown_fields`: Potential typos in test paths (should be zero)

### Step 3: Address Untested Fields

**Systematic Addition Process:**

1. **Group missing fields by category** from `untested_fields` list (auth.*, portal.*, etc.)
2. **Add fields to test struct** with appropriate Go types and HCL tags
3. **Add fields to fixture data** with realistic test values  
4. **Add TestCheckResourceAttr validations** for each field
5. **Re-run tests** until `untested_fields_count` is zero

**Path Normalization Behavior:**

The tracker automatically normalizes test paths:

- `coa_servers.0.ip` → marks `coa_servers.ip` as tested (array index removed)
- `auth_servers.#` → marks `auth_servers` as tested (hash removed)
- `networks.guest.vlan_id` → marks `networks.{key}.vlan_id` as tested (map key replaced)
- `band_24.ant_gain` → marks `band_24.ant_gain` as tested (direct nested path)
- `band_24.0.ant_gain` → marks `band_24.ant_gain` as tested (indexed nested path, index removed)

**Note:** `SingleNestedAttribute` paths in this provider typically use direct access (no `.0` index). Check existing tests for patterns.

### Step 4: Achieve 100% Coverage

- [ ] Zero untested fields required - `untested_fields_count` must be 0
- [ ] Zero unknown fields required - `unknown_fields_count` must be 0 (indicates typos)
- [ ] Verify all test assertions pass
- [ ] Check for schema extraction failures - should be 0

### Field Type Validation Rules

- [ ] Computed-only fields (like `id`): Use `TestCheckResourceAttrSet`
- [ ] All other fields: Use `TestCheckResourceAttr` with expected values
- [ ] Never validate container types - test child attributes only

## Quick Troubleshooting

| Issue | Solution |
| ----- | -------- |
| Missing test cases | Implement both simple_case AND fixture_case patterns |
| HCL generation errors | Only add HCL tags to fields with CTY tags |
| Nested object errors | Validate child attributes only (e.g., `applies.org_id` not `applies`) |
| Field coverage gaps | Use the 4-step verification process above |
| "Attribute 'field.0.subfield' not found" | Remove `.0` - use direct nested paths |
