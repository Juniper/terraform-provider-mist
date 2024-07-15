package resource_org_network

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func VpnTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.NetworkVpnAccessConfig {
	data_map := make(map[string]models.NetworkVpnAccessConfig)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(VpnAccessValue)

		data := models.NetworkVpnAccessConfig{}
		data.AdvertisedSubnet = v_plan.AdvertisedSubnet.ValueStringPointer()
		data.AllowPing = v_plan.AllowPing.ValueBoolPointer()
		data.DestinationNat = destinationNatTerraformToSdk(ctx, diags, v_plan.DestinationNat)
		data.NatPool = v_plan.NatPool.ValueStringPointer()
		data.NoReadvertiseToLanBgp = v_plan.NoReadvertiseToLanBgp.ValueBoolPointer()
		data.NoReadvertiseToLanOspf = v_plan.NoReadvertiseToLanOspf.ValueBoolPointer()
		data.NoReadvertiseToOverlay = v_plan.NoReadvertiseToOverlay.ValueBoolPointer()
		data.OtherVrfs = mist_transform.ListOfStringTerraformToSdk(ctx, v_plan.OtherVrfs)
		data.Routed = v_plan.Routed.ValueBoolPointer()
		data.SourceNat = sourceNatTerraformToSdk(ctx, diags, v_plan.SourceNat)
		data.StaticNat = staticNatTerraformToSdk(ctx, diags, v_plan.StaticNat)
		data.SummarizedSubnet = v_plan.SummarizedSubnet.ValueStringPointer()
		data.SummarizedSubnetToLanBgp = v_plan.SummarizedSubnetToLanBgp.ValueStringPointer()
		data.SummarizedSubnetToLanOspf = v_plan.SummarizedSubnetToLanOspf.ValueStringPointer()
		data_map[k] = data
	}
	return data_map
}
