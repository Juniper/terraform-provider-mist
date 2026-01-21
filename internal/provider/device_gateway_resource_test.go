package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_device_gateway"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDeviceGatewayModel(t *testing.T) {
	type testStep struct {
		config DeviceGatewayModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: DeviceGatewayModel{
						DeviceId: "00000000-0000-0000-1000-f4bfa8426080",
						Name:     "test_device_gateway",
						SiteId:   "2c107c8e-2e06-404a-ba61-e25b5757ecea",
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/device_gateway_resource/device_gateway_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		var FixtureDeviceGatewayModel DeviceGatewayModel
		err = hcl.Decode(&FixtureDeviceGatewayModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureDeviceGatewayModel,
				},
			},
		}
	}

	resourceType := "device_gateway"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_device_gateway.DeviceGatewayResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Skip("Skipping device_gateway tests, as they require a real device.")
		t.Run(tName, func(t *testing.T) {
			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				// siteConfig, siteRef := GetSiteBaseConfig(GetTestOrgId())
				config := step.config

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				// f.Body().SetAttributeRaw("site_id", hclwrite.TokensForIdentifier(siteRef))
				combinedConfig := Render(resourceType, tName, string(f.Bytes()))

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
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps:                    steps,
			})
		})
	}
	if tracker != nil {
		tracker.FieldCoverageReport(t)
	}
}

func (s *DeviceGatewayModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + tName)
	checks.SetTracker(tracker)

	// Always present attributes
	checks.append(t, "TestCheckResourceAttrSet", "site_id")
	checks.append(t, "TestCheckResourceAttr", "name", s.Name)
	checks.append(t, "TestCheckResourceAttrSet", "device_id")

	// Computed-only attributes - check for presence
	checks.append(t, "TestCheckResourceAttrSet", "image1_url")
	checks.append(t, "TestCheckResourceAttrSet", "image2_url")
	checks.append(t, "TestCheckResourceAttrSet", "image3_url")
	checks.append(t, "TestCheckResourceAttrSet", "mac")
	checks.append(t, "TestCheckResourceAttrSet", "model")
	checks.append(t, "TestCheckResourceAttrSet", "org_id")
	checks.append(t, "TestCheckResourceAttrSet", "serial")

	// Optional/configurable attributes with conditional checks
	if len(s.AdditionalConfigCmds) > 0 {
		checks.append(t, "TestCheckResourceAttr", "additional_config_cmds.#", fmt.Sprintf("%d", len(s.AdditionalConfigCmds)))
	}

	if len(s.BgpConfig) > 0 {
		checks.append(t, "TestCheckResourceAttr", "bgp_config.%", fmt.Sprintf("%d", len(s.BgpConfig)))
		for k, v := range s.BgpConfig {
			if v.AuthKey != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.auth_key", k), *v.AuthKey)
			}
			if v.BfdMinimumInterval != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.bfd_minimum_interval", k), fmt.Sprintf("%d", *v.BfdMinimumInterval))
			}
			if v.BfdMultiplier != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.bfd_multiplier", k), fmt.Sprintf("%d", *v.BfdMultiplier))
			}
			if v.DisableBfd != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.disable_bfd", k), fmt.Sprintf("%t", *v.DisableBfd))
			}
			if v.Export != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.export", k), *v.Export)
			}
			if v.ExportPolicy != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.export_policy", k), *v.ExportPolicy)
			}
			if v.ExtendedV4Nexthop != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.extended_v4_nexthop", k), fmt.Sprintf("%t", *v.ExtendedV4Nexthop))
			}
			if v.GracefulRestartTime != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.graceful_restart_time", k), fmt.Sprintf("%d", *v.GracefulRestartTime))
			}
			if v.HoldTime != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.hold_time", k), fmt.Sprintf("%d", *v.HoldTime))
			}
			if v.Import != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.import", k), *v.Import)
			}
			if v.ImportPolicy != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.import_policy", k), *v.ImportPolicy)
			}
			if v.LocalAs != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.local_as", k), *v.LocalAs)
			}
			if v.NeighborAs != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.neighbor_as", k), *v.NeighborAs)
			}
			if len(v.Neighbors) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.neighbors.%%", k), fmt.Sprintf("%d", len(v.Neighbors)))
				// Add detailed neighbor checks
				for neighborKey, neighbor := range v.Neighbors {
					if neighbor.Disabled != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.neighbors.%s.disabled", k, neighborKey), fmt.Sprintf("%t", *neighbor.Disabled))
					}
					if neighbor.ExportPolicy != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.neighbors.%s.export_policy", k, neighborKey), *neighbor.ExportPolicy)
					}
					if neighbor.HoldTime != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.neighbors.%s.hold_time", k, neighborKey), fmt.Sprintf("%d", *neighbor.HoldTime))
					}
					if neighbor.ImportPolicy != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.neighbors.%s.import_policy", k, neighborKey), *neighbor.ImportPolicy)
					}
					if neighbor.MultihopTtl != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.neighbors.%s.multihop_ttl", k, neighborKey), fmt.Sprintf("%d", *neighbor.MultihopTtl))
					}
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.neighbors.%s.neighbor_as", k, neighborKey), neighbor.NeighborAs)

				}
			}
			if len(v.Networks) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.networks.#", k), fmt.Sprintf("%d", len(v.Networks)))
				for i, network := range v.Networks {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.networks.%d", k, i), network)
				}
			}
			if v.NoPrivateAs != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.no_private_as", k), fmt.Sprintf("%t", *v.NoPrivateAs))
			}
			if v.NoReadvertiseToOverlay != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.no_readvertise_to_overlay", k), fmt.Sprintf("%t", *v.NoReadvertiseToOverlay))
			}
			if v.TunnelName != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.tunnel_name", k), *v.TunnelName)
			}
			if v.BgpConfigType != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.type", k), *v.BgpConfigType)
			}

			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.via", k), v.Via)

			if v.VpnName != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.vpn_name", k), *v.VpnName)
			}
			if v.WanName != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("bgp_config.%s.wan_name", k), *v.WanName)
			}
		}
	}

	if s.UrlFilteringDenyMsg != nil {
		checks.append(t, "TestCheckResourceAttr", "url_filtering_deny_msg", *s.UrlFilteringDenyMsg)
	}

	if s.DhcpdConfig != nil {
		if s.DhcpdConfig.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "dhcpd_config.enabled", fmt.Sprintf("%t", *s.DhcpdConfig.Enabled))
		}
		if len(s.DhcpdConfig.Config) > 0 {
			checks.append(t, "TestCheckResourceAttr", "dhcpd_config.config.%", fmt.Sprintf("%d", len(s.DhcpdConfig.Config)))
			for configKey, config := range s.DhcpdConfig.Config {
				configPath := fmt.Sprintf("dhcpd_config.config.%s", configKey)

				// DNS servers
				if len(config.DnsServers) > 0 {
					checks.append(t, "TestCheckResourceAttr", configPath+".dns_servers.#", fmt.Sprintf("%d", len(config.DnsServers)))
					for i, server := range config.DnsServers {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.dns_servers.%d", configPath, i), server)
					}
				}

				// DNS suffix
				if len(config.DnsSuffix) > 0 {
					checks.append(t, "TestCheckResourceAttr", configPath+".dns_suffix.#", fmt.Sprintf("%d", len(config.DnsSuffix)))
					for i, suffix := range config.DnsSuffix {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.dns_suffix.%d", configPath, i), suffix)
					}
				}

				// Fixed bindings
				if len(config.FixedBindings) > 0 {
					checks.append(t, "TestCheckResourceAttr", configPath+".fixed_bindings.%", fmt.Sprintf("%d", len(config.FixedBindings)))
					for bindingKey, binding := range config.FixedBindings {
						bindingPath := fmt.Sprintf("%s.fixed_bindings.%s", configPath, bindingKey)
						if binding.Ip != nil {
							checks.append(t, "TestCheckResourceAttr", bindingPath+".ip", *binding.Ip)
						}
						if binding.Ip6 != nil {
							checks.append(t, "TestCheckResourceAttr", bindingPath+".ip6", *binding.Ip6)
						}
						if binding.Name != nil {
							checks.append(t, "TestCheckResourceAttr", bindingPath+".name", *binding.Name)
						}
					}
				}

				// Gateway and IP ranges
				if config.Gateway != nil {
					checks.append(t, "TestCheckResourceAttr", configPath+".gateway", *config.Gateway)
				}
				if config.IpEnd4 != nil {
					checks.append(t, "TestCheckResourceAttr", configPath+".ip_end", *config.IpEnd4)
				}
				if config.IpStart4 != nil {
					checks.append(t, "TestCheckResourceAttr", configPath+".ip_start", *config.IpStart4)
				}
				if config.Ip6End != nil {
					checks.append(t, "TestCheckResourceAttr", configPath+".ip6_end", *config.Ip6End)
				}
				if config.Ip6Start != nil {
					checks.append(t, "TestCheckResourceAttr", configPath+".ip6_start", *config.Ip6Start)
				}

				// Lease time
				if config.LeaseTime != nil {
					checks.append(t, "TestCheckResourceAttr", configPath+".lease_time", fmt.Sprintf("%d", *config.LeaseTime))
				}

				// Options
				if len(config.Options) > 0 {
					checks.append(t, "TestCheckResourceAttr", configPath+".options.%", fmt.Sprintf("%d", len(config.Options)))
					for optionKey, option := range config.Options {
						optionPath := fmt.Sprintf("%s.options.%s", configPath, optionKey)
						if option.OptionsType != nil {
							checks.append(t, "TestCheckResourceAttr", optionPath+".type", *option.OptionsType)
						}
						if option.Value != nil {
							checks.append(t, "TestCheckResourceAttr", optionPath+".value", *option.Value)
						}
					}
				}

				// Server settings
				if config.ServerIdOverride != nil {
					checks.append(t, "TestCheckResourceAttr", configPath+".server_id_override", fmt.Sprintf("%t", *config.ServerIdOverride))
				}
				if len(config.Servers4) > 0 {
					checks.append(t, "TestCheckResourceAttr", configPath+".servers.#", fmt.Sprintf("%d", len(config.Servers4)))
					for i, server := range config.Servers4 {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.servers.%d", configPath, i), server)
					}
				}
				if len(config.Serversv6) > 0 {
					checks.append(t, "TestCheckResourceAttr", configPath+".serversv6.#", fmt.Sprintf("%d", len(config.Serversv6)))
					for i, server := range config.Serversv6 {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.serversv6.%d", configPath, i), server)
					}
				}

				// Types
				if config.Type4 != nil {
					checks.append(t, "TestCheckResourceAttr", configPath+".type", *config.Type4)
				}
				if config.Type6 != nil {
					checks.append(t, "TestCheckResourceAttr", configPath+".type6", *config.Type6)
				}

				// Vendor encapsulated options
				if len(config.VendorEncapsulated) > 0 {
					checks.append(t, "TestCheckResourceAttr", configPath+".vendor_encapsulated.%", fmt.Sprintf("%d", len(config.VendorEncapsulated)))
					for vendorKey, vendorOption := range config.VendorEncapsulated {
						vendorPath := fmt.Sprintf("%s.vendor_encapsulated.%s", configPath, vendorKey)
						if vendorOption.VendorEncapsulatedType != nil {
							checks.append(t, "TestCheckResourceAttr", vendorPath+".type", *vendorOption.VendorEncapsulatedType)
						}
						if vendorOption.Value != nil {
							checks.append(t, "TestCheckResourceAttr", vendorPath+".value", *vendorOption.Value)
						}
					}
				}
			}
		}
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
		for k, v := range s.ExtraRoutes {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("extra_routes.%s.via", k), v.Via)
		}
	}
	if len(s.ExtraRoutes6) > 0 {
		checks.append(t, "TestCheckResourceAttr", "extra_routes6.%", fmt.Sprintf("%d", len(s.ExtraRoutes6)))
		for k, v := range s.ExtraRoutes6 {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("extra_routes6.%s.via", k), v.Via)
		}
	}
	if s.GatewayMgmt != nil {
		if s.GatewayMgmt.ConfigRevertTimer != nil {
			checks.append(t, "TestCheckResourceAttr", "gateway_mgmt.config_revert_timer", fmt.Sprintf("%d", *s.GatewayMgmt.ConfigRevertTimer))
		}
	}
	if len(s.IdpProfiles) > 0 {
		checks.append(t, "TestCheckResourceAttr", "idp_profiles.%", fmt.Sprintf("%d", len(s.IdpProfiles)))
		for k, v := range s.IdpProfiles {
			if v.BaseProfile != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("idp_profiles.%s.base_profile", k), *v.BaseProfile)
			}
			if v.Id != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("idp_profiles.%s.id", k), *v.Id)
			}
			if v.Name != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("idp_profiles.%s.name", k), *v.Name)
			}
			if v.OrgId != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("idp_profiles.%s.org_id", k), *v.OrgId)
			}
			if len(v.Overwrites) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("idp_profiles.%s.overwrites.#", k), fmt.Sprintf("%d", len(v.Overwrites)))
				for i, overwrite := range v.Overwrites {
					if overwrite.Action != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("idp_profiles.%s.overwrites.%d.action", k, i), *overwrite.Action)
					}
					if overwrite.Name != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("idp_profiles.%s.overwrites.%d.name", k, i), *overwrite.Name)
					}
					if overwrite.IpdProfileOverwriteMatching != nil {
						if len(overwrite.IpdProfileOverwriteMatching.AttackName) > 0 {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("idp_profiles.%s.overwrites.%d.matching.attack_name.#", k, i), fmt.Sprintf("%d", len(overwrite.IpdProfileOverwriteMatching.AttackName)))
							for j, attackName := range overwrite.IpdProfileOverwriteMatching.AttackName {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("idp_profiles.%s.overwrites.%d.matching.attack_name.%d", k, i, j), attackName)
							}
						}
						if len(overwrite.IpdProfileOverwriteMatching.DstSubnet) > 0 {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("idp_profiles.%s.overwrites.%d.matching.dst_subnet.#", k, i), fmt.Sprintf("%d", len(overwrite.IpdProfileOverwriteMatching.DstSubnet)))
							for j, dstSubnet := range overwrite.IpdProfileOverwriteMatching.DstSubnet {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("idp_profiles.%s.overwrites.%d.matching.dst_subnet.%d", k, i, j), dstSubnet)
							}
						}
						if len(overwrite.IpdProfileOverwriteMatching.Severity) > 0 {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("idp_profiles.%s.overwrites.%d.matching.severity.#", k, i), fmt.Sprintf("%d", len(overwrite.IpdProfileOverwriteMatching.Severity)))
							for j, severity := range overwrite.IpdProfileOverwriteMatching.Severity {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("idp_profiles.%s.overwrites.%d.matching.severity.%d", k, i, j), severity)
							}
						}
					}
				}
			}
		}
	}

	if len(s.IpConfigs) > 0 {
		checks.append(t, "TestCheckResourceAttr", "ip_configs.%", fmt.Sprintf("%d", len(s.IpConfigs)))
		for k, v := range s.IpConfigs {
			if v.Ip != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ip_configs.%s.ip", k), *v.Ip)
			}
			if v.Ip6 != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ip_configs.%s.ip6", k), *v.Ip6)
			}
			if v.Netmask != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ip_configs.%s.netmask", k), *v.Netmask)
			}
			if v.Netmask6 != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ip_configs.%s.netmask6", k), *v.Netmask6)
			}
			if len(v.SecondaryIps) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ip_configs.%s.secondary_ips.#", k), fmt.Sprintf("%d", len(v.SecondaryIps)))
				for i, secondaryIp := range v.SecondaryIps {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ip_configs.%s.secondary_ips.%d", k, i), secondaryIp)
				}
			}
			if v.IpConfigsType != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ip_configs.%s.type", k), *v.IpConfigsType)
			}
			if v.Type6 != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ip_configs.%s.type6", k), *v.Type6)
			}
		}
	}

	if s.Managed != nil {
		checks.append(t, "TestCheckResourceAttr", "managed", fmt.Sprintf("%t", *s.Managed))
	}
	if s.MapId != nil && *s.MapId != "" {
		checks.append(t, "TestCheckResourceAttr", "map_id", *s.MapId)
	}
	if s.MspId != nil && *s.MspId != "" {
		checks.append(t, "TestCheckResourceAttr", "msp_id", *s.MspId)
	}

	if len(s.Networks) > 0 {
		checks.append(t, "TestCheckResourceAttr", "networks.#", fmt.Sprintf("%d", len(s.Networks)))
		for i, network := range s.Networks {
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
					for destNatKey, destNat := range network.InternetAccess.InternetAccessDestinationNat {
						if destNat.InternalIp != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.internet_access.destination_nat.%s.internal_ip", i, destNatKey), *destNat.InternalIp)
						}
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.internet_access.destination_nat.%s.name", i, destNatKey), destNat.Name)
						if destNat.Port != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.internet_access.destination_nat.%s.port", i, destNatKey), *destNat.Port)
						}
						if destNat.WanName != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.internet_access.destination_nat.%s.wan_name", i, destNatKey), *destNat.WanName)
						}
					}
				}
				if len(network.InternetAccess.InternetAccessStaticNat) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.internet_access.static_nat.%%", i), fmt.Sprintf("%d", len(network.InternetAccess.InternetAccessStaticNat)))
					for staticNatKey, staticNat := range network.InternetAccess.InternetAccessStaticNat {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.internet_access.static_nat.%s.internal_ip", i, staticNatKey), staticNat.InternalIp)
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.internet_access.static_nat.%s.name", i, staticNatKey), staticNat.Name)
						if staticNat.WanName != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.internet_access.static_nat.%s.wan_name", i, staticNatKey), *staticNat.WanName)
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
					for groupKey, group := range network.Multicast.Groups {
						if group.RpIp != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.multicast.groups.%s.rp_ip", i, groupKey), *group.RpIp)
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
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.subnet", i), network.Subnet)
			if network.Subnet6 != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.subnet6", i), *network.Subnet6)
			}
			if len(network.Tenants) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.tenants.%%", i), fmt.Sprintf("%d", len(network.Tenants)))
				for tenantKey, tenant := range network.Tenants {
					if len(tenant.Addresses) > 0 {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.tenants.%s.addresses.#", i, tenantKey), fmt.Sprintf("%d", len(tenant.Addresses)))
						for j, address := range tenant.Addresses {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.tenants.%s.addresses.%d", i, tenantKey, j), address)
						}
					}
				}
			}
			if network.VlanId != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vlan_id", i), *network.VlanId)
			}
			if len(network.VpnAccess) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%%", i), fmt.Sprintf("%d", len(network.VpnAccess)))
				for vpnKey, vpn := range network.VpnAccess {
					if vpn.AdvertisedSubnet != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.advertised_subnet", i, vpnKey), *vpn.AdvertisedSubnet)
					}
					if vpn.AllowPing != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.allow_ping", i, vpnKey), fmt.Sprintf("%t", *vpn.AllowPing))
					}
					if vpn.NatPool != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.nat_pool", i, vpnKey), *vpn.NatPool)
					}
					if vpn.NoReadvertiseToLanBgp != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.no_readvertise_to_lan_bgp", i, vpnKey), fmt.Sprintf("%t", *vpn.NoReadvertiseToLanBgp))
					}
					if vpn.NoReadvertiseToLanOspf != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.no_readvertise_to_lan_ospf", i, vpnKey), fmt.Sprintf("%t", *vpn.NoReadvertiseToLanOspf))
					}
					if vpn.NoReadvertiseToOverlay != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.no_readvertise_to_overlay", i, vpnKey), fmt.Sprintf("%t", *vpn.NoReadvertiseToOverlay))
					}
					if len(vpn.OtherVrfs) > 0 {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.other_vrfs.#", i, vpnKey), fmt.Sprintf("%d", len(vpn.OtherVrfs)))
						for j, otherVrf := range vpn.OtherVrfs {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.other_vrfs.%d", i, vpnKey, j), otherVrf)
						}
					}
					if vpn.Routed != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.routed", i, vpnKey), fmt.Sprintf("%t", *vpn.Routed))
					}
					if vpn.SourceNat != nil {
						if vpn.SourceNat.ExternalIp != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.source_nat.external_ip", i, vpnKey), *vpn.SourceNat.ExternalIp)
						}
					}
					if vpn.SummarizedSubnet != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.summarized_subnet", i, vpnKey), *vpn.SummarizedSubnet)
					}
					if vpn.SummarizedSubnetToLanBgp != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.summarized_subnet_to_lan_bgp", i, vpnKey), *vpn.SummarizedSubnetToLanBgp)
					}
					if vpn.SummarizedSubnetToLanOspf != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.summarized_subnet_to_lan_ospf", i, vpnKey), *vpn.SummarizedSubnetToLanOspf)
					}
					if len(vpn.VpnAccessDestinationNat) > 0 {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.destination_nat.%%", i, vpnKey), fmt.Sprintf("%d", len(vpn.VpnAccessDestinationNat)))
						for destNatKey, destNat := range vpn.VpnAccessDestinationNat {
							if destNat.InternalIp != nil {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.destination_nat.%s.internal_ip", i, vpnKey, destNatKey), *destNat.InternalIp)
							}
							if destNat.Name != nil {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.destination_nat.%s.name", i, vpnKey, destNatKey), *destNat.Name)
							}
							if destNat.Port != nil {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.destination_nat.%s.port", i, vpnKey, destNatKey), *destNat.Port)
							}
						}
					}
					if len(vpn.VpnAccessStaticNat) > 0 {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.static_nat.%%", i, vpnKey), fmt.Sprintf("%d", len(vpn.VpnAccessStaticNat)))
						for staticNatKey, staticNat := range vpn.VpnAccessStaticNat {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.static_nat.%s.internal_ip", i, vpnKey, staticNatKey), staticNat.InternalIp)
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%d.vpn_access.%s.static_nat.%s.name", i, vpnKey, staticNatKey), staticNat.Name)
						}
					}
				}
			}
		}
	}

	if s.Notes != nil && *s.Notes != "" {
		checks.append(t, "TestCheckResourceAttr", "notes", *s.Notes)
	}

	if len(s.NtpServers) > 0 {
		checks.append(t, "TestCheckResourceAttr", "ntp_servers.#", fmt.Sprintf("%d", len(s.NtpServers)))
		for i, server := range s.NtpServers {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ntp_servers.%d", i), server)
		}
	}

	if s.OobIpConfig != nil {
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
		for k, v := range s.PathPreferences {
			if len(v.Paths) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("path_preferences.%s.paths.#", k), fmt.Sprintf("%d", len(v.Paths)))
				for i, path := range v.Paths {
					if path.Cost != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("path_preferences.%s.paths.%d.cost", k, i), fmt.Sprintf("%d", *path.Cost))
					}
					if path.Disabled != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("path_preferences.%s.paths.%d.disabled", k, i), fmt.Sprintf("%t", *path.Disabled))
					}
					if path.GatewayIp != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("path_preferences.%s.paths.%d.gateway_ip", k, i), *path.GatewayIp)
					}
					if path.InternetAccess != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("path_preferences.%s.paths.%d.internet_access", k, i), fmt.Sprintf("%t", *path.InternetAccess))
					}
					if path.Name != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("path_preferences.%s.paths.%d.name", k, i), *path.Name)
					}
					if len(path.Networks) > 0 {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("path_preferences.%s.paths.%d.networks.#", k, i), fmt.Sprintf("%d", len(path.Networks)))
						for j, network := range path.Networks {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("path_preferences.%s.paths.%d.networks.%d", k, i, j), network)
						}
					}
					if len(path.TargetIps) > 0 {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("path_preferences.%s.paths.%d.target_ips.#", k, i), fmt.Sprintf("%d", len(path.TargetIps)))
						for j, targetIp := range path.TargetIps {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("path_preferences.%s.paths.%d.target_ips.%d", k, i, j), targetIp)
						}
					}
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("path_preferences.%s.paths.%d.type", k, i), path.PathsType)
					if path.WanName != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("path_preferences.%s.paths.%d.wan_name", k, i), *path.WanName)
					}
				}
			}
			if v.Strategy != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("path_preferences.%s.strategy", k), *v.Strategy)
			}
		}
	}

	if len(s.PortConfig) > 0 {
		checks.append(t, "TestCheckResourceAttr", "port_config.%", fmt.Sprintf("%d", len(s.PortConfig)))
		for k, v := range s.PortConfig {
			if v.AeDisableLacp != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ae_disable_lacp", k), fmt.Sprintf("%t", *v.AeDisableLacp))
			}
			if v.AeIdx != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ae_idx", k), *v.AeIdx)
			}
			if v.AeLacpForceUp != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ae_lacp_force_up", k), fmt.Sprintf("%t", *v.AeLacpForceUp))
			}
			if v.Aggregated != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.aggregated", k), fmt.Sprintf("%t", *v.Aggregated))
			}
			if v.Critical != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.critical", k), fmt.Sprintf("%t", *v.Critical))
			}
			if v.Description != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.description", k), *v.Description)
			}
			if v.DisableAutoneg != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.disable_autoneg", k), fmt.Sprintf("%t", *v.DisableAutoneg))
			}
			if v.Disabled != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.disabled", k), fmt.Sprintf("%t", *v.Disabled))
			}
			if v.DslType != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.dsl_type", k), *v.DslType)
			}
			if v.DslVci != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.dsl_vci", k), fmt.Sprintf("%d", *v.DslVci))
			}
			if v.DslVpi != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.dsl_vpi", k), fmt.Sprintf("%d", *v.DslVpi))
			}
			if v.Duplex != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.duplex", k), *v.Duplex)
			}
			if v.LteApn != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.lte_apn", k), *v.LteApn)
			}
			if v.LteAuth != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.lte_auth", k), *v.LteAuth)
			}
			if v.LteBackup != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.lte_backup", k), fmt.Sprintf("%t", *v.LteBackup))
			}
			if v.LtePassword != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.lte_password", k), *v.LtePassword)
			}
			if v.LteUsername != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.lte_username", k), *v.LteUsername)
			}
			if v.Mtu != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.mtu", k), fmt.Sprintf("%d", *v.Mtu))
			}
			if v.Name != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.name", k), *v.Name)
			}
			if len(v.Networks) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.networks.#", k), fmt.Sprintf("%d", len(v.Networks)))
				for i, network := range v.Networks {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.networks.%d", k, i), network)
				}
			}
			if v.OuterVlanId != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.outer_vlan_id", k), fmt.Sprintf("%d", *v.OuterVlanId))
			}
			if v.PoeDisabled != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.poe_disabled", k), fmt.Sprintf("%t", *v.PoeDisabled))
			}
			if v.PortNetwork != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.port_network", k), *v.PortNetwork)
			}
			if v.PreserveDscp != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.preserve_dscp", k), fmt.Sprintf("%t", *v.PreserveDscp))
			}
			if v.Redundant != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.redundant", k), fmt.Sprintf("%t", *v.Redundant))
			}
			if v.RedundantGroup != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.redundant_group", k), fmt.Sprintf("%d", *v.RedundantGroup))
			}
			if v.RethIdx != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.reth_idx", k), *v.RethIdx)
			}
			if v.RethNode != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.reth_node", k), *v.RethNode)
			}
			if len(v.RethNodes) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.reth_nodes.#", k), fmt.Sprintf("%d", len(v.RethNodes)))
				for i, rethNode := range v.RethNodes {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.reth_nodes.%d", k, i), rethNode)
				}
			}
			if v.Speed != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.speed", k), *v.Speed)
			}
			if v.SsrNoVirtualMac != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ssr_no_virtual_mac", k), fmt.Sprintf("%t", *v.SsrNoVirtualMac))
			}
			if v.SvrPortRange != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.svr_port_range", k), *v.SvrPortRange)
			}
			if v.TrafficShaping != nil {
				if len(v.TrafficShaping.ClassPercentages) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.traffic_shaping.class_percentages.#", k), fmt.Sprintf("%d", len(v.TrafficShaping.ClassPercentages)))
					for i, percentage := range v.TrafficShaping.ClassPercentages {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.traffic_shaping.class_percentages.%d", k, i), fmt.Sprintf("%d", percentage))
					}
				}
				if v.TrafficShaping.Enabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.traffic_shaping.enabled", k), fmt.Sprintf("%t", *v.TrafficShaping.Enabled))
				}
				if v.TrafficShaping.MaxTxKbps != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.traffic_shaping.max_tx_kbps", k), fmt.Sprintf("%d", *v.TrafficShaping.MaxTxKbps))
				}
			}
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.usage", k), v.Usage)
			if v.VlanId != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.vlan_id", k), *v.VlanId)
			}
			if v.PortIpConfig != nil {
				if len(v.PortIpConfig.Dns) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ip_config.dns.#", k), fmt.Sprintf("%d", len(v.PortIpConfig.Dns)))
					for i, dns := range v.PortIpConfig.Dns {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ip_config.dns.%d", k, i), dns)
					}
				}
				if len(v.PortIpConfig.DnsSuffix) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ip_config.dns_suffix.#", k), fmt.Sprintf("%d", len(v.PortIpConfig.DnsSuffix)))
					for i, suffix := range v.PortIpConfig.DnsSuffix {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ip_config.dns_suffix.%d", k, i), suffix)
					}
				}
				if v.PortIpConfig.Gateway != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ip_config.gateway", k), *v.PortIpConfig.Gateway)
				}
				if v.PortIpConfig.Gateway6 != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ip_config.gateway6", k), *v.PortIpConfig.Gateway6)
				}
				if v.PortIpConfig.Ip != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ip_config.ip", k), *v.PortIpConfig.Ip)
				}
				if v.PortIpConfig.Ip6 != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ip_config.ip6", k), *v.PortIpConfig.Ip6)
				}
				if v.PortIpConfig.Netmask != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ip_config.netmask", k), *v.PortIpConfig.Netmask)
				}
				if v.PortIpConfig.Netmask6 != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ip_config.netmask6", k), *v.PortIpConfig.Netmask6)
				}
				if v.PortIpConfig.Network != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ip_config.network", k), *v.PortIpConfig.Network)
				}
				if v.PortIpConfig.PoserPassword != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ip_config.poser_password", k), *v.PortIpConfig.PoserPassword)
				}
				if v.PortIpConfig.PppoeAuth != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ip_config.pppoe_auth", k), *v.PortIpConfig.PppoeAuth)
				}
				if v.PortIpConfig.PppoeUsername != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ip_config.pppoe_username", k), *v.PortIpConfig.PppoeUsername)
				}
				if v.PortIpConfig.PortIpConfigType != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ip_config.type", k), *v.PortIpConfig.PortIpConfigType)
				}
				if v.PortIpConfig.Type6 != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ip_config.type6", k), *v.PortIpConfig.Type6)
				}
			}
			if len(v.VpnPaths) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.vpn_paths.%%", k), fmt.Sprintf("%d", len(v.VpnPaths)))
				for vpnPathKey, vpnPath := range v.VpnPaths {
					if vpnPath.BfdProfile != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.vpn_paths.%s.bfd_profile", k, vpnPathKey), *vpnPath.BfdProfile)
					}
					if vpnPath.BfdUseTunnelMode != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.vpn_paths.%s.bfd_use_tunnel_mode", k, vpnPathKey), fmt.Sprintf("%t", *vpnPath.BfdUseTunnelMode))
					}
					if vpnPath.Preference != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.vpn_paths.%s.preference", k, vpnPathKey), fmt.Sprintf("%d", *vpnPath.Preference))
					}
					if vpnPath.Role != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.vpn_paths.%s.role", k, vpnPathKey), *vpnPath.Role)
					}
					if vpnPath.TrafficShaping != nil {
						if len(vpnPath.TrafficShaping.ClassPercentages) > 0 {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.vpn_paths.%s.traffic_shaping.class_percentages.#", k, vpnPathKey), fmt.Sprintf("%d", len(vpnPath.TrafficShaping.ClassPercentages)))
							for i, percentage := range vpnPath.TrafficShaping.ClassPercentages {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.vpn_paths.%s.traffic_shaping.class_percentages.%d", k, vpnPathKey, i), fmt.Sprintf("%d", percentage))
							}
						}
						if vpnPath.TrafficShaping.Enabled != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.vpn_paths.%s.traffic_shaping.enabled", k, vpnPathKey), fmt.Sprintf("%t", *vpnPath.TrafficShaping.Enabled))
						}
						if vpnPath.TrafficShaping.MaxTxKbps != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.vpn_paths.%s.traffic_shaping.max_tx_kbps", k, vpnPathKey), fmt.Sprintf("%d", *vpnPath.TrafficShaping.MaxTxKbps))
						}
					}
				}
			}
			if v.WanArpPolicer != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_arp_policer", k), *v.WanArpPolicer)
			}
			if v.WanDisableSpeedtest != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_disable_speedtest", k), fmt.Sprintf("%t", *v.WanDisableSpeedtest))
			}
			if v.WanExtIp != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_ext_ip", k), *v.WanExtIp)
			}
			if v.WanExtIp6 != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_ext_ip6", k), *v.WanExtIp6)
			}
			if len(v.WanExtraRoutes) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_extra_routes.%%", k), fmt.Sprintf("%d", len(v.WanExtraRoutes)))
				for wanRouteKey, wanRoute := range v.WanExtraRoutes {
					if wanRoute.Via != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_extra_routes.%s.via", k, wanRouteKey), *wanRoute.Via)
					}
				}
			}
			if len(v.WanExtraRoutes6) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_extra_routes6.%%", k), fmt.Sprintf("%d", len(v.WanExtraRoutes6)))
				for wanRoute6Key, wanRoute6 := range v.WanExtraRoutes6 {
					if wanRoute6.Via != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_extra_routes6.%s.via", k, wanRoute6Key), *wanRoute6.Via)
					}
				}
			}
			if len(v.WanNetworks) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_networks.#", k), fmt.Sprintf("%d", len(v.WanNetworks)))
				for i, wanNetwork := range v.WanNetworks {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_networks.%d", k, i), wanNetwork)
				}
			}
			if v.WanProbeOverride != nil {
				if len(v.WanProbeOverride.Ip6s) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_probe_override.ip6s.#", k), fmt.Sprintf("%d", len(v.WanProbeOverride.Ip6s)))
					for i, ip6 := range v.WanProbeOverride.Ip6s {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_probe_override.ip6s.%d", k, i), ip6)
					}
				}
				if len(v.WanProbeOverride.Ips) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_probe_override.ips.#", k), fmt.Sprintf("%d", len(v.WanProbeOverride.Ips)))
					for i, ip := range v.WanProbeOverride.Ips {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_probe_override.ips.%d", k, i), ip)
					}
				}
				if v.WanProbeOverride.ProbeProfile != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_probe_override.probe_profile", k), *v.WanProbeOverride.ProbeProfile)
				}
			}
			if v.WanSourceNat != nil {
				if v.WanSourceNat.Disabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_source_nat.disabled", k), fmt.Sprintf("%t", *v.WanSourceNat.Disabled))
				}
				if v.WanSourceNat.NatPool != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_source_nat.nat_pool", k), *v.WanSourceNat.NatPool)
				}
				if v.WanSourceNat.Nat6Pool != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_source_nat.nat6_pool", k), *v.WanSourceNat.Nat6Pool)
				}
			}
			if v.WanType != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.wan_type", k), *v.WanType)
			}
		}
	}

	if s.PortMirroring != nil {
		if s.PortMirroring.PortMirror != nil {
			if s.PortMirroring.PortMirror.FamilyType != nil {
				checks.append(t, "TestCheckResourceAttr", "port_mirroring.port_mirror.family_type", *s.PortMirroring.PortMirror.FamilyType)
			}
			if len(s.PortMirroring.PortMirror.IngressPortIds) > 0 {
				checks.append(t, "TestCheckResourceAttr", "port_mirroring.port_mirror.ingress_port_ids.#", fmt.Sprintf("%d", len(s.PortMirroring.PortMirror.IngressPortIds)))
				for i, portId := range s.PortMirroring.PortMirror.IngressPortIds {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_mirroring.port_mirror.ingress_port_ids.%d", i), portId)
				}
			}
			if s.PortMirroring.PortMirror.OutputPortId != nil {
				checks.append(t, "TestCheckResourceAttr", "port_mirroring.port_mirror.output_port_id", *s.PortMirroring.PortMirror.OutputPortId)
			}
			if s.PortMirroring.PortMirror.Rate != nil {
				checks.append(t, "TestCheckResourceAttr", "port_mirroring.port_mirror.rate", fmt.Sprintf("%d", *s.PortMirroring.PortMirror.Rate))
			}
			if s.PortMirroring.PortMirror.RunLength != nil {
				checks.append(t, "TestCheckResourceAttr", "port_mirroring.port_mirror.run_length", fmt.Sprintf("%d", *s.PortMirroring.PortMirror.RunLength))
			}
		}
	}
	if s.RouterId != nil && *s.RouterId != "" {
		checks.append(t, "TestCheckResourceAttr", "router_id", *s.RouterId)
	}
	if len(s.RoutingPolicies) > 0 {
		checks.append(t, "TestCheckResourceAttr", "routing_policies.%", fmt.Sprintf("%d", len(s.RoutingPolicies)))
		for k, v := range s.RoutingPolicies {
			if len(v.Terms) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.#", k), fmt.Sprintf("%d", len(v.Terms)))
				for i, term := range v.Terms {
					if term.Actions != nil {
						if term.Actions.Accept != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.actions.accept", k, i), fmt.Sprintf("%t", *term.Actions.Accept))
						}
						if len(term.Actions.AddCommunity) > 0 {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.actions.add_community.#", k, i), fmt.Sprintf("%d", len(term.Actions.AddCommunity)))
							for j, addCommunity := range term.Actions.AddCommunity {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.actions.add_community.%d", k, i, j), addCommunity)
							}
						}
						if len(term.Actions.AddTargetVrfs) > 0 {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.actions.add_target_vrfs.#", k, i), fmt.Sprintf("%d", len(term.Actions.AddTargetVrfs)))
							for j, addTargetVrf := range term.Actions.AddTargetVrfs {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.actions.add_target_vrfs.%d", k, i, j), addTargetVrf)
							}
						}
						if len(term.Actions.Community) > 0 {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.actions.community.#", k, i), fmt.Sprintf("%d", len(term.Actions.Community)))
							for j, community := range term.Actions.Community {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.actions.community.%d", k, i, j), community)
							}
						}
						if len(term.Actions.ExcludeAsPath) > 0 {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.actions.exclude_as_path.#", k, i), fmt.Sprintf("%d", len(term.Actions.ExcludeAsPath)))
							for j, excludeAsPath := range term.Actions.ExcludeAsPath {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.actions.exclude_as_path.%d", k, i, j), excludeAsPath)
							}
						}
						if len(term.Actions.ExcludeCommunity) > 0 {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.actions.exclude_community.#", k, i), fmt.Sprintf("%d", len(term.Actions.ExcludeCommunity)))
							for j, excludeCommunity := range term.Actions.ExcludeCommunity {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.actions.exclude_community.%d", k, i, j), excludeCommunity)
							}
						}
						if len(term.Actions.ExportCommunities) > 0 {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.actions.export_communities.#", k, i), fmt.Sprintf("%d", len(term.Actions.ExportCommunities)))
							for j, exportCommunity := range term.Actions.ExportCommunities {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.actions.export_communities.%d", k, i, j), exportCommunity)
							}
						}
						if term.Actions.LocalPreference != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.actions.local_preference", k, i), *term.Actions.LocalPreference)
						}
						if len(term.Actions.PrependAsPath) > 0 {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.actions.prepend_as_path.#", k, i), fmt.Sprintf("%d", len(term.Actions.PrependAsPath)))
							for j, prependAsPath := range term.Actions.PrependAsPath {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.actions.prepend_as_path.%d", k, i, j), prependAsPath)
							}
						}
					}
					if term.RoutingPolicyTermMatching != nil {
						if len(term.RoutingPolicyTermMatching.AsPath) > 0 {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.as_path.#", k, i), fmt.Sprintf("%d", len(term.RoutingPolicyTermMatching.AsPath)))
							for j, asPath := range term.RoutingPolicyTermMatching.AsPath {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.as_path.%d", k, i, j), asPath)
							}
						}
						if len(term.RoutingPolicyTermMatching.Community) > 0 {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.community.#", k, i), fmt.Sprintf("%d", len(term.RoutingPolicyTermMatching.Community)))
							for j, community := range term.RoutingPolicyTermMatching.Community {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.community.%d", k, i, j), community)
							}
						}
						if len(term.RoutingPolicyTermMatching.Network) > 0 {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.network.#", k, i), fmt.Sprintf("%d", len(term.RoutingPolicyTermMatching.Network)))
							for j, network := range term.RoutingPolicyTermMatching.Network {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.network.%d", k, i, j), network)
							}
						}
						if len(term.RoutingPolicyTermMatching.Prefix) > 0 {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.prefix.#", k, i), fmt.Sprintf("%d", len(term.RoutingPolicyTermMatching.Prefix)))
							for j, prefix := range term.RoutingPolicyTermMatching.Prefix {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.prefix.%d", k, i, j), prefix)
							}
						}
						if len(term.RoutingPolicyTermMatching.Protocol) > 0 {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.protocol.#", k, i), fmt.Sprintf("%d", len(term.RoutingPolicyTermMatching.Protocol)))
							for j, protocol := range term.RoutingPolicyTermMatching.Protocol {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.protocol.%d", k, i, j), protocol)
							}
						}
						if term.RoutingPolicyTermMatching.RouteExists != nil {
							if term.RoutingPolicyTermMatching.RouteExists.Route != nil {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.route_exists.route", k, i), *term.RoutingPolicyTermMatching.RouteExists.Route)
							}
							if term.RoutingPolicyTermMatching.RouteExists.VrfName != nil {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.route_exists.vrf_name", k, i), *term.RoutingPolicyTermMatching.RouteExists.VrfName)
							}
						}
						if len(term.RoutingPolicyTermMatching.VpnNeighborMac) > 0 {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.vpn_neighbor_mac.#", k, i), fmt.Sprintf("%d", len(term.RoutingPolicyTermMatching.VpnNeighborMac)))
							for j, vpnNeighborMac := range term.RoutingPolicyTermMatching.VpnNeighborMac {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.vpn_neighbor_mac.%d", k, i, j), vpnNeighborMac)
							}
						}
						if len(term.RoutingPolicyTermMatching.VpnPath) > 0 {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.vpn_path.#", k, i), fmt.Sprintf("%d", len(term.RoutingPolicyTermMatching.VpnPath)))
							for j, vpnPath := range term.RoutingPolicyTermMatching.VpnPath {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.vpn_path.%d", k, i, j), vpnPath)
							}
						}
						if term.RoutingPolicyTermMatching.VpnPathSla != nil {
							if term.RoutingPolicyTermMatching.VpnPathSla.MaxJitter != nil {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.vpn_path_sla.max_jitter", k, i), fmt.Sprintf("%d", *term.RoutingPolicyTermMatching.VpnPathSla.MaxJitter))
							}
							if term.RoutingPolicyTermMatching.VpnPathSla.MaxLatency != nil {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.vpn_path_sla.max_latency", k, i), fmt.Sprintf("%d", *term.RoutingPolicyTermMatching.VpnPathSla.MaxLatency))
							}
							if term.RoutingPolicyTermMatching.VpnPathSla.MaxLoss != nil {
								checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.%d.matching.vpn_path_sla.max_loss", k, i), fmt.Sprintf("%d", *term.RoutingPolicyTermMatching.VpnPathSla.MaxLoss))
							}
						}
					}
				}
			}
		}
	}
	if len(s.ServicePolicies) > 0 {
		checks.append(t, "TestCheckResourceAttr", "service_policies.#", fmt.Sprintf("%d", len(s.ServicePolicies)))
		for i, v := range s.ServicePolicies {
			if v.Action != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.action", i), *v.Action)
			}
			if v.Antivirus != nil {
				if v.Antivirus.AvprofileId != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.antivirus.avprofile_id", i), *v.Antivirus.AvprofileId)
				}
				if v.Antivirus.Enabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.antivirus.enabled", i), fmt.Sprintf("%t", *v.Antivirus.Enabled))
				}
				if v.Antivirus.Profile != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.antivirus.profile", i), *v.Antivirus.Profile)
				}
			}
			if v.Appqoe != nil {
				if v.Appqoe.Enabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.appqoe.enabled", i), fmt.Sprintf("%t", *v.Appqoe.Enabled))
				}
			}
			if len(v.Ewf) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.ewf.#", i), fmt.Sprintf("%d", len(v.Ewf)))
				for j, ewfV := range v.Ewf {
					if ewfV.AlertOnly != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.ewf.%d.alert_only", i, j), fmt.Sprintf("%t", *ewfV.AlertOnly))
					}
					if ewfV.BlockMessage != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.ewf.%d.block_message", i, j), *ewfV.BlockMessage)
					}
					if ewfV.Enabled != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.ewf.%d.enabled", i, j), fmt.Sprintf("%t", *ewfV.Enabled))
					}
					if ewfV.Profile != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.ewf.%d.profile", i, j), *ewfV.Profile)
					}
				}
			}
			if v.Idp != nil {
				if v.Idp.AlertOnly != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.idp.alert_only", i), fmt.Sprintf("%t", *v.Idp.AlertOnly))
				}
				if v.Idp.Enabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.idp.enabled", i), fmt.Sprintf("%t", *v.Idp.Enabled))
				}
				if v.Idp.IdpprofileId != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.idp.idpprofile_id", i), *v.Idp.IdpprofileId)
				}
				if v.Idp.Profile != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.idp.profile", i), *v.Idp.Profile)
				}
			}
			if v.LocalRouting != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.local_routing", i), fmt.Sprintf("%t", *v.LocalRouting))
			}
			if v.Name != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.name", i), *v.Name)
			}
			if v.PathPreference != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.path_preference", i), *v.PathPreference)
			}
			if v.ServicepolicyId != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.servicepolicy_id", i), *v.ServicepolicyId)
			}
			if len(v.Services) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.services.#", i), fmt.Sprintf("%d", len(v.Services)))
				for j, service := range v.Services {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.services.%d", i, j), service)
				}
			}
			if v.SslProxy != nil {
				if v.SslProxy.CiphersCategory != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.ssl_proxy.ciphers_category", i), *v.SslProxy.CiphersCategory)
				}
				if v.SslProxy.Enabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.ssl_proxy.enabled", i), fmt.Sprintf("%t", *v.SslProxy.Enabled))
				}
			}
			if v.Skyatp != nil {
				if v.Skyatp.DnsDgaDetection != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.skyatp.dns_dga_detection", i), *v.Skyatp.DnsDgaDetection)
				}
				if v.Skyatp.DnsTunnelDetection != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.skyatp.dns_tunnel_detection", i), *v.Skyatp.DnsTunnelDetection)
				}
				if v.Skyatp.HttpInspection != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.skyatp.http_inspection", i), *v.Skyatp.HttpInspection)
				}
				if v.Skyatp.IotDevicePolicy != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.skyatp.iot_device_policy", i), *v.Skyatp.IotDevicePolicy)
				}
			}
			if v.Syslog != nil {
				if v.Syslog.Enabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.syslog.enabled", i), fmt.Sprintf("%t", *v.Syslog.Enabled))
				}
				if len(v.Syslog.ServerNames) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.syslog.server_names.#", i), fmt.Sprintf("%d", len(v.Syslog.ServerNames)))
					for j, serverName := range v.Syslog.ServerNames {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.syslog.server_names.%d", i, j), serverName)
					}
				}
			}
			if len(v.Tenants) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.tenants.#", i), fmt.Sprintf("%d", len(v.Tenants)))
				for j, tenant := range v.Tenants {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("service_policies.%d.tenants.%d", i, j), tenant)
				}
			}
		}
	}
	if len(s.TunnelConfigs) > 0 {
		checks.append(t, "TestCheckResourceAttr", "tunnel_configs.%", fmt.Sprintf("%d", len(s.TunnelConfigs)))
		for i, v := range s.TunnelConfigs {
			if v.AutoProvision != nil {
				if v.AutoProvision.AutoProvisionPrimary != nil {
					if len(v.AutoProvision.AutoProvisionPrimary.ProbeIps) > 0 {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.auto_provision.primary.probe_ips.#", i), fmt.Sprintf("%d", len(v.AutoProvision.AutoProvisionPrimary.ProbeIps)))
						for j, probeIp := range v.AutoProvision.AutoProvisionPrimary.ProbeIps {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.auto_provision.primary.probe_ips.%d", i, j), probeIp)
						}
					}
					if len(v.AutoProvision.AutoProvisionPrimary.WanNames) > 0 {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.auto_provision.primary.wan_names.#", i), fmt.Sprintf("%d", len(v.AutoProvision.AutoProvisionPrimary.WanNames)))
						for j, wanName := range v.AutoProvision.AutoProvisionPrimary.WanNames {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.auto_provision.primary.wan_names.%d", i, j), wanName)
						}
					}
				}
				if v.AutoProvision.AutoProvisionSecondary != nil {
					if len(v.AutoProvision.AutoProvisionSecondary.ProbeIps) > 0 {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.auto_provision.secondary.probe_ips.#", i), fmt.Sprintf("%d", len(v.AutoProvision.AutoProvisionSecondary.ProbeIps)))
						for j, probeIp := range v.AutoProvision.AutoProvisionSecondary.ProbeIps {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.auto_provision.secondary.probe_ips.%d", i, j), probeIp)
						}
					}
					if len(v.AutoProvision.AutoProvisionSecondary.WanNames) > 0 {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.auto_provision.secondary.wan_names.#", i), fmt.Sprintf("%d", len(v.AutoProvision.AutoProvisionSecondary.WanNames)))
						for j, wanName := range v.AutoProvision.AutoProvisionSecondary.WanNames {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.auto_provision.secondary.wan_names.%d", i, j), wanName)
						}
					}
				}
				if v.AutoProvision.Enabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.auto_provision.enabled", i), fmt.Sprintf("%t", *v.AutoProvision.Enabled))
				}
				if v.AutoProvision.Latlng != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.auto_provision.latlng.lat", i), fmt.Sprintf("%f", v.AutoProvision.Latlng.Lat))
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.auto_provision.latlng.lng", i), fmt.Sprintf("%f", v.AutoProvision.Latlng.Lng))
				}
				if v.AutoProvision.Provider != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.auto_provision.provider", i), *v.AutoProvision.Provider)
				}
				if v.AutoProvision.Region != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.auto_provision.region", i), *v.AutoProvision.Region)
				}
				if v.AutoProvision.ServiceConnection != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.auto_provision.service_connection", i), *v.AutoProvision.ServiceConnection)
				}
			}
			if v.IkeLifetime != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.ike_lifetime", i), fmt.Sprintf("%d", *v.IkeLifetime))
			}
			if v.IkeMode != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.ike_mode", i), *v.IkeMode)
			}
			if len(v.IkeProposals) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.ike_proposals.#", i), fmt.Sprintf("%d", len(v.IkeProposals)))
				for j, ikeProposal := range v.IkeProposals {
					if ikeProposal.AuthAlgo != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.ike_proposals.%d.auth_algo", i, j), *ikeProposal.AuthAlgo)
					}
					if ikeProposal.DhGroup != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.ike_proposals.%d.dh_group", i, j), *ikeProposal.DhGroup)
					}
					if ikeProposal.EncAlgo != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.ike_proposals.%d.enc_algo", i, j), *ikeProposal.EncAlgo)
					}
				}
			}
			if v.IpsecLifetime != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.ipsec_lifetime", i), fmt.Sprintf("%d", *v.IpsecLifetime))
			}
			if len(v.IpsecProposals) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.ipsec_proposals.#", i), fmt.Sprintf("%d", len(v.IpsecProposals)))
				for j, ipsecProposal := range v.IpsecProposals {
					if ipsecProposal.AuthAlgo != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.ipsec_proposals.%d.auth_algo", i, j), *ipsecProposal.AuthAlgo)
					}
					if ipsecProposal.DhGroup != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.ipsec_proposals.%d.dh_group", i, j), *ipsecProposal.DhGroup)
					}
					if ipsecProposal.EncAlgo != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.ipsec_proposals.%d.enc_algo", i, j), *ipsecProposal.EncAlgo)
					}
				}
			}
			if v.LocalId != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.local_id", i), *v.LocalId)
			}
			if len(v.LocalSubnets) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.local_subnets.#", i), fmt.Sprintf("%d", len(v.LocalSubnets)))
				for j, localSubnet := range v.LocalSubnets {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.local_subnets.%d", i, j), localSubnet)
				}
			}
			if v.Mode != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.mode", i), *v.Mode)
			}
			if len(v.Networks) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.networks.#", i), fmt.Sprintf("%d", len(v.Networks)))
				for j, network := range v.Networks {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.networks.%d", i, j), network)
				}
			}
			if v.Primary != nil {
				if len(v.Primary.Hosts) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.primary.hosts.#", i), fmt.Sprintf("%d", len(v.Primary.Hosts)))
					for j, host := range v.Primary.Hosts {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.primary.hosts.%d", i, j), host)
					}
				}
				if len(v.Primary.InternalIps) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.primary.internal_ips.#", i), fmt.Sprintf("%d", len(v.Primary.InternalIps)))
					for j, internalIp := range v.Primary.InternalIps {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.primary.internal_ips.%d", i, j), internalIp)
					}
				}
				if len(v.Primary.ProbeIps) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.primary.probe_ips.#", i), fmt.Sprintf("%d", len(v.Primary.ProbeIps)))
					for j, probeIp := range v.Primary.ProbeIps {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.primary.probe_ips.%d", i, j), probeIp)
					}
				}
				if len(v.Primary.RemoteIds) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.primary.remote_ids.#", i), fmt.Sprintf("%d", len(v.Primary.RemoteIds)))
					for j, remoteId := range v.Primary.RemoteIds {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.primary.remote_ids.%d", i, j), remoteId)
					}
				}
				if len(v.Primary.WanNames) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.primary.wan_names.#", i), fmt.Sprintf("%d", len(v.Primary.WanNames)))
					for j, wanName := range v.Primary.WanNames {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.primary.wan_names.%d", i, j), wanName)
					}
				}
			}
			if v.Probe != nil {
				if v.Probe.Interval != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.probe.interval", i), fmt.Sprintf("%d", *v.Probe.Interval))
				}
				if v.Probe.Threshold != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.probe.threshold", i), fmt.Sprintf("%d", *v.Probe.Threshold))
				}
				if v.Probe.Timeout != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.probe.timeout", i), fmt.Sprintf("%d", *v.Probe.Timeout))
				}
				if v.Probe.ProbeType != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.probe.type", i), *v.Probe.ProbeType)
				}
			}
			if v.Protocol != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.protocol", i), *v.Protocol)
			}
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.provider", i), v.Provider)
			if v.Psk != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.psk", i), *v.Psk)
			}
			if len(v.RemoteSubnets) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.remote_subnets.#", i), fmt.Sprintf("%d", len(v.RemoteSubnets)))
				for j, remoteSubnet := range v.RemoteSubnets {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.remote_subnets.%d", i, j), remoteSubnet)
				}
			}
			if v.Secondary != nil {
				if len(v.Secondary.Hosts) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.secondary.hosts.#", i), fmt.Sprintf("%d", len(v.Secondary.Hosts)))
					for j, host := range v.Secondary.Hosts {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.secondary.hosts.%d", i, j), host)
					}
				}
				if len(v.Secondary.InternalIps) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.secondary.internal_ips.#", i), fmt.Sprintf("%d", len(v.Secondary.InternalIps)))
					for j, internalIp := range v.Secondary.InternalIps {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.secondary.internal_ips.%d", i, j), internalIp)
					}
				}
				if len(v.Secondary.ProbeIps) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.secondary.probe_ips.#", i), fmt.Sprintf("%d", len(v.Secondary.ProbeIps)))
					for j, probeIp := range v.Secondary.ProbeIps {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.secondary.probe_ips.%d", i, j), probeIp)
					}
				}
				if len(v.Secondary.RemoteIds) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.secondary.remote_ids.#", i), fmt.Sprintf("%d", len(v.Secondary.RemoteIds)))
					for j, remoteId := range v.Secondary.RemoteIds {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.secondary.remote_ids.%d", i, j), remoteId)
					}
				}
				if len(v.Secondary.WanNames) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.secondary.wan_names.#", i), fmt.Sprintf("%d", len(v.Secondary.WanNames)))
					for j, wanName := range v.Secondary.WanNames {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.secondary.wan_names.%d", i, j), wanName)
					}
				}
			}
			if v.Version != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunnel_configs.%s.version", i), *v.Version)
			}
		}
	}
	if s.TunnelProviderOptions != nil {
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
	if len(s.Vars) > 0 {
		checks.append(t, "TestCheckResourceAttr", "vars.%", fmt.Sprintf("%d", len(s.Vars)))
	}
	if s.VrfConfig != nil {
		if s.VrfConfig.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "vrf_config.enabled", fmt.Sprintf("%t", *s.VrfConfig.Enabled))
		}
	}
	if len(s.VrfInstances) > 0 {
		checks.append(t, "TestCheckResourceAttr", "vrf_instances.%", fmt.Sprintf("%d", len(s.VrfInstances)))
		for k, v := range s.VrfInstances {
			if len(v.Networks) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vrf_instances.%s.networks.#", k), fmt.Sprintf("%d", len(v.Networks)))
				for i, network := range v.Networks {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vrf_instances.%s.networks.%d", k, i), network)
				}
			}
		}
	}
	if s.X != nil {
		checks.append(t, "TestCheckResourceAttr", "x", fmt.Sprintf("%v", *s.X))
	}
	if s.Y != nil {
		checks.append(t, "TestCheckResourceAttr", "y", fmt.Sprintf("%v", *s.Y))
	}

	return checks
}
