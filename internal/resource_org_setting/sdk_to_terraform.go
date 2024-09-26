package resource_org_setting

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, data *models.OrgSetting) (OrgSettingModel, diag.Diagnostics) {
	var state OrgSettingModel
	var diags diag.Diagnostics

	var ap_updown_threshold types.Int64
	var api_policy ApiPolicyValue = NewApiPolicyValueNull()
	// var blacklist_url types.String
	var cacerts types.List = types.ListNull(types.StringType)
	var celona CelonaValue = NewCelonaValueNull()
	var cloudshark CloudsharkValue = NewCloudsharkValueNull()
	var cradlepoint CradlepointValue = NewCradlepointValueNull()
	var device_cert DeviceCertValue = NewDeviceCertValueNull()
	var device_updown_threshold types.Int64
	var disable_pcap types.Bool
	var disable_remote_shell types.Bool
	var gateway_updown_threshold types.Int64
	var installer InstallerValue = NewInstallerValueNull()
	var jcloud JcloudValue = NewJcloudValueNull()
	var juniper JuniperValue = NewJuniperValueNull()
	var mgmt MgmtValue = NewMgmtValueNull()
	var mist_nac MistNacValue = NewMistNacValueNull()
	// var msp_id types.String
	var mxedge_fips_enabled types.Bool
	var mxedge_mgmt MxedgeMgmtValue
	var org_id types.String
	var password_policy PasswordPolicyValue = NewPasswordPolicyValueNull()
	var pcap PcapValue = NewPcapValueNull()
	var port_channelization PortChannelizationValue = NewPortChannelizationValueNull()
	// var pcap_bucket_verified types.Bool
	var security SecurityValue = NewSecurityValueNull()
	var switch_mgmt SwitchMgmtValue = NewSwitchMgmtValueNull()
	var switch_updown_threshold types.Int64
	var synthetic_test SyntheticTestValue = NewSyntheticTestValueNull()
	var ui_idle_timeout types.Int64
	var vpn_options VpnOptionsValue = NewVpnOptionsValueNull()
	var wan_pma WanPmaValue = NewWanPmaValueNull()
	var wired_pma WiredPmaValue = NewWiredPmaValueNull()
	var wireless_pma WirelessPmaValue = NewWirelessPmaValueNull()

	if data.ApUpdownThreshold.Value() != nil {
		ap_updown_threshold = types.Int64Value(int64(*data.ApUpdownThreshold.Value()))
	}
	if data.ApiPolicy != nil {
		api_policy = apiPolicySdkToTerraform(ctx, &diags, data.ApiPolicy)
	}
	// if data.BlacklistUrl != nil {
	// 	blacklist_url = types.StringValue(*data.BlacklistUrl)
	// }
	if data.Cacerts != nil {
		cacerts = mist_transform.ListOfStringSdkToTerraform(ctx, data.Cacerts)
	}
	if data.Celona != nil {
		celona = celonaSdkToTerraform(ctx, &diags, data.Celona)
	}
	if data.Cloudshark != nil {
		cloudshark = cloudsharkSdkToTerraform(ctx, &diags, data.Cloudshark)
	}
	if data.Cradlepoint != nil {
		cradlepoint = cradlepointSdkToTerraform(ctx, &diags, data.Cradlepoint)
	}
	if data.DeviceCert != nil {
		device_cert = deviceCertSdkToTerraform(ctx, &diags, data.DeviceCert)
	}
	if data.DeviceUpdownThreshold.Value() != nil {
		device_updown_threshold = types.Int64Value(int64(*data.DeviceUpdownThreshold.Value()))
	}
	if data.DisablePcap != nil {
		disable_pcap = types.BoolValue(*data.DisablePcap)
	}
	if data.DisableRemoteShell != nil {
		disable_remote_shell = types.BoolValue(*data.DisableRemoteShell)
	}
	if data.GatewayUpdownThreshold.Value() != nil {
		gateway_updown_threshold = types.Int64Value(int64(*data.GatewayUpdownThreshold.Value()))
	}
	if data.Installer != nil {
		installer = installerSdkToTerraform(ctx, &diags, data.Installer)
	}
	if data.Jcloud != nil {
		jcloud = jcloudSdkToTerraform(ctx, &diags, data.Jcloud)
	}
	if data.Juniper != nil {
		juniper = juniperSdkToTerraform(ctx, &diags, data.Juniper)
	}
	if data.Mgmt != nil {
		mgmt = mgmtSdkToTerraform(ctx, &diags, data.Mgmt)
	}
	if data.MistNac != nil {
		mist_nac = mistNacSdkToTerraform(ctx, &diags, data.MistNac)
	}
	// if data.MspId != nil {
	// 	msp_id = types.StringValue(data.MspId.String())
	// }
	if data.MxedgeFipsEnabled != nil {
		mxedge_fips_enabled = types.BoolValue(*data.MxedgeFipsEnabled)
	}
	if data.MxedgeMgmt != nil {
		mxedge_mgmt = mxedgeMgmtSdkToTerraform(ctx, &diags, data.MxedgeMgmt)
	}
	if data.OrgId != nil {
		org_id = types.StringValue(data.OrgId.String())
	}
	if data.PasswordPolicy != nil {
		password_policy = passwordPolicySdkToTerraform(ctx, &diags, data.PasswordPolicy)
	}
	if data.Pcap != nil {
		pcap = pcapSdkToTerraform(ctx, &diags, data.Pcap)
	}
	if data.PortChannelization != nil {
		port_channelization = PortChannelSdkToTerraform(ctx, &diags, data.PortChannelization)
	}
	// if data.PcapBucketVerified != nil {
	// 	pcap_bucket_verified = types.BoolValue(*data.PcapBucketVerified)
	// }
	if data.Security != nil {
		security = securitySdkToTerraform(ctx, &diags, data.Security)
	}
	if data.SwitchMgmt != nil {
		switch_mgmt = switchMgmtSdkToTerraform(ctx, &diags, data.SwitchMgmt)
	}
	if data.SwitchUpdownThreshold.Value() != nil {
		switch_updown_threshold = types.Int64Value(int64(*data.SwitchUpdownThreshold.Value()))
	}
	if data.SyntheticTest != nil {
		synthetic_test = syntheticTestSdkToTerraform(ctx, &diags, data.SyntheticTest)
	}
	if data.UiIdleTimeout != nil {
		ui_idle_timeout = types.Int64Value(int64(*data.UiIdleTimeout))
	}
	if data.VpnOptions != nil {
		vpn_options = vpnOptionsSdkToTerraform(ctx, &diags, data.VpnOptions)
	}
	if data.WanPma != nil && data.WanPma.Enabled != nil {
		wan_pma.Enabled = types.BoolValue(*data.WanPma.Enabled)
	}
	if data.WiredPma != nil && data.WiredPma.Enabled != nil {
		wired_pma.Enabled = types.BoolValue(*data.WiredPma.Enabled)
	}
	if data.WirelessPma != nil && data.WirelessPma.Enabled != nil {
		wireless_pma.Enabled = types.BoolValue(*data.WirelessPma.Enabled)
	}

	state.ApUpdownThreshold = ap_updown_threshold
	state.ApiPolicy = api_policy
	// state.BlacklistUrl = blacklist_url
	state.Cacerts = cacerts
	state.Celona = celona
	state.Cloudshark = cloudshark
	state.Cradlepoint = cradlepoint
	state.DeviceCert = device_cert
	state.DeviceUpdownThreshold = device_updown_threshold
	state.DisablePcap = disable_pcap
	state.DisableRemoteShell = disable_remote_shell
	state.GatewayUpdownThreshold = gateway_updown_threshold
	state.Installer = installer
	state.Jcloud = jcloud
	state.Juniper = juniper
	state.Mgmt = mgmt
	state.MistNac = mist_nac
	// state.MspId = msp_id
	state.MxedgeFipsEnabled = mxedge_fips_enabled
	state.MxedgeMgmt = mxedge_mgmt
	state.OrgId = org_id
	state.PasswordPolicy = password_policy
	state.Pcap = pcap
	state.PortChannelization = port_channelization
	// state.PcapBucketVerified = pcap_bucket_verified
	state.Security = security
	state.SwitchMgmt = switch_mgmt
	state.SwitchUpdownThreshold = switch_updown_threshold
	state.SyntheticTest = synthetic_test
	state.UiIdleTimeout = ui_idle_timeout
	state.VpnOptions = vpn_options
	state.WanPma = wan_pma
	state.WiredPma = wired_pma
	state.WirelessPma = wireless_pma

	return state, diags
}
