package resource_org_setting

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, data *models.OrgSetting) (OrgSettingModel, diag.Diagnostics) {
	var state OrgSettingModel
	var diags diag.Diagnostics

	var allowMist types.Bool
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
	var gatewayTunnelUpdownThreshold types.Int64
	var gatewayUpdownThreshold types.Int64
	var installer = NewInstallerValueNull()
	var jcloud = NewJcloudValueNull()
	var jcloudRa = NewJcloudRaValueNull()
	var juniper = NewJuniperValueNull()
	var juniperSrx = NewJuniperSrxValueNull()
	var junosShellAccess = NewJunosShellAccessValueNull()
	var marvis = NewMarvisValueNull()
	var mgmt = NewMgmtValueNull()
	var mistNac = NewMistNacValueNull()
	// var msp_id types.String
	var mxedgeMgmt MxedgeMgmtValue
	var opticPortConfig = types.MapNull(OpticPortConfigValue{}.Type(ctx))
	var orgId types.String
	var passwordPolicy = NewPasswordPolicyValueNull()
	var pcap = NewPcapValueNull()
	// var pcap_bucket_verified types.Bool
	var security = NewSecurityValueNull()
	var ssr = NewSsrValueNull()
	var switchUpgrade = NewSwitchValueNull() // This is the switch auto upgrade setting
	var switchMgmt = NewSwitchMgmtValueNull()
	var switchUpdownThreshold types.Int64
	var syntheticTest = NewSyntheticTestValueNull()
	var uiIdleTimeout types.Int64
	var uiNoTracking types.Bool
	var vpnOptions = NewVpnOptionsValueNull()
	var wanPma = NewWanPmaValueNull()
	var wiredPma = NewWiredPmaValueNull()
	var wirelessPma = NewWirelessPmaValueNull()

	if data.AllowMist != nil {
		allowMist = types.BoolValue(*data.AllowMist)
	}
	if data.ApUpdownThreshold.Value() != nil {
		apUpdownThreshold = types.Int64Value(int64(*data.ApUpdownThreshold.Value()))
	}
	if !mistutils.IsSdkDataEmpty(data.ApiPolicy) {
		apiPolicy = apiPolicySdkToTerraform(ctx, &diags, data.ApiPolicy)
	}
	// if data.BlacklistUrl != nil {
	// 	blacklist_url = types.StringValue(*data.BlacklistUrl)
	// }
	if len(data.Cacerts) > 0 {
		cacerts = mistutils.ListOfStringSdkToTerraform(data.Cacerts)
	}
	if !mistutils.IsSdkDataEmpty(data.Celona) {
		celona = celonaSdkToTerraform(ctx, &diags, data.Celona)
	}
	if !mistutils.IsSdkDataEmpty(data.Cloudshark) {
		cloudshark = cloudsharkSdkToTerraform(ctx, &diags, data.Cloudshark)
	}
	if !mistutils.IsSdkDataEmpty(data.Cradlepoint) {
		cradlepoint = cradlepointSdkToTerraform(ctx, &diags, data.Cradlepoint)
	}
	if !mistutils.IsSdkDataEmpty(data.DeviceCert) {
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
	if data.GatewayTunnelUpdownThreshold.Value() != nil {
		gatewayTunnelUpdownThreshold = types.Int64Value(int64(*data.GatewayTunnelUpdownThreshold.Value()))
	}
	if data.GatewayUpdownThreshold.Value() != nil {
		gatewayUpdownThreshold = types.Int64Value(int64(*data.GatewayUpdownThreshold.Value()))
	}
	if !mistutils.IsSdkDataEmpty(data.Installer) {
		installer = installerSdkToTerraform(ctx, &diags, data.Installer)
	}
	if !mistutils.IsSdkDataEmpty(data.Jcloud) {
		jcloud = jcloudSdkToTerraform(ctx, &diags, data.Jcloud)
	}
	if !mistutils.IsSdkDataEmpty(data.JcloudRa) {
		jcloudRa = jcloudRaSdkToTerraform(ctx, &diags, data.JcloudRa)
	}
	if !mistutils.IsSdkDataEmpty(data.Juniper) {
		juniper = juniperSdkToTerraform(ctx, &diags, data.Juniper)
	}
	if !mistutils.IsSdkDataEmpty(data.JuniperSrx) {
		juniperSrx = juniperSrxSdkToTerraform(ctx, &diags, data.JuniperSrx)
	}
	if !mistutils.IsSdkDataEmpty(data.JunosShellAccess) {
		junosShellAccess = junosShellAccessSdkToTerraform(ctx, &diags, data.JunosShellAccess)
	}
	if !mistutils.IsSdkDataEmpty(data.Marvis) {
		marvis = marvisSdkToTerraform(ctx, &diags, data.Marvis)
	}
	if !mistutils.IsSdkDataEmpty(data.Mgmt) {
		mgmt = mgmtSdkToTerraform(ctx, &diags, data.Mgmt)
	}
	if !mistutils.IsSdkDataEmpty(data.MistNac) {
		mistNac = mistNacSdkToTerraform(ctx, &diags, data.MistNac)
	}
	// if data.MspId != nil {
	// 	msp_id = types.StringValue(data.MspId.String())
	// }
	if data.MxedgeMgmt != nil {
		mxedgeMgmt = mxedgeMgmtSdkToTerraform(ctx, &diags, data.MxedgeMgmt)
	}
	if data.OpticPortConfig != nil {
		opticPortConfig = opticPortConfigSdkToTerraform(ctx, &diags, data.OpticPortConfig)
	}
	if data.OrgId != nil {
		orgId = types.StringValue(data.OrgId.String())
	}
	if !mistutils.IsSdkDataEmpty(data.PasswordPolicy) {
		passwordPolicy = passwordPolicySdkToTerraform(ctx, &diags, data.PasswordPolicy)
	}
	if !mistutils.IsSdkDataEmpty(data.Pcap) {
		pcap = pcapSdkToTerraform(ctx, &diags, data.Pcap)
	}
	// if data.PcapBucketVerified != nil {
	// 	pcap_bucket_verified = types.BoolValue(*data.PcapBucketVerified)
	// }
	if !mistutils.IsSdkDataEmpty(data.Security) {
		security = securitySdkToTerraform(ctx, &diags, data.Security)
	}
	if data.Switch != nil {
		switchUpgrade = switchSdkToTerraform(ctx, &diags, data.Switch)
	}
	if !mistutils.IsSdkDataEmpty(data.Ssr) {
		ssr = ssrSdkToTerraform(ctx, &diags, data.Ssr)
	}
	if !mistutils.IsSdkDataEmpty(data.SwitchMgmt) {
		switchMgmt = switchMgmtSdkToTerraform(ctx, &diags, data.SwitchMgmt)
	}
	if data.SwitchUpdownThreshold.Value() != nil {
		switchUpdownThreshold = types.Int64Value(int64(*data.SwitchUpdownThreshold.Value()))
	}
	if !mistutils.IsSdkDataEmpty(data.SyntheticTest) {
		syntheticTest = syntheticTestSdkToTerraform(ctx, &diags, data.SyntheticTest)
	}
	if data.UiIdleTimeout != nil {
		uiIdleTimeout = types.Int64Value(int64(*data.UiIdleTimeout))
	}
	if data.UiNoTracking != nil {
		uiNoTracking = types.BoolValue(*data.UiNoTracking)
	}
	if !mistutils.IsSdkDataEmpty(data.VpnOptions) {
		vpnOptions = vpnOptionsSdkToTerraform(ctx, &diags, data.VpnOptions)
	}
	if !mistutils.IsSdkDataEmpty(data.WanPma) && data.WanPma.Enabled != nil {
		var tempDiags diag.Diagnostics
		wanPma, tempDiags = NewWanPmaValue(WanPmaValue{}.AttributeTypes(ctx), map[string]attr.Value{
			"enabled": types.BoolValue(*data.WanPma.Enabled),
		})
		if tempDiags.HasError() {
			wanPma = NewWanPmaValueNull()
		}
	}
	if !mistutils.IsSdkDataEmpty(data.WiredPma) && data.WiredPma.Enabled != nil {
		var tempDiags diag.Diagnostics
		wiredPma, tempDiags = NewWiredPmaValue(WiredPmaValue{}.AttributeTypes(ctx), map[string]attr.Value{
			"enabled": types.BoolValue(*data.WiredPma.Enabled),
		})
		if tempDiags.HasError() {
			wiredPma = NewWiredPmaValueNull()
		}
	}
	if !mistutils.IsSdkDataEmpty(data.WirelessPma) && data.WirelessPma.Enabled != nil {
		var tempDiags diag.Diagnostics
		wirelessPma, tempDiags = NewWirelessPmaValue(WirelessPmaValue{}.AttributeTypes(ctx), map[string]attr.Value{
			"enabled": types.BoolValue(*data.WirelessPma.Enabled),
		})
		if tempDiags.HasError() {
			wirelessPma = NewWirelessPmaValueNull()
		}
	}

	state.AllowMist = allowMist
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
	state.GatewayTunnelUpdownThreshold = gatewayTunnelUpdownThreshold
	state.GatewayUpdownThreshold = gatewayUpdownThreshold
	state.Installer = installer
	state.Jcloud = jcloud
	state.JcloudRa = jcloudRa
	state.Juniper = juniper
	state.JuniperSrx = juniperSrx
	state.JunosShellAccess = junosShellAccess
	state.Marvis = marvis
	state.Mgmt = mgmt
	state.MistNac = mistNac
	// state.MspId = msp_id
	state.MxedgeMgmt = mxedgeMgmt
	state.OpticPortConfig = opticPortConfig
	state.OrgId = orgId
	state.PasswordPolicy = passwordPolicy
	state.Pcap = pcap
	// state.PcapBucketVerified = pcap_bucket_verified
	state.Security = security
	state.Switch = switchUpgrade
	state.Ssr = ssr
	state.SwitchMgmt = switchMgmt
	state.SwitchUpdownThreshold = switchUpdownThreshold
	state.SyntheticTest = syntheticTest
	state.UiIdleTimeout = uiIdleTimeout
	state.UiNoTracking = uiNoTracking
	state.VpnOptions = vpnOptions
	state.WanPma = wanPma
	state.WiredPma = wiredPma
	state.WirelessPma = wirelessPma

	return state, diags
}
