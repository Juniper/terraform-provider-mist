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

func TestOrgWlantemplateModel(t *testing.T) {
	type testStep struct {
		config OrgWlantemplateModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgWlantemplateModel{
						Name:  "TestTemplate",
						OrgId: GetTestOrgId(),
					},
				},
			},
		},
	}

	// Load fixture data following the org_wlan pattern
	b, err := os.ReadFile("fixtures/org_wlantemplate_resource/org_wlantemplate_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		fixtureOrgWlantemplateModel := OrgWlantemplateModel{}
		err = hcl.Decode(&fixtureOrgWlantemplateModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		fixtureOrgWlantemplateModel.OrgId = GetTestOrgId()

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: fixtureOrgWlantemplateModel,
				},
			},
		}
	}

	resourceType := "org_wlantemplate"
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				// Generate Terraform configuration using automated HCL generation
				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&step.config, f.Body())
				combinedConfig := Render("org_wlantemplate", tName, string(f.Bytes()))

				checks := step.config.testChecks(t, resourceType, tName)
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

func (s *OrgWlantemplateModel) testChecks(t testing.TB, rType, tName string) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + tName)

	// Check fields in struct order
	// 1. Applies (nested object)
	if s.Applies != nil {
		if s.Applies.OrgId != nil {
			checks.append(t, "TestCheckResourceAttr", "applies.org_id", *s.Applies.OrgId)
		}

		// applies.site_ids array
		if len(s.Applies.SiteIds) > 0 {
			checks.append(t, "TestCheckResourceAttr", "applies.site_ids.#", fmt.Sprintf("%d", len(s.Applies.SiteIds)))
			for i, id := range s.Applies.SiteIds {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("applies.site_ids.%d", i), id)
			}
		}

		// applies.sitegroup_ids array
		if len(s.Applies.SitegroupIds) > 0 {
			checks.append(t, "TestCheckResourceAttr", "applies.sitegroup_ids.#", fmt.Sprintf("%d", len(s.Applies.SitegroupIds)))
			for i, id := range s.Applies.SitegroupIds {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("applies.sitegroup_ids.%d", i), id)
			}
		}
	}

	// 2. DeviceprofileIds array
	if len(s.DeviceprofileIds) > 0 {
		checks.append(t, "TestCheckResourceAttr", "deviceprofile_ids.#", fmt.Sprintf("%d", len(s.DeviceprofileIds)))
		for i, id := range s.DeviceprofileIds {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("deviceprofile_ids.%d", i), id)
		}
	}

	// 3. Exceptions (nested object)
	if s.Exceptions != nil {
		// exceptions.site_ids array
		if len(s.Exceptions.SiteIds) > 0 {
			checks.append(t, "TestCheckResourceAttr", "exceptions.site_ids.#", fmt.Sprintf("%d", len(s.Exceptions.SiteIds)))
			for i, id := range s.Exceptions.SiteIds {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("exceptions.site_ids.%d", i), id)
			}
		}

		// exceptions.sitegroup_ids array
		if len(s.Exceptions.SitegroupIds) > 0 {
			checks.append(t, "TestCheckResourceAttr", "exceptions.sitegroup_ids.#", fmt.Sprintf("%d", len(s.Exceptions.SitegroupIds)))
			for i, id := range s.Exceptions.SitegroupIds {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("exceptions.sitegroup_ids.%d", i), id)
			}
		}
	}

	// 4. FilterByDeviceprofile
	if s.FilterByDeviceprofile != nil {
		checks.append(t, "TestCheckResourceAttr", "filter_by_deviceprofile", fmt.Sprintf("%t", *s.FilterByDeviceprofile))
	}

	// 5. Id (computed-only)
	checks.append(t, "TestCheckResourceAttrSet", "id")

	// 6. Name (required)
	checks.append(t, "TestCheckResourceAttr", "name", s.Name)

	// 7. OrgId (required top-level field)
	checks.append(t, "TestCheckResourceAttr", "org_id", s.OrgId)

	return checks
}
