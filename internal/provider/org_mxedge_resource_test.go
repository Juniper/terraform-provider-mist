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

func TestOrgMxedgeModel(t *testing.T) {
	type testStep struct {
		config OrgMxedgeModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgMxedgeModel{
						Name:  "Test Org Mxedge",
						Model: "VM",
						OrgId: GetTestOrgId(),
					},
				},
			},
		},
	}

	fixtures, err := os.ReadFile("fixtures/org_mxedge_resource/org_mxedge_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	for i, fixture := range strings.Split(string(fixtures), "␞") {
		fixtureOrgMxedgeModel := OrgMxedgeModel{}
		err = hcl.Decode(&fixtureOrgMxedgeModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		// Need to set org_id to required field as it is used in the url for the resource
		fixtureOrgMxedgeModel.OrgId = GetTestOrgId()
		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: fixtureOrgMxedgeModel,
				},
			},
		}
	}

	resourceType := "org_mxedge"
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
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

func (o *OrgMxedgeModel) testChecks(t testing.TB, rType, tName string) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + tName)

	// Check required fields
	checks.append(t, "TestCheckResourceAttrSet", "org_id")
	checks.append(t, "TestCheckResourceAttr", "name", o.Name)
	checks.append(t, "TestCheckResourceAttr", "model", o.Model)

	// Check optional basic fields
	if o.Note != nil {
		checks.append(t, "TestCheckResourceAttr", "note", *o.Note)
	}
	if o.Magic != nil {
		checks.append(t, "TestCheckResourceAttr", "magic", *o.Magic)
	}
	if o.SiteId != nil {
		checks.append(t, "TestCheckResourceAttr", "site_id", *o.SiteId)
	}
	if o.ForSite != nil {
		checks.append(t, "TestCheckResourceAttr", "for_site", fmt.Sprintf("%t", *o.ForSite))
	}
	if o.MxagentRegistered != nil {
		checks.append(t, "TestCheckResourceAttr", "mxagent_registered", fmt.Sprintf("%t", *o.MxagentRegistered))
	}
	if o.TuntermRegistered != nil {
		checks.append(t, "TestCheckResourceAttr", "tunterm_registered", fmt.Sprintf("%t", *o.TuntermRegistered))
	}
	if o.MxclusterId != nil {
		checks.append(t, "TestCheckResourceAttr", "mxcluster_id", *o.MxclusterId)
	}

	// Check NTP servers
	if len(o.NtpServers) > 0 {
		checks.append(t, "TestCheckResourceAttr", "ntp_servers.#", fmt.Sprintf("%d", len(o.NtpServers)))
		for i, server := range o.NtpServers {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ntp_servers.%d", i), server)
		}
	}

	// Check services
	if len(o.Services) > 0 {
		checks.append(t, "TestCheckResourceAttr", "services.#", fmt.Sprintf("%d", len(o.Services)))
		for i, service := range o.Services {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("services.%d", i), service)
		}
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

	// Check oob_ip_config
	if o.OobIpConfig != nil {
		if o.OobIpConfig.OobIpConfigType != nil {
			checks.append(t, "TestCheckResourceAttr", "oob_ip_config.type", *o.OobIpConfig.OobIpConfigType)
		}
		if o.OobIpConfig.Ip != nil {
			checks.append(t, "TestCheckResourceAttr", "oob_ip_config.ip", *o.OobIpConfig.Ip)
		}
		if o.OobIpConfig.Netmask != nil {
			checks.append(t, "TestCheckResourceAttr", "oob_ip_config.netmask", *o.OobIpConfig.Netmask)
		}
		if o.OobIpConfig.Gateway != nil {
			checks.append(t, "TestCheckResourceAttr", "oob_ip_config.gateway", *o.OobIpConfig.Gateway)
		}
		if len(o.OobIpConfig.Dns) > 0 {
			checks.append(t, "TestCheckResourceAttr", "oob_ip_config.dns.#", fmt.Sprintf("%d", len(o.OobIpConfig.Dns)))
			for i, dns := range o.OobIpConfig.Dns {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("oob_ip_config.dns.%d", i), dns)
			}
		}
		if o.OobIpConfig.Type6 != nil {
			checks.append(t, "TestCheckResourceAttr", "oob_ip_config.type6", *o.OobIpConfig.Type6)
		}
		if o.OobIpConfig.Ip6 != nil {
			checks.append(t, "TestCheckResourceAttr", "oob_ip_config.ip6", *o.OobIpConfig.Ip6)
		}
		if o.OobIpConfig.Netmask6 != nil {
			checks.append(t, "TestCheckResourceAttr", "oob_ip_config.netmask6", *o.OobIpConfig.Netmask6)
		}
		if o.OobIpConfig.Gateway6 != nil {
			checks.append(t, "TestCheckResourceAttr", "oob_ip_config.gateway6", *o.OobIpConfig.Gateway6)
		}
		if o.OobIpConfig.Autoconf6 != nil {
			checks.append(t, "TestCheckResourceAttr", "oob_ip_config.autoconf6", fmt.Sprintf("%t", *o.OobIpConfig.Autoconf6))
		}
		if o.OobIpConfig.Dhcp6 != nil {
			checks.append(t, "TestCheckResourceAttr", "oob_ip_config.dhcp6", fmt.Sprintf("%t", *o.OobIpConfig.Dhcp6))
		}
	}

	// Check proxy
	if o.Proxy != nil && o.Proxy.Url != nil {
		checks.append(t, "TestCheckResourceAttr", "proxy.url", *o.Proxy.Url)
	}

	// Check tunterm_ip_config
	if o.TuntermIpConfig != nil {
		checks.append(t, "TestCheckResourceAttr", "tunterm_ip_config.ip", o.TuntermIpConfig.Ip)
		checks.append(t, "TestCheckResourceAttr", "tunterm_ip_config.netmask", o.TuntermIpConfig.Netmask)
		checks.append(t, "TestCheckResourceAttr", "tunterm_ip_config.gateway", o.TuntermIpConfig.Gateway)
		if o.TuntermIpConfig.Ip6 != nil {
			checks.append(t, "TestCheckResourceAttr", "tunterm_ip_config.ip6", *o.TuntermIpConfig.Ip6)
		}
		if o.TuntermIpConfig.Netmask6 != nil {
			checks.append(t, "TestCheckResourceAttr", "tunterm_ip_config.netmask6", *o.TuntermIpConfig.Netmask6)
		}
		if o.TuntermIpConfig.Gateway6 != nil {
			checks.append(t, "TestCheckResourceAttr", "tunterm_ip_config.gateway6", *o.TuntermIpConfig.Gateway6)
		}
	}

	// Check tunterm_other_ip_configs
	if len(o.TuntermOtherIpConfigs) > 0 {
		checks.append(t, "TestCheckResourceAttr", "tunterm_other_ip_configs.%", fmt.Sprintf("%d", len(o.TuntermOtherIpConfigs)))
		for vlanId, ipConfig := range o.TuntermOtherIpConfigs {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunterm_other_ip_configs.%s.ip", vlanId), ipConfig.Ip)
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunterm_other_ip_configs.%s.netmask", vlanId), ipConfig.Netmask)
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

	// Check tunterm_port_config
	if o.TuntermPortConfig != nil {
		if o.TuntermPortConfig.SeparateUpstreamDownstream != nil {
			checks.append(t, "TestCheckResourceAttr", "tunterm_port_config.separate_upstream_downstream", fmt.Sprintf("%t", *o.TuntermPortConfig.SeparateUpstreamDownstream))
		}
		if o.TuntermPortConfig.UpstreamPortVlanId != nil {
			checks.append(t, "TestCheckResourceAttr", "tunterm_port_config.upstream_port_vlan_id", fmt.Sprintf("%d", *o.TuntermPortConfig.UpstreamPortVlanId))
		}
		if len(o.TuntermPortConfig.UpstreamPorts) > 0 {
			checks.append(t, "TestCheckResourceAttr", "tunterm_port_config.upstream_ports.#", fmt.Sprintf("%d", len(o.TuntermPortConfig.UpstreamPorts)))
			for i, port := range o.TuntermPortConfig.UpstreamPorts {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunterm_port_config.upstream_ports.%d", i), port)
			}
		}
		if len(o.TuntermPortConfig.DownstreamPorts) > 0 {
			checks.append(t, "TestCheckResourceAttr", "tunterm_port_config.downstream_ports.#", fmt.Sprintf("%d", len(o.TuntermPortConfig.DownstreamPorts)))
			for i, port := range o.TuntermPortConfig.DownstreamPorts {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunterm_port_config.downstream_ports.%d", i), port)
			}
		}
	}

	// Check tunterm_switch_config
	if len(o.TuntermSwitchConfig) > 0 {
		checks.append(t, "TestCheckResourceAttr", "tunterm_switch_config.%", fmt.Sprintf("%d", len(o.TuntermSwitchConfig)))
		for portName, switchConfig := range o.TuntermSwitchConfig {
			if switchConfig.PortVlanId != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunterm_switch_config.%s.port_vlan_id", portName), fmt.Sprintf("%d", *switchConfig.PortVlanId))
			}
			if len(switchConfig.VlanIds) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunterm_switch_config.%s.vlan_ids.#", portName), fmt.Sprintf("%d", len(switchConfig.VlanIds)))
				for i, vlanId := range switchConfig.VlanIds {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunterm_switch_config.%s.vlan_ids.%d", portName, i), vlanId)
				}
			}
		}
	}

	// Check tunterm_igmp_snooping_config
	if o.TuntermIgmpSnoopingConfig != nil {
		if o.TuntermIgmpSnoopingConfig.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "tunterm_igmp_snooping_config.enabled", fmt.Sprintf("%t", *o.TuntermIgmpSnoopingConfig.Enabled))
		}
		if len(o.TuntermIgmpSnoopingConfig.VlanIds) > 0 {
			checks.append(t, "TestCheckResourceAttr", "tunterm_igmp_snooping_config.vlan_ids.#", fmt.Sprintf("%d", len(o.TuntermIgmpSnoopingConfig.VlanIds)))
			for i, vlanId := range o.TuntermIgmpSnoopingConfig.VlanIds {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunterm_igmp_snooping_config.vlan_ids.%d", i), fmt.Sprintf("%d", vlanId))
			}
		}
		if o.TuntermIgmpSnoopingConfig.Querier != nil {
			if o.TuntermIgmpSnoopingConfig.Querier.Version != nil {
				checks.append(t, "TestCheckResourceAttr", "tunterm_igmp_snooping_config.querier.version", fmt.Sprintf("%d", *o.TuntermIgmpSnoopingConfig.Querier.Version))
			}
			if o.TuntermIgmpSnoopingConfig.Querier.QueryInterval != nil {
				checks.append(t, "TestCheckResourceAttr", "tunterm_igmp_snooping_config.querier.query_interval", fmt.Sprintf("%d", *o.TuntermIgmpSnoopingConfig.Querier.QueryInterval))
			}
			if o.TuntermIgmpSnoopingConfig.Querier.MaxResponseTime != nil {
				checks.append(t, "TestCheckResourceAttr", "tunterm_igmp_snooping_config.querier.max_response_time", fmt.Sprintf("%d", *o.TuntermIgmpSnoopingConfig.Querier.MaxResponseTime))
			}
			if o.TuntermIgmpSnoopingConfig.Querier.Robustness != nil {
				checks.append(t, "TestCheckResourceAttr", "tunterm_igmp_snooping_config.querier.robustness", fmt.Sprintf("%d", *o.TuntermIgmpSnoopingConfig.Querier.Robustness))
			}
			if o.TuntermIgmpSnoopingConfig.Querier.Mtu != nil {
				checks.append(t, "TestCheckResourceAttr", "tunterm_igmp_snooping_config.querier.mtu", fmt.Sprintf("%d", *o.TuntermIgmpSnoopingConfig.Querier.Mtu))
			}
		}
	}

	// Check tunterm_multicast_config
	if o.TuntermMulticastConfig != nil {
		if o.TuntermMulticastConfig.Mdns != nil {
			if o.TuntermMulticastConfig.Mdns.Enabled != nil {
				checks.append(t, "TestCheckResourceAttr", "tunterm_multicast_config.mdns.enabled", fmt.Sprintf("%t", *o.TuntermMulticastConfig.Mdns.Enabled))
			}
			if len(o.TuntermMulticastConfig.Mdns.VlanIds) > 0 {
				checks.append(t, "TestCheckResourceAttr", "tunterm_multicast_config.mdns.vlan_ids.#", fmt.Sprintf("%d", len(o.TuntermMulticastConfig.Mdns.VlanIds)))
				for i, vlanId := range o.TuntermMulticastConfig.Mdns.VlanIds {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunterm_multicast_config.mdns.vlan_ids.%d", i), vlanId)
				}
			}
		}
		if o.TuntermMulticastConfig.Ssdp != nil {
			if o.TuntermMulticastConfig.Ssdp.Enabled != nil {
				checks.append(t, "TestCheckResourceAttr", "tunterm_multicast_config.ssdp.enabled", fmt.Sprintf("%t", *o.TuntermMulticastConfig.Ssdp.Enabled))
			}
			if len(o.TuntermMulticastConfig.Ssdp.VlanIds) > 0 {
				checks.append(t, "TestCheckResourceAttr", "tunterm_multicast_config.ssdp.vlan_ids.#", fmt.Sprintf("%d", len(o.TuntermMulticastConfig.Ssdp.VlanIds)))
				for i, vlanId := range o.TuntermMulticastConfig.Ssdp.VlanIds {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("tunterm_multicast_config.ssdp.vlan_ids.%d", i), vlanId)
				}
			}
		}
	}

	// Check versions
	if o.Versions != nil {
		if o.Versions.Mxagent != nil {
			checks.append(t, "TestCheckResourceAttr", "versions.mxagent", *o.Versions.Mxagent)
		}
		if o.Versions.Tunterm != nil {
			checks.append(t, "TestCheckResourceAttr", "versions.tunterm", *o.Versions.Tunterm)
		}
	}

	return checks
}
