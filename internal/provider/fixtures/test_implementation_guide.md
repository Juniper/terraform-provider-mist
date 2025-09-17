# Test Implementation Guide for Terraform Provider Mist Resources

## Overview

This guide provides a complete template for implementing test coverage for any resource in the terraform-provider-mist codebase, following established patterns and conventions.

## Part 1: Field Analysis & Discovery

### Extract Complete Schema Fields and Currently Tested Fields

**Follow the field extraction methodology from `/internal/provider/fixtures/field_extraction_guidelines.md`:**

**THE field_extraction_guidelines.md DOCUMENT IS OF UTMOST IMPORTANCE!**

1. **Extract all schema fields** using the automated reflection-based method
2. **Extract currently tested fields** from test execution output

**This process generates two files:**

- `all_schema_fields.txt` - Complete schema field list with proper dot notation
- `current_tested_fields.txt` - Currently validated fields in tests

**Expected outcomes:**

- **100% field accuracy** using reflection-based extraction
- **Proper nested field handling** for complex structures

### Field Analysis Categories

- **Computed fields**: Read-only, populated by API (use `TestCheckResourceAttrSet`)
- **Configurable fields**: User-settable (use `TestCheckResourceAttr` with expected values)
- **Required vs Optional**: Important for test configuration
- **Complex types**: Maps, lists, nested objects requiring special validation patterns

---

## Part 2: Test Implementation

### Prerequisites

- Completed Part 1 field analysis
- Received definitive field list from user (source of truth)
- Identified field categories (computed vs configurable)

### Project Structure & Conventions

#### File Locations

- **Main test file**: `/internal/provider/{resource}_resource_test.go`
- **Test structs file**: `/internal/provider/{resource}_test_structs.go`
- **Fixture file**: `/internal/provider/fixtures/{resource}_resource/{resource}_config.tf`
- **Resource schema**: `/internal/{resource}/{resource}_resource_gen.go`

#### Established Patterns

- Follow the **EXACT SAME PATTERN** as `site_wlan_resource_test.go` for standard resources
- Use the same test structure, provider configuration, and validation approach  
- For resources with nested site references, see the `org_inventory_resource_test.go` pattern as a specialized case

### Step 1: Test Structs Analysis and Updates

**IMPORTANT: Test structs are generated files, but can be enhanced systematically.**

The test structs in `/internal/provider/{resource}_test_structs.go` are automatically generated. You can enhance them in two ways:

1. **Add missing HCL tags**: Add `hcl:"{field_name}"` tags to fields that have `cty` tags but missing `hcl` tags
2. **Add complete missing structs**: Add entire struct definitions that exist in the corresponding `{resource}_resource_gen.go` but are missing from test structs

**Two enhancement patterns:**

#### Pattern 1: Add HCL tags to existing fields

```go
// Example: Add missing hcl tags to existing fields
type MatchingValue struct {
    AuthType     *string  `cty:"auth_type" hcl:"auth_type"`  // ADD hcl tag if missing
    Nactags      []string `cty:"nactags" hcl:"nactags"`      // ADD hcl tag if missing  
}
```

#### Pattern 2: Add complete missing structs

```go
// If AiristaValue exists in device_ap_resource_gen.go but not in test structs:
type AiristaValue struct {
    Enabled *bool   `cty:"enabled" hcl:"enabled"`
    Host    *string `cty:"host" hcl:"host"`  
    Port    *int64  `cty:"port" hcl:"port"`
}
```

**Enhancement Rules:**

- **Compare with resource schema**: Check `{resource}_resource_gen.go` for missing struct definitions
- **Add complete structs only**: Never add partial structs - they must match the resource schema exactly
- **Use proper tags**: Base struct fields need `hcl` tags, nested struct fields need both `cty` and `hcl` tags
- **Verify field coverage improvement**: Adding complete structs can significantly increase testable field coverage

### Step 2: Fixture Configuration

Create `/internal/provider/fixtures/{resource}_resource/{resource}_config.tf` with **realistic, working values**:

```hcl
# Start simple and add complexity gradually
field1 = "test_value"
field2 = true
# Include nested objects but keep them simple initially
nested_field = {
  sub_field1 = "nested_value"
}
```

**Conventions:**

- **Start simple**: Begin with basic required fields and simple optional fields
- **Test incrementally**: Add complex nested objects only after basic fields work
- **Use valid UUIDs**: For fields like `site_ids`, `sitegroup_ids`, use valid UUID format: `"11111111-1111-1111-1111-111111111111"`
- **No comments in fixture files**: Never use any comments in the fixture file
- **Real data when available**: Use real MAC addresses, device IDs when available
- **Separate multiple configs with `␞` delimiter** if needed
- **Handle provider limitations**: Some complex field combinations may cause provider SDK issues - simplify if needed

**Expected Behavior:**

- **Simple configurations should always work**
- **Complex nested objects may need simplification** if they cause provider inconsistency errors
- **Start minimal, expand gradually** to find the working complexity level

### Step 3: Main Test Implementation

Create `/internal/provider/{resource}_resource_test.go` following exact pattern:

```go
package provider

import (
    "fmt"
    "os"
    "strings"
    "testing"

    "github.com/hashicorp/hcl"
    "github.com/hashicorp/hcl/v2/gohcl"
    "github.com/hashicorp/hcl/v2/hclwrite"
    "github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func Test{Resource}Model(t *testing.T) {
    type testStep struct {
        config {Resource}Model
    }

    type testCase struct {
        steps []testStep
    }

    testCases := map[string]testCase{
        "simple_case": {
            steps: []testStep{
                {
                    config: {Resource}Model{
                        // Basic required fields
                    },
                },
            },
        },
    }

    // Load fixture cases
    fixtures, err := os.ReadFile("fixtures/{resource}_resource/{resource}_config.tf")
    if err != nil {
        fmt.Print(err)
    }

    for i, fixture := range strings.Split(string(fixtures), "␞") {
        fixture{Resource}Model := {Resource}Model{}
        err = hcl.Decode(&fixture{Resource}Model, fixture)
        if err != nil {
            fmt.Printf("error decoding hcl: %s\n", err)
        }

        fixture{Resource}Model.OrgId = GetTestOrgId()

        testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
            steps: []testStep{
                {
                    config: fixture{Resource}Model,
                },
            },
        }
    }

    resourceType := "{resource}"
    for tName, tCase := range testCases {
        t.Run(tName, func(t *testing.T) {
            steps := make([]resource.TestStep, len(tCase.steps))
            for i, step := range tCase.steps {
                config := step.config
                config.OrgId = GetTestOrgId()

                f := hclwrite.NewEmptyFile()
                gohcl.EncodeIntoBody(&config, f.Body())
                combinedConfig := Render(resourceType, tName, string(f.Bytes()))

                checks := config.testChecks(t, resourceType, tName)
                chkLog := checks.string()
                stepName := fmt.Sprintf("test case %s step %d", tName, i+1)

                t.Logf("\\n// ------ begin config for %s ------\\n%s// -------- end config for %s ------\\n\\n", stepName, combinedConfig, stepName)
                t.Logf("\\n// ------ begin checks for %s ------\\n%s// -------- end checks for %s ------\\n\\n", stepName, chkLog, stepName)

                steps[i] = resource.TestStep{
                    Config: combinedConfig,
                    Check:  resource.ComposeAggregateTestCheckFunc(checks.checks...),
                }
            }

            resource.Test(t, resource.TestCase{
                ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
                Steps:                    steps,
            })
        })
    }
}

func (o *{Resource}Model) testChecks(t testing.TB, rType, tName string) testChecks {
    checks := newTestChecks(PrefixProviderName(rType) + "." + tName)

    // Required fields
    checks.append(t, "TestCheckResourceAttr", "org_id", o.OrgId)

    // Computed fields (use TestCheckResourceAttrSet)
    checks.append(t, "TestCheckResourceAttrSet", "id")
    
    // IMPLEMENT ALL FIELDS FROM PROVIDED LIST
    // Add all configurable fields from provided list here
    
    return checks
}
```

## Key Patterns & Conventions

### Provider Configuration

The provider configuration is handled through environment variables, not explicit inline configuration:

```go
// Environment variables are used automatically by the test framework:
// MIST_HOST, MIST_API_TOKEN, MIST_TEST_ORG_ID, TF_ACC

// Site configuration is added when needed:
siteConfig, siteRef := GetSiteBaseConfig(GetTestOrgId())
configStr := ""
if siteConfig != "" {
    configStr = siteConfig + "\\n\\n"
}
combinedConfig = configStr + combinedConfig
```

### Site Reference Handling

For standard resources that reference sites directly:

```go
siteConfig, siteRef := GetSiteBaseConfig(GetTestOrgId())
config := step.config

// Set site_id reference for standard site resources
if config.SiteId != nil {
    config.SiteId = &siteRef
}

f := hclwrite.NewEmptyFile()
gohcl.EncodeIntoBody(&config, f.Body())
combinedConfig := siteConfig + "\\n\\n" + Render(resourceType, tName, strings.ReplaceAll(string(f.Bytes()), fmt.Sprintf(`"%s"`, siteRef), siteRef))
```

#### Special Case: Nested Site References (like org_inventory)

For resources with nested objects containing site references:

```go
config := step.config
siteConfig, siteRef := "", ""

// Check if any nested objects need site_id and set up site config
if config.Inventory != nil {
    for key, inventoryItem := range config.Inventory {
        if inventoryItem.SiteId != nil {
            if siteConfig == "" { // only get once even if multiple items use it
                siteConfig, siteRef = GetSiteBaseConfig(GetTestOrgId())
            }
            inventoryItem.SiteId = stringPtr("{site_id}")
            config.Inventory[key] = inventoryItem
        }
    }
}

f := hclwrite.NewEmptyFile()
gohcl.EncodeIntoBody(&config, f.Body())
combinedConfig := Render(resourceType, tName, string(f.Bytes()))

configStr := ""
if siteConfig != "" {
    combinedConfig = strings.ReplaceAll(combinedConfig, "\\"{site_id}\\"", siteRef)
    configStr = siteConfig + "\\n\\n"
}
combinedConfig = configStr + combinedConfig
```

### Validation Patterns

- **Computed fields**: `TestCheckResourceAttrSet` (verifies field exists and has value)
- **Configurable fields**: `TestCheckResourceAttr` (verifies exact expected value)
- **Maps**: Validate count with `.%`, individual keys with `.{key}.{field}`
- **Lists**: Validate count with `.#`, individual items with `.{index}.{field}`

### Skip Patterns

For tests requiring real devices/external data (use sparingly - prefer real data):

```go
// Optional skip pattern - prefer running with real data when available
if strings.HasPrefix(tName, "fixture_case") {
    t.Skip("Skipping fixture case as it requires real devices.")
}
```

### Environment Variables

Standard test environment variables (all must be set):

- `TF_ACC=1`: Enable Terraform acceptance tests
- `MIST_HOST`: API host (e.g., 'api.mistsys.com')
- `MIST_API_TOKEN`: API token (retrieve from your org settings in MIST)
- `MIST_TEST_ORG_ID`: Test organization ID (retrieve from your org settings in MIST)

## Testing Guidelines

### Validation Completeness

1. **All schema fields**: Every field in the resource schema should have corresponding validation
2. **Field types**: Use appropriate validation methods for each data type
3. **Required vs Optional**: Ensure required fields are always validated
4. **Complex structures**: Properly validate nested objects, maps, and lists

### Test Structure

1. **Simple case**: Minimal valid configuration
2. **Fixture cases**: Comprehensive realistic configurations from fixture file
3. **Multiple steps**: Use multiple test steps for update scenarios if needed

### Error Handling

- Log configurations and checks for debugging
- Use descriptive test case names
- Handle HCL decoding errors gracefully

## Critical Success Factors

### Field Coverage Verification

- **100% accuracy**: Use reflection-based field extraction from field_extraction_guidelines.md
- **Source of truth**: Treat provided field list as definitive
- **Complete validation**: Every field must have appropriate test validation
- **Proper categorization**: Computed vs configurable field handling

### Pattern Compliance

- **Exact structure**: Follow site_wlan_resource_test.go pattern precisely for standard resources
- **No deviations**: Use established conventions without modifications
- **Integration ready**: Include provider configuration and site handling
- **Fixtures**: Cover ALL configurable fields with realistic values
- **Specialized cases**: Use org_inventory_resource_test.go pattern for nested site references

## Common Pitfalls to Avoid

1. **Don't try to manually add missing fields** to generated test structs - this will be overwritten
2. **Check for missing complete structs** - field coverage can be significantly improved by adding complete missing struct definitions that exist in the resource schema
3. **Expect varying coverage levels** - coverage depends on generated struct completeness (30-90% range is normal)
4. **Don't forget to add `hcl` tags** to nested struct fields that have `cty` tags
5. **Don't use invalid UUIDs** in fixtures for `site_ids`, `sitegroup_ids` - use valid UUID format
6. **Don't start with complex fixtures** - begin simple and add complexity gradually
7. **Don't skip simplification** when complex nested objects cause provider SDK issues
8. **Don't hardcode site IDs** - use `GetSiteBaseConfig()` pattern when needed
9. **Don't forget provider configuration** - environment variables handle authentication automatically
10. **Don't use wrong validation methods** - computed vs configurable field validation differs
11. **Don't assume all field combinations work** - some complex configurations may cause provider inconsistencies
12. **Don't use fake data when real data is available** - prefer real MAC addresses, device IDs, etc.
13. **Don't create duplicate site names** - ensure test sites are cleaned up or use unique names
14. **Use longer timeouts for tests** - Real API integration tests need time. Use `-timeout=10m` for complex resources
15. **DO NOT FLUFF UP VARIABLE, FUNCTION OR FIELD NAMES WITH "COMPREHENSIVE" PREFIXES AND OTHER NONSENSE!**

**Most Important: Accept that test coverage will be limited by the generated test structs. Focus on testing available fields thoroughly rather than achieving 100% coverage.**

This guide ensures consistent, comprehensive test coverage following all established patterns in the terraform-provider-mist codebase with complete field validation based on the authoritative field list.
