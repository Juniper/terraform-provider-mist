package datasource_org_avprofiles

import (
	"context"
	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l *[]models.Avprofile, elements *[]attr.Value) diag.Diagnostics {
	var diags diag.Diagnostics

	for _, d := range *l {
		elem := servicepolicieSdkToTerraform(ctx, &diags, &d)
		*elements = append(*elements, elem)
	}

	return diags
}

func servicepolicieSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.Avprofile) OrgAvprofilesValue {

	var createdTime basetypes.Float64Value
	var fallbackAction basetypes.StringValue
	var id basetypes.StringValue
	var maxFilesize basetypes.Int64Value
	var mimeWhitelist = types.ListNull(types.StringType)
	var modifiedTime basetypes.Float64Value
	var name basetypes.StringValue
	var orgId basetypes.StringValue
	var protocols = types.ListNull(types.StringType)
	var urlWhitelist = types.ListNull(types.StringType)

	if d.CreatedTime != nil {
		createdTime = types.Float64Value(*d.CreatedTime)
	}
	if d.FallbackAction != nil {
		fallbackAction = types.StringValue(string(*d.FallbackAction))
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.MaxFilesize != nil {
		maxFilesize = types.Int64Value(int64(*d.MaxFilesize))
	}
	if d.MimeWhitelist != nil {
		mimeWhitelist = misttransform.ListOfStringSdkToTerraform(d.MimeWhitelist)
	}
	if d.ModifiedTime != nil {
		modifiedTime = types.Float64Value(*d.ModifiedTime)
	}

	name = types.StringValue(d.Name)

	orgId = types.StringValue(d.OrgId.String())

	if d.Protocols != nil {
		var items []attr.Value
		for _, v := range d.Protocols {
			items = append(items, types.StringValue(string(v)))
		}
		tmp, e := types.ListValue(basetypes.StringType{}, items)
		if e != nil {
			diags.Append(e...)
		} else {
			protocols = tmp
		}
	}
	if d.UrlWhitelist != nil {
		urlWhitelist = misttransform.ListOfStringSdkToTerraform(d.UrlWhitelist)
	}

	dataMapValue := map[string]attr.Value{
		"fallback_action": fallbackAction,
		"created_time":    createdTime,
		"id":              id,
		"max_filesize":    maxFilesize,
		"mime_whitelist":  mimeWhitelist,
		"modified_time":   modifiedTime,
		"name":            name,
		"org_id":          orgId,
		"protocols":       protocols,
		"url_whitelist":   urlWhitelist,
	}
	data, e := NewOrgAvprofilesValue(OrgAvprofilesValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
