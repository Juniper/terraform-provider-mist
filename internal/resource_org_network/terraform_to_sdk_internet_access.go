package resource_org_network

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func destinationNatInternetAccesTerraformToSdk(d basetypes.MapValue) map[string]models.NetworkInternetAccessDestinationNatProperty {
	dataMap := make(map[string]models.NetworkInternetAccessDestinationNatProperty)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		vPlan := vInterface.(InternetAccessDestinationNatValue)
		data := models.NetworkInternetAccessDestinationNatProperty{}
		data.InternalIp = vPlan.InternalIp.ValueStringPointer()
		data.Name = vPlan.Name.ValueStringPointer()
		data.Port = vPlan.Port.ValueStringPointer()
		data.WanName = vPlan.WanName.ValueStringPointer()
		dataMap[k] = data
	}
	return dataMap
}

func staticNatInternetAccesTerraformToSdk(d basetypes.MapValue) map[string]models.NetworkInternetAccessStaticNatProperty {
	dataMap := make(map[string]models.NetworkInternetAccessStaticNatProperty)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		vPlan := vInterface.(InternetAccessStaticNatValue)
		data := models.NetworkInternetAccessStaticNatProperty{}
		data.InternalIp = vPlan.InternalIp.ValueStringPointer()
		data.Name = vPlan.Name.ValueStringPointer()
		data.WanName = vPlan.WanName.ValueStringPointer()
		dataMap[k] = data
	}
	return dataMap
}

func InternetAccessTerraformToSdk(d InternetAccessValue) *models.NetworkInternetAccess {
	data := models.NetworkInternetAccess{}

	if !d.CreateSimpleServicePolicy.IsNull() && !d.CreateSimpleServicePolicy.IsUnknown() {
		data.CreateSimpleServicePolicy = d.CreateSimpleServicePolicy.ValueBoolPointer()
	}
	if !d.InternetAccessDestinationNat.IsNull() && !d.InternetAccessDestinationNat.IsUnknown() {
		data.DestinationNat = destinationNatInternetAccesTerraformToSdk(d.InternetAccessDestinationNat)
	}
	if !d.Enabled.IsNull() && !d.Enabled.IsUnknown() {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if !d.InternetAccessStaticNat.IsNull() && !d.InternetAccessStaticNat.IsUnknown() {
		data.StaticNat = staticNatInternetAccesTerraformToSdk(d.InternetAccessStaticNat)
	}
	if !d.Restricted.IsNull() && !d.Restricted.IsUnknown() {
		data.Restricted = d.Restricted.ValueBoolPointer()
	}

	return &data
}
