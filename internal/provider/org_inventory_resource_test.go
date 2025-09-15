package provider

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"sort"
	"testing"

	resource_org_inventory "github.com/Juniper/terraform-provider-mist/internal/resource_org_inventory"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgInventory(t *testing.T) {
	type testStep struct {
		config OrgInventoryModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgInventoryModel{
						OrgId: GetTestOrgId(),
					},
				},
			},
		},
	}

	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				config := step.config

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				configStr := Render("org_inventory", tName, string(f.Bytes()))

				checks := config.testChecks(t, PrefixProviderName("org_inventory"), tName)
				chkLog := checks.string()
				stepName := fmt.Sprintf("test case %s step %d", tName, i+1)

				// log config and checks here
				t.Logf("\n// ------ begin config for %s ------\n%s// -------- end config for %s ------\n\n", stepName, configStr, stepName)
				t.Logf("\n// ------ begin checks for %s ------\n%s// -------- end checks for %s ------\n\n", stepName, chkLog, stepName)

				steps[i] = resource.TestStep{
					Config: configStr,
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

func (o *OrgInventoryModel) testChecks(t testing.TB, rType, rName string) testChecks {
	checks := newTestChecks(rType + "." + rName)

	// Check required fields
	checks.append(t, "TestCheckResourceAttr", "org_id", o.OrgId)

	return checks
}

// Temporary helper functions for schema extraction
func TestExtractSchema(t *testing.T) {
	ctx := context.Background()

	// Get the schema for org_inventory resource
	schemaObj := resource_org_inventory.OrgInventoryResourceSchema(ctx)

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
