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

func TestOrgServicepolicy(t *testing.T) {
	type testStep struct {
		config OrgServicepolicyModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgServicepolicyModel{
						OrgId: GetTestOrgId(),
						Name:  "test-servicepolicy",
					},
				},
			},
		},
	}

	// Load fixture data following the org_wlan pattern
	b, err := os.ReadFile("fixtures/org_servicepolicy_resource/org_servicepolicy_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		fixtureOrgServicepolicyModel := OrgServicepolicyModel{}
		err = hcl.Decode(&fixtureOrgServicepolicyModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		fixtureOrgServicepolicyModel.OrgId = GetTestOrgId()

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: fixtureOrgServicepolicyModel,
				},
			},
		}
	}

	resourceType := "org_servicepolicy"
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				// Generate Terraform configuration using automated HCL generation
				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&step.config, f.Body())
				configStr := Render(resourceType, tName, string(f.Bytes()))

				checks := step.config.testChecks(t, resourceType, tName)
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

func (o *OrgServicepolicyModel) testChecks(t testing.TB, rType, tName string) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + tName)

	// Check required fields
	checks.append(t, "TestCheckResourceAttr", "org_id", o.OrgId)
	checks.append(t, "TestCheckResourceAttr", "name", o.Name)

	// Check computed field
	checks.append(t, "TestCheckResourceAttrSet", "id")

	// Check optional fields if present
	if o.Action != nil {
		checks.append(t, "TestCheckResourceAttr", "action", *o.Action)
	}
	if o.LocalRouting != nil {
		checks.append(t, "TestCheckResourceAttr", "local_routing", fmt.Sprintf("%t", *o.LocalRouting))
	}
	if o.PathPreference != nil {
		checks.append(t, "TestCheckResourceAttr", "path_preference", *o.PathPreference)
	}

	// Check list fields
	if len(o.Services) > 0 {
		checks.append(t, "TestCheckResourceAttr", "services.#", fmt.Sprintf("%d", len(o.Services)))
		for i, service := range o.Services {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("services.%d", i), service)
		}
	}
	if len(o.Tenants) > 0 {
		checks.append(t, "TestCheckResourceAttr", "tenants.#", fmt.Sprintf("%d", len(o.Tenants)))
		for i, tenant := range o.Tenants {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tenants.%d", i), tenant)
		}
	}

	// Check nested object fields
	if o.Aamw != nil {
		if o.Aamw.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "aamw.enabled", fmt.Sprintf("%t", *o.Aamw.Enabled))
		}
		if o.Aamw.Profile != nil {
			checks.append(t, "TestCheckResourceAttr", "aamw.profile", *o.Aamw.Profile)
		}
		if o.Aamw.AamwprofileId != nil {
			checks.append(t, "TestCheckResourceAttr", "aamw.aamwprofile_id", *o.Aamw.AamwprofileId)
		}
	}

	if o.Antivirus != nil {
		if o.Antivirus.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "antivirus.enabled", fmt.Sprintf("%t", *o.Antivirus.Enabled))
		}
		if o.Antivirus.Profile != nil {
			checks.append(t, "TestCheckResourceAttr", "antivirus.profile", *o.Antivirus.Profile)
		}
		if o.Antivirus.AvprofileId != nil {
			checks.append(t, "TestCheckResourceAttr", "antivirus.avprofile_id", *o.Antivirus.AvprofileId)
		}
	}

	if o.Appqoe != nil {
		if o.Appqoe.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "appqoe.enabled", fmt.Sprintf("%t", *o.Appqoe.Enabled))
		}
	}

	if len(o.Ewf) > 0 {
		checks.append(t, "TestCheckResourceAttr", "ewf.#", fmt.Sprintf("%d", len(o.Ewf)))
		for i, ewf := range o.Ewf {
			if ewf.Enabled != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ewf.%d.enabled", i), fmt.Sprintf("%t", *ewf.Enabled))
			}
			if ewf.AlertOnly != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ewf.%d.alert_only", i), fmt.Sprintf("%t", *ewf.AlertOnly))
			}
			if ewf.Profile != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ewf.%d.profile", i), *ewf.Profile)
			}
			if ewf.BlockMessage != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ewf.%d.block_message", i), *ewf.BlockMessage)
			}
		}
	}

	if o.Idp != nil {
		if o.Idp.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "idp.enabled", fmt.Sprintf("%t", *o.Idp.Enabled))
		}
		if o.Idp.AlertOnly != nil {
			checks.append(t, "TestCheckResourceAttr", "idp.alert_only", fmt.Sprintf("%t", *o.Idp.AlertOnly))
		}
		if o.Idp.Profile != nil {
			checks.append(t, "TestCheckResourceAttr", "idp.profile", *o.Idp.Profile)
		}
		if o.Idp.IdpprofileId != nil {
			checks.append(t, "TestCheckResourceAttr", "idp.idpprofile_id", *o.Idp.IdpprofileId)
		}
	}

	if o.SslProxy != nil {
		if o.SslProxy.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "ssl_proxy.enabled", fmt.Sprintf("%t", *o.SslProxy.Enabled))
		}
		if o.SslProxy.CiphersCategory != nil {
			checks.append(t, "TestCheckResourceAttr", "ssl_proxy.ciphers_category", *o.SslProxy.CiphersCategory)
		}
	}

	return checks
}
