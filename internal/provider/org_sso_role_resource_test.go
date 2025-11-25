package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/resource_org_sso_role"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgSsoRole(t *testing.T) {
	type testStep struct {
		config OrgSsoRoleModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgSsoRoleModel{
						OrgId: GetTestOrgId(),
						Name:  "test-sso-role",
						Privileges: []OrgSsoRolePrivilegesValue{
							{
								Role:  "read",
								Scope: "org",
							},
						},
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/org_sso_role_resource/org_sso_role_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		var FixtureOrgSsoRoleModel OrgSsoRoleModel
		err = hcl.Decode(&FixtureOrgSsoRoleModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		FixtureOrgSsoRoleModel.OrgId = GetTestOrgId()

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureOrgSsoRoleModel,
				},
			},
		}
	}

	var checks testChecks
	resourceType := "org_sso_role"
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				config := step.config
				siteConfig, sitegroupConfig, siteRef, sitegroupRef := "", "", "", ""

				for i, p := range config.Privileges {
					switch p.Scope {
					case "site":
						if siteConfig == "" { // only get once even if multiple privileges use it
							siteConfig, siteRef = GetSiteBaseConfig(GetTestOrgId())
							config.Privileges[i].SiteId = stringPtr("{site_id}")
						}
					case "sitegroup":
						if sitegroupConfig == "" {
							sitegroupConfig, sitegroupRef = GetSitegroupBaseConfig(GetTestOrgId())
							config.Privileges[i].SitegroupId = stringPtr("{sitegroup_id}")
						}
					}
				}

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				combinedConfig := Render(resourceType, tName, string(f.Bytes()))
				configStr := ""
				if siteConfig != "" {
					combinedConfig = strings.ReplaceAll(combinedConfig, "\"{site_id}\"", siteRef)
					configStr = siteConfig + "\n\n" + configStr
				}
				if sitegroupConfig != "" {
					combinedConfig = strings.ReplaceAll(combinedConfig, "\"{sitegroup_id}\"", sitegroupRef)
					configStr = sitegroupConfig + "\n\n" + configStr
				}

				combinedConfig = configStr + combinedConfig

				checks = config.testChecks(t, PrefixProviderName(resourceType), tName)
				chkLog := checks.string()
				stepName := fmt.Sprintf("test case %s step %d", tName, i+1)

				t.Logf("\n// ------ begin config for %s ------\n%s// -------- end config for %s ------\n\n", stepName, combinedConfig, stepName)
				t.Logf("\n// ------ begin checks for %s ------\n%s// -------- end config for %s ------\n\n", stepName, chkLog, stepName)

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
	FieldCoverageReport(t, &checks, "after_test_run")
}

func (o *OrgSsoRoleModel) testChecks(t testing.TB, rType, rName string) testChecks {
	checks := newTestChecks(rType + "." + rName)

	// Track field coverage
	TrackFieldCoverage(t, &checks, "org_sso_role", resource_org_sso_role.OrgSsoRoleResourceSchema)

	// Check required fields
	checks.append(t, "TestCheckResourceAttr", "org_id", o.OrgId)
	checks.append(t, "TestCheckResourceAttr", "name", o.Name)

	// Check computed-only fields (verify they exist)
	checks.append(t, "TestCheckResourceAttrSet", "id")

	// Check privileges array
	if len(o.Privileges) > 0 {
		checks.append(t, "TestCheckResourceAttr", "privileges.#", fmt.Sprintf("%d", len(o.Privileges)))

		// Check all privilege entries
		for i, privilege := range o.Privileges {
			// Check required fields in each privilege
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("privileges.%d.role", i), privilege.Role)
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("privileges.%d.scope", i), privilege.Scope)

			// Check optional fields in each privilege
			if privilege.SiteId != nil {
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("privileges.%d.site_id", i))
			}
			if privilege.SitegroupId != nil {
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("privileges.%d.sitegroup_id", i))
			}

			if len(privilege.Views) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("privileges.%d.views.#", i), fmt.Sprintf("%d", len(privilege.Views)))
				for j, view := range privilege.Views {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("privileges.%d.views.%d", i, j), view)
				}
			}
		}
	}

	return checks
}
