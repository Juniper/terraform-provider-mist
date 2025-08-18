package resource_org_networktemplate

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func remoteSyslogArchiveSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.RemoteSyslogArchive) basetypes.ObjectValue {

	var files basetypes.StringValue
	var size basetypes.StringValue

	if d != nil && d.Files != nil {
		files = mistutils.SyslogFilesAsString(d.Files)
	}
	if d != nil && d.Size != nil {
		size = types.StringValue(*d.Size)
	}

	dataMapValue := map[string]attr.Value{
		"files": files,
		"size":  size,
	}
	data, e := NewArchiveValue(ArchiveValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	o, e := data.ToObjectValue(ctx)
	diags.Append(e...)
	return o
}
func remoteSyslogContentsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.RemoteSyslogContent) basetypes.ListValue {
	var dataList []ContentsValue

	for _, d := range l {
		var facility basetypes.StringValue
		var severity basetypes.StringValue

		if d.Facility != nil {
			facility = types.StringValue(string(*d.Facility))
		}
		if d.Severity != nil {
			severity = types.StringValue(string(*d.Severity))
		}

		dataMapValue := map[string]attr.Value{
			"facility": facility,
			"severity": severity,
		}
		data, e := NewContentsValue(ContentsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	datalistType := ContentsValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, datalistType, dataList)
	diags.Append(e...)
	return r
}
func remoteSyslogConsoleSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.RemoteSyslogConsole) basetypes.ObjectValue {
	var contents = types.ListNull(ContentsValue{}.Type(ctx))

	if d != nil && d.Contents != nil {
		contents = remoteSyslogContentsSdkToTerraform(ctx, diags, d.Contents)
	}

	dataMapValue := map[string]attr.Value{
		"contents": contents,
	}
	data, e := NewConsoleValue(ConsoleValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	o, e := data.ToObjectValue(ctx)
	diags.Append(e...)
	return o
}
func remoteSyslogFilesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.RemoteSyslogFileConfig) basetypes.ListValue {
	var dataList []FilesValue

	for _, d := range l {
		var archive = types.ObjectNull(ArchiveValue{}.AttributeTypes(ctx))
		var contents = types.ListNull(ContentsValue{}.Type(ctx))
		var enableTls basetypes.BoolValue
		var explicitPriority basetypes.BoolValue
		var file basetypes.StringValue
		var match basetypes.StringValue
		var structuredData basetypes.BoolValue

		if d.Archive != nil {
			archive = remoteSyslogArchiveSdkToTerraform(ctx, diags, d.Archive)
		}
		if d.Contents != nil {
			contents = remoteSyslogContentsSdkToTerraform(ctx, diags, d.Contents)
		}
		if d.EnableTls != nil {
			enableTls = types.BoolValue(*d.EnableTls)
		}
		if d.ExplicitPriority != nil {
			explicitPriority = types.BoolValue(*d.ExplicitPriority)
		}
		if d.File != nil {
			file = types.StringValue(*d.File)
		}
		if d.Match != nil {
			match = types.StringValue(*d.Match)
		}
		if d.StructuredData != nil {
			structuredData = types.BoolValue(*d.StructuredData)
		}

		dataMapValue := map[string]attr.Value{
			"archive":           archive,
			"contents":          contents,
			"enable_tls":        enableTls,
			"explicit_priority": explicitPriority,
			"file":              file,
			"match":             match,
			"structured_data":   structuredData,
		}
		data, e := NewFilesValue(FilesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	datalistType := FilesValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, datalistType, dataList)
	diags.Append(e...)

	return r
}
func remoteSyslogServerSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.RemoteSyslogServer) basetypes.ListValue {
	var dataList []ServersValue

	for _, d := range l {
		var contents = types.ListNull(ContentsValue{}.Type(ctx))
		var explicitPriority basetypes.BoolValue
		var facility basetypes.StringValue
		var host basetypes.StringValue
		var match basetypes.StringValue
		var port basetypes.StringValue
		var protocol basetypes.StringValue
		var routingInstance basetypes.StringValue
		var serverName basetypes.StringValue
		var severity basetypes.StringValue
		var sourceAddress basetypes.StringValue
		var structuredData basetypes.BoolValue
		var tag basetypes.StringValue

		if d.Contents != nil {
			contents = remoteSyslogContentsSdkToTerraform(ctx, diags, d.Contents)
		}
		if d.ExplicitPriority != nil {
			explicitPriority = types.BoolValue(*d.ExplicitPriority)
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
			port = mistutils.SyslogPortAsString(d.Port)
		}
		if d.Protocol != nil {
			protocol = types.StringValue(string(*d.Protocol))
		}
		if d.RoutingInstance != nil {
			routingInstance = types.StringValue(*d.RoutingInstance)
		}
		if d.ServerName != nil {
			serverName = types.StringValue(*d.ServerName)
		}
		if d.Severity != nil {
			severity = types.StringValue(string(*d.Severity))
		}
		if d.SourceAddress != nil {
			sourceAddress = types.StringValue(*d.SourceAddress)
		}
		if d.StructuredData != nil {
			structuredData = types.BoolValue(*d.StructuredData)
		}
		if d.Tag != nil {
			tag = types.StringValue(*d.Tag)
		}

		dataMapValue := map[string]attr.Value{
			"contents":          contents,
			"explicit_priority": explicitPriority,
			"facility":          facility,
			"host":              host,
			"match":             match,
			"port":              port,
			"protocol":          protocol,
			"routing_instance":  routingInstance,
			"server_name":       serverName,
			"severity":          severity,
			"source_address":    sourceAddress,
			"structured_data":   structuredData,
			"tag":               tag,
		}
		data, e := NewServersValue(ServersValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	datalistType := ServersValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, datalistType, dataList)
	diags.Append(e...)
	return r
}
func remoteSyslogUsersSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.RemoteSyslogUser) basetypes.ListValue {
	var dataList []UsersValue

	for _, d := range l {

		var contents = types.ListNull(ContentsValue{}.Type(ctx))
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

		dataMapValue := map[string]attr.Value{
			"contents": contents,
			"match":    match,
			"user":     user,
		}
		data, e := NewUsersValue(UsersValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	datalistType := UsersValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, datalistType, dataList)
	diags.Append(e...)
	return r
}

func remoteSyslogSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.RemoteSyslog) RemoteSyslogValue {

	var archive = types.ObjectNull(ArchiveValue{}.AttributeTypes(ctx))
	var caCerts = types.ListNull(types.StringType)
	var console = types.ObjectNull(ConsoleValue{}.AttributeTypes(ctx))
	var enabled basetypes.BoolValue
	var files = types.ListNull(FilesValue{}.Type(ctx))
	var network basetypes.StringValue
	var sendToAllServers basetypes.BoolValue
	var servers = types.ListNull(ServersValue{}.Type(ctx))
	var timeFormat basetypes.StringValue
	var users = types.ListNull(UsersValue{}.Type(ctx))

	if d != nil && d.Archive != nil {
		archive = remoteSyslogArchiveSdkToTerraform(ctx, diags, d.Archive)
	}
	if d != nil && d.Cacerts != nil {
		caCerts = mistutils.ListOfStringSdkToTerraform(d.Cacerts)
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
		sendToAllServers = types.BoolValue(*d.SendToAllServers)
	}
	if d != nil && d.Servers != nil {
		servers = remoteSyslogServerSdkToTerraform(ctx, diags, d.Servers)
	}
	if d != nil && d.TimeFormat != nil {
		timeFormat = types.StringValue(string(*d.TimeFormat))
	}
	if d != nil && d.Users != nil {
		users = remoteSyslogUsersSdkToTerraform(ctx, diags, d.Users)
	}

	dataMapValue := map[string]attr.Value{
		"archive":             archive,
		"cacerts":             caCerts,
		"console":             console,
		"enabled":             enabled,
		"files":               files,
		"network":             network,
		"send_to_all_servers": sendToAllServers,
		"servers":             servers,
		"time_format":         timeFormat,
		"users":               users,
	}
	data, e := NewRemoteSyslogValue(RemoteSyslogValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
