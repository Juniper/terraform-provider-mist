package resource_org_network

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func destinationNatVpnTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.NetworkVpnAccessDestinationNatProperty {
	data_map := make(map[string]models.NetworkVpnAccessDestinationNatProperty)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(VpnAccessDestinationNatValue)
		data := models.NetworkVpnAccessDestinationNatProperty{}
		data.InternalIp = v_plan.InternalIp.ValueStringPointer()
		data.Name = v_plan.Name.ValueStringPointer()
		data.Port = v_plan.Port.ValueStringPointer()
		data_map[k] = data
	}
	return data_map
}

func staticNatVpnTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.NetworkVpnAccessStaticNatProperty {
	data_map := make(map[string]models.NetworkVpnAccessStaticNatProperty)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(VpnAccessStaticNatValue)
		data := models.NetworkVpnAccessStaticNatProperty{}
		data.InternalIp = v_plan.InternalIp.ValueStringPointer()
		data.Name = v_plan.Name.ValueStringPointer()
		data_map[k] = data
	}
	return data_map
}

func VpnTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.NetworkVpnAccessConfig {
	data_map := make(map[string]models.NetworkVpnAccessConfig)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(VpnAccessValue)

		data := models.NetworkVpnAccessConfig{}
		if plan.AdvertisedSubnet.ValueStringPointer() != nil {
			data.AdvertisedSubnet = plan.AdvertisedSubnet.ValueStringPointer()
		}
		if plan.AllowPing.ValueBoolPointer() != nil {
			data.AllowPing = plan.AllowPing.ValueBoolPointer()
		}
		if !plan.VpnAccessDestinationNat.IsNull() && !plan.VpnAccessDestinationNat.IsUnknown() {
			data.DestinationNat = destinationNatVpnTerraformToSdk(ctx, diags, plan.VpnAccessDestinationNat)
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
			data.OtherVrfs = mist_transform.ListOfStringTerraformToSdk(ctx, plan.OtherVrfs)
		}
		if plan.Routed.ValueBoolPointer() != nil {
			data.Routed = plan.Routed.ValueBoolPointer()
		}
		if !plan.SourceNat.IsNull() && !plan.SourceNat.IsUnknown() {
			data.SourceNat = sourceNatTerraformToSdk(ctx, diags, plan.SourceNat)
		}
		if !plan.VpnAccessStaticNat.IsNull() && !plan.VpnAccessStaticNat.IsUnknown() {
			data.StaticNat = staticNatVpnTerraformToSdk(ctx, diags, plan.VpnAccessStaticNat)
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

		data_map[k] = data
	}
	return data_map
}
