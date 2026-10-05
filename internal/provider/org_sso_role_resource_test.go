package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_sso_role"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgSsoRoleModel(t *testing.T) {
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

	resourceType := "org_sso_role"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_org_sso_role.OrgSsoRoleResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				config := step.config
				siteConfig, sitegroupConfig, siteRef, sitegroupRef := "", "", "", ""
				var siteIdxs, sitegroupIdxs []int

				for i, p := range config.Privileges {
					switch p.Scope {
					case "site":
						if siteConfig == "" { // only create the site resource once, but point every site-scoped privilege at it
							siteConfig, siteRef = GetSiteBaseConfig(GetTestOrgId())
						}
						config.Privileges[i].SiteId = stringPtr("{site_id}")
						siteIdxs = append(siteIdxs, i)
					case "sitegroup":
						if sitegroupConfig == "" {
							sitegroupConfig, sitegroupRef = GetSitegroupBaseConfig(GetTestOrgId())
						}
						config.Privileges[i].SitegroupId = stringPtr("{sitegroup_id}")
						sitegroupIdxs = append(sitegroupIdxs, i)
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

				checks := config.testChecks(t, resourceType, tName, tracker, siteIdxs, sitegroupIdxs, siteRef, sitegroupRef)
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
	if tracker != nil {
		tracker.FieldCoverageReport(t)
	}
}

func (o *OrgSsoRoleModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker, siteIdxs, sitegroupIdxs []int, siteRef, sitegroupRef string) testChecks {
	checks := newTestChecks(PrefixProviderName(rType)+"."+tName, tracker)

	var skip []string
	for _, idx := range siteIdxs {
		skip = append(skip, fmt.Sprintf("privileges.%d.site_id", idx))
	}
	for _, idx := range sitegroupIdxs {
		skip = append(skip, fmt.Sprintf("privileges.%d.sitegroup_id", idx))
	}
	appendReflectChecks(t, &checks, o, skip...)

	// The site/sitegroup id are only known after apply, so compare them against
	// the referenced resource's id instead of asserting a literal value.
	for _, idx := range siteIdxs {
		checks.append(t, "TestCheckResourceAttrPair", fmt.Sprintf("privileges.%d.site_id", idx), strings.TrimSuffix(siteRef, ".id"), "id")
	}
	for _, idx := range sitegroupIdxs {
		checks.append(t, "TestCheckResourceAttrPair", fmt.Sprintf("privileges.%d.sitegroup_id", idx), strings.TrimSuffix(sitegroupRef, ".id"), "id")
	}

	return checks
}
