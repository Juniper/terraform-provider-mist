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

	var noReveal basetypes.BoolValue

	if d.NoReveal != nil {
		noReveal = types.BoolValue(*d.NoReveal)
	}

	dataMapValue := map[string]attr.Value{
		"no_reveal": noReveal,
	}
	data, e := NewApiPolicyValue(ApiPolicyValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
