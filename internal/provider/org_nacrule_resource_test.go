package provider

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
	"testing"

	resource_org_nacrule "github.com/Juniper/terraform-provider-mist/internal/resource_org_nacrule"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgNacruleModel(t *testing.T) {
	type testStep struct {
		config OrgNacruleModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgNacruleModel{
						OrgId:  GetTestOrgId(),
						Name:   "test-nacrule",
						Action: "allow",
						Order:  1,
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/org_nacrule_resource/org_nacrule_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		var FixtureOrgNacruleModel OrgNacruleModel
		err = hcl.Decode(&FixtureOrgNacruleModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureOrgNacruleModel,
				},
			},
		}
	}

	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			resourceType := "org_nacrule"

			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				config := step.config
				config.OrgId = GetTestOrgId()
				siteConfig, siteRef := "", ""

				// Check if config needs site_id and set up site config
				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				configStr := string(f.Bytes())

				if strings.Contains(configStr, "{site_id}") || strings.Contains(configStr, "{sitegroup_id}") {
					siteConfig, siteRef = GetSiteBaseConfig(GetTestOrgId())
					// For now, use the same site ID for sitegroup_id as a placeholder
					configStr = strings.ReplaceAll(configStr, "\"{site_id}\"", siteRef)
					configStr = strings.ReplaceAll(configStr, "\"{sitegroup_id}\"", siteRef)
				}

				combinedConfig := Render(resourceType, tName, configStr)

				if siteConfig != "" {
					combinedConfig = siteConfig + "\n\n" + combinedConfig
				}

				checks := config.testChecks(t, resourceType, tName)
				chkLog := checks.string()
				stepName := fmt.Sprintf("test case %s step %d", tName, i+1)

				// log config and checks here
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

func (o *OrgNacruleModel) testChecks(t testing.TB, rType, rName string) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + rName)

	// Computed fields
	checks.append(t, "TestCheckResourceAttrSet", "id")

	// Required fields
	checks.append(t, "TestCheckResourceAttr", "org_id", o.OrgId)
	checks.append(t, "TestCheckResourceAttr", "name", o.Name)
	checks.append(t, "TestCheckResourceAttr", "action", o.Action)
	checks.append(t, "TestCheckResourceAttr", "order", fmt.Sprintf("%d", o.Order))

	// Optional simple fields
	if len(o.ApplyTags) > 0 {
		checks.append(t, "TestCheckResourceAttr", "apply_tags.#", fmt.Sprintf("%d", len(o.ApplyTags)))
		for i, tag := range o.ApplyTags {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("apply_tags.%d", i), tag)
		}
	}

	if o.Enabled != nil {
		checks.append(t, "TestCheckResourceAttr", "enabled", fmt.Sprintf("%t", *o.Enabled))
	}

	if o.GuestAuthState != nil {
		checks.append(t, "TestCheckResourceAttr", "guest_auth_state", *o.GuestAuthState)
	}

	// Matching nested object
	if o.Matching != nil {
		if o.Matching.AuthType != nil {
			checks.append(t, "TestCheckResourceAttr", "matching.auth_type", *o.Matching.AuthType)
		}
		if len(o.Matching.Family) > 0 {
			checks.append(t, "TestCheckResourceAttr", "matching.family.#", fmt.Sprintf("%d", len(o.Matching.Family)))
			for i, family := range o.Matching.Family {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("matching.family.%d", i), family)
			}
		}
		if len(o.Matching.Mfg) > 0 {
			checks.append(t, "TestCheckResourceAttr", "matching.mfg.#", fmt.Sprintf("%d", len(o.Matching.Mfg)))
			for i, mfg := range o.Matching.Mfg {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("matching.mfg.%d", i), mfg)
			}
		}
		if len(o.Matching.Model) > 0 {
			checks.append(t, "TestCheckResourceAttr", "matching.model.#", fmt.Sprintf("%d", len(o.Matching.Model)))
			for i, model := range o.Matching.Model {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("matching.model.%d", i), model)
			}
		}
		if len(o.Matching.Nactags) > 0 {
			checks.append(t, "TestCheckResourceAttr", "matching.nactags.#", fmt.Sprintf("%d", len(o.Matching.Nactags)))
			for i, nactag := range o.Matching.Nactags {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("matching.nactags.%d", i), nactag)
			}
		}
		if len(o.Matching.OsType) > 0 {
			checks.append(t, "TestCheckResourceAttr", "matching.os_type.#", fmt.Sprintf("%d", len(o.Matching.OsType)))
			for i, osType := range o.Matching.OsType {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("matching.os_type.%d", i), osType)
			}
		}
		if len(o.Matching.PortTypes) > 0 {
			checks.append(t, "TestCheckResourceAttr", "matching.port_types.#", fmt.Sprintf("%d", len(o.Matching.PortTypes)))
			for i, portType := range o.Matching.PortTypes {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("matching.port_types.%d", i), portType)
			}
		}
		if len(o.Matching.SiteIds) > 0 {
			checks.append(t, "TestCheckResourceAttr", "matching.site_ids.#", fmt.Sprintf("%d", len(o.Matching.SiteIds)))
			for i, siteId := range o.Matching.SiteIds {
				if siteId == "{site_id}" {
					checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("matching.site_ids.%d", i))
				} else {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("matching.site_ids.%d", i), siteId)
				}
			}
		}
		if len(o.Matching.SitegroupIds) > 0 {
			checks.append(t, "TestCheckResourceAttr", "matching.sitegroup_ids.#", fmt.Sprintf("%d", len(o.Matching.SitegroupIds)))
			for i, sitegroupId := range o.Matching.SitegroupIds {
				if sitegroupId == "{sitegroup_id}" {
					checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("matching.sitegroup_ids.%d", i))
				} else {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("matching.sitegroup_ids.%d", i), sitegroupId)
				}
			}
		}
		if len(o.Matching.Vendor) > 0 {
			checks.append(t, "TestCheckResourceAttr", "matching.vendor.#", fmt.Sprintf("%d", len(o.Matching.Vendor)))
			for i, vendor := range o.Matching.Vendor {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("matching.vendor.%d", i), vendor)
			}
		}
	}

	// NotMatching nested object
	if o.NotMatching != nil {
		if o.NotMatching.AuthType != nil {
			checks.append(t, "TestCheckResourceAttr", "not_matching.auth_type", *o.NotMatching.AuthType)
		}
		if len(o.NotMatching.Family) > 0 {
			checks.append(t, "TestCheckResourceAttr", "not_matching.family.#", fmt.Sprintf("%d", len(o.NotMatching.Family)))
			for i, family := range o.NotMatching.Family {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("not_matching.family.%d", i), family)
			}
		}
		if len(o.NotMatching.Mfg) > 0 {
			checks.append(t, "TestCheckResourceAttr", "not_matching.mfg.#", fmt.Sprintf("%d", len(o.NotMatching.Mfg)))
			for i, mfg := range o.NotMatching.Mfg {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("not_matching.mfg.%d", i), mfg)
			}
		}
		if len(o.NotMatching.Model) > 0 {
			checks.append(t, "TestCheckResourceAttr", "not_matching.model.#", fmt.Sprintf("%d", len(o.NotMatching.Model)))
			for i, model := range o.NotMatching.Model {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("not_matching.model.%d", i), model)
			}
		}
		if len(o.NotMatching.Nactags) > 0 {
			checks.append(t, "TestCheckResourceAttr", "not_matching.nactags.#", fmt.Sprintf("%d", len(o.NotMatching.Nactags)))
			for i, nactag := range o.NotMatching.Nactags {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("not_matching.nactags.%d", i), nactag)
			}
		}
		if len(o.NotMatching.OsType) > 0 {
			checks.append(t, "TestCheckResourceAttr", "not_matching.os_type.#", fmt.Sprintf("%d", len(o.NotMatching.OsType)))
			for i, osType := range o.NotMatching.OsType {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("not_matching.os_type.%d", i), osType)
			}
		}
		if len(o.NotMatching.PortTypes) > 0 {
			checks.append(t, "TestCheckResourceAttr", "not_matching.port_types.#", fmt.Sprintf("%d", len(o.NotMatching.PortTypes)))
			for i, portType := range o.NotMatching.PortTypes {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("not_matching.port_types.%d", i), portType)
			}
		}
		if len(o.NotMatching.SiteIds) > 0 {
			checks.append(t, "TestCheckResourceAttr", "not_matching.site_ids.#", fmt.Sprintf("%d", len(o.NotMatching.SiteIds)))
			for i, siteId := range o.NotMatching.SiteIds {
				if siteId == "{site_id}" {
					checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("not_matching.site_ids.%d", i))
				} else {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("not_matching.site_ids.%d", i), siteId)
				}
			}
		}
		if len(o.NotMatching.SitegroupIds) > 0 {
			checks.append(t, "TestCheckResourceAttr", "not_matching.sitegroup_ids.#", fmt.Sprintf("%d", len(o.NotMatching.SitegroupIds)))
			for i, sitegroupId := range o.NotMatching.SitegroupIds {
				if sitegroupId == "{sitegroup_id}" {
					checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("not_matching.sitegroup_ids.%d", i))
				} else {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("not_matching.sitegroup_ids.%d", i), sitegroupId)
				}
			}
		}
		if len(o.NotMatching.Vendor) > 0 {
			checks.append(t, "TestCheckResourceAttr", "not_matching.vendor.#", fmt.Sprintf("%d", len(o.NotMatching.Vendor)))
			for i, vendor := range o.NotMatching.Vendor {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("not_matching.vendor.%d", i), vendor)
			}
		}
	}

	return checks
}

// Field extraction functions for comprehensive test coverage analysis
func TestExtractSchema(t *testing.T) {
	ctx := context.Background()

	// Get the schema
	schemaObj := resource_org_nacrule.OrgNacruleResourceSchema(ctx)

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
