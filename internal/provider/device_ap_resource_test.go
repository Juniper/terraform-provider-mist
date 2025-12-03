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

func TestDeviceAp(t *testing.T) {
	type testStep struct {
		config DeviceApModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: DeviceApModel{
						DeviceId: "00000000-0000-0000-1000-d420b041d5bf",
						Name:     "test_ap",
						SiteId:   "2c107c8e-2e06-404a-ba61-e25b5757ecea",
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/device_ap_resource/device_ap_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		var FixtureDeviceApModel DeviceApModel
		err = hcl.Decode(&FixtureDeviceApModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureDeviceApModel,
				},
			},
		}
	}

	for tName, tCase := range testCases {
		t.Skip("Skipping device_ap tests, as they require a real device.")
		t.Run(tName, func(t *testing.T) {
			resourceType := "device_ap"

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

func (s *DeviceApModel) testChecks(t testing.TB, rType, rName string) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + rName)

	// Required attributes
	checks.append(t, "TestCheckResourceAttr", "site_id", s.SiteId)
	checks.append(t, "TestCheckResourceAttr", "name", s.Name)
	checks.append(t, "TestCheckResourceAttr", "device_id", s.DeviceId)

	// Computed-only attributes (check presence only)
	checks.append(t, "TestCheckResourceAttrSet", "image1_url")
	checks.append(t, "TestCheckResourceAttrSet", "image2_url")
	checks.append(t, "TestCheckResourceAttrSet", "image3_url")
	checks.append(t, "TestCheckResourceAttrSet", "mac")
	checks.append(t, "TestCheckResourceAttrSet", "model")
	checks.append(t, "TestCheckResourceAttrSet", "org_id")
	checks.append(t, "TestCheckResourceAttrSet", "serial")
	checks.append(t, "TestCheckResourceAttrSet", "type")

	// Optional attributes with conditional checks
	if s.Aeroscout != nil {
		if s.Aeroscout.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "aeroscout.enabled", fmt.Sprintf("%t", *s.Aeroscout.Enabled))
		}
		if s.Aeroscout.Host != nil {
			checks.append(t, "TestCheckResourceAttr", "aeroscout.host", *s.Aeroscout.Host)
		}
		if s.Aeroscout.LocateConnected != nil {
			checks.append(t, "TestCheckResourceAttr", "aeroscout.locate_connected", fmt.Sprintf("%t", *s.Aeroscout.LocateConnected))
		}
		if s.Aeroscout.Port != nil {
			checks.append(t, "TestCheckResourceAttr", "aeroscout.port", fmt.Sprintf("%d", *s.Aeroscout.Port))
		}
	}

	if s.Airista != nil {
		if s.Airista.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "airista.enabled", fmt.Sprintf("%t", *s.Airista.Enabled))
		}
		if s.Airista.Host != nil {
			checks.append(t, "TestCheckResourceAttr", "airista.host", *s.Airista.Host)
		}
		if s.Airista.Port != nil {
			checks.append(t, "TestCheckResourceAttr", "airista.port", fmt.Sprintf("%d", *s.Airista.Port))
		}
	}

	if s.BleConfig != nil {
		if s.BleConfig.BeaconEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "ble_config.beacon_enabled", fmt.Sprintf("%t", *s.BleConfig.BeaconEnabled))
		}
		if s.BleConfig.BeaconRate != nil {
			checks.append(t, "TestCheckResourceAttr", "ble_config.beacon_rate", fmt.Sprintf("%d", *s.BleConfig.BeaconRate))
		}
		if s.BleConfig.BeaconRateMode != nil {
			checks.append(t, "TestCheckResourceAttr", "ble_config.beacon_rate_mode", *s.BleConfig.BeaconRateMode)
		}
		if len(s.BleConfig.BeamDisabled) > 0 {
			for i, beam := range s.BleConfig.BeamDisabled {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ble_config.beam_disabled.%d", i), fmt.Sprintf("%d", beam))
			}
		}
		if s.BleConfig.CustomBlePacketEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "ble_config.custom_ble_packet_enabled", fmt.Sprintf("%t", *s.BleConfig.CustomBlePacketEnabled))
		}
		if s.BleConfig.CustomBlePacketFrame != nil {
			checks.append(t, "TestCheckResourceAttr", "ble_config.custom_ble_packet_frame", *s.BleConfig.CustomBlePacketFrame)
		}
		if s.BleConfig.CustomBlePacketFreqMsec != nil {
			checks.append(t, "TestCheckResourceAttr", "ble_config.custom_ble_packet_freq_msec", fmt.Sprintf("%d", *s.BleConfig.CustomBlePacketFreqMsec))
		}
		if s.BleConfig.EddystoneUidAdvPower != nil {
			checks.append(t, "TestCheckResourceAttr", "ble_config.eddystone_uid_adv_power", fmt.Sprintf("%d", *s.BleConfig.EddystoneUidAdvPower))
		}
		if s.BleConfig.EddystoneUidBeams != nil {
			checks.append(t, "TestCheckResourceAttr", "ble_config.eddystone_uid_beams", *s.BleConfig.EddystoneUidBeams)
		}
		if s.BleConfig.EddystoneUidEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "ble_config.eddystone_uid_enabled", fmt.Sprintf("%t", *s.BleConfig.EddystoneUidEnabled))
		}
		if s.BleConfig.EddystoneUidFreqMsec != nil {
			checks.append(t, "TestCheckResourceAttr", "ble_config.eddystone_uid_freq_msec", fmt.Sprintf("%d", *s.BleConfig.EddystoneUidFreqMsec))
		}
		if s.BleConfig.EddystoneUidInstance != nil {
			checks.append(t, "TestCheckResourceAttr", "ble_config.eddystone_uid_instance", *s.BleConfig.EddystoneUidInstance)
		}
		if s.BleConfig.EddystoneUidNamespace != nil {
			checks.append(t, "TestCheckResourceAttr", "ble_config.eddystone_uid_namespace", *s.BleConfig.EddystoneUidNamespace)
		}
		if s.BleConfig.EddystoneUrlAdvPower != nil {
			checks.append(t, "TestCheckResourceAttr", "ble_config.eddystone_url_adv_power", fmt.Sprintf("%d", *s.BleConfig.EddystoneUrlAdvPower))
		}
		if s.BleConfig.EddystoneUrlBeams != nil {
			checks.append(t, "TestCheckResourceAttr", "ble_config.eddystone_url_beams", *s.BleConfig.EddystoneUrlBeams)
		}
		if s.BleConfig.EddystoneUrlEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "ble_config.eddystone_url_enabled", fmt.Sprintf("%t", *s.BleConfig.EddystoneUrlEnabled))
		}
		if s.BleConfig.EddystoneUrlFreqMsec != nil {
			checks.append(t, "TestCheckResourceAttr", "ble_config.eddystone_url_freq_msec", fmt.Sprintf("%d", *s.BleConfig.EddystoneUrlFreqMsec))
		}
		if s.BleConfig.EddystoneUrlUrl != nil {
			checks.append(t, "TestCheckResourceAttr", "ble_config.eddystone_url_url", *s.BleConfig.EddystoneUrlUrl)
		}
		if s.BleConfig.IbeaconAdvPower != nil {
			checks.append(t, "TestCheckResourceAttr", "ble_config.ibeacon_adv_power", fmt.Sprintf("%d", *s.BleConfig.IbeaconAdvPower))
		}
		if s.BleConfig.IbeaconBeams != nil {
			checks.append(t, "TestCheckResourceAttr", "ble_config.ibeacon_beams", *s.BleConfig.IbeaconBeams)
		}
		if s.BleConfig.IbeaconEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "ble_config.ibeacon_enabled", fmt.Sprintf("%t", *s.BleConfig.IbeaconEnabled))
		}
		if s.BleConfig.IbeaconFreqMsec != nil {
			checks.append(t, "TestCheckResourceAttr", "ble_config.ibeacon_freq_msec", fmt.Sprintf("%d", *s.BleConfig.IbeaconFreqMsec))
		}
		if s.BleConfig.IbeaconMajor != nil {
			checks.append(t, "TestCheckResourceAttr", "ble_config.ibeacon_major", fmt.Sprintf("%d", *s.BleConfig.IbeaconMajor))
		}
		if s.BleConfig.IbeaconMinor != nil {
			checks.append(t, "TestCheckResourceAttr", "ble_config.ibeacon_minor", fmt.Sprintf("%d", *s.BleConfig.IbeaconMinor))
		}
		if s.BleConfig.IbeaconUuid != nil {
			checks.append(t, "TestCheckResourceAttr", "ble_config.ibeacon_uuid", *s.BleConfig.IbeaconUuid)
		}
		if s.BleConfig.Power != nil {
			checks.append(t, "TestCheckResourceAttr", "ble_config.power", fmt.Sprintf("%d", *s.BleConfig.Power))
		}
		if s.BleConfig.PowerMode != nil {
			checks.append(t, "TestCheckResourceAttr", "ble_config.power_mode", *s.BleConfig.PowerMode)
		}
	}

	if s.Centrak != nil {
		if s.Centrak.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "centrak.enabled", fmt.Sprintf("%t", *s.Centrak.Enabled))
		}
	}

	if s.ClientBridge != nil {
		if s.ClientBridge.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "client_bridge.enabled", fmt.Sprintf("%t", *s.ClientBridge.Enabled))
		}
		if s.ClientBridge.Ssid != nil {
			checks.append(t, "TestCheckResourceAttr", "client_bridge.ssid", *s.ClientBridge.Ssid)
		}
		if s.ClientBridge.Auth != nil {
			if s.ClientBridge.Auth.AuthType != nil {
				checks.append(t, "TestCheckResourceAttr", "client_bridge.auth.type", *s.ClientBridge.Auth.AuthType)
			}
			if s.ClientBridge.Auth.Psk != nil {
				checks.append(t, "TestCheckResourceAttr", "client_bridge.auth.psk", *s.ClientBridge.Auth.Psk)
			}
		}
	}

	if s.DisableEth1 != nil {
		checks.append(t, "TestCheckResourceAttr", "disable_eth1", fmt.Sprintf("%t", *s.DisableEth1))
	}
	if s.DisableEth2 != nil {
		checks.append(t, "TestCheckResourceAttr", "disable_eth2", fmt.Sprintf("%t", *s.DisableEth2))
	}
	if s.DisableEth3 != nil {
		checks.append(t, "TestCheckResourceAttr", "disable_eth3", fmt.Sprintf("%t", *s.DisableEth3))
	}
	if s.DisableModule != nil {
		checks.append(t, "TestCheckResourceAttr", "disable_module", fmt.Sprintf("%t", *s.DisableModule))
	}

	if s.EslConfig != nil {
		if s.EslConfig.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "esl_config.enabled", fmt.Sprintf("%t", *s.EslConfig.Enabled))
		}
		if s.EslConfig.Host != nil {
			checks.append(t, "TestCheckResourceAttr", "esl_config.host", *s.EslConfig.Host)
		}
		if s.EslConfig.Port != nil {
			checks.append(t, "TestCheckResourceAttr", "esl_config.port", fmt.Sprintf("%d", *s.EslConfig.Port))
		}
		if s.EslConfig.EslConfigType != nil {
			checks.append(t, "TestCheckResourceAttr", "esl_config.type", *s.EslConfig.EslConfigType)
		}
		if s.EslConfig.Cacert != nil {
			checks.append(t, "TestCheckResourceAttr", "esl_config.cacert", *s.EslConfig.Cacert)
		}
		if s.EslConfig.Channel != nil {
			checks.append(t, "TestCheckResourceAttr", "esl_config.channel", fmt.Sprintf("%d", *s.EslConfig.Channel))
		}
		if s.EslConfig.VerifyCert != nil {
			checks.append(t, "TestCheckResourceAttr", "esl_config.verify_cert", fmt.Sprintf("%t", *s.EslConfig.VerifyCert))
		}
		if s.EslConfig.VlanId != nil {
			checks.append(t, "TestCheckResourceAttr", "esl_config.vlan_id", fmt.Sprintf("%d", *s.EslConfig.VlanId))
		}
	}

	if s.FlowControl != nil {
		checks.append(t, "TestCheckResourceAttr", "flow_control", fmt.Sprintf("%t", *s.FlowControl))
	}
	if s.Height != nil {
		checks.append(t, "TestCheckResourceAttr", "height", fmt.Sprintf("%g", *s.Height))
	}

	if s.IpConfig != nil {
		if s.IpConfig.IpConfigType != nil {
			checks.append(t, "TestCheckResourceAttr", "ip_config.type", *s.IpConfig.IpConfigType)
		}
		if s.IpConfig.Ip != nil {
			checks.append(t, "TestCheckResourceAttr", "ip_config.ip", *s.IpConfig.Ip)
		}
		if s.IpConfig.Netmask != nil {
			checks.append(t, "TestCheckResourceAttr", "ip_config.netmask", *s.IpConfig.Netmask)
		}
		if s.IpConfig.Gateway != nil {
			checks.append(t, "TestCheckResourceAttr", "ip_config.gateway", *s.IpConfig.Gateway)
		}
		if s.IpConfig.Gateway6 != nil {
			checks.append(t, "TestCheckResourceAttr", "ip_config.gateway6", *s.IpConfig.Gateway6)
		}
		if s.IpConfig.Ip6 != nil {
			checks.append(t, "TestCheckResourceAttr", "ip_config.ip6", *s.IpConfig.Ip6)
		}
		if s.IpConfig.Mtu != nil {
			checks.append(t, "TestCheckResourceAttr", "ip_config.mtu", fmt.Sprintf("%d", *s.IpConfig.Mtu))
		}
		if s.IpConfig.Netmask6 != nil {
			checks.append(t, "TestCheckResourceAttr", "ip_config.netmask6", *s.IpConfig.Netmask6)
		}
		if s.IpConfig.Type6 != nil {
			checks.append(t, "TestCheckResourceAttr", "ip_config.type6", *s.IpConfig.Type6)
		}
		if s.IpConfig.VlanId != nil {
			checks.append(t, "TestCheckResourceAttr", "ip_config.vlan_id", fmt.Sprintf("%d", *s.IpConfig.VlanId))
		}
		if len(s.IpConfig.Dns) > 0 {
			for i, dns := range s.IpConfig.Dns {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ip_config.dns.%d", i), dns)
			}
		}
		if len(s.IpConfig.DnsSuffix) > 0 {
			for i, suffix := range s.IpConfig.DnsSuffix {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ip_config.dns_suffix.%d", i), suffix)
			}
		}
	}

	if s.LacpConfig != nil {
		if s.LacpConfig.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "lacp_config.enabled", fmt.Sprintf("%t", *s.LacpConfig.Enabled))
		}
	}

	if s.Led != nil {
		if s.Led.Brightness != nil {
			checks.append(t, "TestCheckResourceAttr", "led.brightness", fmt.Sprintf("%d", *s.Led.Brightness))
		}
		if s.Led.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "led.enabled", fmt.Sprintf("%t", *s.Led.Enabled))
		}
	}

	if s.Locked != nil {
		checks.append(t, "TestCheckResourceAttr", "locked", fmt.Sprintf("%t", *s.Locked))
	}
	if s.MapId != nil {
		checks.append(t, "TestCheckResourceAttr", "map_id", *s.MapId)
	}

	if s.Mesh != nil {
		if len(s.Mesh.Bands) > 0 {
			for i, band := range s.Mesh.Bands {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("mesh.bands.%d", i), band)
			}
		}
		if s.Mesh.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "mesh.enabled", fmt.Sprintf("%t", *s.Mesh.Enabled))
		}
		if s.Mesh.Group != nil {
			checks.append(t, "TestCheckResourceAttr", "mesh.group", fmt.Sprintf("%d", *s.Mesh.Group))
		}
		if s.Mesh.Role != nil {
			checks.append(t, "TestCheckResourceAttr", "mesh.role", *s.Mesh.Role)
		}
	}

	if s.Notes != nil {
		checks.append(t, "TestCheckResourceAttr", "notes", *s.Notes)
	}
	if len(s.NtpServers) > 0 {
		for i, server := range s.NtpServers {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ntp_servers.%d", i), server)
		}
	}
	if s.Orientation != nil {
		checks.append(t, "TestCheckResourceAttr", "orientation", fmt.Sprintf("%d", *s.Orientation))
	}
	if s.PoePassthrough != nil {
		checks.append(t, "TestCheckResourceAttr", "poe_passthrough", fmt.Sprintf("%t", *s.PoePassthrough))
	}

	// Port config map checks
	if s.PortConfig != nil {
		for portName, portCfg := range s.PortConfig {
			portPath := fmt.Sprintf("port_config.%s", portName)
			if portCfg.Disabled != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.disabled", portPath), fmt.Sprintf("%t", *portCfg.Disabled))
			}
			if portCfg.DynamicVlan != nil {
				if portCfg.DynamicVlan.DefaultVlanId != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.dynamic_vlan.default_vlan_id", portPath), fmt.Sprintf("%d", *portCfg.DynamicVlan.DefaultVlanId))
				}
				if portCfg.DynamicVlan.Enabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.dynamic_vlan.enabled", portPath), fmt.Sprintf("%t", *portCfg.DynamicVlan.Enabled))
				}
				if portCfg.DynamicVlan.DynamicVlanType != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.dynamic_vlan.type", portPath), *portCfg.DynamicVlan.DynamicVlanType)
				}
				if len(portCfg.DynamicVlan.Vlans) > 0 {
					for key, vlan := range portCfg.DynamicVlan.Vlans {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.dynamic_vlan.vlans.%s", portPath, key), vlan)
					}
				}
			}
			if portCfg.EnableMacAuth != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.enable_mac_auth", portPath), fmt.Sprintf("%t", *portCfg.EnableMacAuth))
			}
			if portCfg.Forwarding != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.forwarding", portPath), *portCfg.Forwarding)
			}
			if portCfg.MacAuthPreferred != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.mac_auth_preferred", portPath), fmt.Sprintf("%t", *portCfg.MacAuthPreferred))
			}
			if portCfg.MacAuthProtocol != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.mac_auth_protocol", portPath), *portCfg.MacAuthProtocol)
			}
			if portCfg.MistNac != nil {
				if portCfg.MistNac.AcctInterimInterval != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.mist_nac.acct_interim_interval", portPath), fmt.Sprintf("%d", *portCfg.MistNac.AcctInterimInterval))
				}
				if portCfg.MistNac.AuthServersRetries != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.mist_nac.auth_servers_retries", portPath), fmt.Sprintf("%d", *portCfg.MistNac.AuthServersRetries))
				}
				if portCfg.MistNac.AuthServersTimeout != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.mist_nac.auth_servers_timeout", portPath), fmt.Sprintf("%d", *portCfg.MistNac.AuthServersTimeout))
				}
				if portCfg.MistNac.CoaEnabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.mist_nac.coa_enabled", portPath), fmt.Sprintf("%t", *portCfg.MistNac.CoaEnabled))
				}
				if portCfg.MistNac.CoaPort != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.mist_nac.coa_port", portPath), fmt.Sprintf("%d", *portCfg.MistNac.CoaPort))
				}
				if portCfg.MistNac.Enabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.mist_nac.enabled", portPath), fmt.Sprintf("%t", *portCfg.MistNac.Enabled))
				}
				if portCfg.MistNac.FastDot1xTimers != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.mist_nac.fast_dot1x_timers", portPath), fmt.Sprintf("%t", *portCfg.MistNac.FastDot1xTimers))
				}
				if portCfg.MistNac.Network != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.mist_nac.network", portPath), *portCfg.MistNac.Network)
				}
				if portCfg.MistNac.SourceIp != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.mist_nac.source_ip", portPath), *portCfg.MistNac.SourceIp)
				}
			}
			if portCfg.MxTunnelId != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.mx_tunnel_id", portPath), *portCfg.MxTunnelId)
			}
			if portCfg.MxtunnelName != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.mxtunnel_name", portPath), *portCfg.MxtunnelName)
			}
			if portCfg.PortAuth != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.port_auth", portPath), *portCfg.PortAuth)
			}
			if portCfg.PortVlanId != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.port_vlan_id", portPath), fmt.Sprintf("%d", *portCfg.PortVlanId))
			}
			if portCfg.RadiusConfig != nil {
				if portCfg.RadiusConfig.AcctInterimInterval != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radius_config.acct_interim_interval", portPath), fmt.Sprintf("%d", *portCfg.RadiusConfig.AcctInterimInterval))
				}
				if len(portCfg.RadiusConfig.AcctServers) > 0 {
					for i, server := range portCfg.RadiusConfig.AcctServers {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radius_config.acct_servers.%d.host", portPath, i), server.Host)
						if server.KeywrapEnabled != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radius_config.acct_servers.%d.keywrap_enabled", portPath, i), fmt.Sprintf("%t", *server.KeywrapEnabled))
						}
						if server.KeywrapFormat != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radius_config.acct_servers.%d.keywrap_format", portPath, i), *server.KeywrapFormat)
						}
						if server.KeywrapKek != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radius_config.acct_servers.%d.keywrap_kek", portPath, i), *server.KeywrapKek)
						}
						if server.KeywrapMack != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radius_config.acct_servers.%d.keywrap_mack", portPath, i), *server.KeywrapMack)
						}
						if server.Port != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radius_config.acct_servers.%d.port", portPath, i), *server.Port)
						}
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radius_config.acct_servers.%d.secret", portPath, i), server.Secret)
					}
				}
				if len(portCfg.RadiusConfig.AuthServers) > 0 {
					for i, server := range portCfg.RadiusConfig.AuthServers {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radius_config.auth_servers.%d.host", portPath, i), server.Host)
						if server.KeywrapEnabled != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radius_config.auth_servers.%d.keywrap_enabled", portPath, i), fmt.Sprintf("%t", *server.KeywrapEnabled))
						}
						if server.KeywrapFormat != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radius_config.auth_servers.%d.keywrap_format", portPath, i), *server.KeywrapFormat)
						}
						if server.KeywrapKek != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radius_config.auth_servers.%d.keywrap_kek", portPath, i), *server.KeywrapKek)
						}
						if server.KeywrapMack != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radius_config.auth_servers.%d.keywrap_mack", portPath, i), *server.KeywrapMack)
						}
						if server.Port != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radius_config.auth_servers.%d.port", portPath, i), *server.Port)
						}
						if server.RequireMessageAuthenticator != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radius_config.auth_servers.%d.require_message_authenticator", portPath, i), fmt.Sprintf("%t", *server.RequireMessageAuthenticator))
						}
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radius_config.auth_servers.%d.secret", portPath, i), server.Secret)
					}
				}
				if portCfg.RadiusConfig.AuthServersRetries != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radius_config.auth_servers_retries", portPath), fmt.Sprintf("%d", *portCfg.RadiusConfig.AuthServersRetries))
				}
				if portCfg.RadiusConfig.AuthServersTimeout != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radius_config.auth_servers_timeout", portPath), fmt.Sprintf("%d", *portCfg.RadiusConfig.AuthServersTimeout))
				}
				if portCfg.RadiusConfig.CoaEnabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radius_config.coa_enabled", portPath), fmt.Sprintf("%t", *portCfg.RadiusConfig.CoaEnabled))
				}
				if portCfg.RadiusConfig.CoaPort != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radius_config.coa_port", portPath), fmt.Sprintf("%d", *portCfg.RadiusConfig.CoaPort))
				}
				if portCfg.RadiusConfig.Network != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radius_config.network", portPath), *portCfg.RadiusConfig.Network)
				}
				if portCfg.RadiusConfig.SourceIp != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radius_config.source_ip", portPath), *portCfg.RadiusConfig.SourceIp)
				}
			}
			if portCfg.Radsec != nil {
				if portCfg.Radsec.CoaEnabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radsec.coa_enabled", portPath), fmt.Sprintf("%t", *portCfg.Radsec.CoaEnabled))
				}
				if portCfg.Radsec.Enabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radsec.enabled", portPath), fmt.Sprintf("%t", *portCfg.Radsec.Enabled))
				}
				if portCfg.Radsec.IdleTimeout != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radsec.idle_timeout", portPath), *portCfg.Radsec.IdleTimeout)
				}
				if len(portCfg.Radsec.MxclusterIds) > 0 {
					for i, clusterId := range portCfg.Radsec.MxclusterIds {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radsec.mxcluster_ids.%d", portPath, i), clusterId)
					}
				}
				if len(portCfg.Radsec.ProxyHosts) > 0 {
					for i, host := range portCfg.Radsec.ProxyHosts {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radsec.proxy_hosts.%d", portPath, i), host)
					}
				}
				if portCfg.Radsec.ServerName != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radsec.server_name", portPath), *portCfg.Radsec.ServerName)
				}
				if len(portCfg.Radsec.Servers) > 0 {
					for i, server := range portCfg.Radsec.Servers {
						if server.Host != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radsec.servers.%d.host", portPath, i), *server.Host)
						}
						if server.Port != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radsec.servers.%d.port", portPath, i), fmt.Sprintf("%d", *server.Port))
						}
					}
				}
				if portCfg.Radsec.UseMxedge != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radsec.use_mxedge", portPath), fmt.Sprintf("%t", *portCfg.Radsec.UseMxedge))
				}
				if portCfg.Radsec.UseSiteMxedge != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.radsec.use_site_mxedge", portPath), fmt.Sprintf("%t", *portCfg.Radsec.UseSiteMxedge))
				}
			}
			if portCfg.VlanId != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.vlan_id", portPath), fmt.Sprintf("%d", *portCfg.VlanId))
			}
			if len(portCfg.VlanIds) > 0 {
				for i, vlanId := range portCfg.VlanIds {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.vlan_ids.%d", portPath, i), fmt.Sprintf("%d", vlanId))
				}
			}
			if portCfg.WxtunnelId != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.wxtunnel_id", portPath), *portCfg.WxtunnelId)
			}
			if portCfg.WxtunnelRemoteId != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.wxtunnel_remote_id", portPath), *portCfg.WxtunnelRemoteId)
			}
		}
	}

	if s.PwrConfig != nil {
		if s.PwrConfig.Base != nil {
			checks.append(t, "TestCheckResourceAttr", "pwr_config.base", fmt.Sprintf("%d", *s.PwrConfig.Base))
		}
		if s.PwrConfig.PreferUsbOverWifi != nil {
			checks.append(t, "TestCheckResourceAttr", "pwr_config.prefer_usb_over_wifi", fmt.Sprintf("%t", *s.PwrConfig.PreferUsbOverWifi))
		}
	}

	if s.RadioConfig != nil {
		if s.RadioConfig.AllowRrmDisable != nil {
			checks.append(t, "TestCheckResourceAttr", "radio_config.allow_rrm_disable", fmt.Sprintf("%t", *s.RadioConfig.AllowRrmDisable))
		}
		if s.RadioConfig.AntGain24 != nil {
			checks.append(t, "TestCheckResourceAttr", "radio_config.ant_gain_24", fmt.Sprintf("%d", *s.RadioConfig.AntGain24))
		}
		if s.RadioConfig.AntGain5 != nil {
			checks.append(t, "TestCheckResourceAttr", "radio_config.ant_gain_5", fmt.Sprintf("%d", *s.RadioConfig.AntGain5))
		}
		if s.RadioConfig.AntGain6 != nil {
			checks.append(t, "TestCheckResourceAttr", "radio_config.ant_gain_6", fmt.Sprintf("%d", *s.RadioConfig.AntGain6))
		}
		if s.RadioConfig.AntennaSelect != nil {
			checks.append(t, "TestCheckResourceAttr", "radio_config.antenna_select", *s.RadioConfig.AntennaSelect)
		}
		if s.RadioConfig.AntennaMode != nil {
			checks.append(t, "TestCheckResourceAttr", "radio_config.antenna_mode", *s.RadioConfig.AntennaMode)
		}
		if s.RadioConfig.Band24Usage != nil {
			checks.append(t, "TestCheckResourceAttr", "radio_config.band_24_usage", *s.RadioConfig.Band24Usage)
		}
		if s.RadioConfig.FullAutomaticRrm != nil {
			checks.append(t, "TestCheckResourceAttr", "radio_config.full_automatic_rrm", fmt.Sprintf("%t", *s.RadioConfig.FullAutomaticRrm))
		}
		if s.RadioConfig.IndoorUse != nil {
			checks.append(t, "TestCheckResourceAttr", "radio_config.indoor_use", fmt.Sprintf("%t", *s.RadioConfig.IndoorUse))
		}
		if s.RadioConfig.RrmManaged != nil {
			checks.append(t, "TestCheckResourceAttr", "radio_config.rrm_managed", fmt.Sprintf("%t", *s.RadioConfig.RrmManaged))
		}
		if s.RadioConfig.ScanningEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "radio_config.scanning_enabled", fmt.Sprintf("%t", *s.RadioConfig.ScanningEnabled))
		}
		if s.RadioConfig.Band24 != nil {
			if s.RadioConfig.Band24.AllowRrmDisable != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_24.allow_rrm_disable", fmt.Sprintf("%t", *s.RadioConfig.Band24.AllowRrmDisable))
			}
			if s.RadioConfig.Band24.AntGain != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_24.ant_gain", fmt.Sprintf("%d", *s.RadioConfig.Band24.AntGain))
			}
			if s.RadioConfig.Band24.AntennaMode != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_24.antenna_mode", *s.RadioConfig.Band24.AntennaMode)
			}
			if s.RadioConfig.Band24.Bandwidth != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_24.bandwidth", fmt.Sprintf("%d", *s.RadioConfig.Band24.Bandwidth))
			}
			if s.RadioConfig.Band24.Channel != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_24.channel", fmt.Sprintf("%d", *s.RadioConfig.Band24.Channel))
			}
			if len(s.RadioConfig.Band24.Channels) > 0 {
				for i, channel := range s.RadioConfig.Band24.Channels {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radio_config.band_24.channels.%d", i), fmt.Sprintf("%d", channel))
				}
			}
			if s.RadioConfig.Band24.Disabled != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_24.disabled", fmt.Sprintf("%t", *s.RadioConfig.Band24.Disabled))
			}
			if s.RadioConfig.Band24.Power != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_24.power", fmt.Sprintf("%d", *s.RadioConfig.Band24.Power))
			}
			if s.RadioConfig.Band24.PowerMax != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_24.power_max", fmt.Sprintf("%d", *s.RadioConfig.Band24.PowerMax))
			}
			if s.RadioConfig.Band24.PowerMin != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_24.power_min", fmt.Sprintf("%d", *s.RadioConfig.Band24.PowerMin))
			}
			if s.RadioConfig.Band24.Preamble != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_24.preamble", *s.RadioConfig.Band24.Preamble)
			}
		}
		if s.RadioConfig.Band5 != nil {
			if s.RadioConfig.Band5.AllowRrmDisable != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_5.allow_rrm_disable", fmt.Sprintf("%t", *s.RadioConfig.Band5.AllowRrmDisable))
			}
			if s.RadioConfig.Band5.AntGain != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_5.ant_gain", fmt.Sprintf("%d", *s.RadioConfig.Band5.AntGain))
			}
			if s.RadioConfig.Band5.AntennaBeamPattern != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_5.antenna_beam_pattern", *s.RadioConfig.Band5.AntennaBeamPattern)
			}
			if s.RadioConfig.Band5.AntennaMode != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_5.antenna_mode", *s.RadioConfig.Band5.AntennaMode)
			}
			if s.RadioConfig.Band5.Bandwidth != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_5.bandwidth", fmt.Sprintf("%d", *s.RadioConfig.Band5.Bandwidth))
			}
			if s.RadioConfig.Band5.Channel != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_5.channel", fmt.Sprintf("%d", *s.RadioConfig.Band5.Channel))
			}
			if len(s.RadioConfig.Band5.Channels) > 0 {
				for i, channel := range s.RadioConfig.Band5.Channels {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radio_config.band_5.channels.%d", i), fmt.Sprintf("%d", channel))
				}
			}
			if s.RadioConfig.Band5.Disabled != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_5.disabled", fmt.Sprintf("%t", *s.RadioConfig.Band5.Disabled))
			}
			if s.RadioConfig.Band5.Power != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_5.power", fmt.Sprintf("%d", *s.RadioConfig.Band5.Power))
			}
			if s.RadioConfig.Band5.PowerMax != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_5.power_max", fmt.Sprintf("%d", *s.RadioConfig.Band5.PowerMax))
			}
			if s.RadioConfig.Band5.PowerMin != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_5.power_min", fmt.Sprintf("%d", *s.RadioConfig.Band5.PowerMin))
			}
			if s.RadioConfig.Band5.Preamble != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_5.preamble", *s.RadioConfig.Band5.Preamble)
			}
		}
		if s.RadioConfig.Band5On24Radio != nil {
			if s.RadioConfig.Band5On24Radio.AllowRrmDisable != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_5_on_24_radio.allow_rrm_disable", fmt.Sprintf("%t", *s.RadioConfig.Band5On24Radio.AllowRrmDisable))
			}
			if s.RadioConfig.Band5On24Radio.AntGain != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_5_on_24_radio.ant_gain", fmt.Sprintf("%d", *s.RadioConfig.Band5On24Radio.AntGain))
			}
			if s.RadioConfig.Band5On24Radio.AntennaBeamPattern != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_5_on_24_radio.antenna_beam_pattern", *s.RadioConfig.Band5On24Radio.AntennaBeamPattern)
			}
			if s.RadioConfig.Band5On24Radio.AntennaMode != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_5_on_24_radio.antenna_mode", *s.RadioConfig.Band5On24Radio.AntennaMode)
			}
			if s.RadioConfig.Band5On24Radio.Bandwidth != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_5_on_24_radio.bandwidth", fmt.Sprintf("%d", *s.RadioConfig.Band5On24Radio.Bandwidth))
			}
			if s.RadioConfig.Band5On24Radio.Channel != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_5_on_24_radio.channel", fmt.Sprintf("%d", *s.RadioConfig.Band5On24Radio.Channel))
			}
			if len(s.RadioConfig.Band5On24Radio.Channels) > 0 {
				for i, channel := range s.RadioConfig.Band5On24Radio.Channels {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radio_config.band_5_on_24_radio.channels.%d", i), fmt.Sprintf("%d", channel))
				}
			}
			if s.RadioConfig.Band5On24Radio.Disabled != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_5_on_24_radio.disabled", fmt.Sprintf("%t", *s.RadioConfig.Band5On24Radio.Disabled))
			}
			if s.RadioConfig.Band5On24Radio.Power != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_5_on_24_radio.power", fmt.Sprintf("%d", *s.RadioConfig.Band5On24Radio.Power))
			}
			if s.RadioConfig.Band5On24Radio.PowerMax != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_5_on_24_radio.power_max", fmt.Sprintf("%d", *s.RadioConfig.Band5On24Radio.PowerMax))
			}
			if s.RadioConfig.Band5On24Radio.PowerMin != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_5_on_24_radio.power_min", fmt.Sprintf("%d", *s.RadioConfig.Band5On24Radio.PowerMin))
			}
			if s.RadioConfig.Band5On24Radio.Preamble != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_5_on_24_radio.preamble", *s.RadioConfig.Band5On24Radio.Preamble)
			}
		}
		if s.RadioConfig.Band6 != nil {
			if s.RadioConfig.Band6.AllowRrmDisable != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_6.allow_rrm_disable", fmt.Sprintf("%t", *s.RadioConfig.Band6.AllowRrmDisable))
			}
			if s.RadioConfig.Band6.AntGain != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_6.ant_gain", fmt.Sprintf("%d", *s.RadioConfig.Band6.AntGain))
			}
			if s.RadioConfig.Band6.AntennaBeamPattern != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_6.antenna_beam_pattern", *s.RadioConfig.Band6.AntennaBeamPattern)
			}
			if s.RadioConfig.Band6.AntennaMode != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_6.antenna_mode", *s.RadioConfig.Band6.AntennaMode)
			}
			if s.RadioConfig.Band6.Bandwidth != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_6.bandwidth", fmt.Sprintf("%d", *s.RadioConfig.Band6.Bandwidth))
			}
			if s.RadioConfig.Band6.Channel != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_6.channel", fmt.Sprintf("%d", *s.RadioConfig.Band6.Channel))
			}
			if len(s.RadioConfig.Band6.Channels) > 0 {
				for i, channel := range s.RadioConfig.Band6.Channels {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("radio_config.band_6.channels.%d", i), fmt.Sprintf("%d", channel))
				}
			}
			if s.RadioConfig.Band6.Disabled != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_6.disabled", fmt.Sprintf("%t", *s.RadioConfig.Band6.Disabled))
			}
			if s.RadioConfig.Band6.Power != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_6.power", fmt.Sprintf("%d", *s.RadioConfig.Band6.Power))
			}
			if s.RadioConfig.Band6.PowerMax != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_6.power_max", fmt.Sprintf("%d", *s.RadioConfig.Band6.PowerMax))
			}
			if s.RadioConfig.Band6.PowerMin != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_6.power_min", fmt.Sprintf("%d", *s.RadioConfig.Band6.PowerMin))
			}
			if s.RadioConfig.Band6.Preamble != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_6.preamble", *s.RadioConfig.Band6.Preamble)
			}
			if s.RadioConfig.Band6.StandardPower != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_6.standard_power", fmt.Sprintf("%t", *s.RadioConfig.Band6.StandardPower))
			}
		}
	}

	if s.UplinkPortConfig != nil {
		if s.UplinkPortConfig.Dot1x != nil {
			checks.append(t, "TestCheckResourceAttr", "uplink_port_config.dot1x", fmt.Sprintf("%t", *s.UplinkPortConfig.Dot1x))
		}
		if s.UplinkPortConfig.KeepWlansUpIfDown != nil {
			checks.append(t, "TestCheckResourceAttr", "uplink_port_config.keep_wlans_up_if_down", fmt.Sprintf("%t", *s.UplinkPortConfig.KeepWlansUpIfDown))
		}
	}

	if s.UsbConfig != nil {
		if s.UsbConfig.Cacert != nil {
			checks.append(t, "TestCheckResourceAttr", "usb_config.cacert", *s.UsbConfig.Cacert)
		}
		if s.UsbConfig.Channel != nil {
			checks.append(t, "TestCheckResourceAttr", "usb_config.channel", fmt.Sprintf("%d", *s.UsbConfig.Channel))
		}
		if s.UsbConfig.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "usb_config.enabled", fmt.Sprintf("%t", *s.UsbConfig.Enabled))
		}
		if s.UsbConfig.Host != nil {
			checks.append(t, "TestCheckResourceAttr", "usb_config.host", *s.UsbConfig.Host)
		}
		if s.UsbConfig.Port != nil {
			checks.append(t, "TestCheckResourceAttr", "usb_config.port", fmt.Sprintf("%d", *s.UsbConfig.Port))
		}
		if s.UsbConfig.UsbConfigType != nil {
			checks.append(t, "TestCheckResourceAttr", "usb_config.type", *s.UsbConfig.UsbConfigType)
		}
		if s.UsbConfig.VerifyCert != nil {
			checks.append(t, "TestCheckResourceAttr", "usb_config.verify_cert", fmt.Sprintf("%t", *s.UsbConfig.VerifyCert))
		}
		if s.UsbConfig.VlanId != nil {
			checks.append(t, "TestCheckResourceAttr", "usb_config.vlan_id", fmt.Sprintf("%d", *s.UsbConfig.VlanId))
		}
	}

	if s.Vars != nil {
		for key, value := range s.Vars {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vars.%s", key), value)
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
