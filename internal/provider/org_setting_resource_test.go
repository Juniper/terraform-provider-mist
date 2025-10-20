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

func TestOrgSettingModel(t *testing.T) {
	type testStep struct {
		config OrgSettingModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgSettingModel{
						OrgId: GetTestOrgId(),
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/org_setting_resource/org_setting_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		var FixtureOrgSettingModel OrgSettingModel
		err = hcl.Decode(&FixtureOrgSettingModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		FixtureOrgSettingModel.OrgId = GetTestOrgId()

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureOrgSettingModel,
				},
			},
		}
	}

	resourceType := "org_setting"
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {

			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {

				config := step.config

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				configStr := Render(resourceType, tName, string(f.Bytes()))

				checks := config.testChecks(t, PrefixProviderName(resourceType), tName)
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
}

func (o *OrgSettingModel) testChecks(t testing.TB, rType, rName string) testChecks {
	checks := newTestChecks(rType + "." + rName)

	// Check required fields
	checks.append(t, "TestCheckResourceAttr", "org_id", o.OrgId)

	// Check optional integer fields
	if o.ApUpdownThreshold != nil {
		checks.append(t, "TestCheckResourceAttr", "ap_updown_threshold", fmt.Sprintf("%d", *o.ApUpdownThreshold))
	}
	if o.DeviceUpdownThreshold != nil {
		checks.append(t, "TestCheckResourceAttr", "device_updown_threshold", fmt.Sprintf("%d", *o.DeviceUpdownThreshold))
	}
	if o.GatewayUpdownThreshold != nil {
		checks.append(t, "TestCheckResourceAttr", "gateway_updown_threshold", fmt.Sprintf("%d", *o.GatewayUpdownThreshold))
	}
	if o.SwitchUpdownThreshold != nil {
		checks.append(t, "TestCheckResourceAttr", "switch_updown_threshold", fmt.Sprintf("%d", *o.SwitchUpdownThreshold))
	}
	if o.UiIdleTimeout != nil {
		checks.append(t, "TestCheckResourceAttr", "ui_idle_timeout", fmt.Sprintf("%d", *o.UiIdleTimeout))
	}

	// Check boolean fields
	if o.DisablePcap != nil {
		checks.append(t, "TestCheckResourceAttr", "disable_pcap", fmt.Sprintf("%t", *o.DisablePcap))
	}
	if o.DisableRemoteShell != nil {
		checks.append(t, "TestCheckResourceAttr", "disable_remote_shell", fmt.Sprintf("%t", *o.DisableRemoteShell))
	}
	if o.UiNoTracking != nil {
		checks.append(t, "TestCheckResourceAttr", "ui_no_tracking", fmt.Sprintf("%t", *o.UiNoTracking))
	}

	// Check array fields
	if len(o.Cacerts) > 0 {
		checks.append(t, "TestCheckResourceAttr", "cacerts.#", fmt.Sprintf("%d", len(o.Cacerts)))
		for i, cert := range o.Cacerts {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("cacerts.%d", i), cert)
		}
	}

	// Check API Policy nested object
	if o.ApiPolicy != nil {
		if o.ApiPolicy.NoReveal != nil {
			checks.append(t, "TestCheckResourceAttr", "api_policy.no_reveal", fmt.Sprintf("%t", *o.ApiPolicy.NoReveal))
		}
	}

	// Check Celona nested object
	if o.Celona != nil {
		checks.append(t, "TestCheckResourceAttr", "celona.api_key", o.Celona.ApiKey)
		checks.append(t, "TestCheckResourceAttr", "celona.api_prefix", o.Celona.ApiPrefix)
	}

	// Check Cloudshark nested object
	if o.Cloudshark != nil {
		if o.Cloudshark.Apitoken != nil {
			checks.append(t, "TestCheckResourceAttr", "cloudshark.apitoken", *o.Cloudshark.Apitoken)
		}
		if o.Cloudshark.Url != nil {
			checks.append(t, "TestCheckResourceAttr", "cloudshark.url", *o.Cloudshark.Url)
		}
	}

	// Check Device Cert nested object
	if o.DeviceCert != nil {
		checks.append(t, "TestCheckResourceAttr", "device_cert.cert", o.DeviceCert.Cert)
		checks.append(t, "TestCheckResourceAttr", "device_cert.key", o.DeviceCert.Key)
	}

	// Check Installer nested object
	if o.Installer != nil {
		if o.Installer.AllowAllDevices != nil {
			checks.append(t, "TestCheckResourceAttr", "installer.allow_all_devices", fmt.Sprintf("%t", *o.Installer.AllowAllDevices))
		}
		if o.Installer.AllowAllSites != nil {
			checks.append(t, "TestCheckResourceAttr", "installer.allow_all_sites", fmt.Sprintf("%t", *o.Installer.AllowAllSites))
		}
		if len(o.Installer.ExtraSiteIds) > 0 {
			checks.append(t, "TestCheckResourceAttr", "installer.extra_site_ids.#", fmt.Sprintf("%d", len(o.Installer.ExtraSiteIds)))
			for i, siteId := range o.Installer.ExtraSiteIds {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("installer.extra_site_ids.%d", i), siteId)
			}
		}
		if o.Installer.GracePeriod != nil {
			checks.append(t, "TestCheckResourceAttr", "installer.grace_period", fmt.Sprintf("%d", *o.Installer.GracePeriod))
		}
	}

	// Check Jcloud nested object
	if o.Jcloud != nil {
		checks.append(t, "TestCheckResourceAttr", "jcloud.org_apitoken", o.Jcloud.OrgApitoken)
		checks.append(t, "TestCheckResourceAttr", "jcloud.org_apitoken_name", o.Jcloud.OrgApitokenName)
		checks.append(t, "TestCheckResourceAttr", "jcloud.org_id", o.Jcloud.OrgId)
	}

	// Check JcloudRa nested object
	if o.JcloudRa != nil {
		if o.JcloudRa.OrgApitoken != nil {
			checks.append(t, "TestCheckResourceAttr", "jcloud_ra.org_apitoken", *o.JcloudRa.OrgApitoken)
		}
		if o.JcloudRa.OrgApitokenName != nil {
			checks.append(t, "TestCheckResourceAttr", "jcloud_ra.org_apitoken_name", *o.JcloudRa.OrgApitokenName)
		}
		if o.JcloudRa.OrgId != nil {
			checks.append(t, "TestCheckResourceAttr", "jcloud_ra.org_id", *o.JcloudRa.OrgId)
		}
	}

	// Check Management nested object
	if o.Mgmt != nil {
		if len(o.Mgmt.MxtunnelIds) > 0 {
			checks.append(t, "TestCheckResourceAttr", "mgmt.mxtunnel_ids.#", fmt.Sprintf("%d", len(o.Mgmt.MxtunnelIds)))
			for i, tunnelId := range o.Mgmt.MxtunnelIds {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("mgmt.mxtunnel_ids.%d", i), tunnelId)
			}
		}
		if o.Mgmt.UseMxtunnel != nil {
			checks.append(t, "TestCheckResourceAttr", "mgmt.use_mxtunnel", fmt.Sprintf("%t", *o.Mgmt.UseMxtunnel))
		}
		if o.Mgmt.UseWxtunnel != nil {
			checks.append(t, "TestCheckResourceAttr", "mgmt.use_wxtunnel", fmt.Sprintf("%t", *o.Mgmt.UseWxtunnel))
		}
	}

	// Check MistNac nested object - this is a complex nested structure
	if o.MistNac != nil {
		if len(o.MistNac.Cacerts) > 0 {
			checks.append(t, "TestCheckResourceAttr", "mist_nac.cacerts.#", fmt.Sprintf("%d", len(o.MistNac.Cacerts)))
			for i, cert := range o.MistNac.Cacerts {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("mist_nac.cacerts.%d", i), cert)
			}
		}
		if o.MistNac.DefaultIdpId != nil {
			checks.append(t, "TestCheckResourceAttr", "mist_nac.default_idp_id", *o.MistNac.DefaultIdpId)
		}
		if o.MistNac.DisableRsaeAlgorithms != nil {
			checks.append(t, "TestCheckResourceAttr", "mist_nac.disable_rsae_algorithms", fmt.Sprintf("%t", *o.MistNac.DisableRsaeAlgorithms))
		}
		if o.MistNac.EapSslSecurityLevel != nil {
			checks.append(t, "TestCheckResourceAttr", "mist_nac.eap_ssl_security_level", fmt.Sprintf("%d", *o.MistNac.EapSslSecurityLevel))
		}
		if o.MistNac.EuOnly != nil {
			checks.append(t, "TestCheckResourceAttr", "mist_nac.eu_only", fmt.Sprintf("%t", *o.MistNac.EuOnly))
		}
		if o.MistNac.IdpMachineCertLookupField != nil {
			checks.append(t, "TestCheckResourceAttr", "mist_nac.idp_machine_cert_lookup_field", *o.MistNac.IdpMachineCertLookupField)
		}
		if o.MistNac.IdpUserCertLookupField != nil {
			checks.append(t, "TestCheckResourceAttr", "mist_nac.idp_user_cert_lookup_field", *o.MistNac.IdpUserCertLookupField)
		}
		if len(o.MistNac.Idps) > 0 {
			checks.append(t, "TestCheckResourceAttr", "mist_nac.idps.#", fmt.Sprintf("%d", len(o.MistNac.Idps)))
			for i, idp := range o.MistNac.Idps {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("mist_nac.idps.%d.id", i), idp.Id)
				if len(idp.ExcludeRealms) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("mist_nac.idps.%d.exclude_realms.#", i), fmt.Sprintf("%d", len(idp.ExcludeRealms)))
					for j, realm := range idp.ExcludeRealms {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("mist_nac.idps.%d.exclude_realms.%d", i, j), realm)
					}
				}
				if len(idp.UserRealms) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("mist_nac.idps.%d.user_realms.#", i), fmt.Sprintf("%d", len(idp.UserRealms)))
					for j, realm := range idp.UserRealms {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("mist_nac.idps.%d.user_realms.%d", i, j), realm)
					}
				}
			}
		}
		if o.MistNac.ServerCert != nil {
			if o.MistNac.ServerCert.Cert != nil {
				checks.append(t, "TestCheckResourceAttr", "mist_nac.server_cert.cert", *o.MistNac.ServerCert.Cert)
			}
			if o.MistNac.ServerCert.Key != nil {
				checks.append(t, "TestCheckResourceAttr", "mist_nac.server_cert.key", *o.MistNac.ServerCert.Key)
			}
			if o.MistNac.ServerCert.Password != nil {
				checks.append(t, "TestCheckResourceAttr", "mist_nac.server_cert.password", *o.MistNac.ServerCert.Password)
			}
		}
		if o.MistNac.UseIpVersion != nil {
			checks.append(t, "TestCheckResourceAttr", "mist_nac.use_ip_version", *o.MistNac.UseIpVersion)
		}
		if o.MistNac.UseSslPort != nil {
			checks.append(t, "TestCheckResourceAttr", "mist_nac.use_ssl_port", fmt.Sprintf("%t", *o.MistNac.UseSslPort))
		}
	}

	// Check MxedgeMgmt nested object
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
		if o.MxedgeMgmt.OobIpType != nil {
			checks.append(t, "TestCheckResourceAttr", "mxedge_mgmt.oob_ip_type", *o.MxedgeMgmt.OobIpType)
		}
		if o.MxedgeMgmt.OobIpType6 != nil {
			checks.append(t, "TestCheckResourceAttr", "mxedge_mgmt.oob_ip_type6", *o.MxedgeMgmt.OobIpType6)
		}
		if o.MxedgeMgmt.RootPassword != nil {
			checks.append(t, "TestCheckResourceAttr", "mxedge_mgmt.root_password", *o.MxedgeMgmt.RootPassword)
		}
	}

	// Check Marvis nested object
	if o.Marvis != nil {
		if o.Marvis.AutoOperations != nil {
			if o.Marvis.AutoOperations.BouncePortForAbnormalPoeClient != nil {
				checks.append(t, "TestCheckResourceAttr", "marvis.auto_operations.bounce_port_for_abnormal_poe_client", fmt.Sprintf("%t", *o.Marvis.AutoOperations.BouncePortForAbnormalPoeClient))
			}
			if o.Marvis.AutoOperations.DisablePortWhenDdosProtocolViolation != nil {
				checks.append(t, "TestCheckResourceAttr", "marvis.auto_operations.disable_port_when_ddos_protocol_violation", fmt.Sprintf("%t", *o.Marvis.AutoOperations.DisablePortWhenDdosProtocolViolation))
			}
			if o.Marvis.AutoOperations.DisablePortWhenRogueDhcpServerDetected != nil {
				checks.append(t, "TestCheckResourceAttr", "marvis.auto_operations.disable_port_when_rogue_dhcp_server_detected", fmt.Sprintf("%t", *o.Marvis.AutoOperations.DisablePortWhenRogueDhcpServerDetected))
			}
		}
	}

	// Check JuniperSrx nested object
	if o.JuniperSrx != nil {
		if o.JuniperSrx.SrxAutoUpgrade != nil {
			if o.JuniperSrx.SrxAutoUpgrade.Enabled != nil {
				checks.append(t, "TestCheckResourceAttr", "juniper_srx.auto_upgrade.enabled", fmt.Sprintf("%t", *o.JuniperSrx.SrxAutoUpgrade.Enabled))
			}
			if o.JuniperSrx.SrxAutoUpgrade.Snapshot != nil {
				checks.append(t, "TestCheckResourceAttr", "juniper_srx.auto_upgrade.snapshot", fmt.Sprintf("%t", *o.JuniperSrx.SrxAutoUpgrade.Snapshot))
			}
			if len(o.JuniperSrx.SrxAutoUpgrade.CustomVersions) > 0 {
				checks.append(t, "TestCheckResourceAttr", "juniper_srx.auto_upgrade.custom_versions.%", fmt.Sprintf("%d", len(o.JuniperSrx.SrxAutoUpgrade.CustomVersions)))
				for key, version := range o.JuniperSrx.SrxAutoUpgrade.CustomVersions {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("juniper_srx.auto_upgrade.custom_versions.%s", key), version)
				}
			}
		}
	}

	// Check JunosShellAccess nested object
	if o.JunosShellAccess != nil {
		if o.JunosShellAccess.Admin != nil {
			checks.append(t, "TestCheckResourceAttr", "junos_shell_access.admin", *o.JunosShellAccess.Admin)
		}
		if o.JunosShellAccess.Helpdesk != nil {
			checks.append(t, "TestCheckResourceAttr", "junos_shell_access.helpdesk", *o.JunosShellAccess.Helpdesk)
		}
		if o.JunosShellAccess.Read != nil {
			checks.append(t, "TestCheckResourceAttr", "junos_shell_access.read", *o.JunosShellAccess.Read)
		}
		if o.JunosShellAccess.Write != nil {
			checks.append(t, "TestCheckResourceAttr", "junos_shell_access.write", *o.JunosShellAccess.Write)
		}
	}

	// Check Ssr nested object
	if o.Ssr != nil {
		if len(o.Ssr.ConductorHosts) > 0 {
			checks.append(t, "TestCheckResourceAttr", "ssr.conductor_hosts.#", fmt.Sprintf("%d", len(o.Ssr.ConductorHosts)))
			for i, host := range o.Ssr.ConductorHosts {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ssr.conductor_hosts.%d", i), host)
			}
		}
		if o.Ssr.ConductorToken != nil {
			checks.append(t, "TestCheckResourceAttr", "ssr.conductor_token", *o.Ssr.ConductorToken)
		}
		if o.Ssr.DisableStats != nil {
			checks.append(t, "TestCheckResourceAttr", "ssr.disable_stats", fmt.Sprintf("%t", *o.Ssr.DisableStats))
		}
		if o.Ssr.Proxy != nil {
			if o.Ssr.Proxy.Url != nil {
				checks.append(t, "TestCheckResourceAttr", "ssr.proxy.url", *o.Ssr.Proxy.Url)
			}
		}
		if o.Ssr.SsrAutoUpgrade != nil {
			if o.Ssr.SsrAutoUpgrade.Enabled != nil {
				checks.append(t, "TestCheckResourceAttr", "ssr.auto_upgrade.enabled", fmt.Sprintf("%t", *o.Ssr.SsrAutoUpgrade.Enabled))
			}
			if o.Ssr.SsrAutoUpgrade.Channel != nil {
				checks.append(t, "TestCheckResourceAttr", "ssr.auto_upgrade.channel", *o.Ssr.SsrAutoUpgrade.Channel)
			}
			if len(o.Ssr.SsrAutoUpgrade.CustomVersions) > 0 {
				checks.append(t, "TestCheckResourceAttr", "ssr.auto_upgrade.custom_versions.%", fmt.Sprintf("%d", len(o.Ssr.SsrAutoUpgrade.CustomVersions)))
				for key, version := range o.Ssr.SsrAutoUpgrade.CustomVersions {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("ssr.auto_upgrade.custom_versions.%s", key), version)
				}
			}
		}
	}

	// Check Switch nested object
	if o.Switch != nil {
		if o.Switch.AutoUpgrade != nil {
			if o.Switch.AutoUpgrade.Enabled != nil {
				checks.append(t, "TestCheckResourceAttr", "switch.auto_upgrade.enabled", fmt.Sprintf("%t", *o.Switch.AutoUpgrade.Enabled))
			}
			if o.Switch.AutoUpgrade.Snapshot != nil {
				checks.append(t, "TestCheckResourceAttr", "switch.auto_upgrade.snapshot", fmt.Sprintf("%t", *o.Switch.AutoUpgrade.Snapshot))
			}
			if len(o.Switch.AutoUpgrade.CustomVersions) > 0 {
				checks.append(t, "TestCheckResourceAttr", "switch.auto_upgrade.custom_versions.%", fmt.Sprintf("%d", len(o.Switch.AutoUpgrade.CustomVersions)))
				for key, version := range o.Switch.AutoUpgrade.CustomVersions {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("switch.auto_upgrade.custom_versions.%s", key), version)
				}
			}
		}
	}

	// Check Password Policy nested object
	if o.PasswordPolicy != nil {
		if o.PasswordPolicy.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "password_policy.enabled", fmt.Sprintf("%t", *o.PasswordPolicy.Enabled))
		}
		if o.PasswordPolicy.ExpiryInDays != nil {
			checks.append(t, "TestCheckResourceAttr", "password_policy.expiry_in_days", fmt.Sprintf("%d", *o.PasswordPolicy.ExpiryInDays))
		}
		if o.PasswordPolicy.MinLength != nil {
			checks.append(t, "TestCheckResourceAttr", "password_policy.min_length", fmt.Sprintf("%d", *o.PasswordPolicy.MinLength))
		}
		if o.PasswordPolicy.RequiresSpecialChar != nil {
			checks.append(t, "TestCheckResourceAttr", "password_policy.requires_special_char", fmt.Sprintf("%t", *o.PasswordPolicy.RequiresSpecialChar))
		}
		if o.PasswordPolicy.RequiresTwoFactorAuth != nil {
			checks.append(t, "TestCheckResourceAttr", "password_policy.requires_two_factor_auth", fmt.Sprintf("%t", *o.PasswordPolicy.RequiresTwoFactorAuth))
		}
	}

	// Check Pcap nested object
	if o.Pcap != nil {
		if o.Pcap.Bucket != nil {
			checks.append(t, "TestCheckResourceAttr", "pcap.bucket", *o.Pcap.Bucket)
		}
		if o.Pcap.MaxPktLen != nil {
			checks.append(t, "TestCheckResourceAttr", "pcap.max_pkt_len", fmt.Sprintf("%d", *o.Pcap.MaxPktLen))
		}
	}

	// Check Security nested object
	if o.Security != nil {
		if o.Security.DisableLocalSsh != nil {
			checks.append(t, "TestCheckResourceAttr", "security.disable_local_ssh", fmt.Sprintf("%t", *o.Security.DisableLocalSsh))
		}
		if o.Security.FipsZeroizePassword != nil {
			checks.append(t, "TestCheckResourceAttr", "security.fips_zeroize_password", *o.Security.FipsZeroizePassword)
		}
		if o.Security.LimitSshAccess != nil {
			checks.append(t, "TestCheckResourceAttr", "security.limit_ssh_access", fmt.Sprintf("%t", *o.Security.LimitSshAccess))
		}
	}

	// Check SwitchMgmt nested object
	if o.SwitchMgmt != nil {
		if o.SwitchMgmt.ApAffinityThreshold != nil {
			checks.append(t, "TestCheckResourceAttr", "switch_mgmt.ap_affinity_threshold", fmt.Sprintf("%d", *o.SwitchMgmt.ApAffinityThreshold))
		}
	}

	// Check SyntheticTest nested object
	if o.SyntheticTest != nil {
		if o.SyntheticTest.Aggressiveness != nil {
			checks.append(t, "TestCheckResourceAttr", "synthetic_test.aggressiveness", *o.SyntheticTest.Aggressiveness)
		}
		if o.SyntheticTest.Disabled != nil {
			checks.append(t, "TestCheckResourceAttr", "synthetic_test.disabled", fmt.Sprintf("%t", *o.SyntheticTest.Disabled))
		}

		// Check CustomProbes map
		if len(o.SyntheticTest.CustomProbes) > 0 {
			checks.append(t, "TestCheckResourceAttr", "synthetic_test.custom_probes.%", fmt.Sprintf("%d", len(o.SyntheticTest.CustomProbes)))
			for key, probe := range o.SyntheticTest.CustomProbes {
				if probe.Aggressiveness != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("synthetic_test.custom_probes.%s.aggressiveness", key), *probe.Aggressiveness)
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
				if probe.CustomProbesType != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("synthetic_test.custom_probes.%s.type", key), *probe.CustomProbesType)
				}
				if probe.Url != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("synthetic_test.custom_probes.%s.url", key), *probe.Url)
				}
			}
		}

		// Check LanNetworks array
		if len(o.SyntheticTest.LanNetworks) > 0 {
			checks.append(t, "TestCheckResourceAttr", "synthetic_test.lan_networks.#", fmt.Sprintf("%d", len(o.SyntheticTest.LanNetworks)))
			for i, lanNetwork := range o.SyntheticTest.LanNetworks {
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

		if len(o.SyntheticTest.Vlans) > 0 {
			checks.append(t, "TestCheckResourceAttr", "synthetic_test.vlans.#", fmt.Sprintf("%d", len(o.SyntheticTest.Vlans)))
			for i, vlan := range o.SyntheticTest.Vlans {
				if vlan.Disabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("synthetic_test.vlans.%d.disabled", i), fmt.Sprintf("%t", *vlan.Disabled))
				}
				if len(vlan.CustomTestUrls) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("synthetic_test.vlans.%d.custom_test_urls.#", i), fmt.Sprintf("%d", len(vlan.CustomTestUrls)))
					for j, url := range vlan.CustomTestUrls {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("synthetic_test.vlans.%d.custom_test_urls.%d", i, j), url)
					}
				}
				if len(vlan.Probes) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("synthetic_test.vlans.%d.probes.#", i), fmt.Sprintf("%d", len(vlan.Probes)))
					for j, probe := range vlan.Probes {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("synthetic_test.vlans.%d.probes.%d", i, j), probe)
					}
				}
				if len(vlan.VlanIds) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("synthetic_test.vlans.%d.vlan_ids.#", i), fmt.Sprintf("%d", len(vlan.VlanIds)))
					for j, vlanId := range vlan.VlanIds {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("synthetic_test.vlans.%d.vlan_ids.%d", i, j), vlanId)
					}
				}
			}
		}
		if o.SyntheticTest.WanSpeedtest != nil {
			if o.SyntheticTest.WanSpeedtest.Enabled != nil {
				checks.append(t, "TestCheckResourceAttr", "synthetic_test.wan_speedtest.enabled", fmt.Sprintf("%t", *o.SyntheticTest.WanSpeedtest.Enabled))
			}
			if o.SyntheticTest.WanSpeedtest.TimeOfDay != nil {
				checks.append(t, "TestCheckResourceAttr", "synthetic_test.wan_speedtest.time_of_day", *o.SyntheticTest.WanSpeedtest.TimeOfDay)
			}
		}
	}

	// Check VpnOptions nested object
	if o.VpnOptions != nil {
		if o.VpnOptions.AsBase != nil {
			checks.append(t, "TestCheckResourceAttr", "vpn_options.as_base", fmt.Sprintf("%d", *o.VpnOptions.AsBase))
		}
		if o.VpnOptions.EnableIpv6 != nil {
			checks.append(t, "TestCheckResourceAttr", "vpn_options.enable_ipv6", fmt.Sprintf("%t", *o.VpnOptions.EnableIpv6))
		}
		if o.VpnOptions.StSubnet != nil {
			checks.append(t, "TestCheckResourceAttr", "vpn_options.st_subnet", *o.VpnOptions.StSubnet)
		}
	}

	// Check PMA nested objects
	if o.WanPma != nil {
		if o.WanPma.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "wan_pma.enabled", fmt.Sprintf("%t", *o.WanPma.Enabled))
		}
	}
	if o.WiredPma != nil {
		if o.WiredPma.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "wired_pma.enabled", fmt.Sprintf("%t", *o.WiredPma.Enabled))
		}
	}
	if o.WirelessPma != nil {
		if o.WirelessPma.Enabled != nil {
			checks.append(t, "TestCheckResourceAttr", "wireless_pma.enabled", fmt.Sprintf("%t", *o.WirelessPma.Enabled))
		}
	}

	// Check map fields - OpticPortConfig
	if len(o.OpticPortConfig) > 0 {
		checks.append(t, "TestCheckResourceAttr", "optic_port_config.%", fmt.Sprintf("%d", len(o.OpticPortConfig)))
		for key, config := range o.OpticPortConfig {
			if config.Channelized != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("optic_port_config.%s.channelized", key), fmt.Sprintf("%t", *config.Channelized))
			}
			if config.Speed != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("optic_port_config.%s.speed", key), *config.Speed)
			}
		}
	}

	return checks
}
