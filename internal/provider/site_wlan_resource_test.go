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

func TestSiteWlanModel(t *testing.T) {
	type testStep struct {
		config SiteWlanModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: SiteWlanModel{
						Ssid: "TestSSID",
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/site_wlan_resource/site_wlan_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		var FixtureSiteWlanModel SiteWlanModel
		err = hcl.Decode(&FixtureSiteWlanModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureSiteWlanModel,
				},
			},
		}
	}

	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			resourceType := "site_wlan"

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

func (s *SiteWlanModel) testChecks(t testing.TB, rType, rName string) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + rName)
	checks.append(t, "TestCheckResourceAttrSet", "site_id")
	checks.append(t, "TestCheckResourceAttr", "ssid", s.Ssid)

	// Boolean attributes
	if s.AcctImmediateUpdate != nil {
		checks.append(t, "TestCheckResourceAttr", "acct_immediate_update", fmt.Sprintf("%t", *s.AcctImmediateUpdate))
	}
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
	if s.BandSteer != nil {
		checks.append(t, "TestCheckResourceAttr", "band_steer", fmt.Sprintf("%t", *s.BandSteer))
	}
	if s.BandSteerForceBand5 != nil {
		checks.append(t, "TestCheckResourceAttr", "band_steer_force_band5", fmt.Sprintf("%t", *s.BandSteerForceBand5))
	}
	if s.BlockBlacklistClients != nil {
		checks.append(t, "TestCheckResourceAttr", "block_blacklist_clients", fmt.Sprintf("%t", *s.BlockBlacklistClients))
	}
	if s.ClientLimitDownEnabled != nil {
		checks.append(t, "TestCheckResourceAttr", "client_limit_down_enabled", fmt.Sprintf("%t", *s.ClientLimitDownEnabled))
	}
	if s.ClientLimitUpEnabled != nil {
		checks.append(t, "TestCheckResourceAttr", "client_limit_up_enabled", fmt.Sprintf("%t", *s.ClientLimitUpEnabled))
	}
	if s.Disable11ax != nil {
		checks.append(t, "TestCheckResourceAttr", "disable_11ax", fmt.Sprintf("%t", *s.Disable11ax))
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
	if s.DisableWhenGatewayUnreachable != nil {
		checks.append(t, "TestCheckResourceAttr", "disable_when_gateway_unreachable", fmt.Sprintf("%t", *s.DisableWhenGatewayUnreachable))
	}
	if s.DisableWhenMxtunnelDown != nil {
		checks.append(t, "TestCheckResourceAttr", "disable_when_mxtunnel_down", fmt.Sprintf("%t", *s.DisableWhenMxtunnelDown))
	}
	if s.DisableWmm != nil {
		checks.append(t, "TestCheckResourceAttr", "disable_wmm", fmt.Sprintf("%t", *s.DisableWmm))
	}
	if s.EnableLocalKeycaching != nil {
		checks.append(t, "TestCheckResourceAttr", "enable_local_keycaching", fmt.Sprintf("%t", *s.EnableLocalKeycaching))
	}
	if s.EnableWirelessBridging != nil {
		checks.append(t, "TestCheckResourceAttr", "enable_wireless_bridging", fmt.Sprintf("%t", *s.EnableWirelessBridging))
	}
	if s.EnableWirelessBridgingDhcpTracking != nil {
		checks.append(t, "TestCheckResourceAttr", "enable_wireless_bridging_dhcp_tracking", fmt.Sprintf("%t", *s.EnableWirelessBridgingDhcpTracking))
	}
	if s.Enabled != nil {
		checks.append(t, "TestCheckResourceAttr", "enabled", fmt.Sprintf("%t", *s.Enabled))
	}
	if s.FastDot1xTimers != nil {
		checks.append(t, "TestCheckResourceAttr", "fast_dot1x_timers", fmt.Sprintf("%t", *s.FastDot1xTimers))
	}
	if s.HideSsid != nil {
		checks.append(t, "TestCheckResourceAttr", "hide_ssid", fmt.Sprintf("%t", *s.HideSsid))
	}
	if s.HostnameIe != nil {
		checks.append(t, "TestCheckResourceAttr", "hostname_ie", fmt.Sprintf("%t", *s.HostnameIe))
	}
	if s.Isolation != nil {
		checks.append(t, "TestCheckResourceAttr", "isolation", fmt.Sprintf("%t", *s.Isolation))
	}
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
	if s.NoStaticDns != nil {
		checks.append(t, "TestCheckResourceAttr", "no_static_dns", fmt.Sprintf("%t", *s.NoStaticDns))
	}
	if s.NoStaticIp != nil {
		checks.append(t, "TestCheckResourceAttr", "no_static_ip", fmt.Sprintf("%t", *s.NoStaticIp))
	}
	if s.ReconnectClientsWhenRoamingMxcluster != nil {
		checks.append(t, "TestCheckResourceAttr", "reconnect_clients_when_roaming_mxcluster", fmt.Sprintf("%t", *s.ReconnectClientsWhenRoamingMxcluster))
	}
	if s.SleExcluded != nil {
		checks.append(t, "TestCheckResourceAttr", "sle_excluded", fmt.Sprintf("%t", *s.SleExcluded))
	}
	if s.UseEapolV1 != nil {
		checks.append(t, "TestCheckResourceAttr", "use_eapol_v1", fmt.Sprintf("%t", *s.UseEapolV1))
	}
	if s.VlanEnabled != nil {
		checks.append(t, "TestCheckResourceAttr", "vlan_enabled", fmt.Sprintf("%t", *s.VlanEnabled))
	}
	if s.VlanPooling != nil {
		checks.append(t, "TestCheckResourceAttr", "vlan_pooling", fmt.Sprintf("%t", *s.VlanPooling))
	}
	if s.WlanLimitDownEnabled != nil {
		checks.append(t, "TestCheckResourceAttr", "wlan_limit_down_enabled", fmt.Sprintf("%t", *s.WlanLimitDownEnabled))
	}
	if s.WlanLimitUpEnabled != nil {
		checks.append(t, "TestCheckResourceAttr", "wlan_limit_up_enabled", fmt.Sprintf("%t", *s.WlanLimitUpEnabled))
	}

	// Integer attributes
	if s.AcctInterimInterval != nil {
		checks.append(t, "TestCheckResourceAttr", "acct_interim_interval", fmt.Sprintf("%d", *s.AcctInterimInterval))
	}
	if s.AuthServersRetries != nil {
		checks.append(t, "TestCheckResourceAttr", "auth_servers_retries", fmt.Sprintf("%d", *s.AuthServersRetries))
	}
	if s.AuthServersTimeout != nil {
		checks.append(t, "TestCheckResourceAttr", "auth_servers_timeout", fmt.Sprintf("%d", *s.AuthServersTimeout))
	}
	if s.ClientLimitDown != nil {
		checks.append(t, "TestCheckResourceAttr", "client_limit_down", *s.ClientLimitDown)
	}
	if s.ClientLimitUp != nil {
		checks.append(t, "TestCheckResourceAttr", "client_limit_up", *s.ClientLimitUp)
	}
	if s.Dtim != nil {
		checks.append(t, "TestCheckResourceAttr", "dtim", fmt.Sprintf("%d", *s.Dtim))
	}
	if s.MaxIdletime != nil {
		checks.append(t, "TestCheckResourceAttr", "max_idletime", fmt.Sprintf("%d", *s.MaxIdletime))
	}
	if s.MaxNumClients != nil {
		checks.append(t, "TestCheckResourceAttr", "max_num_clients", fmt.Sprintf("%d", *s.MaxNumClients))
	}
	if s.WlanLimitDown != nil {
		checks.append(t, "TestCheckResourceAttr", "wlan_limit_down", *s.WlanLimitDown)
	}
	if s.WlanLimitUp != nil {
		checks.append(t, "TestCheckResourceAttr", "wlan_limit_up", *s.WlanLimitUp)
	}

	// String attributes
	if s.ApplyTo != nil {
		checks.append(t, "TestCheckResourceAttr", "apply_to", *s.ApplyTo)
	}
	if s.AuthServerSelection != nil {
		checks.append(t, "TestCheckResourceAttr", "auth_server_selection", *s.AuthServerSelection)
	}
	if s.AuthServersNasId != nil {
		checks.append(t, "TestCheckResourceAttr", "auth_servers_nas_id", *s.AuthServersNasId)
	}
	if s.AuthServersNasIp != nil {
		checks.append(t, "TestCheckResourceAttr", "auth_servers_nas_ip", *s.AuthServersNasIp)
	}
	if s.Interface != nil {
		checks.append(t, "TestCheckResourceAttr", "interface", *s.Interface)
	}
	if s.RoamMode != nil {
		checks.append(t, "TestCheckResourceAttr", "roam_mode", *s.RoamMode)
	}
	if s.VlanId != nil {
		checks.append(t, "TestCheckResourceAttr", "vlan_id", *s.VlanId)
	}
	if s.WxtunnelId != nil {
		checks.append(t, "TestCheckResourceAttr", "wxtunnel_id", *s.WxtunnelId)
	}
	if s.WxtunnelRemoteId != nil {
		checks.append(t, "TestCheckResourceAttr", "wxtunnel_remote_id", *s.WxtunnelRemoteId)
	}

	// List attributes
	if len(s.ApIds) > 0 {
		checks.append(t, "TestCheckResourceAttr", "ap_ids.#", fmt.Sprintf("%d", len(s.ApIds)))
		for i, apId := range s.ApIds {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ap_ids.%d", i), apId)
		}
	}
	if len(s.Bands) > 0 {
		checks.append(t, "TestCheckResourceAttr", "bands.#", fmt.Sprintf("%d", len(s.Bands)))
		for i, band := range s.Bands {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bands.%d", i), band)
		}
	}
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
	if len(s.VlanIds) > 0 {
		checks.append(t, "TestCheckResourceAttr", "vlan_ids.#", fmt.Sprintf("%d", len(s.VlanIds)))
		for i, vlanId := range s.VlanIds {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vlan_ids.%d", i), vlanId)
		}
	}
	if len(s.WxtagIds) > 0 {
		checks.append(t, "TestCheckResourceAttr", "wxtag_ids.#", fmt.Sprintf("%d", len(s.WxtagIds)))
		for i, wxtagId := range s.WxtagIds {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("wxtag_ids.%d", i), wxtagId)
		}
	}

	// Complex nested attributes - use element value key checks for map attributes
	if s.Airwatch != nil {
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
	if s.AppLimit != nil {
		if len(s.AppLimit.Apps) > 0 {
			checks.append(t, "TestCheckResourceAttr", "app_limit.apps.%", fmt.Sprintf("%d", len(s.AppLimit.Apps)))
			for appKey, appValue := range s.AppLimit.Apps {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("app_limit.apps.%s", appKey), fmt.Sprintf("%d", appValue))
			}
		}
		if s.AppLimit.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "app_limit.enabled", fmt.Sprintf("%t", *s.AppLimit.Enabled))
		}
		if len(s.AppLimit.WxtagIds) > 0 {
			checks.append(t, "TestCheckResourceAttr", "app_limit.wxtag_ids.%", fmt.Sprintf("%d", len(s.AppLimit.WxtagIds)))
			for wxtagKey, wxtagValue := range s.AppLimit.WxtagIds {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("app_limit.wxtag_ids.%s", wxtagKey), fmt.Sprintf("%d", wxtagValue))
			}
		}
	}
	if s.AppQos != nil {
		if len(s.AppQos.Apps) > 0 {
			checks.append(t, "TestCheckResourceAttr", "app_qos.apps.%", fmt.Sprintf("%d", len(s.AppQos.Apps)))
			for appKey, app := range s.AppQos.Apps {
				prefix := fmt.Sprintf("app_qos.apps.%s", appKey)
				if app.Dscp != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".dscp", *app.Dscp)
				}
				if app.DstSubnet != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".dst_subnet", *app.DstSubnet)
				}
				if app.SrcSubnet != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".src_subnet", *app.SrcSubnet)
				}
			}
		}
		if s.AppQos.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "app_qos.enabled", fmt.Sprintf("%t", *s.AppQos.Enabled))
		}
		if len(s.AppQos.Others) > 0 {
			checks.append(t, "TestCheckResourceAttr", "app_qos.others.#", fmt.Sprintf("%d", len(s.AppQos.Others)))
			for i, other := range s.AppQos.Others {
				prefix := fmt.Sprintf("app_qos.others.%d", i)
				if other.Dscp != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".dscp", *other.Dscp)
				}
				if other.DstSubnet != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".dst_subnet", *other.DstSubnet)
				}
				if other.PortRanges != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".port_ranges", *other.PortRanges)
				}
				if other.Protocol != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".protocol", *other.Protocol)
				}
				if other.SrcSubnet != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".src_subnet", *other.SrcSubnet)
				}
			}
		}
	}
	if s.Auth != nil {
		if s.Auth.AnticlogThreshold != nil {
			checks.append(t, "TestCheckResourceAttr", "auth.anticlog_threshold", fmt.Sprintf("%d", *s.Auth.AnticlogThreshold))
		}
		if s.Auth.EapReauth != nil {
			checks.append(t, "TestCheckResourceAttr", "auth.eap_reauth", fmt.Sprintf("%t", *s.Auth.EapReauth))
		}
		if s.Auth.EnableMacAuth != nil {
			checks.append(t, "TestCheckResourceAttr", "auth.enable_mac_auth", fmt.Sprintf("%t", *s.Auth.EnableMacAuth))
		}
		if s.Auth.KeyIdx != nil {
			checks.append(t, "TestCheckResourceAttr", "auth.key_idx", fmt.Sprintf("%d", *s.Auth.KeyIdx))
		}
		if len(s.Auth.Keys) > 0 {
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
		if len(s.Auth.Pairwise) > 0 {
			checks.append(t, "TestCheckResourceAttr", "auth.pairwise.#", fmt.Sprintf("%d", len(s.Auth.Pairwise)))
			for i, pairwise := range s.Auth.Pairwise {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("auth.pairwise.%d", i), pairwise)
			}
		}
		if s.Auth.PrivateWlan != nil {
			checks.append(t, "TestCheckResourceAttr", "auth.private_wlan", fmt.Sprintf("%t", *s.Auth.PrivateWlan))
		}
		if s.Auth.Psk != nil {
			checks.append(t, "TestCheckResourceAttr", "auth.psk", *s.Auth.Psk)
		}
		if s.Auth.AuthType != nil {
			checks.append(t, "TestCheckResourceAttr", "auth.type", *s.Auth.AuthType)
		}
		if s.Auth.WepAsSecondaryAuth != nil {
			checks.append(t, "TestCheckResourceAttr", "auth.wep_as_secondary_auth", fmt.Sprintf("%t", *s.Auth.WepAsSecondaryAuth))
		}
	}
	if s.Bonjour != nil {
		if len(s.Bonjour.AdditionalVlanIds) > 0 {
			checks.append(t, "TestCheckResourceAttr", "bonjour.additional_vlan_ids.#", fmt.Sprintf("%d", len(s.Bonjour.AdditionalVlanIds)))
			for i, vlanId := range s.Bonjour.AdditionalVlanIds {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bonjour.additional_vlan_ids.%d", i), vlanId)
			}
		}
		if s.Bonjour.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "bonjour.enabled", fmt.Sprintf("%t", *s.Bonjour.Enabled))
		}
		if len(s.Bonjour.Services) > 0 {
			checks.append(t, "TestCheckResourceAttr", "bonjour.services.%", fmt.Sprintf("%d", len(s.Bonjour.Services)))
			for serviceKey, service := range s.Bonjour.Services {
				prefix := fmt.Sprintf("bonjour.services.%s", serviceKey)
				if service.DisableLocal != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".disable_local", fmt.Sprintf("%t", *service.DisableLocal))
				}
				if len(service.RadiusGroups) > 0 {
					checks.append(t, "TestCheckResourceAttr", prefix+".radius_groups.#", fmt.Sprintf("%d", len(service.RadiusGroups)))
					for i, group := range service.RadiusGroups {
						checks.append(t, "TestCheckResourceAttr", prefix+fmt.Sprintf(".radius_groups.%d", i), group)
					}
				}
				if service.Scope != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".scope", *service.Scope)
				}
			}
		}
	}
	if s.CiscoCwa != nil {
		if len(s.CiscoCwa.AllowedHostnames) > 0 {
			checks.append(t, "TestCheckResourceAttr", "cisco_cwa.allowed_hostnames.#", fmt.Sprintf("%d", len(s.CiscoCwa.AllowedHostnames)))
			for i, hostname := range s.CiscoCwa.AllowedHostnames {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("cisco_cwa.allowed_hostnames.%d", i), hostname)
			}
		}
		if len(s.CiscoCwa.AllowedSubnets) > 0 {
			checks.append(t, "TestCheckResourceAttr", "cisco_cwa.allowed_subnets.#", fmt.Sprintf("%d", len(s.CiscoCwa.AllowedSubnets)))
			for i, subnet := range s.CiscoCwa.AllowedSubnets {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("cisco_cwa.allowed_subnets.%d", i), subnet)
			}
		}
		if len(s.CiscoCwa.BlockedSubnets) > 0 {
			checks.append(t, "TestCheckResourceAttr", "cisco_cwa.blocked_subnets.#", fmt.Sprintf("%d", len(s.CiscoCwa.BlockedSubnets)))
			for i, subnet := range s.CiscoCwa.BlockedSubnets {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("cisco_cwa.blocked_subnets.%d", i), subnet)
			}
		}
		if s.CiscoCwa.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "cisco_cwa.enabled", fmt.Sprintf("%t", *s.CiscoCwa.Enabled))
		}
	}
	if s.DnsServerRewrite != nil {
		if s.DnsServerRewrite.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "dns_server_rewrite.enabled", fmt.Sprintf("%t", *s.DnsServerRewrite.Enabled))
		}
		if len(s.DnsServerRewrite.RadiusGroups) > 0 {
			checks.append(t, "TestCheckResourceAttr", "dns_server_rewrite.radius_groups.%", fmt.Sprintf("%d", len(s.DnsServerRewrite.RadiusGroups)))
			for groupKey, groupValue := range s.DnsServerRewrite.RadiusGroups {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dns_server_rewrite.radius_groups.%s", groupKey), groupValue)
			}
		}
	}
	if s.DynamicPsk != nil {
		if s.DynamicPsk.DefaultPsk != nil {
			checks.append(t, "TestCheckResourceAttr", "dynamic_psk.default_psk", *s.DynamicPsk.DefaultPsk)
		}
		if s.DynamicPsk.DefaultVlanId != nil {
			checks.append(t, "TestCheckResourceAttr", "dynamic_psk.default_vlan_id", *s.DynamicPsk.DefaultVlanId)
		}
		if s.DynamicPsk.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "dynamic_psk.enabled", fmt.Sprintf("%t", *s.DynamicPsk.Enabled))
		}
		if s.DynamicPsk.ForceLookup != nil {
			checks.append(t, "TestCheckResourceAttr", "dynamic_psk.force_lookup", fmt.Sprintf("%t", *s.DynamicPsk.ForceLookup))
		}
		if s.DynamicPsk.Source != nil {
			checks.append(t, "TestCheckResourceAttr", "dynamic_psk.source", *s.DynamicPsk.Source)
		}
	}
	if s.DynamicVlan != nil {
		if len(s.DynamicVlan.DefaultVlanIds) > 0 {
			checks.append(t, "TestCheckResourceAttr", "dynamic_vlan.default_vlan_ids.#", fmt.Sprintf("%d", len(s.DynamicVlan.DefaultVlanIds)))
			for i, vlanId := range s.DynamicVlan.DefaultVlanIds {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dynamic_vlan.default_vlan_ids.%d", i), vlanId)
			}
		}
		if s.DynamicVlan.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "dynamic_vlan.enabled", fmt.Sprintf("%t", *s.DynamicVlan.Enabled))
		}
		if len(s.DynamicVlan.LocalVlanIds) > 0 {
			checks.append(t, "TestCheckResourceAttr", "dynamic_vlan.local_vlan_ids.#", fmt.Sprintf("%d", len(s.DynamicVlan.LocalVlanIds)))
			for i, vlanId := range s.DynamicVlan.LocalVlanIds {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dynamic_vlan.local_vlan_ids.%d", i), vlanId)
			}
		}
		if s.DynamicVlan.DynamicVlanType != nil {
			checks.append(t, "TestCheckResourceAttr", "dynamic_vlan.type", *s.DynamicVlan.DynamicVlanType)
		}
		if len(s.DynamicVlan.Vlans) > 0 {
			checks.append(t, "TestCheckResourceAttr", "dynamic_vlan.vlans.%", fmt.Sprintf("%d", len(s.DynamicVlan.Vlans)))
			for vlanKey, vlanValue := range s.DynamicVlan.Vlans {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dynamic_vlan.vlans.%s", vlanKey), vlanValue)
			}
		}
	}
	if s.Hotspot20 != nil {
		if len(s.Hotspot20.DomainName) > 0 {
			checks.append(t, "TestCheckResourceAttr", "hotspot20.domain_name.#", fmt.Sprintf("%d", len(s.Hotspot20.DomainName)))
			for i, domain := range s.Hotspot20.DomainName {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("hotspot20.domain_name.%d", i), domain)
			}
		}
		if s.Hotspot20.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "hotspot20.enabled", fmt.Sprintf("%t", *s.Hotspot20.Enabled))
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
		if s.InjectDhcpOption82.CircuitId != nil {
			checks.append(t, "TestCheckResourceAttr", "inject_dhcp_option_82.circuit_id", *s.InjectDhcpOption82.CircuitId)
		}
		if s.InjectDhcpOption82.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "inject_dhcp_option_82.enabled", fmt.Sprintf("%t", *s.InjectDhcpOption82.Enabled))
		}
	}
	if s.MistNac != nil {
		if s.MistNac.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "mist_nac.enabled", fmt.Sprintf("%t", *s.MistNac.Enabled))
		}
		if s.MistNac.Network != nil {
			checks.append(t, "TestCheckResourceAttr", "mist_nac.network", *s.MistNac.Network)
		}
	}
	if s.Portal != nil {
		if s.Portal.AllowWlanIdRoam != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.allow_wlan_id_roam", fmt.Sprintf("%t", *s.Portal.AllowWlanIdRoam))
		}
		if s.Portal.AmazonClientId != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.amazon_client_id", *s.Portal.AmazonClientId)
		}
		if s.Portal.AmazonClientSecret != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.amazon_client_secret", *s.Portal.AmazonClientSecret)
		}
		if len(s.Portal.AmazonEmailDomains) > 0 {
			checks.append(t, "TestCheckResourceAttr", "portal.amazon_email_domains.#", fmt.Sprintf("%d", len(s.Portal.AmazonEmailDomains)))
			for i, domain := range s.Portal.AmazonEmailDomains {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal.amazon_email_domains.%d", i), domain)
			}
		}
		if s.Portal.AmazonEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.amazon_enabled", fmt.Sprintf("%t", *s.Portal.AmazonEnabled))
		}
		if s.Portal.AmazonExpire != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.amazon_expire", fmt.Sprintf("%d", *s.Portal.AmazonExpire))
		}
		if s.Portal.Auth != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.auth", *s.Portal.Auth)
		}
		if s.Portal.AzureClientId != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.azure_client_id", *s.Portal.AzureClientId)
		}
		if s.Portal.AzureClientSecret != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.azure_client_secret", *s.Portal.AzureClientSecret)
		}
		if s.Portal.AzureEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.azure_enabled", fmt.Sprintf("%t", *s.Portal.AzureEnabled))
		}
		if s.Portal.AzureExpire != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.azure_expire", fmt.Sprintf("%d", *s.Portal.AzureExpire))
		}
		if s.Portal.AzureTenantId != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.azure_tenant_id", *s.Portal.AzureTenantId)
		}
		if s.Portal.EmailEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.email_enabled", fmt.Sprintf("%t", *s.Portal.EmailEnabled))
		}
		if s.Portal.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.enabled", fmt.Sprintf("%t", *s.Portal.Enabled))
		}
		if s.Portal.Expire != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.expire", fmt.Sprintf("%d", *s.Portal.Expire))
		}
		if s.Portal.ExternalPortalUrl != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.external_portal_url", *s.Portal.ExternalPortalUrl)
		}
		if s.Portal.FacebookClientId != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.facebook_client_id", *s.Portal.FacebookClientId)
		}
		if s.Portal.FacebookClientSecret != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.facebook_client_secret", *s.Portal.FacebookClientSecret)
		}
		if len(s.Portal.FacebookEmailDomains) > 0 {
			checks.append(t, "TestCheckResourceAttr", "portal.facebook_email_domains.#", fmt.Sprintf("%d", len(s.Portal.FacebookEmailDomains)))
			for i, domain := range s.Portal.FacebookEmailDomains {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal.facebook_email_domains.%d", i), domain)
			}
		}
		if s.Portal.FacebookEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.facebook_enabled", fmt.Sprintf("%t", *s.Portal.FacebookEnabled))
		}
		if s.Portal.FacebookExpire != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.facebook_expire", fmt.Sprintf("%d", *s.Portal.FacebookExpire))
		}
		if s.Portal.Forward != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.forward", fmt.Sprintf("%t", *s.Portal.Forward))
		}
		if s.Portal.ForwardUrl != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.forward_url", *s.Portal.ForwardUrl)
		}
		if s.Portal.GoogleClientId != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.google_client_id", *s.Portal.GoogleClientId)
		}
		if s.Portal.GoogleClientSecret != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.google_client_secret", *s.Portal.GoogleClientSecret)
		}
		if len(s.Portal.GoogleEmailDomains) > 0 {
			checks.append(t, "TestCheckResourceAttr", "portal.google_email_domains.#", fmt.Sprintf("%d", len(s.Portal.GoogleEmailDomains)))
			for i, domain := range s.Portal.GoogleEmailDomains {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal.google_email_domains.%d", i), domain)
			}
		}
		if s.Portal.GoogleEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.google_enabled", fmt.Sprintf("%t", *s.Portal.GoogleEnabled))
		}
		if s.Portal.GoogleExpire != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.google_expire", fmt.Sprintf("%d", *s.Portal.GoogleExpire))
		}
		if s.Portal.MicrosoftClientId != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.microsoft_client_id", *s.Portal.MicrosoftClientId)
		}
		if s.Portal.MicrosoftClientSecret != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.microsoft_client_secret", *s.Portal.MicrosoftClientSecret)
		}
		if len(s.Portal.MicrosoftEmailDomains) > 0 {
			checks.append(t, "TestCheckResourceAttr", "portal.microsoft_email_domains.#", fmt.Sprintf("%d", len(s.Portal.MicrosoftEmailDomains)))
			for i, domain := range s.Portal.MicrosoftEmailDomains {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal.microsoft_email_domains.%d", i), domain)
			}
		}
		if s.Portal.MicrosoftEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.microsoft_enabled", fmt.Sprintf("%t", *s.Portal.MicrosoftEnabled))
		}
		if s.Portal.MicrosoftExpire != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.microsoft_expire", fmt.Sprintf("%d", *s.Portal.MicrosoftExpire))
		}
		if s.Portal.PassphraseEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.passphrase_enabled", fmt.Sprintf("%t", *s.Portal.PassphraseEnabled))
		}
		if s.Portal.PassphraseExpire != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.passphrase_expire", fmt.Sprintf("%d", *s.Portal.PassphraseExpire))
		}
		if s.Portal.Password != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.password", *s.Portal.Password)
		}
		if s.Portal.Privacy != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.privacy", fmt.Sprintf("%t", *s.Portal.Privacy))
		}
		if s.Portal.SmsEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.sms_enabled", fmt.Sprintf("%t", *s.Portal.SmsEnabled))
		}
		if s.Portal.SmsExpire != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.sms_expire", fmt.Sprintf("%d", *s.Portal.SmsExpire))
		}
		if s.Portal.SmsMessageFormat != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.sms_message_format", *s.Portal.SmsMessageFormat)
		}
		if s.Portal.SmsProvider != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.sms_provider", *s.Portal.SmsProvider)
		}
		if s.Portal.SponsorEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.sponsor_enabled", fmt.Sprintf("%t", *s.Portal.SponsorEnabled))
		}
		if s.Portal.SponsorExpire != nil {
			checks.append(t, "TestCheckResourceAttr", "portal.sponsor_expire", fmt.Sprintf("%d", *s.Portal.SponsorExpire))
		}
		if len(s.Portal.Sponsors) > 0 {
			checks.append(t, "TestCheckResourceAttr", "portal.sponsors.%", fmt.Sprintf("%d", len(s.Portal.Sponsors)))
			for sponsorKey, sponsorValue := range s.Portal.Sponsors {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal.sponsors.%s", sponsorKey), sponsorValue)
			}
		}
	}
	if s.Qos != nil {
		if s.Qos.Class != nil {
			checks.append(t, "TestCheckResourceAttr", "qos.class", *s.Qos.Class)
		}
		if s.Qos.Overwrite != nil {
			checks.append(t, "TestCheckResourceAttr", "qos.overwrite", fmt.Sprintf("%t", *s.Qos.Overwrite))
		}
	}
	if s.Radsec != nil {
		if s.Radsec.CoaEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "radsec.coa_enabled", fmt.Sprintf("%t", *s.Radsec.CoaEnabled))
		}
		if s.Radsec.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "radsec.enabled", fmt.Sprintf("%t", *s.Radsec.Enabled))
		}
		if s.Radsec.IdleTimeout != nil {
			checks.append(t, "TestCheckResourceAttr", "radsec.idle_timeout", *s.Radsec.IdleTimeout)
		}
		if len(s.Radsec.MxclusterIds) > 0 {
			checks.append(t, "TestCheckResourceAttr", "radsec.mxcluster_ids.#", fmt.Sprintf("%d", len(s.Radsec.MxclusterIds)))
			for i, clusterId := range s.Radsec.MxclusterIds {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radsec.mxcluster_ids.%d", i), clusterId)
			}
		}
		if len(s.Radsec.ProxyHosts) > 0 {
			checks.append(t, "TestCheckResourceAttr", "radsec.proxy_hosts.#", fmt.Sprintf("%d", len(s.Radsec.ProxyHosts)))
			for i, host := range s.Radsec.ProxyHosts {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radsec.proxy_hosts.%d", i), host)
			}
		}
		if s.Radsec.ServerName != nil {
			checks.append(t, "TestCheckResourceAttr", "radsec.server_name", *s.Radsec.ServerName)
		}
		if len(s.Radsec.Servers) > 0 {
			checks.append(t, "TestCheckResourceAttr", "radsec.servers.#", fmt.Sprintf("%d", len(s.Radsec.Servers)))
			for i, server := range s.Radsec.Servers {
				prefix := fmt.Sprintf("radsec.servers.%d", i)
				if server.Host != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".host", *server.Host)
				}
				if server.Port != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".port", fmt.Sprintf("%d", *server.Port))
				}
			}
		}
		if s.Radsec.UseMxedge != nil {
			checks.append(t, "TestCheckResourceAttr", "radsec.use_mxedge", fmt.Sprintf("%t", *s.Radsec.UseMxedge))
		}
		if s.Radsec.UseSiteMxedge != nil {
			checks.append(t, "TestCheckResourceAttr", "radsec.use_site_mxedge", fmt.Sprintf("%t", *s.Radsec.UseSiteMxedge))
		}
	}
	if s.Schedule != nil {
		if s.Schedule.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "schedule.enabled", fmt.Sprintf("%t", *s.Schedule.Enabled))
		}
		if s.Schedule.Hours != nil {
			if s.Schedule.Hours.Fri != nil {
				checks.append(t, "TestCheckResourceAttr", "schedule.hours.fri", *s.Schedule.Hours.Fri)
			}
			if s.Schedule.Hours.Mon != nil {
				checks.append(t, "TestCheckResourceAttr", "schedule.hours.mon", *s.Schedule.Hours.Mon)
			}
			if s.Schedule.Hours.Sat != nil {
				checks.append(t, "TestCheckResourceAttr", "schedule.hours.sat", *s.Schedule.Hours.Sat)
			}
			if s.Schedule.Hours.Sun != nil {
				checks.append(t, "TestCheckResourceAttr", "schedule.hours.sun", *s.Schedule.Hours.Sun)
			}
			if s.Schedule.Hours.Thu != nil {
				checks.append(t, "TestCheckResourceAttr", "schedule.hours.thu", *s.Schedule.Hours.Thu)
			}
			if s.Schedule.Hours.Tue != nil {
				checks.append(t, "TestCheckResourceAttr", "schedule.hours.tue", *s.Schedule.Hours.Tue)
			}
			if s.Schedule.Hours.Wed != nil {
				checks.append(t, "TestCheckResourceAttr", "schedule.hours.wed", *s.Schedule.Hours.Wed)
			}
		}
	}

	// List of complex objects
	if len(s.AcctServers) > 0 {
		checks.append(t, "TestCheckResourceAttr", "acct_servers.#", fmt.Sprintf("%d", len(s.AcctServers)))
		for i, server := range s.AcctServers {
			prefix := fmt.Sprintf("acct_servers.%d", i)
			checks.append(t, "TestCheckResourceAttr", prefix+".host", server.Host)
			if server.KeywrapEnabled != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".keywrap_enabled", fmt.Sprintf("%t", *server.KeywrapEnabled))
			}
			if server.KeywrapFormat != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".keywrap_format", *server.KeywrapFormat)
			}
			if server.KeywrapKek != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".keywrap_kek", *server.KeywrapKek)
			}
			if server.KeywrapMack != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".keywrap_mack", *server.KeywrapMack)
			}
			if server.Port != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".port", *server.Port)
			}
			checks.append(t, "TestCheckResourceAttr", prefix+".secret", server.Secret)
		}
	}
	if len(s.AuthServers) > 0 {
		checks.append(t, "TestCheckResourceAttr", "auth_servers.#", fmt.Sprintf("%d", len(s.AuthServers)))
		for i, server := range s.AuthServers {
			prefix := fmt.Sprintf("auth_servers.%d", i)
			checks.append(t, "TestCheckResourceAttr", prefix+".host", server.Host)
			if server.KeywrapEnabled != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".keywrap_enabled", fmt.Sprintf("%t", *server.KeywrapEnabled))
			}
			if server.KeywrapFormat != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".keywrap_format", *server.KeywrapFormat)
			}
			if server.KeywrapKek != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".keywrap_kek", *server.KeywrapKek)
			}
			if server.KeywrapMack != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".keywrap_mack", *server.KeywrapMack)
			}
			if server.Port != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".port", *server.Port)
			}
			if server.RequireMessageAuthenticator != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".require_message_authenticator", fmt.Sprintf("%t", *server.RequireMessageAuthenticator))
			}
			checks.append(t, "TestCheckResourceAttr", prefix+".secret", server.Secret)
		}
	}
	if len(s.CoaServers) > 0 {
		checks.append(t, "TestCheckResourceAttr", "coa_servers.#", fmt.Sprintf("%d", len(s.CoaServers)))
		for i, server := range s.CoaServers {
			prefix := fmt.Sprintf("coa_servers.%d", i)
			if server.DisableEventTimestampCheck != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".disable_event_timestamp_check", fmt.Sprintf("%t", *server.DisableEventTimestampCheck))
			}
			if server.Enabled != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".enabled", fmt.Sprintf("%t", *server.Enabled))
			}
			checks.append(t, "TestCheckResourceAttr", prefix+".ip", server.Ip)
			if server.Port != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".port", *server.Port)
			}
			checks.append(t, "TestCheckResourceAttr", prefix+".secret", server.Secret)
		}
	}

	// Map attributes
	if len(s.Rateset) > 0 {
		checks.append(t, "TestCheckResourceAttr", "rateset.%", fmt.Sprintf("%d", len(s.Rateset)))
		for key, rateset := range s.Rateset {
			prefix := fmt.Sprintf("rateset.%s", key)
			if rateset.Ht != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".ht", *rateset.Ht)
			}
			if len(rateset.Legacy) > 0 {
				checks.append(t, "TestCheckResourceAttr", prefix+".legacy.#", fmt.Sprintf("%d", len(rateset.Legacy)))
				for i, legacy := range rateset.Legacy {
					checks.append(t, "TestCheckResourceAttr", prefix+fmt.Sprintf(".legacy.%d", i), legacy)
				}
			}
			if rateset.MinRssi != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".min_rssi", fmt.Sprintf("%d", *rateset.MinRssi))
			}
			if rateset.Template != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".template", *rateset.Template)
			}
			if rateset.Vht != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".vht", *rateset.Vht)
			}
		}
	}

	return checks
}
