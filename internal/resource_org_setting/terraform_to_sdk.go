package resource_org_setting

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
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
		data.ApiPolicy = apiPolicyTerraformToSdk(ctx, &diags, plan.ApiPolicy)
	} else {
		unset["-api_policy"] = ""
	}

	if !plan.Cacerts.IsNull() && !plan.Cacerts.IsUnknown() {
		data.Cacerts = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Cacerts)
	} else {
		unset["-cacerts"] = ""
	}

	if !plan.Celona.IsNull() && !plan.Celona.IsUnknown() {
		data.Celona = celonaTerraformToSdk(ctx, &diags, plan.Celona)
	} else {
		unset["-celona"] = ""
	}

	if !plan.Cloudshark.IsNull() && !plan.Cloudshark.IsUnknown() {
		data.Cloudshark = cloudsharkTerraformToSdk(ctx, &diags, plan.Cloudshark)
	} else {
		unset["-cloudshark"] = ""
	}

	if !plan.Cradlepoint.IsNull() && !plan.Cradlepoint.IsUnknown() {
		data.Cradlepoint = cradlepointTerraformToSdk(ctx, &diags, plan.Cradlepoint)
	} else {
		unset["-cradlepoint"] = ""
	}

	if !plan.DeviceCert.IsNull() && !plan.DeviceCert.IsUnknown() {
		data.DeviceCert = deviceCertTerraformToSdk(ctx, &diags, plan.DeviceCert)
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
		data.Installer = installerTerraformToSdk(ctx, &diags, plan.Installer)
	} else {
		unset["-installer"] = ""
	}

	if !plan.Jcloud.IsNull() && !plan.Jcloud.IsUnknown() {
		data.Jcloud = jcloudTerraformToSdk(ctx, &diags, plan.Jcloud)
	} else {
		unset["-jcloud"] = ""
	}

	if !plan.Juniper.IsNull() && !plan.Juniper.IsUnknown() {
		data.Juniper = juniperTerraformToSdk(ctx, &diags, plan.Juniper)
	} else {
		unset["-juniper"] = ""
	}

	if !plan.Mgmt.IsNull() && !plan.Mgmt.IsUnknown() {
		data.Mgmt = mgmtTerraformToSdk(ctx, &diags, plan.Mgmt)
	} else {
		unset["-mgmt"] = ""
	}

	if !plan.MistNac.IsNull() && !plan.MistNac.IsUnknown() {
		data.MistNac = mistNacTerraformToSdk(ctx, &diags, plan.MistNac)
	} else {
		unset["-mist_nac"] = ""
	}

	if plan.MxedgeFipsEnabled.ValueBoolPointer() != nil {
		data.MxedgeFipsEnabled = plan.MxedgeFipsEnabled.ValueBoolPointer()
	} else {
		unset["-mxedge_fips_enabled"] = ""
	}

	if !plan.MxedgeMgmt.IsNull() && !plan.MxedgeMgmt.IsUnknown() {
		data.MxedgeMgmt = mxEdgeMgmtTerraformToSdk(ctx, &diags, plan.MxedgeMgmt)
	} else {
		unset["-mxedge_mgmt"] = ""
	}

	if !plan.PasswordPolicy.IsNull() && !plan.PasswordPolicy.IsUnknown() {
		data.PasswordPolicy = passwordPolicyTerraformToSdk(ctx, &diags, plan.PasswordPolicy)
	} else {
		unset["-password_policy"] = ""
	}

	if !plan.Pcap.IsNull() && !plan.Pcap.IsUnknown() {
		data.Pcap = pcapTerraformToSdk(ctx, &diags, plan.Pcap)
	} else {
		unset["-pcap"] = ""
	}

	if !plan.Security.IsNull() && !plan.Security.IsUnknown() {
		data.Security = securityTerraformToSdk(ctx, &diags, plan.Security)
	} else {
		unset["-security"] = ""
	}

	if !plan.SwitchMgmt.IsNull() && !plan.SwitchMgmt.IsUnknown() {
		data.SwitchMgmt = switchMgmtTerraformToSdk(ctx, &diags, plan.SwitchMgmt)
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

	if !plan.VpnOptions.IsNull() && !plan.VpnOptions.IsUnknown() {
		data.VpnOptions = vpnOptionsTerraformToSdk(ctx, &diags, plan.VpnOptions)
	} else {
		unset["-vpn_options"] = ""
	}

	data.AdditionalProperties = unset

	return &data, diags
}
