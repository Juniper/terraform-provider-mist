package resource_org_networktemplate

import (
	"context"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// //////////////////////////////////
// ////////// CLIENTS
func snmpClientListSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.SnmpConfigClientList) basetypes.ListValue {
	var dataList []ClientListValue
	for _, d := range l {

		var clientListName basetypes.StringValue
		var clients = misttransform.ListOfStringSdkToTerraformEmpty()

		if d.ClientListName != nil {
			clientListName = types.StringValue(*d.ClientListName)
		}
		if d.Clients != nil {
			clients = misttransform.ListOfStringSdkToTerraform(d.Clients)
		}

		dataMapValue := map[string]attr.Value{
			"client_list_name": clientListName,
			"clients":          clients,
		}
		data, e := NewClientListValue(ClientListValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, ClientListValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}

// //////////////////////////////////
// ////////// TRAP GROUPS
func snmpTrapGroupsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.SnmpConfigTrapGroup) basetypes.ListValue {
	var dataList []TrapGroupsValue
	for _, d := range l {
		var categories = misttransform.ListOfStringSdkToTerraformEmpty()
		var groupName basetypes.StringValue
		var targets = misttransform.ListOfStringSdkToTerraformEmpty()
		var version basetypes.StringValue

		if d.Categories != nil {
			categories = misttransform.ListOfStringSdkToTerraform(d.Categories)
		}
		if d.GroupName != nil {
			groupName = types.StringValue(*d.GroupName)
		}
		if d.Targets != nil {
			targets = misttransform.ListOfStringSdkToTerraform(d.Targets)
		}
		if d.Version != nil {
			version = types.StringValue(string(*d.Version))
		}

		dataMapValue := map[string]attr.Value{
			"categories": categories,
			"group_name": groupName,
			"targets":    targets,
			"version":    version,
		}
		data, e := NewTrapGroupsValue(TrapGroupsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, TrapGroupsValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}

// //////////////////////////////////
// ////////// V2
func snmpV2cSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.SnmpConfigV2cConfig) basetypes.ListValue {
	var dataList []V2cConfigValue
	for _, d := range l {
		var authorization basetypes.StringValue
		var clientListName basetypes.StringValue
		var communityName basetypes.StringValue
		var view basetypes.StringValue

		if d.Authorization != nil {
			authorization = types.StringValue(*d.Authorization)
		}
		if d.ClientListName != nil {
			clientListName = types.StringValue(*d.ClientListName)
		}
		if d.CommunityName != nil {
			communityName = types.StringValue(*d.CommunityName)
		}
		if d.View != nil {
			view = types.StringValue(*d.View)
		}

		dataMapValue := map[string]attr.Value{
			"authorization":    authorization,
			"client_list_name": clientListName,
			"community_name":   communityName,
			"view":             view,
		}
		data, e := NewV2cConfigValue(V2cConfigValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, V2cConfigValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}

// //////////////////////////////////
// ////////// V3
// NOTIFY
func snmpV3NotifySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.Snmpv3ConfigNotifyItems) basetypes.ListValue {
	var dataList []NotifyValue
	for _, d := range l {
		var name basetypes.StringValue
		var tag basetypes.StringValue
		var notifyType basetypes.StringValue

		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}
		if d.Tag != nil {
			tag = types.StringValue(*d.Tag)
		}
		if d.Type != nil {
			notifyType = types.StringValue(string(*d.Type))
		}

		dataMapValue := map[string]attr.Value{
			"name": name,
			"tag":  tag,
			"type": notifyType,
		}
		data, e := NewNotifyValue(NotifyValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, NotifyValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}

// NOTIFY Filter
func snmpV3NotifyFilterContentSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.Snmpv3ConfigNotifyFilterItemContent) basetypes.ListValue {
	var dataList []Snmpv3ContentsValue
	for _, d := range l {

		var include basetypes.BoolValue
		var oid basetypes.StringValue

		if d.Include != nil {
			include = types.BoolValue(*d.Include)
		}
		if d.Oid != nil {
			oid = types.StringValue(*d.Oid)
		}

		dataMapValue := map[string]attr.Value{
			"include": include,
			"oid":     oid,
		}
		data, e := NewSnmpv3ContentsValue(Snmpv3ContentsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, Snmpv3ContentsValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}
func snmpV3NotifyFilterSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.Snmpv3ConfigNotifyFilterItem) basetypes.ListValue {
	var dataList []NotifyFilterValue
	for _, d := range l {
		var contents = snmpV3NotifyFilterContentSdkToTerraform(ctx, diags, d.Contents)
		var profileName basetypes.StringValue

		if d.ProfileName != nil {
			profileName = types.StringValue(*d.ProfileName)
		}

		dataMapValue := map[string]attr.Value{
			"contents":     contents,
			"profile_name": profileName,
		}
		data, e := NewNotifyFilterValue(NotifyFilterValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, NotifyFilterValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}

// TARGET ADDRESS
func snmpV3TargetAddressSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.Snmpv3ConfigTargetAddressItem) basetypes.ListValue {
	var dataList []TargetAddressValue
	for _, d := range l {
		var address basetypes.StringValue
		var addressMask basetypes.StringValue
		var port basetypes.Int64Value
		var tagList basetypes.StringValue
		var targetAddressName basetypes.StringValue
		var targetParameters basetypes.StringValue

		if d.Address != nil {
			address = types.StringValue(*d.Address)
		}
		if d.AddressMask != nil {
			addressMask = types.StringValue(*d.AddressMask)
		}
		if d.Port != nil {
			port = types.Int64Value(int64(*d.Port))
		}
		if d.TagList != nil {
			tagList = types.StringValue(*d.TagList)
		}
		if d.TargetAddressName != nil {
			targetAddressName = types.StringValue(*d.TargetAddressName)
		}
		if d.TargetParameters != nil {
			targetParameters = types.StringValue(*d.TargetParameters)
		}

		dataMapValue := map[string]attr.Value{
			"address":             address,
			"address_mask":        addressMask,
			"port":                port,
			"tag_list":            tagList,
			"target_address_name": targetAddressName,
			"target_parameters":   targetParameters,
		}
		data, e := NewTargetAddressValue(TargetAddressValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, TargetAddressValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}

// TARGET PARAMETERS
func snmpV3TargetParametersSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.Snmpv3ConfigTargetParam) basetypes.ListValue {
	var dataList []TargetParametersValue
	for _, d := range l {
		var messageProcessingModel basetypes.StringValue
		var name basetypes.StringValue
		var notifyFilter basetypes.StringValue
		var securityLevel basetypes.StringValue
		var securityModel basetypes.StringValue
		var securityName basetypes.StringValue

		if d.MessageProcessingModel != nil {
			messageProcessingModel = types.StringValue(string(*d.MessageProcessingModel))
		}
		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}
		if d.NotifyFilter != nil {
			notifyFilter = types.StringValue(*d.NotifyFilter)
		}
		if d.SecurityLevel != nil {
			securityLevel = types.StringValue(string(*d.SecurityLevel))
		}
		if d.SecurityModel != nil {
			securityModel = types.StringValue(string(*d.SecurityModel))
		}
		if d.SecurityName != nil {
			securityName = types.StringValue(*d.SecurityName)
		}

		dataMapValue := map[string]attr.Value{
			"message_processing_model": messageProcessingModel,
			"name":                     name,
			"notify_filter":            notifyFilter,
			"security_level":           securityLevel,
			"security_model":           securityModel,
			"security_name":            securityName,
		}
		data, e := NewTargetParametersValue(TargetParametersValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, TargetParametersValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}

// USM
func snmpV3UsmUsersSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.SnmpUsmpUser) basetypes.ListValue {
	var dataList []Snmpv3UsersValue
	for _, d := range l {
		var authenticationPassword basetypes.StringValue
		var authenticationType basetypes.StringValue
		var encryptionPassword basetypes.StringValue
		var encryptionType basetypes.StringValue
		var name basetypes.StringValue

		if d.AuthenticationPassword != nil {
			authenticationPassword = types.StringValue(*d.AuthenticationPassword)
		}
		if d.AuthenticationType != nil {
			authenticationType = types.StringValue(string(*d.AuthenticationType))
		}
		if d.EncryptionPassword != nil {
			encryptionPassword = types.StringValue(*d.EncryptionPassword)
		}
		if d.EncryptionType != nil {
			encryptionType = types.StringValue(string(*d.EncryptionType))
		}
		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}

		dataMapValue := map[string]attr.Value{
			"authentication_password": authenticationPassword,
			"authentication_type":     authenticationType,
			"encryption_password":     encryptionPassword,
			"encryption_type":         encryptionType,
			"name":                    name,
		}
		data, e := NewSnmpv3UsersValue(Snmpv3UsersValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, Snmpv3UsersValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}
func snmpV3UsmSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SnmpUsm) basetypes.ObjectValue {

	var engineType basetypes.StringValue
	var engineid basetypes.StringValue
	var users = snmpV3UsmUsersSdkToTerraform(ctx, diags, d.Users)

	if d.EngineType != nil {
		engineType = types.StringValue(string(*d.EngineType))
	}
	if d.EngineId != nil {
		engineid = types.StringValue(*d.EngineId)
	}

	dataMapValue := map[string]attr.Value{
		"engine_type": engineType,
		"engineid":    engineid,
		"users":       users,
	}
	//data, e := NewUsmValue(data_map_attr_type, dataMapValue)
	data, e := basetypes.NewObjectValue(UsmValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

// VACM ACCESS
func snmpV3VacmAccessPrefixListSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.SnmpVacmAccessItemPrefixListItem) basetypes.ListValue {
	var dataList []PrefixListValue
	for _, d := range l {
		var contextPrefix basetypes.StringValue
		var notifyView basetypes.StringValue
		var readView basetypes.StringValue
		var securityLevel basetypes.StringValue
		var securityModel basetypes.StringValue
		var prefixType basetypes.StringValue
		var writeView basetypes.StringValue

		if d.ContextPrefix != nil {
			contextPrefix = types.StringValue(*d.ContextPrefix)
		}
		if d.NotifyView != nil {
			notifyView = types.StringValue(*d.NotifyView)
		}
		if d.ReadView != nil {
			readView = types.StringValue(*d.ReadView)
		}
		if d.SecurityLevel != nil {
			securityLevel = types.StringValue(string(*d.SecurityLevel))
		}
		if d.SecurityModel != nil {
			securityModel = types.StringValue(string(*d.SecurityModel))
		}
		if d.Type != nil {
			prefixType = types.StringValue(string(*d.Type))
		}
		if d.WriteView != nil {
			writeView = types.StringValue(*d.WriteView)
		}

		dataMapValue := map[string]attr.Value{
			"context_prefix": contextPrefix,
			"notify_view":    notifyView,
			"read_view":      readView,
			"security_level": securityLevel,
			"security_model": securityModel,
			"type":           prefixType,
			"write_view":     writeView,
		}
		data, e := NewPrefixListValue(PrefixListValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, PrefixListValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}
func snmpV3VacmAccessSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.SnmpVacmAccessItem) basetypes.ListValue {
	var dataList []AccessValue
	for _, d := range l {
		var groupName basetypes.StringValue
		var prefixList = snmpV3VacmAccessPrefixListSdkToTerraform(ctx, diags, d.PrefixList)

		if d.GroupName != nil {
			groupName = types.StringValue(*d.GroupName)
		}

		dataMapValue := map[string]attr.Value{
			"group_name":  groupName,
			"prefix_list": prefixList,
		}
		data, e := NewAccessValue(AccessValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, AccessValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}

// VACM SECURITY TO GROUP
func snmpV3VacmSecurityToGroupContentSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.SnmpVacmSecurityToGroupContentItem) basetypes.ListValue {
	var dataList []Snmpv3VacmContentValue
	for _, d := range l {
		var group basetypes.StringValue
		var securityName basetypes.StringValue

		if d.Group != nil {
			group = types.StringValue(*d.Group)
		}
		if d.SecurityName != nil {
			securityName = types.StringValue(*d.SecurityName)
		}

		dataMapValue := map[string]attr.Value{
			"group":         group,
			"security_name": securityName,
		}
		data, e := NewSnmpv3VacmContentValue(Snmpv3VacmContentValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, Snmpv3VacmContentValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}
func snmpV3VacmSecurityToGroupSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SnmpVacmSecurityToGroup) basetypes.ObjectValue {
	content := snmpV3VacmSecurityToGroupContentSdkToTerraform(ctx, diags, d.Content)

	var securityModel basetypes.StringValue

	if d.SecurityModel != nil {
		securityModel = types.StringValue(string(*d.SecurityModel))
	}

	dataMapValue := map[string]attr.Value{
		"security_model": securityModel,
		"content":        content,
	}
	// data, e := NewSecurityToGroupValue(data_map_attr_type, dataMapValue)
	data, e := basetypes.NewObjectValue(SecurityToGroupValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

// VACM
func snmpV3VacmSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SnmpVacm) basetypes.ObjectValue {
	access := snmpV3VacmAccessSdkToTerraform(ctx, diags, d.Access)
	securityToGroup := snmpV3VacmSecurityToGroupSdkToTerraform(ctx, diags, d.SecurityToGroup)

	rAttrType := VacmValue{}.AttributeTypes(ctx)
	rAttrValue := map[string]attr.Value{
		"access":            access,
		"security_to_group": securityToGroup,
	}
	r, e := basetypes.NewObjectValue(rAttrType, rAttrValue)
	diags.Append(e...)
	return r
}

// V3
func snmpV3SdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.Snmpv3Config) basetypes.ObjectValue {
	var notify = types.ListNull(NotifyValue{}.Type(ctx))
	var notifyFilter = types.ListNull(NotifyFilterValue{}.Type(ctx))
	var targetAddress = types.ListNull(TargetAddressValue{}.Type(ctx))
	var targetParameters = types.ListNull(TargetParametersValue{}.Type(ctx))
	var usm = types.ObjectNull(UsersValue{}.AttributeTypes(ctx))
	var vacm = types.ObjectNull(VacmValue{}.AttributeTypes(ctx))

	if d.Notify != nil {
		notify = snmpV3NotifySdkToTerraform(ctx, diags, d.Notify)
	}
	if d.NotifyFilter != nil {
		notifyFilter = snmpV3NotifyFilterSdkToTerraform(ctx, diags, d.NotifyFilter)
	}
	if d.TargetAddress != nil {
		targetAddress = snmpV3TargetAddressSdkToTerraform(ctx, diags, d.TargetAddress)
	}
	if d.TargetParameters != nil {
		targetParameters = snmpV3TargetParametersSdkToTerraform(ctx, diags, d.TargetParameters)
	}
	if d.Usm != nil {
		usm = snmpV3UsmSdkToTerraform(ctx, diags, d.Usm)
	}
	if d.Vacm != nil {
		vacm = snmpV3VacmSdkToTerraform(ctx, diags, d.Vacm)

	}

	rAttrType := V3ConfigValue{}.AttributeTypes(ctx)
	rAttrValue := map[string]attr.Value{
		"notify":            notify,
		"notify_filter":     notifyFilter,
		"target_address":    targetAddress,
		"target_parameters": targetParameters,
		"usm":               usm,
		"vacm":              vacm,
	}

	data, e := basetypes.NewObjectValue(rAttrType, rAttrValue)
	diags.Append(e...)
	return data
}

// ////////////////////////////////
// VIEWS
func snmpConfigViewsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.SnmpConfigView) basetypes.ListValue {
	var dataList []ViewsValue
	for _, d := range l {
		var include basetypes.BoolValue
		var oid basetypes.StringValue
		var viewName basetypes.StringValue

		if d.Include != nil {
			include = types.BoolValue(*d.Include)
		}
		if d.Oid != nil {
			oid = types.StringValue(*d.Oid)
		}
		if d.ViewName != nil {
			viewName = types.StringValue(*d.ViewName)
		}

		dataMapValue := map[string]attr.Value{
			"include":   include,
			"oid":       oid,
			"view_name": viewName,
		}
		data, e := NewViewsValue(ViewsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, ViewsValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}

// //////////////////////////////////
// ////////// MAIN
func snmpConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SnmpConfig) SnmpConfigValue {

	var clientList = types.ListNull(ClientListValue{}.Type(ctx))
	var contact basetypes.StringValue
	var description basetypes.StringValue
	var enabled basetypes.BoolValue
	var engineId basetypes.StringValue
	var location basetypes.StringValue
	var name basetypes.StringValue
	var network basetypes.StringValue
	var trapGroups = types.ListNull(TrapGroupsValue{}.Type(ctx))
	var v2cConfig = types.ListNull(V2cConfigValue{}.Type(ctx))
	var v3Config = types.ObjectNull(V3ConfigValue{}.AttributeTypes(ctx))
	var views = types.ListNull(ViewsValue{}.Type(ctx))

	if d != nil {
		clientList = snmpClientListSdkToTerraform(ctx, diags, d.ClientList)

		if d.Contact != nil {
			contact = types.StringValue(*d.Contact)
		}
		if d.Description != nil {
			description = types.StringValue(*d.Description)
		}
		if d.Enabled != nil {
			enabled = types.BoolValue(*d.Enabled)
		}
		if d.EngineId != nil {
			engineId = types.StringValue(string(*d.EngineId))
		}
		if d.Location != nil {
			location = types.StringValue(*d.Location)
		}
		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}
		if d.Network != nil {
			network = types.StringValue(*d.Network)
		}
		if d.TrapGroups != nil {
			trapGroups = snmpTrapGroupsSdkToTerraform(ctx, diags, d.TrapGroups)
		}
		if d.V2cConfig != nil {
			v2cConfig = snmpV2cSdkToTerraform(ctx, diags, d.V2cConfig)
		}
		if d.V3Config != nil {
			v3Config = snmpV3SdkToTerraform(ctx, diags, d.V3Config)
		}
		if d.Views != nil {
			views = snmpConfigViewsSdkToTerraform(ctx, diags, d.Views)
		}
	}

	dataMapValue := map[string]attr.Value{
		"client_list": clientList,
		"contact":     contact,
		"description": description,
		"enabled":     enabled,
		"engine_id":   engineId,
		"location":    location,
		"name":        name,
		"network":     network,
		"trap_groups": trapGroups,
		"v2c_config":  v2cConfig,
		"v3_config":   v3Config,
		"views":       views,
	}
	data, e := NewSnmpConfigValue(SnmpConfigValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
