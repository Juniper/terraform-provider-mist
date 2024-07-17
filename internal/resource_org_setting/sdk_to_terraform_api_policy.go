package resource_org_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func apiPolicySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.OrgSettingApiPolicy) ApiPolicyValue {

	var no_reveal basetypes.BoolValue

	if d.NoReveal != nil {
		no_reveal = types.BoolValue(*d.NoReveal)
	}

	data_map_attr_type := ApiPolicyValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"no_reveal": no_reveal,
	}
	data, e := NewApiPolicyValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
