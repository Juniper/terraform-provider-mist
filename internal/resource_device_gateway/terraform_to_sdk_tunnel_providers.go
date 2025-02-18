package resource_device_gateway

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func tunnelProviderOptionsJseTerraformToSdk(ctx context.Context, d basetypes.ObjectValue) models.TunnelProviderOptionsJse {
	data := models.TunnelProviderOptionsJse{}
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewJseValueMust(d.AttributeTypes(ctx), d.Attributes())
		if plan.OrgName.ValueStringPointer() != nil {
			data.OrgName = plan.OrgName.ValueStringPointer()
		}
		if plan.NumUsers.ValueInt64Pointer() != nil {
			data.NumUsers = models.ToPointer(int(plan.NumUsers.ValueInt64()))
		}
		return data
	}
}

func tunnelProviderOptionsZscalerSubLocationTerraformToSdk(d basetypes.ListValue) []models.TunnelProviderOptionsZscalerSubLocation {
	var dataList []models.TunnelProviderOptionsZscalerSubLocation
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(SubLocationsValue)
		data := models.TunnelProviderOptionsZscalerSubLocation{}
		if plan.AupBlockInternetUntilAccepted.ValueBoolPointer() != nil {
			data.AupBlockInternetUntilAccepted = plan.AupBlockInternetUntilAccepted.ValueBoolPointer()
		}
		if plan.AupEnabled.ValueBoolPointer() != nil {
			data.AupEnabled = plan.AupEnabled.ValueBoolPointer()
		}
		if plan.AupForceSslInspection.ValueBoolPointer() != nil {
			data.AupForceSslInspection = plan.AupForceSslInspection.ValueBoolPointer()
		}
		if plan.AupTimeoutInDays.ValueInt64Pointer() != nil {
			data.AupTimeoutInDays = models.ToPointer(int(plan.AupTimeoutInDays.ValueInt64()))
		}
		if plan.AuthRequired.ValueBoolPointer() != nil {
			data.AuthRequired = plan.AuthRequired.ValueBoolPointer()
		}
		if plan.CautionEnabled.ValueBoolPointer() != nil {
			data.CautionEnabled = plan.CautionEnabled.ValueBoolPointer()
		}
		if plan.DnBandwidth.ValueFloat64Pointer() != nil {
			data.DnBandwidth = models.NewOptional(plan.DnBandwidth.ValueFloat64Pointer())
		}
		if plan.IdleTimeInMinutes.ValueInt64Pointer() != nil {
			data.IdleTimeInMinutes = models.ToPointer(int(plan.IdleTimeInMinutes.ValueInt64()))
		}
		if plan.Name.ValueStringPointer() != nil {
			data.Name = plan.Name.ValueStringPointer()
		}
		if plan.OfwEnabled.ValueBoolPointer() != nil {
			data.OfwEnabled = plan.OfwEnabled.ValueBoolPointer()
		}
		if plan.SurrogateIp.ValueBoolPointer() != nil {
			data.SurrogateIP = plan.SurrogateIp.ValueBoolPointer()
		}
		if plan.SurrogateIpEnforcedForKnownBrowsers.ValueBoolPointer() != nil {
			data.SurrogateIPEnforcedForKnownBrowsers = plan.SurrogateIpEnforcedForKnownBrowsers.ValueBoolPointer()
		}
		if plan.SurrogateRefreshTimeInMinutes.ValueInt64Pointer() != nil {
			data.SurrogateRefreshTimeInMinutes = models.ToPointer(int(plan.SurrogateRefreshTimeInMinutes.ValueInt64()))
		}
		if plan.UpBandwidth.ValueFloat64Pointer() != nil {
			data.UpBandwidth = models.NewOptional(plan.UpBandwidth.ValueFloat64Pointer())
		}

		dataList = append(dataList, data)
	}
	return dataList
}

func tunnelProviderOptionsZscalerTerraformToSdk(ctx context.Context, d basetypes.ObjectValue) models.TunnelProviderOptionsZscaler {
	data := models.TunnelProviderOptionsZscaler{}
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewZscalerValueMust(d.AttributeTypes(ctx), d.Attributes())
		if plan.AupBlockInternetUntilAccepted.ValueBoolPointer() != nil {
			data.AupBlockInternetUntilAccepted = plan.AupBlockInternetUntilAccepted.ValueBoolPointer()
		}
		if plan.AupEnabled.ValueBoolPointer() != nil {
			data.AupEnabled = plan.AupEnabled.ValueBoolPointer()
		}
		if plan.AupForceSslInspection.ValueBoolPointer() != nil {
			data.AupForceSslInspection = plan.AupForceSslInspection.ValueBoolPointer()
		}
		if plan.AupTimeoutInDays.ValueInt64Pointer() != nil {
			data.AupTimeoutInDays = models.ToPointer(int(plan.AupTimeoutInDays.ValueInt64()))
		}
		if plan.AuthRequired.ValueBoolPointer() != nil {
			data.AuthRequired = plan.AuthRequired.ValueBoolPointer()
		}
		if plan.CautionEnabled.ValueBoolPointer() != nil {
			data.CautionEnabled = plan.CautionEnabled.ValueBoolPointer()
		}
		if plan.DnBandwidth.ValueFloat64Pointer() != nil {
			data.DnBandwidth = models.NewOptional(plan.DnBandwidth.ValueFloat64Pointer())
		}
		if plan.IdleTimeInMinutes.ValueInt64Pointer() != nil {
			data.IdleTimeInMinutes = models.ToPointer(int(plan.IdleTimeInMinutes.ValueInt64()))
		}
		if !plan.SubLocations.IsNull() && !plan.SubLocations.IsUnknown() {
			subLocations := tunnelProviderOptionsZscalerSubLocationTerraformToSdk(plan.SubLocations)
			data.SubLocations = subLocations
		}
		if plan.OfwEnabled.ValueBoolPointer() != nil {
			data.OfwEnabled = plan.OfwEnabled.ValueBoolPointer()
		}
		if plan.SurrogateIp.ValueBoolPointer() != nil {
			data.SurrogateIP = plan.SurrogateIp.ValueBoolPointer()
		}
		if plan.SurrogateIpEnforcedForKnownBrowsers.ValueBoolPointer() != nil {
			data.SurrogateIPEnforcedForKnownBrowsers = plan.SurrogateIpEnforcedForKnownBrowsers.ValueBoolPointer()
		}
		if plan.SurrogateRefreshTimeInMinutes.ValueInt64Pointer() != nil {
			data.SurrogateRefreshTimeInMinutes = models.ToPointer(int(plan.SurrogateRefreshTimeInMinutes.ValueInt64()))
		}
		if plan.UpBandwidth.ValueFloat64Pointer() != nil {
			data.UpBandwidth = models.NewOptional(plan.UpBandwidth.ValueFloat64Pointer())
		}
		if plan.XffForwardEnabled.ValueBoolPointer() != nil {
			data.XffForwardEnabled = plan.XffForwardEnabled.ValueBoolPointer()
		}

		return data
	}
}

func tunnelProviderOptionsTerraformToSdk(ctx context.Context, d TunnelProviderOptionsValue) *models.TunnelProviderOptions {

	data := models.TunnelProviderOptions{}

	jse := tunnelProviderOptionsJseTerraformToSdk(ctx, d.Jse)
	if !d.Jse.IsNull() && !d.Jse.IsUnknown() {
		data.Jse = &jse
	}

	zscaler := tunnelProviderOptionsZscalerTerraformToSdk(ctx, d.Zscaler)
	if !d.Zscaler.IsNull() && !d.Zscaler.IsUnknown() {
		data.Zscaler = &zscaler
	}

	return &data
}
