package resource_org_deviceprofile_ap

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, data *models.DeviceprofileAp) (OrgDeviceprofileApModel, diag.Diagnostics) {
	var state OrgDeviceprofileApModel
	var diags diag.Diagnostics

	var aeroscout = NewAeroscoutValueNull()
	var airista = NewAiristaValueNull()
	var bleConfig = NewBleConfigValueNull()
	var disableEth1 types.Bool
	var disableEth2 types.Bool
	var disableEth3 types.Bool
	var disableModule types.Bool
	var eslConfig = NewEslConfigValueNull()
	var ipConfig = NewIpConfigValueNull()
	var lacpConfig = NewLacpConfigValueNull()
	var led = NewLedValueNull()
	var mesh = NewMeshValueNull()
	var name types.String
	var ntpServers = types.ListNull(types.StringType)
	var orgId types.String
	var poePassthrough types.Bool
	var profileId types.String
	var portConfig = types.MapNull(PortConfigValue{}.Type(ctx))
	var pwrConfig = NewPwrConfigValueNull()
	var radioConfig = NewRadioConfigValueNull()
	var siteId types.String
	var uplinkPortConfig = NewUplinkPortConfigValueNull()
	var usbConfig = NewUsbConfigValueNull()
	var vars = types.MapNull(types.StringType)

	var profileType types.String

	if !mistutils.IsSdkDataEmpty(data.Aeroscout) {
		aeroscout = aeroscoutSdkToTerraform(ctx, &diags, data.Aeroscout)
	}
	if !mistutils.IsSdkDataEmpty(data.Airista) {
		airista = airistaSdkToTerraform(ctx, &diags, data.Airista)
	}
	if !mistutils.IsSdkDataEmpty(data.BleConfig) {
		bleConfig = bleConfigSdkToTerraform(ctx, &diags, data.BleConfig)
	}
	if data.DisableEth1 != nil {
		disableEth1 = types.BoolValue(*data.DisableEth1)
	}
	if data.DisableEth2 != nil {
		disableEth2 = types.BoolValue(*data.DisableEth2)
	}
	if data.DisableEth3 != nil {
		disableEth3 = types.BoolValue(*data.DisableEth3)
	}
	if data.DisableModule != nil {
		disableModule = types.BoolValue(*data.DisableModule)
	}
	if !mistutils.IsSdkDataEmpty(data.EslConfig) {
		eslConfig = eslSdkToTerraform(ctx, &diags, data.EslConfig)
	}
	if data.Id != nil {
		profileId = types.StringValue(data.Id.String())
	}
	if !mistutils.IsSdkDataEmpty(data.IpConfig) {
		ipConfig = ipConfigSdkToTerraform(ctx, &diags, data.IpConfig)
	}
	if !mistutils.IsSdkDataEmpty(data.LacpConfig) {
		lacpConfig = lacpConfigSdkToTerraform(ctx, &diags, data.LacpConfig)
	}
	if !mistutils.IsSdkDataEmpty(data.Led) {
		led = ledSdkToTerraform(ctx, &diags, data.Led)
	}
	if !mistutils.IsSdkDataEmpty(data.Mesh) {
		mesh = meshSdkToTerraform(ctx, &diags, data.Mesh)
	}
	if data.Name.Value() != nil {
		name = types.StringValue(*data.Name.Value())
	}

	if len(data.NtpServers) > 0 {
		ntpServers = mistutils.ListOfStringSdkToTerraform(data.NtpServers)
	}
	if data.OrgId != nil {
		orgId = types.StringValue(data.OrgId.String())
	}
	if data.PoePassthrough != nil {
		poePassthrough = types.BoolValue(*data.PoePassthrough)
	}
	if data.PortConfig != nil {
		portConfig = portConfigSdkToTerraform(ctx, &diags, data.PortConfig)
	}
	if !mistutils.IsSdkDataEmpty(data.PwrConfig) {
		pwrConfig = pwrConfigSdkToTerraform(ctx, &diags, data.PwrConfig)
	}
	if !mistutils.IsSdkDataEmpty(data.RadioConfig) {
		radioConfig = radioConfigSdkToTerraform(ctx, &diags, data.RadioConfig)
	}
	if data.SiteId != nil {
		siteId = types.StringValue(data.SiteId.String())
	}
	if !mistutils.IsSdkDataEmpty(data.UplinkPortConfig) {
		uplinkPortConfig = uplinkPortConfigSdkToTerraform(ctx, &diags, data.UplinkPortConfig)
	}
	if !mistutils.IsSdkDataEmpty(data.UsbConfig) {
		usbConfig = usbConfigSdkToTerraform(ctx, &diags, data.UsbConfig)
	}
	if len(data.Vars) > 0 {
		vars = varsSdkToTerraform(ctx, &diags, data.Vars)
	}

	profileType = types.StringValue(data.Type)

	state.Aeroscout = aeroscout
	state.Airista = airista
	state.BleConfig = bleConfig
	state.DisableEth1 = disableEth1
	state.DisableEth2 = disableEth2
	state.DisableEth3 = disableEth3
	state.DisableModule = disableModule
	state.EslConfig = eslConfig
	state.Id = profileId
	state.IpConfig = ipConfig
	state.LacpConfig = lacpConfig
	state.Led = led
	state.Mesh = mesh
	state.Name = name
	state.NtpServers = ntpServers
	state.OrgId = orgId
	state.PoePassthrough = poePassthrough
	state.PortConfig = portConfig
	state.PwrConfig = pwrConfig
	state.RadioConfig = radioConfig
	state.SiteId = siteId
	state.UplinkPortConfig = uplinkPortConfig
	state.UsbConfig = usbConfig
	state.Vars = vars
	state.Type = profileType

	return state, diags
}
