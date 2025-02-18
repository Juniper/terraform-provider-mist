package resource_device_ap

import (
	"context"

	mistlist "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, data *models.DeviceAp) (DeviceApModel, diag.Diagnostics) {
	var state DeviceApModel
	var diags diag.Diagnostics

	var aeroscout = NewAeroscoutValueNull()
	var bleConfig = NewBleConfigValueNull()
	var centrak = NewCentrakValueNull()
	var clientBridge = NewClientBridgeValueNull()
	var disableEth1 types.Bool
	var disableEth2 types.Bool
	var disableEth3 types.Bool
	var disableModule types.Bool
	var eslConfig = NewEslConfigValueNull()
	var flowControl = types.BoolValue(false)
	var height types.Float64
	var deviceId types.String
	var image1Url = types.StringValue("not_present")
	var image2Url = types.StringValue("not_present")
	var image3Url = types.StringValue("not_present")
	var ipConfig = NewIpConfigValueNull()
	var led = NewLedValueNull()
	var locked types.Bool
	var mapId types.String
	var mesh = NewMeshValueNull()
	var name types.String
	var notes types.String
	var ntpServers = types.ListNull(types.StringType)
	var orgId types.String
	var orientation types.Int64
	var poePassthrough types.Bool
	var pwrConfig = NewPwrConfigValueNull()
	var radioConfig = NewRadioConfigValueNull()
	var siteId types.String
	var uplinkPortConfig = NewUplinkPortConfigValueNull()
	var usbConfig = NewUsbConfigValueNull()
	var vars = types.MapNull(types.StringType)
	var x types.Float64
	var y types.Float64

	var deviceType types.String
	var serial types.String
	var mac types.String
	var model types.String

	if data.Aeroscout != nil {
		aeroscout = aeroscoutSdkToTerraform(ctx, &diags, data.Aeroscout)
	}
	if data.BleConfig != nil {
		bleConfig = bleConfigsSdkToTerraform(ctx, &diags, data.BleConfig)
	}
	if data.Centrak != nil {
		centrak = centrakSdkToTerraform(ctx, &diags, data.Centrak)
	}
	if data.ClientBridge != nil {
		clientBridge = clientBridgeSdkToTerraform(ctx, &diags, data.ClientBridge)
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
	if data.EslConfig != nil {
		eslConfig = eslSdkToTerraform(ctx, &diags, data.EslConfig)
	}
	if data.FlowControl != nil {
		flowControl = types.BoolValue(*data.FlowControl)
	}
	if data.Height != nil {
		height = types.Float64Value(*data.Height)
	}
	if data.Id != nil {
		deviceId = types.StringValue(data.Id.String())
	}
	if data.Image1Url.Value() != nil {
		image1Url = types.StringValue("present")
	}
	if data.Image2Url.Value() != nil {
		image2Url = types.StringValue("present")
	}
	if data.Image3Url.Value() != nil {
		image3Url = types.StringValue("present")
	}
	if data.IpConfig != nil {
		ipConfig = ipConfigSdkToTerraform(ctx, &diags, data.IpConfig)
	}
	if data.Led != nil {
		led = ledSdkToTerraform(ctx, &diags, data.Led)
	}
	if data.Locked != nil {
		locked = types.BoolValue(*data.Locked)
	}
	if data.MapId != nil {
		mapId = types.StringValue(data.MapId.String())
	}
	if data.Mesh != nil {
		mesh = meshSdkToTerraform(ctx, &diags, data.Mesh)
	}
	if data.Name != nil {
		name = types.StringValue(*data.Name)
	}
	if data.Notes != nil {
		notes = types.StringValue(*data.Notes)
	}
	if data.NtpServers != nil {
		ntpServers = mistlist.ListOfStringSdkToTerraform(data.NtpServers)
	}
	if data.OrgId != nil {
		orgId = types.StringValue(data.OrgId.String())
	}
	if data.Orientation != nil {
		orientation = types.Int64Value(int64(*data.Orientation))
	}
	if data.PoePassthrough != nil {
		poePassthrough = types.BoolValue(*data.PoePassthrough)
	}
	if data.PwrConfig != nil {
		pwrConfig = pwrConfigSdkToTerraform(ctx, &diags, data.PwrConfig)
	}
	if data.RadioConfig != nil {
		radioConfig = radioConfigSdkToTerraform(ctx, &diags, data.RadioConfig)
	}
	if data.SiteId != nil {
		siteId = types.StringValue(data.SiteId.String())
	}
	if data.UplinkPortConfig != nil {
		uplinkPortConfig = uplinkPortConfigSdkToTerraform(ctx, &diags, data.UplinkPortConfig)
	}
	if data.UsbConfig != nil {
		usbConfig = usbConfigSdkToTerraform(ctx, &diags, data.UsbConfig)
	}
	if data.Vars != nil {
		vars = varsSdkToTerraform(ctx, &diags, data.Vars)
	}
	if data.X != nil {
		x = types.Float64Value(*data.X)
	}
	if data.Y != nil {
		y = types.Float64Value(*data.Y)
	}

	deviceType = types.StringValue(data.Type)

	if data.Serial != nil {
		serial = types.StringValue(*data.Serial)
	}

	if data.Mac != nil {
		mac = types.StringValue(*data.Mac)
	}

	if data.Model != nil {
		model = types.StringValue(*data.Model)
	}
	state.Aeroscout = aeroscout
	state.BleConfig = bleConfig
	state.Centrak = centrak
	state.ClientBridge = clientBridge
	state.DisableEth1 = disableEth1
	state.DisableEth2 = disableEth2
	state.DisableEth3 = disableEth3
	state.DisableModule = disableModule
	state.EslConfig = eslConfig
	state.FlowControl = flowControl
	state.Height = height
	state.DeviceId = deviceId
	state.Image1Url = image1Url
	state.Image2Url = image2Url
	state.Image3Url = image3Url
	state.IpConfig = ipConfig
	state.Led = led
	state.Locked = locked
	state.MapId = mapId
	state.Mesh = mesh
	state.Name = name
	state.NtpServers = ntpServers
	state.Notes = notes
	state.Orientation = orientation
	state.OrgId = orgId
	state.PoePassthrough = poePassthrough
	state.PwrConfig = pwrConfig
	state.RadioConfig = radioConfig
	state.SiteId = siteId
	state.UplinkPortConfig = uplinkPortConfig
	state.UsbConfig = usbConfig
	state.Vars = vars
	state.X = x
	state.Y = y
	state.Type = deviceType
	state.Serial = serial
	state.Mac = mac
	state.Model = model

	return state, diags
}
