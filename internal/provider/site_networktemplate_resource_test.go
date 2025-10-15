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

func TestSiteNetworktemplateModel(t *testing.T) {
	type testStep struct {
		config SiteNetworktemplateModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: SiteNetworktemplateModel{},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/site_networktemplate_resource/site_networktemplate_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		var FixtureSiteNetworktemplateModel SiteNetworktemplateModel

		err = hcl.Decode(&FixtureSiteNetworktemplateModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureSiteNetworktemplateModel,
				},
			},
		}
	}

	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			resourceType := "site_networktemplate"

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

func (s *SiteNetworktemplateModel) testChecks(t testing.TB, rType, rName string) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + rName)
	checks.append(t, "TestCheckResourceAttrSet", "site_id")

	// Check acl_policies list if present
	if len(s.AclPolicies) > 0 {
		checks.append(t, "TestCheckResourceAttr", "acl_policies.#", fmt.Sprintf("%d", len(s.AclPolicies)))
		for i, policy := range s.AclPolicies {
			prefix := fmt.Sprintf("acl_policies.%d", i)
			if policy.Name != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".name", *policy.Name)
			}
			if len(policy.SrcTags) > 0 {
				checks.append(t, "TestCheckResourceAttr", prefix+".src_tags.#", fmt.Sprintf("%d", len(policy.SrcTags)))
				for j, tag := range policy.SrcTags {
					checks.append(t, "TestCheckResourceAttr", prefix+fmt.Sprintf(".src_tags.%d", j), tag)
				}
			}
			if len(policy.Actions) > 0 {
				checks.append(t, "TestCheckResourceAttr", prefix+".actions.#", fmt.Sprintf("%d", len(policy.Actions)))
				for j, action := range policy.Actions {
					actionPrefix := prefix + fmt.Sprintf(".actions.%d", j)
					if action.Action != nil {
						checks.append(t, "TestCheckResourceAttr", actionPrefix+".action", *action.Action)
					}
					checks.append(t, "TestCheckResourceAttr", actionPrefix+".dst_tag", action.DstTag)
				}
			}
		}
	}

	// Check acl_tags map if present
	if len(s.AclTags) > 0 {
		checks.append(t, "TestCheckResourceAttr", "acl_tags.%", fmt.Sprintf("%d", len(s.AclTags)))
		for key, tag := range s.AclTags {
			prefix := fmt.Sprintf("acl_tags.%s", key)
			if len(tag.EtherTypes) > 0 {
				checks.append(t, "TestCheckResourceAttr", prefix+".ether_types.#", fmt.Sprintf("%d", len(tag.EtherTypes)))
				for i, etherType := range tag.EtherTypes {
					checks.append(t, "TestCheckResourceAttr", prefix+fmt.Sprintf(".ether_types.%d", i), etherType)
				}
			}
			if tag.GbpTag != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".gbp_tag", fmt.Sprintf("%d", *tag.GbpTag))
			}
			if len(tag.Macs) > 0 {
				checks.append(t, "TestCheckResourceAttr", prefix+".macs.#", fmt.Sprintf("%d", len(tag.Macs)))
				for i, mac := range tag.Macs {
					checks.append(t, "TestCheckResourceAttr", prefix+fmt.Sprintf(".macs.%d", i), mac)
				}
			}
			if tag.Network != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".network", *tag.Network)
			}
			if tag.PortUsage != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".port_usage", *tag.PortUsage)
			}
			if tag.RadiusGroup != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".radius_group", *tag.RadiusGroup)
			}
			if len(tag.Specs) > 0 {
				checks.append(t, "TestCheckResourceAttr", prefix+".specs.#", fmt.Sprintf("%d", len(tag.Specs)))
				for i, spec := range tag.Specs {
					specPrefix := prefix + fmt.Sprintf(".specs.%d", i)
					if spec.PortRange != nil {
						checks.append(t, "TestCheckResourceAttr", specPrefix+".port_range", *spec.PortRange)
					}
					if spec.Protocol != nil {
						checks.append(t, "TestCheckResourceAttr", specPrefix+".protocol", *spec.Protocol)
					}
				}
			}
			if len(tag.Subnets) > 0 {
				checks.append(t, "TestCheckResourceAttr", prefix+".subnets.#", fmt.Sprintf("%d", len(tag.Subnets)))
				for i, subnet := range tag.Subnets {
					checks.append(t, "TestCheckResourceAttr", prefix+fmt.Sprintf(".subnets.%d", i), subnet)
				}
			}
			checks.append(t, "TestCheckResourceAttr", prefix+".type", tag.AclTagsType)
		}
	}

	// Check additional_config_cmds list if present
	if len(s.AdditionalConfigCmds) > 0 {
		checks.append(t, "TestCheckResourceAttr", "additional_config_cmds.#", fmt.Sprintf("%d", len(s.AdditionalConfigCmds)))
		for i, cmd := range s.AdditionalConfigCmds {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("additional_config_cmds.%d", i), cmd)
		}
	}
	if s.AutoUpgradeLinecard != nil {
		checks.append(t, "TestCheckResourceAttr", "auto_upgrade_linecard", fmt.Sprintf("%t", *s.AutoUpgradeLinecard))
	}
	if s.DefaultPortUsage != nil {
		checks.append(t, "TestCheckResourceAttr", "default_port_usage", *s.DefaultPortUsage)
	}
	// Check dhcp_snooping if present
	if s.DhcpSnooping != nil {
		if s.DhcpSnooping.AllNetworks != nil {
			checks.append(t, "TestCheckResourceAttr", "dhcp_snooping.all_networks", fmt.Sprintf("%t", *s.DhcpSnooping.AllNetworks))
		}
		if s.DhcpSnooping.EnableArpSpoofCheck != nil {
			checks.append(t, "TestCheckResourceAttr", "dhcp_snooping.enable_arp_spoof_check", fmt.Sprintf("%t", *s.DhcpSnooping.EnableArpSpoofCheck))
		}
		if s.DhcpSnooping.EnableIpSourceGuard != nil {
			checks.append(t, "TestCheckResourceAttr", "dhcp_snooping.enable_ip_source_guard", fmt.Sprintf("%t", *s.DhcpSnooping.EnableIpSourceGuard))
		}
		if s.DhcpSnooping.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "dhcp_snooping.enabled", fmt.Sprintf("%t", *s.DhcpSnooping.Enabled))
		}
		if len(s.DhcpSnooping.Networks) > 0 {
			checks.append(t, "TestCheckResourceAttr", "dhcp_snooping.networks.#", fmt.Sprintf("%d", len(s.DhcpSnooping.Networks)))
		}
	}

	// Check disabled_system_defined_port_usages list if present
	if len(s.DisabledSystemDefinedPortUsages) > 0 {
		checks.append(t, "TestCheckResourceAttr", "disabled_system_defined_port_usages.#", fmt.Sprintf("%d", len(s.DisabledSystemDefinedPortUsages)))
		for i, usage := range s.DisabledSystemDefinedPortUsages {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("disabled_system_defined_port_usages.%d", i), usage)
		}
	}

	// Check dns_servers list if present
	if len(s.DnsServers) > 0 {
		checks.append(t, "TestCheckResourceAttr", "dns_servers.#", fmt.Sprintf("%d", len(s.DnsServers)))
		for i, server := range s.DnsServers {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dns_servers.%d", i), server)
		}
	}

	// Check dns_suffix list if present
	if len(s.DnsSuffix) > 0 {
		checks.append(t, "TestCheckResourceAttr", "dns_suffix.#", fmt.Sprintf("%d", len(s.DnsSuffix)))
		for i, suffix := range s.DnsSuffix {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dns_suffix.%d", i), suffix)
		}
	}

	// Check extra_routes map if present
	if len(s.ExtraRoutes) > 0 {
		checks.append(t, "TestCheckResourceAttr", "extra_routes.%", fmt.Sprintf("%d", len(s.ExtraRoutes)))
		for key, route := range s.ExtraRoutes {
			prefix := fmt.Sprintf("extra_routes.%s", key)
			if route.Discard != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".discard", fmt.Sprintf("%t", *route.Discard))
			}
			if route.Metric != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".metric", fmt.Sprintf("%d", *route.Metric))
			}
			if len(route.NextQualified) > 0 {
				checks.append(t, "TestCheckResourceAttr", prefix+".next_qualified.%", fmt.Sprintf("%d", len(route.NextQualified)))
				for qualKey, qual := range route.NextQualified {
					qualPrefix := prefix + fmt.Sprintf(".next_qualified.%s", qualKey)
					if qual.Metric != nil {
						checks.append(t, "TestCheckResourceAttr", qualPrefix+".metric", fmt.Sprintf("%d", *qual.Metric))
					}
					if qual.Preference != nil {
						checks.append(t, "TestCheckResourceAttr", qualPrefix+".preference", fmt.Sprintf("%d", *qual.Preference))
					}
				}
			}
			if route.NoResolve != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".no_resolve", fmt.Sprintf("%t", *route.NoResolve))
			}
			if route.Preference != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".preference", fmt.Sprintf("%d", *route.Preference))
			}
			checks.append(t, "TestCheckResourceAttr", prefix+".via", route.Via)
		}
	}

	// Check extra_routes6 map if present
	if len(s.ExtraRoutes6) > 0 {
		checks.append(t, "TestCheckResourceAttr", "extra_routes6.%", fmt.Sprintf("%d", len(s.ExtraRoutes6)))
		for key, route := range s.ExtraRoutes6 {
			prefix := fmt.Sprintf("extra_routes6.%s", key)
			if route.Discard != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".discard", fmt.Sprintf("%t", *route.Discard))
			}
			if route.Metric != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".metric", fmt.Sprintf("%d", *route.Metric))
			}
			if len(route.NextQualified) > 0 {
				checks.append(t, "TestCheckResourceAttr", prefix+".next_qualified.%", fmt.Sprintf("%d", len(route.NextQualified)))
				for qualKey, qual := range route.NextQualified {
					qualPrefix := prefix + fmt.Sprintf(".next_qualified.%s", qualKey)
					if qual.Metric != nil {
						checks.append(t, "TestCheckResourceAttr", qualPrefix+".metric", fmt.Sprintf("%d", *qual.Metric))
					}
					if qual.Preference != nil {
						checks.append(t, "TestCheckResourceAttr", qualPrefix+".preference", fmt.Sprintf("%d", *qual.Preference))
					}
				}
			}
			if route.NoResolve != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".no_resolve", fmt.Sprintf("%t", *route.NoResolve))
			}
			if route.Preference != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".preference", fmt.Sprintf("%d", *route.Preference))
			}
			checks.append(t, "TestCheckResourceAttr", prefix+".via", route.Via)
		}
	}

	// Check mist_nac if present
	if s.MistNac != nil {
		if s.MistNac.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "mist_nac.enabled", fmt.Sprintf("%t", *s.MistNac.Enabled))
		}
		if s.MistNac.Network != nil {
			checks.append(t, "TestCheckResourceAttr", "mist_nac.network", *s.MistNac.Network)
		}
	}

	// Check networks map if present
	if len(s.Networks) > 0 {
		checks.append(t, "TestCheckResourceAttr", "networks.%", fmt.Sprintf("%d", len(s.Networks)))
		for key, network := range s.Networks {
			prefix := fmt.Sprintf("networks.%s", key)
			if network.Gateway != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".gateway", *network.Gateway)
			}
			if network.Gateway6 != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".gateway6", *network.Gateway6)
			}
			if network.Isolation != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".isolation", fmt.Sprintf("%t", *network.Isolation))
			}
			if network.IsolationVlanId != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".isolation_vlan_id", *network.IsolationVlanId)
			}
			if network.Subnet != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".subnet", *network.Subnet)
			}
			if network.Subnet6 != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".subnet6", *network.Subnet6)
			}
			checks.append(t, "TestCheckResourceAttr", prefix+".vlan_id", network.VlanId)
		}
	}

	// Check ntp_servers list if present
	if len(s.NtpServers) > 0 {
		checks.append(t, "TestCheckResourceAttr", "ntp_servers.#", fmt.Sprintf("%d", len(s.NtpServers)))
		for i, server := range s.NtpServers {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ntp_servers.%d", i), server)
		}
	}

	// Check ospf_areas map if present
	if len(s.OspfAreas) > 0 {
		checks.append(t, "TestCheckResourceAttr", "ospf_areas.%", fmt.Sprintf("%d", len(s.OspfAreas)))
		for key, area := range s.OspfAreas {
			prefix := fmt.Sprintf("ospf_areas.%s", key)
			if area.IncludeLoopback != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".include_loopback", fmt.Sprintf("%t", *area.IncludeLoopback))
			}
			if area.OspfAreasType != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".type", *area.OspfAreasType)
			}
			if len(area.OspfNetworks) > 0 {
				checks.append(t, "TestCheckResourceAttr", prefix+".networks.%", fmt.Sprintf("%d", len(area.OspfNetworks)))
				for netKey, network := range area.OspfNetworks {
					netPrefix := prefix + fmt.Sprintf(".networks.%s", netKey)
					if len(network.AuthKeys) > 0 {
						checks.append(t, "TestCheckResourceAttr", netPrefix+".auth_keys.%", fmt.Sprintf("%d", len(network.AuthKeys)))
						for authKey, authValue := range network.AuthKeys {
							checks.append(t, "TestCheckResourceAttr", netPrefix+fmt.Sprintf(".auth_keys.%s", authKey), authValue)
						}
					}
					if network.AuthPassword != nil {
						checks.append(t, "TestCheckResourceAttr", netPrefix+".auth_password", *network.AuthPassword)
					}
					if network.AuthType != nil {
						checks.append(t, "TestCheckResourceAttr", netPrefix+".auth_type", *network.AuthType)
					}
					if network.BfdMinimumInterval != nil {
						checks.append(t, "TestCheckResourceAttr", netPrefix+".bfd_minimum_interval", fmt.Sprintf("%d", *network.BfdMinimumInterval))
					}
					if network.DeadInterval != nil {
						checks.append(t, "TestCheckResourceAttr", netPrefix+".dead_interval", fmt.Sprintf("%d", *network.DeadInterval))
					}
					if network.ExportPolicy != nil {
						checks.append(t, "TestCheckResourceAttr", netPrefix+".export_policy", *network.ExportPolicy)
					}
					if network.HelloInterval != nil {
						checks.append(t, "TestCheckResourceAttr", netPrefix+".hello_interval", fmt.Sprintf("%d", *network.HelloInterval))
					}
					if network.ImportPolicy != nil {
						checks.append(t, "TestCheckResourceAttr", netPrefix+".import_policy", *network.ImportPolicy)
					}
					if network.InterfaceType != nil {
						checks.append(t, "TestCheckResourceAttr", netPrefix+".interface_type", *network.InterfaceType)
					}
					if network.Metric != nil {
						checks.append(t, "TestCheckResourceAttr", netPrefix+".metric", fmt.Sprintf("%d", *network.Metric))
					}
					if network.NoReadvertiseToOverlay != nil {
						checks.append(t, "TestCheckResourceAttr", netPrefix+".no_readvertise_to_overlay", fmt.Sprintf("%t", *network.NoReadvertiseToOverlay))
					}
					if network.Passive != nil {
						checks.append(t, "TestCheckResourceAttr", netPrefix+".passive", fmt.Sprintf("%t", *network.Passive))
					}
				}
			}
		}
	}

	// Check port_mirroring map if present
	if len(s.PortMirroring) > 0 {
		checks.append(t, "TestCheckResourceAttr", "port_mirroring.%", fmt.Sprintf("%d", len(s.PortMirroring)))
		for key, portMirror := range s.PortMirroring {
			prefix := fmt.Sprintf("port_mirroring.%s", key)
			if len(portMirror.InputNetworksIngress) > 0 {
				checks.append(t, "TestCheckResourceAttr", prefix+".input_networks_ingress.#", fmt.Sprintf("%d", len(portMirror.InputNetworksIngress)))
				for i, network := range portMirror.InputNetworksIngress {
					checks.append(t, "TestCheckResourceAttr", prefix+fmt.Sprintf(".input_networks_ingress.%d", i), network)
				}
			}
			if len(portMirror.InputPortIdsEgress) > 0 {
				checks.append(t, "TestCheckResourceAttr", prefix+".input_port_ids_egress.#", fmt.Sprintf("%d", len(portMirror.InputPortIdsEgress)))
				for i, portId := range portMirror.InputPortIdsEgress {
					checks.append(t, "TestCheckResourceAttr", prefix+fmt.Sprintf(".input_port_ids_egress.%d", i), portId)
				}
			}
			if len(portMirror.InputPortIdsIngress) > 0 {
				checks.append(t, "TestCheckResourceAttr", prefix+".input_port_ids_ingress.#", fmt.Sprintf("%d", len(portMirror.InputPortIdsIngress)))
				for i, portId := range portMirror.InputPortIdsIngress {
					checks.append(t, "TestCheckResourceAttr", prefix+fmt.Sprintf(".input_port_ids_ingress.%d", i), portId)
				}
			}
			if portMirror.OutputIpAddress != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".output_ip_address", *portMirror.OutputIpAddress)
			}
			if portMirror.OutputNetwork != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".output_network", *portMirror.OutputNetwork)
			}
			if portMirror.OutputPortId != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".output_port_id", *portMirror.OutputPortId)
			}
		}
	}

	// Check port_usages map if present
	if len(s.PortUsages) > 0 {
		checks.append(t, "TestCheckResourceAttr", "port_usages.%", fmt.Sprintf("%d", len(s.PortUsages)))
		for key, usage := range s.PortUsages {
			prefix := fmt.Sprintf("port_usages.%s", key)
			if usage.AllNetworks != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".all_networks", fmt.Sprintf("%t", *usage.AllNetworks))
			}
			if usage.AllowDhcpd != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".allow_dhcpd", fmt.Sprintf("%t", *usage.AllowDhcpd))
			}
			if usage.AllowMultipleSupplicants != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".allow_multiple_supplicants", fmt.Sprintf("%t", *usage.AllowMultipleSupplicants))
			}
			if usage.BypassAuthWhenServerDown != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".bypass_auth_when_server_down", fmt.Sprintf("%t", *usage.BypassAuthWhenServerDown))
			}
			if usage.BypassAuthWhenServerDownForUnknownClient != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".bypass_auth_when_server_down_for_unknown_client", fmt.Sprintf("%t", *usage.BypassAuthWhenServerDownForUnknownClient))
			}
			if usage.Description != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".description", *usage.Description)
			}
			if usage.DisableAutoneg != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".disable_autoneg", fmt.Sprintf("%t", *usage.DisableAutoneg))
			}
			if usage.Disabled != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".disabled", fmt.Sprintf("%t", *usage.Disabled))
			}
			if usage.Duplex != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".duplex", *usage.Duplex)
			}
			if len(usage.DynamicVlanNetworks) > 0 {
				checks.append(t, "TestCheckResourceAttr", prefix+".dynamic_vlan_networks.#", fmt.Sprintf("%d", len(usage.DynamicVlanNetworks)))
				for i, network := range usage.DynamicVlanNetworks {
					checks.append(t, "TestCheckResourceAttr", prefix+fmt.Sprintf(".dynamic_vlan_networks.%d", i), network)
				}
			}
			if usage.EnableMacAuth != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".enable_mac_auth", fmt.Sprintf("%t", *usage.EnableMacAuth))
			}
			if usage.EnableQos != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".enable_qos", fmt.Sprintf("%t", *usage.EnableQos))
			}
			if usage.GuestNetwork != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".guest_network", *usage.GuestNetwork)
			}
			if usage.InterSwitchLink != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".inter_switch_link", fmt.Sprintf("%t", *usage.InterSwitchLink))
			}
			if usage.MacAuthOnly != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".mac_auth_only", fmt.Sprintf("%t", *usage.MacAuthOnly))
			}
			if usage.MacAuthPreferred != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".mac_auth_preferred", fmt.Sprintf("%t", *usage.MacAuthPreferred))
			}
			if usage.MacAuthProtocol != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".mac_auth_protocol", *usage.MacAuthProtocol)
			}
			if usage.MacLimit != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".mac_limit", *usage.MacLimit)
			}
			if usage.Mode != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".mode", *usage.Mode)
			}
			if usage.Mtu != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".mtu", *usage.Mtu)
			}
			if len(usage.Networks) > 0 {
				checks.append(t, "TestCheckResourceAttr", prefix+".networks.#", fmt.Sprintf("%d", len(usage.Networks)))
				for i, network := range usage.Networks {
					checks.append(t, "TestCheckResourceAttr", prefix+fmt.Sprintf(".networks.%d", i), network)
				}
			}
			if usage.PersistMac != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".persist_mac", fmt.Sprintf("%t", *usage.PersistMac))
			}
			if usage.PoeDisabled != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".poe_disabled", fmt.Sprintf("%t", *usage.PoeDisabled))
			}
			if usage.PortAuth != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".port_auth", *usage.PortAuth)
			}
			if usage.PortNetwork != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".port_network", *usage.PortNetwork)
			}
			if usage.ReauthInterval != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".reauth_interval", *usage.ReauthInterval)
			}
			if usage.ResetDefaultWhen != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".reset_default_when", *usage.ResetDefaultWhen)
			}
			if len(usage.Rules) > 0 {
				checks.append(t, "TestCheckResourceAttr", prefix+".rules.#", fmt.Sprintf("%d", len(usage.Rules)))
				for i, rule := range usage.Rules {
					rulePrefix := prefix + fmt.Sprintf(".rules.%d", i)
					if rule.Equals != nil {
						checks.append(t, "TestCheckResourceAttr", rulePrefix+".equals", *rule.Equals)
					}
					if len(rule.EqualsAny) > 0 {
						checks.append(t, "TestCheckResourceAttr", rulePrefix+".equals_any.#", fmt.Sprintf("%d", len(rule.EqualsAny)))
						for j, equals := range rule.EqualsAny {
							checks.append(t, "TestCheckResourceAttr", rulePrefix+fmt.Sprintf(".equals_any.%d", j), equals)
						}
					}
					if rule.Expression != nil {
						checks.append(t, "TestCheckResourceAttr", rulePrefix+".expression", *rule.Expression)
					}
					checks.append(t, "TestCheckResourceAttr", rulePrefix+".src", rule.Src)
					if rule.Usage != nil {
						checks.append(t, "TestCheckResourceAttr", rulePrefix+".usage", *rule.Usage)
					}
				}
			}
			if usage.ServerFailNetwork != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".server_fail_network", *usage.ServerFailNetwork)
			}
			if usage.ServerRejectNetwork != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".server_reject_network", *usage.ServerRejectNetwork)
			}
			if usage.Speed != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".speed", *usage.Speed)
			}
			if usage.StormControl != nil {
				if usage.StormControl.NoBroadcast != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".storm_control.no_broadcast", fmt.Sprintf("%t", *usage.StormControl.NoBroadcast))
				}
				if usage.StormControl.NoMulticast != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".storm_control.no_multicast", fmt.Sprintf("%t", *usage.StormControl.NoMulticast))
				}
				if usage.StormControl.NoRegisteredMulticast != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".storm_control.no_registered_multicast", fmt.Sprintf("%t", *usage.StormControl.NoRegisteredMulticast))
				}
				if usage.StormControl.NoUnknownUnicast != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".storm_control.no_unknown_unicast", fmt.Sprintf("%t", *usage.StormControl.NoUnknownUnicast))
				}
				if usage.StormControl.Percentage != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".storm_control.percentage", fmt.Sprintf("%d", *usage.StormControl.Percentage))
				}
			}
			if usage.StpDisable != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".stp_disable", fmt.Sprintf("%t", *usage.StpDisable))
			}
			if usage.StpEdge != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".stp_edge", fmt.Sprintf("%t", *usage.StpEdge))
			}
			if usage.StpNoRootPort != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".stp_no_root_port", fmt.Sprintf("%t", *usage.StpNoRootPort))
			}
			if usage.StpP2p != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".stp_p2p", fmt.Sprintf("%t", *usage.StpP2p))
			}
			if usage.StpRequired != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".stp_required", fmt.Sprintf("%t", *usage.StpRequired))
			}
			if usage.UiEvpntopoId != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".ui_evpntopo_id", *usage.UiEvpntopoId)
			}
			if usage.UseVstp != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".use_vstp", fmt.Sprintf("%t", *usage.UseVstp))
			}
			if usage.VoipNetwork != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".voip_network", *usage.VoipNetwork)
			}
		}
	}

	// Check radius_config if present
	if s.RadiusConfig != nil {
		if s.RadiusConfig.AcctImmediateUpdate != nil {
			checks.append(t, "TestCheckResourceAttr", "radius_config.acct_immediate_update", fmt.Sprintf("%t", *s.RadiusConfig.AcctImmediateUpdate))
		}
		if s.RadiusConfig.AcctInterimInterval != nil {
			checks.append(t, "TestCheckResourceAttr", "radius_config.acct_interim_interval", fmt.Sprintf("%d", *s.RadiusConfig.AcctInterimInterval))
		}
		if len(s.RadiusConfig.AcctServers) > 0 {
			checks.append(t, "TestCheckResourceAttr", "radius_config.acct_servers.#", fmt.Sprintf("%d", len(s.RadiusConfig.AcctServers)))
			for i, server := range s.RadiusConfig.AcctServers {
				prefix := fmt.Sprintf("radius_config.acct_servers.%d", i)
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
		if s.RadiusConfig.AuthServerSelection != nil {
			checks.append(t, "TestCheckResourceAttr", "radius_config.auth_server_selection", *s.RadiusConfig.AuthServerSelection)
		}
		if len(s.RadiusConfig.AuthServers) > 0 {
			checks.append(t, "TestCheckResourceAttr", "radius_config.auth_servers.#", fmt.Sprintf("%d", len(s.RadiusConfig.AuthServers)))
			for i, server := range s.RadiusConfig.AuthServers {
				prefix := fmt.Sprintf("radius_config.auth_servers.%d", i)
				checks.append(t, "TestCheckResourceAttr", prefix+".host", server.Host)
				if server.KeywrapEnabled != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".keywrap_enabled", fmt.Sprintf("%t", *server.KeywrapEnabled))
				}
			}
		}
		if s.RadiusConfig.AuthServersRetries != nil {
			checks.append(t, "TestCheckResourceAttr", "radius_config.auth_servers_retries", fmt.Sprintf("%d", *s.RadiusConfig.AuthServersRetries))
		}
		if s.RadiusConfig.AuthServersTimeout != nil {
			checks.append(t, "TestCheckResourceAttr", "radius_config.auth_servers_timeout", fmt.Sprintf("%d", *s.RadiusConfig.AuthServersTimeout))
		}
		if s.RadiusConfig.CoaEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "radius_config.coa_enabled", fmt.Sprintf("%t", *s.RadiusConfig.CoaEnabled))
		}
		if s.RadiusConfig.CoaPort != nil {
			checks.append(t, "TestCheckResourceAttr", "radius_config.coa_port", *s.RadiusConfig.CoaPort)
		}
		if s.RadiusConfig.FastDot1xTimers != nil {
			checks.append(t, "TestCheckResourceAttr", "radius_config.fast_dot1x_timers", fmt.Sprintf("%t", *s.RadiusConfig.FastDot1xTimers))
		}
		if s.RadiusConfig.Network != nil {
			checks.append(t, "TestCheckResourceAttr", "radius_config.network", *s.RadiusConfig.Network)
		}
		if s.RadiusConfig.SourceIp != nil {
			checks.append(t, "TestCheckResourceAttr", "radius_config.source_ip", *s.RadiusConfig.SourceIp)
		}
	}

	// Check remote_syslog if present
	if s.RemoteSyslog != nil {
		if s.RemoteSyslog.Archive != nil {
			if s.RemoteSyslog.Archive.Files != nil {
				checks.append(t, "TestCheckResourceAttr", "remote_syslog.archive.files", *s.RemoteSyslog.Archive.Files)
			}
			if s.RemoteSyslog.Archive.Size != nil {
				checks.append(t, "TestCheckResourceAttr", "remote_syslog.archive.size", *s.RemoteSyslog.Archive.Size)
			}
		}
		if len(s.RemoteSyslog.Cacerts) > 0 {
			checks.append(t, "TestCheckResourceAttr", "remote_syslog.cacerts.#", fmt.Sprintf("%d", len(s.RemoteSyslog.Cacerts)))
			for i, cert := range s.RemoteSyslog.Cacerts {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.cacerts.%d", i), cert)
			}
		}
		if s.RemoteSyslog.Console != nil && len(s.RemoteSyslog.Console.Contents) > 0 {
			checks.append(t, "TestCheckResourceAttr", "remote_syslog.console.contents.#", fmt.Sprintf("%d", len(s.RemoteSyslog.Console.Contents)))
		}
		if s.RemoteSyslog.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "remote_syslog.enabled", fmt.Sprintf("%t", *s.RemoteSyslog.Enabled))
		}
		if len(s.RemoteSyslog.Files) > 0 {
			checks.append(t, "TestCheckResourceAttr", "remote_syslog.files.#", fmt.Sprintf("%d", len(s.RemoteSyslog.Files)))
			for i, file := range s.RemoteSyslog.Files {
				prefix := fmt.Sprintf("remote_syslog.files.%d", i)
				if file.Archive != nil {
					if file.Archive.Files != nil {
						checks.append(t, "TestCheckResourceAttr", prefix+".archive.files", *file.Archive.Files)
					}
					if file.Archive.Size != nil {
						checks.append(t, "TestCheckResourceAttr", prefix+".archive.size", *file.Archive.Size)
					}
				}
				if len(file.Contents) > 0 {
					checks.append(t, "TestCheckResourceAttr", prefix+".contents.#", fmt.Sprintf("%d", len(file.Contents)))
					for j, content := range file.Contents {
						contentPrefix := prefix + fmt.Sprintf(".contents.%d", j)
						if content.Facility != nil {
							checks.append(t, "TestCheckResourceAttr", contentPrefix+".facility", *content.Facility)
						}
						if content.Severity != nil {
							checks.append(t, "TestCheckResourceAttr", contentPrefix+".severity", *content.Severity)
						}
					}
				}
				if file.EnableTls != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".enable_tls", fmt.Sprintf("%t", *file.EnableTls))
				}
				if file.ExplicitPriority != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".explicit_priority", fmt.Sprintf("%t", *file.ExplicitPriority))
				}
				if file.File != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".file", *file.File)
				}
				if file.Match != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".match", *file.Match)
				}
				if file.StructuredData != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".structured_data", fmt.Sprintf("%t", *file.StructuredData))
				}
			}
		}
		if s.RemoteSyslog.Network != nil {
			checks.append(t, "TestCheckResourceAttr", "remote_syslog.network", *s.RemoteSyslog.Network)
		}
		if s.RemoteSyslog.SendToAllServers != nil {
			checks.append(t, "TestCheckResourceAttr", "remote_syslog.send_to_all_servers", fmt.Sprintf("%t", *s.RemoteSyslog.SendToAllServers))
		}
		if len(s.RemoteSyslog.Servers) > 0 {
			checks.append(t, "TestCheckResourceAttr", "remote_syslog.servers.#", fmt.Sprintf("%d", len(s.RemoteSyslog.Servers)))
			for i, server := range s.RemoteSyslog.Servers {
				prefix := fmt.Sprintf("remote_syslog.servers.%d", i)
				if len(server.Contents) > 0 {
					checks.append(t, "TestCheckResourceAttr", prefix+".contents.#", fmt.Sprintf("%d", len(server.Contents)))
					for j, content := range server.Contents {
						contentPrefix := prefix + fmt.Sprintf(".contents.%d", j)
						if content.Facility != nil {
							checks.append(t, "TestCheckResourceAttr", contentPrefix+".facility", *content.Facility)
						}
						if content.Severity != nil {
							checks.append(t, "TestCheckResourceAttr", contentPrefix+".severity", *content.Severity)
						}
					}
				}
				if server.ExplicitPriority != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".explicit_priority", fmt.Sprintf("%t", *server.ExplicitPriority))
				}
				if server.Facility != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".facility", *server.Facility)
				}
				if server.Host != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".host", *server.Host)
				}
				if server.Match != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".match", *server.Match)
				}
				if server.Port != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".port", *server.Port)
				}
				if server.Protocol != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".protocol", *server.Protocol)
				}
				if server.RoutingInstance != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".routing_instance", *server.RoutingInstance)
				}
				if server.ServerName != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".server_name", *server.ServerName)
				}
				if server.Severity != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".severity", *server.Severity)
				}
				if server.SourceAddress != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".source_address", *server.SourceAddress)
				}
				if server.StructuredData != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".structured_data", fmt.Sprintf("%t", *server.StructuredData))
				}
				if server.Tag != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".tag", *server.Tag)
				}
			}
		}
		if s.RemoteSyslog.TimeFormat != nil {
			checks.append(t, "TestCheckResourceAttr", "remote_syslog.time_format", *s.RemoteSyslog.TimeFormat)
		}
		if len(s.RemoteSyslog.Users) > 0 {
			checks.append(t, "TestCheckResourceAttr", "remote_syslog.users.#", fmt.Sprintf("%d", len(s.RemoteSyslog.Users)))
			for i, user := range s.RemoteSyslog.Users {
				prefix := fmt.Sprintf("remote_syslog.users.%d", i)
				if len(user.Contents) > 0 {
					checks.append(t, "TestCheckResourceAttr", prefix+".contents.#", fmt.Sprintf("%d", len(user.Contents)))
					for j, content := range user.Contents {
						contentPrefix := prefix + fmt.Sprintf(".contents.%d", j)
						if content.Facility != nil {
							checks.append(t, "TestCheckResourceAttr", contentPrefix+".facility", *content.Facility)
						}
						if content.Severity != nil {
							checks.append(t, "TestCheckResourceAttr", contentPrefix+".severity", *content.Severity)
						}
					}
				}
				if user.Match != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".match", *user.Match)
				}
				if user.User != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".user", *user.User)
				}
			}
		}
	}

	// Check snmp_config if present
	if s.SnmpConfig != nil {
		if len(s.SnmpConfig.ClientList) > 0 {
			checks.append(t, "TestCheckResourceAttr", "snmp_config.client_list.#", fmt.Sprintf("%d", len(s.SnmpConfig.ClientList)))
			for i, client := range s.SnmpConfig.ClientList {
				prefix := fmt.Sprintf("snmp_config.client_list.%d", i)
				if client.ClientListName != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".client_list_name", *client.ClientListName)
				}
				if len(client.Clients) > 0 {
					checks.append(t, "TestCheckResourceAttr", prefix+".clients.#", fmt.Sprintf("%d", len(client.Clients)))
					for j, clientAddr := range client.Clients {
						checks.append(t, "TestCheckResourceAttr", prefix+fmt.Sprintf(".clients.%d", j), clientAddr)
					}
				}
			}
		}
		if s.SnmpConfig.Contact != nil {
			checks.append(t, "TestCheckResourceAttr", "snmp_config.contact", *s.SnmpConfig.Contact)
		}
		if s.SnmpConfig.Description != nil {
			checks.append(t, "TestCheckResourceAttr", "snmp_config.description", *s.SnmpConfig.Description)
		}
		if s.SnmpConfig.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "snmp_config.enabled", fmt.Sprintf("%t", *s.SnmpConfig.Enabled))
		}
		if s.SnmpConfig.EngineId != nil {
			checks.append(t, "TestCheckResourceAttr", "snmp_config.engine_id", *s.SnmpConfig.EngineId)
		}
		if s.SnmpConfig.EngineIdType != nil {
			checks.append(t, "TestCheckResourceAttr", "snmp_config.engine_id_type", *s.SnmpConfig.EngineIdType)
		}
		if s.SnmpConfig.Location != nil {
			checks.append(t, "TestCheckResourceAttr", "snmp_config.location", *s.SnmpConfig.Location)
		}
		if s.SnmpConfig.Name != nil {
			checks.append(t, "TestCheckResourceAttr", "snmp_config.name", *s.SnmpConfig.Name)
		}
		if s.SnmpConfig.Network != nil {
			checks.append(t, "TestCheckResourceAttr", "snmp_config.network", *s.SnmpConfig.Network)
		}
		if len(s.SnmpConfig.TrapGroups) > 0 {
			checks.append(t, "TestCheckResourceAttr", "snmp_config.trap_groups.#", fmt.Sprintf("%d", len(s.SnmpConfig.TrapGroups)))
			for i, group := range s.SnmpConfig.TrapGroups {
				prefix := fmt.Sprintf("snmp_config.trap_groups.%d", i)
				if len(group.Categories) > 0 {
					checks.append(t, "TestCheckResourceAttr", prefix+".categories.#", fmt.Sprintf("%d", len(group.Categories)))
					for j, category := range group.Categories {
						checks.append(t, "TestCheckResourceAttr", prefix+fmt.Sprintf(".categories.%d", j), category)
					}
				}
				if group.GroupName != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".group_name", *group.GroupName)
				}
				if len(group.Targets) > 0 {
					checks.append(t, "TestCheckResourceAttr", prefix+".targets.#", fmt.Sprintf("%d", len(group.Targets)))
					for j, target := range group.Targets {
						checks.append(t, "TestCheckResourceAttr", prefix+fmt.Sprintf(".targets.%d", j), target)
					}
				}
				if group.Version != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".version", *group.Version)
				}
			}
		}
		if len(s.SnmpConfig.V2cConfig) > 0 {
			checks.append(t, "TestCheckResourceAttr", "snmp_config.v2c_config.#", fmt.Sprintf("%d", len(s.SnmpConfig.V2cConfig)))
			for i, config := range s.SnmpConfig.V2cConfig {
				prefix := fmt.Sprintf("snmp_config.v2c_config.%d", i)
				if config.Authorization != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".authorization", *config.Authorization)
				}
				if config.ClientListName != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".client_list_name", *config.ClientListName)
				}
				if config.CommunityName != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".community_name", *config.CommunityName)
				}
				if config.View != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".view", *config.View)
				}
			}
		}
		if s.SnmpConfig.V3Config != nil {
			if len(s.SnmpConfig.V3Config.Notify) > 0 {
				checks.append(t, "TestCheckResourceAttr", "snmp_config.v3_config.notify.#", fmt.Sprintf("%d", len(s.SnmpConfig.V3Config.Notify)))
				for i, notify := range s.SnmpConfig.V3Config.Notify {
					prefix := fmt.Sprintf("snmp_config.v3_config.notify.%d", i)
					checks.append(t, "TestCheckResourceAttr", prefix+".name", notify.Name)
					checks.append(t, "TestCheckResourceAttr", prefix+".tag", notify.Tag)
					checks.append(t, "TestCheckResourceAttr", prefix+".type", notify.NotifyType)
				}
			}
			if len(s.SnmpConfig.V3Config.NotifyFilter) > 0 {
				checks.append(t, "TestCheckResourceAttr", "snmp_config.v3_config.notify_filter.#", fmt.Sprintf("%d", len(s.SnmpConfig.V3Config.NotifyFilter)))
				for i, filter := range s.SnmpConfig.V3Config.NotifyFilter {
					prefix := fmt.Sprintf("snmp_config.v3_config.notify_filter.%d", i)
					if filter.ProfileName != nil {
						checks.append(t, "TestCheckResourceAttr", prefix+".profile_name", *filter.ProfileName)
					}
					if len(filter.Snmpv3Contents) > 0 {
						checks.append(t, "TestCheckResourceAttr", prefix+".contents.#", fmt.Sprintf("%d", len(filter.Snmpv3Contents)))
						for j, content := range filter.Snmpv3Contents {
							contentPrefix := prefix + fmt.Sprintf(".contents.%d", j)
							if content.Include != nil {
								checks.append(t, "TestCheckResourceAttr", contentPrefix+".include", fmt.Sprintf("%t", *content.Include))
							}
							checks.append(t, "TestCheckResourceAttr", contentPrefix+".oid", content.Oid)
						}
					}
				}
			}
			if len(s.SnmpConfig.V3Config.TargetAddress) > 0 {
				checks.append(t, "TestCheckResourceAttr", "snmp_config.v3_config.target_address.#", fmt.Sprintf("%d", len(s.SnmpConfig.V3Config.TargetAddress)))
				for i, target := range s.SnmpConfig.V3Config.TargetAddress {
					prefix := fmt.Sprintf("snmp_config.v3_config.target_address.%d", i)
					checks.append(t, "TestCheckResourceAttr", prefix+".address", target.Address)
					checks.append(t, "TestCheckResourceAttr", prefix+".address_mask", target.AddressMask)
					if target.Port != nil {
						checks.append(t, "TestCheckResourceAttr", prefix+".port", *target.Port)
					}
					if target.TagList != nil {
						checks.append(t, "TestCheckResourceAttr", prefix+".tag_list", *target.TagList)
					}
					checks.append(t, "TestCheckResourceAttr", prefix+".target_address_name", target.TargetAddressName)
					if target.TargetParameters != nil {
						checks.append(t, "TestCheckResourceAttr", prefix+".target_parameters", *target.TargetParameters)
					}
				}
			}
			if len(s.SnmpConfig.V3Config.TargetParameters) > 0 {
				checks.append(t, "TestCheckResourceAttr", "snmp_config.v3_config.target_parameters.#", fmt.Sprintf("%d", len(s.SnmpConfig.V3Config.TargetParameters)))
				for i, param := range s.SnmpConfig.V3Config.TargetParameters {
					prefix := fmt.Sprintf("snmp_config.v3_config.target_parameters.%d", i)
					checks.append(t, "TestCheckResourceAttr", prefix+".message_processing_model", param.MessageProcessingModel)
					checks.append(t, "TestCheckResourceAttr", prefix+".name", param.Name)
					if param.NotifyFilter != nil {
						checks.append(t, "TestCheckResourceAttr", prefix+".notify_filter", *param.NotifyFilter)
					}
					if param.SecurityLevel != nil {
						checks.append(t, "TestCheckResourceAttr", prefix+".security_level", *param.SecurityLevel)
					}
					if param.SecurityModel != nil {
						checks.append(t, "TestCheckResourceAttr", prefix+".security_model", *param.SecurityModel)
					}
					if param.SecurityName != nil {
						checks.append(t, "TestCheckResourceAttr", prefix+".security_name", *param.SecurityName)
					}
				}
			}
			if len(s.SnmpConfig.V3Config.Usm) > 0 {
				checks.append(t, "TestCheckResourceAttr", "snmp_config.v3_config.usm.#", fmt.Sprintf("%d", len(s.SnmpConfig.V3Config.Usm)))
				for i, usm := range s.SnmpConfig.V3Config.Usm {
					prefix := fmt.Sprintf("snmp_config.v3_config.usm.%d", i)
					checks.append(t, "TestCheckResourceAttr", prefix+".engine_type", usm.EngineType)
					if usm.RemoteEngineId != nil {
						checks.append(t, "TestCheckResourceAttr", prefix+".remote_engine_id", *usm.RemoteEngineId)
					}
					if len(usm.Snmpv3Users) > 0 {
						checks.append(t, "TestCheckResourceAttr", prefix+".users.#", fmt.Sprintf("%d", len(usm.Snmpv3Users)))
						for j, user := range usm.Snmpv3Users {
							userPrefix := prefix + fmt.Sprintf(".users.%d", j)
							if user.AuthenticationPassword != nil {
								checks.append(t, "TestCheckResourceAttr", userPrefix+".authentication_password", *user.AuthenticationPassword)
							}
							if user.AuthenticationType != nil {
								checks.append(t, "TestCheckResourceAttr", userPrefix+".authentication_type", *user.AuthenticationType)
							}
							if user.EncryptionPassword != nil {
								checks.append(t, "TestCheckResourceAttr", userPrefix+".encryption_password", *user.EncryptionPassword)
							}
							if user.EncryptionType != nil {
								checks.append(t, "TestCheckResourceAttr", userPrefix+".encryption_type", *user.EncryptionType)
							}
							if user.Name != nil {
								checks.append(t, "TestCheckResourceAttr", userPrefix+".name", *user.Name)
							}
						}
					}
				}
			}
			if s.SnmpConfig.V3Config.Vacm != nil {
				if len(s.SnmpConfig.V3Config.Vacm.Access) > 0 {
					checks.append(t, "TestCheckResourceAttr", "snmp_config.v3_config.vacm.access.#", fmt.Sprintf("%d", len(s.SnmpConfig.V3Config.Vacm.Access)))
					for i, access := range s.SnmpConfig.V3Config.Vacm.Access {
						prefix := fmt.Sprintf("snmp_config.v3_config.vacm.access.%d", i)
						if access.GroupName != nil {
							checks.append(t, "TestCheckResourceAttr", prefix+".group_name", *access.GroupName)
						}
						if len(access.PrefixList) > 0 {
							checks.append(t, "TestCheckResourceAttr", prefix+".prefix_list.#", fmt.Sprintf("%d", len(access.PrefixList)))
							for j, prefixItem := range access.PrefixList {
								prefixPrefix := prefix + fmt.Sprintf(".prefix_list.%d", j)
								if prefixItem.ContextPrefix != nil {
									checks.append(t, "TestCheckResourceAttr", prefixPrefix+".context_prefix", *prefixItem.ContextPrefix)
								}
								if prefixItem.NotifyView != nil {
									checks.append(t, "TestCheckResourceAttr", prefixPrefix+".notify_view", *prefixItem.NotifyView)
								}
								if prefixItem.ReadView != nil {
									checks.append(t, "TestCheckResourceAttr", prefixPrefix+".read_view", *prefixItem.ReadView)
								}
								if prefixItem.SecurityLevel != nil {
									checks.append(t, "TestCheckResourceAttr", prefixPrefix+".security_level", *prefixItem.SecurityLevel)
								}
								if prefixItem.SecurityModel != nil {
									checks.append(t, "TestCheckResourceAttr", prefixPrefix+".security_model", *prefixItem.SecurityModel)
								}
								if prefixItem.PrefixListType != nil {
									checks.append(t, "TestCheckResourceAttr", prefixPrefix+".type", *prefixItem.PrefixListType)
								}
								if prefixItem.WriteView != nil {
									checks.append(t, "TestCheckResourceAttr", prefixPrefix+".write_view", *prefixItem.WriteView)
								}
							}
						}
					}
				}
				if s.SnmpConfig.V3Config.Vacm.SecurityToGroup != nil {
					if s.SnmpConfig.V3Config.Vacm.SecurityToGroup.SecurityModel != nil {
						checks.append(t, "TestCheckResourceAttr", "snmp_config.v3_config.vacm.security_to_group.security_model", *s.SnmpConfig.V3Config.Vacm.SecurityToGroup.SecurityModel)
					}
					if len(s.SnmpConfig.V3Config.Vacm.SecurityToGroup.Snmpv3VacmContent) > 0 {
						checks.append(t, "TestCheckResourceAttr", "snmp_config.v3_config.vacm.security_to_group.content.#", fmt.Sprintf("%d", len(s.SnmpConfig.V3Config.Vacm.SecurityToGroup.Snmpv3VacmContent)))
						for i, content := range s.SnmpConfig.V3Config.Vacm.SecurityToGroup.Snmpv3VacmContent {
							prefix := fmt.Sprintf("snmp_config.v3_config.vacm.security_to_group.content.%d", i)
							if content.Group != nil {
								checks.append(t, "TestCheckResourceAttr", prefix+".group", *content.Group)
							}
							if content.SecurityName != nil {
								checks.append(t, "TestCheckResourceAttr", prefix+".security_name", *content.SecurityName)
							}
						}
					}
				}
			}
		}
		if len(s.SnmpConfig.Views) > 0 {
			checks.append(t, "TestCheckResourceAttr", "snmp_config.views.#", fmt.Sprintf("%d", len(s.SnmpConfig.Views)))
			for i, view := range s.SnmpConfig.Views {
				prefix := fmt.Sprintf("snmp_config.views.%d", i)
				if view.Include != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".include", fmt.Sprintf("%t", *view.Include))
				}
				if view.Oid != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".oid", *view.Oid)
				}
				if view.ViewName != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".view_name", *view.ViewName)
				}
			}
		}
	}

	// Check switch_matching if present
	if s.SwitchMatching != nil {
		if s.SwitchMatching.Enable != nil {
			checks.append(t, "TestCheckResourceAttr", "switch_matching.enable", fmt.Sprintf("%t", *s.SwitchMatching.Enable))
		}
		if len(s.SwitchMatching.MatchingRules) > 0 {
			checks.append(t, "TestCheckResourceAttr", "switch_matching.rules.#", fmt.Sprintf("%d", len(s.SwitchMatching.MatchingRules)))
			for i, rule := range s.SwitchMatching.MatchingRules {
				prefix := fmt.Sprintf("switch_matching.rules.%d", i)
				if len(rule.AdditionalConfigCmds) > 0 {
					checks.append(t, "TestCheckResourceAttr", prefix+".additional_config_cmds.#", fmt.Sprintf("%d", len(rule.AdditionalConfigCmds)))
					for j, cmd := range rule.AdditionalConfigCmds {
						checks.append(t, "TestCheckResourceAttr", prefix+fmt.Sprintf(".additional_config_cmds.%d", j), cmd)
					}
				}
				if rule.IpConfig != nil {
					if rule.IpConfig.Network != nil {
						checks.append(t, "TestCheckResourceAttr", prefix+".ip_config.network", *rule.IpConfig.Network)
					}
					if rule.IpConfig.IpConfigType != nil {
						checks.append(t, "TestCheckResourceAttr", prefix+".ip_config.type", *rule.IpConfig.IpConfigType)
					}
				}
				if rule.MatchModel != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".match_model", *rule.MatchModel)
				}
				if rule.MatchName != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".match_name", *rule.MatchName)
				}
				if rule.MatchNameOffset != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".match_name_offset", fmt.Sprintf("%d", *rule.MatchNameOffset))
				}
				if rule.MatchRole != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".match_role", *rule.MatchRole)
				}
				if rule.Name != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".name", *rule.Name)
				}
				if rule.OobIpConfig != nil {
					if rule.OobIpConfig.OobIpConfigType != nil {
						checks.append(t, "TestCheckResourceAttr", prefix+".oob_ip_config.type", *rule.OobIpConfig.OobIpConfigType)
					}
					if rule.OobIpConfig.UseMgmtVrf != nil {
						checks.append(t, "TestCheckResourceAttr", prefix+".oob_ip_config.use_mgmt_vrf", fmt.Sprintf("%t", *rule.OobIpConfig.UseMgmtVrf))
					}
					if rule.OobIpConfig.UseMgmtVrfForHostOut != nil {
						checks.append(t, "TestCheckResourceAttr", prefix+".oob_ip_config.use_mgmt_vrf_for_host_out", fmt.Sprintf("%t", *rule.OobIpConfig.UseMgmtVrfForHostOut))
					}
				}
				if len(rule.PortConfig) > 0 {
					checks.append(t, "TestCheckResourceAttr", prefix+".port_config.%", fmt.Sprintf("%d", len(rule.PortConfig)))
				}
				if len(rule.PortMirroring) > 0 {
					checks.append(t, "TestCheckResourceAttr", prefix+".port_mirroring.%", fmt.Sprintf("%d", len(rule.PortMirroring)))
				}
				if rule.StpConfig != nil {
					if rule.StpConfig.BridgePriority != nil {
						checks.append(t, "TestCheckResourceAttr", prefix+".stp_config.bridge_priority", *rule.StpConfig.BridgePriority)
					}
				}
			}
		}
	}

	// Check switch_mgmt if present
	if s.SwitchMgmt != nil {
		if s.SwitchMgmt.ApAffinityThreshold != nil {
			checks.append(t, "TestCheckResourceAttr", "switch_mgmt.ap_affinity_threshold", fmt.Sprintf("%d", *s.SwitchMgmt.ApAffinityThreshold))
		}
		if s.SwitchMgmt.CliBanner != nil {
			checks.append(t, "TestCheckResourceAttr", "switch_mgmt.cli_banner", *s.SwitchMgmt.CliBanner)
		}
		if s.SwitchMgmt.CliIdleTimeout != nil {
			checks.append(t, "TestCheckResourceAttr", "switch_mgmt.cli_idle_timeout", fmt.Sprintf("%d", *s.SwitchMgmt.CliIdleTimeout))
		}
		if s.SwitchMgmt.ConfigRevertTimer != nil {
			checks.append(t, "TestCheckResourceAttr", "switch_mgmt.config_revert_timer", fmt.Sprintf("%d", *s.SwitchMgmt.ConfigRevertTimer))
		}
		if s.SwitchMgmt.DhcpOptionFqdn != nil {
			checks.append(t, "TestCheckResourceAttr", "switch_mgmt.dhcp_option_fqdn", fmt.Sprintf("%t", *s.SwitchMgmt.DhcpOptionFqdn))
		}
		if s.SwitchMgmt.DisableOobDownAlarm != nil {
			checks.append(t, "TestCheckResourceAttr", "switch_mgmt.disable_oob_down_alarm", fmt.Sprintf("%t", *s.SwitchMgmt.DisableOobDownAlarm))
		}
		if s.SwitchMgmt.FipsEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "switch_mgmt.fips_enabled", fmt.Sprintf("%t", *s.SwitchMgmt.FipsEnabled))
		}
		if len(s.SwitchMgmt.LocalAccounts) > 0 {
			checks.append(t, "TestCheckResourceAttr", "switch_mgmt.local_accounts.%", fmt.Sprintf("%d", len(s.SwitchMgmt.LocalAccounts)))
			for key, account := range s.SwitchMgmt.LocalAccounts {
				prefix := fmt.Sprintf("switch_mgmt.local_accounts.%s", key)
				if account.Password != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".password", *account.Password)
				}
				if account.Role != nil {
					checks.append(t, "TestCheckResourceAttr", prefix+".role", *account.Role)
				}
			}
		}
		if s.SwitchMgmt.MxedgeProxyHost != nil {
			checks.append(t, "TestCheckResourceAttr", "switch_mgmt.mxedge_proxy_host", *s.SwitchMgmt.MxedgeProxyHost)
		}
		if s.SwitchMgmt.MxedgeProxyPort != nil {
			checks.append(t, "TestCheckResourceAttr", "switch_mgmt.mxedge_proxy_port", *s.SwitchMgmt.MxedgeProxyPort)
		}
		if s.SwitchMgmt.ProtectRe != nil {
			if len(s.SwitchMgmt.ProtectRe.AllowedServices) > 0 {
				checks.append(t, "TestCheckResourceAttr", "switch_mgmt.protect_re.allowed_services.#", fmt.Sprintf("%d", len(s.SwitchMgmt.ProtectRe.AllowedServices)))
				for i, service := range s.SwitchMgmt.ProtectRe.AllowedServices {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("switch_mgmt.protect_re.allowed_services.%d", i), service)
				}
			}
			if len(s.SwitchMgmt.ProtectRe.Custom) > 0 {
				checks.append(t, "TestCheckResourceAttr", "switch_mgmt.protect_re.custom.#", fmt.Sprintf("%d", len(s.SwitchMgmt.ProtectRe.Custom)))
				for i, custom := range s.SwitchMgmt.ProtectRe.Custom {
					prefix := fmt.Sprintf("switch_mgmt.protect_re.custom.%d", i)
					if custom.PortRange != nil {
						checks.append(t, "TestCheckResourceAttr", prefix+".port_range", *custom.PortRange)
					}
					if custom.Protocol != nil {
						checks.append(t, "TestCheckResourceAttr", prefix+".protocol", *custom.Protocol)
					}
					if len(custom.Subnets) > 0 {
						checks.append(t, "TestCheckResourceAttr", prefix+".subnets.#", fmt.Sprintf("%d", len(custom.Subnets)))
						for j, subnet := range custom.Subnets {
							checks.append(t, "TestCheckResourceAttr", prefix+fmt.Sprintf(".subnets.%d", j), subnet)
						}
					}
				}
			}
			if s.SwitchMgmt.ProtectRe.Enabled != nil {
				checks.append(t, "TestCheckResourceAttr", "switch_mgmt.protect_re.enabled", fmt.Sprintf("%t", *s.SwitchMgmt.ProtectRe.Enabled))
			}
			if s.SwitchMgmt.ProtectRe.HitCount != nil {
				checks.append(t, "TestCheckResourceAttr", "switch_mgmt.protect_re.hit_count", fmt.Sprintf("%t", *s.SwitchMgmt.ProtectRe.HitCount))
			}
			if len(s.SwitchMgmt.ProtectRe.TrustedHosts) > 0 {
				checks.append(t, "TestCheckResourceAttr", "switch_mgmt.protect_re.trusted_hosts.#", fmt.Sprintf("%d", len(s.SwitchMgmt.ProtectRe.TrustedHosts)))
				for i, host := range s.SwitchMgmt.ProtectRe.TrustedHosts {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("switch_mgmt.protect_re.trusted_hosts.%d", i), host)
				}
			}
		}
		if s.SwitchMgmt.RemoveExistingConfigs != nil {
			checks.append(t, "TestCheckResourceAttr", "switch_mgmt.remove_existing_configs", fmt.Sprintf("%t", *s.SwitchMgmt.RemoveExistingConfigs))
		}
		if s.SwitchMgmt.RootPassword != nil {
			checks.append(t, "TestCheckResourceAttr", "switch_mgmt.root_password", *s.SwitchMgmt.RootPassword)
		}
		if s.SwitchMgmt.Tacacs != nil {
			if s.SwitchMgmt.Tacacs.DefaultRole != nil {
				checks.append(t, "TestCheckResourceAttr", "switch_mgmt.tacacs.default_role", *s.SwitchMgmt.Tacacs.DefaultRole)
			}
			if s.SwitchMgmt.Tacacs.Enabled != nil {
				checks.append(t, "TestCheckResourceAttr", "switch_mgmt.tacacs.enabled", fmt.Sprintf("%t", *s.SwitchMgmt.Tacacs.Enabled))
			}
			if s.SwitchMgmt.Tacacs.Network != nil {
				checks.append(t, "TestCheckResourceAttr", "switch_mgmt.tacacs.network", *s.SwitchMgmt.Tacacs.Network)
			}
			if len(s.SwitchMgmt.Tacacs.TacacctServers) > 0 {
				checks.append(t, "TestCheckResourceAttr", "switch_mgmt.tacacs.acct_servers.#", fmt.Sprintf("%d", len(s.SwitchMgmt.Tacacs.TacacctServers)))
				for i, server := range s.SwitchMgmt.Tacacs.TacacctServers {
					prefix := fmt.Sprintf("switch_mgmt.tacacs.acct_servers.%d", i)
					if server.Host != nil {
						checks.append(t, "TestCheckResourceAttr", prefix+".host", *server.Host)
					}
					if server.Port != nil {
						checks.append(t, "TestCheckResourceAttr", prefix+".port", *server.Port)
					}
					if server.Secret != nil {
						checks.append(t, "TestCheckResourceAttr", prefix+".secret", *server.Secret)
					}
					if server.Timeout != nil {
						checks.append(t, "TestCheckResourceAttr", prefix+".timeout", fmt.Sprintf("%d", *server.Timeout))
					}
				}
			}
			if len(s.SwitchMgmt.Tacacs.TacplusServers) > 0 {
				checks.append(t, "TestCheckResourceAttr", "switch_mgmt.tacacs.tacplus_servers.#", fmt.Sprintf("%d", len(s.SwitchMgmt.Tacacs.TacplusServers)))
				for i, server := range s.SwitchMgmt.Tacacs.TacplusServers {
					prefix := fmt.Sprintf("switch_mgmt.tacacs.tacplus_servers.%d", i)
					if server.Host != nil {
						checks.append(t, "TestCheckResourceAttr", prefix+".host", *server.Host)
					}
					if server.Port != nil {
						checks.append(t, "TestCheckResourceAttr", prefix+".port", *server.Port)
					}
					if server.Secret != nil {
						checks.append(t, "TestCheckResourceAttr", prefix+".secret", *server.Secret)
					}
					if server.Timeout != nil {
						checks.append(t, "TestCheckResourceAttr", prefix+".timeout", fmt.Sprintf("%d", *server.Timeout))
					}
				}
			}
		}
		if s.SwitchMgmt.UseMxedgeProxy != nil {
			checks.append(t, "TestCheckResourceAttr", "switch_mgmt.use_mxedge_proxy", fmt.Sprintf("%t", *s.SwitchMgmt.UseMxedgeProxy))
		}
	}

	// Check vrf_config if present
	if s.VrfConfig != nil {
		if s.VrfConfig.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "vrf_config.enabled", fmt.Sprintf("%t", *s.VrfConfig.Enabled))
		}
	}

	// Check vrf_instances map if present
	if len(s.VrfInstances) > 0 {
		checks.append(t, "TestCheckResourceAttr", "vrf_instances.%", fmt.Sprintf("%d", len(s.VrfInstances)))
		for key, vrf := range s.VrfInstances {
			prefix := fmt.Sprintf("vrf_instances.%s", key)
			if vrf.EvpnAutoLoopbackSubnet != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".evpn_auto_loopback_subnet", *vrf.EvpnAutoLoopbackSubnet)
			}
			if vrf.EvpnAutoLoopbackSubnet6 != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".evpn_auto_loopback_subnet6", *vrf.EvpnAutoLoopbackSubnet6)
			}
			if len(vrf.Networks) > 0 {
				checks.append(t, "TestCheckResourceAttr", prefix+".networks.#", fmt.Sprintf("%d", len(vrf.Networks)))
				for i, network := range vrf.Networks {
					checks.append(t, "TestCheckResourceAttr", prefix+fmt.Sprintf(".networks.%d", i), network)
				}
			}
			if len(vrf.VrfExtraRoutes) > 0 {
				checks.append(t, "TestCheckResourceAttr", prefix+".extra_routes.%", fmt.Sprintf("%d", len(vrf.VrfExtraRoutes)))
				for routeKey, route := range vrf.VrfExtraRoutes {
					routePrefix := prefix + fmt.Sprintf(".extra_routes.%s", routeKey)
					checks.append(t, "TestCheckResourceAttr", routePrefix+".via", route.Via)
				}
			}
			if len(vrf.VrfExtraRoutes6) > 0 {
				checks.append(t, "TestCheckResourceAttr", prefix+".extra_routes6.%", fmt.Sprintf("%d", len(vrf.VrfExtraRoutes6)))
				for routeKey, route := range vrf.VrfExtraRoutes6 {
					routePrefix := prefix + fmt.Sprintf(".extra_routes6.%s", routeKey)
					if route.Via != nil {
						checks.append(t, "TestCheckResourceAttr", routePrefix+".via", *route.Via)
					}
				}
			}
		}
	}

	return checks
}
