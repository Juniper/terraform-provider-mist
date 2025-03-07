package resource_org_network

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func destinationNatVpnTerraformToSdk(d basetypes.MapValue) map[string]models.NetworkVpnAccessDestinationNatProperty {
	dataMap := make(map[string]models.NetworkVpnAccessDestinationNatProperty)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		vPlan := vInterface.(VpnAccessDestinationNatValue)
		data := models.NetworkVpnAccessDestinationNatProperty{}
		data.InternalIp = vPlan.InternalIp.ValueStringPointer()
		data.Name = vPlan.Name.ValueStringPointer()
		data.Port = vPlan.Port.ValueStringPointer()
		dataMap[k] = data
	}
	return dataMap
}

func staticNatVpnTerraformToSdk(d basetypes.MapValue) map[string]models.NetworkVpnAccessStaticNatProperty {
	dataMap := make(map[string]models.NetworkVpnAccessStaticNatProperty)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		vPlan := vInterface.(VpnAccessStaticNatValue)
		data := models.NetworkVpnAccessStaticNatProperty{}
		data.InternalIp = vPlan.InternalIp.ValueStringPointer()
		data.Name = vPlan.Name.ValueStringPointer()
		dataMap[k] = data
	}
	return dataMap
}

func VpnTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.NetworkVpnAccessConfig {
	dataMap := make(map[string]models.NetworkVpnAccessConfig)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(VpnAccessValue)

		data := models.NetworkVpnAccessConfig{}
		if plan.AdvertisedSubnet.ValueStringPointer() != nil {
			data.AdvertisedSubnet = plan.AdvertisedSubnet.ValueStringPointer()
		}
		if plan.AllowPing.ValueBoolPointer() != nil {
			data.AllowPing = plan.AllowPing.ValueBoolPointer()
		}
		if !plan.VpnAccessDestinationNat.IsNull() && !plan.VpnAccessDestinationNat.IsUnknown() {
			data.DestinationNat = destinationNatVpnTerraformToSdk(plan.VpnAccessDestinationNat)
		}
		if plan.NatPool.ValueStringPointer() != nil {
			data.NatPool = plan.NatPool.ValueStringPointer()
		}
		if plan.NoReadvertiseToLanBgp.ValueBoolPointer() != nil {
			data.NoReadvertiseToLanBgp = plan.NoReadvertiseToLanBgp.ValueBoolPointer()
		}
		if plan.NoReadvertiseToLanOspf.ValueBoolPointer() != nil {
			data.NoReadvertiseToLanOspf = plan.NoReadvertiseToLanOspf.ValueBoolPointer()
		}
		if plan.NoReadvertiseToOverlay.ValueBoolPointer() != nil {
			data.NoReadvertiseToOverlay = plan.NoReadvertiseToOverlay.ValueBoolPointer()
		}
		if !plan.OtherVrfs.IsNull() && !plan.OtherVrfs.IsUnknown() {
			data.OtherVrfs = mistutils.ListOfStringTerraformToSdk(plan.OtherVrfs)
		}
		if plan.Routed.ValueBoolPointer() != nil {
			data.Routed = plan.Routed.ValueBoolPointer()
		}
		if !plan.SourceNat.IsNull() && !plan.SourceNat.IsUnknown() {
			data.SourceNat = sourceNatTerraformToSdk(ctx, diags, plan.SourceNat)
		}
		if !plan.VpnAccessStaticNat.IsNull() && !plan.VpnAccessStaticNat.IsUnknown() {
			data.StaticNat = staticNatVpnTerraformToSdk(plan.VpnAccessStaticNat)
		}
		if plan.SummarizedSubnet.ValueStringPointer() != nil {
			data.SummarizedSubnet = plan.SummarizedSubnet.ValueStringPointer()
		}
		if plan.SummarizedSubnetToLanBgp.ValueStringPointer() != nil {
			data.SummarizedSubnetToLanBgp = plan.SummarizedSubnetToLanBgp.ValueStringPointer()
		}
		if plan.SummarizedSubnetToLanOspf.ValueStringPointer() != nil {
			data.SummarizedSubnetToLanOspf = plan.SummarizedSubnetToLanOspf.ValueStringPointer()
		}

		dataMap[k] = data
	}
	return dataMap
}
