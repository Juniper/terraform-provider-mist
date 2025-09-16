# Comprehensive Test Implementation Guide for Terraform Provider Mist Resources

## Overview
This guide provides a complete template for implementing comprehensive test coverage for any resource in the terraform-provider-mist codebase, following established patterns and conventions.

## Part 1: Field Analysis & Discovery

### Extract Complete Schema Fields and Currently Tested Fields

**Follow the field extraction methodology from `/internal/provider/fixtures/TESTING_GUIDELINES.md`:**

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

### Step 1: Test Structs Creation

Create `/internal/provider/{resource}_test_structs.go` with **ONLY configurable fields** from the provided field list:

```go
package provider

import (
    "github.com/hashicorp/terraform-plugin-framework/types"
)

type {Resource}Model struct {
    // Include ONLY configurable fields for test manipulation
    // Use dual tagging: `cty:"{field_name}" hcl:"{field_name}"`
    OrgId  string                    `cty:"org_id" hcl:"org_id"`
    Field1 *string                   `cty:"field1" hcl:"field1"`
    Field2 *bool                     `cty:"field2" hcl:"field2"`
    // Map/complex types as needed based on provided field list
}

// Include nested struct types if needed based on provided field list
type NestedValue struct {
    SubField1 *string `cty:"sub_field1" hcl:"sub_field1"`
    SubField2 *int64  `cty:"sub_field2" hcl:"sub_field2"`
}
```

**Key Conventions:**

- Use pointer types (`*string`, `*bool`, `*int64`) for optional fields
- Include `cty` and `hcl` tags for both configuration parsing and encoding
- Only include configurable fields (not computed/read-only fields)
- Follow the exact field list provided as source of truth

### Step 2: Fixture Configuration

Create `/internal/provider/fixtures/{resource}_resource/{resource}_config.tf` with comprehensive coverage:

```hcl
# Comprehensive realistic configuration covering all configurable fields
field1 = "test_value"
field2 = true
nested_field = {
  sub_field1 = "nested_value"
  sub_field2 = 123
}
```

**Conventions:**

- No terraform resource boilerplate (added by test framework)
- Use realistic test values for maximum coverage
- Set boolean fields to `true` for comprehensive testing
- Include all configurable fields from the provided list
- Separate multiple configs with `␞` delimiter
- For resources requiring real devices/data, use real values when available (like MAC addresses from actual inventory)

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
                        OrgId: GetTestOrgId(),
                        // Minimal required configuration
                    },
                },
            },
        },
    }

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
            // Skip if requires real devices/data (optional)
            // if strings.HasPrefix(tName, "fixture_case") {
            //     t.Skip("Skipping fixture case as it requires real devices.")
            // }

            steps := make([]resource.TestStep, len(tCase.steps))
            for i, step := range tCase.steps {
                siteConfig, siteRef := GetSiteBaseConfig(GetTestOrgId())
                config := step.config

                // Set site_id reference for standard site resources
                if config.SiteId != nil {
                    config.SiteId = &siteRef
                }

                f := hclwrite.NewEmptyFile()
                gohcl.EncodeIntoBody(&config, f.Body())
                combinedConfig := siteConfig + "\n\n" + Render(resourceType, tName, strings.ReplaceAll(string(f.Bytes()), fmt.Sprintf(`"%s"`, siteRef), siteRef))

                checks := config.testChecks(t, "{resource}", tName)
                chkLog := checks.string()
                stepName := fmt.Sprintf("test case %s step %d", tName, i+1)

                t.Logf("\n// ------ begin config for %s ------\n%s// -------- end config for %s ------\n\n", stepName, combinedConfig, stepName)
                t.Logf("\n// ------ begin checks for %s ------\n%s// -------- end checks for %s ------\n\n", stepName, chkLog, stepName)

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

    // IMPLEMENT ALL FIELDS FROM PROVIDED LIST
    // Computed fields (use TestCheckResourceAttrSet)
    checks.append(t, "TestCheckResourceAttrSet", "id")
    // Add all computed fields from provided list here

    // Configurable fields (use TestCheckResourceAttr with expected values)
    // Add all configurable fields from provided list here
    
    // Example patterns for different field types:
    // if o.StringField != nil {
    //     checks.append(t, "TestCheckResourceAttr", "string_field", *o.StringField)
    // }
    // if o.BoolField != nil {
    //     checks.append(t, "TestCheckResourceAttr", "bool_field", fmt.Sprintf("%t", *o.BoolField))
    // }
    // if o.IntField != nil {
    //     checks.append(t, "TestCheckResourceAttr", "int_field", fmt.Sprintf("%d", *o.IntField))
    // }

    // Handle complex types (maps, lists, nested objects)
    // For maps:
    // if len(o.MapField) > 0 {
    //     checks.append(t, "TestCheckResourceAttr", "map_field.%", fmt.Sprintf("%d", len(o.MapField)))
    //     for key, value := range o.MapField {
    //         checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("map_field.%s.sub_field", key), *value.SubField)
    //     }
    // }

    // For lists:
    // if len(o.ListField) > 0 {
    //     checks.append(t, "TestCheckResourceAttr", "list_field.#", fmt.Sprintf("%d", len(o.ListField)))
    //     for i, item := range o.ListField {
    //         checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("list_field.%d.sub_field", i), *item.SubField)
    //     }
    // }

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
    configStr = siteConfig + "\n\n"
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
combinedConfig := siteConfig + "\n\n" + Render(resourceType, tName, strings.ReplaceAll(string(f.Bytes()), fmt.Sprintf(`"%s"`, siteRef), siteRef))
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
            // Set placeholder for site_id in nested object
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
    combinedConfig = strings.ReplaceAll(combinedConfig, "\"{site_id}\"", siteRef)
    configStr = siteConfig + "\n\n"
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
    t.Skip("Skipping fixture case as it requires real devices with valid MAC addresses.")
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

## Example Usage Prompts

### Part 1 Prompt: Field Analysis

```
Please extract complete schema fields and currently tested fields for {RESOURCE_NAME} following the methodology in `/internal/provider/fixtures/TESTING_GUIDELINES.md` and internal/provider/fixtures/COMPREHENSIVE_TEST_IMPLEMENTATION_GUIDE.md:

1. Extract all schema fields from /internal/{resource}/{resource}_resource_gen.go using the automated reflection-based method
2. Extract currently tested fields from test execution output
3. Generate two files:
   - `all_schema_fields.txt` - Complete schema field list with proper dot notation
   - `current_tested_fields.txt` - Currently validated fields in tests
4. Identify which fields are computed vs configurable

Use the reflection-based extraction method from TESTING_GUIDELINES.md for 100% accuracy with nested fields.
```

### Part 2 Prompt: Test Implementation (Use after receiving field list)

```
Please implement comprehensive test coverage for {RESOURCE_NAME} using the provided field list as the source of truth.

FIELD LIST PROVIDED:
[User provides the definitive field list here]

IMPLEMENTATION REQUIREMENTS:
1. Create test structs in /internal/provider/{resource}_test_structs.go with ONLY configurable fields from the provided list
2. Create comprehensive fixture in /internal/provider/fixtures/{resource}_resource/{resource}_config.tf covering all configurable fields
3. Implement main test in /internal/provider/{resource}_resource_test.go following the EXACT SAME PATTERN as site_wlan_resource_test.go (or org_inventory_resource_test.go for nested site references)
4. Ensure comprehensive validation of ALL fields from the provided list using appropriate TestCheck methods:
   - Computed fields: TestCheckResourceAttrSet
   - Configurable fields: TestCheckResourceAttr with expected values
5. Handle complex types (maps, lists, nested objects) with proper validation patterns
6. Add skip logic if the resource requires real external data
7. Include provider configuration for integration testing
8. Use GetSiteBaseConfig() pattern for site references

Follow the established patterns exactly - do not deviate from the template structure.
```

## Key Patterns & Conventions

### Provider Configuration (Implementation Pattern)
Environment variables handle provider authentication automatically:

```go
// These environment variables are used by the test framework:
// TF_ACC=1, MIST_HOST, MIST_API_TOKEN, MIST_TEST_ORG_ID

// Site configuration is conditionally added:
configStr := ""
if siteConfig != "" {
    configStr = siteConfig + "\n\n"
}
combinedConfig = configStr + combinedConfig
```

### Site Reference Handling (Nested Objects)
For resources with nested objects containing site references:

```go
// Pattern for nested objects like inventory maps:
if config.Inventory != nil {
    for key, inventoryItem := range config.Inventory {
        if inventoryItem.SiteId != nil {
            if siteConfig == "" {
                siteConfig, siteRef = GetSiteBaseConfig(GetTestOrgId())
            }
            inventoryItem.SiteId = stringPtr("{site_id}")
            config.Inventory[key] = inventoryItem
        }
    }
}

// After HCL encoding, replace placeholders:
combinedConfig = strings.ReplaceAll(combinedConfig, "\"{site_id}\"", siteRef)
```

### Validation Patterns

- **Computed fields**: `TestCheckResourceAttrSet` (verifies field exists and has value)
- **Configurable fields**: `TestCheckResourceAttr` (verifies exact expected value)
- **Maps**: Validate count with `.%`, individual keys with `.{key}.{field}`
- **Lists**: Validate count with `.#`, individual items with `.{index}.{field}`

### Skip Patterns
For tests requiring real devices/external data:

```go
t.Skip("Skipping by default as test requires real devices/data.")
```

### Environment Variables
Standard test environment variables:

- `MIST_HOST`: API host
- `MIST_API_TOKEN`: API token
- `MIST_TEST_ORG_ID`: Test organization ID

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

- **100% accuracy**: Use reflection-based field extraction from TESTING_GUIDELINES.md
- **Source of truth**: Treat provided field list as definitive
- **Complete validation**: Every field must have appropriate test validation
- **Proper categorization**: Computed vs configurable field handling

### Pattern Compliance

- **Exact structure**: Follow site_wlan_resource_test.go pattern precisely for standard resources
- **No deviations**: Use established conventions without modifications
- **Integration ready**: Include provider configuration and site handling
- **Comprehensive fixtures**: Cover all configurable fields with realistic values
- **Specialized cases**: Use org_inventory_resource_test.go pattern for nested site references

## Common Pitfalls to Avoid

1. **Don't include computed fields** in test structs - they can't be configured
2. **Don't forget dual tags** (`cty` and `hcl`) on struct fields
3. **Don't hardcode site IDs** - use `GetSiteBaseConfig()` pattern
4. **Don't skip provider configuration** - environment variables handle authentication automatically
5. **Don't use wrong validation methods** - computed vs configurable field validation differs
6. **Don't forget complex type validation** - maps and lists need count + individual item validation
7. **Don't deviate from provided field list** - it is the source of truth
8. **Don't skip field coverage verification** - ensure 100% validation coverage
9. **Don't use fake data when real data is available** - prefer real MAC addresses, device IDs, etc.
10. **Don't create duplicate site names** - ensure test sites are cleaned up or use unique names
11. **Ask user to set all required environment variables if you find them unset** - TF_ACC, MIST_HOST, MIST_API_TOKEN, MIST_TEST_ORG_ID

This guide ensures consistent, comprehensive test coverage following all established patterns in the terraform-provider-mist codebase with complete field validation based on the authoritative field list.
