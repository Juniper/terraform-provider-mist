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

func TestOrgNacidpModel(t *testing.T) {
	type testStep struct {
		config OrgNacidpModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgNacidpModel{
						OrgId:            GetTestOrgId(),
						Name:             "test-nacidp",
						IdpType:          "ldap",
						LdapBaseDn:       stringPtr("dc=example,dc=com"),
						LdapBindDn:       stringPtr("cn=admin,dc=example,dc=com"),
						LdapBindPassword: stringPtr("password123"),
						LdapType:         stringPtr("azure"),
						LdapServerHosts:  []string{"ldap.example.com"},
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/org_nacidp_resource/org_nacidp_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	fixtures := strings.Split(string(b), "␞")

	for i, fixture := range fixtures {
		var FixtureOrgNacidpModel OrgNacidpModel
		err = hcl.Decode(&FixtureOrgNacidpModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		FixtureOrgNacidpModel.OrgId = GetTestOrgId()

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureOrgNacidpModel,
				},
			},
		}
	}

	resourceType := "org_nacidp"
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {

			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				config := step.config

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				combinedConfig := Render(resourceType, tName, string(f.Bytes()))

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

func (o *OrgNacidpModel) testChecks(t testing.TB, rType, rName string) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + rName)

	// Computed fields
	checks.append(t, "TestCheckResourceAttrSet", "id")

	// Required fields
	checks.append(t, "TestCheckResourceAttr", "org_id", o.OrgId)
	checks.append(t, "TestCheckResourceAttr", "name", o.Name)
	checks.append(t, "TestCheckResourceAttr", "idp_type", o.IdpType)

	// Optional string fields
	if o.GroupFilter != nil {
		checks.append(t, "TestCheckResourceAttr", "group_filter", *o.GroupFilter)
	}
	if o.LdapBaseDn != nil {
		checks.append(t, "TestCheckResourceAttr", "ldap_base_dn", *o.LdapBaseDn)
	}
	if o.LdapBindDn != nil {
		checks.append(t, "TestCheckResourceAttr", "ldap_bind_dn", *o.LdapBindDn)
	}
	if o.LdapBindPassword != nil {
		checks.append(t, "TestCheckResourceAttr", "ldap_bind_password", *o.LdapBindPassword)
	}
	if o.LdapClientCert != nil {
		checks.append(t, "TestCheckResourceAttr", "ldap_client_cert", *o.LdapClientCert)
	}
	if o.LdapClientKey != nil {
		checks.append(t, "TestCheckResourceAttr", "ldap_client_key", *o.LdapClientKey)
	}
	if o.LdapGroupAttr != nil {
		checks.append(t, "TestCheckResourceAttr", "ldap_group_attr", *o.LdapGroupAttr)
	}
	if o.LdapGroupDn != nil {
		checks.append(t, "TestCheckResourceAttr", "ldap_group_dn", *o.LdapGroupDn)
	}
	if o.LdapType != nil {
		checks.append(t, "TestCheckResourceAttr", "ldap_type", *o.LdapType)
	}
	if o.LdapUserFilter != nil {
		checks.append(t, "TestCheckResourceAttr", "ldap_user_filter", *o.LdapUserFilter)
	}
	if o.MemberFilter != nil {
		checks.append(t, "TestCheckResourceAttr", "member_filter", *o.MemberFilter)
	}
	if o.OauthCcClientId != nil {
		checks.append(t, "TestCheckResourceAttr", "oauth_cc_client_id", *o.OauthCcClientId)
	}
	if o.OauthCcClientSecret != nil {
		checks.append(t, "TestCheckResourceAttr", "oauth_cc_client_secret", *o.OauthCcClientSecret)
	}
	if o.OauthDiscoveryUrl != nil {
		checks.append(t, "TestCheckResourceAttr", "oauth_discovery_url", *o.OauthDiscoveryUrl)
	}
	if o.OauthPingIdentityRegion != nil {
		checks.append(t, "TestCheckResourceAttr", "oauth_ping_identity_region", *o.OauthPingIdentityRegion)
	}
	if o.OauthRopcClientId != nil {
		checks.append(t, "TestCheckResourceAttr", "oauth_ropc_client_id", *o.OauthRopcClientId)
	}
	if o.OauthRopcClientSecret != nil {
		checks.append(t, "TestCheckResourceAttr", "oauth_ropc_client_secret", *o.OauthRopcClientSecret)
	}
	if o.OauthTenantId != nil {
		checks.append(t, "TestCheckResourceAttr", "oauth_tenant_id", *o.OauthTenantId)
	}
	if o.OauthType != nil {
		checks.append(t, "TestCheckResourceAttr", "oauth_type", *o.OauthType)
	}
	if o.ScimSecretToken != nil {
		checks.append(t, "TestCheckResourceAttr", "scim_secret_token", *o.ScimSecretToken)
	}

	// Boolean fields
	if o.LdapResolveGroups != nil {
		checks.append(t, "TestCheckResourceAttr", "ldap_resolve_groups", fmt.Sprintf("%t", *o.LdapResolveGroups))
	}
	if o.ScimEnabled != nil {
		checks.append(t, "TestCheckResourceAttr", "scim_enabled", fmt.Sprintf("%t", *o.ScimEnabled))
	}

	// List fields
	if len(o.LdapCacerts) > 0 {
		checks.append(t, "TestCheckResourceAttr", "ldap_cacerts.#", fmt.Sprintf("%d", len(o.LdapCacerts)))
		for i, cert := range o.LdapCacerts {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ldap_cacerts.%d", i), cert)
		}
	}
	if len(o.LdapServerHosts) > 0 {
		checks.append(t, "TestCheckResourceAttr", "ldap_server_hosts.#", fmt.Sprintf("%d", len(o.LdapServerHosts)))
		for i, host := range o.LdapServerHosts {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ldap_server_hosts.%d", i), host)
		}
	}

	return checks
}
