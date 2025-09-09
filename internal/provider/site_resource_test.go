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

	b, err := os.ReadFile("fixtures/site_resource/site_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {

		var FixtureSiteModel SiteModel
		err = hcl.Decode(&FixtureSiteModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		FixtureSiteModel.OrgId = GetTestOrgId()

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureSiteModel,
				},
			},
		}
	}

	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			resourceType := "site"

			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				config := step.config

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				configStr := Render(resourceType, tName, string(f.Bytes()))

				checks := config.testChecks(t, resourceType, tName)
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

func (s *SiteModel) testChecks(t testing.TB, rType, rName string) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + rName)
	checks.append(t, "TestCheckResourceAttr", "address", s.Address)
	checks.append(t, "TestCheckResourceAttr", "name", s.Name)
	checks.append(t, "TestCheckResourceAttr", "org_id", s.OrgId)

	// Conditional checks for optional parameters
	if s.AlarmtemplateId != nil {
		checks.append(t, "TestCheckResourceAttr", "alarmtemplate_id", *s.AlarmtemplateId)
	}
	if s.AptemplateId != nil {
		checks.append(t, "TestCheckResourceAttr", "aptemplate_id", *s.AptemplateId)
	}
	if s.CountryCode != nil {
		checks.append(t, "TestCheckResourceAttr", "country_code", *s.CountryCode)
	}
	if s.GatewaytemplateId != nil {
		checks.append(t, "TestCheckResourceAttr", "gatewaytemplate_id", *s.GatewaytemplateId)
	}
	if s.Latlng != nil {
		checks.append(t, "TestCheckResourceAttr", "latlng.lat", fmt.Sprintf("%g", s.Latlng.Lat))
		checks.append(t, "TestCheckResourceAttr", "latlng.lng", fmt.Sprintf("%g", s.Latlng.Lng))
	}
	if s.NetworktemplateId != nil {
		checks.append(t, "TestCheckResourceAttr", "networktemplate_id", *s.NetworktemplateId)
	}
	if s.Notes != nil {
		checks.append(t, "TestCheckResourceAttr", "notes", *s.Notes)
	}
	if s.RftemplateId != nil {
		checks.append(t, "TestCheckResourceAttr", "rftemplate_id", *s.RftemplateId)
	}
	if s.SecpolicyId != nil {
		checks.append(t, "TestCheckResourceAttr", "secpolicy_id", *s.SecpolicyId)
	}
	if len(s.SitegroupIds) > 0 {
		for i, id := range s.SitegroupIds {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("sitegroup_ids.%d", i), id)
		}
	}
	if s.SitetemplateId != nil {
		checks.append(t, "TestCheckResourceAttr", "sitetemplate_id", *s.SitetemplateId)
	}
	if s.Timezone != nil {
		checks.append(t, "TestCheckResourceAttr", "timezone", *s.Timezone)
	}

	return checks
}
