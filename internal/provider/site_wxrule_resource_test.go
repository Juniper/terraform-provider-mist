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

func TestSiteWxruleModel(t *testing.T) {
	type testStep struct {
		config SiteWxruleModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: SiteWxruleModel{
						Action: "allow",
						Order:  1,
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/site_wxrule_resource/site_wxrule_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {

		var FixtureSiteWxruleModel SiteWxruleModel
		err = hcl.Decode(&FixtureSiteWxruleModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureSiteWxruleModel,
				},
			},
		}
	}

	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			resourceType := "site_wxrule"

			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				siteConfig, siteRef := GetSiteBaseConfig(GetTestOrgId())
				config := step.config

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				f.Body().SetAttributeRaw("site_id", hclwrite.TokensForIdentifier(siteRef))
				combinedConfig := siteConfig + "\n\n" + Render(resourceType, tName, string(f.Bytes()))

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

func (s *SiteWxruleModel) testChecks(t testing.TB, rType, rName string) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + rName)
	checks.append(t, "TestCheckResourceAttrSet", "site_id")

	// Required string attributes
	checks.append(t, "TestCheckResourceAttr", "action", s.Action)
	checks.append(t, "TestCheckResourceAttr", "order", fmt.Sprintf("%d", s.Order))

	// Optional boolean attributes
	if s.Enabled != nil {
		checks.append(t, "TestCheckResourceAttr", "enabled", fmt.Sprintf("%t", *s.Enabled))
	}

	// Optional list attributes
	if len(s.ApplyTags) > 0 {
		checks.append(t, "TestCheckResourceAttr", "apply_tags.#", fmt.Sprintf("%d", len(s.ApplyTags)))
		for i, tag := range s.ApplyTags {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("apply_tags.%d", i), tag)
		}
	}

	if len(s.BlockedApps) > 0 {
		checks.append(t, "TestCheckResourceAttr", "blocked_apps.#", fmt.Sprintf("%d", len(s.BlockedApps)))
		for i, app := range s.BlockedApps {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("blocked_apps.%d", i), app)
		}
	}

	if len(s.DstAllowWxtags) > 0 {
		checks.append(t, "TestCheckResourceAttr", "dst_allow_wxtags.#", fmt.Sprintf("%d", len(s.DstAllowWxtags)))
		for i, tag := range s.DstAllowWxtags {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dst_allow_wxtags.%d", i), tag)
		}
	}

	if len(s.DstDenyWxtags) > 0 {
		checks.append(t, "TestCheckResourceAttr", "dst_deny_wxtags.#", fmt.Sprintf("%d", len(s.DstDenyWxtags)))
		for i, tag := range s.DstDenyWxtags {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dst_deny_wxtags.%d", i), tag)
		}
	}

	if len(s.DstWxtags) > 0 {
		checks.append(t, "TestCheckResourceAttr", "dst_wxtags.#", fmt.Sprintf("%d", len(s.DstWxtags)))
		for i, tag := range s.DstWxtags {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dst_wxtags.%d", i), tag)
		}
	}

	if len(s.SrcWxtags) > 0 {
		checks.append(t, "TestCheckResourceAttr", "src_wxtags.#", fmt.Sprintf("%d", len(s.SrcWxtags)))
		for i, tag := range s.SrcWxtags {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("src_wxtags.%d", i), tag)
		}
	}

	return checks
}
