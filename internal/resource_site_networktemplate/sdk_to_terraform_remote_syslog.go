package resource_site_networktemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func remoteSyslogArchiveSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.RemoteSyslogArchive) basetypes.ObjectValue {

	var files basetypes.Int64Value
	var size basetypes.StringValue

	if d != nil && d.Files != nil {
		files = types.Int64Value(int64(*d.Files))
	}
	if d != nil && d.Size != nil {
		size = types.StringValue(*d.Size)
	}

	data_map_attr_type := ArchiveValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"files": files,
		"size":  size,
	}
	data, e := NewArchiveValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	o, e := data.ToObjectValue(ctx)
	diags.Append(e...)
	return o
}
func remoteSyslogContentsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.RemoteSyslogContent) basetypes.ListValue {
	var data_list = []ContentsValue{}

	for _, d := range l {
		var facility basetypes.StringValue
		var severity basetypes.StringValue

		if d.Facility != nil {
			facility = types.StringValue(string(*d.Facility))
		}
		if d.Severity != nil {
			severity = types.StringValue(string(*d.Severity))
		}

		data_map_attr_type := ContentsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"facility": facility,
			"severity": severity,
		}
		data, e := NewContentsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	data_list_type := ContentsValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}
func remoteSyslogConsoleSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.RemoteSyslogConsole) basetypes.ObjectValue {
	var contents basetypes.ListValue = types.ListNull(ContentsValue{}.Type(ctx))

	if d != nil && d.Contents != nil {
		contents = remoteSyslogContentsSdkToTerraform(ctx, diags, d.Contents)
	}

	data_map_attr_type := ConsoleValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"contents": contents,
	}
	data, e := NewConsoleValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	o, e := data.ToObjectValue(ctx)
	diags.Append(e...)
	return o
}
func remoteSyslogFilesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.RemoteSyslogFileConfig) basetypes.ListValue {
	var data_list = []FilesValue{}

	for _, d := range l {
		var archive basetypes.ObjectValue = types.ObjectNull(ArchiveValue{}.AttributeTypes(ctx))
		var contents basetypes.ListValue = types.ListNull(ContentsValue{}.Type(ctx))
		var explicit_priority basetypes.BoolValue
		var file basetypes.StringValue
		var match basetypes.StringValue
		var structured_data basetypes.BoolValue

		if d.Archive != nil {
			archive = remoteSyslogArchiveSdkToTerraform(ctx, diags, d.Archive)
		}
		if d.Contents != nil {
			contents = remoteSyslogContentsSdkToTerraform(ctx, diags, d.Contents)
		}
		if d.ExplicitPriority != nil {
			explicit_priority = types.BoolValue(*d.ExplicitPriority)
		}
		if d.File != nil {
			file = types.StringValue(*d.File)
		}
		if d.Match != nil {
			match = types.StringValue(*d.Match)
		}
		if d.StructuredData != nil {
			structured_data = types.BoolValue(*d.StructuredData)
		}

		data_map_attr_type := FilesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"archive":           archive,
			"contents":          contents,
			"explicit_priority": explicit_priority,
			"file":              file,
			"match":             match,
			"structured_data":   structured_data,
		}
		data, e := NewFilesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	data_list_type := FilesValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)

	return r
}
func remoteSyslogServerSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.RemoteSyslogServer) basetypes.ListValue {
	var data_list = []ServersValue{}

	for _, d := range l {
		var contents basetypes.ListValue = types.ListNull(ContentsValue{}.Type(ctx))
		var explicit_priority basetypes.BoolValue
		var facility basetypes.StringValue
		var host basetypes.StringValue
		var match basetypes.StringValue
		var port basetypes.Int64Value
		var protocol basetypes.StringValue
		var routing_instance basetypes.StringValue
		var severity basetypes.StringValue
		var source_address basetypes.StringValue
		var structured_data basetypes.BoolValue
		var tag basetypes.StringValue

		if d.Contents != nil {
			contents = remoteSyslogContentsSdkToTerraform(ctx, diags, d.Contents)
		}
		if d.ExplicitPriority != nil {
			explicit_priority = types.BoolValue(*d.ExplicitPriority)
		}
		if d.Facility != nil {
			facility = types.StringValue(string(*d.Facility))
		}
		if d.Host != nil {
			host = types.StringValue(*d.Host)
		}
		if d.Match != nil {
			match = types.StringValue(*d.Match)
		}
		if d.Port != nil {
			port = types.Int64Value(int64(*d.Port))
		}
		if d.Protocol != nil {
			protocol = types.StringValue(string(*d.Protocol))
		}
		if d.RoutingInstance != nil {
			routing_instance = types.StringValue(*d.RoutingInstance)
		}
		if d.Severity != nil {
			severity = types.StringValue(string(*d.Severity))
		}
		if d.SourceAddress != nil {
			source_address = types.StringValue(*d.SourceAddress)
		}
		if d.StructuredData != nil {
			structured_data = types.BoolValue(*d.StructuredData)
		}
		if d.Tag != nil {
			tag = types.StringValue(*d.Tag)
		}

		data_map_attr_type := ServersValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"contents":          contents,
			"explicit_priority": explicit_priority,
			"facility":          facility,
			"host":              host,
			"match":             match,
			"port":              port,
			"protocol":          protocol,
			"routing_instance":  routing_instance,
			"severity":          severity,
			"source_address":    source_address,
			"structured_data":   structured_data,
			"tag":               tag,
		}
		data, e := NewServersValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	data_list_type := ServersValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}
func remoteSyslogUsersSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.RemoteSyslogUser) basetypes.ListValue {
	var data_list = []UsersValue{}

	for _, d := range l {

		var contents basetypes.ListValue = types.ListNull(ContentsValue{}.Type(ctx))
		var match basetypes.StringValue
		var user basetypes.StringValue

		if d.Contents != nil {
			contents = remoteSyslogContentsSdkToTerraform(ctx, diags, d.Contents)
		}
		if d.Match != nil {
			match = types.StringValue(*d.Match)
		}
		if d.User != nil {
			user = types.StringValue(*d.User)
		}

		data_map_attr_type := UsersValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"contents": contents,
			"match":    match,
			"user":     user,
		}
		data, e := NewUsersValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	data_list_type := UsersValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}

func remoteSyslogSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.RemoteSyslog) RemoteSyslogValue {

	var archive basetypes.ObjectValue = types.ObjectNull(ArchiveValue{}.AttributeTypes(ctx))
	var console basetypes.ObjectValue = types.ObjectNull(ConsoleValue{}.AttributeTypes(ctx))
	var enabled basetypes.BoolValue
	var files basetypes.ListValue = types.ListNull(FilesValue{}.Type(ctx))
	var network basetypes.StringValue
	var send_to_all_servers basetypes.BoolValue
	var servers basetypes.ListValue = types.ListNull(ServersValue{}.Type(ctx))
	var time_format basetypes.StringValue
	var users basetypes.ListValue = types.ListNull(UsersValue{}.Type(ctx))

	if d != nil && d.Archive != nil {
		archive = remoteSyslogArchiveSdkToTerraform(ctx, diags, d.Archive)
	}
	if d != nil && d.Console != nil {
		console = remoteSyslogConsoleSdkToTerraform(ctx, diags, d.Console)
	}
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.Files != nil {
		files = remoteSyslogFilesSdkToTerraform(ctx, diags, d.Files)
	}
	if d != nil && d.Network != nil {
		network = types.StringValue(*d.Network)
	}
	if d != nil && d.SendToAllServers != nil {
		send_to_all_servers = types.BoolValue(*d.SendToAllServers)
	}
	if d != nil && d.Servers != nil {
		servers = remoteSyslogServerSdkToTerraform(ctx, diags, d.Servers)
	}
	if d != nil && d.TimeFormat != nil {
		time_format = types.StringValue(string(*d.TimeFormat))
	}
	if d != nil && d.Users != nil {
		users = remoteSyslogUsersSdkToTerraform(ctx, diags, d.Users)
	}

	data_map_attr_type := RemoteSyslogValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"archive":             archive,
		"console":             console,
		"enabled":             enabled,
		"files":               files,
		"network":             network,
		"send_to_all_servers": send_to_all_servers,
		"servers":             servers,
		"time_format":         time_format,
		"users":               users,
	}
	data, e := NewRemoteSyslogValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
