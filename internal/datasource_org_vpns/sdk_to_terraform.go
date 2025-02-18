package datasource_org_vpns

import (
	"context"
	"math/big"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l *[]models.Vpn, elements *[]attr.Value) diag.Diagnostics {
	var diags diag.Diagnostics

	for _, d := range *l {
		elem := vpnSdkToTerraform(ctx, &diags, &d)
		*elements = append(*elements, elem)
	}

	return diags
}

func vpnSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.Vpn) OrgVpnsValue {
	var createdTime basetypes.NumberValue
	var id basetypes.StringValue
	var modifiedTime basetypes.NumberValue
	var name types.String
	var orgId types.String
	var paths = types.MapNull(PathsValue{}.Type(ctx))

	if d.CreatedTime != nil {
		createdTime = types.NumberValue(big.NewFloat(*d.CreatedTime))
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.ModifiedTime != nil {
		modifiedTime = types.NumberValue(big.NewFloat(*d.ModifiedTime))
	}

	name = types.StringValue(d.Name)

	if d.OrgId != nil {
		orgId = types.StringValue(d.OrgId.String())
	}
	if d.Paths != nil && len(d.Paths) > 0 {
		paths = vpnPathsSdkToTerraform(ctx, diags, d.Paths)
	}

	dataMapValue := map[string]attr.Value{
		"created_time":  createdTime,
		"id":            id,
		"modified_time": modifiedTime,
		"name":          name,
		"org_id":        orgId,
		"paths":         paths,
	}
	data, e := NewOrgVpnsValue(OrgVpnsValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func vpnPathsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.VpnPath) basetypes.MapValue {
	mapAttrValues := make(map[string]attr.Value)
	for k, d := range m {
		var bfdProfile basetypes.StringValue
		var ip basetypes.StringValue
		var pod basetypes.Int64Value

		if d.BfdProfile != nil {
			bfdProfile = types.StringValue(string(*d.BfdProfile))
		}
		if d.Ip != nil {
			ip = types.StringValue(*d.Ip)
		}
		if d.Pod != nil {
			pod = types.Int64Value(int64(*d.Pod))
		}

		dataMapValue := map[string]attr.Value{
			"bfd_profile": bfdProfile,
			"ip":          ip,
			"pod":         pod,
		}
		data, e := NewPathsValue(PathsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		mapAttrValues[k] = data
	}
	stateResult, e := types.MapValueFrom(ctx, PathsValue{}.Type(ctx), mapAttrValues)
	diags.Append(e...)
	return stateResult
}
