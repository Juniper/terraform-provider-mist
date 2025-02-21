package resource_org_avprofile

import (
	"context"
	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/hashicorp/terraform-plugin-framework/attr"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, data *models.Avprofile) (OrgAvprofileModel, diag.Diagnostics) {
	var state OrgAvprofileModel
	var diags diag.Diagnostics

	var fallbackAction basetypes.StringValue
	var id basetypes.StringValue
	var maxFilesize basetypes.Int64Value
	var mimeWhitelist basetypes.ListValue
	var name basetypes.StringValue
	var orgId basetypes.StringValue
	var protocols basetypes.ListValue
	var urlWhitelist basetypes.ListValue

	if data.FallbackAction != nil {
		fallbackAction = types.StringValue(string(*data.FallbackAction))
	}
	if data.Id != nil {
		id = types.StringValue(data.Id.String())
	}
	if data.MaxFilesize != nil {
		maxFilesize = types.Int64Value(int64(*data.MaxFilesize))
	}
	if data.MimeWhitelist != nil {
		mimeWhitelist = misttransform.ListOfStringSdkToTerraform(data.MimeWhitelist)
	}

	name = types.StringValue(data.Name)

	orgId = types.StringValue(data.OrgId.String())

	if data.Protocols != nil {
		var items []attr.Value
		for _, v := range data.Protocols {
			items = append(items, types.StringValue(string(v)))
		}
		tmp, e := types.ListValue(basetypes.StringType{}, items)
		if e != nil {
			diags.Append(e...)
		} else {
			protocols = tmp
		}
	}
	if data.UrlWhitelist != nil {
		urlWhitelist = misttransform.ListOfStringSdkToTerraform(data.UrlWhitelist)
	}

	state.FallbackAction = fallbackAction
	state.Id = id
	state.MaxFilesize = maxFilesize
	state.MimeWhitelist = mimeWhitelist
	state.Name = name
	state.OrgId = orgId
	state.Protocols = protocols
	state.UrlWhitelist = urlWhitelist

	return state, diags
}
