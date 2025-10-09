package resource_org_setting

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func TerraformToSdk(ctx context.Context, plan *OrgSettingModel) (*models.OrgSetting, diag.Diagnostics) {
	data := models.OrgSetting{}
	var diags diag.Diagnostics
	unset := make(map[string]interface{})

	if plan.ApUpdownThreshold.ValueInt64Pointer() != nil {
		data.ApUpdownThreshold = models.NewOptional(models.ToPointer(int(plan.ApUpdownThreshold.ValueInt64())))
	} else {
		unset["-ap_updown_threshold"] = ""
	}

	if !plan.ApiPolicy.IsNull() && !plan.ApiPolicy.IsUnknown() {
		data.ApiPolicy = apiPolicyTerraformToSdk(plan.ApiPolicy)
	} else {
		unset["-api_policy"] = ""
	}

	if !plan.Cacerts.IsNull() && !plan.Cacerts.IsUnknown() {
		data.Cacerts = mistutils.ListOfStringTerraformToSdk(plan.Cacerts)
	} else {
		unset["-cacerts"] = ""
	}

	if !plan.Celona.IsNull() && !plan.Celona.IsUnknown() {
		data.Celona = celonaTerraformToSdk(plan.Celona)
	} else {
		unset["-celona"] = ""
	}

	if !plan.Cloudshark.IsNull() && !plan.Cloudshark.IsUnknown() {
		data.Cloudshark = cloudsharkTerraformToSdk(plan.Cloudshark)
	} else {
		unset["-cloudshark"] = ""
	}

	if !plan.DeviceCert.IsNull() && !plan.DeviceCert.IsUnknown() {
		data.DeviceCert = deviceCertTerraformToSdk(plan.DeviceCert)
	} else {
		unset["-device_cert"] = ""
	}

	if plan.DeviceUpdownThreshold.ValueInt64Pointer() != nil {
		data.DeviceUpdownThreshold = models.NewOptional(models.ToPointer(int(plan.DeviceUpdownThreshold.ValueInt64())))
	} else {
		unset["-device_updown_threshold"] = ""
	}

	if plan.DisablePcap.ValueBoolPointer() != nil {
		data.DisablePcap = plan.DisablePcap.ValueBoolPointer()
	} else {
		unset["-disable_pcap"] = ""
	}

	if plan.DisableRemoteShell.ValueBoolPointer() != nil {
		data.DisableRemoteShell = plan.DisableRemoteShell.ValueBoolPointer()
	} else {
		unset["-disable_remote_shell"] = ""
	}

	if plan.GatewayUpdownThreshold.ValueInt64Pointer() != nil {
		data.GatewayUpdownThreshold = models.NewOptional(models.ToPointer(int(plan.GatewayUpdownThreshold.ValueInt64())))
	} else {
		unset["-gateway_updown_threshold"] = ""
	}

	if !plan.Installer.IsNull() && !plan.Installer.IsUnknown() {
		data.Installer = installerTerraformToSdk(plan.Installer)
	} else {
		unset["-installer"] = ""
	}

	if !plan.JcloudRa.IsNull() && !plan.JcloudRa.IsUnknown() {
		data.JcloudRa = jcloudRaTerraformToSdk(plan.JcloudRa)
	} else {
		unset["-jcloud_ra"] = ""
	}

	if !plan.Jcloud.IsNull() && !plan.Jcloud.IsUnknown() {
		data.Jcloud = jcloudTerraformToSdk(plan.Jcloud)
	} else {
		unset["-jcloud"] = ""
	}

	if !plan.Juniper.IsNull() && !plan.Juniper.IsUnknown() {
		data.Juniper = juniperTerraformToSdk(plan.Juniper)
	} else {
		unset["-juniper"] = ""
	}

	if !plan.JuniperSrx.IsNull() && !plan.JuniperSrx.IsUnknown() {
		data.JuniperSrx = juniperSrxTerraformToSdk(ctx, &diags, plan.JuniperSrx)
	} else {
		unset["-juniper_srx"] = ""
	}

	if !plan.JunosShellAccess.IsNull() && !plan.JunosShellAccess.IsUnknown() {
		data.JunosShellAccess = junosShellAccessTerraformToSdk(plan.JunosShellAccess)
	} else {
		unset["-junos_shell_access"] = ""
	}

	if !plan.Marvis.IsNull() && !plan.Marvis.IsUnknown() {
		data.Marvis = marvisTerraformToSdk(ctx, plan.Marvis)
	} else {
		unset["-marvis"] = ""
	}

	if !plan.Mgmt.IsNull() && !plan.Mgmt.IsUnknown() {
		data.Mgmt = mgmtTerraformToSdk(plan.Mgmt)
	} else {
		unset["-mgmt"] = ""
	}

	if !plan.MistNac.IsNull() && !plan.MistNac.IsUnknown() {
		data.MistNac = mistNacTerraformToSdk(ctx, &diags, plan.MistNac)
	} else {
		unset["-mist_nac"] = ""
	}

	if !plan.MxedgeMgmt.IsNull() && !plan.MxedgeMgmt.IsUnknown() {
		data.MxedgeMgmt = mxEdgeMgmtTerraformToSdk(plan.MxedgeMgmt)
	} else {
		unset["-mxedge_mgmt"] = ""
	}

	if !plan.OpticPortConfig.IsNull() && !plan.OpticPortConfig.IsUnknown() {
		data.OpticPortConfig = opticPortConfigTerraformToSdk(plan.OpticPortConfig)
	} else {
		unset["-optic_port_config"] = ""
	}

	if !plan.PasswordPolicy.IsNull() && !plan.PasswordPolicy.IsUnknown() {
		data.PasswordPolicy = passwordPolicyTerraformToSdk(plan.PasswordPolicy)
	} else {
		unset["-password_policy"] = ""
	}

	if !plan.Pcap.IsNull() && !plan.Pcap.IsUnknown() {
		data.Pcap = pcapTerraformToSdk(plan.Pcap)
	} else {
		unset["-pcap"] = ""
	}

	if !plan.Security.IsNull() && !plan.Security.IsUnknown() {
		data.Security = securityTerraformToSdk(plan.Security)
	} else {
		unset["-security"] = ""
	}

	if !plan.Ssr.IsNull() && !plan.Ssr.IsUnknown() {
		data.Ssr = ssrTerraformToSdk(plan.Ssr)
	} else {
		unset["-ssr"] = ""
	}

	if !plan.Switch.IsNull() && !plan.Switch.IsUnknown() {
		data.Switch = switchTerraformToSdk(ctx, plan.Switch)
	} else {
		unset["-switch"] = ""
	}

	if !plan.SwitchMgmt.IsNull() && !plan.SwitchMgmt.IsUnknown() {
		data.SwitchMgmt = switchMgmtTerraformToSdk(plan.SwitchMgmt)
	} else {
		unset["-switch_mgmt"] = ""
	}

	if plan.SwitchUpdownThreshold.ValueInt64Pointer() != nil {
		data.SwitchUpdownThreshold = models.NewOptional(models.ToPointer(int(plan.SwitchUpdownThreshold.ValueInt64())))
	} else {
		unset["-switch_updown_threshold"] = ""
	}

	if !plan.SyntheticTest.IsNull() && !plan.SyntheticTest.IsUnknown() {
		data.SyntheticTest = syntheticTestTerraformToSdk(ctx, &diags, plan.SyntheticTest)
	} else {
		unset["-synthetic_test"] = ""
	}

	if plan.UiIdleTimeout.ValueInt64Pointer() != nil {
		data.UiIdleTimeout = models.ToPointer(int(plan.UiIdleTimeout.ValueInt64()))
	} else {
		unset["-ui_idle_timeout"] = ""
	}

	if plan.UiNoTracking.ValueBoolPointer() != nil {
		data.UiNoTracking = plan.UiNoTracking.ValueBoolPointer()
	} else {
		unset["-ui_no_tracking"] = ""
	}

	if !plan.VpnOptions.IsNull() && !plan.VpnOptions.IsUnknown() {
		data.VpnOptions = vpnOptionsTerraformToSdk(plan.VpnOptions)
	} else {
		unset["-vpn_options"] = ""
	}

	if !plan.WanPma.IsNull() && !plan.WanPma.IsUnknown() {
		if plan.WanPma.Enabled.ValueBoolPointer() != nil {
			data.WanPma.Enabled = plan.WanPma.Enabled.ValueBoolPointer()
		}
	} else {
		unset["-wan_pma"] = ""
	}

	if !plan.WiredPma.IsNull() && !plan.WiredPma.IsUnknown() {
		if plan.WiredPma.Enabled.ValueBoolPointer() != nil {
			data.WiredPma.Enabled = plan.WiredPma.Enabled.ValueBoolPointer()
		}
	} else {
		unset["-wired_pma"] = ""
	}

	if !plan.WirelessPma.IsNull() && !plan.WirelessPma.IsUnknown() {
		if plan.WirelessPma.Enabled.ValueBoolPointer() != nil {
			data.WirelessPma.Enabled = plan.WirelessPma.Enabled.ValueBoolPointer()
		}
	} else {
		unset["-wireless_pma"] = ""
	}

	data.AdditionalProperties = unset

	return &data, diags
}
