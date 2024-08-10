package resource_device_switch

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// //////////////////////////////////
// ////////// CLIENTS
func snmpClientListSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.SnmpConfigClientList) basetypes.ListValue {
	var data_list = []ClientListValue{}
	for _, d := range l {

		var client_list_name basetypes.StringValue
		var clients basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)

		if d.ClientListName != nil {
			client_list_name = types.StringValue(*d.ClientListName)
		}
		if d.Clients != nil {
			clients = mist_transform.ListOfStringSdkToTerraform(ctx, d.Clients)
		}

		data_map_attr_type := ClientListValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"client_list_name": client_list_name,
			"clients":          clients,
		}
		data, e := NewClientListValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, ClientListValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}

// //////////////////////////////////
// ////////// TRAP GROUPS
func snmpTrapGroupsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.SnmpConfigTrapGroup) basetypes.ListValue {
	var data_list = []TrapGroupsValue{}
	for _, d := range l {
		var categories basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
		var group_name basetypes.StringValue
		var targets basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
		var version basetypes.StringValue

		if d.Categories != nil {
			categories = mist_transform.ListOfStringSdkToTerraform(ctx, d.Categories)
		}
		if d.GroupName != nil {
			group_name = types.StringValue(*d.GroupName)
		}
		if d.Targets != nil {
			targets = mist_transform.ListOfStringSdkToTerraform(ctx, d.Targets)
		}
		if d.Version != nil {
			version = types.StringValue(string(*d.Version))
		}

		data_map_attr_type := TrapGroupsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"categories": categories,
			"group_name": group_name,
			"targets":    targets,
			"version":    version,
		}
		data, e := NewTrapGroupsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, TrapGroupsValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}

// //////////////////////////////////
// ////////// V2
func snmpV2cSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.SnmpConfigV2cConfig) basetypes.ListValue {
	var data_list = []V2cConfigValue{}
	for _, d := range l {
		var authorization basetypes.StringValue
		var client_list_name basetypes.StringValue
		var community_name basetypes.StringValue
		var view basetypes.StringValue

		if d.Authorization != nil {
			authorization = types.StringValue(*d.Authorization)
		}
		if d.ClientListName != nil {
			client_list_name = types.StringValue(*d.ClientListName)
		}
		if d.CommunityName != nil {
			community_name = types.StringValue(*d.CommunityName)
		}
		if d.View != nil {
			view = types.StringValue(*d.View)
		}

		data_map_attr_type := V2cConfigValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"authorization":    authorization,
			"client_list_name": client_list_name,
			"community_name":   community_name,
			"view":             view,
		}
		data, e := NewV2cConfigValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, V2cConfigValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}

// //////////////////////////////////
// ////////// V3
// NOTIFY
func snmpV3NotifySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.Snmpv3ConfigNotifyItems) basetypes.ListValue {
	var data_list = []NotifyValue{}
	for _, d := range l {
		var name basetypes.StringValue
		var tag basetypes.StringValue
		var notify_type basetypes.StringValue

		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}
		if d.Tag != nil {
			tag = types.StringValue(*d.Tag)
		}
		if d.Type != nil {
			notify_type = types.StringValue(string(*d.Type))
		}

		data_map_attr_type := NotifyValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"name": name,
			"tag":  tag,
			"type": notify_type,
		}
		data, e := NewNotifyValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, NotifyValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}

// NOTIFY Filter
func snmpV3NotifyFilterContentSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.Snmpv3ConfigNotifyFilterItemContent) basetypes.ListValue {
	var data_list = []Snmpv3ContentsValue{}
	for _, d := range l {

		var include basetypes.BoolValue
		var oid basetypes.StringValue

		if d.Include != nil {
			include = types.BoolValue(*d.Include)
		}
		if d.Oid != nil {
			oid = types.StringValue(*d.Oid)
		}

		data_map_attr_type := Snmpv3ContentsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"include": include,
			"oid":     oid,
		}
		data, e := NewSnmpv3ContentsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, Snmpv3ContentsValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}
func snmpV3NotifyFilterSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.Snmpv3ConfigNotifyFilterItem) basetypes.ListValue {
	var data_list = []NotifyFilterValue{}
	for _, d := range l {
		var contents basetypes.ListValue = snmpV3NotifyFilterContentSdkToTerraform(ctx, diags, d.Contents)
		var profile_name basetypes.StringValue

		if d.ProfileName != nil {
			profile_name = types.StringValue(*d.ProfileName)
		}

		data_map_attr_type := NotifyFilterValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"contents":     contents,
			"profile_name": profile_name,
		}
		data, e := NewNotifyFilterValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, NotifyFilterValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}

// TARGET ADDRESS
func snmpV3TargetAddressSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.Snmpv3ConfigTargetAddressItem) basetypes.ListValue {
	var data_list = []TargetAddressValue{}
	for _, d := range l {
		var address basetypes.StringValue
		var address_mask basetypes.StringValue
		var port basetypes.Int64Value
		var tag_list basetypes.StringValue
		var target_address_name basetypes.StringValue
		var target_parameters basetypes.StringValue

		if d.Address != nil {
			address = types.StringValue(*d.Address)
		}
		if d.AddressMask != nil {
			address_mask = types.StringValue(*d.AddressMask)
		}
		if d.Port != nil {
			port = types.Int64Value(int64(*d.Port))
		}
		if d.TagList != nil {
			tag_list = types.StringValue(*d.TagList)
		}
		if d.TargetAddressName != nil {
			target_address_name = types.StringValue(*d.TargetAddressName)
		}
		if d.TargetParameters != nil {
			target_parameters = types.StringValue(*d.TargetParameters)
		}

		data_map_attr_type := TargetAddressValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"address":             address,
			"address_mask":        address_mask,
			"port":                port,
			"tag_list":            tag_list,
			"target_address_name": target_address_name,
			"target_parameters":   target_parameters,
		}
		data, e := NewTargetAddressValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, TargetAddressValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}

// TARGET PARAMETERS
func snmpV3TargetParametersSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.Snmpv3ConfigTargetParam) basetypes.ListValue {
	var data_list = []TargetParametersValue{}
	for _, d := range l {
		var message_processing_model basetypes.StringValue
		var name basetypes.StringValue
		var notify_filter basetypes.StringValue
		var security_level basetypes.StringValue
		var security_model basetypes.StringValue
		var security_name basetypes.StringValue

		if d.MessageProcessingModel != nil {
			message_processing_model = types.StringValue(string(*d.MessageProcessingModel))
		}
		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}
		if d.NotifyFilter != nil {
			notify_filter = types.StringValue(*d.NotifyFilter)
		}
		if d.SecurityLevel != nil {
			security_level = types.StringValue(string(*d.SecurityLevel))
		}
		if d.SecurityModel != nil {
			security_model = types.StringValue(string(*d.SecurityModel))
		}
		if d.SecurityName != nil {
			security_name = types.StringValue(*d.SecurityName)
		}

		data_map_attr_type := TargetParametersValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"message_processing_model": message_processing_model,
			"name":                     name,
			"notify_filter":            notify_filter,
			"security_level":           security_level,
			"security_model":           security_model,
			"security_name":            security_name,
		}
		data, e := NewTargetParametersValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, TargetParametersValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}

// USM
func snmpV3UsmUsersSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.SnmpUsmpUser) basetypes.ListValue {
	var data_list = []Snmpv3UsersValue{}
	for _, d := range l {
		var authentication_password basetypes.StringValue
		var authentication_type basetypes.StringValue
		var encryption_password basetypes.StringValue
		var encryption_type basetypes.StringValue
		var name basetypes.StringValue

		if d.AuthenticationPassword != nil {
			authentication_password = types.StringValue(*d.AuthenticationPassword)
		}
		if d.AuthenticationType != nil {
			authentication_type = types.StringValue(string(*d.AuthenticationType))
		}
		if d.EncryptionPassword != nil {
			encryption_password = types.StringValue(*d.EncryptionPassword)
		}
		if d.EncryptionType != nil {
			encryption_type = types.StringValue(string(*d.EncryptionType))
		}
		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}

		data_map_attr_type := Snmpv3UsersValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"authentication_password": authentication_password,
			"authentication_type":     authentication_type,
			"encryption_password":     encryption_password,
			"encryption_type":         encryption_type,
			"name":                    name,
		}
		data, e := NewSnmpv3UsersValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, Snmpv3UsersValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}
func snmpV3UsmSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SnmpUsm) basetypes.ObjectValue {

	var engine_type basetypes.StringValue
	var engineid basetypes.StringValue
	var users basetypes.ListValue = snmpV3UsmUsersSdkToTerraform(ctx, diags, d.Users)

	if d.EngineType != nil {
		engine_type = types.StringValue(string(*d.EngineType))
	}
	if d.EngineId != nil {
		engineid = types.StringValue(*d.EngineId)
	}

	data_map_attr_type := UsmValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"engine_type": engine_type,
		"engineid":    engineid,
		"users":       users,
	}
	//data, e := NewUsmValue(data_map_attr_type, data_map_value)
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

// VACM ACCESS
func snmpV3VacmAccessPrefixListSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.SnmpVacmAccessItemPrefixListItem) basetypes.ListValue {
	var data_list = []PrefixListValue{}
	for _, d := range l {
		var context_prefix basetypes.StringValue
		var notify_view basetypes.StringValue
		var read_view basetypes.StringValue
		var security_level basetypes.StringValue
		var security_model basetypes.StringValue
		var prefix_type basetypes.StringValue
		var write_view basetypes.StringValue

		if d.ContextPrefix != nil {
			context_prefix = types.StringValue(*d.ContextPrefix)
		}
		if d.NotifyView != nil {
			notify_view = types.StringValue(*d.NotifyView)
		}
		if d.ReadView != nil {
			read_view = types.StringValue(*d.ReadView)
		}
		if d.SecurityLevel != nil {
			security_level = types.StringValue(string(*d.SecurityLevel))
		}
		if d.SecurityModel != nil {
			security_model = types.StringValue(string(*d.SecurityModel))
		}
		if d.Type != nil {
			prefix_type = types.StringValue(string(*d.Type))
		}
		if d.WriteView != nil {
			write_view = types.StringValue(*d.WriteView)
		}

		data_map_attr_type := PrefixListValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"context_prefix": context_prefix,
			"notify_view":    notify_view,
			"read_view":      read_view,
			"security_level": security_level,
			"security_model": security_model,
			"type":           prefix_type,
			"write_view":     write_view,
		}
		data, e := NewPrefixListValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, PrefixListValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}
func snmpV3VacmAccessSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.SnmpVacmAccessItem) basetypes.ListValue {
	var data_list = []AccessValue{}
	for _, d := range l {
		var group_name basetypes.StringValue
		var prefix_list basetypes.ListValue = snmpV3VacmAccessPrefixListSdkToTerraform(ctx, diags, d.PrefixList)

		if d.GroupName != nil {
			group_name = types.StringValue(*d.GroupName)
		}

		data_map_attr_type := AccessValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"group_name":  group_name,
			"prefix_list": prefix_list,
		}
		data, e := NewAccessValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, AccessValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}

// VACM SECURITY TO GROUP
func snmpV3VacmSecurityToGroupContentSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.SnmpVacmSecurityToGroupContentItem) basetypes.ListValue {
	var data_list = []Snmpv3VacmContentValue{}
	for _, d := range l {
		var group basetypes.StringValue
		var security_name basetypes.StringValue

		if d.Group != nil {
			group = types.StringValue(*d.Group)
		}
		if d.SecurityName != nil {
			security_name = types.StringValue(*d.SecurityName)
		}

		data_map_attr_type := Snmpv3VacmContentValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"group":         group,
			"security_name": security_name,
		}
		data, e := NewSnmpv3VacmContentValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, Snmpv3VacmContentValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}
func snmpV3VacmSecurityToGroupSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SnmpVacmSecurityToGroup) basetypes.ObjectValue {
	content := snmpV3VacmSecurityToGroupContentSdkToTerraform(ctx, diags, d.Content)

	var security_model basetypes.StringValue

	if d.SecurityModel != nil {
		security_model = types.StringValue(string(*d.SecurityModel))
	}

	data_map_attr_type := SecurityToGroupValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"security_model": security_model,
		"content":        content,
	}
	// data, e := NewSecurityToGroupValue(data_map_attr_type, data_map_value)
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

// VACM
func snmpV3VacmSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SnmpVacm) basetypes.ObjectValue {
	access := snmpV3VacmAccessSdkToTerraform(ctx, diags, d.Access)
	security_to_group := snmpV3VacmSecurityToGroupSdkToTerraform(ctx, diags, d.SecurityToGroup)

	r_attr_type := VacmValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"access":            access,
		"security_to_group": security_to_group,
	}
	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}

// V3
func snmpV3SdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.Snmpv3Config) basetypes.ObjectValue {
	var notify basetypes.ListValue = types.ListNull(NotifyValue{}.Type(ctx))
	var notify_filter basetypes.ListValue = types.ListNull(NotifyFilterValue{}.Type(ctx))
	var target_address basetypes.ListValue = types.ListNull(TargetAddressValue{}.Type(ctx))
	var target_parameters basetypes.ListValue = types.ListNull(TargetParametersValue{}.Type(ctx))
	var usm basetypes.ObjectValue = types.ObjectNull(UsersValue{}.AttributeTypes(ctx))
	var vacm basetypes.ObjectValue = types.ObjectNull(VacmValue{}.AttributeTypes(ctx))

	if d.Notify != nil {
		notify = snmpV3NotifySdkToTerraform(ctx, diags, d.Notify)
	}
	if d.NotifyFilter != nil {
		notify_filter = snmpV3NotifyFilterSdkToTerraform(ctx, diags, d.NotifyFilter)
	}
	if d.TargetAddress != nil {
		target_address = snmpV3TargetAddressSdkToTerraform(ctx, diags, d.TargetAddress)
	}
	if d.TargetParameters != nil {
		target_parameters = snmpV3TargetParametersSdkToTerraform(ctx, diags, d.TargetParameters)
	}
	if d.Usm != nil {
		usm = snmpV3UsmSdkToTerraform(ctx, diags, d.Usm)
	}
	if d.Vacm != nil {
		vacm = snmpV3VacmSdkToTerraform(ctx, diags, d.Vacm)

	}

	r_attr_type := V3ConfigValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"notify":            notify,
		"notify_filter":     notify_filter,
		"target_address":    target_address,
		"target_parameters": target_parameters,
		"usm":               usm,
		"vacm":              vacm,
	}

	data, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return data
}

// ////////////////////////////////
// VIEWS
func snmpConfigViewsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.SnmpConfigView) basetypes.ListValue {
	var data_list = []ViewsValue{}
	for _, d := range l {
		var include basetypes.BoolValue
		var oid basetypes.StringValue
		var view_name basetypes.StringValue

		if d.Include != nil {
			include = types.BoolValue(*d.Include)
		}
		if d.Oid != nil {
			oid = types.StringValue(*d.Oid)
		}
		if d.ViewName != nil {
			view_name = types.StringValue(*d.ViewName)
		}

		data_map_attr_type := ViewsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"include":   include,
			"oid":       oid,
			"view_name": view_name,
		}
		data, e := NewViewsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, ViewsValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}

// //////////////////////////////////
// ////////// MAIN
func snmpConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SnmpConfig) SnmpConfigValue {

	var client_list basetypes.ListValue = types.ListNull(ClientListValue{}.Type(ctx))
	var contact basetypes.StringValue
	var description basetypes.StringValue
	var enabled basetypes.BoolValue
	var engine_id basetypes.StringValue
	var location basetypes.StringValue
	var name basetypes.StringValue
	var network basetypes.StringValue
	var trap_groups basetypes.ListValue = types.ListNull(TrapGroupsValue{}.Type(ctx))
	var v2c_config basetypes.ListValue = types.ListNull(V2cConfigValue{}.Type(ctx))
	var v3_config basetypes.ObjectValue = types.ObjectNull(V3ConfigValue{}.AttributeTypes(ctx))
	var views basetypes.ListValue = types.ListNull(ViewsValue{}.Type(ctx))

	if d != nil {
		client_list = snmpClientListSdkToTerraform(ctx, diags, d.ClientList)

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
			engine_id = types.StringValue(string(*d.EngineId))
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
			trap_groups = snmpTrapGroupsSdkToTerraform(ctx, diags, d.TrapGroups)
		}
		if d.V2cConfig != nil {
			v2c_config = snmpV2cSdkToTerraform(ctx, diags, d.V2cConfig)
		}
		if d.V3Config != nil {
			v3_config = snmpV3SdkToTerraform(ctx, diags, d.V3Config)
		}
		if d.Views != nil {
			views = snmpConfigViewsSdkToTerraform(ctx, diags, d.Views)
		}
	}

	data_map_attr_type := SnmpConfigValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"client_list": client_list,
		"contact":     contact,
		"description": description,
		"enabled":     enabled,
		"engine_id":   engine_id,
		"location":    location,
		"name":        name,
		"network":     network,
		"trap_groups": trap_groups,
		"v2c_config":  v2c_config,
		"v3_config":   v3_config,
		"views":       views,
	}
	data, e := NewSnmpConfigValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
