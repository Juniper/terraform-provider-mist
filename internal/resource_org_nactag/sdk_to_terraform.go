package resource_org_nactag

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func SdkToTerraform(data *models.NacTag) (OrgNactagModel, diag.Diagnostics) {
	var state OrgNactagModel
	var diags diag.Diagnostics

	var allowUsermacOverride types.Bool
	var egressVlanNames = types.ListNull(types.StringType)
	var gbpTag types.Int64
	var id types.String
	var match types.String
	var matchAll types.Bool
	var name types.String
	var orgId types.String
	var radiusAttrs = types.ListNull(types.StringType)
	var radiusGroup types.String
	var radiusVendorAttrs = types.ListNull(types.StringType)
	var sessionTimeout types.Int64
	var typeNactag types.String
	var values = types.ListNull(types.StringType)
	var vlan types.String

	if data.AllowUsermacOverride != nil {
		allowUsermacOverride = types.BoolValue(*data.AllowUsermacOverride)
	}
	if data.EgressVlanNames != nil {
		egressVlanNames = misttransform.ListOfStringSdkToTerraform(data.EgressVlanNames)
	}
	if data.GbpTag != nil {
		gbpTag = types.Int64Value(int64(*data.GbpTag))
	}
	if data.Id != nil {
		id = types.StringValue(data.Id.String())
	}
	if data.Match != nil {
		match = types.StringValue(string(*data.Match))
	}
	if data.MatchAll != nil {
		matchAll = types.BoolValue(*data.MatchAll)
	}

	name = types.StringValue(data.Name)

	if data.OrgId != nil {
		orgId = types.StringValue(data.OrgId.String())
	}
	if data.RadiusAttrs != nil {
		radiusAttrs = misttransform.ListOfStringSdkToTerraform(data.RadiusAttrs)
	}
	if data.RadiusGroup != nil {
		radiusGroup = types.StringValue(*data.RadiusGroup)
	}
	if data.RadiusVendorAttrs != nil {
		radiusVendorAttrs = misttransform.ListOfStringSdkToTerraform(data.RadiusVendorAttrs)
	}
	if data.SessionTimeout != nil {
		sessionTimeout = types.Int64Value(int64(*data.SessionTimeout))
	}
	typeNactag = types.StringValue(string(data.Type))
	if data.Values != nil {
		values = misttransform.ListOfStringSdkToTerraform(data.Values)
	}
	if data.Vlan != nil {
		vlan = types.StringValue(*data.Vlan)
	}

	state.AllowUsermacOverride = allowUsermacOverride
	state.EgressVlanNames = egressVlanNames
	state.GbpTag = gbpTag
	state.Id = id
	state.Match = match
	state.MatchAll = matchAll
	state.Name = name
	state.OrgId = orgId
	state.RadiusAttrs = radiusAttrs
	state.RadiusGroup = radiusGroup
	state.RadiusVendorAttrs = radiusVendorAttrs
	state.SessionTimeout = sessionTimeout
	state.Type = typeNactag
	state.Values = values
	state.Vlan = vlan

	return state, diags
}
