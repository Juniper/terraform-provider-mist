package resource_org

import (
	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(data models.Org) (OrgModel, diag.Diagnostics) {
	var state OrgModel
	var diags diag.Diagnostics

	var alarmtemplateId basetypes.StringValue
	var allowMist basetypes.BoolValue
	var orggroupIds = types.ListNull(types.StringType)
	var mspId basetypes.StringValue
	var mspName basetypes.StringValue
	var sessionExpiry basetypes.Int64Value
	state.Id = types.StringValue(data.Id.String())

	if data.AlarmtemplateId.Value() != nil {
		alarmtemplateId = types.StringValue(data.AlarmtemplateId.Value().String())
	}
	if data.AllowMist != nil {
		allowMist = types.BoolValue(*data.AllowMist)
	}
	if data.OrggroupIds != nil {
		orggroupIds = misttransform.ListOfUuidSdkToTerraform(data.OrggroupIds)
	}
	if data.MspId != nil {
		mspId = types.StringValue(data.MspId.String())
	}
	if data.MspName != nil {
		mspName = types.StringPointerValue(data.MspName)
	}
	if data.SessionExpiry != nil {
		sessionExpiry = types.Int64Value(int64(*data.SessionExpiry))
	}

	state.AlarmtemplateId = alarmtemplateId
	state.AllowMist = allowMist
	state.OrggroupIds = orggroupIds
	state.MspId = mspId
	state.MspName = mspName
	state.Name = types.StringValue(data.Name)
	state.SessionExpiry = sessionExpiry

	var items []attr.Value
	for _, item := range data.OrggroupIds {
		tmp := types.StringValue(item.String())
		items = append(items, tmp)
	}
	state.OrggroupIds, _ = types.ListValue(types.StringType, items)
	return state, diags
}
