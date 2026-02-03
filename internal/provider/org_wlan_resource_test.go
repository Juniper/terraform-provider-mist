// WIP
package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_wlan"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgWlanModel(t *testing.T) {
	type testStep struct {
		config OrgWlanModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgWlanModel{
						OrgId: GetTestOrgId(),
						Ssid:  "TestSSID",
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/org_wlan_resource/org_wlan_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		fixtureOrgWlanModel := OrgWlanModel{}
		err = hcl.Decode(&fixtureOrgWlanModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		fixtureOrgWlanModel.OrgId = GetTestOrgId()

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: fixtureOrgWlanModel,
				},
			},
		}
	}

	resourceType := "org_wlan"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_org_wlan.OrgWlanResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			templateName := "test_template"

			// Create single-step tests with combined config (template + WLAN)
			// Since WLANs require a template, we create both in the same config
			// but focus our checks on the WLAN resource being tested
			steps := make([]resource.TestStep, len(tCase.steps))

			for i, step := range tCase.steps {
				// Generate combined config: template + WLAN
				combinedConfig := generateOrgWlanTestConfig(templateName, tName, step.config)

				// Focus checks on the WLAN resource (template is just a prerequisite)
				checks := step.config.testChecks(t, resourceType, tName, tracker)
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
	if tracker != nil {
		tracker.FieldCoverageReport(t)
	}
}

// generateOrgWlanTestConfig creates a combined configuration with both a WLAN template and a WLAN
// This handles the prerequisite that WLANs require a template to exist
func generateOrgWlanTestConfig(templateName, wlanName string, wlanConfig OrgWlanModel) string {
	// Create the prerequisite WLAN template
	templateConfig := OrgWlantemplateModel{
		Name:  "Test_WLAN_Template",
		OrgId: wlanConfig.OrgId,
	}

	f := hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(&templateConfig, f.Body())
	templateConfigStr := Render("org_wlantemplate", templateName, string(f.Bytes()))

	// Create the WLAN that references the template
	templateRef := fmt.Sprintf("mist_org_wlantemplate.%s.id", templateName)

	f = hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(&wlanConfig, f.Body())

	// Add the template_id attribute to the body before rendering
	f.Body().SetAttributeRaw("template_id", hclwrite.TokensForIdentifier(templateRef))
	wlanConfigStr := Render("org_wlan", wlanName, string(f.Bytes()))

	// Combine both configs
	return templateConfigStr + "\n\n" + wlanConfigStr
}

func (s *OrgWlanModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType)+"."+tName, tracker)

	// Always check required fields
	checks.append(t, "TestCheckResourceAttrSet", "id")
	checks.append(t, "TestCheckResourceAttrSet", "org_id")
	checks.append(t, "TestCheckResourceAttr", "org_id", s.OrgId)
	checks.append(t, "TestCheckResourceAttr", "ssid", s.Ssid)
	checks.append(t, "TestCheckResourceAttrSet", "template_id")
	checks.append(t, "TestCheckResourceAttr", "msp_id", "")

	// Core boolean settings
	if s.Enabled != nil {
		checks.append(t, "TestCheckResourceAttr", "enabled", fmt.Sprintf("%t", *s.Enabled))
	}
	if s.HideSsid != nil {
		checks.append(t, "TestCheckResourceAttr", "hide_ssid", fmt.Sprintf("%t", *s.HideSsid))
	}
	if s.Isolation != nil {
		checks.append(t, "TestCheckResourceAttr", "isolation", fmt.Sprintf("%t", *s.Isolation))
	}

	// Accounting settings
	if s.AcctImmediateUpdate != nil {
		checks.append(t, "TestCheckResourceAttr", "acct_immediate_update", fmt.Sprintf("%t", *s.AcctImmediateUpdate))
	}
	if s.AcctInterimInterval != nil {
		checks.append(t, "TestCheckResourceAttr", "acct_interim_interval", fmt.Sprintf("%d", *s.AcctInterimInterval))
	}

	// Network protocol settings
	if s.AllowIpv6Ndp != nil {
		checks.append(t, "TestCheckResourceAttr", "allow_ipv6_ndp", fmt.Sprintf("%t", *s.AllowIpv6Ndp))
	}
	if s.AllowMdns != nil {
		checks.append(t, "TestCheckResourceAttr", "allow_mdns", fmt.Sprintf("%t", *s.AllowMdns))
	}
	if s.AllowSsdp != nil {
		checks.append(t, "TestCheckResourceAttr", "allow_ssdp", fmt.Sprintf("%t", *s.AllowSsdp))
	}
	if s.ArpFilter != nil {
		checks.append(t, "TestCheckResourceAttr", "arp_filter", fmt.Sprintf("%t", *s.ArpFilter))
	}

	// Band steering
	if s.BandSteer != nil {
		checks.append(t, "TestCheckResourceAttr", "band_steer", fmt.Sprintf("%t", *s.BandSteer))
	}
	if s.BandSteerForceBand5 != nil {
		checks.append(t, "TestCheckResourceAttr", "band_steer_force_band5", fmt.Sprintf("%t", *s.BandSteerForceBand5))
	}

	// Client management
	if s.BlockBlacklistClients != nil {
		checks.append(t, "TestCheckResourceAttr", "block_blacklist_clients", fmt.Sprintf("%t", *s.BlockBlacklistClients))
	}
	if s.MaxNumClients != nil {
		checks.append(t, "TestCheckResourceAttr", "max_num_clients", fmt.Sprintf("%d", *s.MaxNumClients))
	}
	if s.MaxIdletime != nil {
		checks.append(t, "TestCheckResourceAttr", "max_idletime", fmt.Sprintf("%d", *s.MaxIdletime))
	}

	// Wireless protocol settings
	if s.Disable11ax != nil {
		checks.append(t, "TestCheckResourceAttr", "disable_11ax", fmt.Sprintf("%t", *s.Disable11ax))
	}
	if s.Disable11be != nil {
		checks.append(t, "TestCheckResourceAttr", "disable_11be", fmt.Sprintf("%t", *s.Disable11be))
	}
	if s.DisableHtVhtRates != nil {
		checks.append(t, "TestCheckResourceAttr", "disable_ht_vht_rates", fmt.Sprintf("%t", *s.DisableHtVhtRates))
	}
	if s.DisableUapsd != nil {
		checks.append(t, "TestCheckResourceAttr", "disable_uapsd", fmt.Sprintf("%t", *s.DisableUapsd))
	}
	if s.DisableV1RoamNotify != nil {
		checks.append(t, "TestCheckResourceAttr", "disable_v1_roam_notify", fmt.Sprintf("%t", *s.DisableV1RoamNotify))
	}
	if s.DisableV2RoamNotify != nil {
		checks.append(t, "TestCheckResourceAttr", "disable_v2_roam_notify", fmt.Sprintf("%t", *s.DisableV2RoamNotify))
	}
	if s.DisableWmm != nil {
		checks.append(t, "TestCheckResourceAttr", "disable_wmm", fmt.Sprintf("%t", *s.DisableWmm))
	}
	if s.Dtim != nil {
		checks.append(t, "TestCheckResourceAttr", "dtim", fmt.Sprintf("%d", *s.Dtim))
	}

	// Wireless features
	if s.EnableLocalKeycaching != nil {
		checks.append(t, "TestCheckResourceAttr", "enable_local_keycaching", fmt.Sprintf("%t", *s.EnableLocalKeycaching))
	}
	if s.EnableWirelessBridging != nil {
		checks.append(t, "TestCheckResourceAttr", "enable_wireless_bridging", fmt.Sprintf("%t", *s.EnableWirelessBridging))
	}
	if s.EnableWirelessBridgingDhcpTracking != nil {
		checks.append(t, "TestCheckResourceAttr", "enable_wireless_bridging_dhcp_tracking", fmt.Sprintf("%t", *s.EnableWirelessBridgingDhcpTracking))
	}
	if s.FastDot1xTimers != nil {
		checks.append(t, "TestCheckResourceAttr", "fast_dot1x_timers", fmt.Sprintf("%t", *s.FastDot1xTimers))
	}
	if s.HostnameIe != nil {
		checks.append(t, "TestCheckResourceAttr", "hostname_ie", fmt.Sprintf("%t", *s.HostnameIe))
	}

	// Network isolation
	if s.L2Isolation != nil {
		checks.append(t, "TestCheckResourceAttr", "l2_isolation", fmt.Sprintf("%t", *s.L2Isolation))
	}
	if s.LegacyOverds != nil {
		checks.append(t, "TestCheckResourceAttr", "legacy_overds", fmt.Sprintf("%t", *s.LegacyOverds))
	}
	if s.LimitBcast != nil {
		checks.append(t, "TestCheckResourceAttr", "limit_bcast", fmt.Sprintf("%t", *s.LimitBcast))
	}
	if s.LimitProbeResponse != nil {
		checks.append(t, "TestCheckResourceAttr", "limit_probe_response", fmt.Sprintf("%t", *s.LimitProbeResponse))
	}

	// Static settings
	if s.NoStaticDns != nil {
		checks.append(t, "TestCheckResourceAttr", "no_static_dns", fmt.Sprintf("%t", *s.NoStaticDns))
	}
	if s.NoStaticIp != nil {
		checks.append(t, "TestCheckResourceAttr", "no_static_ip", fmt.Sprintf("%t", *s.NoStaticIp))
	}

	// Roaming and clustering
	if s.ReconnectClientsWhenRoamingMxcluster != nil {
		checks.append(t, "TestCheckResourceAttr", "reconnect_clients_when_roaming_mxcluster", fmt.Sprintf("%t", *s.ReconnectClientsWhenRoamingMxcluster))
	}
	if s.RoamMode != nil {
		checks.append(t, "TestCheckResourceAttr", "roam_mode", *s.RoamMode)
	}

	// SLE and monitoring
	if s.SleExcluded != nil {
		checks.append(t, "TestCheckResourceAttr", "sle_excluded", fmt.Sprintf("%t", *s.SleExcluded))
	}
	if s.UseEapolV1 != nil {
		checks.append(t, "TestCheckResourceAttr", "use_eapol_v1", fmt.Sprintf("%t", *s.UseEapolV1))
	}

	// Interface
	if s.Interface != nil {
		checks.append(t, "TestCheckResourceAttr", "interface", *s.Interface)
	}

	// VLAN settings
	if s.VlanEnabled != nil {
		checks.append(t, "TestCheckResourceAttr", "vlan_enabled", fmt.Sprintf("%t", *s.VlanEnabled))
	}
	if s.VlanId != nil {
		checks.append(t, "TestCheckResourceAttr", "vlan_id", *s.VlanId)
	}
	if s.VlanPooling != nil {
		checks.append(t, "TestCheckResourceAttr", "vlan_pooling", fmt.Sprintf("%t", *s.VlanPooling))
	}
	if len(s.VlanIds) > 0 {
		checks.append(t, "TestCheckResourceAttr", "vlan_ids.#", fmt.Sprintf("%d", len(s.VlanIds)))
		for i, vlanId := range s.VlanIds {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vlan_ids.%d", i), vlanId)
		}
	}

	// Bandwidth limits - WLAN level
	if s.WlanLimitDownEnabled != nil {
		checks.append(t, "TestCheckResourceAttr", "wlan_limit_down_enabled", fmt.Sprintf("%t", *s.WlanLimitDownEnabled))
	}
	if s.WlanLimitDown != nil {
		checks.append(t, "TestCheckResourceAttr", "wlan_limit_down", *s.WlanLimitDown)
	}
	if s.WlanLimitUpEnabled != nil {
		checks.append(t, "TestCheckResourceAttr", "wlan_limit_up_enabled", fmt.Sprintf("%t", *s.WlanLimitUpEnabled))
	}
	if s.WlanLimitUp != nil {
		checks.append(t, "TestCheckResourceAttr", "wlan_limit_up", *s.WlanLimitUp)
	}

	// Bandwidth limits - Client level
	if s.ClientLimitDownEnabled != nil {
		checks.append(t, "TestCheckResourceAttr", "client_limit_down_enabled", fmt.Sprintf("%t", *s.ClientLimitDownEnabled))
	}
	if s.ClientLimitDown != nil {
		checks.append(t, "TestCheckResourceAttr", "client_limit_down", *s.ClientLimitDown)
	}
	if s.ClientLimitUpEnabled != nil {
		checks.append(t, "TestCheckResourceAttr", "client_limit_up_enabled", fmt.Sprintf("%t", *s.ClientLimitUpEnabled))
	}
	if s.ClientLimitUp != nil {
		checks.append(t, "TestCheckResourceAttr", "client_limit_up", *s.ClientLimitUp)
	}

	// Auth server settings
	if s.AuthServerSelection != nil {
		checks.append(t, "TestCheckResourceAttr", "auth_server_selection", *s.AuthServerSelection)
	}
	if s.AuthServersNasId != nil {
		checks.append(t, "TestCheckResourceAttr", "auth_servers_nas_id", *s.AuthServersNasId)
	}
	if s.AuthServersNasIp != nil {
		checks.append(t, "TestCheckResourceAttr", "auth_servers_nas_ip", *s.AuthServersNasIp)
	}
	if s.AuthServersRetries != nil {
		checks.append(t, "TestCheckResourceAttr", "auth_servers_retries", fmt.Sprintf("%d", *s.AuthServersRetries))
	}
	if s.AuthServersTimeout != nil {
		checks.append(t, "TestCheckResourceAttr", "auth_servers_timeout", fmt.Sprintf("%d", *s.AuthServersTimeout))
	}

	// AP IDs array
	if len(s.ApIds) > 0 {
		checks.append(t, "TestCheckResourceAttr", "ap_ids.#", fmt.Sprintf("%d", len(s.ApIds)))
		for i, apId := range s.ApIds {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ap_ids.%d", i), apId)
		}
	}

	// Bands array
	if len(s.Bands) > 0 {
		checks.append(t, "TestCheckResourceAttr", "bands.#", fmt.Sprintf("%d", len(s.Bands)))
		for i, band := range s.Bands {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bands.%d", i), band)
		}
	}

	// ApplyTo string
	if s.ApplyTo != nil {
		checks.append(t, "TestCheckResourceAttr", "apply_to", *s.ApplyTo)
	}

	// WX Tunnel settings
	if s.WxtunnelId != nil {
		checks.append(t, "TestCheckResourceAttr", "wxtunnel_id", *s.WxtunnelId)
	}
	if s.WxtunnelRemoteId != nil {
		checks.append(t, "TestCheckResourceAttr", "wxtunnel_remote_id", *s.WxtunnelRemoteId)
	}

	// Tunnel disable settings
	if s.DisableWhenGatewayUnreachable != nil {
		checks.append(t, "TestCheckResourceAttr", "disable_when_gateway_unreachable", fmt.Sprintf("%t", *s.DisableWhenGatewayUnreachable))
	}
	if s.DisableWhenMxtunnelDown != nil {
		checks.append(t, "TestCheckResourceAttr", "disable_when_mxtunnel_down", fmt.Sprintf("%t", *s.DisableWhenMxtunnelDown))
	}

	// MX Tunnel arrays
	if len(s.MxtunnelIds) > 0 {
		checks.append(t, "TestCheckResourceAttr", "mxtunnel_ids.#", fmt.Sprintf("%d", len(s.MxtunnelIds)))
		for i, mxtunnelId := range s.MxtunnelIds {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("mxtunnel_ids.%d", i), mxtunnelId)
		}
	}
	if len(s.MxtunnelName) > 0 {
		checks.append(t, "TestCheckResourceAttr", "mxtunnel_name.#", fmt.Sprintf("%d", len(s.MxtunnelName)))
		for i, mxtunnelName := range s.MxtunnelName {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("mxtunnel_name.%d", i), mxtunnelName)
		}
	}

	// Portal hostname arrays
	if len(s.PortalAllowedHostnames) > 0 {
		checks.append(t, "TestCheckResourceAttr", "portal_allowed_hostnames.#", fmt.Sprintf("%d", len(s.PortalAllowedHostnames)))
		for i, hostname := range s.PortalAllowedHostnames {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_allowed_hostnames.%d", i), hostname)
		}
	}
	if len(s.PortalAllowedSubnets) > 0 {
		checks.append(t, "TestCheckResourceAttr", "portal_allowed_subnets.#", fmt.Sprintf("%d", len(s.PortalAllowedSubnets)))
		for i, subnet := range s.PortalAllowedSubnets {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_allowed_subnets.%d", i), subnet)
		}
	}
	if len(s.PortalDeniedHostnames) > 0 {
		checks.append(t, "TestCheckResourceAttr", "portal_denied_hostnames.#", fmt.Sprintf("%d", len(s.PortalDeniedHostnames)))
		for i, hostname := range s.PortalDeniedHostnames {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_denied_hostnames.%d", i), hostname)
		}
	}

	// WX Tag IDs array
	if len(s.WxtagIds) > 0 {
		checks.append(t, "TestCheckResourceAttr", "wxtag_ids.#", fmt.Sprintf("%d", len(s.WxtagIds)))
		for i, wxtagId := range s.WxtagIds {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("wxtag_ids.%d", i), wxtagId)
		}
	}

	// Auth servers array
	if len(s.AuthServers) > 0 {
		checks.append(t, "TestCheckResourceAttr", "auth_servers.#", fmt.Sprintf("%d", len(s.AuthServers)))
		for i, server := range s.AuthServers {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("auth_servers.%d.host", i), server.Host)
			if server.Port != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("auth_servers.%d.port", i), *server.Port)
			}
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("auth_servers.%d.secret", i), server.Secret)
			if server.KeywrapEnabled != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("auth_servers.%d.keywrap_enabled", i), fmt.Sprintf("%t", *server.KeywrapEnabled))
			}
			if server.KeywrapFormat != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("auth_servers.%d.keywrap_format", i), *server.KeywrapFormat)
			}
			if server.KeywrapKek != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("auth_servers.%d.keywrap_kek", i), *server.KeywrapKek)
			}
			if server.KeywrapMack != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("auth_servers.%d.keywrap_mack", i), *server.KeywrapMack)
			}
			if server.RequireMessageAuthenticator != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("auth_servers.%d.require_message_authenticator", i), fmt.Sprintf("%t", *server.RequireMessageAuthenticator))
			}
		}
	}

	// Acct servers array
	if len(s.AcctServers) > 0 {
		checks.append(t, "TestCheckResourceAttr", "acct_servers.#", fmt.Sprintf("%d", len(s.AcctServers)))
		for i, server := range s.AcctServers {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acct_servers.%d.host", i), server.Host)
			if server.Port != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acct_servers.%d.port", i), *server.Port)
			}
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acct_servers.%d.secret", i), server.Secret)
			if server.KeywrapEnabled != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acct_servers.%d.keywrap_enabled", i), fmt.Sprintf("%t", *server.KeywrapEnabled))
			}
			if server.KeywrapFormat != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acct_servers.%d.keywrap_format", i), *server.KeywrapFormat)
			}
			if server.KeywrapKek != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acct_servers.%d.keywrap_kek", i), *server.KeywrapKek)
			}
			if server.KeywrapMack != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acct_servers.%d.keywrap_mack", i), *server.KeywrapMack)
			}
		}
	}

	// CoA servers array validation
	if len(s.CoaServers) > 0 {
		checks.append(t, "TestCheckResourceAttr", "coa_servers.#", fmt.Sprintf("%d", len(s.CoaServers)))
		for i, server := range s.CoaServers {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("coa_servers.%d.ip", i), server.Ip)
			if server.Port != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("coa_servers.%d.port", i), *server.Port)
			}
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("coa_servers.%d.secret", i), server.Secret)
			if server.Enabled != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("coa_servers.%d.enabled", i), fmt.Sprintf("%t", *server.Enabled))
			}
			if server.DisableEventTimestampCheck != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("coa_servers.%d.disable_event_timestamp_check", i), fmt.Sprintf("%t", *server.DisableEventTimestampCheck))
			}
		}
	}

	if s.Hotspot20 != nil {
		if s.Hotspot20.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "hotspot20.enabled", fmt.Sprintf("%t", *s.Hotspot20.Enabled))
		}
		if len(s.Hotspot20.DomainName) > 0 {
			checks.append(t, "TestCheckResourceAttr", "hotspot20.domain_name.#", fmt.Sprintf("%d", len(s.Hotspot20.DomainName)))
			for i, domain := range s.Hotspot20.DomainName {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("hotspot20.domain_name.%d", i), domain)
			}
		}
		if len(s.Hotspot20.NaiRealms) > 0 {
			checks.append(t, "TestCheckResourceAttr", "hotspot20.nai_realms.#", fmt.Sprintf("%d", len(s.Hotspot20.NaiRealms)))
			for i, realm := range s.Hotspot20.NaiRealms {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("hotspot20.nai_realms.%d", i), realm)
			}
		}
		if len(s.Hotspot20.Operators) > 0 {
			checks.append(t, "TestCheckResourceAttr", "hotspot20.operators.#", fmt.Sprintf("%d", len(s.Hotspot20.Operators)))
			for i, operator := range s.Hotspot20.Operators {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("hotspot20.operators.%d", i), operator)
			}
		}
		if len(s.Hotspot20.Rcoi) > 0 {
			checks.append(t, "TestCheckResourceAttr", "hotspot20.rcoi.#", fmt.Sprintf("%d", len(s.Hotspot20.Rcoi)))
			for i, rcoi := range s.Hotspot20.Rcoi {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("hotspot20.rcoi.%d", i), rcoi)
			}
		}
		if s.Hotspot20.VenueName != nil {
			checks.append(t, "TestCheckResourceAttr", "hotspot20.venue_name", *s.Hotspot20.VenueName)
		}
	}

	if s.InjectDhcpOption82 != nil {
		if s.InjectDhcpOption82.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "inject_dhcp_option_82.enabled", fmt.Sprintf("%t", *s.InjectDhcpOption82.Enabled))
		}
		if s.InjectDhcpOption82.CircuitId != nil {
			checks.append(t, "TestCheckResourceAttr", "inject_dhcp_option_82.circuit_id", *s.InjectDhcpOption82.CircuitId)
		}
	}

	if s.AppLimit != nil {
		if s.AppLimit.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "app_limit.enabled", fmt.Sprintf("%t", *s.AppLimit.Enabled))
		}
		if s.AppLimit.Apps != nil {
			for appName, limit := range s.AppLimit.Apps {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("app_limit.apps.%s", appName), fmt.Sprintf("%d", limit))
			}
		}
		if s.AppLimit.WxtagIds != nil {
			for wxtagName, limit := range s.AppLimit.WxtagIds {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("app_limit.wxtag_ids.%s", wxtagName), fmt.Sprintf("%d", limit))
			}
		}
	}

	// App QoS settings
	if s.AppQos != nil {
		if s.AppQos.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "app_qos.enabled", fmt.Sprintf("%t", *s.AppQos.Enabled))
		}
		if s.AppQos.Apps != nil {
			for appName, appConfig := range s.AppQos.Apps {
				if appConfig.Dscp != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("app_qos.apps.%s.dscp", appName), *appConfig.Dscp)
				}
				if appConfig.DstSubnet != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("app_qos.apps.%s.dst_subnet", appName), *appConfig.DstSubnet)
				}
				if appConfig.SrcSubnet != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("app_qos.apps.%s.src_subnet", appName), *appConfig.SrcSubnet)
				}
			}
		}
		if s.AppQos.Others != nil {
			checks.append(t, "TestCheckResourceAttr", "app_qos.others.#", fmt.Sprintf("%d", len(s.AppQos.Others)))
			for i, othersConfig := range s.AppQos.Others {
				if othersConfig.Dscp != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("app_qos.others.%d.dscp", i), *othersConfig.Dscp)
				}
				if othersConfig.DstSubnet != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("app_qos.others.%d.dst_subnet", i), *othersConfig.DstSubnet)
				}
				if othersConfig.PortRanges != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("app_qos.others.%d.port_ranges", i), *othersConfig.PortRanges)
				}
				if othersConfig.Protocol != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("app_qos.others.%d.protocol", i), *othersConfig.Protocol)
				}
				if othersConfig.SrcSubnet != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("app_qos.others.%d.src_subnet", i), *othersConfig.SrcSubnet)
				}
			}
		}
	}

	// QoS settings
	if s.Qos != nil {
		if s.Qos.Class != nil {
			checks.append(t, "TestCheckResourceAttr", "qos.class", *s.Qos.Class)
		}
		if s.Qos.Overwrite != nil {
			checks.append(t, "TestCheckResourceAttr", "qos.overwrite", fmt.Sprintf("%t", *s.Qos.Overwrite))
		}
	}

	// Schedule settings
	if s.Schedule != nil {
		checks.append(t, "TestCheckResourceAttrSet", "schedule.%")
		if s.Schedule.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "schedule.enabled", fmt.Sprintf("%t", *s.Schedule.Enabled))
		}
		if s.Schedule.Hours != nil {
			checks.append(t, "TestCheckResourceAttrSet", "schedule.hours.%")
			if s.Schedule.Hours.Mon != nil {
				checks.append(t, "TestCheckResourceAttr", "schedule.hours.mon", *s.Schedule.Hours.Mon)
			}
			if s.Schedule.Hours.Tue != nil {
				checks.append(t, "TestCheckResourceAttr", "schedule.hours.tue", *s.Schedule.Hours.Tue)
			}
			if s.Schedule.Hours.Wed != nil {
				checks.append(t, "TestCheckResourceAttr", "schedule.hours.wed", *s.Schedule.Hours.Wed)
			}
			if s.Schedule.Hours.Thu != nil {
				checks.append(t, "TestCheckResourceAttr", "schedule.hours.thu", *s.Schedule.Hours.Thu)
			}
			if s.Schedule.Hours.Fri != nil {
				checks.append(t, "TestCheckResourceAttr", "schedule.hours.fri", *s.Schedule.Hours.Fri)
			}
			if s.Schedule.Hours.Sat != nil {
				checks.append(t, "TestCheckResourceAttr", "schedule.hours.sat", *s.Schedule.Hours.Sat)
			}
			if s.Schedule.Hours.Sun != nil {
				checks.append(t, "TestCheckResourceAttr", "schedule.hours.sun", *s.Schedule.Hours.Sun)
			}
		}
	}

	// Auth object validation
	if s.Auth != nil {
		checks.append(t, "TestCheckResourceAttrSet", "auth.%")
		if s.Auth.AuthType != nil {
			checks.append(t, "TestCheckResourceAttr", "auth.type", *s.Auth.AuthType)
		}
		if s.Auth.EnableMacAuth != nil {
			checks.append(t, "TestCheckResourceAttr", "auth.enable_mac_auth", fmt.Sprintf("%t", *s.Auth.EnableMacAuth))
		}
		if len(s.Auth.Pairwise) > 0 {
			checks.append(t, "TestCheckResourceAttr", "auth.pairwise.#", fmt.Sprintf("%d", len(s.Auth.Pairwise)))
			for i, pairwise := range s.Auth.Pairwise {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("auth.pairwise.%d", i), pairwise)
			}
		}
		if s.Auth.AnticlogThreshold != nil {
			checks.append(t, "TestCheckResourceAttr", "auth.anticlog_threshold", fmt.Sprintf("%d", *s.Auth.AnticlogThreshold))
		}
		if s.Auth.EapReauth != nil {
			checks.append(t, "TestCheckResourceAttr", "auth.eap_reauth", fmt.Sprintf("%t", *s.Auth.EapReauth))
		}
		if s.Auth.KeyIdx != nil {
			checks.append(t, "TestCheckResourceAttr", "auth.key_idx", fmt.Sprintf("%d", *s.Auth.KeyIdx))
		}
		if s.Auth.Keys != nil {
			checks.append(t, "TestCheckResourceAttr", "auth.keys.#", fmt.Sprintf("%d", len(s.Auth.Keys)))
			for i, key := range s.Auth.Keys {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("auth.keys.%d", i), key)
			}
		}
		if s.Auth.MultiPskOnly != nil {
			checks.append(t, "TestCheckResourceAttr", "auth.multi_psk_only", fmt.Sprintf("%t", *s.Auth.MultiPskOnly))
		}
		if s.Auth.Owe != nil {
			checks.append(t, "TestCheckResourceAttr", "auth.owe", *s.Auth.Owe)
		}
		if s.Auth.Psk != nil {
			checks.append(t, "TestCheckResourceAttr", "auth.psk", *s.Auth.Psk)
		}
		if s.Auth.WepAsSecondaryAuth != nil {
			checks.append(t, "TestCheckResourceAttr", "auth.wep_as_secondary_auth", fmt.Sprintf("%t", *s.Auth.WepAsSecondaryAuth))
		}
		if s.Auth.PrivateWlan != nil {
			checks.append(t, "TestCheckResourceAttr", "auth.private_wlan", fmt.Sprintf("%t", *s.Auth.PrivateWlan))
		}
	}

	// Airwatch object validation
	if s.Airwatch != nil {
		checks.append(t, "TestCheckResourceAttrSet", "airwatch.%")
		if s.Airwatch.ApiKey != nil {
			checks.append(t, "TestCheckResourceAttr", "airwatch.api_key", *s.Airwatch.ApiKey)
		}
		if s.Airwatch.ConsoleUrl != nil {
			checks.append(t, "TestCheckResourceAttr", "airwatch.console_url", *s.Airwatch.ConsoleUrl)
		}
		if s.Airwatch.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "airwatch.enabled", fmt.Sprintf("%t", *s.Airwatch.Enabled))
		}
		if s.Airwatch.Password != nil {
			checks.append(t, "TestCheckResourceAttr", "airwatch.password", *s.Airwatch.Password)
		}
		if s.Airwatch.Username != nil {
			checks.append(t, "TestCheckResourceAttr", "airwatch.username", *s.Airwatch.Username)
		}
	}

	// Bonjour object validation
	if s.Bonjour != nil {
		checks.append(t, "TestCheckResourceAttrSet", "bonjour.%")
		if s.Bonjour.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "bonjour.enabled", fmt.Sprintf("%t", *s.Bonjour.Enabled))
		}
		if s.Bonjour.AdditionalVlanIds != nil {
			checks.append(t, "TestCheckResourceAttr", "bonjour.additional_vlan_ids.#", fmt.Sprintf("%d", len(s.Bonjour.AdditionalVlanIds)))
			for i, vlan := range s.Bonjour.AdditionalVlanIds {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bonjour.additional_vlan_ids.%d", i), vlan)
			}
		}
		if s.Bonjour.Services != nil {
			checks.append(t, "TestCheckResourceAttrSet", "bonjour.services.%")
			serviceCount := 0
			for serviceName, service := range s.Bonjour.Services {
				if service.DisableLocal != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bonjour.services.%s.disable_local", serviceName), fmt.Sprintf("%t", *service.DisableLocal))
				}
				if service.RadiusGroups != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bonjour.services.%s.radius_groups.#", serviceName), fmt.Sprintf("%d", len(service.RadiusGroups)))
					for i, group := range service.RadiusGroups {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bonjour.services.%s.radius_groups.%d", serviceName, i), group)
					}
				}
				if service.Scope != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bonjour.services.%s.scope", serviceName), *service.Scope)
				}
				serviceCount++
			}
		}
	}

	// CiscoCwa object validation
	if s.CiscoCwa != nil {
		if s.CiscoCwa.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "cisco_cwa.enabled", fmt.Sprintf("%t", *s.CiscoCwa.Enabled))
		}
		if s.CiscoCwa.AllowedHostnames != nil {
			checks.append(t, "TestCheckResourceAttr", "cisco_cwa.allowed_hostnames.#", fmt.Sprintf("%d", len(s.CiscoCwa.AllowedHostnames)))
			for i, hostname := range s.CiscoCwa.AllowedHostnames {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("cisco_cwa.allowed_hostnames.%d", i), hostname)
			}
		}
		if s.CiscoCwa.AllowedSubnets != nil {
			checks.append(t, "TestCheckResourceAttr", "cisco_cwa.allowed_subnets.#", fmt.Sprintf("%d", len(s.CiscoCwa.AllowedSubnets)))
			for i, subnet := range s.CiscoCwa.AllowedSubnets {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("cisco_cwa.allowed_subnets.%d", i), subnet)
			}
		}
		if s.CiscoCwa.BlockedSubnets != nil {
			checks.append(t, "TestCheckResourceAttr", "cisco_cwa.blocked_subnets.#", fmt.Sprintf("%d", len(s.CiscoCwa.BlockedSubnets)))
			for i, subnet := range s.CiscoCwa.BlockedSubnets {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("cisco_cwa.blocked_subnets.%d", i), subnet)
			}
		}
	}

	// DnsServerRewrite object validation
	if s.DnsServerRewrite != nil {
		if s.DnsServerRewrite.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "dns_server_rewrite.enabled", fmt.Sprintf("%t", *s.DnsServerRewrite.Enabled))
		}
		if s.DnsServerRewrite.RadiusGroups != nil {
			for groupName, dnsServer := range s.DnsServerRewrite.RadiusGroups {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dns_server_rewrite.radius_groups.%s", groupName), dnsServer)
			}
		}
	}

	// DynamicPsk object validation
	if s.DynamicPsk != nil {
		if s.DynamicPsk.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "dynamic_psk.enabled", fmt.Sprintf("%t", *s.DynamicPsk.Enabled))
		}
		if s.DynamicPsk.Source != nil {
			checks.append(t, "TestCheckResourceAttr", "dynamic_psk.source", *s.DynamicPsk.Source)
		}
		if s.DynamicPsk.DefaultPsk != nil {
			checks.append(t, "TestCheckResourceAttr", "dynamic_psk.default_psk", *s.DynamicPsk.DefaultPsk)
		}
		if s.DynamicPsk.DefaultVlanId != nil {
			checks.append(t, "TestCheckResourceAttr", "dynamic_psk.default_vlan_id", *s.DynamicPsk.DefaultVlanId)
		}
		if s.DynamicPsk.ForceLookup != nil {
			checks.append(t, "TestCheckResourceAttr", "dynamic_psk.force_lookup", fmt.Sprintf("%t", *s.DynamicPsk.ForceLookup))
		}
	}

	// DynamicVlan object validation
	if s.DynamicVlan != nil {
		if s.DynamicVlan.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "dynamic_vlan.enabled", fmt.Sprintf("%t", *s.DynamicVlan.Enabled))
		}
		if s.DynamicVlan.DynamicVlanType != nil {
			checks.append(t, "TestCheckResourceAttr", "dynamic_vlan.type", *s.DynamicVlan.DynamicVlanType)
		}
		if len(s.DynamicVlan.DefaultVlanIds) > 0 {
			checks.append(t, "TestCheckResourceAttr", "dynamic_vlan.default_vlan_ids.#", fmt.Sprintf("%d", len(s.DynamicVlan.DefaultVlanIds)))
			for i, vlanId := range s.DynamicVlan.DefaultVlanIds {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dynamic_vlan.default_vlan_ids.%d", i), vlanId)
			}
		}
		if len(s.DynamicVlan.LocalVlanIds) > 0 {
			checks.append(t, "TestCheckResourceAttr", "dynamic_vlan.local_vlan_ids.#", fmt.Sprintf("%d", len(s.DynamicVlan.LocalVlanIds)))
			for i, vlanId := range s.DynamicVlan.LocalVlanIds {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dynamic_vlan.local_vlan_ids.%d", i), vlanId)
			}
		}
		if len(s.DynamicVlan.Vlans) > 0 {
			checks.append(t, "TestCheckResourceAttr", "dynamic_vlan.vlans.%", fmt.Sprintf("%d", len(s.DynamicVlan.Vlans)))
			for vlanId, interfaceName := range s.DynamicVlan.Vlans {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dynamic_vlan.vlans.%s", vlanId), interfaceName)
			}
		}
	}

	// MistNac object validation
	if s.MistNac != nil {
		if s.MistNac.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "mist_nac.enabled", fmt.Sprintf("%t", *s.MistNac.Enabled))
		}
	}

	// Portal API secret validation - only check when portal.auth is set to "external"
	if s.Portal != nil && s.Portal.Auth != nil && *s.Portal.Auth == "external" {
		checks.append(t, "TestCheckResourceAttrSet", "portal_api_secret")
	}

	// Portal SSO URL validation - only check when portal.auth is set to "sso"
	if s.Portal != nil && s.Portal.Auth != nil && *s.Portal.Auth == "sso" {
		checks.append(t, "TestCheckResourceAttrSet", "portal_sso_url")
	}

	// NOTE: portal_image field testing is handled in TestOrgWlanPortalImageModel
	// since it requires portal image resource creation and API propagation timing

	// Portal object validation
	if s.Portal != nil {
		checks.append(t, "TestCheckResourceAttrSet", "portal.%")
		if s.Portal.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.enabled", fmt.Sprintf("%t", *s.Portal.Enabled))
		}
		if s.Portal.Auth != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.auth", *s.Portal.Auth)
		}

		// Amazon OAuth settings
		if s.Portal.AmazonEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.amazon_enabled", fmt.Sprintf("%t", *s.Portal.AmazonEnabled))
		}
		if s.Portal.AmazonClientId != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.amazon_client_id", *s.Portal.AmazonClientId)
		}
		if s.Portal.AmazonClientSecret != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.amazon_client_secret", *s.Portal.AmazonClientSecret)
		}
		if s.Portal.AmazonExpire != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.amazon_expire", fmt.Sprintf("%d", *s.Portal.AmazonExpire))
		}
		if len(s.Portal.AmazonEmailDomains) > 0 {
			checks.append(t, "TestCheckResourceAttr", "portal.amazon_email_domains.#", fmt.Sprintf("%d", len(s.Portal.AmazonEmailDomains)))
			for i, domain := range s.Portal.AmazonEmailDomains {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal.amazon_email_domains.%d", i), domain)
			}
		}

		// Azure OAuth settings
		if s.Portal.AzureEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.azure_enabled", fmt.Sprintf("%t", *s.Portal.AzureEnabled))
		}
		if s.Portal.AzureClientId != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.azure_client_id", *s.Portal.AzureClientId)
		}
		if s.Portal.AzureClientSecret != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.azure_client_secret", *s.Portal.AzureClientSecret)
		}
		if s.Portal.AzureExpire != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.azure_expire", fmt.Sprintf("%d", *s.Portal.AzureExpire))
		}
		if s.Portal.AzureTenantId != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.azure_tenant_id", *s.Portal.AzureTenantId)
		}

		// Facebook OAuth settings
		if s.Portal.FacebookEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.facebook_enabled", fmt.Sprintf("%t", *s.Portal.FacebookEnabled))
		}
		if s.Portal.FacebookClientId != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.facebook_client_id", *s.Portal.FacebookClientId)
		}
		if s.Portal.FacebookClientSecret != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.facebook_client_secret", *s.Portal.FacebookClientSecret)
		}
		if s.Portal.FacebookExpire != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.facebook_expire", fmt.Sprintf("%d", *s.Portal.FacebookExpire))
		}
		if len(s.Portal.FacebookEmailDomains) > 0 {
			checks.append(t, "TestCheckResourceAttr", "portal.facebook_email_domains.#", fmt.Sprintf("%d", len(s.Portal.FacebookEmailDomains)))
			for i, domain := range s.Portal.FacebookEmailDomains {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal.facebook_email_domains.%d", i), domain)
			}
		}

		// Google OAuth settings
		if s.Portal.GoogleEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.google_enabled", fmt.Sprintf("%t", *s.Portal.GoogleEnabled))
		}
		if s.Portal.GoogleClientId != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.google_client_id", *s.Portal.GoogleClientId)
		}
		if s.Portal.GoogleClientSecret != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.google_client_secret", *s.Portal.GoogleClientSecret)
		}
		if s.Portal.GoogleExpire != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.google_expire", fmt.Sprintf("%d", *s.Portal.GoogleExpire))
		}
		if len(s.Portal.GoogleEmailDomains) > 0 {
			checks.append(t, "TestCheckResourceAttr", "portal.google_email_domains.#", fmt.Sprintf("%d", len(s.Portal.GoogleEmailDomains)))
			for i, domain := range s.Portal.GoogleEmailDomains {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal.google_email_domains.%d", i), domain)
			}
		}

		// Microsoft OAuth settings
		if s.Portal.MicrosoftEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.microsoft_enabled", fmt.Sprintf("%t", *s.Portal.MicrosoftEnabled))
		}
		if s.Portal.MicrosoftClientId != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.microsoft_client_id", *s.Portal.MicrosoftClientId)
		}
		if s.Portal.MicrosoftClientSecret != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.microsoft_client_secret", *s.Portal.MicrosoftClientSecret)
		}
		if s.Portal.MicrosoftExpire != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.microsoft_expire", fmt.Sprintf("%d", *s.Portal.MicrosoftExpire))
		}
		if len(s.Portal.MicrosoftEmailDomains) > 0 {
			checks.append(t, "TestCheckResourceAttr", "portal.microsoft_email_domains.#", fmt.Sprintf("%d", len(s.Portal.MicrosoftEmailDomains)))
			for i, domain := range s.Portal.MicrosoftEmailDomains {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal.microsoft_email_domains.%d", i), domain)
			}
		}

		// SMS settings
		if s.Portal.SmsEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.sms_enabled", fmt.Sprintf("%t", *s.Portal.SmsEnabled))
		}
		if s.Portal.SmsExpire != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.sms_expire", fmt.Sprintf("%d", *s.Portal.SmsExpire))
		}
		if s.Portal.SmsProvider != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.sms_provider", *s.Portal.SmsProvider)
		}
		if s.Portal.SmsMessageFormat != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.sms_message_format", *s.Portal.SmsMessageFormat)
		}

		// SMSGlobal settings
		if s.Portal.SmsglobalApiKey != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.smsglobal_api_key", *s.Portal.SmsglobalApiKey)
		}
		if s.Portal.SmsglobalApiSecret != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.smsglobal_api_secret", *s.Portal.SmsglobalApiSecret)
		}

		// Twilio settings
		if s.Portal.TwilioAuthToken != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.twilio_auth_token", *s.Portal.TwilioAuthToken)
		}
		if s.Portal.TwilioPhoneNumber != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.twilio_phone_number", *s.Portal.TwilioPhoneNumber)
		}
		if s.Portal.TwilioSid != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.twilio_sid", *s.Portal.TwilioSid)
		}

		// Other SMS providers
		if s.Portal.ClickatellApiKey != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.clickatell_api_key", *s.Portal.ClickatellApiKey)
		}
		if s.Portal.GupshupPassword != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.gupshup_password", *s.Portal.GupshupPassword)
		}
		if s.Portal.GupshupUserid != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.gupshup_userid", *s.Portal.GupshupUserid)
		}
		if s.Portal.PuzzelPassword != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.puzzel_password", *s.Portal.PuzzelPassword)
		}
		if s.Portal.PuzzelServiceId != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.puzzel_service_id", *s.Portal.PuzzelServiceId)
		}
		if s.Portal.PuzzelUsername != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.puzzel_username", *s.Portal.PuzzelUsername)
		}
		if s.Portal.TelstraClientId != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.telstra_client_id", *s.Portal.TelstraClientId)
		}
		if s.Portal.TelstraClientSecret != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.telstra_client_secret", *s.Portal.TelstraClientSecret)
		}

		// Broadnet settings
		if s.Portal.BroadnetPassword != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.broadnet_password", *s.Portal.BroadnetPassword)
		}
		if s.Portal.BroadnetSid != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.broadnet_sid", *s.Portal.BroadnetSid)
		}
		if s.Portal.BroadnetUserId != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.broadnet_user_id", *s.Portal.BroadnetUserId)
		}

		// Email and passphrase settings
		if s.Portal.EmailEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.email_enabled", fmt.Sprintf("%t", *s.Portal.EmailEnabled))
		}
		if s.Portal.PassphraseEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.passphrase_enabled", fmt.Sprintf("%t", *s.Portal.PassphraseEnabled))
		}
		if s.Portal.PassphraseExpire != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.passphrase_expire", fmt.Sprintf("%d", *s.Portal.PassphraseExpire))
		}

		// Sponsor settings
		if s.Portal.SponsorEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.sponsor_enabled", fmt.Sprintf("%t", *s.Portal.SponsorEnabled))
		}
		if s.Portal.SponsorExpire != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.sponsor_expire", fmt.Sprintf("%d", *s.Portal.SponsorExpire))
		}
		if s.Portal.SponsorAutoApprove != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.sponsor_auto_approve", fmt.Sprintf("%t", *s.Portal.SponsorAutoApprove))
		}
		if s.Portal.SponsorNotifyAll != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.sponsor_notify_all", fmt.Sprintf("%t", *s.Portal.SponsorNotifyAll))
		}
		if s.Portal.SponsorStatusNotify != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.sponsor_status_notify", fmt.Sprintf("%t", *s.Portal.SponsorStatusNotify))
		}
		if s.Portal.SponsorLinkValidityDuration != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.sponsor_link_validity_duration", *s.Portal.SponsorLinkValidityDuration)
		}
		if s.Portal.PredefinedSponsorsEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.predefined_sponsors_enabled", fmt.Sprintf("%t", *s.Portal.PredefinedSponsorsEnabled))
		}
		if s.Portal.PredefinedSponsorsHideEmail != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.predefined_sponsors_hide_email", fmt.Sprintf("%t", *s.Portal.PredefinedSponsorsHideEmail))
		}

		// Sponsors map
		if len(s.Portal.Sponsors) > 0 {
			checks.append(t, "TestCheckResourceAttr", "portal.sponsors.%", fmt.Sprintf("%d", len(s.Portal.Sponsors)))
			for key, sponsor := range s.Portal.Sponsors {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal.sponsors.%s", key), sponsor)
			}
		}
		if len(s.Portal.SponsorEmailDomains) > 0 {
			checks.append(t, "TestCheckResourceAttr", "portal.sponsor_email_domains.#", fmt.Sprintf("%d", len(s.Portal.SponsorEmailDomains)))
			for i, domain := range s.Portal.SponsorEmailDomains {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal.sponsor_email_domains.%d", i), domain)
			}
		}

		// SSO settings
		if s.Portal.SsoDefaultRole != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.sso_default_role", *s.Portal.SsoDefaultRole)
		}
		if s.Portal.SsoForcedRole != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.sso_forced_role", *s.Portal.SsoForcedRole)
		}
		if s.Portal.SsoIdpCert != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.sso_idp_cert", *s.Portal.SsoIdpCert)
		}
		if s.Portal.SsoIdpSignAlgo != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.sso_idp_sign_algo", *s.Portal.SsoIdpSignAlgo)
		}
		if s.Portal.SsoIdpSsoUrl != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.sso_idp_sso_url", *s.Portal.SsoIdpSsoUrl)
		}
		if s.Portal.SsoIssuer != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.sso_issuer", *s.Portal.SsoIssuer)
		}
		if s.Portal.SsoNameidFormat != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.sso_nameid_format", *s.Portal.SsoNameidFormat)
		}

		// Portal behavior settings
		if s.Portal.AllowWlanIdRoam != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.allow_wlan_id_roam", fmt.Sprintf("%t", *s.Portal.AllowWlanIdRoam))
		}
		if s.Portal.BypassWhenCloudDown != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.bypass_when_cloud_down", fmt.Sprintf("%t", *s.Portal.BypassWhenCloudDown))
		}
		if s.Portal.CrossSite != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.cross_site", fmt.Sprintf("%t", *s.Portal.CrossSite))
		}
		if s.Portal.Forward != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.forward", fmt.Sprintf("%t", *s.Portal.Forward))
		}
		if s.Portal.ForwardUrl != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.forward_url", *s.Portal.ForwardUrl)
		}
		if s.Portal.ExternalPortalUrl != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.external_portal_url", *s.Portal.ExternalPortalUrl)
		}
		if s.Portal.Privacy != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.privacy", fmt.Sprintf("%t", *s.Portal.Privacy))
		}

		// Session settings
		if s.Portal.Expire != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.expire", fmt.Sprintf("%d", *s.Portal.Expire))
		}
		if s.Portal.Password != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.password", *s.Portal.Password)
		}
	}

	// Radsec object validation
	if s.Radsec != nil {
		if s.Radsec.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "radsec.enabled", fmt.Sprintf("%t", *s.Radsec.Enabled))
		}
		if s.Radsec.CoaEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "radsec.coa_enabled", fmt.Sprintf("%t", *s.Radsec.CoaEnabled))
		}
		if s.Radsec.IdleTimeout != nil {
			checks.append(t, "TestCheckResourceAttr", "radsec.idle_timeout", *s.Radsec.IdleTimeout)
		}
		if s.Radsec.ServerName != nil {
			checks.append(t, "TestCheckResourceAttr", "radsec.server_name", *s.Radsec.ServerName)
		}
		if s.Radsec.UseMxedge != nil {
			checks.append(t, "TestCheckResourceAttr", "radsec.use_mxedge", fmt.Sprintf("%t", *s.Radsec.UseMxedge))
		}
		if s.Radsec.UseSiteMxedge != nil {
			checks.append(t, "TestCheckResourceAttr", "radsec.use_site_mxedge", fmt.Sprintf("%t", *s.Radsec.UseSiteMxedge))
		}

		// Validate MxclusterIds array
		if len(s.Radsec.MxclusterIds) > 0 {
			checks.append(t, "TestCheckResourceAttr", "radsec.mxcluster_ids.#", fmt.Sprintf("%d", len(s.Radsec.MxclusterIds)))
			for i, id := range s.Radsec.MxclusterIds {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radsec.mxcluster_ids.%d", i), id)
			}
		}

		// Validate ProxyHosts array
		if len(s.Radsec.ProxyHosts) > 0 {
			checks.append(t, "TestCheckResourceAttr", "radsec.proxy_hosts.#", fmt.Sprintf("%d", len(s.Radsec.ProxyHosts)))
			for i, host := range s.Radsec.ProxyHosts {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radsec.proxy_hosts.%d", i), host)
			}
		}

		// Validate Servers array
		if len(s.Radsec.Servers) > 0 {
			checks.append(t, "TestCheckResourceAttr", "radsec.servers.#", fmt.Sprintf("%d", len(s.Radsec.Servers)))
			for i, server := range s.Radsec.Servers {
				if server.Host != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radsec.servers.%d.host", i), *server.Host)
				}
				if server.Port != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radsec.servers.%d.port", i), fmt.Sprintf("%d", *server.Port))
				}
			}
		}
	}

	// Rateset map validation
	if len(s.Rateset) > 0 {
		for band, rateset := range s.Rateset {
			if rateset.Template != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("rateset.%s.template", band), *rateset.Template)
			}
			if rateset.MinRssi != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("rateset.%s.min_rssi", band), fmt.Sprintf("%d", *rateset.MinRssi))
			}
			if len(rateset.Legacy) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("rateset.%s.legacy.#", band), fmt.Sprintf("%d", len(rateset.Legacy)))
				for i, legacy := range rateset.Legacy {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("rateset.%s.legacy.%d", band, i), legacy)
				}
			}
			if rateset.Ht != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("rateset.%s.ht", band), *rateset.Ht)
			}
			if rateset.Vht != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("rateset.%s.vht", band), *rateset.Vht)
			}
			if rateset.He != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("rateset.%s.he", band), *rateset.He)
			}
			if rateset.Eht != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("rateset.%s.eht", band), *rateset.Eht)
			}
		}
	}

	// NOTE: AppLimit field testing skipped due to map[string]int64 type conversion issues
	// in the test framework. The field requires special handling for map types.

	return checks
}
