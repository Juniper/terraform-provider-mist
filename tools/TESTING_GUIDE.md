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
        resource_org_networktemplate.OrgNetworktemplateResourceSchema(t.Context()).Attributes,
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

**Note:** The tracker automatically normalizes test paths (see "Path Normalization Examples" above). For example, `band_24.0.ant_gain` marks `band_24.ant_gain` as tested since array indices are removed.

### Step 4: Achieve 100% Coverage

- [ ] Zero untested fields required - `untested_fields_count` must be 0
- [ ] Zero unknown fields required - `unknown_fields_count` must be 0 (indicates typos)
- [ ] Verify all test assertions pass
- [ ] Check for schema extraction failures - should be 0

### Field Type Validation Rules

- [ ] Computed-only fields (like `id`): Use `TestCheckResourceAttrSet`
- [ ] All other fields: Use `TestCheckResourceAttr` with expected values
- [ ] Never validate container types - test child attributes only

## Quick Test Update Guide

### Step 1: Add/Update Test Checks

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

### Step 2: Update Fixture Files

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

## Best Practices

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

```hcl
# Good - descriptive
name = "ble_ibeacon_full_config"
name = "port_config_with_radius"

# Bad - unclear
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
```

### 6. Validate Against Real API

When possible, test configurations should reflect valid API scenarios:

``` hcl
# Valid iBeacon config
ble_config = {
  ibeacon_enabled = true
  ibeacon_uuid    = "f3f51b3e-b3c4-4c3e-b3c4-4c3e4c3e4c3e"  # Valid UUID
  ibeacon_major   = 100    # Valid range: 0-65535
  ibeacon_minor   = 200    # Valid range: 0-65535
}

# Invalid config
ble_config = {
  ibeacon_enabled = true
  ibeacon_major   = 70000  # Out of range!
}
```

### Reference Implementations

**Good Examples to Follow:**

- **Primary Pattern**: `org_wlantemplate_resource_test.go` - Complete dual test case implementation
- **Field Coverage**: `org_wlan_portal_template_resource_test.go` - 100% coverage methodology (225 fields)
- **Nested Objects**: `org_wxtag_resource_test.go` - Complex nested array validation
- **Test Structs**: `org_wlan_test_structs.go` - Proper HCL tag usage
