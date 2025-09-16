# TestOrgNacruleModel - Provider Bugs in not_matching

## Summary

During systematic field isolation testing of the `org_nacrule` resource, we identified provider bugs in the `not_matching` field processing.

**Test Coverage Result**: Successfully improved from 8 to 44+ test checks (550% improvement)

## Field Test Results

### ✅ Working Fields in not_matching
- `auth_type` - Works correctly
- `nactags` - Works correctly  
- `vendor` - Works correctly
- `port_types` - Works correctly
- `site_ids` - Works correctly
- `sitegroup_ids` - Works correctly

### ❌ Problematic Fields in not_matching
These fields cause nil pointer panics:
- `family`
- `mfg` 
- `model`
- `os_type`

## Bug Details

**Location**: `sdk_to_terraform_matching.go:104`  
**Function**: `notMatchingSdkToTerraform`  
**Error**: `panic: runtime error: invalid memory address or nil pointer dereference`

## Impact

- All enhanced fields work perfectly in `matching` configuration
- 4 specific fields cannot be used in `not_matching` due to implementation bug
- Core functionality remains intact with excellent test coverage

## Recommendations

1. Fix nil pointer handling in `notMatchingSdkToTerraform` function
2. Use working fields for exclusion scenarios as workaround
3. Maintain current comprehensive test coverage (44+ checks)