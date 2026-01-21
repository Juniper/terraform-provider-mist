package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_gatewaytemplate"
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

	resourceType := "org_gatewaytemplate"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_org_gatewaytemplate.OrgGatewaytemplateResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {

			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				config := step.config

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				configStr := Render(resourceType, tName, string(f.Bytes()))

				checks := config.testChecks(t, resourceType, tName, tracker)
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
	if tracker != nil {
		tracker.FieldCoverageReport(t)
	}
}

func (s *OrgGatewaytemplateModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + tName)
	checks.SetTracker(tracker)
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
						if binding.Ip != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.fixed_bindings.%s.ip", configKey, bindingKey), *binding.Ip)
						}
						if binding.Name != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.fixed_bindings.%s.name", configKey, bindingKey), *binding.Name)
						}
					}
				}
				if config.Gateway != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.gateway", configKey), *config.Gateway)
				}
				if config.IpEnd4 != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.ip_end", configKey), *config.IpEnd4)
				}
				if config.Ip6End != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.ip6_end", configKey), *config.Ip6End)
				}
				if config.IpStart4 != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.ip_start", configKey), *config.IpStart4)
				}
				if config.Ip6Start != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.ip6_start", configKey), *config.Ip6Start)
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
				if len(config.Servers4) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.servers.#", configKey), fmt.Sprintf("%d", len(config.Servers4)))
					for i, server := range config.Servers4 {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.servers.%d", configKey, i), server)
					}
				}
				if len(config.Serversv6) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.serversv6.#", configKey), fmt.Sprintf("%d", len(config.Serversv6)))
					for i, server := range config.Serversv6 {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.serversv6.%d", configKey, i), server)
					}
				}
				if config.Type4 != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.type", configKey), *config.Type4)
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
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("extra_routes.%s.via", key), route.Via)
		}
	}
	if len(s.ExtraRoutes6) > 0 {
		checks.append(t, "TestCheckResourceAttr", "extra_routes6.%", fmt.Sprintf("%d", len(s.ExtraRoutes6)))
		for key, route := range s.ExtraRoutes6 {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("extra_routes6.%s.via", key), route.Via)
		}
	}
	if len(s.IdpProfiles) > 0 {
		checks.append(t, "TestCheckResourceAttr", "idp_profiles.%", fmt.Sprintf("%d", len(s.IdpProfiles)))
		for key, idpProfile := range s.IdpProfiles {
			if idpProfile.BaseProfile != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("idp_profiles.%s.base_profile", key), *idpProfile.BaseProfile)
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
			// Required attributes
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.subnet", i), network.Subnet)

			// Optional attributes
			if network.DisallowMistServices != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.disallow_mist_services", i), fmt.Sprintf("%t", *network.DisallowMistServices))
			}
			if network.Gateway != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.gateway", i), *network.Gateway)
			}
			if network.Gateway6 != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.gateway6", i), *network.Gateway6)
			}
			if network.InternalAccess != nil {
				if network.InternalAccess.Enabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.internal_access.enabled", i), fmt.Sprintf("%t", *network.InternalAccess.Enabled))
				}
			}
			if network.InternetAccess != nil {
				if network.InternetAccess.CreateSimpleServicePolicy != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.internet_access.create_simple_service_policy", i), fmt.Sprintf("%t", *network.InternetAccess.CreateSimpleServicePolicy))
				}
				if network.InternetAccess.Enabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.internet_access.enabled", i), fmt.Sprintf("%t", *network.InternetAccess.Enabled))
				}
				if network.InternetAccess.Restricted != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.internet_access.restricted", i), fmt.Sprintf("%t", *network.InternetAccess.Restricted))
				}
				if len(network.InternetAccess.InternetAccessDestinationNat) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.internet_access.destination_nat.%%", i), fmt.Sprintf("%d", len(network.InternetAccess.InternetAccessDestinationNat)))
					for key, destNat := range network.InternetAccess.InternetAccessDestinationNat {
						if destNat.InternalIp != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.internet_access.destination_nat.%s.internal_ip", i, key), *destNat.InternalIp)
						}
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.internet_access.destination_nat.%s.name", i, key), destNat.Name)
						if destNat.Port != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.internet_access.destination_nat.%s.port", i, key), *destNat.Port)
						}
						if destNat.WanName != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.internet_access.destination_nat.%s.wan_name", i, key), *destNat.WanName)
						}
					}
				}
				if len(network.InternetAccess.InternetAccessStaticNat) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.internet_access.static_nat.%%", i), fmt.Sprintf("%d", len(network.InternetAccess.InternetAccessStaticNat)))
					for key, staticNat := range network.InternetAccess.InternetAccessStaticNat {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.internet_access.static_nat.%s.internal_ip", i, key), staticNat.InternalIp)
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.internet_access.static_nat.%s.name", i, key), staticNat.Name)
						if staticNat.WanName != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.internet_access.static_nat.%s.wan_name", i, key), *staticNat.WanName)
						}
					}
				}
			}
			if network.Isolation != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.isolation", i), fmt.Sprintf("%t", *network.Isolation))
			}
			if network.Multicast != nil {
				if network.Multicast.DisableIgmp != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.multicast.disable_igmp", i), fmt.Sprintf("%t", *network.Multicast.DisableIgmp))
				}
				if network.Multicast.Enabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.multicast.enabled", i), fmt.Sprintf("%t", *network.Multicast.Enabled))
				}
				if len(network.Multicast.Groups) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.multicast.groups.%%", i), fmt.Sprintf("%d", len(network.Multicast.Groups)))
					for key, group := range network.Multicast.Groups {
						if group.RpIp != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.multicast.groups.%s.rp_ip", i, key), *group.RpIp)
						}
					}
				}
			}
			if network.Name != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.name", i), *network.Name)
			}
			if len(network.RoutedForNetworks) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.routed_for_networks.#", i), fmt.Sprintf("%d", len(network.RoutedForNetworks)))
				for j, routedNetwork := range network.RoutedForNetworks {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.routed_for_networks.%d", i, j), routedNetwork)
				}
			}
			if network.Subnet6 != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.subnet6", i), *network.Subnet6)
			}
			if len(network.Tenants) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.tenants.%%", i), fmt.Sprintf("%d", len(network.Tenants)))
				for key, tenant := range network.Tenants {
					if len(tenant.Addresses) > 0 {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.tenants.%s.addresses.#", i, key), fmt.Sprintf("%d", len(tenant.Addresses)))
						for j, address := range tenant.Addresses {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.tenants.%s.addresses.%d", i, key, j), address)
						}
					}
				}
			}
			if network.VlanId != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vlan_id", i), *network.VlanId)
			}
			if len(network.VpnAccess) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%%", i), fmt.Sprintf("%d", len(network.VpnAccess)))
				for key, vpnAccess := range network.VpnAccess {
					if vpnAccess.AdvertisedSubnet != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.advertised_subnet", i, key), *vpnAccess.AdvertisedSubnet)
					}
					if vpnAccess.AllowPing != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.allow_ping", i, key), fmt.Sprintf("%t", *vpnAccess.AllowPing))
					}
					if vpnAccess.NatPool != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.nat_pool", i, key), *vpnAccess.NatPool)
					}
					if vpnAccess.NoReadvertiseToLanBgp != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.no_readvertise_to_lan_bgp", i, key), fmt.Sprintf("%t", *vpnAccess.NoReadvertiseToLanBgp))
					}
					if vpnAccess.NoReadvertiseToLanOspf != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.no_readvertise_to_lan_ospf", i, key), fmt.Sprintf("%t", *vpnAccess.NoReadvertiseToLanOspf))
					}
					if vpnAccess.NoReadvertiseToOverlay != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.no_readvertise_to_overlay", i, key), fmt.Sprintf("%t", *vpnAccess.NoReadvertiseToOverlay))
					}
					if len(vpnAccess.OtherVrfs) > 0 {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.other_vrfs.#", i, key), fmt.Sprintf("%d", len(vpnAccess.OtherVrfs)))
						for j, vrf := range vpnAccess.OtherVrfs {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.other_vrfs.%d", i, key, j), vrf)
						}
					}
					if vpnAccess.Routed != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.routed", i, key), fmt.Sprintf("%t", *vpnAccess.Routed))
					}
					if vpnAccess.SourceNat != nil {
						if vpnAccess.SourceNat.ExternalIp != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.source_nat.external_ip", i, key), *vpnAccess.SourceNat.ExternalIp)
						}
					}
					if vpnAccess.SummarizedSubnet != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.summarized_subnet", i, key), *vpnAccess.SummarizedSubnet)
					}
					if vpnAccess.SummarizedSubnetToLanBgp != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.summarized_subnet_to_lan_bgp", i, key), *vpnAccess.SummarizedSubnetToLanBgp)
					}
					if vpnAccess.SummarizedSubnetToLanOspf != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.summarized_subnet_to_lan_ospf", i, key), *vpnAccess.SummarizedSubnetToLanOspf)
					}
					if len(vpnAccess.VpnAccessDestinationNat) > 0 {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.destination_nat.%%", i, key), fmt.Sprintf("%d", len(vpnAccess.VpnAccessDestinationNat)))
						for destKey, destNat := range vpnAccess.VpnAccessDestinationNat {
							if destNat.InternalIp != nil {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.destination_nat.%s.internal_ip", i, key, destKey), *destNat.InternalIp)
							}
							if destNat.Name != nil {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.destination_nat.%s.name", i, key, destKey), *destNat.Name)
							}
							if destNat.Port != nil {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.destination_nat.%s.port", i, key, destKey), *destNat.Port)
							}
						}
					}
					if len(vpnAccess.VpnAccessStaticNat) > 0 {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.static_nat.%%", i, key), fmt.Sprintf("%d", len(vpnAccess.VpnAccessStaticNat)))
						for staticKey, staticNat := range vpnAccess.VpnAccessStaticNat {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.static_nat.%s.internal_ip", i, key, staticKey), staticNat.InternalIp)
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.static_nat.%s.name", i, key, staticKey), staticNat.Name)
						}
					}
				}
			}
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
		if s.OobIpConfig.Node1 != nil {
			if s.OobIpConfig.Node1.Gateway != nil {
				checks.append(t, "TestCheckResourceAttr", "oob_ip_config.node1.gateway", *s.OobIpConfig.Node1.Gateway)
			}
			if s.OobIpConfig.Node1.Ip != nil {
				checks.append(t, "TestCheckResourceAttr", "oob_ip_config.node1.ip", *s.OobIpConfig.Node1.Ip)
			}
			if s.OobIpConfig.Node1.Netmask != nil {
				checks.append(t, "TestCheckResourceAttr", "oob_ip_config.node1.netmask", *s.OobIpConfig.Node1.Netmask)
			}
			if s.OobIpConfig.Node1.Node1Type != nil {
				checks.append(t, "TestCheckResourceAttr", "oob_ip_config.node1.type", *s.OobIpConfig.Node1.Node1Type)
			}
			if s.OobIpConfig.Node1.UseMgmtVrf != nil {
				checks.append(t, "TestCheckResourceAttr", "oob_ip_config.node1.use_mgmt_vrf", fmt.Sprintf("%t", *s.OobIpConfig.Node1.UseMgmtVrf))
			}
			if s.OobIpConfig.Node1.UseMgmtVrfForHostOut != nil {
				checks.append(t, "TestCheckResourceAttr", "oob_ip_config.node1.use_mgmt_vrf_for_host_out", fmt.Sprintf("%t", *s.OobIpConfig.Node1.UseMgmtVrfForHostOut))
			}
			if s.OobIpConfig.Node1.VlanId != nil {
				checks.append(t, "TestCheckResourceAttr", "oob_ip_config.node1.vlan_id", *s.OobIpConfig.Node1.VlanId)
			}
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
		if s.OobIpConfig.VlanId != nil {
			checks.append(t, "TestCheckResourceAttr", "oob_ip_config.vlan_id", *s.OobIpConfig.VlanId)
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
			// Required field
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.usage", key), portConfig.Usage)

			// Optional fields
			if portConfig.AeDisableLacp != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ae_disable_lacp", key), fmt.Sprintf("%t", *portConfig.AeDisableLacp))
			}
			if portConfig.AeIdx != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ae_idx", key), *portConfig.AeIdx)
			}
			if portConfig.AeLacpForceUp != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ae_lacp_force_up", key), fmt.Sprintf("%t", *portConfig.AeLacpForceUp))
			}
			if portConfig.Aggregated != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.aggregated", key), fmt.Sprintf("%t", *portConfig.Aggregated))
			}
			if portConfig.Critical != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.critical", key), fmt.Sprintf("%t", *portConfig.Critical))
			}
			if portConfig.Description != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.description", key), *portConfig.Description)
			}
			if portConfig.DisableAutoneg != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.disable_autoneg", key), fmt.Sprintf("%t", *portConfig.DisableAutoneg))
			}
			if portConfig.Disabled != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.disabled", key), fmt.Sprintf("%t", *portConfig.Disabled))
			}
			if portConfig.DslType != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.dsl_type", key), *portConfig.DslType)
			}
			if portConfig.DslVci != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.dsl_vci", key), fmt.Sprintf("%d", *portConfig.DslVci))
			}
			if portConfig.DslVpi != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.dsl_vpi", key), fmt.Sprintf("%d", *portConfig.DslVpi))
			}
			if portConfig.Duplex != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.duplex", key), *portConfig.Duplex)
			}
			if portConfig.PortIpConfig != nil {
				if len(portConfig.PortIpConfig.Dns) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ip_config.dns.#", key), fmt.Sprintf("%d", len(portConfig.PortIpConfig.Dns)))
					for i, dns := range portConfig.PortIpConfig.Dns {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ip_config.dns.%d", key, i), dns)
					}
				}
				if len(portConfig.PortIpConfig.DnsSuffix) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ip_config.dns_suffix.#", key), fmt.Sprintf("%d", len(portConfig.PortIpConfig.DnsSuffix)))
					for i, suffix := range portConfig.PortIpConfig.DnsSuffix {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ip_config.dns_suffix.%d", key, i), suffix)
					}
				}
				if portConfig.PortIpConfig.Gateway != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ip_config.gateway", key), *portConfig.PortIpConfig.Gateway)
				}
				if portConfig.PortIpConfig.Gateway6 != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ip_config.gateway6", key), *portConfig.PortIpConfig.Gateway6)
				}
				if portConfig.PortIpConfig.Ip != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ip_config.ip", key), *portConfig.PortIpConfig.Ip)
				}
				if portConfig.PortIpConfig.Ip6 != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ip_config.ip6", key), *portConfig.PortIpConfig.Ip6)
				}
				if portConfig.PortIpConfig.Netmask != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ip_config.netmask", key), *portConfig.PortIpConfig.Netmask)
				}
				if portConfig.PortIpConfig.Netmask6 != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ip_config.netmask6", key), *portConfig.PortIpConfig.Netmask6)
				}
				if portConfig.PortIpConfig.Network != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ip_config.network", key), *portConfig.PortIpConfig.Network)
				}
				if portConfig.PortIpConfig.PortIpConfigType != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ip_config.type", key), *portConfig.PortIpConfig.PortIpConfigType)
				}
				if portConfig.PortIpConfig.Type6 != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ip_config.type6", key), *portConfig.PortIpConfig.Type6)
				}
			}
			if len(portConfig.Networks) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.networks.#", key), fmt.Sprintf("%d", len(portConfig.Networks)))
				for i, network := range portConfig.Networks {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.networks.%d", key, i), network)
				}
			}
			if portConfig.Mtu != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.mtu", key), fmt.Sprintf("%d", *portConfig.Mtu))
			}
			if portConfig.Name != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.name", key), *portConfig.Name)
			}
			if portConfig.OuterVlanId != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.outer_vlan_id", key), fmt.Sprintf("%d", *portConfig.OuterVlanId))
			}
			if portConfig.PoeDisabled != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.poe_disabled", key), fmt.Sprintf("%t", *portConfig.PoeDisabled))
			}
			if portConfig.PortNetwork != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.port_network", key), *portConfig.PortNetwork)
			}
			if portConfig.PreserveDscp != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.preserve_dscp", key), fmt.Sprintf("%t", *portConfig.PreserveDscp))
			}
			if portConfig.Redundant != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.redundant", key), fmt.Sprintf("%t", *portConfig.Redundant))
			}
			if portConfig.Speed != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.speed", key), *portConfig.Speed)
			}
			if portConfig.VlanId != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.vlan_id", key), *portConfig.VlanId)
			}
			if portConfig.WanType != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_type", key), *portConfig.WanType)
			}
			// LTE-related fields
			if portConfig.LteApn != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.lte_apn", key), *portConfig.LteApn)
			}
			if portConfig.LteAuth != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.lte_auth", key), *portConfig.LteAuth)
			}
			if portConfig.LteBackup != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.lte_backup", key), fmt.Sprintf("%t", *portConfig.LteBackup))
			}
			if portConfig.LtePassword != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.lte_password", key), *portConfig.LtePassword)
			}
			if portConfig.LteUsername != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.lte_username", key), *portConfig.LteUsername)
			}
			// RETH-related fields
			if portConfig.RedundantGroup != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.redundant_group", key), fmt.Sprintf("%d", *portConfig.RedundantGroup))
			}
			if portConfig.RethIdx != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.reth_idx", key), *portConfig.RethIdx)
			}
			if portConfig.RethNode != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.reth_node", key), *portConfig.RethNode)
			}
			if len(portConfig.RethNodes) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.reth_nodes.#", key), fmt.Sprintf("%d", len(portConfig.RethNodes)))
				for i, node := range portConfig.RethNodes {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.reth_nodes.%d", key, i), node)
				}
			}
			// SSR and other advanced fields
			if portConfig.SsrNoVirtualMac != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ssr_no_virtual_mac", key), fmt.Sprintf("%t", *portConfig.SsrNoVirtualMac))
			}
			if portConfig.SvrPortRange != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.svr_port_range", key), *portConfig.SvrPortRange)
			}
			// WAN-related fields
			if portConfig.WanArpPolicer != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_arp_policer", key), *portConfig.WanArpPolicer)
			}
			if portConfig.WanDisableSpeedtest != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_disable_speedtest", key), fmt.Sprintf("%t", *portConfig.WanDisableSpeedtest))
			}
			if portConfig.WanExtIp != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_ext_ip", key), *portConfig.WanExtIp)
			}
			if portConfig.WanExtIp6 != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_ext_ip6", key), *portConfig.WanExtIp6)
			}
			if len(portConfig.WanNetworks) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_networks.#", key), fmt.Sprintf("%d", len(portConfig.WanNetworks)))
				for i, network := range portConfig.WanNetworks {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_networks.%d", key, i), network)
				}
			}
			// TrafficShaping nested structure
			if portConfig.TrafficShaping != nil {
				if len(portConfig.TrafficShaping.ClassPercentages) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.traffic_shaping.class_percentages.#", key), fmt.Sprintf("%d", len(portConfig.TrafficShaping.ClassPercentages)))
					for i, percentage := range portConfig.TrafficShaping.ClassPercentages {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.traffic_shaping.class_percentages.%d", key, i), fmt.Sprintf("%d", percentage))
					}
				}
				if portConfig.TrafficShaping.Enabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.traffic_shaping.enabled", key), fmt.Sprintf("%t", *portConfig.TrafficShaping.Enabled))
				}
				if portConfig.TrafficShaping.MaxTxKbps != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.traffic_shaping.max_tx_kbps", key), fmt.Sprintf("%d", *portConfig.TrafficShaping.MaxTxKbps))
				}
			}
			// VpnPaths nested map structure
			if len(portConfig.VpnPaths) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.vpn_paths.%%", key), fmt.Sprintf("%d", len(portConfig.VpnPaths)))
				for vpnKey, vpnPath := range portConfig.VpnPaths {
					if vpnPath.BfdProfile != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.vpn_paths.%s.bfd_profile", key, vpnKey), *vpnPath.BfdProfile)
					}
					if vpnPath.BfdUseTunnelMode != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.vpn_paths.%s.bfd_use_tunnel_mode", key, vpnKey), fmt.Sprintf("%t", *vpnPath.BfdUseTunnelMode))
					}
					if vpnPath.Preference != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.vpn_paths.%s.preference", key, vpnKey), fmt.Sprintf("%d", *vpnPath.Preference))
					}
					if vpnPath.Role != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.vpn_paths.%s.role", key, vpnKey), *vpnPath.Role)
					}
					if vpnPath.TrafficShaping != nil {
						if len(vpnPath.TrafficShaping.ClassPercentages) > 0 {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.vpn_paths.%s.traffic_shaping.class_percentages.#", key, vpnKey), fmt.Sprintf("%d", len(vpnPath.TrafficShaping.ClassPercentages)))
							for i, percentage := range vpnPath.TrafficShaping.ClassPercentages {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.vpn_paths.%s.traffic_shaping.class_percentages.%d", key, vpnKey, i), fmt.Sprintf("%d", percentage))
							}
						}
						if vpnPath.TrafficShaping.Enabled != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.vpn_paths.%s.traffic_shaping.enabled", key, vpnKey), fmt.Sprintf("%t", *vpnPath.TrafficShaping.Enabled))
						}
						if vpnPath.TrafficShaping.MaxTxKbps != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.vpn_paths.%s.traffic_shaping.max_tx_kbps", key, vpnKey), fmt.Sprintf("%d", *vpnPath.TrafficShaping.MaxTxKbps))
						}
					}
				}
			}
			// WanExtraRoutes nested map structure
			if len(portConfig.WanExtraRoutes) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_extra_routes.%%", key), fmt.Sprintf("%d", len(portConfig.WanExtraRoutes)))
				for routeKey, route := range portConfig.WanExtraRoutes {
					if route.Via != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_extra_routes.%s.via", key, routeKey), *route.Via)
					}
				}
			}
			// WanExtraRoutes6 nested map structure
			if len(portConfig.WanExtraRoutes6) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_extra_routes6.%%", key), fmt.Sprintf("%d", len(portConfig.WanExtraRoutes6)))
				for routeKey, route := range portConfig.WanExtraRoutes6 {
					if route.Via != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_extra_routes6.%s.via", key, routeKey), *route.Via)
					}
				}
			}
			// WanProbeOverride nested structure
			if portConfig.WanProbeOverride != nil {
				if len(portConfig.WanProbeOverride.Ip6s) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_probe_override.ip6s.#", key), fmt.Sprintf("%d", len(portConfig.WanProbeOverride.Ip6s)))
					for i, ip6 := range portConfig.WanProbeOverride.Ip6s {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_probe_override.ip6s.%d", key, i), ip6)
					}
				}
				if len(portConfig.WanProbeOverride.Ips) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_probe_override.ips.#", key), fmt.Sprintf("%d", len(portConfig.WanProbeOverride.Ips)))
					for i, ip := range portConfig.WanProbeOverride.Ips {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_probe_override.ips.%d", key, i), ip)
					}
				}
				if portConfig.WanProbeOverride.ProbeProfile != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_probe_override.probe_profile", key), *portConfig.WanProbeOverride.ProbeProfile)
				}
			}
			// WanSourceNat nested structure
			if portConfig.WanSourceNat != nil {
				if portConfig.WanSourceNat.Disabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_source_nat.disabled", key), fmt.Sprintf("%t", *portConfig.WanSourceNat.Disabled))
				}
				if portConfig.WanSourceNat.Nat6Pool != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_source_nat.nat6_pool", key), *portConfig.WanSourceNat.Nat6Pool)
				}
				if portConfig.WanSourceNat.NatPool != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_source_nat.nat_pool", key), *portConfig.WanSourceNat.NatPool)
				}
			}
		}
	}
	if s.RouterId != nil {
		checks.append(t, "TestCheckResourceAttr", "router_id", *s.RouterId)
	}
	if len(s.ServicePolicies) > 0 {
		checks.append(t, "TestCheckResourceAttr", "service_policies.#", fmt.Sprintf("%d", len(s.ServicePolicies)))
		for i, policy := range s.ServicePolicies {
			if policy.Action != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.action", i), *policy.Action)
			}
			if policy.Idp != nil {
				if policy.Idp.Enabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.idp.enabled", i), fmt.Sprintf("%t", *policy.Idp.Enabled))
				}
				if policy.Idp.Profile != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.idp.profile", i), *policy.Idp.Profile)
				}
				if policy.Idp.AlertOnly != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.idp.alert_only", i), fmt.Sprintf("%t", *policy.Idp.AlertOnly))
				}
				if policy.Idp.IdpprofileId != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.idp.idpprofile_id", i), *policy.Idp.IdpprofileId)
				}
			}
			if policy.LocalRouting != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.local_routing", i), fmt.Sprintf("%t", *policy.LocalRouting))
			}
			if policy.Name != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.name", i), *policy.Name)
			}
			if policy.PathPreference != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.path_preference", i), *policy.PathPreference)
			}
			if len(policy.Services) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.services.#", i), fmt.Sprintf("%d", len(policy.Services)))
				for j, service := range policy.Services {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.services.%d", i, j), service)
				}
			}
			if len(policy.Tenants) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.tenants.#", i), fmt.Sprintf("%d", len(policy.Tenants)))
				for j, tenant := range policy.Tenants {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.tenants.%d", i, j), tenant)
				}
			}
			if policy.Skyatp != nil {
				if policy.Skyatp.DnsDgaDetection != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.skyatp.dns_dga_detection", i), *policy.Skyatp.DnsDgaDetection)
				}
				if policy.Skyatp.DnsTunnelDetection != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.skyatp.dns_tunnel_detection", i), *policy.Skyatp.DnsTunnelDetection)
				}
				if policy.Skyatp.HttpInspection != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.skyatp.http_inspection", i), *policy.Skyatp.HttpInspection)
				}
				if policy.Skyatp.IotDevicePolicy != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.skyatp.iot_device_policy", i), *policy.Skyatp.IotDevicePolicy)
				}
			}
			if policy.Syslog != nil {
				if policy.Syslog.Enabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.syslog.enabled", i), fmt.Sprintf("%t", *policy.Syslog.Enabled))
				}
				if len(policy.Syslog.ServerNames) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.syslog.server_names.#", i), fmt.Sprintf("%d", len(policy.Syslog.ServerNames)))
					for j, serverName := range policy.Syslog.ServerNames {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.syslog.server_names.%d", i, j), serverName)
					}
				}
			}
		}
	}
	if len(s.TunnelConfigs) > 0 {
		checks.append(t, "TestCheckResourceAttr", "tunnel_configs.%", fmt.Sprintf("%d", len(s.TunnelConfigs)))
		for key, tunnelConfig := range s.TunnelConfigs {
			if tunnelConfig.AutoProvision != nil {
				if tunnelConfig.AutoProvision.Enabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.auto_provision.enabled", key), fmt.Sprintf("%t", *tunnelConfig.AutoProvision.Enabled))
				}
				if tunnelConfig.AutoProvision.Latlng != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.auto_provision.latlng.lat", key), fmt.Sprintf("%g", tunnelConfig.AutoProvision.Latlng.Lat))
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.auto_provision.latlng.lng", key), fmt.Sprintf("%g", tunnelConfig.AutoProvision.Latlng.Lng))
				}
				if tunnelConfig.AutoProvision.AutoProvisionPrimary != nil {
					if len(tunnelConfig.AutoProvision.AutoProvisionPrimary.WanNames) > 0 {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.auto_provision.primary.wan_names.#", key), fmt.Sprintf("%d", len(tunnelConfig.AutoProvision.AutoProvisionPrimary.WanNames)))
						for i, wanName := range tunnelConfig.AutoProvision.AutoProvisionPrimary.WanNames {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.auto_provision.primary.wan_names.%d", key, i), wanName)
						}
					}
					if len(tunnelConfig.AutoProvision.AutoProvisionPrimary.ProbeIps) > 0 {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.auto_provision.primary.probe_ips.#", key), fmt.Sprintf("%d", len(tunnelConfig.AutoProvision.AutoProvisionPrimary.ProbeIps)))
						for i, probeIp := range tunnelConfig.AutoProvision.AutoProvisionPrimary.ProbeIps {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.auto_provision.primary.probe_ips.%d", key, i), probeIp)
						}
					}
				}
				if tunnelConfig.AutoProvision.AutoProvisionSecondary != nil {
					if len(tunnelConfig.AutoProvision.AutoProvisionSecondary.WanNames) > 0 {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.auto_provision.secondary.wan_names.#", key), fmt.Sprintf("%d", len(tunnelConfig.AutoProvision.AutoProvisionSecondary.WanNames)))
						for i, wanName := range tunnelConfig.AutoProvision.AutoProvisionSecondary.WanNames {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.auto_provision.secondary.wan_names.%d", key, i), wanName)
						}
					}
					if len(tunnelConfig.AutoProvision.AutoProvisionSecondary.ProbeIps) > 0 {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.auto_provision.secondary.probe_ips.#", key), fmt.Sprintf("%d", len(tunnelConfig.AutoProvision.AutoProvisionSecondary.ProbeIps)))
						for i, probeIp := range tunnelConfig.AutoProvision.AutoProvisionSecondary.ProbeIps {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.auto_provision.secondary.probe_ips.%d", key, i), probeIp)
						}
					}
				}
				if tunnelConfig.AutoProvision.Provider != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.auto_provision.provider", key), *tunnelConfig.AutoProvision.Provider)
				}
				if tunnelConfig.AutoProvision.Region != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.auto_provision.region", key), *tunnelConfig.AutoProvision.Region)
				}
				if tunnelConfig.AutoProvision.ServiceConnection != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.auto_provision.service_connection", key), *tunnelConfig.AutoProvision.ServiceConnection)
				}
			}
			if tunnelConfig.IkeProposals != nil {
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
			if tunnelConfig.IpsecProposals != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.ipsec_proposals.#", key), fmt.Sprintf("%d", len(tunnelConfig.IpsecProposals)))
				for i, proposal := range tunnelConfig.IpsecProposals {
					if proposal.AuthAlgo != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.ipsec_proposals.%d.auth_algo", key, i), *proposal.AuthAlgo)
					}
					if proposal.EncAlgo != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.ipsec_proposals.%d.enc_algo", key, i), *proposal.EncAlgo)
					}
					if proposal.DhGroup != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.ipsec_proposals.%d.dh_group", key, i), *proposal.DhGroup)
					}
				}
			}
			if tunnelConfig.LocalId != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.local_id", key), *tunnelConfig.LocalId)
			}
			if tunnelConfig.Mode != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.mode", key), *tunnelConfig.Mode)
			}
			if tunnelConfig.IkeLifetime != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.ike_lifetime", key), fmt.Sprintf("%d", *tunnelConfig.IkeLifetime))
			}
			if tunnelConfig.IkeMode != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.ike_mode", key), *tunnelConfig.IkeMode)
			}
			if tunnelConfig.IpsecLifetime != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.ipsec_lifetime", key), fmt.Sprintf("%d", *tunnelConfig.IpsecLifetime))
			}
			if tunnelConfig.Probe != nil {
				if tunnelConfig.Probe.Interval != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.probe.interval", key), fmt.Sprintf("%d", *tunnelConfig.Probe.Interval))
				}
				if tunnelConfig.Probe.Threshold != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.probe.threshold", key), fmt.Sprintf("%d", *tunnelConfig.Probe.Threshold))
				}
				if tunnelConfig.Probe.Timeout != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.probe.timeout", key), fmt.Sprintf("%d", *tunnelConfig.Probe.Timeout))
				}
				if tunnelConfig.Probe.ProbeType != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.probe.type", key), *tunnelConfig.Probe.ProbeType)
				}
			}
			if tunnelConfig.Protocol != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.protocol", key), *tunnelConfig.Protocol)
			}
			// Provider is required field (not pointer)
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.provider", key), tunnelConfig.Provider)
			if tunnelConfig.Psk != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.psk", key), *tunnelConfig.Psk)
			}
			if tunnelConfig.Version != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.version", key), *tunnelConfig.Version)
			}
			// Additional tunnel config fields that exist in the struct
			if len(tunnelConfig.LocalSubnets) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.local_subnets.#", key), fmt.Sprintf("%d", len(tunnelConfig.LocalSubnets)))
				for i, subnet := range tunnelConfig.LocalSubnets {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.local_subnets.%d", key, i), subnet)
				}
			}
			if len(tunnelConfig.RemoteSubnets) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.remote_subnets.#", key), fmt.Sprintf("%d", len(tunnelConfig.RemoteSubnets)))
				for i, subnet := range tunnelConfig.RemoteSubnets {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.remote_subnets.%d", key, i), subnet)
				}
			}
			if len(tunnelConfig.Networks) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.networks.#", key), fmt.Sprintf("%d", len(tunnelConfig.Networks)))
				for i, network := range tunnelConfig.Networks {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.networks.%d", key, i), network)
				}
			}
		}
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
	if s.Type != nil {
		checks.append(t, "TestCheckResourceAttr", "type", *s.Type)
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
				checks.append(t, "TestCheckResourceAttr", "tunnel_provider_options.zscaler.dn_bandwidth", fmt.Sprintf("%g", *s.TunnelProviderOptions.Zscaler.DnBandwidth))
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
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_provider_options.zscaler.sub_locations.%d.dn_bandwidth", i), fmt.Sprintf("%g", *subLocation.DnBandwidth))
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
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_provider_options.zscaler.sub_locations.%d.up_bandwidth", i), fmt.Sprintf("%g", *subLocation.UpBandwidth))
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
				checks.append(t, "TestCheckResourceAttr", "tunnel_provider_options.zscaler.up_bandwidth", fmt.Sprintf("%g", *s.TunnelProviderOptions.Zscaler.UpBandwidth))
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
			if len(vrfInstance.Networks) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vrf_instances.%s.networks.#", key), fmt.Sprintf("%d", len(vrfInstance.Networks)))
				for i, network := range vrfInstance.Networks {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vrf_instances.%s.networks.%d", key, i), network)
				}
			}
		}
	}

	return checks
}
