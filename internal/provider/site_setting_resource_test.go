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

func TestSiteSettingModel(t *testing.T) {
	type testStep struct {
		config SiteSettingModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: SiteSettingModel{},
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
		var FixtureSiteSettingModel SiteSettingModel

		err = hcl.Decode(&FixtureSiteSettingModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureSiteSettingModel,
				},
			},
		}
	}

	resourceType := "site_setting"
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

func (s *SiteSettingModel) testChecks(t testing.TB, rType, rName string) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + rName)
	checks.append(t, "TestCheckResourceAttrSet", "site_id")

	// Conditional checks for optional parameters
	if s.Analytic != nil {
		if s.Analytic.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "analytic.enabled", fmt.Sprintf("%t", *s.Analytic.Enabled))
		}
	}
	if s.ApUpdownThreshold != nil {
		checks.append(t, "TestCheckResourceAttr", "ap_updown_threshold", fmt.Sprintf("%d", *s.ApUpdownThreshold))
	}
	if s.AutoUpgrade != nil {
		if len(s.AutoUpgrade.CustomVersions) > 0 {
			for key, value := range s.AutoUpgrade.CustomVersions {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("auto_upgrade.custom_versions.%s", key), value)
			}
		}
		if s.AutoUpgrade.DayOfWeek != nil {
			checks.append(t, "TestCheckResourceAttr", "auto_upgrade.day_of_week", *s.AutoUpgrade.DayOfWeek)
		}
		if s.AutoUpgrade.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "auto_upgrade.enabled", fmt.Sprintf("%t", *s.AutoUpgrade.Enabled))
		}
		if s.AutoUpgrade.TimeOfDay != nil {
			checks.append(t, "TestCheckResourceAttr", "auto_upgrade.time_of_day", *s.AutoUpgrade.TimeOfDay)
		}
		if s.AutoUpgrade.Version != nil {
			checks.append(t, "TestCheckResourceAttr", "auto_upgrade.version", *s.AutoUpgrade.Version)
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
	if s.ConfigAutoRevert != nil {
		checks.append(t, "TestCheckResourceAttr", "config_auto_revert", fmt.Sprintf("%t", *s.ConfigAutoRevert))
	}
	if s.ConfigPushPolicy != nil {
		if s.ConfigPushPolicy.NoPush != nil {
			checks.append(t, "TestCheckResourceAttr", "config_push_policy.no_push", fmt.Sprintf("%t", *s.ConfigPushPolicy.NoPush))
		}
		if s.ConfigPushPolicy.PushWindow != nil {
			if s.ConfigPushPolicy.PushWindow.Enabled != nil {
				checks.append(t, "TestCheckResourceAttr", "config_push_policy.push_window.enabled", fmt.Sprintf("%t", *s.ConfigPushPolicy.PushWindow.Enabled))
			}
			if s.ConfigPushPolicy.PushWindow.Hours != nil {
				if s.ConfigPushPolicy.PushWindow.Hours.Fri != nil {
					checks.append(t, "TestCheckResourceAttr", "config_push_policy.push_window.hours.fri", *s.ConfigPushPolicy.PushWindow.Hours.Fri)
				}
				if s.ConfigPushPolicy.PushWindow.Hours.Mon != nil {
					checks.append(t, "TestCheckResourceAttr", "config_push_policy.push_window.hours.mon", *s.ConfigPushPolicy.PushWindow.Hours.Mon)
				}
				if s.ConfigPushPolicy.PushWindow.Hours.Sat != nil {
					checks.append(t, "TestCheckResourceAttr", "config_push_policy.push_window.hours.sat", *s.ConfigPushPolicy.PushWindow.Hours.Sat)
				}
				if s.ConfigPushPolicy.PushWindow.Hours.Sun != nil {
					checks.append(t, "TestCheckResourceAttr", "config_push_policy.push_window.hours.sun", *s.ConfigPushPolicy.PushWindow.Hours.Sun)
				}
				if s.ConfigPushPolicy.PushWindow.Hours.Thu != nil {
					checks.append(t, "TestCheckResourceAttr", "config_push_policy.push_window.hours.thu", *s.ConfigPushPolicy.PushWindow.Hours.Thu)
				}
				if s.ConfigPushPolicy.PushWindow.Hours.Tue != nil {
					checks.append(t, "TestCheckResourceAttr", "config_push_policy.push_window.hours.tue", *s.ConfigPushPolicy.PushWindow.Hours.Tue)
				}
				if s.ConfigPushPolicy.PushWindow.Hours.Wed != nil {
					checks.append(t, "TestCheckResourceAttr", "config_push_policy.push_window.hours.wed", *s.ConfigPushPolicy.PushWindow.Hours.Wed)
				}
			}
		}
	}
	if s.CriticalUrlMonitoring != nil {
		if s.CriticalUrlMonitoring.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "critical_url_monitoring.enabled", fmt.Sprintf("%t", *s.CriticalUrlMonitoring.Enabled))
		}
		if len(s.CriticalUrlMonitoring.Monitors) > 0 {
			checks.append(t, "TestCheckResourceAttr", "critical_url_monitoring.monitors.#", fmt.Sprintf("%d", len(s.CriticalUrlMonitoring.Monitors)))
			for i, monitor := range s.CriticalUrlMonitoring.Monitors {
				if monitor.Url != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("critical_url_monitoring.monitors.%d.url", i), *monitor.Url)
				}
				if monitor.VlanId != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("critical_url_monitoring.monitors.%d.vlan_id", i), *monitor.VlanId)
				}
			}
		}
	}
	if s.DeviceUpdownThreshold != nil {
		checks.append(t, "TestCheckResourceAttr", "device_updown_threshold", fmt.Sprintf("%d", *s.DeviceUpdownThreshold))
	}
	if s.Engagement != nil {
		if s.Engagement.DwellTagNames != nil {
			if s.Engagement.DwellTagNames.Bounce != nil {
				checks.append(t, "TestCheckResourceAttr", "engagement.dwell_tag_names.bounce", *s.Engagement.DwellTagNames.Bounce)
			}
			if s.Engagement.DwellTagNames.Engaged != nil {
				checks.append(t, "TestCheckResourceAttr", "engagement.dwell_tag_names.engaged", *s.Engagement.DwellTagNames.Engaged)
			}
			if s.Engagement.DwellTagNames.Passerby != nil {
				checks.append(t, "TestCheckResourceAttr", "engagement.dwell_tag_names.passerby", *s.Engagement.DwellTagNames.Passerby)
			}
			if s.Engagement.DwellTagNames.Stationed != nil {
				checks.append(t, "TestCheckResourceAttr", "engagement.dwell_tag_names.stationed", *s.Engagement.DwellTagNames.Stationed)
			}
		}
		if s.Engagement.DwellTags != nil {
			if s.Engagement.DwellTags.Bounce != nil {
				checks.append(t, "TestCheckResourceAttr", "engagement.dwell_tags.bounce", *s.Engagement.DwellTags.Bounce)
			}
			if s.Engagement.DwellTags.Engaged != nil {
				checks.append(t, "TestCheckResourceAttr", "engagement.dwell_tags.engaged", *s.Engagement.DwellTags.Engaged)
			}
			if s.Engagement.DwellTags.Passerby != nil {
				checks.append(t, "TestCheckResourceAttr", "engagement.dwell_tags.passerby", *s.Engagement.DwellTags.Passerby)
			}
			if s.Engagement.DwellTags.Stationed != nil {
				checks.append(t, "TestCheckResourceAttr", "engagement.dwell_tags.stationed", *s.Engagement.DwellTags.Stationed)
			}
		}
		if s.Engagement.Hours != nil {
			if s.Engagement.Hours.Fri != nil {
				checks.append(t, "TestCheckResourceAttr", "engagement.hours.fri", *s.Engagement.Hours.Fri)
			}
			if s.Engagement.Hours.Mon != nil {
				checks.append(t, "TestCheckResourceAttr", "engagement.hours.mon", *s.Engagement.Hours.Mon)
			}
			if s.Engagement.Hours.Sat != nil {
				checks.append(t, "TestCheckResourceAttr", "engagement.hours.sat", *s.Engagement.Hours.Sat)
			}
			if s.Engagement.Hours.Sun != nil {
				checks.append(t, "TestCheckResourceAttr", "engagement.hours.sun", *s.Engagement.Hours.Sun)
			}
			if s.Engagement.Hours.Thu != nil {
				checks.append(t, "TestCheckResourceAttr", "engagement.hours.thu", *s.Engagement.Hours.Thu)
			}
			if s.Engagement.Hours.Tue != nil {
				checks.append(t, "TestCheckResourceAttr", "engagement.hours.tue", *s.Engagement.Hours.Tue)
			}
			if s.Engagement.Hours.Wed != nil {
				checks.append(t, "TestCheckResourceAttr", "engagement.hours.wed", *s.Engagement.Hours.Wed)
			}
		}
		if s.Engagement.MaxDwell != nil {
			checks.append(t, "TestCheckResourceAttr", "engagement.max_dwell", fmt.Sprintf("%d", *s.Engagement.MaxDwell))
		}
		if s.Engagement.MinDwell != nil {
			checks.append(t, "TestCheckResourceAttr", "engagement.min_dwell", fmt.Sprintf("%d", *s.Engagement.MinDwell))
		}
	}
	if s.GatewayMgmt != nil {
		if len(s.GatewayMgmt.AdminSshkeys) > 0 {
			checks.append(t, "TestCheckResourceAttr", "gateway_mgmt.admin_sshkeys.#", fmt.Sprintf("%d", len(s.GatewayMgmt.AdminSshkeys)))
			for i, key := range s.GatewayMgmt.AdminSshkeys {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("gateway_mgmt.admin_sshkeys.%d", i), key)
			}
		}
		if s.GatewayMgmt.AppProbing != nil {
			if len(s.GatewayMgmt.AppProbing.Apps) > 0 {
				checks.append(t, "TestCheckResourceAttr", "gateway_mgmt.app_probing.apps.#", fmt.Sprintf("%d", len(s.GatewayMgmt.AppProbing.Apps)))
				for i, app := range s.GatewayMgmt.AppProbing.Apps {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("gateway_mgmt.app_probing.apps.%d", i), app)
				}
			}
			if len(s.GatewayMgmt.AppProbing.CustomApps) > 0 {
				checks.append(t, "TestCheckResourceAttr", "gateway_mgmt.app_probing.custom_apps.#", fmt.Sprintf("%d", len(s.GatewayMgmt.AppProbing.CustomApps)))
				for i, customApp := range s.GatewayMgmt.AppProbing.CustomApps {
					if customApp.AppType != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("gateway_mgmt.app_probing.custom_apps.%d.app_type", i), *customApp.AppType)
					}
					if len(customApp.Hostnames) > 0 {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("gateway_mgmt.app_probing.custom_apps.%d.hostnames.#", i), fmt.Sprintf("%d", len(customApp.Hostnames)))
						for j, hostname := range customApp.Hostnames {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("gateway_mgmt.app_probing.custom_apps.%d.hostnames.%d", i, j), hostname)
						}
					}
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("gateway_mgmt.app_probing.custom_apps.%d.name", i), customApp.Name)
					if customApp.Network != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("gateway_mgmt.app_probing.custom_apps.%d.network", i), *customApp.Network)
					}
					if customApp.PacketSize != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("gateway_mgmt.app_probing.custom_apps.%d.packet_size", i), fmt.Sprintf("%d", *customApp.PacketSize))
					}
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("gateway_mgmt.app_probing.custom_apps.%d.protocol", i), customApp.Protocol)
					if customApp.Vrf != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("gateway_mgmt.app_probing.custom_apps.%d.vrf", i), *customApp.Vrf)
					}
				}
			}
			if s.GatewayMgmt.AppProbing.Enabled != nil {
				checks.append(t, "TestCheckResourceAttr", "gateway_mgmt.app_probing.enabled", fmt.Sprintf("%t", *s.GatewayMgmt.AppProbing.Enabled))
			}
		}
		if s.GatewayMgmt.AppUsage != nil {
			checks.append(t, "TestCheckResourceAttr", "gateway_mgmt.app_usage", fmt.Sprintf("%t", *s.GatewayMgmt.AppUsage))
		}
		if s.GatewayMgmt.AutoSignatureUpdate != nil {
			if s.GatewayMgmt.AutoSignatureUpdate.DayOfWeek != nil {
				checks.append(t, "TestCheckResourceAttr", "gateway_mgmt.auto_signature_update.day_of_week", *s.GatewayMgmt.AutoSignatureUpdate.DayOfWeek)
			}
			if s.GatewayMgmt.AutoSignatureUpdate.Enable != nil {
				checks.append(t, "TestCheckResourceAttr", "gateway_mgmt.auto_signature_update.enable", fmt.Sprintf("%t", *s.GatewayMgmt.AutoSignatureUpdate.Enable))
			}
			if s.GatewayMgmt.AutoSignatureUpdate.TimeOfDay != nil {
				checks.append(t, "TestCheckResourceAttr", "gateway_mgmt.auto_signature_update.time_of_day", *s.GatewayMgmt.AutoSignatureUpdate.TimeOfDay)
			}
		}
		if s.GatewayMgmt.ConfigRevertTimer != nil {
			checks.append(t, "TestCheckResourceAttr", "gateway_mgmt.config_revert_timer", fmt.Sprintf("%d", *s.GatewayMgmt.ConfigRevertTimer))
		}
		if s.GatewayMgmt.DisableConsole != nil {
			checks.append(t, "TestCheckResourceAttr", "gateway_mgmt.disable_console", fmt.Sprintf("%t", *s.GatewayMgmt.DisableConsole))
		}
		if s.GatewayMgmt.DisableOob != nil {
			checks.append(t, "TestCheckResourceAttr", "gateway_mgmt.disable_oob", fmt.Sprintf("%t", *s.GatewayMgmt.DisableOob))
		}
		if len(s.GatewayMgmt.ProbeHosts) > 0 {
			checks.append(t, "TestCheckResourceAttr", "gateway_mgmt.probe_hosts.#", fmt.Sprintf("%d", len(s.GatewayMgmt.ProbeHosts)))
			for i, host := range s.GatewayMgmt.ProbeHosts {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("gateway_mgmt.probe_hosts.%d", i), host)
			}
		}
		if s.GatewayMgmt.ProtectRe != nil {
			if len(s.GatewayMgmt.ProtectRe.AllowedServices) > 0 {
				checks.append(t, "TestCheckResourceAttr", "gateway_mgmt.protect_re.allowed_services.#", fmt.Sprintf("%d", len(s.GatewayMgmt.ProtectRe.AllowedServices)))
				for i, service := range s.GatewayMgmt.ProtectRe.AllowedServices {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("gateway_mgmt.protect_re.allowed_services.%d", i), service)
				}
			}
			if len(s.GatewayMgmt.ProtectRe.Custom) > 0 {
				checks.append(t, "TestCheckResourceAttr", "gateway_mgmt.protect_re.custom.#", fmt.Sprintf("%d", len(s.GatewayMgmt.ProtectRe.Custom)))
				for i, custom := range s.GatewayMgmt.ProtectRe.Custom {
					if custom.PortRange != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("gateway_mgmt.protect_re.custom.%d.port_range", i), *custom.PortRange)
					}
					if custom.Protocol != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("gateway_mgmt.protect_re.custom.%d.protocol", i), *custom.Protocol)
					}
					if len(custom.Subnets) > 0 {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("gateway_mgmt.protect_re.custom.%d.subnets.#", i), fmt.Sprintf("%d", len(custom.Subnets)))
						for j, subnet := range custom.Subnets {
							checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("gateway_mgmt.protect_re.custom.%d.subnets.%d", i, j), subnet)
						}
					}
				}
			}
			if s.GatewayMgmt.ProtectRe.Enabled != nil {
				checks.append(t, "TestCheckResourceAttr", "gateway_mgmt.protect_re.enabled", fmt.Sprintf("%t", *s.GatewayMgmt.ProtectRe.Enabled))
			}
			if s.GatewayMgmt.ProtectRe.HitCount != nil {
				checks.append(t, "TestCheckResourceAttr", "gateway_mgmt.protect_re.hit_count", fmt.Sprintf("%t", *s.GatewayMgmt.ProtectRe.HitCount))
			}
			if len(s.GatewayMgmt.ProtectRe.TrustedHosts) > 0 {
				checks.append(t, "TestCheckResourceAttr", "gateway_mgmt.protect_re.trusted_hosts.#", fmt.Sprintf("%d", len(s.GatewayMgmt.ProtectRe.TrustedHosts)))
				for i, host := range s.GatewayMgmt.ProtectRe.TrustedHosts {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("gateway_mgmt.protect_re.trusted_hosts.%d", i), host)
				}
			}
		}
		if s.GatewayMgmt.RootPassword != nil {
			checks.append(t, "TestCheckResourceAttr", "gateway_mgmt.root_password", *s.GatewayMgmt.RootPassword)
		}
		if s.GatewayMgmt.SecurityLogSourceAddress != nil {
			checks.append(t, "TestCheckResourceAttr", "gateway_mgmt.security_log_source_address", *s.GatewayMgmt.SecurityLogSourceAddress)
		}
		if s.GatewayMgmt.SecurityLogSourceInterface != nil {
			checks.append(t, "TestCheckResourceAttr", "gateway_mgmt.security_log_source_interface", *s.GatewayMgmt.SecurityLogSourceInterface)
		}
	}
	if s.GatewayUpdownThreshold != nil {
		checks.append(t, "TestCheckResourceAttr", "gateway_updown_threshold", fmt.Sprintf("%d", *s.GatewayUpdownThreshold))
	}
	if s.JuniperSrx != nil {
		if len(s.JuniperSrx.Gateways) > 0 {
			checks.append(t, "TestCheckResourceAttr", "juniper_srx.gateways.#", fmt.Sprintf("%d", len(s.JuniperSrx.Gateways)))
			for i, gateway := range s.JuniperSrx.Gateways {
				if gateway.ApiKey != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("juniper_srx.gateways.%d.api_key", i), *gateway.ApiKey)
				}
				if gateway.ApiUrl != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("juniper_srx.gateways.%d.api_url", i), *gateway.ApiUrl)
				}
			}
		}
		if s.JuniperSrx.SendMistNacUserInfo != nil {
			checks.append(t, "TestCheckResourceAttr", "juniper_srx.send_mist_nac_user_info", fmt.Sprintf("%t", *s.JuniperSrx.SendMistNacUserInfo))
		}
		if s.JuniperSrx.SrxAutoUpgrade != nil {
			if s.JuniperSrx.SrxAutoUpgrade.Enabled != nil {
				checks.append(t, "TestCheckResourceAttr", "juniper_srx.auto_upgrade.enabled", fmt.Sprintf("%t", *s.JuniperSrx.SrxAutoUpgrade.Enabled))
			}
			if s.JuniperSrx.SrxAutoUpgrade.Snapshot != nil {
				checks.append(t, "TestCheckResourceAttr", "juniper_srx.auto_upgrade.snapshot", fmt.Sprintf("%t", *s.JuniperSrx.SrxAutoUpgrade.Snapshot))
			}
			if len(s.JuniperSrx.SrxAutoUpgrade.CustomVersions) > 0 {
				checks.append(t, "TestCheckResourceAttr", "juniper_srx.auto_upgrade.custom_versions.%", fmt.Sprintf("%d", len(s.JuniperSrx.SrxAutoUpgrade.CustomVersions)))
				for key, version := range s.JuniperSrx.SrxAutoUpgrade.CustomVersions {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("juniper_srx.auto_upgrade.custom_versions.%s", key), version)
				}
			}
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
	if s.Occupancy != nil {
		if s.Occupancy.AssetsEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "occupancy.assets_enabled", fmt.Sprintf("%t", *s.Occupancy.AssetsEnabled))
		}
		if s.Occupancy.ClientsEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "occupancy.clients_enabled", fmt.Sprintf("%t", *s.Occupancy.ClientsEnabled))
		}
		if s.Occupancy.MinDuration != nil {
			checks.append(t, "TestCheckResourceAttr", "occupancy.min_duration", fmt.Sprintf("%d", *s.Occupancy.MinDuration))
		}
		if s.Occupancy.SdkclientsEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "occupancy.sdkclients_enabled", fmt.Sprintf("%t", *s.Occupancy.SdkclientsEnabled))
		}
		if s.Occupancy.UnconnectedClientsEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "occupancy.unconnected_clients_enabled", fmt.Sprintf("%t", *s.Occupancy.UnconnectedClientsEnabled))
		}
	}
	if s.PersistConfigOnDevice != nil {
		checks.append(t, "TestCheckResourceAttr", "persist_config_on_device", fmt.Sprintf("%t", *s.PersistConfigOnDevice))
	}
	if s.Proxy != nil && s.Proxy.Url != nil {
		checks.append(t, "TestCheckResourceAttr", "proxy.url", *s.Proxy.Url)
	}
	if s.RemoveExistingConfigs != nil {
		checks.append(t, "TestCheckResourceAttr", "remove_existing_configs", fmt.Sprintf("%t", *s.RemoveExistingConfigs))
	}
	if s.ReportGatt != nil {
		checks.append(t, "TestCheckResourceAttr", "report_gatt", fmt.Sprintf("%t", *s.ReportGatt))
	}
	if s.Rogue != nil {
		if s.Rogue.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "rogue.enabled", fmt.Sprintf("%t", *s.Rogue.Enabled))
		}
		if s.Rogue.HoneypotEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "rogue.honeypot_enabled", fmt.Sprintf("%t", *s.Rogue.HoneypotEnabled))
		}
		if s.Rogue.MinDuration != nil {
			checks.append(t, "TestCheckResourceAttr", "rogue.min_duration", fmt.Sprintf("%d", *s.Rogue.MinDuration))
		}
		if s.Rogue.MinRssi != nil {
			checks.append(t, "TestCheckResourceAttr", "rogue.min_rssi", fmt.Sprintf("%d", *s.Rogue.MinRssi))
		}
		if len(s.Rogue.WhitelistedBssids) > 0 {
			for i, bssid := range s.Rogue.WhitelistedBssids {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("rogue.whitelisted_bssids.%d", i), bssid)
			}
		}
		if len(s.Rogue.WhitelistedSsids) > 0 {
			for i, ssid := range s.Rogue.WhitelistedSsids {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("rogue.whitelisted_ssids.%d", i), ssid)
			}
		}
	}
	if s.Rtsa != nil {
		if s.Rtsa.AppWaking != nil {
			checks.append(t, "TestCheckResourceAttr", "rtsa.app_waking", fmt.Sprintf("%t", *s.Rtsa.AppWaking))
		}
		if s.Rtsa.DisableDeadReckoning != nil {
			checks.append(t, "TestCheckResourceAttr", "rtsa.disable_dead_reckoning", fmt.Sprintf("%t", *s.Rtsa.DisableDeadReckoning))
		}
		if s.Rtsa.DisablePressureSensor != nil {
			checks.append(t, "TestCheckResourceAttr", "rtsa.disable_pressure_sensor", fmt.Sprintf("%t", *s.Rtsa.DisablePressureSensor))
		}
		if s.Rtsa.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "rtsa.enabled", fmt.Sprintf("%t", *s.Rtsa.Enabled))
		}
		if s.Rtsa.TrackAsset != nil {
			checks.append(t, "TestCheckResourceAttr", "rtsa.track_asset", fmt.Sprintf("%t", *s.Rtsa.TrackAsset))
		}
	}
	if s.SimpleAlert != nil {
		if s.SimpleAlert.ArpFailure != nil {
			if s.SimpleAlert.ArpFailure.ClientCount != nil {
				checks.append(t, "TestCheckResourceAttr", "simple_alert.arp_failure.client_count", fmt.Sprintf("%d", *s.SimpleAlert.ArpFailure.ClientCount))
			}
			if s.SimpleAlert.ArpFailure.Duration != nil {
				checks.append(t, "TestCheckResourceAttr", "simple_alert.arp_failure.duration", fmt.Sprintf("%d", *s.SimpleAlert.ArpFailure.Duration))
			}
			if s.SimpleAlert.ArpFailure.IncidentCount != nil {
				checks.append(t, "TestCheckResourceAttr", "simple_alert.arp_failure.incident_count", fmt.Sprintf("%d", *s.SimpleAlert.ArpFailure.IncidentCount))
			}
		}
		if s.SimpleAlert.DhcpFailure != nil {
			if s.SimpleAlert.DhcpFailure.ClientCount != nil {
				checks.append(t, "TestCheckResourceAttr", "simple_alert.dhcp_failure.client_count", fmt.Sprintf("%d", *s.SimpleAlert.DhcpFailure.ClientCount))
			}
			if s.SimpleAlert.DhcpFailure.Duration != nil {
				checks.append(t, "TestCheckResourceAttr", "simple_alert.dhcp_failure.duration", fmt.Sprintf("%d", *s.SimpleAlert.DhcpFailure.Duration))
			}
			if s.SimpleAlert.DhcpFailure.IncidentCount != nil {
				checks.append(t, "TestCheckResourceAttr", "simple_alert.dhcp_failure.incident_count", fmt.Sprintf("%d", *s.SimpleAlert.DhcpFailure.IncidentCount))
			}
		}
		if s.SimpleAlert.DnsFailure != nil {
			if s.SimpleAlert.DnsFailure.ClientCount != nil {
				checks.append(t, "TestCheckResourceAttr", "simple_alert.dns_failure.client_count", fmt.Sprintf("%d", *s.SimpleAlert.DnsFailure.ClientCount))
			}
			if s.SimpleAlert.DnsFailure.Duration != nil {
				checks.append(t, "TestCheckResourceAttr", "simple_alert.dns_failure.duration", fmt.Sprintf("%d", *s.SimpleAlert.DnsFailure.Duration))
			}
			if s.SimpleAlert.DnsFailure.IncidentCount != nil {
				checks.append(t, "TestCheckResourceAttr", "simple_alert.dns_failure.incident_count", fmt.Sprintf("%d", *s.SimpleAlert.DnsFailure.IncidentCount))
			}
		}
	}
	if s.Skyatp != nil {
		if s.Skyatp.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "skyatp.enabled", fmt.Sprintf("%t", *s.Skyatp.Enabled))
		}
		if s.Skyatp.SendIpMacMapping != nil {
			checks.append(t, "TestCheckResourceAttr", "skyatp.send_ip_mac_mapping", fmt.Sprintf("%t", *s.Skyatp.SendIpMacMapping))
		}
	}
	if s.SrxApp != nil && s.SrxApp.Enabled != nil {
		checks.append(t, "TestCheckResourceAttr", "srx_app.enabled", fmt.Sprintf("%t", *s.SrxApp.Enabled))
	}
	if len(s.SshKeys) > 0 {
		for i, key := range s.SshKeys {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ssh_keys.%d", i), key)
		}
	}
	if s.Ssr != nil {
		if len(s.Ssr.ConductorHosts) > 0 {
			for i, host := range s.Ssr.ConductorHosts {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ssr.conductor_hosts.%d", i), host)
			}
		}
		if s.Ssr.ConductorToken != nil {
			checks.append(t, "TestCheckResourceAttr", "ssr.conductor_token", *s.Ssr.ConductorToken)
		}
		if s.Ssr.DisableStats != nil {
			checks.append(t, "TestCheckResourceAttr", "ssr.disable_stats", fmt.Sprintf("%t", *s.Ssr.DisableStats))
		}
		if s.Ssr.Proxy != nil && s.Ssr.Proxy.Url != nil {
			checks.append(t, "TestCheckResourceAttr", "ssr.proxy.url", *s.Ssr.Proxy.Url)
		}
		if s.Ssr.SsrAutoUpgrade != nil {
			if s.Ssr.SsrAutoUpgrade.Enabled != nil {
				checks.append(t, "TestCheckResourceAttr", "ssr.auto_upgrade.enabled", fmt.Sprintf("%t", *s.Ssr.SsrAutoUpgrade.Enabled))
			}
			if s.Ssr.SsrAutoUpgrade.Channel != nil {
				checks.append(t, "TestCheckResourceAttr", "ssr.auto_upgrade.channel", *s.Ssr.SsrAutoUpgrade.Channel)
			}
			if len(s.Ssr.SsrAutoUpgrade.CustomVersions) > 0 {
				checks.append(t, "TestCheckResourceAttr", "ssr.auto_upgrade.custom_versions.%", fmt.Sprintf("%d", len(s.Ssr.SsrAutoUpgrade.CustomVersions)))
				for key, version := range s.Ssr.SsrAutoUpgrade.CustomVersions {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ssr.auto_upgrade.custom_versions.%s", key), version)
				}
			}
		}
	}
	if s.SwitchUpdownThreshold != nil {
		checks.append(t, "TestCheckResourceAttr", "switch_updown_threshold", fmt.Sprintf("%d", *s.SwitchUpdownThreshold))
	}
	if s.SyntheticTest != nil {
		if s.SyntheticTest.Disabled != nil {
			checks.append(t, "TestCheckResourceAttr", "synthetic_test.disabled", fmt.Sprintf("%t", *s.SyntheticTest.Disabled))
		}
		if s.SyntheticTest.Aggressiveness != nil {
			checks.append(t, "TestCheckResourceAttr", "synthetic_test.aggressiveness", *s.SyntheticTest.Aggressiveness)
		}
		if len(s.SyntheticTest.CustomProbes) > 0 {
			for key, probe := range s.SyntheticTest.CustomProbes {
				if probe.CustomProbesType != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("synthetic_test.custom_probes.%s.custom_probes_type", key), *probe.CustomProbesType)
				}
				if probe.Host != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("synthetic_test.custom_probes.%s.host", key), *probe.Host)
				}
				if probe.Port != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("synthetic_test.custom_probes.%s.port", key), fmt.Sprintf("%d", *probe.Port))
				}
				if probe.Threshold != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("synthetic_test.custom_probes.%s.threshold", key), fmt.Sprintf("%d", *probe.Threshold))
				}
				if probe.Aggressiveness != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("synthetic_test.custom_probes.%s.aggressiveness", key), *probe.Aggressiveness)
				}
				if probe.Url != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("synthetic_test.custom_probes.%s.url", key), *probe.Url)
				}
			}
		}
		if len(s.SyntheticTest.LanNetworks) > 0 {
			checks.append(t, "TestCheckResourceAttr", "synthetic_test.lan_networks.#", fmt.Sprintf("%d", len(s.SyntheticTest.LanNetworks)))
			for i, lanNetwork := range s.SyntheticTest.LanNetworks {
				if len(lanNetwork.Networks) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("synthetic_test.lan_networks.%d.networks.#", i), fmt.Sprintf("%d", len(lanNetwork.Networks)))
					for j, network := range lanNetwork.Networks {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("synthetic_test.lan_networks.%d.networks.%d", i, j), network)
					}
				}
				if len(lanNetwork.Probes) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("synthetic_test.lan_networks.%d.probes.#", i), fmt.Sprintf("%d", len(lanNetwork.Probes)))
					for j, probe := range lanNetwork.Probes {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("synthetic_test.lan_networks.%d.probes.%d", i, j), probe)
					}
				}
			}
		}
		if len(s.SyntheticTest.Vlans) > 0 {
			for i, vlan := range s.SyntheticTest.Vlans {
				if len(vlan.VlanIds) > 0 {
					for j, vlanId := range vlan.VlanIds {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("synthetic_test.vlans.%d.vlan_ids.%d", i, j), vlanId)
					}
				}
				if vlan.Disabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("synthetic_test.vlans.%d.disabled", i), fmt.Sprintf("%t", *vlan.Disabled))
				}
				if len(vlan.Probes) > 0 {
					for j, probe := range vlan.Probes {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("synthetic_test.vlans.%d.probes.%d", i, j), probe)
					}
				}
				if len(vlan.CustomTestUrls) > 0 {
					for j, url := range vlan.CustomTestUrls {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("synthetic_test.vlans.%d.custom_test_urls.%d", i, j), url)
					}
				}
			}
		}
		if s.SyntheticTest.WanSpeedtest != nil {
			if s.SyntheticTest.WanSpeedtest.Enabled != nil {
				checks.append(t, "TestCheckResourceAttr", "synthetic_test.wan_speedtest.enabled", fmt.Sprintf("%t", *s.SyntheticTest.WanSpeedtest.Enabled))
			}
			if s.SyntheticTest.WanSpeedtest.TimeOfDay != nil {
				checks.append(t, "TestCheckResourceAttr", "synthetic_test.wan_speedtest.time_of_day", *s.SyntheticTest.WanSpeedtest.TimeOfDay)
			}
		}
	}
	if s.TrackAnonymousDevices != nil {
		checks.append(t, "TestCheckResourceAttr", "track_anonymous_devices", fmt.Sprintf("%t", *s.TrackAnonymousDevices))
	}
	if s.UplinkPortConfig != nil {
		if s.UplinkPortConfig.Dot1x != nil {
			checks.append(t, "TestCheckResourceAttr", "uplink_port_config.dot1x", fmt.Sprintf("%t", *s.UplinkPortConfig.Dot1x))
		}
		if s.UplinkPortConfig.KeepWlansUpIfDown != nil {
			checks.append(t, "TestCheckResourceAttr", "uplink_port_config.keep_wlans_up_if_down", fmt.Sprintf("%t", *s.UplinkPortConfig.KeepWlansUpIfDown))
		}
	}
	if len(s.Vars) > 0 {
		for key, value := range s.Vars {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vars.%s", key), value)
		}
	}
	if s.Vna != nil && s.Vna.Enabled != nil {
		checks.append(t, "TestCheckResourceAttr", "vna.enabled", fmt.Sprintf("%t", *s.Vna.Enabled))
	}
	if len(s.VsInstance) > 0 {
		for key, vsInstance := range s.VsInstance {
			if len(vsInstance.Networks) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vs_instance.%s.networks.#", key), fmt.Sprintf("%d", len(vsInstance.Networks)))
				for i, network := range vsInstance.Networks {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("vs_instance.%s.networks.%d", key, i), network)
				}
			} else {
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("vs_instance.%s", key))
			}
		}
	}
	if s.WanVna != nil && s.WanVna.Enabled != nil {
		checks.append(t, "TestCheckResourceAttr", "wan_vna.enabled", fmt.Sprintf("%t", *s.WanVna.Enabled))
	}
	if s.Wids != nil && s.Wids.RepeatedAuthFailures != nil {
		if s.Wids.RepeatedAuthFailures.Duration != nil {
			checks.append(t, "TestCheckResourceAttr", "wids.repeated_auth_failures.duration", fmt.Sprintf("%d", *s.Wids.RepeatedAuthFailures.Duration))
		}
		if s.Wids.RepeatedAuthFailures.Threshold != nil {
			checks.append(t, "TestCheckResourceAttr", "wids.repeated_auth_failures.threshold", fmt.Sprintf("%d", *s.Wids.RepeatedAuthFailures.Threshold))
		}
	}
	if s.Wifi != nil {
		if s.Wifi.CiscoEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "wifi.cisco_enabled", fmt.Sprintf("%t", *s.Wifi.CiscoEnabled))
		}
		if s.Wifi.Disable11k != nil {
			checks.append(t, "TestCheckResourceAttr", "wifi.disable_11k", fmt.Sprintf("%t", *s.Wifi.Disable11k))
		}
		if s.Wifi.DisableRadiosWhenPowerConstrained != nil {
			checks.append(t, "TestCheckResourceAttr", "wifi.disable_radios_when_power_constrained", fmt.Sprintf("%t", *s.Wifi.DisableRadiosWhenPowerConstrained))
		}
		if s.Wifi.EnableArpSpoofCheck != nil {
			checks.append(t, "TestCheckResourceAttr", "wifi.enable_arp_spoof_check", fmt.Sprintf("%t", *s.Wifi.EnableArpSpoofCheck))
		}
		if s.Wifi.EnableSharedRadioScanning != nil {
			checks.append(t, "TestCheckResourceAttr", "wifi.enable_shared_radio_scanning", fmt.Sprintf("%t", *s.Wifi.EnableSharedRadioScanning))
		}
		if s.Wifi.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "wifi.enabled", fmt.Sprintf("%t", *s.Wifi.Enabled))
		}
		if s.Wifi.LocateConnected != nil {
			checks.append(t, "TestCheckResourceAttr", "wifi.locate_connected", fmt.Sprintf("%t", *s.Wifi.LocateConnected))
		}
		if s.Wifi.LocateUnconnected != nil {
			checks.append(t, "TestCheckResourceAttr", "wifi.locate_unconnected", fmt.Sprintf("%t", *s.Wifi.LocateUnconnected))
		}
		if s.Wifi.MeshAllowDfs != nil {
			checks.append(t, "TestCheckResourceAttr", "wifi.mesh_allow_dfs", fmt.Sprintf("%t", *s.Wifi.MeshAllowDfs))
		}
		if s.Wifi.MeshEnableCrm != nil {
			checks.append(t, "TestCheckResourceAttr", "wifi.mesh_enable_crm", fmt.Sprintf("%t", *s.Wifi.MeshEnableCrm))
		}
		if s.Wifi.MeshEnabled != nil {
			checks.append(t, "TestCheckResourceAttr", "wifi.mesh_enabled", fmt.Sprintf("%t", *s.Wifi.MeshEnabled))
		}
		if s.Wifi.MeshPsk != nil {
			checks.append(t, "TestCheckResourceAttr", "wifi.mesh_psk", *s.Wifi.MeshPsk)
		}
		if s.Wifi.MeshSsid != nil {
			checks.append(t, "TestCheckResourceAttr", "wifi.mesh_ssid", *s.Wifi.MeshSsid)
		}
		if s.Wifi.ProxyArp != nil {
			checks.append(t, "TestCheckResourceAttr", "wifi.proxy_arp", *s.Wifi.ProxyArp)
		}
	}
	if s.WiredVna != nil && s.WiredVna.Enabled != nil {
		checks.append(t, "TestCheckResourceAttr", "wired_vna.enabled", fmt.Sprintf("%t", *s.WiredVna.Enabled))
	}
	if s.ZoneOccupancyAlert != nil {
		if len(s.ZoneOccupancyAlert.EmailNotifiers) > 0 {
			for i, email := range s.ZoneOccupancyAlert.EmailNotifiers {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("zone_occupancy_alert.email_notifiers.%d", i), email)
			}
		}
		if s.ZoneOccupancyAlert.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "zone_occupancy_alert.enabled", fmt.Sprintf("%t", *s.ZoneOccupancyAlert.Enabled))
		}
		if s.ZoneOccupancyAlert.Threshold != nil {
			checks.append(t, "TestCheckResourceAttr", "zone_occupancy_alert.threshold", fmt.Sprintf("%d", *s.ZoneOccupancyAlert.Threshold))
		}
	}

	return checks
}
