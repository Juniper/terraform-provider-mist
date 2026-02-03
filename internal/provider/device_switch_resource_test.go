package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_device_switch"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDeviceSwitchModel(t *testing.T) {
	resourceType := "device_switch"
	t.Skipf("Skipping %s tests, as they require a real device.", resourceType)

	type testStep struct {
		config DeviceSwitchModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: DeviceSwitchModel{
						DeviceId: "",
						Name:     "test_device_switch",
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/site_setting_resource/site_setting_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		var FixtureDeviceSwitchModel DeviceSwitchModel
		err = hcl.Decode(&FixtureDeviceSwitchModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureDeviceSwitchModel,
				},
			},
		}
	}

	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_device_switch.DeviceSwitchResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				siteConfig, siteRef := GetSiteBaseConfig(GetTestOrgId())
				config := step.config

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				f.Body().SetAttributeRaw("site_id", hclwrite.TokensForIdentifier(siteRef))
				combinedConfig := siteConfig + "\n\n" + Render(resourceType, tName, string(f.Bytes()))

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

func (s *DeviceSwitchModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType)+"."+tName, tracker)

	// Required string attributes
	checks.append(t, "TestCheckResourceAttrSet", "site_id")
	checks.append(t, "TestCheckResourceAttr", "name", s.Name)
	checks.append(t, "TestCheckResourceAttrSet", "device_id")

	// Computed-only attributes - check for presence
	checks.append(t, "TestCheckResourceAttrSet", "mac")
	checks.append(t, "TestCheckResourceAttrSet", "model")
	checks.append(t, "TestCheckResourceAttrSet", "org_id")
	checks.append(t, "TestCheckResourceAttrSet", "serial")
	checks.append(t, "TestCheckResourceAttrSet", "type")

	// Conditional checks for all DeviceSwitchModel attributes
	if len(s.AclPolicies) > 0 {
		checks.append(t, "TestCheckResourceAttrSet", "acl_policies")
		// Check nested attributes of AclPolicies
		for i, policy := range s.AclPolicies {
			if policy.Name != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acl_policies.%d.name", i), *policy.Name)
			}
			if len(policy.SrcTags) > 0 {
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("acl_policies.%d.src_tags", i))
				// Check individual source tag values
				for j, tag := range policy.SrcTags {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acl_policies.%d.src_tags.%d", i, j), tag)
				}
			}
			if len(policy.Actions) > 0 {
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("acl_policies.%d.actions", i))
				// Check nested Actions
				for j, action := range policy.Actions {
					if action.Action != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acl_policies.%d.actions.%d.action", i, j), *action.Action)
					}
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acl_policies.%d.actions.%d.dst_tag", i, j), action.DstTag)
				}
			}
		}
	}
	if len(s.AclTags) > 0 {
		checks.append(t, "TestCheckResourceAttrSet", "acl_tags")
		// Check nested attributes of AclTags
		for key, tag := range s.AclTags {
			if len(tag.EtherTypes) > 0 {
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("acl_tags.%s.ether_types", key))
				// Check individual ether type values
				for i, etherType := range tag.EtherTypes {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acl_tags.%s.ether_types.%d", key, i), etherType)
				}
			}
			if tag.GbpTag != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acl_tags.%s.gbp_tag", key), fmt.Sprintf("%d", *tag.GbpTag))
			}
			if len(tag.Macs) > 0 {
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("acl_tags.%s.macs", key))
				// Check individual MAC address values
				for i, mac := range tag.Macs {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acl_tags.%s.macs.%d", key, i), mac)
				}
			}
			if tag.Network != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acl_tags.%s.network", key), *tag.Network)
			}
			if tag.PortUsage != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acl_tags.%s.port_usage", key), *tag.PortUsage)
			}
			if tag.RadiusGroup != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acl_tags.%s.radius_group", key), *tag.RadiusGroup)
			}
			if len(tag.Specs) > 0 {
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("acl_tags.%s.specs", key))
				// Check nested Specs
				for i, spec := range tag.Specs {
					if spec.PortRange != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acl_tags.%s.specs.%d.port_range", key, i), *spec.PortRange)
					}
					if spec.Protocol != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acl_tags.%s.specs.%d.protocol", key, i), *spec.Protocol)
					}
				}
			}
			if len(tag.Subnets) > 0 {
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("acl_tags.%s.subnets", key))
				// Check individual subnet values
				for i, subnet := range tag.Subnets {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acl_tags.%s.subnets.%d", key, i), subnet)
				}
			}
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acl_tags.%s.type", key), tag.AclTagsType)
		}
	}
	if len(s.AdditionalConfigCmds) > 0 {
		checks.append(t, "TestCheckResourceAttrSet", "additional_config_cmds")
		// Check individual additional config commands
		for i, cmd := range s.AdditionalConfigCmds {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("additional_config_cmds.%d", i), cmd)
		}
	}
	// BgpConfig map
	if len(s.BgpConfig) > 0 {
		for key, bgp := range s.BgpConfig {
			basePath := fmt.Sprintf("bgp_config.%s", key)
			if bgp.AuthKey != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".auth_key", *bgp.AuthKey)
			}
			if bgp.BfdMinimumInterval != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".bfd_minimum_interval", fmt.Sprintf("%d", *bgp.BfdMinimumInterval))
			}
			if bgp.ExportPolicy != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".export_policy", *bgp.ExportPolicy)
			}
			if bgp.HoldTime != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".hold_time", fmt.Sprintf("%d", *bgp.HoldTime))
			}

			if bgp.ImportPolicy != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".import_policy", *bgp.ImportPolicy)
			}
			checks.append(t, "TestCheckResourceAttr", basePath+".local_as", bgp.LocalAs)
			if len(bgp.Networks) > 0 {
				checks.append(t, "TestCheckResourceAttr", basePath+".networks.#", fmt.Sprintf("%d", len(bgp.Networks)))
				for i, network := range bgp.Networks {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.networks.%d", basePath, i), network)
				}
			}
			checks.append(t, "TestCheckResourceAttr", basePath+".type", bgp.BgpConfigType)
			// BgpConfig neighbors
			if len(bgp.Neighbors) > 0 {
				for neighborKey, neighbor := range bgp.Neighbors {
					neighborPath := fmt.Sprintf("%s.neighbors.%s", basePath, neighborKey)
					if neighbor.ExportPolicy != nil {
						checks.append(t, "TestCheckResourceAttr", neighborPath+".export_policy", *neighbor.ExportPolicy)
					}
					if neighbor.HoldTime != nil {
						checks.append(t, "TestCheckResourceAttr", neighborPath+".hold_time", fmt.Sprintf("%d", *neighbor.HoldTime))
					}
					if neighbor.ImportPolicy != nil {
						checks.append(t, "TestCheckResourceAttr", neighborPath+".import_policy", *neighbor.ImportPolicy)
					}
					if neighbor.MultihopTtl != nil {
						checks.append(t, "TestCheckResourceAttr", neighborPath+".multihop_ttl", fmt.Sprintf("%d", *neighbor.MultihopTtl))
					}
					checks.append(t, "TestCheckResourceAttr", neighborPath+".neighbor_as", neighbor.NeighborAs)
				}
			}
		}
	}
	if s.DefaultPortUsage != nil {
		checks.append(t, "TestCheckResourceAttr", "default_port_usage", *s.DefaultPortUsage)
	}
	if s.DhcpSnooping != nil {
		checks.append(t, "TestCheckResourceAttrSet", "dhcp_snooping")
		// Check nested attributes of DhcpSnooping
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
			checks.append(t, "TestCheckResourceAttrSet", "dhcp_snooping.networks")
			// Check individual DHCP snooping network values
			for i, network := range s.DhcpSnooping.Networks {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcp_snooping.networks.%d", i), network)
			}
		}
	}
	if s.DhcpdConfig != nil {
		checks.append(t, "TestCheckResourceAttrSet", "dhcpd_config")
		// Check nested attributes of DhcpdConfig
		if s.DhcpdConfig.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "dhcpd_config.enabled", fmt.Sprintf("%t", *s.DhcpdConfig.Enabled))
		}
		if len(s.DhcpdConfig.Config) > 0 {
			checks.append(t, "TestCheckResourceAttrSet", "dhcpd_config.config")
			// Check nested Config values
			for key, config := range s.DhcpdConfig.Config {
				if len(config.DnsServers) > 0 {
					checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("dhcpd_config.config.%s.dns_servers", key))
					// Check individual DHCP config DNS server values
					for i, server := range config.DnsServers {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.dns_servers.%d", key, i), server)
					}
				}
				if len(config.DnsSuffix) > 0 {
					checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("dhcpd_config.config.%s.dns_suffix", key))
					// Check individual DHCP config DNS suffix values
					for i, suffix := range config.DnsSuffix {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.dns_suffix.%d", key, i), suffix)
					}
				}
				if config.Gateway != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.gateway", key), *config.Gateway)
				}
				if config.IpEnd != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.ip_end", key), *config.IpEnd)
				}
				if config.IpEnd6 != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.ip_end6", key), *config.IpEnd6)
				}
				if config.IpStart != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.ip_start", key), *config.IpStart)
				}
				if config.IpStart6 != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.ip_start6", key), *config.IpStart6)
				}
				if config.LeaseTime != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.lease_time", key), fmt.Sprintf("%d", *config.LeaseTime))
				}
				if config.ServerIdOverride != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.server_id_override", key), fmt.Sprintf("%t", *config.ServerIdOverride))
				}
				if len(config.Servers) > 0 {
					checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("dhcpd_config.config.%s.servers", key))
					// Check individual DHCP config server values
					for i, server := range config.Servers {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.servers.%d", key, i), server)
					}
				}
				if len(config.Servers6) > 0 {
					checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("dhcpd_config.config.%s.servers6", key))
					// Check individual DHCP config IPv6 server values
					for i, server := range config.Servers6 {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.servers6.%d", key, i), server)
					}
				}
				if config.ConfigType != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.type", key), *config.ConfigType)
				}
				if config.Type6 != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.type6", key), *config.Type6)
				}
				if len(config.FixedBindings) > 0 {
					checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("dhcpd_config.config.%s.fixed_bindings", key))
					for bindingKey, binding := range config.FixedBindings {
						prefix := fmt.Sprintf("dhcpd_config.config.%s.fixed_bindings.%s", key, bindingKey)
						if binding.Ip != nil {
							checks.append(t, "TestCheckResourceAttr", prefix+".ip", *binding.Ip)
						}
						if binding.Ip6 != nil {
							checks.append(t, "TestCheckResourceAttr", prefix+".ip6", *binding.Ip6)
						}
						if binding.Name != nil {
							checks.append(t, "TestCheckResourceAttr", prefix+".name", *binding.Name)
						}
					}
				}
				if len(config.Options) > 0 {
					checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("dhcpd_config.config.%s.options", key))
				}
				if len(config.VendorEncapsulated) > 0 {
					checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("dhcpd_config.config.%s.vendor_encapsulated", key))
				}
			}
		}
	}
	if s.DisableAutoConfig != nil {
		checks.append(t, "TestCheckResourceAttr", "disable_auto_config", fmt.Sprintf("%t", *s.DisableAutoConfig))
	}
	if len(s.DnsServers) > 0 {
		checks.append(t, "TestCheckResourceAttrSet", "dns_servers")
		// Check individual DNS server values
		for i, server := range s.DnsServers {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dns_servers.%d", i), server)
		}
	}
	if len(s.DnsSuffix) > 0 {
		checks.append(t, "TestCheckResourceAttrSet", "dns_suffix")
		// Check individual DNS suffix values
		for i, suffix := range s.DnsSuffix {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dns_suffix.%d", i), suffix)
		}
	}
	if len(s.ExtraRoutes) > 0 {
		checks.append(t, "TestCheckResourceAttrSet", "extra_routes")
		// Check nested attributes of ExtraRoutes
		for key, route := range s.ExtraRoutes {
			if route.Discard != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("extra_routes.%s.discard", key), fmt.Sprintf("%t", *route.Discard))
			}
			if route.Metric != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("extra_routes.%s.metric", key), fmt.Sprintf("%d", *route.Metric))
			}
			if route.NoResolve != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("extra_routes.%s.no_resolve", key), fmt.Sprintf("%t", *route.NoResolve))
			}
			if route.Preference != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("extra_routes.%s.preference", key), fmt.Sprintf("%d", *route.Preference))
			}
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("extra_routes.%s.via", key), route.Via)
			if len(route.NextQualified) > 0 {
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("extra_routes.%s.next_qualified", key))
				// Check NextQualified nested values
				for nqKey, nq := range route.NextQualified {
					if nq.Metric != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("extra_routes.%s.next_qualified.%s.metric", key, nqKey), fmt.Sprintf("%d", *nq.Metric))
					}
					if nq.Preference != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("extra_routes.%s.next_qualified.%s.preference", key, nqKey), fmt.Sprintf("%d", *nq.Preference))
					}
				}
			}
		}
	}
	if len(s.ExtraRoutes6) > 0 {
		checks.append(t, "TestCheckResourceAttrSet", "extra_routes6")
		// Check nested attributes of ExtraRoutes6
		for key, route := range s.ExtraRoutes6 {
			if route.Discard != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("extra_routes6.%s.discard", key), fmt.Sprintf("%t", *route.Discard))
			}
			if route.Metric != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("extra_routes6.%s.metric", key), fmt.Sprintf("%d", *route.Metric))
			}
			if route.NoResolve != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("extra_routes6.%s.no_resolve", key), fmt.Sprintf("%t", *route.NoResolve))
			}
			if route.Preference != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("extra_routes6.%s.preference", key), fmt.Sprintf("%d", *route.Preference))
			}
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("extra_routes6.%s.via", key), route.Via)
			if len(route.NextQualified) > 0 {
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("extra_routes6.%s.next_qualified", key))
				// Check NextQualified nested values for IPv6
				for nqKey, nq := range route.NextQualified {
					if nq.Metric != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("extra_routes6.%s.next_qualified.%s.metric", key, nqKey), fmt.Sprintf("%d", *nq.Metric))
					}
					if nq.Preference != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("extra_routes6.%s.next_qualified.%s.preference", key, nqKey), fmt.Sprintf("%d", *nq.Preference))
					}
				}
			}
		}
	}
	if s.IpConfig != nil {
		checks.append(t, "TestCheckResourceAttrSet", "ip_config")
		// Check nested attributes of IpConfig
		if len(s.IpConfig.Dns) > 0 {
			checks.append(t, "TestCheckResourceAttrSet", "ip_config.dns")
			// Check individual IP config DNS values
			for i, dns := range s.IpConfig.Dns {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ip_config.dns.%d", i), dns)
			}
		}
		if len(s.IpConfig.DnsSuffix) > 0 {
			checks.append(t, "TestCheckResourceAttrSet", "ip_config.dns_suffix")
			// Check individual IP config DNS suffix values
			for i, suffix := range s.IpConfig.DnsSuffix {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ip_config.dns_suffix.%d", i), suffix)
			}
		}
		if s.IpConfig.Gateway != nil {
			checks.append(t, "TestCheckResourceAttr", "ip_config.gateway", *s.IpConfig.Gateway)
		}
		if s.IpConfig.Ip != nil {
			checks.append(t, "TestCheckResourceAttr", "ip_config.ip", *s.IpConfig.Ip)
		}
		if s.IpConfig.Netmask != nil {
			checks.append(t, "TestCheckResourceAttr", "ip_config.netmask", *s.IpConfig.Netmask)
		}
		if s.IpConfig.Network != nil {
			checks.append(t, "TestCheckResourceAttr", "ip_config.network", *s.IpConfig.Network)
		}
		if s.IpConfig.IpConfigType != nil {
			checks.append(t, "TestCheckResourceAttr", "ip_config.type", *s.IpConfig.IpConfigType)
		}
	}
	if len(s.LocalPortConfig) > 0 {
		checks.append(t, "TestCheckResourceAttrSet", "local_port_config")
		// Check nested attributes of LocalPortConfig
		for key, portConfig := range s.LocalPortConfig {
			if portConfig.AllNetworks != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("local_port_config.%s.all_networks", key), fmt.Sprintf("%t", *portConfig.AllNetworks))
			}
			if portConfig.AllowDhcpd != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("local_port_config.%s.allow_dhcpd", key), fmt.Sprintf("%t", *portConfig.AllowDhcpd))
			}
			if portConfig.Description != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("local_port_config.%s.description", key), *portConfig.Description)
			}
			if portConfig.DisableAutoneg != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("local_port_config.%s.disable_autoneg", key), fmt.Sprintf("%t", *portConfig.DisableAutoneg))
			}
			if portConfig.Disabled != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("local_port_config.%s.disabled", key), fmt.Sprintf("%t", *portConfig.Disabled))
			}
			if portConfig.Duplex != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("local_port_config.%s.duplex", key), *portConfig.Duplex)
			}
			if portConfig.EnableQos != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("local_port_config.%s.enable_qos", key), fmt.Sprintf("%t", *portConfig.EnableQos))
			}
			if portConfig.MacLimit != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("local_port_config.%s.mac_limit", key), fmt.Sprintf("%d", *portConfig.MacLimit))
			}
			if portConfig.Mode != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("local_port_config.%s.mode", key), *portConfig.Mode)
			}
			if portConfig.Mtu != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("local_port_config.%s.mtu", key), fmt.Sprintf("%d", *portConfig.Mtu))
			}
			if len(portConfig.Networks) > 0 {
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("local_port_config.%s.networks", key))
				// Check individual local port config network values
				for i, network := range portConfig.Networks {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("local_port_config.%s.networks.%d", key, i), network)
				}
			}
			if portConfig.Note != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("local_port_config.%s.note", key), *portConfig.Note)
			}
			if portConfig.PortNetwork != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("local_port_config.%s.port_network", key), *portConfig.PortNetwork)
			}
			if portConfig.Speed != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("local_port_config.%s.speed", key), *portConfig.Speed)
			}
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("local_port_config.%s.usage", key), portConfig.Usage)
			if portConfig.StormControl != nil {
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("local_port_config.%s.storm_control", key))
			}
		}
	}
	if s.Managed != nil {
		checks.append(t, "TestCheckResourceAttr", "managed", fmt.Sprintf("%t", *s.Managed))
	}
	if s.MapId != nil {
		checks.append(t, "TestCheckResourceAttr", "map_id", *s.MapId)
	}
	if s.MistNac != nil {
		checks.append(t, "TestCheckResourceAttrSet", "mist_nac")
		// Check nested attributes of MistNac
		if s.MistNac.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "mist_nac.enabled", fmt.Sprintf("%t", *s.MistNac.Enabled))
		}
		if s.MistNac.Network != nil {
			checks.append(t, "TestCheckResourceAttr", "mist_nac.network", *s.MistNac.Network)
		}
	}
	if len(s.Networks) > 0 {
		checks.append(t, "TestCheckResourceAttrSet", "networks")
		// Check nested attributes of Networks
		for key, network := range s.Networks {
			if network.Gateway != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%s.gateway", key), *network.Gateway)
			}
			if network.Gateway6 != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%s.gateway6", key), *network.Gateway6)
			}
			if network.Isolation != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%s.isolation", key), fmt.Sprintf("%t", *network.Isolation))
			}
			if network.IsolationVlanId != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%s.isolation_vlan_id", key), *network.IsolationVlanId)
			}
			if network.Subnet != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%s.subnet", key), *network.Subnet)
			}
			if network.Subnet6 != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%s.subnet6", key), *network.Subnet6)
			}
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("networks.%s.vlan_id", key), network.VlanId)
		}
	}
	if s.Notes != nil {
		checks.append(t, "TestCheckResourceAttr", "notes", *s.Notes)
	}
	if len(s.NtpServers) > 0 {
		checks.append(t, "TestCheckResourceAttrSet", "ntp_servers")
		// Check individual NTP server values
		for i, server := range s.NtpServers {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ntp_servers.%d", i), server)
		}
	}
	if s.OobIpConfig != nil {
		checks.append(t, "TestCheckResourceAttrSet", "oob_ip_config")
		// Check nested attributes of OobIpConfig
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
	if len(s.OspfAreas) > 0 {
		checks.append(t, "TestCheckResourceAttrSet", "ospf_areas")
		// Check nested attributes of OspfAreas
		for key, area := range s.OspfAreas {
			if area.IncludeLoopback != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ospf_areas.%s.include_loopback", key), fmt.Sprintf("%t", *area.IncludeLoopback))
			}
			if area.OspfAreasType != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ospf_areas.%s.type", key), *area.OspfAreasType)
			}
			if len(area.OspfNetworks) > 0 {
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("ospf_areas.%s.networks", key))
				// Check nested OspfNetworks
				for netKey, network := range area.OspfNetworks {
					if len(network.AuthKeys) > 0 {
						checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("ospf_areas.%s.networks.%s.auth_keys", key, netKey))
						// Check individual auth key values
						for i, authKey := range network.AuthKeys {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ospf_areas.%s.networks.%s.auth_keys.%s", key, netKey, i), authKey)
						}
					}
					if network.AuthPassword != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ospf_areas.%s.networks.%s.auth_password", key, netKey), *network.AuthPassword)
					}
					if network.AuthType != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ospf_areas.%s.networks.%s.auth_type", key, netKey), *network.AuthType)
					}
					if network.BfdMinimumInterval != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ospf_areas.%s.networks.%s.bfd_minimum_interval", key, netKey), fmt.Sprintf("%d", *network.BfdMinimumInterval))
					}
					if network.DeadInterval != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ospf_areas.%s.networks.%s.dead_interval", key, netKey), fmt.Sprintf("%d", *network.DeadInterval))
					}
					if network.HelloInterval != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ospf_areas.%s.networks.%s.hello_interval", key, netKey), fmt.Sprintf("%d", *network.HelloInterval))
					}
					if network.InterfaceType != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ospf_areas.%s.networks.%s.interface_type", key, netKey), *network.InterfaceType)
					}
					if network.Metric != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ospf_areas.%s.networks.%s.metric", key, netKey), fmt.Sprintf("%d", *network.Metric))
					}
					if network.Passive != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ospf_areas.%s.networks.%s.passive", key, netKey), fmt.Sprintf("%t", *network.Passive))
					}
				}
			}
		}
	}
	if s.OspfConfig != nil {
		checks.append(t, "TestCheckResourceAttrSet", "ospf_config")
		// Check nested attributes of OspfConfig
		if s.OspfConfig.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "ospf_config.enabled", fmt.Sprintf("%t", *s.OspfConfig.Enabled))
		}
		if s.OspfConfig.ExportPolicy != nil {
			checks.append(t, "TestCheckResourceAttr", "ospf_config.export_policy", *s.OspfConfig.ExportPolicy)
		}
		if s.OspfConfig.ImportPolicy != nil {
			checks.append(t, "TestCheckResourceAttr", "ospf_config.import_policy", *s.OspfConfig.ImportPolicy)
		}
		if s.OspfConfig.ReferenceBandwidth != nil {
			checks.append(t, "TestCheckResourceAttr", "ospf_config.reference_bandwidth", *s.OspfConfig.ReferenceBandwidth)
		}
		if len(s.OspfConfig.Areas) > 0 {
			checks.append(t, "TestCheckResourceAttrSet", "ospf_config.areas")
			// Check nested Areas
			for areaKey, area := range s.OspfConfig.Areas {
				if area.NoSummary != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ospf_config.areas.%s.no_summary", areaKey), fmt.Sprintf("%t", *area.NoSummary))
				}
			}
		}
	}
	if len(s.OtherIpConfigs) > 0 {
		checks.append(t, "TestCheckResourceAttrSet", "other_ip_configs")
		// Check nested attributes of OtherIpConfigs
		for key, config := range s.OtherIpConfigs {
			if config.EvpnAnycast != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("other_ip_configs.%s.evpn_anycast", key), fmt.Sprintf("%t", *config.EvpnAnycast))
			}
			if config.Ip != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("other_ip_configs.%s.ip", key), *config.Ip)
			}
			if config.Ip6 != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("other_ip_configs.%s.ip6", key), *config.Ip6)
			}
			if config.Netmask != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("other_ip_configs.%s.netmask", key), *config.Netmask)
			}
			if config.Netmask6 != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("other_ip_configs.%s.netmask6", key), *config.Netmask6)
			}
			if config.OtherIpConfigsType != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("other_ip_configs.%s.type", key), *config.OtherIpConfigsType)
			}
			if config.Type6 != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("other_ip_configs.%s.type6", key), *config.Type6)
			}
		}
	}
	if len(s.PortConfig) > 0 {
		checks.append(t, "TestCheckResourceAttrSet", "port_config")
		// Check nested attributes of PortConfig
		for key, port := range s.PortConfig {
			if port.AeDisableLacp != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ae_disable_lacp", key), fmt.Sprintf("%t", *port.AeDisableLacp))
			}
			if port.AeIdx != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ae_idx", key), fmt.Sprintf("%d", *port.AeIdx))
			}
			if port.AeLacpSlow != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ae_lacp_slow", key), fmt.Sprintf("%t", *port.AeLacpSlow))
			}
			if port.Aggregated != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.aggregated", key), fmt.Sprintf("%t", *port.Aggregated))
			}
			if port.Critical != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.critical", key), fmt.Sprintf("%t", *port.Critical))
			}
			if port.Description != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.description", key), *port.Description)
			}
			if port.DisableAutoneg != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.disable_autoneg", key), fmt.Sprintf("%t", *port.DisableAutoneg))
			}
			if port.Duplex != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.duplex", key), *port.Duplex)
			}
			if port.DynamicUsage != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.dynamic_usage", key), *port.DynamicUsage)
			}
			if port.Mtu != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.mtu", key), fmt.Sprintf("%d", *port.Mtu))
			}
			if len(port.Networks) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.networks.#", key), fmt.Sprintf("%d", len(port.Networks)))
				for i, network := range port.Networks {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.networks.%d", key, i), network)
				}
			}
			if port.PoeDisabled != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.poe_disabled", key), fmt.Sprintf("%t", *port.PoeDisabled))
			}
			if port.PortNetwork != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.port_network", key), *port.PortNetwork)
			}
			if port.Speed != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.speed", key), *port.Speed)
			}
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.usage", key), port.Usage)
		}
	}
	if len(s.PortConfigOverwrite) > 0 {
		checks.append(t, "TestCheckResourceAttrSet", "port_config_overwrite")
		// Check nested attributes of PortConfigOverwrite
		for key, port := range s.PortConfigOverwrite {
			if port.Description != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config_overwrite.%s.description", key), *port.Description)
			}
			if port.Disabled != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config_overwrite.%s.disabled", key), fmt.Sprintf("%t", *port.Disabled))
			}
			if port.Duplex != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config_overwrite.%s.duplex", key), *port.Duplex)
			}
			if port.MacLimit != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config_overwrite.%s.mac_limit", key), *port.MacLimit)
			}
			if port.PoeDisabled != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config_overwrite.%s.poe_disabled", key), fmt.Sprintf("%t", *port.PoeDisabled))
			}
			if port.PortNetwork != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config_overwrite.%s.port_network", key), *port.PortNetwork)
			}
			if port.Speed != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config_overwrite.%s.speed", key), *port.Speed)
			}
		}
	}
	if len(s.PortMirroring) > 0 {
		checks.append(t, "TestCheckResourceAttrSet", "port_mirroring")
		// Check nested attributes of PortMirroring
		for key, mirror := range s.PortMirroring {
			if len(mirror.InputNetworksIngress) > 0 {
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("port_mirroring.%s.input_networks_ingress", key))
				// Check individual input networks ingress values
				for i, network := range mirror.InputNetworksIngress {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_mirroring.%s.input_networks_ingress.%d", key, i), network)
				}
			}
			if len(mirror.InputPortIdsEgress) > 0 {
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("port_mirroring.%s.input_port_ids_egress", key))
				// Check individual input port IDs egress values
				for i, portId := range mirror.InputPortIdsEgress {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_mirroring.%s.input_port_ids_egress.%d", key, i), portId)
				}
			}
			if len(mirror.InputPortIdsIngress) > 0 {
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("port_mirroring.%s.input_port_ids_ingress", key))
				// Check individual input port IDs ingress values
				for i, portId := range mirror.InputPortIdsIngress {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_mirroring.%s.input_port_ids_ingress.%d", key, i), portId)
				}
			}
			if mirror.OutputIpAddress != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_mirroring.%s.output_ip_address", key), *mirror.OutputIpAddress)
			}
			if mirror.OutputNetwork != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_mirroring.%s.output_network", key), *mirror.OutputNetwork)
			}
			if mirror.OutputPortId != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_mirroring.%s.output_port_id", key), *mirror.OutputPortId)
			}
		}
	}
	if len(s.PortUsages) > 0 {
		checks.append(t, "TestCheckResourceAttrSet", "port_usages")
		// Check nested attributes of PortUsages (key attributes only due to size)
		for key, usage := range s.PortUsages {
			if usage.AllNetworks != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.all_networks", key), fmt.Sprintf("%t", *usage.AllNetworks))
			}
			if usage.AllowDhcpd != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.allow_dhcpd", key), fmt.Sprintf("%t", *usage.AllowDhcpd))
			}
			if usage.BypassAuthWhenServerDownForVoip != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.bypass_auth_when_server_down_for_voip", key), fmt.Sprintf("%t", *usage.BypassAuthWhenServerDownForVoip))
			}
			if usage.Description != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.description", key), *usage.Description)
			}
			if usage.Disabled != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.disabled", key), fmt.Sprintf("%t", *usage.Disabled))
			}
			if usage.EnableQos != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.enable_qos", key), fmt.Sprintf("%t", *usage.EnableQos))
			}
			if usage.MacLimit != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.mac_limit", key), *usage.MacLimit)
			}
			if usage.Mode != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.mode", key), *usage.Mode)
			}
			if len(usage.Networks) > 0 {
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("port_usages.%s.networks", key))
				// Check individual port usage network values
				for i, network := range usage.Networks {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.networks.%d", key, i), network)
				}
			}
			if usage.PortNetwork != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.port_network", key), *usage.PortNetwork)
			}
			if usage.PoePriority != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.poe_priority", key), *usage.PoePriority)
			}
			if usage.Speed != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.speed", key), *usage.Speed)
			}
			if len(usage.Rules) > 0 {
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("port_usages.%s.rules", key))
				// Check nested Rules
				for i, rule := range usage.Rules {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.rules.%d.src", key, i), rule.Src)
					if rule.Equals != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.rules.%d.equals", key, i), *rule.Equals)
					}
					if len(rule.EqualsAny) > 0 {
						checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("port_usages.%s.rules.%d.equals_any", key, i))
						// Check individual equals any values
						for j, equalsAnyValue := range rule.EqualsAny {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.rules.%d.equals_any.%d", key, i, j), equalsAnyValue)
						}
					}
					if rule.Usage != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.rules.%d.usage", key, i), *rule.Usage)
					}
				}
			}
			if usage.StormControl != nil {
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("port_usages.%s.storm_control", key))
			}
			if usage.StpDisable != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.stp_disable", key), fmt.Sprintf("%t", *usage.StpDisable))
			}
			if usage.StpRequired != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.stp_required", key), fmt.Sprintf("%t", *usage.StpRequired))
			}
		}
	}
	if s.RadiusConfig != nil {
		checks.append(t, "TestCheckResourceAttrSet", "radius_config")
		// Check nested attributes of RadiusConfig
		if s.RadiusConfig.AcctImmediateUpdate != nil {
			checks.append(t, "TestCheckResourceAttr", "radius_config.acct_immediate_update", fmt.Sprintf("%t", *s.RadiusConfig.AcctImmediateUpdate))
		}
		if s.RadiusConfig.AcctInterimInterval != nil {
			checks.append(t, "TestCheckResourceAttr", "radius_config.acct_interim_interval", fmt.Sprintf("%d", *s.RadiusConfig.AcctInterimInterval))
		}
		if len(s.RadiusConfig.AcctServers) > 0 {
			checks.append(t, "TestCheckResourceAttrSet", "radius_config.acct_servers")
			// Check nested AcctServers
			for i, server := range s.RadiusConfig.AcctServers {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radius_config.acct_servers.%d.host", i), server.Host)
				if server.KeywrapEnabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radius_config.acct_servers.%d.keywrap_enabled", i), fmt.Sprintf("%t", *server.KeywrapEnabled))
				}
				if server.KeywrapFormat != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radius_config.acct_servers.%d.keywrap_format", i), *server.KeywrapFormat)
				}
				if server.KeywrapKek != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radius_config.acct_servers.%d.keywrap_kek", i), *server.KeywrapKek)
				}
				if server.KeywrapMack != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radius_config.acct_servers.%d.keywrap_mack", i), *server.KeywrapMack)
				}
				if server.Port != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radius_config.acct_servers.%d.port", i), *server.Port)
				}
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radius_config.acct_servers.%d.secret", i), server.Secret)
			}
		}
		if s.RadiusConfig.AuthServerSelection != nil {
			checks.append(t, "TestCheckResourceAttr", "radius_config.auth_server_selection", *s.RadiusConfig.AuthServerSelection)
		}
		if len(s.RadiusConfig.AuthServers) > 0 {
			checks.append(t, "TestCheckResourceAttrSet", "radius_config.auth_servers")
			// Check nested AuthServers
			for i, server := range s.RadiusConfig.AuthServers {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radius_config.auth_servers.%d.host", i), server.Host)
				if server.KeywrapEnabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radius_config.auth_servers.%d.keywrap_enabled", i), fmt.Sprintf("%t", *server.KeywrapEnabled))
				}
				if server.KeywrapFormat != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radius_config.auth_servers.%d.keywrap_format", i), *server.KeywrapFormat)
				}
				if server.KeywrapKek != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radius_config.auth_servers.%d.keywrap_kek", i), *server.KeywrapKek)
				}
				if server.KeywrapMack != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radius_config.auth_servers.%d.keywrap_mack", i), *server.KeywrapMack)
				}
				if server.Port != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radius_config.auth_servers.%d.port", i), *server.Port)
				}
				if server.RequireMessageAuthenticator != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radius_config.auth_servers.%d.require_message_authenticator", i), fmt.Sprintf("%t", *server.RequireMessageAuthenticator))
				}
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radius_config.auth_servers.%d.secret", i), server.Secret)
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
	if s.RemoteSyslog != nil {
		checks.append(t, "TestCheckResourceAttrSet", "remote_syslog")
		// Check nested attributes of RemoteSyslog
		if s.RemoteSyslog.Archive != nil {
			checks.append(t, "TestCheckResourceAttrSet", "remote_syslog.archive")
			// Check nested Archive
			if s.RemoteSyslog.Archive.Files != nil {
				checks.append(t, "TestCheckResourceAttr", "remote_syslog.archive.files", *s.RemoteSyslog.Archive.Files)
			}
			if s.RemoteSyslog.Archive.Size != nil {
				checks.append(t, "TestCheckResourceAttr", "remote_syslog.archive.size", *s.RemoteSyslog.Archive.Size)
			}
		}
		if len(s.RemoteSyslog.Cacerts) > 0 {
			checks.append(t, "TestCheckResourceAttrSet", "remote_syslog.cacerts")
			// Check individual CA certificate values
			for i, cacert := range s.RemoteSyslog.Cacerts {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.cacerts.%d", i), cacert)
			}
		}
		if s.RemoteSyslog.Console != nil {
			checks.append(t, "TestCheckResourceAttrSet", "remote_syslog.console")
			// Check nested Console Contents
			if len(s.RemoteSyslog.Console.Contents) > 0 {
				checks.append(t, "TestCheckResourceAttrSet", "remote_syslog.console.contents")
				for i, content := range s.RemoteSyslog.Console.Contents {
					if content.Facility != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.console.contents.%d.facility", i), *content.Facility)
					}
					if content.Severity != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.console.contents.%d.severity", i), *content.Severity)
					}
				}
			}
		}
		if s.RemoteSyslog.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "remote_syslog.enabled", fmt.Sprintf("%t", *s.RemoteSyslog.Enabled))
		}
		if len(s.RemoteSyslog.Files) > 0 {
			checks.append(t, "TestCheckResourceAttrSet", "remote_syslog.files")
			// Check nested Files
			for i, file := range s.RemoteSyslog.Files {
				if file.Archive != nil {
					checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("remote_syslog.files.%d.archive", i))
					if file.Archive.Files != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.files.%d.archive.files", i), *file.Archive.Files)
					}
					if file.Archive.Size != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.files.%d.archive.size", i), *file.Archive.Size)
					}
				}
				if len(file.Contents) > 0 {
					checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("remote_syslog.files.%d.contents", i))
				}
				if file.EnableTls != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.files.%d.enable_tls", i), fmt.Sprintf("%t", *file.EnableTls))
				}
				if file.ExplicitPriority != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.files.%d.explicit_priority", i), fmt.Sprintf("%t", *file.ExplicitPriority))
				}
				if file.File != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.files.%d.file", i), *file.File)
				}
				if file.Match != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.files.%d.match", i), *file.Match)
				}
				if file.StructuredData != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.files.%d.structured_data", i), fmt.Sprintf("%t", *file.StructuredData))
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
			checks.append(t, "TestCheckResourceAttrSet", "remote_syslog.servers")
			// Check nested Servers
			for i, server := range s.RemoteSyslog.Servers {
				if len(server.Contents) > 0 {
					checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("remote_syslog.servers.%d.contents", i))
					for j, content := range server.Contents {
						if content.Facility != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.servers.%d.contents.%d.facility", i, j), *content.Facility)
						}
						if content.Severity != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.servers.%d.contents.%d.severity", i, j), *content.Severity)
						}
					}
				}
				if server.ExplicitPriority != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.servers.%d.explicit_priority", i), fmt.Sprintf("%t", *server.ExplicitPriority))
				}
				if server.Facility != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.servers.%d.facility", i), *server.Facility)
				}
				if server.Host != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.servers.%d.host", i), *server.Host)
				}
				if server.Match != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.servers.%d.match", i), *server.Match)
				}
				if server.Port != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.servers.%d.port", i), *server.Port)
				}
				if server.Protocol != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.servers.%d.protocol", i), *server.Protocol)
				}
				if server.RoutingInstance != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.servers.%d.routing_instance", i), *server.RoutingInstance)
				}
				if server.ServerName != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.servers.%d.server_name", i), *server.ServerName)
				}
				if server.Severity != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.servers.%d.severity", i), *server.Severity)
				}
				if server.SourceAddress != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.servers.%d.source_address", i), *server.SourceAddress)
				}
				if server.StructuredData != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.servers.%d.structured_data", i), fmt.Sprintf("%t", *server.StructuredData))
				}
				if server.Tag != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.servers.%d.tag", i), *server.Tag)
				}
			}
		}
		if s.RemoteSyslog.TimeFormat != nil {
			checks.append(t, "TestCheckResourceAttr", "remote_syslog.time_format", *s.RemoteSyslog.TimeFormat)
		}
		if len(s.RemoteSyslog.Users) > 0 {
			checks.append(t, "TestCheckResourceAttrSet", "remote_syslog.users")
			// Check nested Users
			for i, user := range s.RemoteSyslog.Users {
				if len(user.Contents) > 0 {
					checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("remote_syslog.users.%d.contents", i))
					for j, content := range user.Contents {
						if content.Facility != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.users.%d.contents.%d.facility", i, j), *content.Facility)
						}
						if content.Severity != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.users.%d.contents.%d.severity", i, j), *content.Severity)
						}
					}
				}
				if user.Match != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.users.%d.match", i), *user.Match)
				}
				if user.User != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.users.%d.user", i), *user.User)
				}
			}
		}
	}
	if s.Role != nil {
		checks.append(t, "TestCheckResourceAttr", "role", *s.Role)
	}
	if s.RouterId != nil {
		checks.append(t, "TestCheckResourceAttr", "router_id", *s.RouterId)
	}

	// Check routing_policies if present
	if len(s.RoutingPolicies) > 0 {
		checks.append(t, "TestCheckResourceAttr", "routing_policies.%", fmt.Sprintf("%d", len(s.RoutingPolicies)))
		for k, v := range s.RoutingPolicies {
			if len(v.Terms) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.#", k), fmt.Sprintf("%d", len(v.Terms)))
				// Check for terms by name using TestCheckTypeSetElemNestedAttrs to handle ordering
				for _, term := range v.Terms {
					termChecks := make(map[string]string)
					termChecks["name"] = term.Name

					if term.RoutingPolicyTermActions != nil {
						if term.RoutingPolicyTermActions.Accept != nil {
							termChecks["actions.accept"] = fmt.Sprintf("%t", *term.RoutingPolicyTermActions.Accept)
						}
						if term.RoutingPolicyTermActions.LocalPreference != nil {
							termChecks["actions.local_preference"] = *term.RoutingPolicyTermActions.LocalPreference
						}
						if len(term.RoutingPolicyTermActions.Community) > 0 {
							termChecks["actions.community.#"] = fmt.Sprintf("%d", len(term.RoutingPolicyTermActions.Community))
							for j, community := range term.RoutingPolicyTermActions.Community {
								termChecks[fmt.Sprintf("actions.community.%d", j)] = community
							}
						}
						if len(term.RoutingPolicyTermActions.PrependAsPath) > 0 {
							termChecks["actions.prepend_as_path.#"] = fmt.Sprintf("%d", len(term.RoutingPolicyTermActions.PrependAsPath))
							for j, prependAsPath := range term.RoutingPolicyTermActions.PrependAsPath {
								termChecks[fmt.Sprintf("actions.prepend_as_path.%d", j)] = prependAsPath
							}
						}
					}
					if term.Matching != nil {
						if len(term.Matching.AsPath) > 0 {
							termChecks["matching.as_path.#"] = fmt.Sprintf("%d", len(term.Matching.AsPath))
							for j, asPath := range term.Matching.AsPath {
								termChecks[fmt.Sprintf("matching.as_path.%d", j)] = asPath
							}
						}
						if len(term.Matching.Community) > 0 {
							termChecks["matching.community.#"] = fmt.Sprintf("%d", len(term.Matching.Community))
							for j, community := range term.Matching.Community {
								termChecks[fmt.Sprintf("matching.community.%d", j)] = community
							}
						}
						if len(term.Matching.Prefix) > 0 {
							termChecks["matching.prefix.#"] = fmt.Sprintf("%d", len(term.Matching.Prefix))
							for j, prefix := range term.Matching.Prefix {
								termChecks[fmt.Sprintf("matching.prefix.%d", j)] = prefix
							}
						}
						if len(term.Matching.Protocol) > 0 {
							termChecks["matching.protocol.#"] = fmt.Sprintf("%d", len(term.Matching.Protocol))
							for j, protocol := range term.Matching.Protocol {
								termChecks[fmt.Sprintf("matching.protocol.%d", j)] = protocol
							}
						}
					}
					checks.appendSetNestedCheck(t, fmt.Sprintf("routing_policies.%s.terms.*", k), termChecks)
				}
			}
		}
	}

	if s.SnmpConfig != nil {
		checks.append(t, "TestCheckResourceAttrSet", "snmp_config")
		// Check nested attributes of SnmpConfig
		if len(s.SnmpConfig.ClientList) > 0 {
			checks.append(t, "TestCheckResourceAttrSet", "snmp_config.client_list")
			// Check nested ClientList
			for i, client := range s.SnmpConfig.ClientList {
				if client.ClientListName != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("snmp_config.client_list.%d.client_list_name", i), *client.ClientListName)
				}
				if len(client.Clients) > 0 {
					checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("snmp_config.client_list.%d.clients", i))
					// Check individual client values
					for j, clientValue := range client.Clients {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("snmp_config.client_list.%d.clients.%d", i, j), clientValue)
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
			checks.append(t, "TestCheckResourceAttrSet", "snmp_config.trap_groups")
			// Check nested TrapGroups
			for i, trap := range s.SnmpConfig.TrapGroups {
				if len(trap.Categories) > 0 {
					checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("snmp_config.trap_groups.%d.categories", i))
					// Check individual category values
					for j, category := range trap.Categories {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("snmp_config.trap_groups.%d.categories.%d", i, j), category)
					}
				}
				if trap.GroupName != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("snmp_config.trap_groups.%d.group_name", i), *trap.GroupName)
				}
				if len(trap.Targets) > 0 {
					checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("snmp_config.trap_groups.%d.targets", i))
					// Check individual target values
					for j, target := range trap.Targets {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("snmp_config.trap_groups.%d.targets.%d", i, j), target)
					}
				}
				if trap.Version != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("snmp_config.trap_groups.%d.version", i), *trap.Version)
				}
			}
		}
		if len(s.SnmpConfig.V2cConfig) > 0 {
			checks.append(t, "TestCheckResourceAttrSet", "snmp_config.v2c_config")
			// Check nested V2cConfig
			for i, v2c := range s.SnmpConfig.V2cConfig {
				if v2c.Authorization != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("snmp_config.v2c_config.%d.authorization", i), *v2c.Authorization)
				}
				if v2c.ClientListName != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("snmp_config.v2c_config.%d.client_list_name", i), *v2c.ClientListName)
				}
				if v2c.CommunityName != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("snmp_config.v2c_config.%d.community_name", i), *v2c.CommunityName)
				}
				if v2c.View != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("snmp_config.v2c_config.%d.view", i), *v2c.View)
				}
			}
		}
		if s.SnmpConfig.V3Config != nil {
			checks.append(t, "TestCheckResourceAttrSet", "snmp_config.v3_config")
			// Check nested V3Config (key attributes due to complexity)
			if len(s.SnmpConfig.V3Config.Notify) > 0 {
				checks.append(t, "TestCheckResourceAttrSet", "snmp_config.v3_config.notify")
				for i, notify := range s.SnmpConfig.V3Config.Notify {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("snmp_config.v3_config.notify.%d.name", i), notify.Name)
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("snmp_config.v3_config.notify.%d.tag", i), notify.Tag)
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("snmp_config.v3_config.notify.%d.type", i), notify.NotifyType)
				}
			}
			if len(s.SnmpConfig.V3Config.TargetAddress) > 0 {
				checks.append(t, "TestCheckResourceAttrSet", "snmp_config.v3_config.target_address")
				for i, target := range s.SnmpConfig.V3Config.TargetAddress {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("snmp_config.v3_config.target_address.%d.address", i), target.Address)
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("snmp_config.v3_config.target_address.%d.address_mask", i), target.AddressMask)
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("snmp_config.v3_config.target_address.%d.target_address_name", i), target.TargetAddressName)
				}
			}
			if len(s.SnmpConfig.V3Config.Usm) > 0 {
				checks.append(t, "TestCheckResourceAttrSet", "snmp_config.v3_config.usm")
				for i, usm := range s.SnmpConfig.V3Config.Usm {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("snmp_config.v3_config.usm.%d.engine_type", i), usm.EngineType)
					if len(usm.Snmpv3Users) > 0 {
						checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("snmp_config.v3_config.usm.%d.users", i))
					}
				}
			}
		}
		if len(s.SnmpConfig.Views) > 0 {
			checks.append(t, "TestCheckResourceAttrSet", "snmp_config.views")
			// Check nested Views
			for i, view := range s.SnmpConfig.Views {
				if view.Include != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("snmp_config.views.%d.include", i), fmt.Sprintf("%t", *view.Include))
				}
				if view.Oid != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("snmp_config.views.%d.oid", i), *view.Oid)
				}
				if view.ViewName != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("snmp_config.views.%d.view_name", i), *view.ViewName)
				}
			}
		}
	}
	if s.StpConfig != nil {
		checks.append(t, "TestCheckResourceAttrSet", "stp_config")
		// Check nested attributes of StpConfig
		if s.StpConfig.BridgePriority != nil {
			checks.append(t, "TestCheckResourceAttr", "stp_config.bridge_priority", *s.StpConfig.BridgePriority)
		}
	}
	if s.SwitchMgmt != nil {
		checks.append(t, "TestCheckResourceAttrSet", "switch_mgmt")
		// Check nested attributes of SwitchMgmt
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
			checks.append(t, "TestCheckResourceAttrSet", "switch_mgmt.local_accounts")
			// Check nested LocalAccounts
			for key, account := range s.SwitchMgmt.LocalAccounts {
				if account.Password != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("switch_mgmt.local_accounts.%s.password", key), *account.Password)
				}
				if account.Role != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("switch_mgmt.local_accounts.%s.role", key), *account.Role)
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
			checks.append(t, "TestCheckResourceAttrSet", "switch_mgmt.protect_re")
			// Check nested ProtectRe
			if len(s.SwitchMgmt.ProtectRe.AllowedServices) > 0 {
				checks.append(t, "TestCheckResourceAttrSet", "switch_mgmt.protect_re.allowed_services")
				// Check individual allowed service values
				for i, service := range s.SwitchMgmt.ProtectRe.AllowedServices {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("switch_mgmt.protect_re.allowed_services.%d", i), service)
				}
			}
			if len(s.SwitchMgmt.ProtectRe.Custom) > 0 {
				checks.append(t, "TestCheckResourceAttrSet", "switch_mgmt.protect_re.custom")
				// Check nested Custom
				for i, custom := range s.SwitchMgmt.ProtectRe.Custom {
					if custom.PortRange != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("switch_mgmt.protect_re.custom.%d.port_range", i), *custom.PortRange)
					}
					if custom.Protocol != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("switch_mgmt.protect_re.custom.%d.protocol", i), *custom.Protocol)
					}
					if len(custom.Subnets) > 0 {
						checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("switch_mgmt.protect_re.custom.%d.subnets", i))
						// Check individual subnet values
						for j, subnet := range custom.Subnets {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("switch_mgmt.protect_re.custom.%d.subnets.%d", i, j), subnet)
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
				checks.append(t, "TestCheckResourceAttrSet", "switch_mgmt.protect_re.trusted_hosts")
				// Check individual trusted host values
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
			checks.append(t, "TestCheckResourceAttrSet", "switch_mgmt.tacacs")
			// Check nested Tacacs
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
				checks.append(t, "TestCheckResourceAttrSet", "switch_mgmt.tacacs.acct_servers")
				// Check nested TacacctServers
				for i, server := range s.SwitchMgmt.Tacacs.TacacctServers {
					if server.Host != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("switch_mgmt.tacacs.acct_servers.%d.host", i), *server.Host)
					}
					if server.Port != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("switch_mgmt.tacacs.acct_servers.%d.port", i), *server.Port)
					}
					if server.Secret != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("switch_mgmt.tacacs.acct_servers.%d.secret", i), *server.Secret)
					}
					if server.Timeout != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("switch_mgmt.tacacs.acct_servers.%d.timeout", i), fmt.Sprintf("%d", *server.Timeout))
					}
				}
			}
			if len(s.SwitchMgmt.Tacacs.TacplusServers) > 0 {
				checks.append(t, "TestCheckResourceAttrSet", "switch_mgmt.tacacs.tacplus_servers")
				// Check nested TacplusServers
				for i, server := range s.SwitchMgmt.Tacacs.TacplusServers {
					if server.Host != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("switch_mgmt.tacacs.tacplus_servers.%d.host", i), *server.Host)
					}
					if server.Port != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("switch_mgmt.tacacs.tacplus_servers.%d.port", i), *server.Port)
					}
					if server.Secret != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("switch_mgmt.tacacs.tacplus_servers.%d.secret", i), *server.Secret)
					}
					if server.Timeout != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("switch_mgmt.tacacs.tacplus_servers.%d.timeout", i), fmt.Sprintf("%d", *server.Timeout))
					}
				}
			}
		}
		if s.SwitchMgmt.UseMxedgeProxy != nil {
			checks.append(t, "TestCheckResourceAttr", "switch_mgmt.use_mxedge_proxy", fmt.Sprintf("%t", *s.SwitchMgmt.UseMxedgeProxy))
		}
	}
	if s.UseRouterIdAsSourceIp != nil {
		checks.append(t, "TestCheckResourceAttr", "use_router_id_as_source_ip", fmt.Sprintf("%t", *s.UseRouterIdAsSourceIp))
	}
	if len(s.Vars) > 0 {
		checks.append(t, "TestCheckResourceAttrSet", "vars")
		// Check nested attributes of Vars (simple string map)
		for key, value := range s.Vars {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vars.%s", key), value)
		}
	}
	if s.VirtualChassis != nil {
		checks.append(t, "TestCheckResourceAttrSet", "virtual_chassis")
		// Check nested attributes of VirtualChassis
		if s.VirtualChassis.Preprovisioned != nil {
			checks.append(t, "TestCheckResourceAttr", "virtual_chassis.preprovisioned", fmt.Sprintf("%t", *s.VirtualChassis.Preprovisioned))
		}
		if len(s.VirtualChassis.Members) > 0 {
			checks.append(t, "TestCheckResourceAttrSet", "virtual_chassis.members")
			// Check nested Members
			for i, member := range s.VirtualChassis.Members {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("virtual_chassis.members.%d.mac", i), member.Mac)
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("virtual_chassis.members.%d.member_id", i), fmt.Sprintf("%d", member.MemberId))
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("virtual_chassis.members.%d.vc_role", i), member.VcRole)
			}
		}
	}
	if s.VrfConfig != nil {
		checks.append(t, "TestCheckResourceAttrSet", "vrf_config")
		// Check nested attributes of VrfConfig
		if s.VrfConfig.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "vrf_config.enabled", fmt.Sprintf("%t", *s.VrfConfig.Enabled))
		}
	}
	if len(s.VrfInstances) > 0 {
		checks.append(t, "TestCheckResourceAttrSet", "vrf_instances")
		// Check nested attributes of VrfInstances
		for key, vrf := range s.VrfInstances {
			if vrf.EvpnAutoLoopbackSubnet != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vrf_instances.%s.evpn_auto_loopback_subnet", key), *vrf.EvpnAutoLoopbackSubnet)
			}
			if vrf.EvpnAutoLoopbackSubnet6 != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vrf_instances.%s.evpn_auto_loopback_subnet6", key), *vrf.EvpnAutoLoopbackSubnet6)
			}
			if len(vrf.Networks) > 0 {
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("vrf_instances.%s.networks", key))
			}
			if len(vrf.VrfExtraRoutes) > 0 {
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("vrf_instances.%s.extra_routes", key))
				// Check nested VrfExtraRoutes
				for routeKey, route := range vrf.VrfExtraRoutes {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vrf_instances.%s.extra_routes.%s.via", key, routeKey), route.Via)
				}
			}
			if len(vrf.VrfExtraRoutes6) > 0 {
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("vrf_instances.%s.extra_routes6", key))
				// Check nested VrfExtraRoutes6
				for routeKey, route := range vrf.VrfExtraRoutes6 {
					if route.Via != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vrf_instances.%s.extra_routes6.%s.via", key, routeKey), *route.Via)
					}
				}
			}
		}
	}
	if s.VrrpConfig != nil {
		checks.append(t, "TestCheckResourceAttrSet", "vrrp_config")
		// Check nested attributes of VrrpConfig
		if s.VrrpConfig.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "vrrp_config.enabled", fmt.Sprintf("%t", *s.VrrpConfig.Enabled))
		}
		if len(s.VrrpConfig.Groups) > 0 {
			checks.append(t, "TestCheckResourceAttrSet", "vrrp_config.groups")
			// Check nested Groups
			for groupKey, group := range s.VrrpConfig.Groups {
				if group.Preempt != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vrrp_config.groups.%s.preempt", groupKey), fmt.Sprintf("%t", *group.Preempt))
				}
				if group.Priority != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vrrp_config.groups.%s.priority", groupKey), fmt.Sprintf("%d", *group.Priority))
				}
			}
		}
	}
	if s.X != nil {
		checks.append(t, "TestCheckResourceAttr", "x", fmt.Sprintf("%g", *s.X))
	}
	if s.Y != nil {
		checks.append(t, "TestCheckResourceAttr", "y", fmt.Sprintf("%g", *s.Y))
	}

	return checks
}
