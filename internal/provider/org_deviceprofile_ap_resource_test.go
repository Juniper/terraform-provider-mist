package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/resource_org_deviceprofile_ap"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgDeviceprofileApModel(t *testing.T) {
	type testStep struct {
		config OrgDeviceprofileApModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgDeviceprofileApModel{
						OrgId: GetTestOrgId(),
						Name:  "test_ap",
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/org_deviceprofile_ap_resource/org_deviceprofile_ap_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		var FixtureOrgDeviceprofileApModel OrgDeviceprofileApModel

		err = hcl.Decode(&FixtureOrgDeviceprofileApModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		FixtureOrgDeviceprofileApModel.OrgId = GetTestOrgId()

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureOrgDeviceprofileApModel,
				},
			},
		}
	}

	resourceType := "org_deviceprofile_ap"
	var checks testChecks
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {

			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				config := step.config

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				configStr := Render(resourceType, tName, string(f.Bytes()))

				checks = config.testChecks(t, resourceType, tName)
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
	FieldCoverageReport(t, &checks)
}

func (s *OrgDeviceprofileApModel) testChecks(t testing.TB, rType, rName string) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + rName)
	TrackFieldCoverage(t, &checks, "org_deviceprofile_ap", resource_org_deviceprofile_ap.OrgDeviceprofileApResourceSchema)
	// Required parameters
	checks.append(t, "TestCheckResourceAttrSet", "org_id")
	checks.append(t, "TestCheckResourceAttr", "name", s.Name)

	// Optional parameters
	if s.Aeroscout != nil {
		checks.append(t, "TestCheckResourceAttrSet", "aeroscout.%")
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
	if s.BleConfig != nil {
		checks.append(t, "TestCheckResourceAttrSet", "ble_config.%")
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
			checks.append(t, "TestCheckResourceAttr", "ble_config.beam_disabled.#", fmt.Sprintf("%d", len(s.BleConfig.BeamDisabled)))
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
		checks.append(t, "TestCheckResourceAttrSet", "esl_config.%")
		if s.EslConfig.Cacert != nil {
			checks.append(t, "TestCheckResourceAttr", "esl_config.cacert", *s.EslConfig.Cacert)
		}
		if s.EslConfig.Channel != nil {
			checks.append(t, "TestCheckResourceAttr", "esl_config.channel", fmt.Sprintf("%d", *s.EslConfig.Channel))
		}
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
		if s.EslConfig.VerifyCert != nil {
			checks.append(t, "TestCheckResourceAttr", "esl_config.verify_cert", fmt.Sprintf("%t", *s.EslConfig.VerifyCert))
		}
		if s.EslConfig.VlanId != nil {
			checks.append(t, "TestCheckResourceAttr", "esl_config.vlan_id", fmt.Sprintf("%d", *s.EslConfig.VlanId))
		}
	}
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
			for i, dnsSuffix := range s.IpConfig.DnsSuffix {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ip_config.dns_suffix.%d", i), dnsSuffix)
			}
		}
		if s.IpConfig.Gateway != nil {
			checks.append(t, "TestCheckResourceAttr", "ip_config.gateway", *s.IpConfig.Gateway)
		}
		if s.IpConfig.Gateway6 != nil {
			checks.append(t, "TestCheckResourceAttr", "ip_config.gateway6", *s.IpConfig.Gateway6)
		}
		if s.IpConfig.Ip != nil {
			checks.append(t, "TestCheckResourceAttr", "ip_config.ip", *s.IpConfig.Ip)
		}
		if s.IpConfig.Ip6 != nil {
			checks.append(t, "TestCheckResourceAttr", "ip_config.ip6", *s.IpConfig.Ip6)
		}
		if s.IpConfig.Mtu != nil {
			checks.append(t, "TestCheckResourceAttr", "ip_config.mtu", fmt.Sprintf("%d", *s.IpConfig.Mtu))
		}
		if s.IpConfig.Netmask != nil {
			checks.append(t, "TestCheckResourceAttr", "ip_config.netmask", *s.IpConfig.Netmask)
		}
		if s.IpConfig.Netmask6 != nil {
			checks.append(t, "TestCheckResourceAttr", "ip_config.netmask6", *s.IpConfig.Netmask6)
		}
		if s.IpConfig.IpConfigType != nil {
			checks.append(t, "TestCheckResourceAttr", "ip_config.type", *s.IpConfig.IpConfigType)
		}
		if s.IpConfig.Type6 != nil {
			checks.append(t, "TestCheckResourceAttr", "ip_config.type6", *s.IpConfig.Type6)
		}
		if s.IpConfig.VlanId != nil {
			checks.append(t, "TestCheckResourceAttr", "ip_config.vlan_id", fmt.Sprintf("%d", *s.IpConfig.VlanId))
		}
	}
	if s.Led != nil {
		checks.append(t, "TestCheckResourceAttrSet", "led.%")
		if s.Led.Brightness != nil {
			checks.append(t, "TestCheckResourceAttr", "led.brightness", fmt.Sprintf("%d", *s.Led.Brightness))
		}
		if s.Led.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "led.enabled", fmt.Sprintf("%t", *s.Led.Enabled))
		}
	}
	if s.Mesh != nil {
		checks.append(t, "TestCheckResourceAttrSet", "mesh.%")
		if len(s.Mesh.Bands) > 0 {
			checks.append(t, "TestCheckResourceAttr", "mesh.bands.#", fmt.Sprintf("%d", len(s.Mesh.Bands)))
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
	if len(s.NtpServers) > 0 {
		checks.append(t, "TestCheckResourceAttr", "ntp_servers.#", fmt.Sprintf("%d", len(s.NtpServers)))
		for i, server := range s.NtpServers {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ntp_servers.%d", i), server)
		}
	}
	if s.PoePassthrough != nil {
		checks.append(t, "TestCheckResourceAttr", "poe_passthrough", fmt.Sprintf("%t", *s.PoePassthrough))
	}
	if s.PwrConfig != nil {
		checks.append(t, "TestCheckResourceAttrSet", "pwr_config.%")
		if s.PwrConfig.Base != nil {
			checks.append(t, "TestCheckResourceAttr", "pwr_config.base", fmt.Sprintf("%d", *s.PwrConfig.Base))
		}
		if s.PwrConfig.PreferUsbOverWifi != nil {
			checks.append(t, "TestCheckResourceAttr", "pwr_config.prefer_usb_over_wifi", fmt.Sprintf("%t", *s.PwrConfig.PreferUsbOverWifi))
		}
	}
	if s.RadioConfig != nil {
		checks.append(t, "TestCheckResourceAttrSet", "radio_config.%")
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
		if s.RadioConfig.RrmManaged != nil {
			checks.append(t, "TestCheckResourceAttr", "radio_config.rrm_managed", fmt.Sprintf("%t", *s.RadioConfig.RrmManaged))
		}
		if s.RadioConfig.AntennaMode != nil {
			checks.append(t, "TestCheckResourceAttr", "radio_config.antenna_mode", *s.RadioConfig.AntennaMode)
		}
		if s.RadioConfig.Band24 != nil {
			checks.append(t, "TestCheckResourceAttrSet", "radio_config.band_24.%")
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
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_24.channels.#", fmt.Sprintf("%d", len(s.RadioConfig.Band24.Channels)))
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
		if s.RadioConfig.Band24Usage != nil {
			checks.append(t, "TestCheckResourceAttr", "radio_config.band_24_usage", *s.RadioConfig.Band24Usage)
		}
		if s.RadioConfig.Band5 != nil {
			checks.append(t, "TestCheckResourceAttrSet", "radio_config.band_5.%")
			if s.RadioConfig.Band5.AllowRrmDisable != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_5.allow_rrm_disable", fmt.Sprintf("%t", *s.RadioConfig.Band5.AllowRrmDisable))
			}
			if s.RadioConfig.Band5.AntGain != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_5.ant_gain", fmt.Sprintf("%d", *s.RadioConfig.Band5.AntGain))
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
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_5.channels.#", fmt.Sprintf("%d", len(s.RadioConfig.Band5.Channels)))
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
			if s.RadioConfig.Band5.AntennaBeamPattern != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_5.antenna_beam_pattern", *s.RadioConfig.Band5.AntennaBeamPattern)
			}
		}
		if s.RadioConfig.Band5On24Radio != nil {
			checks.append(t, "TestCheckResourceAttrSet", "radio_config.band_5_on_24_radio.%")
			if s.RadioConfig.Band5On24Radio.AllowRrmDisable != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_5_on_24_radio.allow_rrm_disable", fmt.Sprintf("%t", *s.RadioConfig.Band5On24Radio.AllowRrmDisable))
			}
			if s.RadioConfig.Band5On24Radio.AntGain != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_5_on_24_radio.ant_gain", fmt.Sprintf("%d", *s.RadioConfig.Band5On24Radio.AntGain))
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
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_5_on_24_radio.channels.#", fmt.Sprintf("%d", len(s.RadioConfig.Band5On24Radio.Channels)))
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
			if s.RadioConfig.Band5On24Radio.AntennaBeamPattern != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_5_on_24_radio.antenna_beam_pattern", *s.RadioConfig.Band5On24Radio.AntennaBeamPattern)
			}
		}
		if s.RadioConfig.Band6 != nil {
			checks.append(t, "TestCheckResourceAttrSet", "radio_config.band_6.%")
			if s.RadioConfig.Band6.AllowRrmDisable != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_6.allow_rrm_disable", fmt.Sprintf("%t", *s.RadioConfig.Band6.AllowRrmDisable))
			}
			if s.RadioConfig.Band6.AntGain != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_6.ant_gain", fmt.Sprintf("%d", *s.RadioConfig.Band6.AntGain))
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
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_6.channels.#", fmt.Sprintf("%d", len(s.RadioConfig.Band6.Channels)))
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
			if s.RadioConfig.Band6.AntennaBeamPattern != nil {
				checks.append(t, "TestCheckResourceAttr", "radio_config.band_6.antenna_beam_pattern", *s.RadioConfig.Band6.AntennaBeamPattern)
			}
		}
		if s.RadioConfig.FullAutomaticRrm != nil {
			checks.append(t, "TestCheckResourceAttr", "radio_config.full_automatic_rrm", fmt.Sprintf("%t", *s.RadioConfig.FullAutomaticRrm))
		}
		if s.RadioConfig.IndoorUse != nil {
			checks.append(t, "TestCheckResourceAttr", "radio_config.indoor_use", fmt.Sprintf("%t", *s.RadioConfig.IndoorUse))
		}
		if s.RadioConfig.ScanningEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "radio_config.scanning_enabled", fmt.Sprintf("%t", *s.RadioConfig.ScanningEnabled))
		}
	}
	if s.SiteId != nil {
		checks.append(t, "TestCheckResourceAttr", "site_id", *s.SiteId)
	}
	if s.UplinkPortConfig != nil {
		checks.append(t, "TestCheckResourceAttrSet", "uplink_port_config.%")
		if s.UplinkPortConfig.Dot1x != nil {
			checks.append(t, "TestCheckResourceAttr", "uplink_port_config.dot1x", fmt.Sprintf("%t", *s.UplinkPortConfig.Dot1x))
		}
		if s.UplinkPortConfig.KeepWlansUpIfDown != nil {
			checks.append(t, "TestCheckResourceAttr", "uplink_port_config.keep_wlans_up_if_down", fmt.Sprintf("%t", *s.UplinkPortConfig.KeepWlansUpIfDown))
		}
	}
	if s.UsbConfig != nil {
		checks.append(t, "TestCheckResourceAttrSet", "usb_config.%")
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
	if s.Airista != nil {
		checks.append(t, "TestCheckResourceAttrSet", "airista.%")
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
	if s.LacpConfig != nil {
		checks.append(t, "TestCheckResourceAttrSet", "lacp_config.%")
		if s.LacpConfig.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "lacp_config.enabled", fmt.Sprintf("%t", *s.LacpConfig.Enabled))
		}
	}
	if len(s.PortConfig) > 0 {
		checks.append(t, "TestCheckResourceAttr", "port_config.%", fmt.Sprintf("%d", len(s.PortConfig)))
		for portName, portCfg := range s.PortConfig {
			portPrefix := fmt.Sprintf("port_config.%s", portName)
			checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("%s.%%", portPrefix))

			if portCfg.Disabled != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.disabled", portPrefix), fmt.Sprintf("%t", *portCfg.Disabled))
			}
			if portCfg.Forwarding != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.forwarding", portPrefix), *portCfg.Forwarding)
			}
			if portCfg.EnableMacAuth != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.enable_mac_auth", portPrefix), fmt.Sprintf("%t", *portCfg.EnableMacAuth))
			}
			if portCfg.MacAuthPreferred != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.mac_auth_preferred", portPrefix), fmt.Sprintf("%t", *portCfg.MacAuthPreferred))
			}
			if portCfg.MacAuthProtocol != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.mac_auth_protocol", portPrefix), *portCfg.MacAuthProtocol)
			}
			if portCfg.MxTunnelId != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.mx_tunnel_id", portPrefix), *portCfg.MxTunnelId)
			}
			if portCfg.MxtunnelName != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.mxtunnel_name", portPrefix), *portCfg.MxtunnelName)
			}
			if portCfg.PortAuth != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.port_auth", portPrefix), *portCfg.PortAuth)
			}
			if portCfg.PortVlanId != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.port_vlan_id", portPrefix), fmt.Sprintf("%d", *portCfg.PortVlanId))
			}
			if portCfg.VlanId != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.vlan_id", portPrefix), fmt.Sprintf("%d", *portCfg.VlanId))
			}
			if portCfg.VlanIds != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.vlan_ids", portPrefix), *portCfg.VlanIds)
			}
			if portCfg.WxtunnelId != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.wxtunnel_id", portPrefix), *portCfg.WxtunnelId)
			}
			if portCfg.WxtunnelRemoteId != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.wxtunnel_remote_id", portPrefix), *portCfg.WxtunnelRemoteId)
			}

			// Dynamic VLAN checks
			if portCfg.DynamicVlan != nil {
				dynVlanPrefix := fmt.Sprintf("%s.dynamic_vlan", portPrefix)
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("%s.%%", dynVlanPrefix))
				if portCfg.DynamicVlan.Enabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.enabled", dynVlanPrefix), fmt.Sprintf("%t", *portCfg.DynamicVlan.Enabled))
				}
				if portCfg.DynamicVlan.DefaultVlanId != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.default_vlan_id", dynVlanPrefix), fmt.Sprintf("%d", *portCfg.DynamicVlan.DefaultVlanId))
				}
				if portCfg.DynamicVlan.DynamicVlanType != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.type", dynVlanPrefix), *portCfg.DynamicVlan.DynamicVlanType)
				}
				if len(portCfg.DynamicVlan.Vlans) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.vlans.%%", dynVlanPrefix), fmt.Sprintf("%d", len(portCfg.DynamicVlan.Vlans)))
					for key, value := range portCfg.DynamicVlan.Vlans {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.vlans.%s", dynVlanPrefix, key), value)
					}
				}
			}

			// Mist NAC checks
			if portCfg.MistNac != nil {
				mistNacPrefix := fmt.Sprintf("%s.mist_nac", portPrefix)
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("%s.%%", mistNacPrefix))
				if portCfg.MistNac.Enabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.enabled", mistNacPrefix), fmt.Sprintf("%t", *portCfg.MistNac.Enabled))
				}
				if portCfg.MistNac.AcctInterimInterval != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.acct_interim_interval", mistNacPrefix), fmt.Sprintf("%d", *portCfg.MistNac.AcctInterimInterval))
				}
				if portCfg.MistNac.AuthServersRetries != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.auth_servers_retries", mistNacPrefix), fmt.Sprintf("%d", *portCfg.MistNac.AuthServersRetries))
				}
				if portCfg.MistNac.AuthServersTimeout != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.auth_servers_timeout", mistNacPrefix), fmt.Sprintf("%d", *portCfg.MistNac.AuthServersTimeout))
				}
				if portCfg.MistNac.CoaEnabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.coa_enabled", mistNacPrefix), fmt.Sprintf("%t", *portCfg.MistNac.CoaEnabled))
				}
				if portCfg.MistNac.CoaPort != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.coa_port", mistNacPrefix), fmt.Sprintf("%d", *portCfg.MistNac.CoaPort))
				}
				if portCfg.MistNac.FastDot1xTimers != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.fast_dot1x_timers", mistNacPrefix), fmt.Sprintf("%t", *portCfg.MistNac.FastDot1xTimers))
				}
				if portCfg.MistNac.Network != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.network", mistNacPrefix), *portCfg.MistNac.Network)
				}
				if portCfg.MistNac.SourceIp != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.source_ip", mistNacPrefix), *portCfg.MistNac.SourceIp)
				}
			}

			// RADIUS Config checks
			if portCfg.RadiusConfig != nil {
				radiusPrefix := fmt.Sprintf("%s.radius_config", portPrefix)
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("%s.%%", radiusPrefix))
				if portCfg.RadiusConfig.AcctInterimInterval != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.acct_interim_interval", radiusPrefix), fmt.Sprintf("%d", *portCfg.RadiusConfig.AcctInterimInterval))
				}
				if portCfg.RadiusConfig.AuthServersRetries != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.auth_servers_retries", radiusPrefix), fmt.Sprintf("%d", *portCfg.RadiusConfig.AuthServersRetries))
				}
				if portCfg.RadiusConfig.AuthServersTimeout != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.auth_servers_timeout", radiusPrefix), fmt.Sprintf("%d", *portCfg.RadiusConfig.AuthServersTimeout))
				}
				if portCfg.RadiusConfig.CoaEnabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.coa_enabled", radiusPrefix), fmt.Sprintf("%t", *portCfg.RadiusConfig.CoaEnabled))
				}
				if portCfg.RadiusConfig.CoaPort != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.coa_port", radiusPrefix), fmt.Sprintf("%d", *portCfg.RadiusConfig.CoaPort))
				}
				if portCfg.RadiusConfig.Network != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.network", radiusPrefix), *portCfg.RadiusConfig.Network)
				}
				if portCfg.RadiusConfig.SourceIp != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.source_ip", radiusPrefix), *portCfg.RadiusConfig.SourceIp)
				}

				// Auth servers checks
				if len(portCfg.RadiusConfig.AuthServers) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.auth_servers.#", radiusPrefix), fmt.Sprintf("%d", len(portCfg.RadiusConfig.AuthServers)))
					for i, authSrv := range portCfg.RadiusConfig.AuthServers {
						authSrvPrefix := fmt.Sprintf("%s.auth_servers.%d", radiusPrefix, i)
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.host", authSrvPrefix), authSrv.Host)
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.secret", authSrvPrefix), authSrv.Secret)
						if authSrv.Port != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.port", authSrvPrefix), *authSrv.Port)
						}
						if authSrv.KeywrapEnabled != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.keywrap_enabled", authSrvPrefix), fmt.Sprintf("%t", *authSrv.KeywrapEnabled))
						}
						if authSrv.KeywrapFormat != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.keywrap_format", authSrvPrefix), *authSrv.KeywrapFormat)
						}
						if authSrv.KeywrapKek != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.keywrap_kek", authSrvPrefix), *authSrv.KeywrapKek)
						}
						if authSrv.KeywrapMack != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.keywrap_mack", authSrvPrefix), *authSrv.KeywrapMack)
						}
						if authSrv.RequireMessageAuthenticator != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.require_message_authenticator", authSrvPrefix), fmt.Sprintf("%t", *authSrv.RequireMessageAuthenticator))
						}
					}
				}

				// Acct servers checks
				if len(portCfg.RadiusConfig.AcctServers) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.acct_servers.#", radiusPrefix), fmt.Sprintf("%d", len(portCfg.RadiusConfig.AcctServers)))
					for i, acctSrv := range portCfg.RadiusConfig.AcctServers {
						acctSrvPrefix := fmt.Sprintf("%s.acct_servers.%d", radiusPrefix, i)
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.host", acctSrvPrefix), acctSrv.Host)
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.secret", acctSrvPrefix), acctSrv.Secret)
						if acctSrv.Port != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.port", acctSrvPrefix), *acctSrv.Port)
						}
						if acctSrv.KeywrapEnabled != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.keywrap_enabled", acctSrvPrefix), fmt.Sprintf("%t", *acctSrv.KeywrapEnabled))
						}
						if acctSrv.KeywrapFormat != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.keywrap_format", acctSrvPrefix), *acctSrv.KeywrapFormat)
						}
						if acctSrv.KeywrapKek != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.keywrap_kek", acctSrvPrefix), *acctSrv.KeywrapKek)
						}
						if acctSrv.KeywrapMack != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.keywrap_mack", acctSrvPrefix), *acctSrv.KeywrapMack)
						}
					}
				}
			}

			// RADSEC checks
			if portCfg.Radsec != nil {
				radsecPrefix := fmt.Sprintf("%s.radsec", portPrefix)
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("%s.%%", radsecPrefix))
				if portCfg.Radsec.Enabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.enabled", radsecPrefix), fmt.Sprintf("%t", *portCfg.Radsec.Enabled))
				}
				if portCfg.Radsec.CoaEnabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.coa_enabled", radsecPrefix), fmt.Sprintf("%t", *portCfg.Radsec.CoaEnabled))
				}
				if portCfg.Radsec.IdleTimeout != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.idle_timeout", radsecPrefix), *portCfg.Radsec.IdleTimeout)
				}
				if portCfg.Radsec.ServerName != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.server_name", radsecPrefix), *portCfg.Radsec.ServerName)
				}
				if portCfg.Radsec.UseMxedge != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.use_mxedge", radsecPrefix), fmt.Sprintf("%t", *portCfg.Radsec.UseMxedge))
				}
				if portCfg.Radsec.UseSiteMxedge != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.use_site_mxedge", radsecPrefix), fmt.Sprintf("%t", *portCfg.Radsec.UseSiteMxedge))
				}
				if len(portCfg.Radsec.MxclusterIds) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.mxcluster_ids.#", radsecPrefix), fmt.Sprintf("%d", len(portCfg.Radsec.MxclusterIds)))
					for i, clusterId := range portCfg.Radsec.MxclusterIds {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.mxcluster_ids.%d", radsecPrefix, i), clusterId)
					}
				}
				if len(portCfg.Radsec.ProxyHosts) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.proxy_hosts.#", radsecPrefix), fmt.Sprintf("%d", len(portCfg.Radsec.ProxyHosts)))
					for i, proxyHost := range portCfg.Radsec.ProxyHosts {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.proxy_hosts.%d", radsecPrefix, i), proxyHost)
					}
				}
				if len(portCfg.Radsec.Servers) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.servers.#", radsecPrefix), fmt.Sprintf("%d", len(portCfg.Radsec.Servers)))
					for i, server := range portCfg.Radsec.Servers {
						serverPrefix := fmt.Sprintf("%s.servers.%d", radsecPrefix, i)
						if server.Host != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.host", serverPrefix), *server.Host)
						}
						if server.Port != nil {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("%s.port", serverPrefix), fmt.Sprintf("%d", *server.Port))
						}
					}
				}
			}
		}
	}
	if len(s.Vars) > 0 {
		checks.append(t, "TestCheckResourceAttr", "vars.%", fmt.Sprintf("%d", len(s.Vars)))
		for key, value := range s.Vars {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vars.%s", key), value)
		}
	}

	return checks
}
