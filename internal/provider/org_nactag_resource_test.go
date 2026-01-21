package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_nactag"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgNactagModel(t *testing.T) {
	type testStep struct {
		config OrgNactagModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgNactagModel{
						OrgId: GetTestOrgId(),
						Name:  "test-nactag",
						Type:  "vlan",
						Vlan:  stringPtr("10"),
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/org_nactag_resource/org_nactag_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b)
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		var FixtureOrgNactagModel OrgNactagModel
		err = hcl.Decode(&FixtureOrgNactagModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureOrgNactagModel,
				},
			},
		}
	}

	resourceType := "org_nactag"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_org_nactag.OrgNactagResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {

			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				config := step.config
				config.OrgId = GetTestOrgId()

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				combinedConfig := Render(resourceType, tName, string(f.Bytes()))

				checks := config.testChecks(t, resourceType, tName, tracker)
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
	tracker.FieldCoverageReport(t)
}

func (o *OrgNactagModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + tName)
	checks.SetTracker(tracker)

	// Computed fields
	checks.append(t, "TestCheckResourceAttrSet", "id")

	// Required fields
	checks.append(t, "TestCheckResourceAttr", "org_id", o.OrgId)
	checks.append(t, "TestCheckResourceAttr", "name", o.Name)
	checks.append(t, "TestCheckResourceAttr", "type", o.Type)

	// Optional configurable fields
	if o.AllowUsermacOverride != nil {
		checks.append(t, "TestCheckResourceAttr", "allow_usermac_override", fmt.Sprintf("%t", *o.AllowUsermacOverride))
	}

	if len(o.EgressVlanNames) > 0 {
		checks.append(t, "TestCheckResourceAttr", "egress_vlan_names.#", fmt.Sprintf("%d", len(o.EgressVlanNames)))
		for i, vlanName := range o.EgressVlanNames {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("egress_vlan_names.%d", i), vlanName)
		}
	}

	if o.GbpTag != nil {
		checks.append(t, "TestCheckResourceAttr", "gbp_tag", *o.GbpTag)
	}

	if o.Match != nil {
		checks.append(t, "TestCheckResourceAttr", "match", *o.Match)
	}

	if o.MatchAll != nil {
		checks.append(t, "TestCheckResourceAttr", "match_all", fmt.Sprintf("%t", *o.MatchAll))
	}

	if o.NacportalId != nil {
		checks.append(t, "TestCheckResourceAttr", "nacportal_id", *o.NacportalId)
	}

	if len(o.RadiusAttrs) > 0 {
		checks.append(t, "TestCheckResourceAttr", "radius_attrs.#", fmt.Sprintf("%d", len(o.RadiusAttrs)))
		for i, attr := range o.RadiusAttrs {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radius_attrs.%d", i), attr)
		}
	}

	if o.RadiusGroup != nil {
		checks.append(t, "TestCheckResourceAttr", "radius_group", *o.RadiusGroup)
	}

	if len(o.RadiusVendorAttrs) > 0 {
		checks.append(t, "TestCheckResourceAttr", "radius_vendor_attrs.#", fmt.Sprintf("%d", len(o.RadiusVendorAttrs)))
		for i, attr := range o.RadiusVendorAttrs {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radius_vendor_attrs.%d", i), attr)
		}
	}

	if o.SessionTimeout != nil {
		checks.append(t, "TestCheckResourceAttr", "session_timeout", fmt.Sprintf("%d", *o.SessionTimeout))
	}

	if o.UsernameAttr != nil {
		checks.append(t, "TestCheckResourceAttr", "username_attr", *o.UsernameAttr)
	}

	if len(o.Values) > 0 {
		checks.append(t, "TestCheckResourceAttr", "values.#", fmt.Sprintf("%d", len(o.Values)))
		for i, value := range o.Values {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("values.%d", i), value)
		}
	}

	if o.Vlan != nil {
		checks.append(t, "TestCheckResourceAttr", "vlan", *o.Vlan)
	}

	return checks
}
