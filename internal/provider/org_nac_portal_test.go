package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_nac_portal"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgNacPortalModel(t *testing.T) {
	type testStep struct {
		config OrgNacPortalModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgNacPortalModel{
						OrgId: GetTestOrgId(),
						Name:  "test-nac-portal",
						Type:  stringPtr("guest_portal"),
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/org_nac_portal_resource/org_nac_portal_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b)
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		var FixtureOrgNacPortalModel OrgNacPortalModel
		err = hcl.Decode(&FixtureOrgNacPortalModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureOrgNacPortalModel,
				},
			},
		}
	}

	resourceType := "org_nac_portal"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_org_nac_portal.OrgNacPortalResourceSchema(t.Context()).Attributes)
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
	if tracker != nil {
		tracker.FieldCoverageReport(t)
	}
}

func (o *OrgNacPortalModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType)+"."+tName, tracker)

	// Computed fields
	checks.append(t, "TestCheckResourceAttrSet", "id")

	// Required fields
	checks.append(t, "TestCheckResourceAttr", "org_id", o.OrgId)
	checks.append(t, "TestCheckResourceAttr", "name", o.Name)

	// Optional configurable fields
	if o.AccessType != nil {
		checks.append(t, "TestCheckResourceAttr", "access_type", *o.AccessType)
	}

	if len(o.AdditionalCacerts) > 0 {
		checks.append(t, "TestCheckResourceAttr", "additional_cacerts.#", fmt.Sprintf("%d", len(o.AdditionalCacerts)))
		for i, cert := range o.AdditionalCacerts {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("additional_cacerts.%d", i), cert)
		}
	}

	if len(o.AdditionalNacServerName) > 0 {
		checks.append(t, "TestCheckResourceAttr", "additional_nac_server_name.#", fmt.Sprintf("%d", len(o.AdditionalNacServerName)))
		for i, name := range o.AdditionalNacServerName {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("additional_nac_server_name.%d", i), name)
		}
	}

	if o.CertExpireTime != nil {
		checks.append(t, "TestCheckResourceAttr", "cert_expire_time", fmt.Sprintf("%d", *o.CertExpireTime))
	}

	if o.EapType != nil {
		checks.append(t, "TestCheckResourceAttr", "eap_type", *o.EapType)
	}

	if o.EnableTelemetry != nil {
		checks.append(t, "TestCheckResourceAttr", "enable_telemetry", fmt.Sprintf("%t", *o.EnableTelemetry))
	}

	if o.ExpiryNotificationTime != nil {
		checks.append(t, "TestCheckResourceAttr", "expiry_notification_time", fmt.Sprintf("%d", *o.ExpiryNotificationTime))
	}

	if o.NotifyExpiry != nil {
		checks.append(t, "TestCheckResourceAttr", "notify_expiry", fmt.Sprintf("%t", *o.NotifyExpiry))
	}

	// Portal nested object
	if o.Portal != nil {
		if o.Portal.Auth != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.auth", *o.Portal.Auth)
		}

		if o.Portal.Expire != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.expire", fmt.Sprintf("%d", *o.Portal.Expire))
		}

		if o.Portal.ExternalPortalUrl != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.external_portal_url", *o.Portal.ExternalPortalUrl)
		}

		if o.Portal.ForceReconnect != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.force_reconnect", fmt.Sprintf("%t", *o.Portal.ForceReconnect))
		}

		if o.Portal.Forward != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.forward", fmt.Sprintf("%t", *o.Portal.Forward))
		}

		if o.Portal.ForwardUrl != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.forward_url", *o.Portal.ForwardUrl)
		}

		if o.Portal.MaxNumDevices != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.max_num_devices", fmt.Sprintf("%d", *o.Portal.MaxNumDevices))
		}

		if o.Portal.Privacy != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.privacy", fmt.Sprintf("%t", *o.Portal.Privacy))
		}
	}

	if o.Ssid != nil {
		checks.append(t, "TestCheckResourceAttr", "ssid", *o.Ssid)
	}

	// SSO nested object
	if o.Sso != nil {
		if o.Sso.IdpCert != nil {
			checks.append(t, "TestCheckResourceAttr", "sso.idp_cert", *o.Sso.IdpCert)
		}

		if o.Sso.IdpSignAlgo != nil {
			checks.append(t, "TestCheckResourceAttr", "sso.idp_sign_algo", *o.Sso.IdpSignAlgo)
		}

		if o.Sso.IdpSsoUrl != nil {
			checks.append(t, "TestCheckResourceAttr", "sso.idp_sso_url", *o.Sso.IdpSsoUrl)
		}

		if o.Sso.Issuer != nil {
			checks.append(t, "TestCheckResourceAttr", "sso.issuer", *o.Sso.Issuer)
		}

		if o.Sso.NameidFormat != nil {
			checks.append(t, "TestCheckResourceAttr", "sso.nameid_format", *o.Sso.NameidFormat)
		}

		if len(o.Sso.SsoRoleMatching) > 0 {
			checks.append(t, "TestCheckResourceAttr", "sso.sso_role_matching.#", fmt.Sprintf("%d", len(o.Sso.SsoRoleMatching)))
			for i, roleMatch := range o.Sso.SsoRoleMatching {
				if roleMatch.Assigned != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("sso.sso_role_matching.%d.assigned", i), *roleMatch.Assigned)
				}
				if roleMatch.Match != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("sso.sso_role_matching.%d.match", i), *roleMatch.Match)
				}
			}
		}

		if o.Sso.UseSsoRoleForCert != nil {
			checks.append(t, "TestCheckResourceAttr", "sso.use_sso_role_for_cert", fmt.Sprintf("%t", *o.Sso.UseSsoRoleForCert))
		}
	}

	if o.Tos != nil {
		checks.append(t, "TestCheckResourceAttr", "tos", *o.Tos)
	}

	if o.Type != nil {
		checks.append(t, "TestCheckResourceAttr", "type", *o.Type)
	}

	return checks
}
