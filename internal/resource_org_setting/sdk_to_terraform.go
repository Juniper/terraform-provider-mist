package resource_org_setting

import (
	"context"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, data *models.OrgSetting) (OrgSettingModel, diag.Diagnostics) {
	var state OrgSettingModel
	var diags diag.Diagnostics

	var apUpdownThreshold types.Int64
	var apiPolicy = NewApiPolicyValueNull()
	// var blacklist_url types.String
	var cacerts = types.ListNull(types.StringType)
	var celona = NewCelonaValueNull()
	var cloudshark = NewCloudsharkValueNull()
	var cradlepoint = NewCradlepointValueNull()
	var deviceCert = NewDeviceCertValueNull()
	var deviceUpdownThreshold types.Int64
	var disablePcap types.Bool
	var disableRemoteShell types.Bool
	var gatewayUpdownThreshold types.Int64
	var installer = NewInstallerValueNull()
	var jcloud = NewJcloudValueNull()
	var jcloudRa = NewJcloudRaValueNull()
	var juniper = NewJuniperValueNull()
	var junosShellAccess = NewJunosShellAccessValueNull()
	var mgmt = NewMgmtValueNull()
	var mistNac = NewMistNacValueNull()
	// var msp_id types.String
	var mxedgeFipsEnabled types.Bool
	var mxedgeMgmt MxedgeMgmtValue
	var opticPortConfig = types.MapNull(OpticPortConfigValue{}.Type(ctx))
	var orgId types.String
	var passwordPolicy = NewPasswordPolicyValueNull()
	var pcap = NewPcapValueNull()
	// var pcap_bucket_verified types.Bool
	var security = NewSecurityValueNull()
	var switchMgmt = NewSwitchMgmtValueNull()
	var switchUpdownThreshold types.Int64
	var syntheticTest = NewSyntheticTestValueNull()
	var uiIdleTimeout types.Int64
	var vpnOptions = NewVpnOptionsValueNull()
	var wanPma = NewWanPmaValueNull()
	var wiredPma = NewWiredPmaValueNull()
	var wirelessPma = NewWirelessPmaValueNull()

	if data.ApUpdownThreshold.Value() != nil {
		apUpdownThreshold = types.Int64Value(int64(*data.ApUpdownThreshold.Value()))
	}
	if data.ApiPolicy != nil {
		apiPolicy = apiPolicySdkToTerraform(ctx, &diags, data.ApiPolicy)
	}
	// if data.BlacklistUrl != nil {
	// 	blacklist_url = types.StringValue(*data.BlacklistUrl)
	// }
	if data.Cacerts != nil {
		cacerts = misttransform.ListOfStringSdkToTerraform(data.Cacerts)
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
		deviceCert = deviceCertSdkToTerraform(ctx, &diags, data.DeviceCert)
	}
	if data.DeviceUpdownThreshold.Value() != nil {
		deviceUpdownThreshold = types.Int64Value(int64(*data.DeviceUpdownThreshold.Value()))
	}
	if data.DisablePcap != nil {
		disablePcap = types.BoolValue(*data.DisablePcap)
	}
	if data.DisableRemoteShell != nil {
		disableRemoteShell = types.BoolValue(*data.DisableRemoteShell)
	}
	if data.GatewayUpdownThreshold.Value() != nil {
		gatewayUpdownThreshold = types.Int64Value(int64(*data.GatewayUpdownThreshold.Value()))
	}
	if data.Installer != nil {
		installer = installerSdkToTerraform(ctx, &diags, data.Installer)
	}
	if data.Jcloud != nil {
		jcloud = jcloudSdkToTerraform(ctx, &diags, data.Jcloud)
	}
	if data.JcloudRa != nil {
		jcloudRa = jcloudRaSdkToTerraform(ctx, &diags, data.JcloudRa)
	}
	if data.Juniper != nil {
		juniper = juniperSdkToTerraform(ctx, &diags, data.Juniper)
	}
	if data.JunosShellAccess != nil {
		junosShellAccess = junosShellAccessSdkToTerraform(ctx, &diags, data.JunosShellAccess)
	}
	if data.Mgmt != nil {
		mgmt = mgmtSdkToTerraform(ctx, &diags, data.Mgmt)
	}
	if data.MistNac != nil {
		mistNac = mistNacSdkToTerraform(ctx, &diags, data.MistNac)
	}
	// if data.MspId != nil {
	// 	msp_id = types.StringValue(data.MspId.String())
	// }
	if data.MxedgeFipsEnabled != nil {
		mxedgeFipsEnabled = types.BoolValue(*data.MxedgeFipsEnabled)
	}
	if data.MxedgeMgmt != nil {
		mxedgeMgmt = mxedgeMgmtSdkToTerraform(ctx, &diags, data.MxedgeMgmt)
	}
	if data.OpticPortConfig != nil {
		opticPortConfig = opticPortConfigSdkToTerraform(ctx, &diags, data.OpticPortConfig)
	}
	if data.OrgId != nil {
		orgId = types.StringValue(data.OrgId.String())
	}
	if data.PasswordPolicy != nil {
		passwordPolicy = passwordPolicySdkToTerraform(ctx, &diags, data.PasswordPolicy)
	}
	if data.Pcap != nil {
		pcap = pcapSdkToTerraform(ctx, &diags, data.Pcap)
	}
	// if data.PcapBucketVerified != nil {
	// 	pcap_bucket_verified = types.BoolValue(*data.PcapBucketVerified)
	// }
	if data.Security != nil {
		security = securitySdkToTerraform(ctx, &diags, data.Security)
	}
	if data.SwitchMgmt != nil {
		switchMgmt = switchMgmtSdkToTerraform(ctx, &diags, data.SwitchMgmt)
	}
	if data.SwitchUpdownThreshold.Value() != nil {
		switchUpdownThreshold = types.Int64Value(int64(*data.SwitchUpdownThreshold.Value()))
	}
	if data.SyntheticTest != nil {
		syntheticTest = syntheticTestSdkToTerraform(ctx, &diags, data.SyntheticTest)
	}
	if data.UiIdleTimeout != nil {
		uiIdleTimeout = types.Int64Value(int64(*data.UiIdleTimeout))
	}
	if data.VpnOptions != nil {
		vpnOptions = vpnOptionsSdkToTerraform(ctx, &diags, data.VpnOptions)
	}
	if data.WanPma != nil && data.WanPma.Enabled != nil {
		wanPma.Enabled = types.BoolValue(*data.WanPma.Enabled)
	}
	if data.WiredPma != nil && data.WiredPma.Enabled != nil {
		wiredPma.Enabled = types.BoolValue(*data.WiredPma.Enabled)
	}
	if data.WirelessPma != nil && data.WirelessPma.Enabled != nil {
		wirelessPma.Enabled = types.BoolValue(*data.WirelessPma.Enabled)
	}

	state.ApUpdownThreshold = apUpdownThreshold
	state.ApiPolicy = apiPolicy
	// state.BlacklistUrl = blacklist_url
	state.Cacerts = cacerts
	state.Celona = celona
	state.Cloudshark = cloudshark
	state.Cradlepoint = cradlepoint
	state.DeviceCert = deviceCert
	state.DeviceUpdownThreshold = deviceUpdownThreshold
	state.DisablePcap = disablePcap
	state.DisableRemoteShell = disableRemoteShell
	state.GatewayUpdownThreshold = gatewayUpdownThreshold
	state.Installer = installer
	state.Jcloud = jcloud
	state.JcloudRa = jcloudRa
	state.Juniper = juniper
	state.JunosShellAccess = junosShellAccess
	state.Mgmt = mgmt
	state.MistNac = mistNac
	// state.MspId = msp_id
	state.MxedgeFipsEnabled = mxedgeFipsEnabled
	state.MxedgeMgmt = mxedgeMgmt
	state.OpticPortConfig = opticPortConfig
	state.OrgId = orgId
	state.PasswordPolicy = passwordPolicy
	state.Pcap = pcap
	// state.PcapBucketVerified = pcap_bucket_verified
	state.Security = security
	state.SwitchMgmt = switchMgmt
	state.SwitchUpdownThreshold = switchUpdownThreshold
	state.SyntheticTest = syntheticTest
	state.UiIdleTimeout = uiIdleTimeout
	state.VpnOptions = vpnOptions
	state.WanPma = wanPma
	state.WiredPma = wiredPma
	state.WirelessPma = wirelessPma

	return state, diags
}
