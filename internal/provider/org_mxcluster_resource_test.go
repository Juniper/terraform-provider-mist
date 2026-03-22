package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_mxcluster"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgMxclusterModel(t *testing.T) {
	type testStep struct {
		config OrgMxclusterModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgMxclusterModel{
						Name:  "Test Org Mxcluster",
						OrgId: GetTestOrgId(),
					},
				},
			},
		},
	}

	fixtures, err := os.ReadFile("fixtures/org_mxcluster_resource/org_mxcluster_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	for i, fixture := range strings.Split(string(fixtures), "␞") {
		fixtureOrgMxclusterModel := OrgMxclusterModel{}
		err = hcl.Decode(&fixtureOrgMxclusterModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		// Need to set org_id to required field as it is used in the url for the resource
		fixtureOrgMxclusterModel.OrgId = GetTestOrgId()
		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: fixtureOrgMxclusterModel,
				},
			},
		}
	}

	resourceType := "org_mxcluster"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_org_mxcluster.OrgMxclusterResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				config := step.config
				siteConfig, siteRef := "", ""

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())

				combinedConfig := Render(resourceType, tName, string(f.Bytes()))
				if config.SiteId != nil && *config.SiteId != "" {
					siteConfig, siteRef = GetSiteBaseConfig(GetTestOrgId())
				}

				configStr := ""
				if siteConfig != "" {
					combinedConfig = strings.ReplaceAll(combinedConfig, "\"{site_id}\"", siteRef)
					configStr = siteConfig + "\n\n"
				}
				combinedConfig = configStr + combinedConfig

				checks := config.testChecks(t, resourceType, tName, tracker)
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
				IsUnitTest:               true,
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps:                    steps,
			})
		})
	}
	if tracker != nil {
		tracker.FieldCoverageReport(t)
	}
}

func (o *OrgMxclusterModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType)+"."+tName, tracker)

	// Check required fields
	checks.append(t, "TestCheckResourceAttrSet", "org_id")
	checks.append(t, "TestCheckResourceAttr", "name", o.Name)

	// Check optional basic fields
	if o.SiteId != nil {
		checks.append(t, "TestCheckResourceAttrSet", "site_id")
	}

	// Check mist_das
	if o.MistDas != nil {
		if o.MistDas.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "mist_das.enabled", fmt.Sprintf("%t", *o.MistDas.Enabled))
		}
		if len(o.MistDas.CoaServers) > 0 {
			checks.append(t, "TestCheckResourceAttr", "mist_das.coa_servers.#", fmt.Sprintf("%d", len(o.MistDas.CoaServers)))
			for i, coaServer := range o.MistDas.CoaServers {
				if coaServer.DisableEventTimestampCheck != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("mist_das.coa_servers.%d.disable_event_timestamp_check", i), fmt.Sprintf("%t", *coaServer.DisableEventTimestampCheck))
				}
				if coaServer.Enabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("mist_das.coa_servers.%d.enabled", i), fmt.Sprintf("%t", *coaServer.Enabled))
				}
				if coaServer.Host != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("mist_das.coa_servers.%d.host", i), *coaServer.Host)
				}
				if coaServer.Port != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("mist_das.coa_servers.%d.port", i), fmt.Sprintf("%d", *coaServer.Port))
				}
				if coaServer.RequireMessageAuthenticator != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("mist_das.coa_servers.%d.require_message_authenticator", i), fmt.Sprintf("%t", *coaServer.RequireMessageAuthenticator))
				}
				if coaServer.Secret != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("mist_das.coa_servers.%d.secret", i), *coaServer.Secret)
				}
			}
		}
	}

	// Check mist_nac
	if o.MistDas != nil {
		if o.MistNac.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "mist_nac.enabled", fmt.Sprintf("%t", *o.MistNac.Enabled))
		}
		if o.MistNac.AcctServerPort != nil {
			checks.append(t, "TestCheckResourceAttr", "mist_nac.acct_server_port", fmt.Sprintf("%d", *o.MistNac.AcctServerPort))
		}
		if o.MistNac.AuthServerPort != nil {
			checks.append(t, "TestCheckResourceAttr", "mist_nac.auth_server_port", fmt.Sprintf("%d", *o.MistNac.AuthServerPort))
		}
		if o.MistNac.Secret != nil {
			checks.append(t, "TestCheckResourceAttr", "mist_nac.secret", *o.MistNac.Secret)
		}
		// if len(o.MistNac.ClientIps) > 0 {
		// 	checks.append(t, "TestCheckResourceAttr", "mist_nac.client_ips.%", fmt.Sprintf("%d", len(o.MistNac.ClientIps)))
		// }
	}

	// Check radsec
	if o.Radsec != nil {
		if o.Radsec.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "radsec.enabled", fmt.Sprintf("%t", *o.Radsec.Enabled))
		}
		if o.Radsec.MatchSsid != nil {
			checks.append(t, "TestCheckResourceAttr", "radsec.match_ssid", fmt.Sprintf("%t", *o.Radsec.MatchSsid))
		}
		if o.Radsec.NasIpSource != nil {
			checks.append(t, "TestCheckResourceAttr", "radsec.nas_ip_source", *o.Radsec.NasIpSource)
		}
		if o.Radsec.ServerSelection != nil {
			checks.append(t, "TestCheckResourceAttr", "radsec.server_selection", *o.Radsec.ServerSelection)
		}
		if o.Radsec.SrcIpSource != nil {
			checks.append(t, "TestCheckResourceAttr", "radsec.src_ip_source", *o.Radsec.SrcIpSource)
		}
		if len(o.Radsec.ProxyHosts) > 0 {
			checks.append(t, "TestCheckResourceAttr", "radsec.proxy_hosts.#", fmt.Sprintf("%d", len(o.Radsec.ProxyHosts)))
			for i, host := range o.Radsec.ProxyHosts {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radsec.proxy_hosts.%d", i), host)
			}
		}
		if len(o.Radsec.AcctServers) > 0 {
			checks.append(t, "TestCheckResourceAttr", "radsec.acct_servers.#", fmt.Sprintf("%d", len(o.Radsec.AcctServers)))
			for i, server := range o.Radsec.AcctServers {
				if server.Host != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radsec.acct_servers.%d.host", i), *server.Host)
				}
				if server.Port != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radsec.acct_servers.%d.port", i), fmt.Sprintf("%d", *server.Port))
				}
				if server.Secret != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radsec.acct_servers.%d.secret", i), *server.Secret)
				}
				if len(server.Ssids) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radsec.acct_servers.%d.ssids.#", i), fmt.Sprintf("%d", len(server.Ssids)))
				}
			}
		}
		if len(o.Radsec.AuthServers) > 0 {
			checks.append(t, "TestCheckResourceAttr", "radsec.auth_servers.#", fmt.Sprintf("%d", len(o.Radsec.AuthServers)))
			for i, server := range o.Radsec.AuthServers {
				if server.Host != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radsec.auth_servers.%d.host", i), *server.Host)
				}
				if server.Port != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radsec.auth_servers.%d.port", i), fmt.Sprintf("%d", *server.Port))
				}
				if server.Secret != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radsec.auth_servers.%d.secret", i), *server.Secret)
				}
				if server.InbandStatusCheck != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radsec.auth_servers.%d.inband_status_check", i), fmt.Sprintf("%t", *server.InbandStatusCheck))
				}
				if server.InbandStatusInterval != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radsec.auth_servers.%d.inband_status_interval", i), fmt.Sprintf("%d", *server.InbandStatusInterval))
				}
				if server.KeywrapEnabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radsec.auth_servers.%d.keywrap_enabled", i), fmt.Sprintf("%t", *server.KeywrapEnabled))
				}
				if server.KeywrapFormat != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radsec.auth_servers.%d.keywrap_format", i), *server.KeywrapFormat)
				}
				if server.KeywrapKek != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radsec.auth_servers.%d.keywrap_kek", i), *server.KeywrapKek)
				}
				if server.KeywrapMack != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radsec.auth_servers.%d.keywrap_mack", i), *server.KeywrapMack)
				}
				if server.Retry != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radsec.auth_servers.%d.retry", i), fmt.Sprintf("%d", *server.Retry))
				}
				if server.Timeout != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radsec.auth_servers.%d.timeout", i), fmt.Sprintf("%d", *server.Timeout))
				}
				if len(server.Ssids) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radsec.auth_servers.%d.ssids.#", i), fmt.Sprintf("%d", len(server.Ssids)))
				}
			}
		}
	}

	// Check radsec_tls
	if o.RadsecTls != nil {
		if o.RadsecTls.Keypair != nil {
			checks.append(t, "TestCheckResourceAttr", "radsec_tls.keypair", *o.RadsecTls.Keypair)
		}
	}

	// Check tunterm_ap_subnets
	if len(o.TuntermApSubnets) > 0 {
		checks.append(t, "TestCheckResourceAttr", "tunterm_ap_subnets.#", fmt.Sprintf("%d", len(o.TuntermApSubnets)))
		for i, subnet := range o.TuntermApSubnets {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunterm_ap_subnets.%d", i), subnet)
		}
	}

	// Check tunterm_hosts
	if len(o.TuntermHosts) > 0 {
		checks.append(t, "TestCheckResourceAttr", "tunterm_hosts.#", fmt.Sprintf("%d", len(o.TuntermHosts)))
		for i, host := range o.TuntermHosts {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunterm_hosts.%d", i), host)
		}
	}

	// Check tunterm_hosts_order
	if len(o.TuntermHostsOrder) > 0 {
		checks.append(t, "TestCheckResourceAttr", "tunterm_hosts_order.#", fmt.Sprintf("%d", len(o.TuntermHostsOrder)))
		for i, order := range o.TuntermHostsOrder {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunterm_hosts_order.%d", i), fmt.Sprintf("%d", order))
		}
	}

	// Check tunterm_hosts_selection
	if o.TuntermHostsSelection != nil {
		checks.append(t, "TestCheckResourceAttr", "tunterm_hosts_selection", *o.TuntermHostsSelection)
	}

	// Check tunterm_monitoring_disabled
	if o.TuntermMonitoringDisabled != nil {
		checks.append(t, "TestCheckResourceAttr", "tunterm_monitoring_disabled", fmt.Sprintf("%t", *o.TuntermMonitoringDisabled))
	}

	// Check mxedge_mgmt
	if o.MxedgeMgmt != nil {
		if o.MxedgeMgmt.ConfigAutoRevert != nil {
			checks.append(t, "TestCheckResourceAttr", "mxedge_mgmt.config_auto_revert", fmt.Sprintf("%t", *o.MxedgeMgmt.ConfigAutoRevert))
		}
		if o.MxedgeMgmt.FipsEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "mxedge_mgmt.fips_enabled", fmt.Sprintf("%t", *o.MxedgeMgmt.FipsEnabled))
		}
		if o.MxedgeMgmt.MistPassword != nil {
			checks.append(t, "TestCheckResourceAttr", "mxedge_mgmt.mist_password", *o.MxedgeMgmt.MistPassword)
		}
		if o.MxedgeMgmt.RootPassword != nil {
			checks.append(t, "TestCheckResourceAttr", "mxedge_mgmt.root_password", *o.MxedgeMgmt.RootPassword)
		}
		if o.MxedgeMgmt.OobIpType != nil {
			checks.append(t, "TestCheckResourceAttr", "mxedge_mgmt.oob_ip_type", *o.MxedgeMgmt.OobIpType)
		}
		if o.MxedgeMgmt.OobIpType6 != nil {
			checks.append(t, "TestCheckResourceAttr", "mxedge_mgmt.oob_ip_type6", *o.MxedgeMgmt.OobIpType6)
		}
	}

	// Check proxy
	if o.Proxy != nil {
		if o.Proxy.Url != nil {
			checks.append(t, "TestCheckResourceAttr", "proxy.url", *o.Proxy.Url)
		}
		if o.Proxy.Disabled != nil {
			checks.append(t, "TestCheckResourceAttr", "proxy.disabled", fmt.Sprintf("%t", *o.Proxy.Disabled))
		}
	}

	// Check tunterm_extra_routes
	if len(o.TuntermExtraRoutes) > 0 {
		checks.append(t, "TestCheckResourceAttr", "tunterm_extra_routes.%", fmt.Sprintf("%d", len(o.TuntermExtraRoutes)))
		for cidr, route := range o.TuntermExtraRoutes {
			if route.Via != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunterm_extra_routes.%s.via", cidr), *route.Via)
			}
		}
	}

	// Check tunterm_dhcpd_config
	if len(o.TuntermDhcpdConfig) > 0 {
		checks.append(t, "TestCheckResourceAttr", "tunterm_dhcpd_config.%", fmt.Sprintf("%d", len(o.TuntermDhcpdConfig)))
		for vlanId, dhcpConfig := range o.TuntermDhcpdConfig {
			if dhcpConfig.Enabled != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunterm_dhcpd_config.%s.enabled", vlanId), fmt.Sprintf("%t", *dhcpConfig.Enabled))
			}
			if dhcpConfig.TuntermDhcpdConfigType != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunterm_dhcpd_config.%s.type", vlanId), *dhcpConfig.TuntermDhcpdConfigType)
			}
			if len(dhcpConfig.Servers) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunterm_dhcpd_config.%s.servers.#", vlanId), fmt.Sprintf("%d", len(dhcpConfig.Servers)))
				for i, server := range dhcpConfig.Servers {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunterm_dhcpd_config.%s.servers.%d", vlanId, i), server)
				}
			}
		}
	}

	// Check tunterm_monitoring
	if len(o.TuntermMonitoring) > 0 {
		checks.append(t, "TestCheckResourceAttr", "tunterm_monitoring.#", fmt.Sprintf("%d", len(o.TuntermMonitoring)))
		for i, monitoringGroup := range o.TuntermMonitoring {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunterm_monitoring.%d.#", i), fmt.Sprintf("%d", len(monitoringGroup)))
			for j, monitoringItem := range monitoringGroup {
				if monitoringItem.Host != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunterm_monitoring.%d.%d.host", i, j), *monitoringItem.Host)
				}
				if monitoringItem.Port != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunterm_monitoring.%d.%d.port", i, j), fmt.Sprintf("%d", *monitoringItem.Port))
				}
				if monitoringItem.Protocol != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunterm_monitoring.%d.%d.protocol", i, j), *monitoringItem.Protocol)
				}
				if monitoringItem.SrcVlanId != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunterm_monitoring.%d.%d.src_vlan_id", i, j), fmt.Sprintf("%d", *monitoringItem.SrcVlanId))
				}
				if monitoringItem.Timeout != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunterm_monitoring.%d.%d.timeout", i, j), fmt.Sprintf("%d", *monitoringItem.Timeout))
				}
			}
		}
	}

	return checks
}
