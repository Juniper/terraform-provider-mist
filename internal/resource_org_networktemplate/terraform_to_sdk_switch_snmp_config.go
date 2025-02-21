package resource_org_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

// //////////////////////////////////
// ////////// CLIENTS
func snmpConfigClientListTerraformToSdk(d basetypes.ListValue) []models.SnmpConfigClientList {
	var dataList []models.SnmpConfigClientList
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(ClientListValue)
		data := models.SnmpConfigClientList{}
		if !plan.Clients.IsNull() && !plan.Clients.IsUnknown() {
			data.Clients = misttransform.ListOfStringTerraformToSdk(plan.Clients)
		}
		if plan.ClientListName.ValueStringPointer() != nil {
			data.ClientListName = plan.ClientListName.ValueStringPointer()
		}

		dataList = append(dataList, data)
	}

	return dataList
}

// //////////////////////////////////
// ////////// TRAPS
func snmpConfigTrapGroupsTerraformToSdk(d basetypes.ListValue) []models.SnmpConfigTrapGroup {
	var dataList []models.SnmpConfigTrapGroup
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(TrapGroupsValue)
		data := models.SnmpConfigTrapGroup{}
		if !plan.Categories.IsNull() && !plan.Categories.IsUnknown() {
			data.Categories = misttransform.ListOfStringTerraformToSdk(plan.Categories)
		}
		if plan.GroupName.ValueStringPointer() != nil {
			data.GroupName = plan.GroupName.ValueStringPointer()
		}
		if !plan.Targets.IsNull() && !plan.Targets.IsUnknown() {
			data.Targets = misttransform.ListOfStringTerraformToSdk(plan.Targets)
		}
		if plan.Version.ValueStringPointer() != nil {
			data.Version = (*models.SnmpConfigTrapVerionEnum)(plan.Version.ValueStringPointer())
		}

		dataList = append(dataList, data)
	}

	return dataList
}

// //////////////////////////////////
// ////////// V2c
func snmpConfigV2cTerraformToSdk(d basetypes.ListValue) []models.SnmpConfigV2cConfig {
	var dataList []models.SnmpConfigV2cConfig
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(V2cConfigValue)
		data := models.SnmpConfigV2cConfig{}
		if plan.Authorization.ValueStringPointer() != nil {
			data.Authorization = plan.Authorization.ValueStringPointer()
		}
		if plan.ClientListName.ValueStringPointer() != nil {
			data.ClientListName = plan.ClientListName.ValueStringPointer()
		}
		if plan.CommunityName.ValueStringPointer() != nil {
			data.CommunityName = plan.CommunityName.ValueStringPointer()
		}
		if plan.View.ValueStringPointer() != nil {
			data.View = plan.View.ValueStringPointer()
		}

		dataList = append(dataList, data)
	}

	return dataList
}

// //////////////////////////////////
// ////////// V3
// V3 NOTIFY
func snmpConfigV3NotifyTerraformToSdk(d basetypes.ListValue) []models.Snmpv3ConfigNotifyItems {
	var dataList []models.Snmpv3ConfigNotifyItems
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(NotifyValue)
		data := models.Snmpv3ConfigNotifyItems{}
		if plan.Name.ValueStringPointer() != nil {
			data.Name = plan.Name.ValueStringPointer()
		}
		if plan.Tag.ValueStringPointer() != nil {
			data.Tag = plan.Tag.ValueStringPointer()
		}
		if plan.NotifyType.ValueStringPointer() != nil {
			data.Type = (*models.Snmpv3ConfigNotifyTypeEnum)(plan.NotifyType.ValueStringPointer())
		}

		dataList = append(dataList, data)
	}

	return dataList
}

func snmpConfigV3NotifyFilterContentTerraformToSdk(d basetypes.ListValue) []models.Snmpv3ConfigNotifyFilterItemContent {
	var dataList []models.Snmpv3ConfigNotifyFilterItemContent
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(Snmpv3ContentsValue)
		data := models.Snmpv3ConfigNotifyFilterItemContent{}

		if plan.Include.ValueBoolPointer() != nil {
			data.Include = plan.Include.ValueBoolPointer()
		}
		if plan.Oid.ValueStringPointer() != nil {
			data.Oid = plan.Oid.ValueStringPointer()
		}

		dataList = append(dataList, data)
	}

	return dataList
}

func snmpConfigV3NotifyFilterTerraformToSdk(d basetypes.ListValue) []models.Snmpv3ConfigNotifyFilterItem {
	var dataList []models.Snmpv3ConfigNotifyFilterItem
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(NotifyFilterValue)
		data := models.Snmpv3ConfigNotifyFilterItem{}

		if !plan.Snmpv3Contents.IsNull() && !plan.Snmpv3Contents.IsUnknown() {
			data.Contents = snmpConfigV3NotifyFilterContentTerraformToSdk(plan.Snmpv3Contents)
		}
		if plan.ProfileName.ValueStringPointer() != nil {
			data.ProfileName = plan.ProfileName.ValueStringPointer()
		}

		dataList = append(dataList, data)
	}

	return dataList
}

// V3 TARGETS
func snmpConfigV3TargetAddressTerraformToSdk(d basetypes.ListValue) []models.Snmpv3ConfigTargetAddressItem {
	var dataList []models.Snmpv3ConfigTargetAddressItem
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(TargetAddressValue)
		data := models.Snmpv3ConfigTargetAddressItem{}

		if plan.Address.ValueStringPointer() != nil {
			data.Address = plan.Address.ValueStringPointer()
		}
		if plan.AddressMask.ValueStringPointer() != nil {
			data.AddressMask = plan.AddressMask.ValueStringPointer()
		}
		if plan.Port.ValueStringPointer() != nil {
			data.Port = models.NewOptional(plan.Port.ValueStringPointer())
		}
		if plan.TagList.ValueStringPointer() != nil {
			data.TagList = plan.TagList.ValueStringPointer()
		}
		if plan.TargetAddressName.ValueStringPointer() != nil {
			data.TargetAddressName = plan.TargetAddressName.ValueStringPointer()
		}
		if plan.TargetParameters.ValueStringPointer() != nil {
			data.TargetParameters = plan.TargetParameters.ValueStringPointer()
		}

		dataList = append(dataList, data)
	}

	return dataList
}

func snmpConfigV3TargetParametersTerraformToSdk(d basetypes.ListValue) []models.Snmpv3ConfigTargetParam {
	var dataList []models.Snmpv3ConfigTargetParam
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(TargetParametersValue)
		data := models.Snmpv3ConfigTargetParam{}

		if plan.MessageProcessingModel.ValueStringPointer() != nil {
			data.MessageProcessingModel = (*models.Snmpv3ConfigTargetParamMessProcessModelEnum)(plan.MessageProcessingModel.ValueStringPointer())
		}
		if plan.Name.ValueStringPointer() != nil {
			data.Name = plan.Name.ValueStringPointer()
		}
		if plan.NotifyFilter.ValueStringPointer() != nil {
			data.NotifyFilter = plan.NotifyFilter.ValueStringPointer()
		}
		if plan.SecurityLevel.ValueStringPointer() != nil {
			data.SecurityLevel = (*models.Snmpv3ConfigTargetParamSecurityLevelEnum)(plan.SecurityLevel.ValueStringPointer())
		}
		if plan.SecurityModel.ValueStringPointer() != nil {
			data.SecurityModel = (*models.Snmpv3ConfigTargetParamSecurityModelEnum)(plan.SecurityModel.ValueStringPointer())
		}
		if plan.SecurityName.ValueStringPointer() != nil {
			data.SecurityName = plan.SecurityName.ValueStringPointer()
		}

		dataList = append(dataList, data)
	}

	return dataList
}

// V3 USM
func snmpConfigV3UsmUsersTerraformToSdk(d basetypes.ListValue) []models.SnmpUsmpUser {
	var dataList []models.SnmpUsmpUser
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(Snmpv3UsersValue)
		data := models.SnmpUsmpUser{}

		if plan.AuthenticationPassword.ValueStringPointer() != nil {
			data.AuthenticationPassword = plan.AuthenticationPassword.ValueStringPointer()
		}
		if plan.AuthenticationType.ValueStringPointer() != nil {
			data.AuthenticationType = (*models.SnmpUsmpUserAuthenticationTypeEnum)(plan.AuthenticationType.ValueStringPointer())
		}
		if plan.EncryptionPassword.ValueStringPointer() != nil {
			data.EncryptionPassword = plan.EncryptionPassword.ValueStringPointer()
		}
		if plan.EncryptionType.ValueStringPointer() != nil {
			data.EncryptionType = (*models.SnmpUsmpUserEncryptionTypeEnum)(plan.EncryptionType.ValueStringPointer())
		}
		if plan.Name.ValueStringPointer() != nil {
			data.Name = plan.Name.ValueStringPointer()
		}

		dataList = append(dataList, data)
	}

	return dataList
}

func snmpConfigV3UsmTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.SnmpUsm {
	var dataList []models.SnmpUsm
	for _, v := range d.Elements() {

		var vInterface interface{} = v
		plan := vInterface.(UsmValue)
		data := models.SnmpUsm{}

		if plan.EngineType.ValueStringPointer() != nil {
			data.EngineType = (*models.SnmpUsmEngineTypeEnum)(plan.EngineType.ValueStringPointer())
		}
		if plan.RemoteEngineId.ValueStringPointer() != nil {
			data.RemoteEngineId = plan.RemoteEngineId.ValueStringPointer()
		}
		if !plan.Snmpv3Users.IsNull() && !plan.Snmpv3Users.IsUnknown() {
			data.Users = snmpConfigV3UsmUsersTerraformToSdk(plan.Snmpv3Users)
		}
		dataList = append(dataList, data)
	}

	return dataList
}

// V3 VACM ACCESS
func snmpConfigV3VacmAccessPrefixTerraformToSdk(d basetypes.ListValue) []models.SnmpVacmAccessItemPrefixListItem {
	var dataList []models.SnmpVacmAccessItemPrefixListItem
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(PrefixListValue)
		data := models.SnmpVacmAccessItemPrefixListItem{}

		if plan.ContextPrefix.ValueStringPointer() != nil {
			data.ContextPrefix = plan.ContextPrefix.ValueStringPointer()
		}
		if plan.NotifyView.ValueStringPointer() != nil {
			data.NotifyView = plan.NotifyView.ValueStringPointer()
		}
		if plan.ReadView.ValueStringPointer() != nil {
			data.ReadView = plan.ReadView.ValueStringPointer()
		}
		if plan.SecurityLevel.ValueStringPointer() != nil {
			data.SecurityLevel = (*models.SnmpVacmAccessItemPrefixListItemLevelEnum)(plan.SecurityLevel.ValueStringPointer())
		}
		if plan.SecurityModel.ValueStringPointer() != nil {
			data.SecurityModel = (*models.SnmpVacmAccessItemPrefixListItemModelEnum)(plan.SecurityModel.ValueStringPointer())
		}
		if plan.PrefixListType.ValueStringPointer() != nil {
			data.Type = (*models.SnmpVacmAccessItemTypeEnum)(plan.PrefixListType.ValueStringPointer())
		}
		if plan.WriteView.ValueStringPointer() != nil {
			data.WriteView = plan.WriteView.ValueStringPointer()
		}

		dataList = append(dataList, data)
	}

	return dataList
}
func snmpConfigV3VacmAccessTerraformToSdk(d basetypes.ListValue) []models.SnmpVacmAccessItem {
	var dataList []models.SnmpVacmAccessItem
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(AccessValue)
		data := models.SnmpVacmAccessItem{}

		if plan.GroupName.ValueStringPointer() != nil {
			data.GroupName = plan.GroupName.ValueStringPointer()
		}
		if !plan.PrefixList.IsNull() && !plan.PrefixList.IsUnknown() {
			data.PrefixList = snmpConfigV3VacmAccessPrefixTerraformToSdk(plan.PrefixList)
		}

		dataList = append(dataList, data)
	}

	return dataList
}

// V3 VACM SEC TO GROUP
func snmpConfigV3VacmSecurityToGroupContentTerraformToSdk(d basetypes.ListValue) []models.SnmpVacmSecurityToGroupContentItem {
	var dataList []models.SnmpVacmSecurityToGroupContentItem
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(Snmpv3VacmContentValue)
		data := models.SnmpVacmSecurityToGroupContentItem{}

		if plan.Group.ValueStringPointer() != nil {
			data.Group = plan.Group.ValueStringPointer()
		}
		if plan.SecurityName.ValueStringPointer() != nil {
			data.SecurityName = plan.SecurityName.ValueStringPointer()
		}

		dataList = append(dataList, data)
	}

	return dataList
}
func snmpConfigV3VacmSecurityToGroupTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.SnmpVacmSecurityToGroup {
	data := models.SnmpVacmSecurityToGroup{}
	if !d.IsNull() && !d.IsUnknown() {
		plan, e := NewSecurityToGroupValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if plan.SecurityModel.ValueStringPointer() != nil {
				data.SecurityModel = (*models.SnmpVacmSecurityModelEnum)(plan.SecurityModel.ValueStringPointer())
			}
			if !plan.Snmpv3VacmContent.IsNull() && !plan.Snmpv3VacmContent.IsUnknown() {
				data.Content = snmpConfigV3VacmSecurityToGroupContentTerraformToSdk(plan.Snmpv3VacmContent)
			}
		}
	}
	return &data
}
func snmpConfigV3VacmTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.SnmpVacm {
	data := models.SnmpVacm{}
	if !d.IsNull() && !d.IsUnknown() {
		plan, e := NewVacmValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if !plan.Access.IsNull() && !plan.Access.IsUnknown() {
				data.Access = snmpConfigV3VacmAccessTerraformToSdk(plan.Access)
			}
			if !plan.SecurityToGroup.IsNull() && !plan.SecurityToGroup.IsUnknown() {
				data.SecurityToGroup = snmpConfigV3VacmSecurityToGroupTerraformToSdk(ctx, diags, plan.SecurityToGroup)
			}
		}
	}
	return &data
}

// V3 MAIN
func snmpConfigV3TerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.Snmpv3Config {
	data := models.Snmpv3Config{}
	if !d.IsNull() && !d.IsUnknown() {
		plan, e := NewV3ConfigValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if !plan.Notify.IsNull() && !plan.Notify.IsUnknown() {
				data.Notify = snmpConfigV3NotifyTerraformToSdk(plan.Notify)
			}
			if !plan.NotifyFilter.IsNull() && !plan.NotifyFilter.IsUnknown() {
				data.NotifyFilter = snmpConfigV3NotifyFilterTerraformToSdk(plan.NotifyFilter)
			}
			if !plan.TargetAddress.IsNull() && !plan.TargetAddress.IsUnknown() {
				data.TargetAddress = snmpConfigV3TargetAddressTerraformToSdk(plan.TargetAddress)
			}
			if !plan.TargetParameters.IsNull() && !plan.TargetParameters.IsUnknown() {
				data.TargetParameters = snmpConfigV3TargetParametersTerraformToSdk(plan.TargetParameters)
			}
			if !plan.Usm.IsNull() && !plan.Usm.IsUnknown() {
				data.Usm = snmpConfigV3UsmTerraformToSdk(ctx, diags, plan.Usm)
			}
			if !plan.Vacm.IsNull() && !plan.Vacm.IsUnknown() {
				data.Vacm = snmpConfigV3VacmTerraformToSdk(ctx, diags, plan.Vacm)
			}
		}
	}
	return &data
}

// //////////////////////////////////////////////
// ////////// VIEWS
func snmpConfigViewsTerraformToSdk(d basetypes.ListValue) []models.SnmpConfigView {
	var dataList []models.SnmpConfigView
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(ViewsValue)
		data := models.SnmpConfigView{}

		if plan.Include.ValueBoolPointer() != nil {
			data.Include = plan.Include.ValueBoolPointer()
		}
		if plan.Oid.ValueStringPointer() != nil {
			data.Oid = plan.Oid.ValueStringPointer()
		}
		if plan.ViewName.ValueStringPointer() != nil {
			data.ViewName = plan.ViewName.ValueStringPointer()
		}

		dataList = append(dataList, data)
	}

	return dataList
}

// //////////////////////////////////////////////
// ////////// MAIN
func snmpConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SnmpConfigValue) *models.SnmpConfig {

	data := models.SnmpConfig{}
	if !d.ClientList.IsNull() && !d.ClientList.IsUnknown() {
		data.ClientList = snmpConfigClientListTerraformToSdk(d.ClientList)
	}
	if d.Contact.ValueStringPointer() != nil {
		data.Contact = d.Contact.ValueStringPointer()
	}
	if d.Description.ValueStringPointer() != nil {
		data.Description = d.Description.ValueStringPointer()
	}
	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if d.EngineId.ValueStringPointer() != nil {
		data.EngineId = d.EngineId.ValueStringPointer()
	}
	if d.Location.ValueStringPointer() != nil {
		data.Location = d.Location.ValueStringPointer()
	}
	if d.Name.ValueStringPointer() != nil {
		data.Name = d.Name.ValueStringPointer()
	}
	if d.Network.ValueStringPointer() != nil {
		data.Network = d.Network.ValueStringPointer()
	}
	if !d.TrapGroups.IsNull() && !d.TrapGroups.IsUnknown() {
		data.TrapGroups = snmpConfigTrapGroupsTerraformToSdk(d.TrapGroups)
	}
	if !d.V2cConfig.IsNull() && !d.V2cConfig.IsUnknown() {
		data.V2cConfig = snmpConfigV2cTerraformToSdk(d.V2cConfig)
	}
	if !d.V3Config.IsNull() && !d.V3Config.IsUnknown() {
		data.V3Config = snmpConfigV3TerraformToSdk(ctx, diags, d.V3Config)
	}
	if !d.Views.IsNull() && !d.Views.IsUnknown() {
		data.Views = snmpConfigViewsTerraformToSdk(d.Views)
	}

	return &data
}
