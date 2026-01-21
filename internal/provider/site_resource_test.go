package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_site"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestSiteModel(t *testing.T) {
	type testStep struct {
		config SiteModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: SiteModel{
						Address: "test-address",
						Name:    "test-site",
						OrgId:   GetTestOrgId(),
					},
				},
			},
		},
	}

	// Load fixture data
	b, err := os.ReadFile("fixtures/site_resource/site_resource_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b)
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		fixtureSiteModel := SiteModel{}
		err = hcl.Decode(&fixtureSiteModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		fixtureSiteModel.OrgId = GetTestOrgId()

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: fixtureSiteModel,
				},
			},
		}
	}

	resourceType := "site"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_site.SiteResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				// Generate Terraform configuration using automated HCL generation
				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&step.config, f.Body())
				configStr := Render("site", tName, string(f.Bytes()))

				checks := step.config.testChecks(t, resourceType, tName, tracker)
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
	if tracker != nil {
		tracker.FieldCoverageReport(t)
	}
}

func (s *SiteModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + tName)
	checks.SetTracker(tracker)

	// Check fields in struct order
	// 1. Address (required)
	checks.append(t, "TestCheckResourceAttr", "address", s.Address)

	// 2. AlarmtemplateId (optional)
	if s.AlarmtemplateId != nil {
		checks.append(t, "TestCheckResourceAttr", "alarmtemplate_id", *s.AlarmtemplateId)
	}

	// 3. AptemplateId (optional)
	if s.AptemplateId != nil {
		checks.append(t, "TestCheckResourceAttr", "aptemplate_id", *s.AptemplateId)
	}

	// 4. CountryCode (optional)
	if s.CountryCode != nil {
		checks.append(t, "TestCheckResourceAttr", "country_code", *s.CountryCode)
	}

	// 5. GatewaytemplateId (optional)
	if s.GatewaytemplateId != nil {
		checks.append(t, "TestCheckResourceAttr", "gatewaytemplate_id", *s.GatewaytemplateId)
	}

	// 6. Id (computed-only)
	checks.append(t, "TestCheckResourceAttrSet", "id")

	// 7. Latlng (optional nested object) - test child attributes only
	if s.Latlng != nil {
		checks.append(t, "TestCheckResourceAttr", "latlng.lat", fmt.Sprintf("%g", s.Latlng.Lat))
		checks.append(t, "TestCheckResourceAttr", "latlng.lng", fmt.Sprintf("%g", s.Latlng.Lng))
	}

	// 8. Name (required)
	checks.append(t, "TestCheckResourceAttr", "name", s.Name)

	// 9. NetworktemplateId (optional)
	if s.NetworktemplateId != nil {
		checks.append(t, "TestCheckResourceAttr", "networktemplate_id", *s.NetworktemplateId)
	}

	// 10. Notes (optional)
	if s.Notes != nil {
		checks.append(t, "TestCheckResourceAttr", "notes", *s.Notes)
	}

	// 11. OrgId (required)
	checks.append(t, "TestCheckResourceAttr", "org_id", s.OrgId)

	// 12. RftemplateId (optional)
	if s.RftemplateId != nil {
		checks.append(t, "TestCheckResourceAttr", "rftemplate_id", *s.RftemplateId)
	}

	// 13. SecpolicyId (optional)
	if s.SecpolicyId != nil {
		checks.append(t, "TestCheckResourceAttr", "secpolicy_id", *s.SecpolicyId)
	}

	// 14. SitegroupIds (optional array)
	if len(s.SitegroupIds) > 0 {
		checks.append(t, "TestCheckResourceAttr", "sitegroup_ids.#", fmt.Sprintf("%d", len(s.SitegroupIds)))
		for i, id := range s.SitegroupIds {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("sitegroup_ids.%d", i), id)
		}
	}

	// 15. SitetemplateId (optional)
	if s.SitetemplateId != nil {
		checks.append(t, "TestCheckResourceAttr", "sitetemplate_id", *s.SitetemplateId)
	}

	// 16. Timezone (optional)
	if s.Timezone != nil {
		checks.append(t, "TestCheckResourceAttr", "timezone", *s.Timezone)
	}

	// 17. Tzoffset (computed-only)
	checks.append(t, "TestCheckResourceAttrSet", "tzoffset")

	return checks
}
