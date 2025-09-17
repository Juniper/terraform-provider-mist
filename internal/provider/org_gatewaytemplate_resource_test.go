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

func TestOrgGatewaytemplateModel(t *testing.T) {
	type testStep struct {
		config OrgGatewaytemplateModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgGatewaytemplateModel{
						Name:  "test_gateway_template",
						OrgId: GetTestOrgId(),
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/org_gatewaytemplate_resource/org_gatewaytemplate_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		var FixtureOrgGatewaytemplateModel OrgGatewaytemplateModel
		err = hcl.Decode(&FixtureOrgGatewaytemplateModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		FixtureOrgGatewaytemplateModel.OrgId = GetTestOrgId()

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureOrgGatewaytemplateModel,
				},
			},
		}
	}

	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			resourceType := "org_gatewaytemplate"

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

func (s *OrgGatewaytemplateModel) testChecks(t testing.TB, rType, rName string) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + rName)
	// Required parameters
	checks.append(t, "TestCheckResourceAttrSet", "org_id")
	checks.append(t, "TestCheckResourceAttr", "name", s.Name)

	// Optional parameters
	if len(s.AdditionalConfigCmds) > 0 {
		checks.append(t, "TestCheckResourceAttr", "additional_config_cmds.#", fmt.Sprintf("%d", len(s.AdditionalConfigCmds)))
		for i, cmd := range s.AdditionalConfigCmds {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("additional_config_cmds.%d", i), cmd)
		}
	}
	if len(s.BgpConfig) > 0 {
		checks.append(t, "TestCheckResourceAttr", "bgp_config.%", fmt.Sprintf("%d", len(s.BgpConfig)))
		for key, bgpConfig := range s.BgpConfig {
			if bgpConfig.AuthKey != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.auth_key", key), *bgpConfig.AuthKey)
			}
			if bgpConfig.BfdMinimumInterval != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.bfd_minimum_interval", key), fmt.Sprintf("%d", *bgpConfig.BfdMinimumInterval))
			}
			if bgpConfig.BfdMultiplier != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.bfd_multiplier", key), fmt.Sprintf("%d", *bgpConfig.BfdMultiplier))
			}
			if bgpConfig.DisableBfd != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.disable_bfd", key), fmt.Sprintf("%t", *bgpConfig.DisableBfd))
			}
			if bgpConfig.Export != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.export", key), *bgpConfig.Export)
			}
			if bgpConfig.ExportPolicy != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.export_policy", key), *bgpConfig.ExportPolicy)
			}
			if bgpConfig.ExtendedV4Nexthop != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.extended_v4_nexthop", key), fmt.Sprintf("%t", *bgpConfig.ExtendedV4Nexthop))
			}
			if bgpConfig.GracefulRestartTime != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.graceful_restart_time", key), fmt.Sprintf("%d", *bgpConfig.GracefulRestartTime))
			}
			if bgpConfig.HoldTime != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.hold_time", key), fmt.Sprintf("%d", *bgpConfig.HoldTime))
			}
			if bgpConfig.Import != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.import", key), *bgpConfig.Import)
			}
			if bgpConfig.ImportPolicy != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.import_policy", key), *bgpConfig.ImportPolicy)
			}
			if bgpConfig.LocalAs != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.local_as", key), *bgpConfig.LocalAs)
			}
			if bgpConfig.NeighborAs != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.neighbor_as", key), *bgpConfig.NeighborAs)
			}
			if len(bgpConfig.Neighbors) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.neighbors.%%", key), fmt.Sprintf("%d", len(bgpConfig.Neighbors)))
				for neighborKey, neighbor := range bgpConfig.Neighbors {
					if neighbor.Disabled != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.neighbors.%s.disabled", key, neighborKey), fmt.Sprintf("%t", *neighbor.Disabled))
					}
					if neighbor.ExportPolicy != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.neighbors.%s.export_policy", key, neighborKey), *neighbor.ExportPolicy)
					}
					if neighbor.HoldTime != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.neighbors.%s.hold_time", key, neighborKey), fmt.Sprintf("%d", *neighbor.HoldTime))
					}
					if neighbor.ImportPolicy != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.neighbors.%s.import_policy", key, neighborKey), *neighbor.ImportPolicy)
					}
					if neighbor.MultihopTtl != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.neighbors.%s.multihop_ttl", key, neighborKey), fmt.Sprintf("%d", *neighbor.MultihopTtl))
					}
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.neighbors.%s.neighbor_as", key, neighborKey), neighbor.NeighborAs)
				}
			}
			if len(bgpConfig.Networks) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.networks.#", key), fmt.Sprintf("%d", len(bgpConfig.Networks)))
				for i, network := range bgpConfig.Networks {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.networks.%d", key, i), network)
				}
			}
			if bgpConfig.NoPrivateAs != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.no_private_as", key), fmt.Sprintf("%t", *bgpConfig.NoPrivateAs))
			}
			if bgpConfig.NoReadvertiseToOverlay != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.no_readvertise_to_overlay", key), fmt.Sprintf("%t", *bgpConfig.NoReadvertiseToOverlay))
			}
			if bgpConfig.TunnelName != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.tunnel_name", key), *bgpConfig.TunnelName)
			}
			if bgpConfig.BgpConfigType != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.type", key), *bgpConfig.BgpConfigType)
			}
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.via", key), bgpConfig.Via)
			if bgpConfig.VpnName != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.vpn_name", key), *bgpConfig.VpnName)
			}
			if bgpConfig.WanName != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.wan_name", key), *bgpConfig.WanName)
			}
		}
	}
	if s.DhcpdConfig != nil {
		checks.append(t, "TestCheckResourceAttrSet", "dhcpd_config.%")
		if s.DhcpdConfig.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "dhcpd_config.enabled", fmt.Sprintf("%t", *s.DhcpdConfig.Enabled))
		}
		if len(s.DhcpdConfig.Config) > 0 {
			checks.append(t, "TestCheckResourceAttr", "dhcpd_config.config.%", fmt.Sprintf("%d", len(s.DhcpdConfig.Config)))
			for configKey, config := range s.DhcpdConfig.Config {
				if len(config.DnsServers) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.dns_servers.#", configKey), fmt.Sprintf("%d", len(config.DnsServers)))
					for i, server := range config.DnsServers {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.dns_servers.%d", configKey, i), server)
					}
				}
				if len(config.DnsSuffix) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.dns_suffix.#", configKey), fmt.Sprintf("%d", len(config.DnsSuffix)))
					for i, suffix := range config.DnsSuffix {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.dns_suffix.%d", configKey, i), suffix)
					}
				}
				if len(config.FixedBindings) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.fixed_bindings.%%", configKey), fmt.Sprintf("%d", len(config.FixedBindings)))
					for bindingKey, binding := range config.FixedBindings {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.fixed_bindings.%s.ip", configKey, bindingKey), binding.Ip)
						if binding.Name != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.fixed_bindings.%s.name", configKey, bindingKey), *binding.Name)
						}
					}
				}
				if config.Gateway != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.gateway", configKey), *config.Gateway)
				}
				if config.IpEnd != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.ip_end", configKey), *config.IpEnd)
				}
				if config.IpEnd6 != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.ip_end6", configKey), *config.IpEnd6)
				}
				if config.IpStart != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.ip_start", configKey), *config.IpStart)
				}
				if config.IpStart6 != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.ip_start6", configKey), *config.IpStart6)
				}
				if config.LeaseTime != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.lease_time", configKey), fmt.Sprintf("%d", *config.LeaseTime))
				}
				if len(config.Options) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.options.%%", configKey), fmt.Sprintf("%d", len(config.Options)))
					for optionKey, option := range config.Options {
						if option.OptionsType != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.options.%s.type", configKey, optionKey), *option.OptionsType)
						}
						if option.Value != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.options.%s.value", configKey, optionKey), *option.Value)
						}
					}
				}
				if config.ServerIdOverride != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.server_id_override", configKey), fmt.Sprintf("%t", *config.ServerIdOverride))
				}
				if len(config.Servers) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.servers.#", configKey), fmt.Sprintf("%d", len(config.Servers)))
					for i, server := range config.Servers {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.servers.%d", configKey, i), server)
					}
				}
				if len(config.Servers6) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.servers6.#", configKey), fmt.Sprintf("%d", len(config.Servers6)))
					for i, server := range config.Servers6 {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.servers6.%d", configKey, i), server)
					}
				}
				if config.ConfigType != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.type", configKey), *config.ConfigType)
				}
				if config.Type6 != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.type6", configKey), *config.Type6)
				}
				if len(config.VendorEncapsulated) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.vendor_encapsulated.%%", configKey), fmt.Sprintf("%d", len(config.VendorEncapsulated)))
					for vendorKey, vendor := range config.VendorEncapsulated {
						if vendor.VendorEncapsulatedType != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.vendor_encapsulated.%s.type", configKey, vendorKey), *vendor.VendorEncapsulatedType)
						}
						if vendor.Value != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.vendor_encapsulated.%s.value", configKey, vendorKey), *vendor.Value)
						}
					}
				}
			}
		}
	}
	if s.DnsOverride != nil {
		checks.append(t, "TestCheckResourceAttr", "dns_override", fmt.Sprintf("%t", *s.DnsOverride))
	}
	if len(s.DnsServers) > 0 {
		checks.append(t, "TestCheckResourceAttr", "dns_servers.#", fmt.Sprintf("%d", len(s.DnsServers)))
		for i, server := range s.DnsServers {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dns_servers.%d", i), server)
		}
	}
	if len(s.DnsSuffix) > 0 {
		checks.append(t, "TestCheckResourceAttr", "dns_suffix.#", fmt.Sprintf("%d", len(s.DnsSuffix)))
		for i, suffix := range s.DnsSuffix {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dns_suffix.%d", i), suffix)
		}
	}
	if len(s.ExtraRoutes) > 0 {
		checks.append(t, "TestCheckResourceAttr", "extra_routes.%", fmt.Sprintf("%d", len(s.ExtraRoutes)))
		for key, route := range s.ExtraRoutes {
			if route.Discard != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("extra_routes.%s.discard", key), fmt.Sprintf("%t", *route.Discard))
			}
			if route.Metric != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("extra_routes.%s.metric", key), fmt.Sprintf("%d", *route.Metric))
			}
			if len(route.NextQualified) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("extra_routes.%s.next_qualified.%%", key), fmt.Sprintf("%d", len(route.NextQualified)))
				for qualifiedKey, qualified := range route.NextQualified {
					if qualified.Metric != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("extra_routes.%s.next_qualified.%s.metric", key, qualifiedKey), fmt.Sprintf("%d", *qualified.Metric))
					}
					if qualified.Preference != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("extra_routes.%s.next_qualified.%s.preference", key, qualifiedKey), fmt.Sprintf("%d", *qualified.Preference))
					}
				}
			}
			if route.NoResolve != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("extra_routes.%s.no_resolve", key), fmt.Sprintf("%t", *route.NoResolve))
			}
			if route.Preference != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("extra_routes.%s.preference", key), fmt.Sprintf("%d", *route.Preference))
			}
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("extra_routes.%s.via", key), route.Via)
		}
	}
	if len(s.IdpProfiles) > 0 {
		checks.append(t, "TestCheckResourceAttr", "idp_profiles.%", fmt.Sprintf("%d", len(s.IdpProfiles)))
		for key, idpProfile := range s.IdpProfiles {
			if idpProfile.BaseProfile != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("idp_profiles.%s.base_profile", key), *idpProfile.BaseProfile)
			}
			if idpProfile.Id != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("idp_profiles.%s.id", key), *idpProfile.Id)
			}
			if idpProfile.Name != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("idp_profiles.%s.name", key), *idpProfile.Name)
			}
			if idpProfile.OrgId != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("idp_profiles.%s.org_id", key), *idpProfile.OrgId)
			}
			if len(idpProfile.Overwrites) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("idp_profiles.%s.overwrites.#", key), fmt.Sprintf("%d", len(idpProfile.Overwrites)))
				for i, overwrite := range idpProfile.Overwrites {
					if overwrite.Action != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("idp_profiles.%s.overwrites.%d.action", key, i), *overwrite.Action)
					}
					if overwrite.IpdProfileOverwriteMatching != nil {
						if len(overwrite.IpdProfileOverwriteMatching.AttackName) > 0 {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("idp_profiles.%s.overwrites.%d.matching.attack_name.#", key, i), fmt.Sprintf("%d", len(overwrite.IpdProfileOverwriteMatching.AttackName)))
							for j, attackName := range overwrite.IpdProfileOverwriteMatching.AttackName {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("idp_profiles.%s.overwrites.%d.matching.attack_name.%d", key, i, j), attackName)
							}
						}
						if len(overwrite.IpdProfileOverwriteMatching.DstSubnet) > 0 {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("idp_profiles.%s.overwrites.%d.matching.dst_subnet.#", key, i), fmt.Sprintf("%d", len(overwrite.IpdProfileOverwriteMatching.DstSubnet)))
							for j, subnet := range overwrite.IpdProfileOverwriteMatching.DstSubnet {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("idp_profiles.%s.overwrites.%d.matching.dst_subnet.%d", key, i, j), subnet)
							}
						}
						if len(overwrite.IpdProfileOverwriteMatching.Severity) > 0 {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("idp_profiles.%s.overwrites.%d.matching.severity.#", key, i), fmt.Sprintf("%d", len(overwrite.IpdProfileOverwriteMatching.Severity)))
							for j, severity := range overwrite.IpdProfileOverwriteMatching.Severity {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("idp_profiles.%s.overwrites.%d.matching.severity.%d", key, i, j), severity)
							}
						}
					}
				}
			}
		}
	}
	if len(s.IpConfigs) > 0 {
		checks.append(t, "TestCheckResourceAttr", "ip_configs.%", fmt.Sprintf("%d", len(s.IpConfigs)))
		for key, ipConfig := range s.IpConfigs {
			if ipConfig.Ip != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ip_configs.%s.ip", key), *ipConfig.Ip)
			}
			if ipConfig.Ip6 != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ip_configs.%s.ip6", key), *ipConfig.Ip6)
			}
			if ipConfig.Netmask != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ip_configs.%s.netmask", key), *ipConfig.Netmask)
			}
			if ipConfig.Netmask6 != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ip_configs.%s.netmask6", key), *ipConfig.Netmask6)
			}
			if len(ipConfig.SecondaryIps) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ip_configs.%s.secondary_ips.#", key), fmt.Sprintf("%d", len(ipConfig.SecondaryIps)))
				for i, secondaryIp := range ipConfig.SecondaryIps {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ip_configs.%s.secondary_ips.%d", key, i), secondaryIp)
				}
			}
			if ipConfig.IpConfigsType != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ip_configs.%s.type", key), *ipConfig.IpConfigsType)
			}
			if ipConfig.Type6 != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ip_configs.%s.type6", key), *ipConfig.Type6)
			}
		}
	}
	if len(s.Networks) > 0 {
		checks.append(t, "TestCheckResourceAttr", "networks.#", fmt.Sprintf("%d", len(s.Networks)))
		for i, network := range s.Networks {
			if network.Gateway != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.gateway", i), *network.Gateway)
			}
			if network.Gateway6 != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.gateway6", i), *network.Gateway6)
			}
			if network.Isolation != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.isolation", i), fmt.Sprintf("%t", *network.Isolation))
			}
			if network.IsolationVlanId != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.isolation_vlan_id", i), *network.IsolationVlanId)
			}
			if network.Subnet != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.subnet", i), *network.Subnet)
			}
			if network.Subnet6 != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.subnet6", i), *network.Subnet6)
			}
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vlan_id", i), network.VlanId)
		}
	}
	if s.NtpOverride != nil {
		checks.append(t, "TestCheckResourceAttr", "ntp_override", fmt.Sprintf("%t", *s.NtpOverride))
	}
	if len(s.NtpServers) > 0 {
		checks.append(t, "TestCheckResourceAttr", "ntp_servers.#", fmt.Sprintf("%d", len(s.NtpServers)))
		for i, server := range s.NtpServers {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ntp_servers.%d", i), server)
		}
	}
	if s.OobIpConfig != nil {
		checks.append(t, "TestCheckResourceAttrSet", "oob_ip_config.%")
		if s.OobIpConfig.Gateway != nil {
			checks.append(t, "TestCheckResourceAttr", "oob_ip_config.gateway", *s.OobIpConfig.Gateway)
		}
		if s.OobIpConfig.Ip != nil {
			checks.append(t, "TestCheckResourceAttr", "oob_ip_config.ip", *s.OobIpConfig.Ip)
		}
		if s.OobIpConfig.Netmask != nil {
			checks.append(t, "TestCheckResourceAttr", "oob_ip_config.netmask", *s.OobIpConfig.Netmask)
		}
		if s.OobIpConfig.Network != nil {
			checks.append(t, "TestCheckResourceAttr", "oob_ip_config.network", *s.OobIpConfig.Network)
		}
		if s.OobIpConfig.OobIpConfigType != nil {
			checks.append(t, "TestCheckResourceAttr", "oob_ip_config.type", *s.OobIpConfig.OobIpConfigType)
		}
		if s.OobIpConfig.UseMgmtVrf != nil {
			checks.append(t, "TestCheckResourceAttr", "oob_ip_config.use_mgmt_vrf", fmt.Sprintf("%t", *s.OobIpConfig.UseMgmtVrf))
		}
		if s.OobIpConfig.UseMgmtVrfForHostOut != nil {
			checks.append(t, "TestCheckResourceAttr", "oob_ip_config.use_mgmt_vrf_for_host_out", fmt.Sprintf("%t", *s.OobIpConfig.UseMgmtVrfForHostOut))
		}
	}
	if s.RouterId != nil {
		checks.append(t, "TestCheckResourceAttr", "router_id", *s.RouterId)
	}
	if len(s.RoutingPolicies) > 0 {
		checks.append(t, "TestCheckResourceAttr", "routing_policies.%", fmt.Sprintf("%d", len(s.RoutingPolicies)))
		for key, policy := range s.RoutingPolicies {
			if len(policy.Terms) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.#", key), fmt.Sprintf("%d", len(policy.Terms)))
				for i, term := range policy.Terms {
					if term.Actions != nil {
						if term.Actions.Accept != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.actions.accept", key, i), fmt.Sprintf("%t", *term.Actions.Accept))
						}
						if len(term.Actions.AddCommunity) > 0 {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.actions.add_community.#", key, i), fmt.Sprintf("%d", len(term.Actions.AddCommunity)))
							for j, addItem := range term.Actions.AddCommunity {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.actions.add_community.%d", key, i, j), addItem)
							}
						}
						if len(term.Actions.Community) > 0 {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.actions.community.#", key, i), fmt.Sprintf("%d", len(term.Actions.Community)))
							for j, setItem := range term.Actions.Community {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.actions.community.%d", key, i, j), setItem)
							}
						}
						if term.Actions.LocalPreference != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.actions.local_preference", key, i), *term.Actions.LocalPreference)
						}
						if len(term.Actions.PrependAsPath) > 0 {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.actions.prepend_as_path.#", key, i), fmt.Sprintf("%d", len(term.Actions.PrependAsPath)))
							for j, prependAsPath := range term.Actions.PrependAsPath {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.actions.prepend_as_path.%d", key, i, j), prependAsPath)
							}
						}
					}
					if term.RoutingPolicyTermMatching != nil {
						if len(term.RoutingPolicyTermMatching.AsPath) > 0 {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.as_path.#", key, i), fmt.Sprintf("%d", len(term.RoutingPolicyTermMatching.AsPath)))
							for j, asPath := range term.RoutingPolicyTermMatching.AsPath {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.as_path.%d", key, i, j), asPath)
							}
						}
						if len(term.RoutingPolicyTermMatching.Community) > 0 {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.community.#", key, i), fmt.Sprintf("%d", len(term.RoutingPolicyTermMatching.Community)))
							for j, community := range term.RoutingPolicyTermMatching.Community {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.community.%d", key, i, j), community)
							}
						}
						if len(term.RoutingPolicyTermMatching.Network) > 0 {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.network.#", key, i), fmt.Sprintf("%d", len(term.RoutingPolicyTermMatching.Network)))
							for j, network := range term.RoutingPolicyTermMatching.Network {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.network.%d", key, i, j), network)
							}
						}
						if len(term.RoutingPolicyTermMatching.Prefix) > 0 {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.prefix.#", key, i), fmt.Sprintf("%d", len(term.RoutingPolicyTermMatching.Prefix)))
							for j, prefix := range term.RoutingPolicyTermMatching.Prefix {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.prefix.%d", key, i, j), prefix)
							}
						}
						if len(term.RoutingPolicyTermMatching.Protocol) > 0 {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.protocol.#", key, i), fmt.Sprintf("%d", len(term.RoutingPolicyTermMatching.Protocol)))
							for j, protocol := range term.RoutingPolicyTermMatching.Protocol {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.protocol.%d", key, i, j), protocol)
							}
						}
						if term.RoutingPolicyTermMatching.RouteExists != nil {
							if term.RoutingPolicyTermMatching.RouteExists.Route != nil {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.route_exists.route", key, i), *term.RoutingPolicyTermMatching.RouteExists.Route)
							}
							if term.RoutingPolicyTermMatching.RouteExists.VrfName != nil {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.route_exists.vrf_name", key, i), *term.RoutingPolicyTermMatching.RouteExists.VrfName)
							}
						}
						if len(term.RoutingPolicyTermMatching.VpnPath) > 0 {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.vpn_path.#", key, i), fmt.Sprintf("%d", len(term.RoutingPolicyTermMatching.VpnPath)))
							for j, vpnPath := range term.RoutingPolicyTermMatching.VpnPath {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.vpn_path.%d", key, i, j), vpnPath)
							}
						}
						if term.RoutingPolicyTermMatching.VpnPathSla != nil {
							if term.RoutingPolicyTermMatching.VpnPathSla.MaxJitter != nil {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.vpn_path_sla.max_jitter", key, i), fmt.Sprintf("%d", *term.RoutingPolicyTermMatching.VpnPathSla.MaxJitter))
							}
							if term.RoutingPolicyTermMatching.VpnPathSla.MaxLatency != nil {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.vpn_path_sla.max_latency", key, i), fmt.Sprintf("%d", *term.RoutingPolicyTermMatching.VpnPathSla.MaxLatency))
							}
							if term.RoutingPolicyTermMatching.VpnPathSla.MaxLoss != nil {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.vpn_path_sla.max_loss", key, i), fmt.Sprintf("%d", *term.RoutingPolicyTermMatching.VpnPathSla.MaxLoss))
							}
						}
					}
				}
			}
		}
	}
	if len(s.ServicePolicies) > 0 {
		checks.append(t, "TestCheckResourceAttr", "service_policies.#", fmt.Sprintf("%d", len(s.ServicePolicies)))
		for i, servicePolicy := range s.ServicePolicies {
			if servicePolicy.Action != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.action", i), *servicePolicy.Action)
			}
			if servicePolicy.Idp != nil {
				if servicePolicy.Idp.AlertOnly != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.idp.alert_only", i), fmt.Sprintf("%t", *servicePolicy.Idp.AlertOnly))
				}
				if servicePolicy.Idp.Enabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.idp.enabled", i), fmt.Sprintf("%t", *servicePolicy.Idp.Enabled))
				}
				if servicePolicy.Idp.IdpprofileId != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.idp.idpprofile_id", i), *servicePolicy.Idp.IdpprofileId)
				}
				if servicePolicy.Idp.Profile != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.idp.profile", i), *servicePolicy.Idp.Profile)
				}
			}
			if servicePolicy.Name != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.name", i), *servicePolicy.Name)
			}
			if servicePolicy.PathPreference != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.path_preference", i), *servicePolicy.PathPreference)
			}
			if len(servicePolicy.Services) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.services.#", i), fmt.Sprintf("%d", len(servicePolicy.Services)))
				for j, service := range servicePolicy.Services {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.services.%d", i, j), service)
				}
			}
			if len(servicePolicy.Tenants) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.tenants.#", i), fmt.Sprintf("%d", len(servicePolicy.Tenants)))
				for j, tenant := range servicePolicy.Tenants {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.tenants.%d", i, j), tenant)
				}
			}
		}
	}
	if s.Type != nil {
		checks.append(t, "TestCheckResourceAttr", "type", *s.Type)
	}
	if len(s.TunnelConfigs) > 0 {
		checks.append(t, "TestCheckResourceAttr", "tunnel_configs.%", fmt.Sprintf("%d", len(s.TunnelConfigs)))
		for key, tunnelConfig := range s.TunnelConfigs {
			if tunnelConfig.AutoProvision != nil {
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("tunnel_configs.%s.auto_provision", key))
			}
			if tunnelConfig.IkeLifetime != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.ike_lifetime", key), fmt.Sprintf("%d", *tunnelConfig.IkeLifetime))
			}
			if tunnelConfig.IkeMode != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.ike_mode", key), *tunnelConfig.IkeMode)
			}
			if len(tunnelConfig.IkeProposals) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.ike_proposals.#", key), fmt.Sprintf("%d", len(tunnelConfig.IkeProposals)))
				for i, proposal := range tunnelConfig.IkeProposals {
					if proposal.AuthAlgo != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.ike_proposals.%d.auth_algo", key, i), *proposal.AuthAlgo)
					}
					if proposal.DhGroup != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.ike_proposals.%d.dh_group", key, i), *proposal.DhGroup)
					}
					if proposal.EncAlgo != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.ike_proposals.%d.enc_algo", key, i), *proposal.EncAlgo)
					}
				}
			}
			if tunnelConfig.IpsecLifetime != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.ipsec_lifetime", key), fmt.Sprintf("%d", *tunnelConfig.IpsecLifetime))
			}
			if len(tunnelConfig.IpsecProposals) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.ipsec_proposals.#", key), fmt.Sprintf("%d", len(tunnelConfig.IpsecProposals)))
				for i, proposal := range tunnelConfig.IpsecProposals {
					if proposal.AuthAlgo != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.ipsec_proposals.%d.auth_algo", key, i), *proposal.AuthAlgo)
					}
					if proposal.DhGroup != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.ipsec_proposals.%d.dh_group", key, i), *proposal.DhGroup)
					}
					if proposal.EncAlgo != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.ipsec_proposals.%d.enc_algo", key, i), *proposal.EncAlgo)
					}
				}
			}
			if tunnelConfig.LocalId != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.local_id", key), *tunnelConfig.LocalId)
			}
			if tunnelConfig.Mode != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.mode", key), *tunnelConfig.Mode)
			}
			if len(tunnelConfig.Networks) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.networks.#", key), fmt.Sprintf("%d", len(tunnelConfig.Networks)))
				for i, network := range tunnelConfig.Networks {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.networks.%d", key, i), network)
				}
			}
			if tunnelConfig.Psk != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.psk", key), *tunnelConfig.Psk)
			}
			if tunnelConfig.Probe != nil {
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("tunnel_configs.%s.probe.%%", key))
				if tunnelConfig.Probe.Interval != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.probe.interval", key), fmt.Sprintf("%d", *tunnelConfig.Probe.Interval))
				}
				if tunnelConfig.Probe.Threshold != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.probe.threshold", key), fmt.Sprintf("%d", *tunnelConfig.Probe.Threshold))
				}
				if tunnelConfig.Probe.ProbeType != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.probe.type", key), *tunnelConfig.Probe.ProbeType)
				}
			}
			if tunnelConfig.Protocol != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.protocol", key), *tunnelConfig.Protocol)
			}
			if tunnelConfig.Provider != "" {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.provider", key), tunnelConfig.Provider)
			}
			if tunnelConfig.Psk != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.psk", key), *tunnelConfig.Psk)
			}

			if tunnelConfig.Version != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.version", key), *tunnelConfig.Version)
			}
		}
	}
	if s.TunnelProviderOptions != nil {
		checks.append(t, "TestCheckResourceAttrSet", "tunnel_provider_options.%")
		if s.TunnelProviderOptions.Jse != nil {
			if s.TunnelProviderOptions.Jse.NumUsers != nil {
				checks.append(t, "TestCheckResourceAttr", "tunnel_provider_options.jse.num_users", fmt.Sprintf("%d", *s.TunnelProviderOptions.Jse.NumUsers))
			}
			if s.TunnelProviderOptions.Jse.OrgName != nil {
				checks.append(t, "TestCheckResourceAttr", "tunnel_provider_options.jse.org_name", *s.TunnelProviderOptions.Jse.OrgName)
			}
		}
		if s.TunnelProviderOptions.Prisma != nil {
			if s.TunnelProviderOptions.Prisma.ServiceAccountName != nil {
				checks.append(t, "TestCheckResourceAttr", "tunnel_provider_options.prisma.service_account_name", *s.TunnelProviderOptions.Prisma.ServiceAccountName)
			}
		}
		if s.TunnelProviderOptions.Zscaler != nil {
			if s.TunnelProviderOptions.Zscaler.AupBlockInternetUntilAccepted != nil {
				checks.append(t, "TestCheckResourceAttr", "tunnel_provider_options.zscaler.aup_block_internet_until_accepted", fmt.Sprintf("%t", *s.TunnelProviderOptions.Zscaler.AupBlockInternetUntilAccepted))
			}
			if s.TunnelProviderOptions.Zscaler.AupEnabled != nil {
				checks.append(t, "TestCheckResourceAttr", "tunnel_provider_options.zscaler.aup_enabled", fmt.Sprintf("%t", *s.TunnelProviderOptions.Zscaler.AupEnabled))
			}
			if s.TunnelProviderOptions.Zscaler.AupForceSslInspection != nil {
				checks.append(t, "TestCheckResourceAttr", "tunnel_provider_options.zscaler.aup_force_ssl_inspection", fmt.Sprintf("%t", *s.TunnelProviderOptions.Zscaler.AupForceSslInspection))
			}
			if s.TunnelProviderOptions.Zscaler.AupTimeoutInDays != nil {
				checks.append(t, "TestCheckResourceAttr", "tunnel_provider_options.zscaler.aup_timeout_in_days", fmt.Sprintf("%d", *s.TunnelProviderOptions.Zscaler.AupTimeoutInDays))
			}
			if s.TunnelProviderOptions.Zscaler.AuthRequired != nil {
				checks.append(t, "TestCheckResourceAttr", "tunnel_provider_options.zscaler.auth_required", fmt.Sprintf("%t", *s.TunnelProviderOptions.Zscaler.AuthRequired))
			}
			if s.TunnelProviderOptions.Zscaler.CautionEnabled != nil {
				checks.append(t, "TestCheckResourceAttr", "tunnel_provider_options.zscaler.caution_enabled", fmt.Sprintf("%t", *s.TunnelProviderOptions.Zscaler.CautionEnabled))
			}
			if s.TunnelProviderOptions.Zscaler.DnBandwidth != nil {
				checks.append(t, "TestCheckResourceAttr", "tunnel_provider_options.zscaler.dn_bandwidth", fmt.Sprintf("%f", *s.TunnelProviderOptions.Zscaler.DnBandwidth))
			}
			if s.TunnelProviderOptions.Zscaler.IdleTimeInMinutes != nil {
				checks.append(t, "TestCheckResourceAttr", "tunnel_provider_options.zscaler.idle_time_in_minutes", fmt.Sprintf("%d", *s.TunnelProviderOptions.Zscaler.IdleTimeInMinutes))
			}
			if s.TunnelProviderOptions.Zscaler.OfwEnabled != nil {
				checks.append(t, "TestCheckResourceAttr", "tunnel_provider_options.zscaler.ofw_enabled", fmt.Sprintf("%t", *s.TunnelProviderOptions.Zscaler.OfwEnabled))
			}
			if len(s.TunnelProviderOptions.Zscaler.SubLocations) > 0 {
				checks.append(t, "TestCheckResourceAttr", "tunnel_provider_options.zscaler.sub_locations.#", fmt.Sprintf("%d", len(s.TunnelProviderOptions.Zscaler.SubLocations)))
				for i, subLocation := range s.TunnelProviderOptions.Zscaler.SubLocations {
					if subLocation.AupBlockInternetUntilAccepted != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_provider_options.zscaler.sub_locations.%d.aup_block_internet_until_accepted", i), fmt.Sprintf("%t", *subLocation.AupBlockInternetUntilAccepted))
					}
					if subLocation.AupEnabled != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_provider_options.zscaler.sub_locations.%d.aup_enabled", i), fmt.Sprintf("%t", *subLocation.AupEnabled))
					}
					if subLocation.AupForceSslInspection != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_provider_options.zscaler.sub_locations.%d.aup_force_ssl_inspection", i), fmt.Sprintf("%t", *subLocation.AupForceSslInspection))
					}
					if subLocation.AupTimeoutInDays != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_provider_options.zscaler.sub_locations.%d.aup_timeout_in_days", i), fmt.Sprintf("%d", *subLocation.AupTimeoutInDays))
					}
					if subLocation.AuthRequired != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_provider_options.zscaler.sub_locations.%d.auth_required", i), fmt.Sprintf("%t", *subLocation.AuthRequired))
					}
					if subLocation.CautionEnabled != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_provider_options.zscaler.sub_locations.%d.caution_enabled", i), fmt.Sprintf("%t", *subLocation.CautionEnabled))
					}
					if subLocation.DnBandwidth != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_provider_options.zscaler.sub_locations.%d.dn_bandwidth", i), fmt.Sprintf("%f", *subLocation.DnBandwidth))
					}
					if subLocation.IdleTimeInMinutes != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_provider_options.zscaler.sub_locations.%d.idle_time_in_minutes", i), fmt.Sprintf("%d", *subLocation.IdleTimeInMinutes))
					}
					if subLocation.Name != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_provider_options.zscaler.sub_locations.%d.name", i), *subLocation.Name)
					}
					if subLocation.OfwEnabled != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_provider_options.zscaler.sub_locations.%d.ofw_enabled", i), fmt.Sprintf("%t", *subLocation.OfwEnabled))
					}
					if subLocation.SurrogateIp != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_provider_options.zscaler.sub_locations.%d.surrogate_ip", i), fmt.Sprintf("%t", *subLocation.SurrogateIp))
					}
					if subLocation.SurrogateIpEnforcedForKnownBrowsers != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_provider_options.zscaler.sub_locations.%d.surrogate_ip_enforced_for_known_browsers", i), fmt.Sprintf("%t", *subLocation.SurrogateIpEnforcedForKnownBrowsers))
					}
					if subLocation.SurrogateRefreshTimeInMinutes != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_provider_options.zscaler.sub_locations.%d.surrogate_refresh_time_in_minutes", i), fmt.Sprintf("%d", *subLocation.SurrogateRefreshTimeInMinutes))
					}
					if subLocation.UpBandwidth != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_provider_options.zscaler.sub_locations.%d.up_bandwidth", i), fmt.Sprintf("%f", *subLocation.UpBandwidth))
					}
				}
			}
			if s.TunnelProviderOptions.Zscaler.SurrogateIp != nil {
				checks.append(t, "TestCheckResourceAttr", "tunnel_provider_options.zscaler.surrogate_ip", fmt.Sprintf("%t", *s.TunnelProviderOptions.Zscaler.SurrogateIp))
			}
			if s.TunnelProviderOptions.Zscaler.SurrogateIpEnforcedForKnownBrowsers != nil {
				checks.append(t, "TestCheckResourceAttr", "tunnel_provider_options.zscaler.surrogate_ip_enforced_for_known_browsers", fmt.Sprintf("%t", *s.TunnelProviderOptions.Zscaler.SurrogateIpEnforcedForKnownBrowsers))
			}
			if s.TunnelProviderOptions.Zscaler.SurrogateRefreshTimeInMinutes != nil {
				checks.append(t, "TestCheckResourceAttr", "tunnel_provider_options.zscaler.surrogate_refresh_time_in_minutes", fmt.Sprintf("%d", *s.TunnelProviderOptions.Zscaler.SurrogateRefreshTimeInMinutes))
			}
			if s.TunnelProviderOptions.Zscaler.UpBandwidth != nil {
				checks.append(t, "TestCheckResourceAttr", "tunnel_provider_options.zscaler.up_bandwidth", fmt.Sprintf("%f", *s.TunnelProviderOptions.Zscaler.UpBandwidth))
			}
			if s.TunnelProviderOptions.Zscaler.XffForwardEnabled != nil {
				checks.append(t, "TestCheckResourceAttr", "tunnel_provider_options.zscaler.xff_forward_enabled", fmt.Sprintf("%t", *s.TunnelProviderOptions.Zscaler.XffForwardEnabled))
			}
		}
	}
	if s.VrfConfig != nil {
		checks.append(t, "TestCheckResourceAttrSet", "vrf_config.%")
		if s.VrfConfig.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "vrf_config.enabled", fmt.Sprintf("%t", *s.VrfConfig.Enabled))
		}
	}
	if len(s.VrfInstances) > 0 {
		checks.append(t, "TestCheckResourceAttr", "vrf_instances.%", fmt.Sprintf("%d", len(s.VrfInstances)))
		for key, vrfInstance := range s.VrfInstances {
			if vrfInstance.EvpnAutoLoopbackSubnet != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vrf_instances.%s.evpn_auto_loopback_subnet", key), *vrfInstance.EvpnAutoLoopbackSubnet)
			}
			if vrfInstance.EvpnAutoLoopbackSubnet6 != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vrf_instances.%s.evpn_auto_loopback_subnet6", key), *vrfInstance.EvpnAutoLoopbackSubnet6)
			}
			if len(vrfInstance.Networks) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vrf_instances.%s.networks.#", key), fmt.Sprintf("%d", len(vrfInstance.Networks)))
				for i, network := range vrfInstance.Networks {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vrf_instances.%s.networks.%d", key, i), network)
				}
			}
			if len(vrfInstance.VrfExtraRoutes) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vrf_instances.%s.extra_routes.%%", key), fmt.Sprintf("%d", len(vrfInstance.VrfExtraRoutes)))
				for routeKey, route := range vrfInstance.VrfExtraRoutes {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vrf_instances.%s.extra_routes.%s.via", key, routeKey), route.Via)
				}
			}
			if len(vrfInstance.VrfExtraRoutes6) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vrf_instances.%s.extra_routes6.%%", key), fmt.Sprintf("%d", len(vrfInstance.VrfExtraRoutes6)))
				for routeKey, route6 := range vrfInstance.VrfExtraRoutes6 {
					if route6.Via != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vrf_instances.%s.extra_routes6.%s.via", key, routeKey), *route6.Via)
					}
				}
			}
		}
	}
	if len(s.PathPreferences) > 0 {
		checks.append(t, "TestCheckResourceAttr", "path_preferences.%", fmt.Sprintf("%d", len(s.PathPreferences)))
		for key, pathPref := range s.PathPreferences {
			if len(pathPref.Paths) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("path_preferences.%s.paths.#", key), fmt.Sprintf("%d", len(pathPref.Paths)))
				for i, path := range pathPref.Paths {
					if path.Cost != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("path_preferences.%s.paths.%d.cost", key, i), fmt.Sprintf("%d", *path.Cost))
					}
					if path.Disabled != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("path_preferences.%s.paths.%d.disabled", key, i), fmt.Sprintf("%t", *path.Disabled))
					}
					if path.GatewayIp != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("path_preferences.%s.paths.%d.gateway_ip", key, i), *path.GatewayIp)
					}
					if path.InternetAccess != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("path_preferences.%s.paths.%d.internet_access", key, i), fmt.Sprintf("%t", *path.InternetAccess))
					}
					if path.Name != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("path_preferences.%s.paths.%d.name", key, i), *path.Name)
					}
					if len(path.Networks) > 0 {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("path_preferences.%s.paths.%d.networks.#", key, i), fmt.Sprintf("%d", len(path.Networks)))
						for j, network := range path.Networks {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("path_preferences.%s.paths.%d.networks.%d", key, i, j), network)
						}
					}
					if len(path.TargetIps) > 0 {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("path_preferences.%s.paths.%d.target_ips.#", key, i), fmt.Sprintf("%d", len(path.TargetIps)))
						for j, targetIp := range path.TargetIps {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("path_preferences.%s.paths.%d.target_ips.%d", key, i, j), targetIp)
						}
					}
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("path_preferences.%s.paths.%d.type", key, i), path.PathsType)
					if path.WanName != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("path_preferences.%s.paths.%d.wan_name", key, i), *path.WanName)
					}
				}
			}
			if pathPref.Strategy != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("path_preferences.%s.strategy", key), *pathPref.Strategy)
			}
		}
	}
	if len(s.PortConfig) > 0 {
		checks.append(t, "TestCheckResourceAttr", "port_config.%", fmt.Sprintf("%d", len(s.PortConfig)))
		for key, portConfig := range s.PortConfig {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.usage", key), portConfig.Usage)
		}
	}

	return checks
}
