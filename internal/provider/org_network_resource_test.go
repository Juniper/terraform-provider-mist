package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_network"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgNetworkModel(t *testing.T) {
	type testStep struct {
		config OrgNetworkModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgNetworkModel{
						OrgId:  GetTestOrgId(),
						Name:   "test-network1",
						Subnet: "10.4.0.0/24",
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/org_network_resource/org_network_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		FixtureOrgNetworkModel := OrgNetworkModel{}

		err = hcl.Decode(&FixtureOrgNetworkModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		FixtureOrgNetworkModel.OrgId = GetTestOrgId()

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureOrgNetworkModel,
				},
			},
		}
	}

	resourceType := "org_network"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_org_network.OrgNetworkResourceSchema(t.Context()).Attributes)
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
				t.Logf("\n// ------ begin checks for %s ------\n%s// -------- end config for %s ------\n\n", stepName, chkLog, stepName)

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
	tracker.FieldCoverageReport(t)
}

func (o *OrgNetworkModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + tName)
	checks.SetTracker(tracker)

	// Check required fields
	checks.append(t, "TestCheckResourceAttr", "org_id", o.OrgId)
	checks.append(t, "TestCheckResourceAttr", "name", o.Name)
	checks.append(t, "TestCheckResourceAttr", "subnet", o.Subnet)

	// Optional basic attributes
	if o.DisallowMistServices != nil {
		checks.append(t, "TestCheckResourceAttr", "disallow_mist_services", fmt.Sprintf("%t", *o.DisallowMistServices))
	}
	if o.Gateway != nil {
		checks.append(t, "TestCheckResourceAttr", "gateway", *o.Gateway)
	}
	if o.Gateway6 != nil {
		checks.append(t, "TestCheckResourceAttr", "gateway6", *o.Gateway6)
	}
	if o.Isolation != nil {
		checks.append(t, "TestCheckResourceAttr", "isolation", fmt.Sprintf("%t", *o.Isolation))
	}
	if o.Subnet6 != nil {
		checks.append(t, "TestCheckResourceAttr", "subnet6", *o.Subnet6)
	}
	if o.VlanId != nil {
		checks.append(t, "TestCheckResourceAttr", "vlan_id", *o.VlanId)
	}

	// Routed for networks
	if len(o.RoutedForNetworks) > 0 {
		checks.append(t, "TestCheckResourceAttr", "routed_for_networks.#", fmt.Sprintf("%d", len(o.RoutedForNetworks)))
		for i, network := range o.RoutedForNetworks {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("routed_for_networks.%d", i), network)
		}
	}

	// Internal Access
	if o.InternalAccess != nil {
		if o.InternalAccess.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "internal_access.enabled", fmt.Sprintf("%t", *o.InternalAccess.Enabled))
		}
	}

	// Internet Access
	if o.InternetAccess != nil {
		if o.InternetAccess.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "internet_access.enabled", fmt.Sprintf("%t", *o.InternetAccess.Enabled))
		}
		if o.InternetAccess.CreateSimpleServicePolicy != nil {
			checks.append(t, "TestCheckResourceAttr", "internet_access.create_simple_service_policy", fmt.Sprintf("%t", *o.InternetAccess.CreateSimpleServicePolicy))
		}
		if o.InternetAccess.Restricted != nil {
			checks.append(t, "TestCheckResourceAttr", "internet_access.restricted", fmt.Sprintf("%t", *o.InternetAccess.Restricted))
		}

		// Internet Access Destination NAT
		if len(o.InternetAccess.InternetAccessDestinationNat) > 0 {
			for key, destNat := range o.InternetAccess.InternetAccessDestinationNat {
				basePath := fmt.Sprintf("internet_access.destination_nat.%s", key)
				if destNat.InternalIp != nil {
					checks.append(t, "TestCheckResourceAttr", basePath+".internal_ip", *destNat.InternalIp)
				}
				if destNat.Name != nil {
					checks.append(t, "TestCheckResourceAttr", basePath+".name", *destNat.Name)
				}
				if destNat.Port != nil {
					checks.append(t, "TestCheckResourceAttr", basePath+".port", *destNat.Port)
				}
				if destNat.WanName != nil {
					checks.append(t, "TestCheckResourceAttr", basePath+".wan_name", *destNat.WanName)
				}
			}
		}

		// Internet Access Static NAT
		if len(o.InternetAccess.InternetAccessStaticNat) > 0 {
			for key, staticNat := range o.InternetAccess.InternetAccessStaticNat {
				basePath := fmt.Sprintf("internet_access.static_nat.%s", key)
				checks.append(t, "TestCheckResourceAttr", basePath+".internal_ip", staticNat.InternalIp)
				checks.append(t, "TestCheckResourceAttr", basePath+".name", staticNat.Name)
				if staticNat.WanName != nil {
					checks.append(t, "TestCheckResourceAttr", basePath+".wan_name", *staticNat.WanName)
				}
			}
		}
	}

	// Multicast
	if o.Multicast != nil {
		if o.Multicast.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "multicast.enabled", fmt.Sprintf("%t", *o.Multicast.Enabled))
		}
		if o.Multicast.DisableIgmp != nil {
			checks.append(t, "TestCheckResourceAttr", "multicast.disable_igmp", fmt.Sprintf("%t", *o.Multicast.DisableIgmp))
		}

		// Multicast Groups
		if len(o.Multicast.Groups) > 0 {
			for key, group := range o.Multicast.Groups {
				basePath := fmt.Sprintf("multicast.groups.%s", key)
				if group.RpIp != nil {
					checks.append(t, "TestCheckResourceAttr", basePath+".rp_ip", *group.RpIp)
				}
			}
		}
	}

	// Tenants
	if len(o.Tenants) > 0 {
		for key, tenant := range o.Tenants {
			basePath := fmt.Sprintf("tenants.%s", key)
			if len(tenant.Addresses) > 0 {
				checks.append(t, "TestCheckResourceAttr", basePath+".addresses.#", fmt.Sprintf("%d", len(tenant.Addresses)))
				for i, address := range tenant.Addresses {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.addresses.%d", basePath, i), address)
				}
			}
		}
	}

	// VPN Access
	if len(o.VpnAccess) > 0 {
		for key, vpnAccess := range o.VpnAccess {
			basePath := fmt.Sprintf("vpn_access.%s", key)

			if vpnAccess.AdvertisedSubnet != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".advertised_subnet", *vpnAccess.AdvertisedSubnet)
			}
			if vpnAccess.AllowPing != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".allow_ping", fmt.Sprintf("%t", *vpnAccess.AllowPing))
			}
			if vpnAccess.NatPool != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".nat_pool", *vpnAccess.NatPool)
			}
			if vpnAccess.NoReadvertiseToLanBgp != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".no_readvertise_to_lan_bgp", fmt.Sprintf("%t", *vpnAccess.NoReadvertiseToLanBgp))
			}
			if vpnAccess.NoReadvertiseToLanOspf != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".no_readvertise_to_lan_ospf", fmt.Sprintf("%t", *vpnAccess.NoReadvertiseToLanOspf))
			}
			if vpnAccess.NoReadvertiseToOverlay != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".no_readvertise_to_overlay", fmt.Sprintf("%t", *vpnAccess.NoReadvertiseToOverlay))
			}
			if vpnAccess.Routed != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".routed", fmt.Sprintf("%t", *vpnAccess.Routed))
			}
			if vpnAccess.SummarizedSubnet != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".summarized_subnet", *vpnAccess.SummarizedSubnet)
			}
			if vpnAccess.SummarizedSubnetToLanBgp != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".summarized_subnet_to_lan_bgp", *vpnAccess.SummarizedSubnetToLanBgp)
			}
			if vpnAccess.SummarizedSubnetToLanOspf != nil {
				checks.append(t, "TestCheckResourceAttr", basePath+".summarized_subnet_to_lan_ospf", *vpnAccess.SummarizedSubnetToLanOspf)
			}

			// Other VRFs
			if len(vpnAccess.OtherVrfs) > 0 {
				checks.append(t, "TestCheckResourceAttr", basePath+".other_vrfs.#", fmt.Sprintf("%d", len(vpnAccess.OtherVrfs)))
				for i, vrf := range vpnAccess.OtherVrfs {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.other_vrfs.%d", basePath, i), vrf)
				}
			}

			// Source NAT
			if vpnAccess.SourceNat != nil {
				if vpnAccess.SourceNat.ExternalIp != nil {
					checks.append(t, "TestCheckResourceAttr", basePath+".source_nat.external_ip", *vpnAccess.SourceNat.ExternalIp)
				}
			}

			// VPN Access Destination NAT
			if len(vpnAccess.VpnAccessDestinationNat) > 0 {
				for destKey, destNat := range vpnAccess.VpnAccessDestinationNat {
					destBasePath := fmt.Sprintf("%s.destination_nat.%s", basePath, destKey)
					if destNat.InternalIp != nil {
						checks.append(t, "TestCheckResourceAttr", destBasePath+".internal_ip", *destNat.InternalIp)
					}
					if destNat.Name != nil {
						checks.append(t, "TestCheckResourceAttr", destBasePath+".name", *destNat.Name)
					}
					if destNat.Port != nil {
						checks.append(t, "TestCheckResourceAttr", destBasePath+".port", *destNat.Port)
					}
				}
			}

			// VPN Access Static NAT
			if len(vpnAccess.VpnAccessStaticNat) > 0 {
				for staticKey, staticNat := range vpnAccess.VpnAccessStaticNat {
					staticBasePath := fmt.Sprintf("%s.static_nat.%s", basePath, staticKey)
					checks.append(t, "TestCheckResourceAttr", staticBasePath+".internal_ip", staticNat.InternalIp)
					checks.append(t, "TestCheckResourceAttr", staticBasePath+".name", staticNat.Name)
				}
			}
		}
	}

	return checks
}
