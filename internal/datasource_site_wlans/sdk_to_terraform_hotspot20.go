package datasource_site_wlans

import (
	"context"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func hotspot20SdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.WlanHotspot20) basetypes.ObjectValue {
	var domainName = misttransform.ListOfStringSdkToTerraformEmpty()
	var enabled basetypes.BoolValue
	var naiRealms = misttransform.ListOfStringSdkToTerraformEmpty()
	var operators = misttransform.ListOfStringSdkToTerraformEmpty()
	var rcoi = misttransform.ListOfStringSdkToTerraformEmpty()
	var venueName basetypes.StringValue

	if d != nil && d.DomainName != nil {
		domainName = misttransform.ListOfStringSdkToTerraform(d.DomainName)
	}
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.NaiRealms != nil {
		naiRealms = misttransform.ListOfStringSdkToTerraform(d.NaiRealms)
	}
	if d != nil && d.Operators != nil {
		var operatorsList []attr.Value
		for _, v := range d.Operators {
			operatorsList = append(operatorsList, types.StringValue(string(v)))
		}
		operators = types.ListValueMust(basetypes.StringType{}, operatorsList)
	}
	if d != nil && d.Rcoi != nil {
		rcoi = misttransform.ListOfStringSdkToTerraform(d.Rcoi)
	}
	if d != nil && d.VenueName != nil {
		venueName = types.StringValue(*d.VenueName)
	}

	dataMapValue := map[string]attr.Value{
		"domain_name": domainName,
		"enabled":     enabled,
		"nai_realms":  naiRealms,
		"operators":   operators,
		"rcoi":        rcoi,
		"venue_name":  venueName,
	}
	data, e := basetypes.NewObjectValue(Hotspot20Value{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
