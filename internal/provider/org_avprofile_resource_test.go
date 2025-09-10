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

func TestOrgAvprofileModel(t *testing.T) {
	type testStep struct {
		config OrgAvprofileModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgAvprofileModel{
						OrgId:     GetTestOrgId(),
						Name:      "test_avprofile",
						Protocols: []string{"http", "ftp"},
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/org_avprofile_resource/org_avprofile_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		var FixtureOrgAvprofileModel OrgAvprofileModel

		err = hcl.Decode(&FixtureOrgAvprofileModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureOrgAvprofileModel,
				},
			},
		}
	}

	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			resourceType := "org_avprofile"

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

func (s *OrgAvprofileModel) testChecks(t testing.TB, rType, rName string) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + rName)
	// Required parameters
	checks.append(t, "TestCheckResourceAttrSet", "org_id")
	checks.append(t, "TestCheckResourceAttr", "name", s.Name)
	checks.append(t, "TestCheckResourceAttr", "protocols.#", fmt.Sprintf("%d", len(s.Protocols)))
	for i, prot := range s.Protocols {
		checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("protocols.%d", i), prot)
	}

	// Optional parameters
	if s.FallbackAction != nil {
		checks.append(t, "TestCheckResourceAttr", "fallback_action", *s.FallbackAction)
	}
	if s.MaxFilesize != nil {
		checks.append(t, "TestCheckResourceAttr", "max_filesize", fmt.Sprintf("%d", *s.MaxFilesize))
	}
	if len(s.MimeWhitelist) > 0 {
		checks.append(t, "TestCheckResourceAttr", "mime_whitelist.#", fmt.Sprintf("%d", len(s.MimeWhitelist)))
		for i, mime := range s.MimeWhitelist {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("mime_whitelist.%d", i), mime)
		}
	}
	if len(s.UrlWhitelist) > 0 {
		checks.append(t, "TestCheckResourceAttr", "url_whitelist.#", fmt.Sprintf("%d", len(s.UrlWhitelist)))
		for i, url := range s.UrlWhitelist {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("url_whitelist.%d", i), url)
		}
	}

	return checks
}
