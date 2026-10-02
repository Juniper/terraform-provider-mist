package resource_org_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func cacertsConfigsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.OrgSettingCacertsConfig) basetypes.ListValue {
	var dataList []CacertsConfigsValue
	for _, d := range l {
		var crlEnabled basetypes.BoolValue
		var crlUrl basetypes.StringValue
		var name basetypes.StringValue
		var ocspEnabled basetypes.BoolValue
		var ocspUrl basetypes.StringValue

		cert := types.StringValue(d.Cert)
		if d.CrlEnabled != nil {
			crlEnabled = types.BoolValue(*d.CrlEnabled)
		}
		if d.CrlUrl != nil {
			crlUrl = types.StringValue(*d.CrlUrl)
		}
		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}
		if d.OcspEnabled != nil {
			ocspEnabled = types.BoolValue(*d.OcspEnabled)
		}
		if d.OcspUrl != nil {
			ocspUrl = types.StringValue(*d.OcspUrl)
		}

		dataMapValue := map[string]attr.Value{
			"cert":         cert,
			"crl_enabled":  crlEnabled,
			"crl_url":      crlUrl,
			"name":         name,
			"ocsp_enabled": ocspEnabled,
			"ocsp_url":     ocspUrl,
		}
		data, e := NewCacertsConfigsValue(CacertsConfigsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, CacertsConfigsValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}
