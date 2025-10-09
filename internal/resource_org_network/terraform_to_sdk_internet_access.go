package resource_org_network

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func destinationNatInternetAccessTerraformToSdk(d basetypes.MapValue) map[string]models.NetworkInternetAccessDestinationNatProperty {
	dataMap := make(map[string]models.NetworkInternetAccessDestinationNatProperty)
	for k, v := range d.Elements() {
		// Extract attributes directly from the ObjectValue instead of casting to specific type
		if objVal, ok := v.(basetypes.ObjectValue); ok {
			data := models.NetworkInternetAccessDestinationNatProperty{}
			attrs := objVal.Attributes()

			if internalIp, exists := attrs["internal_ip"]; exists {
				if strVal, ok := internalIp.(basetypes.StringValue); ok {
					data.InternalIp = strVal.ValueStringPointer()
				}
			}
			if name, exists := attrs["name"]; exists {
				if strVal, ok := name.(basetypes.StringValue); ok {
					data.Name = strVal.ValueStringPointer()
				}
			}
			if port, exists := attrs["port"]; exists {
				if strVal, ok := port.(basetypes.StringValue); ok {
					data.Port = strVal.ValueStringPointer()
				}
			}
			if wanName, exists := attrs["wan_name"]; exists {
				if strVal, ok := wanName.(basetypes.StringValue); ok {
					data.WanName = strVal.ValueStringPointer()
				}
			}
			dataMap[k] = data
		}
	}
	return dataMap
}

func staticNatInternetAccessTerraformToSdk(d basetypes.MapValue) map[string]models.NetworkInternetAccessStaticNatProperty {
	dataMap := make(map[string]models.NetworkInternetAccessStaticNatProperty)
	for k, v := range d.Elements() {
		// Extract attributes directly from the ObjectValue instead of casting to specific type
		if objVal, ok := v.(basetypes.ObjectValue); ok {
			data := models.NetworkInternetAccessStaticNatProperty{}
			attrs := objVal.Attributes()

			if internalIp, exists := attrs["internal_ip"]; exists {
				if strVal, ok := internalIp.(basetypes.StringValue); ok {
					data.InternalIp = strVal.ValueStringPointer()
				}
			}
			if name, exists := attrs["name"]; exists {
				if strVal, ok := name.(basetypes.StringValue); ok {
					data.Name = strVal.ValueStringPointer()
				}
			}
			if wanName, exists := attrs["wan_name"]; exists {
				if strVal, ok := wanName.(basetypes.StringValue); ok {
					data.WanName = strVal.ValueStringPointer()
				}
			}
			dataMap[k] = data
		}
	}
	return dataMap
}

func internetAccessTerraformToSdk(d InternetAccessValue) *models.NetworkInternetAccess {
	data := models.NetworkInternetAccess{}

	if !d.CreateSimpleServicePolicy.IsNull() && !d.CreateSimpleServicePolicy.IsUnknown() {
		data.CreateSimpleServicePolicy = d.CreateSimpleServicePolicy.ValueBoolPointer()
	}
	if !d.InternetAccessDestinationNat.IsNull() && !d.InternetAccessDestinationNat.IsUnknown() {
		data.DestinationNat = destinationNatInternetAccessTerraformToSdk(d.InternetAccessDestinationNat)
	}
	if !d.Enabled.IsNull() && !d.Enabled.IsUnknown() {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if !d.InternetAccessStaticNat.IsNull() && !d.InternetAccessStaticNat.IsUnknown() {
		data.StaticNat = staticNatInternetAccessTerraformToSdk(d.InternetAccessStaticNat)
	}
	if !d.Restricted.IsNull() && !d.Restricted.IsUnknown() {
		data.Restricted = d.Restricted.ValueBoolPointer()
	}

	return &data
}
