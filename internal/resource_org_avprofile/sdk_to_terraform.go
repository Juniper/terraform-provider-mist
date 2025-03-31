package resource_org_avprofile

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/hashicorp/terraform-plugin-framework/attr"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(_ context.Context, data *models.Avprofile) (OrgAvprofileModel, diag.Diagnostics) {
	var state OrgAvprofileModel
	var diags diag.Diagnostics

	var fallbackAction basetypes.StringValue
	var id basetypes.StringValue
	var maxFilesize basetypes.Int64Value
	var mimeWhitelist = types.ListValueMust(types.StringType, []attr.Value{})
	var name basetypes.StringValue
	var orgId basetypes.StringValue
	var protocols = types.ListNull(types.StringType)
	var urlWhitelist = types.ListValueMust(types.StringType, []attr.Value{})

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
		mimeWhitelist = mistutils.ListOfStringSdkToTerraform(data.MimeWhitelist)
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
		urlWhitelist = mistutils.ListOfStringSdkToTerraform(data.UrlWhitelist)
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
