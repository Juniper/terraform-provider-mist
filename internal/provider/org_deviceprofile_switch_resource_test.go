package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_deviceprofile_switch"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgDeviceprofileSwitchModel(t *testing.T) {
	type testStep struct {
		config OrgDeviceprofileSwitchModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgDeviceprofileSwitchModel{
						Name:  "test_switch_profile",
						OrgId: GetTestOrgId(),
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/org_deviceprofile_switch_resource/org_deviceprofile_switch_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b)
	fixtures := strings.Split(str, "␞")
	for i, fixture := range fixtures {
		var FixtureOrgDeviceprofileSwitchModel OrgDeviceprofileSwitchModel
		err = hcl.Decode(&FixtureOrgDeviceprofileSwitchModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		FixtureOrgDeviceprofileSwitchModel.OrgId = GetTestOrgId()
		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureOrgDeviceprofileSwitchModel,
				},
			},
		}
	}

	resourceType := "org_deviceprofile_switch"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_org_deviceprofile_switch.OrgDeviceprofileSwitchResourceSchema(t.Context()).Attributes)
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

func (s *OrgDeviceprofileSwitchModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType)+"."+tName, tracker)

	// Required parameters
	checks.append(t, "TestCheckResourceAttrSet", "org_id")
	checks.append(t, "TestCheckResourceAttr", "name", s.Name)

	// Computed
	checks.append(t, "TestCheckResourceAttrSet", "id")
	checks.append(t, "TestCheckResourceAttr", "type", s.Type)

	// acl_policies
	if len(s.AclPolicies) > 0 {
		checks.append(t, "TestCheckResourceAttr", "acl_policies.#", fmt.Sprintf("%d", len(s.AclPolicies)))
		for i, policy := range s.AclPolicies {
			if policy.Name != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acl_policies.%d.name", i), *policy.Name)
			}
			if len(policy.SrcTags) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acl_policies.%d.src_tags.#", i), fmt.Sprintf("%d", len(policy.SrcTags)))
				for j, tag := range policy.SrcTags {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acl_policies.%d.src_tags.%d", i, j), tag)
				}
			}
			if len(policy.Actions) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acl_policies.%d.actions.#", i), fmt.Sprintf("%d", len(policy.Actions)))
				for j, action := range policy.Actions {
					if action.Action != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acl_policies.%d.actions.%d.action", i, j), *action.Action)
					}
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acl_policies.%d.actions.%d.dst_tag", i, j), action.DstTag)
				}
			}
		}
	}

	// acl_tags
	if len(s.AclTags) > 0 {
		checks.append(t, "TestCheckResourceAttr", "acl_tags.%", fmt.Sprintf("%d", len(s.AclTags)))
		for key, tag := range s.AclTags {
			if len(tag.EtherTypes) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acl_tags.%s.ether_types.#", key), fmt.Sprintf("%d", len(tag.EtherTypes)))
				for i, etherType := range tag.EtherTypes {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acl_tags.%s.ether_types.%d", key, i), etherType)
				}
			}
			if tag.GbpTag != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acl_tags.%s.gbp_tag", key), fmt.Sprintf("%d", *tag.GbpTag))
			}
			if len(tag.Macs) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acl_tags.%s.macs.#", key), fmt.Sprintf("%d", len(tag.Macs)))
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
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acl_tags.%s.specs.#", key), fmt.Sprintf("%d", len(tag.Specs)))
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
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acl_tags.%s.subnets.#", key), fmt.Sprintf("%d", len(tag.Subnets)))
				for i, subnet := range tag.Subnets {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acl_tags.%s.subnets.%d", key, i), subnet)
				}
			}
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("acl_tags.%s.type", key), tag.AclTagsType)
		}
	}

	// additional_config_cmds
	if len(s.AdditionalConfigCmds) > 0 {
		checks.append(t, "TestCheckResourceAttr", "additional_config_cmds.#", fmt.Sprintf("%d", len(s.AdditionalConfigCmds)))
		for i, cmd := range s.AdditionalConfigCmds {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("additional_config_cmds.%d", i), cmd)
		}
	}

	// dhcp_snooping
	if s.DhcpSnooping != nil {
		checks.append(t, "TestCheckResourceAttrSet", "dhcp_snooping.%")
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
			for i, network := range s.DhcpSnooping.Networks {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcp_snooping.networks.%d", i), network)
			}
		}
	}

	// dhcpd_config
	if s.DhcpdConfig != nil {
		checks.append(t, "TestCheckResourceAttrSet", "dhcpd_config.%")
		if s.DhcpdConfig.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "dhcpd_config.enabled", fmt.Sprintf("%t", *s.DhcpdConfig.Enabled))
		}
		if len(s.DhcpdConfig.Config) > 0 {
			checks.append(t, "TestCheckResourceAttr", "dhcpd_config.config.%", fmt.Sprintf("%d", len(s.DhcpdConfig.Config)))
			for key, config := range s.DhcpdConfig.Config {
				if len(config.DnsServers) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.dns_servers.#", key), fmt.Sprintf("%d", len(config.DnsServers)))
					for i, server := range config.DnsServers {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.dns_servers.%d", key, i), server)
					}
				}
				if len(config.DnsSuffix) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.dns_suffix.#", key), fmt.Sprintf("%d", len(config.DnsSuffix)))
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
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.servers.#", key), fmt.Sprintf("%d", len(config.Servers)))
					for i, server := range config.Servers {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.servers.%d", key, i), server)
					}
				}
				if len(config.Servers6) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.servers6.#", key), fmt.Sprintf("%d", len(config.Servers6)))
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
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.fixed_bindings.%%", key), fmt.Sprintf("%d", len(config.FixedBindings)))
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
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.options.%%", key), fmt.Sprintf("%d", len(config.Options)))
				}
				if len(config.VendorEncapsulated) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dhcpd_config.config.%s.vendor_encapsulated.%%", key), fmt.Sprintf("%d", len(config.VendorEncapsulated)))
				}
			}
		}
	}

	// dns_servers
	if len(s.DnsServers) > 0 {
		checks.append(t, "TestCheckResourceAttr", "dns_servers.#", fmt.Sprintf("%d", len(s.DnsServers)))
		for i, server := range s.DnsServers {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dns_servers.%d", i), server)
		}
	}

	// dns_suffix
	if len(s.DnsSuffix) > 0 {
		checks.append(t, "TestCheckResourceAttr", "dns_suffix.#", fmt.Sprintf("%d", len(s.DnsSuffix)))
		for i, suffix := range s.DnsSuffix {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("dns_suffix.%d", i), suffix)
		}
	}

	// evpn_config (unique to deviceprofile_switch)
	if s.EvpnConfig != nil {
		checks.append(t, "TestCheckResourceAttrSet", "evpn_config.%")
		if s.EvpnConfig.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "evpn_config.enabled", fmt.Sprintf("%t", *s.EvpnConfig.Enabled))
		}
		if s.EvpnConfig.Role != nil {
			checks.append(t, "TestCheckResourceAttr", "evpn_config.role", *s.EvpnConfig.Role)
		}
	}

	// extra_routes
	if len(s.ExtraRoutes) > 0 {
		checks.append(t, "TestCheckResourceAttr", "extra_routes.%", fmt.Sprintf("%d", len(s.ExtraRoutes)))
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
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("extra_routes.%s.next_qualified.%%", key), fmt.Sprintf("%d", len(route.NextQualified)))
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

	// extra_routes6
	if len(s.ExtraRoutes6) > 0 {
		checks.append(t, "TestCheckResourceAttr", "extra_routes6.%", fmt.Sprintf("%d", len(s.ExtraRoutes6)))
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
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("extra_routes6.%s.next_qualified.%%", key), fmt.Sprintf("%d", len(route.NextQualified)))
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

	// iot_config (unique to deviceprofile_switch)
	if len(s.IotConfig) > 0 {
		checks.append(t, "TestCheckResourceAttr", "iot_config.%", fmt.Sprintf("%d", len(s.IotConfig)))
		for key, iot := range s.IotConfig {
			if iot.AlarmClass != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("iot_config.%s.alarm_class", key), *iot.AlarmClass)
			}
			if iot.Enabled != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("iot_config.%s.enabled", key), fmt.Sprintf("%t", *iot.Enabled))
			}
			if iot.InputSrc != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("iot_config.%s.input_src", key), *iot.InputSrc)
			}
			if iot.Name != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("iot_config.%s.name", key), *iot.Name)
			}
		}
	}

	// ip_config
	if s.IpConfig != nil {
		checks.append(t, "TestCheckResourceAttrSet", "ip_config.%")
		if len(s.IpConfig.Dns) > 0 {
			checks.append(t, "TestCheckResourceAttr", "ip_config.dns.#", fmt.Sprintf("%d", len(s.IpConfig.Dns)))
			for i, dns := range s.IpConfig.Dns {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ip_config.dns.%d", i), dns)
			}
		}
		if len(s.IpConfig.DnsSuffix) > 0 {
			checks.append(t, "TestCheckResourceAttr", "ip_config.dns_suffix.#", fmt.Sprintf("%d", len(s.IpConfig.DnsSuffix)))
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

	// mist_nac
	if s.MistNac != nil {
		checks.append(t, "TestCheckResourceAttrSet", "mist_nac.%")
		if s.MistNac.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "mist_nac.enabled", fmt.Sprintf("%t", *s.MistNac.Enabled))
		}
		if s.MistNac.Network != nil {
			checks.append(t, "TestCheckResourceAttr", "mist_nac.network", *s.MistNac.Network)
		}
	}

	// networks
	if len(s.Networks) > 0 {
		checks.append(t, "TestCheckResourceAttr", "networks.%", fmt.Sprintf("%d", len(s.Networks)))
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

	// ntp_servers
	if len(s.NtpServers) > 0 {
		checks.append(t, "TestCheckResourceAttr", "ntp_servers.#", fmt.Sprintf("%d", len(s.NtpServers)))
		for i, server := range s.NtpServers {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ntp_servers.%d", i), server)
		}
	}

	// oob_ip_config
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

	// ospf_areas
	if len(s.OspfAreas) > 0 {
		checks.append(t, "TestCheckResourceAttr", "ospf_areas.%", fmt.Sprintf("%d", len(s.OspfAreas)))
		for key, area := range s.OspfAreas {
			if area.IncludeLoopback != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ospf_areas.%s.include_loopback", key), fmt.Sprintf("%t", *area.IncludeLoopback))
			}
			if area.OspfAreasType != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ospf_areas.%s.type", key), *area.OspfAreasType)
			}
			if len(area.OspfNetworks) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ospf_areas.%s.networks.%%", key), fmt.Sprintf("%d", len(area.OspfNetworks)))
				for netKey, net := range area.OspfNetworks {
					if net.AuthType != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ospf_areas.%s.networks.%s.auth_type", key, netKey), *net.AuthType)
					}
					if net.BfdMinimumInterval != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ospf_areas.%s.networks.%s.bfd_minimum_interval", key, netKey), fmt.Sprintf("%d", *net.BfdMinimumInterval))
					}
					if net.DeadInterval != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ospf_areas.%s.networks.%s.dead_interval", key, netKey), fmt.Sprintf("%d", *net.DeadInterval))
					}
					if net.HelloInterval != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ospf_areas.%s.networks.%s.hello_interval", key, netKey), fmt.Sprintf("%d", *net.HelloInterval))
					}
					if net.InterfaceType != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ospf_areas.%s.networks.%s.interface_type", key, netKey), *net.InterfaceType)
					}
					if net.Metric != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ospf_areas.%s.networks.%s.metric", key, netKey), fmt.Sprintf("%d", *net.Metric))
					}
					if net.Passive != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ospf_areas.%s.networks.%s.passive", key, netKey), fmt.Sprintf("%t", *net.Passive))
					}
				}
			}
		}
	}

	// other_ip_configs
	if len(s.OtherIpConfigs) > 0 {
		checks.append(t, "TestCheckResourceAttr", "other_ip_configs.%", fmt.Sprintf("%d", len(s.OtherIpConfigs)))
		for key, cfg := range s.OtherIpConfigs {
			if cfg.EvpnAnycast != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("other_ip_configs.%s.evpn_anycast", key), fmt.Sprintf("%t", *cfg.EvpnAnycast))
			}
			if cfg.Ip != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("other_ip_configs.%s.ip", key), *cfg.Ip)
			}
			if cfg.Ip6 != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("other_ip_configs.%s.ip6", key), *cfg.Ip6)
			}
			if cfg.Netmask != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("other_ip_configs.%s.netmask", key), *cfg.Netmask)
			}
			if cfg.Netmask6 != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("other_ip_configs.%s.netmask6", key), *cfg.Netmask6)
			}
			if cfg.OtherIpConfigsType != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("other_ip_configs.%s.type", key), *cfg.OtherIpConfigsType)
			}
			if cfg.Type6 != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("other_ip_configs.%s.type6", key), *cfg.Type6)
			}
		}
	}

	// port_config
	if len(s.PortConfig) > 0 {
		checks.append(t, "TestCheckResourceAttr", "port_config.%", fmt.Sprintf("%d", len(s.PortConfig)))
		for key, port := range s.PortConfig {
			if port.AeDisableLacp != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ae_disable_lacp", key), fmt.Sprintf("%t", *port.AeDisableLacp))
			}
			if port.AeIdx != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ae_idx", key), fmt.Sprintf("%d", *port.AeIdx))
			}
			if port.AeLacpForceUp != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ae_lacp_force_up", key), fmt.Sprintf("%t", *port.AeLacpForceUp))
			}
			if port.AeLacpPassive != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.ae_lacp_passive", key), fmt.Sprintf("%t", *port.AeLacpPassive))
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
			if port.Esilag != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.esilag", key), fmt.Sprintf("%t", *port.Esilag))
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
			if port.NoLocalOverwrite != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_config.%s.no_local_overwrite", key), fmt.Sprintf("%t", *port.NoLocalOverwrite))
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

	// port_mirroring
	if len(s.PortMirroring) > 0 {
		checks.append(t, "TestCheckResourceAttr", "port_mirroring.%", fmt.Sprintf("%d", len(s.PortMirroring)))
		for key, mirror := range s.PortMirroring {
			if len(mirror.InputNetworksIngress) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_mirroring.%s.input_networks_ingress.#", key), fmt.Sprintf("%d", len(mirror.InputNetworksIngress)))
				for i, network := range mirror.InputNetworksIngress {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_mirroring.%s.input_networks_ingress.%d", key, i), network)
				}
			}
			if len(mirror.InputPortIdsEgress) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_mirroring.%s.input_port_ids_egress.#", key), fmt.Sprintf("%d", len(mirror.InputPortIdsEgress)))
				for i, portId := range mirror.InputPortIdsEgress {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_mirroring.%s.input_port_ids_egress.%d", key, i), portId)
				}
			}
			if len(mirror.InputPortIdsIngress) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_mirroring.%s.input_port_ids_ingress.#", key), fmt.Sprintf("%d", len(mirror.InputPortIdsIngress)))
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

	// port_usages
	if len(s.PortUsages) > 0 {
		checks.append(t, "TestCheckResourceAttr", "port_usages.%", fmt.Sprintf("%d", len(s.PortUsages)))
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
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.networks.#", key), fmt.Sprintf("%d", len(usage.Networks)))
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
			if usage.PoeKeepStateWhenReboot != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.poe_keep_state_when_reboot", key), fmt.Sprintf("%t", *usage.PoeKeepStateWhenReboot))
			}
			if usage.Speed != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.speed", key), *usage.Speed)
			}
			if len(usage.Rules) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.rules.#", key), fmt.Sprintf("%d", len(usage.Rules)))
				for i, rule := range usage.Rules {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.rules.%d.src", key, i), rule.Src)
					if rule.Equals != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.rules.%d.equals", key, i), *rule.Equals)
					}
					if len(rule.EqualsAny) > 0 {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.rules.%d.equals_any.#", key, i), fmt.Sprintf("%d", len(rule.EqualsAny)))
						for j, equalsAnyValue := range rule.EqualsAny {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.rules.%d.equals_any.%d", key, i, j), equalsAnyValue)
						}
					}
					if rule.Usage != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.rules.%d.usage", key, i), *rule.Usage)
					}
				}
			}
			if usage.ServerFailNetwork != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.server_fail_network", key), *usage.ServerFailNetwork)
			}
			if usage.ServerFailRetryInterval != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.server_fail_retry_interval", key), fmt.Sprintf("%d", *usage.ServerFailRetryInterval))
			}
			if usage.StormControl != nil {
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("port_usages.%s.storm_control.%%", key))
			}
			if usage.StpDisable != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.stp_disable", key), fmt.Sprintf("%t", *usage.StpDisable))
			}
			if usage.StpRequired != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.stp_required", key), fmt.Sprintf("%t", *usage.StpRequired))
			}
			if usage.VoipNetwork != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("port_usages.%s.voip_network", key), *usage.VoipNetwork)
			}
		}
	}

	// radius_config
	if s.RadiusConfig != nil {
		checks.append(t, "TestCheckResourceAttrSet", "radius_config.%")
		if s.RadiusConfig.AcctImmediateUpdate != nil {
			checks.append(t, "TestCheckResourceAttr", "radius_config.acct_immediate_update", fmt.Sprintf("%t", *s.RadiusConfig.AcctImmediateUpdate))
		}
		if s.RadiusConfig.AcctInterimInterval != nil {
			checks.append(t, "TestCheckResourceAttr", "radius_config.acct_interim_interval", fmt.Sprintf("%d", *s.RadiusConfig.AcctInterimInterval))
		}
		if len(s.RadiusConfig.AcctServers) > 0 {
			checks.append(t, "TestCheckResourceAttr", "radius_config.acct_servers.#", fmt.Sprintf("%d", len(s.RadiusConfig.AcctServers)))
			for i, server := range s.RadiusConfig.AcctServers {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radius_config.acct_servers.%d.host", i), server.Host)
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
			checks.append(t, "TestCheckResourceAttr", "radius_config.auth_servers.#", fmt.Sprintf("%d", len(s.RadiusConfig.AuthServers)))
			for i, server := range s.RadiusConfig.AuthServers {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radius_config.auth_servers.%d.host", i), server.Host)
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

	// remote_syslog
	if s.RemoteSyslog != nil {
		checks.append(t, "TestCheckResourceAttrSet", "remote_syslog.%")
		if s.RemoteSyslog.Archive != nil {
			checks.append(t, "TestCheckResourceAttrSet", "remote_syslog.archive.%")
			if s.RemoteSyslog.Archive.Files != nil {
				checks.append(t, "TestCheckResourceAttr", "remote_syslog.archive.files", *s.RemoteSyslog.Archive.Files)
			}
			if s.RemoteSyslog.Archive.Size != nil {
				checks.append(t, "TestCheckResourceAttr", "remote_syslog.archive.size", *s.RemoteSyslog.Archive.Size)
			}
		}
		if len(s.RemoteSyslog.Cacerts) > 0 {
			checks.append(t, "TestCheckResourceAttr", "remote_syslog.cacerts.#", fmt.Sprintf("%d", len(s.RemoteSyslog.Cacerts)))
			for i, cacert := range s.RemoteSyslog.Cacerts {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.cacerts.%d", i), cacert)
			}
		}
		if s.RemoteSyslog.Console != nil {
			checks.append(t, "TestCheckResourceAttrSet", "remote_syslog.console.%")
			if len(s.RemoteSyslog.Console.Contents) > 0 {
				checks.append(t, "TestCheckResourceAttr", "remote_syslog.console.contents.#", fmt.Sprintf("%d", len(s.RemoteSyslog.Console.Contents)))
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
			checks.append(t, "TestCheckResourceAttr", "remote_syslog.files.#", fmt.Sprintf("%d", len(s.RemoteSyslog.Files)))
			for i, file := range s.RemoteSyslog.Files {
				if file.Archive != nil {
					checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("remote_syslog.files.%d.archive.%%", i))
				}
				if len(file.Contents) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.files.%d.contents.#", i), fmt.Sprintf("%d", len(file.Contents)))
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
			checks.append(t, "TestCheckResourceAttr", "remote_syslog.servers.#", fmt.Sprintf("%d", len(s.RemoteSyslog.Servers)))
			for i, server := range s.RemoteSyslog.Servers {
				if len(server.Contents) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.servers.%d.contents.#", i), fmt.Sprintf("%d", len(server.Contents)))
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
			checks.append(t, "TestCheckResourceAttr", "remote_syslog.users.#", fmt.Sprintf("%d", len(s.RemoteSyslog.Users)))
			for i, user := range s.RemoteSyslog.Users {
				if len(user.Contents) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("remote_syslog.users.%d.contents.#", i), fmt.Sprintf("%d", len(user.Contents)))
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

	// routing_policies
	if len(s.RoutingPolicies) > 0 {
		checks.append(t, "TestCheckResourceAttr", "routing_policies.%", fmt.Sprintf("%d", len(s.RoutingPolicies)))
		for k, v := range s.RoutingPolicies {
			if len(v.Terms) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routing_policies.%s.terms.#", k), fmt.Sprintf("%d", len(v.Terms)))
				for _, term := range v.Terms {
					termChecks := make(map[string]string)
					termChecks["name"] = term.Name
					if term.RoutingPolicyTermActions != nil {
						if term.RoutingPolicyTermActions.Accept != nil {
							termChecks["routing_policy_term_actions.accept"] = fmt.Sprintf("%t", *term.RoutingPolicyTermActions.Accept)
						}
						if term.RoutingPolicyTermActions.LocalPreference != nil {
							termChecks["routing_policy_term_actions.local_preference"] = *term.RoutingPolicyTermActions.LocalPreference
						}
					}
					if term.Matching != nil {
						if len(term.Matching.Protocol) > 0 {
							termChecks["matching.protocol.#"] = fmt.Sprintf("%d", len(term.Matching.Protocol))
						}
					}
					checks.appendSetNestedCheck(t, fmt.Sprintf("routing_policies.%s.terms.*", k), termChecks)
				}
			}
		}
	}

	// snmp_config
	if s.SnmpConfig != nil {
		checks.append(t, "TestCheckResourceAttrSet", "snmp_config.%")
		if s.SnmpConfig.Contact != nil {
			checks.append(t, "TestCheckResourceAttr", "snmp_config.contact", *s.SnmpConfig.Contact)
		}
		if s.SnmpConfig.Description != nil {
			checks.append(t, "TestCheckResourceAttr", "snmp_config.description", *s.SnmpConfig.Description)
		}
		if s.SnmpConfig.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "snmp_config.enabled", fmt.Sprintf("%t", *s.SnmpConfig.Enabled))
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
		if len(s.SnmpConfig.V2cConfig) > 0 {
			checks.append(t, "TestCheckResourceAttr", "snmp_config.v2c_config.#", fmt.Sprintf("%d", len(s.SnmpConfig.V2cConfig)))
			for i, v2c := range s.SnmpConfig.V2cConfig {
				if v2c.CommunityName != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("snmp_config.v2c_config.%d.community_name", i), *v2c.CommunityName)
				}
				if v2c.Authorization != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("snmp_config.v2c_config.%d.authorization", i), *v2c.Authorization)
				}
			}
		}
		if len(s.SnmpConfig.TrapGroups) > 0 {
			checks.append(t, "TestCheckResourceAttr", "snmp_config.trap_groups.#", fmt.Sprintf("%d", len(s.SnmpConfig.TrapGroups)))
			for i, trap := range s.SnmpConfig.TrapGroups {
				if trap.GroupName != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("snmp_config.trap_groups.%d.group_name", i), *trap.GroupName)
				}
				if trap.Version != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("snmp_config.trap_groups.%d.version", i), *trap.Version)
				}
				if len(trap.Targets) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("snmp_config.trap_groups.%d.targets.#", i), fmt.Sprintf("%d", len(trap.Targets)))
					for j, target := range trap.Targets {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("snmp_config.trap_groups.%d.targets.%d", i, j), target)
					}
				}
			}
		}
	}

	// stp_config
	if s.StpConfig != nil {
		checks.append(t, "TestCheckResourceAttrSet", "stp_config.%")
		if s.StpConfig.BridgePriority != nil {
			checks.append(t, "TestCheckResourceAttr", "stp_config.bridge_priority", *s.StpConfig.BridgePriority)
		}
	}

	// switch_mgmt
	if s.SwitchMgmt != nil {
		checks.append(t, "TestCheckResourceAttrSet", "switch_mgmt.%")
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
			checks.append(t, "TestCheckResourceAttrSet", "switch_mgmt.protect_re.%")
			if s.SwitchMgmt.ProtectRe.Enabled != nil {
				checks.append(t, "TestCheckResourceAttr", "switch_mgmt.protect_re.enabled", fmt.Sprintf("%t", *s.SwitchMgmt.ProtectRe.Enabled))
			}
			if len(s.SwitchMgmt.ProtectRe.AllowedServices) > 0 {
				checks.append(t, "TestCheckResourceAttr", "switch_mgmt.protect_re.allowed_services.#", fmt.Sprintf("%d", len(s.SwitchMgmt.ProtectRe.AllowedServices)))
				for i, service := range s.SwitchMgmt.ProtectRe.AllowedServices {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("switch_mgmt.protect_re.allowed_services.%d", i), service)
				}
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
			checks.append(t, "TestCheckResourceAttrSet", "switch_mgmt.tacacs.%")
			if s.SwitchMgmt.Tacacs.DefaultRole != nil {
				checks.append(t, "TestCheckResourceAttr", "switch_mgmt.tacacs.default_role", *s.SwitchMgmt.Tacacs.DefaultRole)
			}
			if s.SwitchMgmt.Tacacs.Enabled != nil {
				checks.append(t, "TestCheckResourceAttr", "switch_mgmt.tacacs.enabled", fmt.Sprintf("%t", *s.SwitchMgmt.Tacacs.Enabled))
			}
			if s.SwitchMgmt.Tacacs.Network != nil {
				checks.append(t, "TestCheckResourceAttr", "switch_mgmt.tacacs.network", *s.SwitchMgmt.Tacacs.Network)
			}
			if len(s.SwitchMgmt.Tacacs.TacplusServers) > 0 {
				checks.append(t, "TestCheckResourceAttr", "switch_mgmt.tacacs.tacplus_servers.#", fmt.Sprintf("%d", len(s.SwitchMgmt.Tacacs.TacplusServers)))
				for i, server := range s.SwitchMgmt.Tacacs.TacplusServers {
					if server.Host != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("switch_mgmt.tacacs.tacplus_servers.%d.host", i), *server.Host)
					}
					if server.Secret != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("switch_mgmt.tacacs.tacplus_servers.%d.secret", i), *server.Secret)
					}
				}
			}
		}
		if s.SwitchMgmt.UseMxedgeProxy != nil {
			checks.append(t, "TestCheckResourceAttr", "switch_mgmt.use_mxedge_proxy", fmt.Sprintf("%t", *s.SwitchMgmt.UseMxedgeProxy))
		}
	}

	// use_router_id_as_source_ip
	if s.UseRouterIdAsSourceIp != nil {
		checks.append(t, "TestCheckResourceAttr", "use_router_id_as_source_ip", fmt.Sprintf("%t", *s.UseRouterIdAsSourceIp))
	}

	// vrf_config
	if s.VrfConfig != nil {
		checks.append(t, "TestCheckResourceAttrSet", "vrf_config.%")
		if s.VrfConfig.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "vrf_config.enabled", fmt.Sprintf("%t", *s.VrfConfig.Enabled))
		}
	}

	// vrf_instances
	if len(s.VrfInstances) > 0 {
		checks.append(t, "TestCheckResourceAttr", "vrf_instances.%", fmt.Sprintf("%d", len(s.VrfInstances)))
		for key, vrf := range s.VrfInstances {
			if vrf.EvpnAutoLoopbackSubnet != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vrf_instances.%s.evpn_auto_loopback_subnet", key), *vrf.EvpnAutoLoopbackSubnet)
			}
			if vrf.EvpnAutoLoopbackSubnet6 != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vrf_instances.%s.evpn_auto_loopback_subnet6", key), *vrf.EvpnAutoLoopbackSubnet6)
			}
			if len(vrf.Networks) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vrf_instances.%s.networks.#", key), fmt.Sprintf("%d", len(vrf.Networks)))
			}
			if len(vrf.VrfExtraRoutes) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vrf_instances.%s.extra_routes.%%", key), fmt.Sprintf("%d", len(vrf.VrfExtraRoutes)))
				for routeKey, route := range vrf.VrfExtraRoutes {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vrf_instances.%s.extra_routes.%s.via", key, routeKey), route.Via)
				}
			}
			if len(vrf.VrfExtraRoutes6) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vrf_instances.%s.extra_routes6.%%", key), fmt.Sprintf("%d", len(vrf.VrfExtraRoutes6)))
				for routeKey, route := range vrf.VrfExtraRoutes6 {
					if route.Via != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vrf_instances.%s.extra_routes6.%s.via", key, routeKey), *route.Via)
					}
				}
			}
		}
	}

	// vrrp_config
	if s.VrrpConfig != nil {
		checks.append(t, "TestCheckResourceAttrSet", "vrrp_config.%")
		if s.VrrpConfig.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "vrrp_config.enabled", fmt.Sprintf("%t", *s.VrrpConfig.Enabled))
		}
		if len(s.VrrpConfig.Groups) > 0 {
			checks.append(t, "TestCheckResourceAttr", "vrrp_config.groups.%", fmt.Sprintf("%d", len(s.VrrpConfig.Groups)))
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

	return checks
}
