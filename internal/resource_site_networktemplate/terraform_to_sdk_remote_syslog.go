package resource_site_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func remoteSyslogConfigArchiveTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.RemoteSyslogArchive {
	data := models.RemoteSyslogArchive{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		item, e := NewArchiveValue(ArchiveValue{}.AttributeTypes(ctx), d.Attributes())
		diags.Append(e...)
		var item_interface interface{} = item
		item_obj := item_interface.(ArchiveValue)
		if item_obj.Files.ValueInt64Pointer() != nil {
			data.Files = models.ToPointer(int(item_obj.Files.ValueInt64()))
		}
		if item_obj.Size.ValueStringPointer() != nil {
			data.Size = models.ToPointer(item_obj.Size.ValueString())
		}
		return &data
	}
}
func remoteSyslogArchiveTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.RemoteSyslogArchive {
	data := models.RemoteSyslogArchive{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		item, e := NewArchiveValue(ArchiveValue{}.AttributeTypes(ctx), d.Attributes())
		diags.Append(e...)
		var item_interface interface{} = item
		item_obj := item_interface.(ArchiveValue)
		if item_obj.Files.ValueInt64Pointer() != nil {
			data.Files = models.ToPointer(int(item_obj.Files.ValueInt64()))
		}
		if item_obj.Size.ValueStringPointer() != nil {
			data.Size = models.ToPointer(item_obj.Size.ValueString())
		}
		return &data
	}
}
func remoteSyslogContentTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.RemoteSyslogContent {
	var data []models.RemoteSyslogContent
	for _, v := range d.Elements() {
		var item_interface interface{} = v
		item_in := item_interface.(ContentsValue)
		item_out := models.RemoteSyslogContent{}
		facility := models.ToPointer(models.RemoteSyslogFacilityEnum(item_in.Facility.ValueString()))
		severity := models.ToPointer(models.RemoteSyslogSeverityEnum(item_in.Severity.ValueString()))
		item_out.Facility = models.ToPointer(*facility)
		item_out.Severity = models.ToPointer(*severity)
		data = append(data, item_out)
	}
	return data
}
func remoteSyslogConsoleTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.RemoteSyslogConsole {
	data := models.RemoteSyslogConsole{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		item_obj, e := NewConsoleValue(d.AttributeTypes(ctx), d.Attributes())
		diags.Append(e...)
		// var item_interface interface{} = d
		// item_obj := item_interface.(ConsoleValue)

		data.Contents = remoteSyslogContentTerraformToSdk(ctx, diags, item_obj.Contents)
		return &data
	}
}

func remoteSyslogFilesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.RemoteSyslogFileConfig {

	var data []models.RemoteSyslogFileConfig
	for _, item := range d.Elements() {
		var item_interface interface{} = item
		item_obj := item_interface.(FilesValue)

		data_item := models.RemoteSyslogFileConfig{}
		data_item.Archive = remoteSyslogConfigArchiveTerraformToSdk(ctx, diags, item_obj.Archive)
		data_item.Contents = remoteSyslogContentTerraformToSdk(ctx, diags, item_obj.Contents)
		if item_obj.ExplicitPriority.ValueBoolPointer() != nil {
			data_item.ExplicitPriority = models.ToPointer(item_obj.ExplicitPriority.ValueBool())
		}
		if item_obj.File.ValueStringPointer() != nil {
			data_item.File = models.ToPointer(item_obj.File.ValueString())
		}
		if item_obj.Match.ValueStringPointer() != nil {
			data_item.Match = models.ToPointer(item_obj.Match.ValueString())
		}
		if item_obj.StructuredData.ValueBoolPointer() != nil {
			data_item.StructuredData = models.ToPointer(item_obj.StructuredData.ValueBool())
		}
		data = append(data, data_item)
	}

	return data
}

func remoteSyslogServersTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.RemoteSyslogServer {

	var data []models.RemoteSyslogServer
	for _, item := range d.Elements() {
		var item_interface interface{} = item
		item_obj := item_interface.(ServersValue)

		data_item := models.RemoteSyslogServer{}
		data_item.Contents = remoteSyslogContentTerraformToSdk(ctx, diags, item_obj.Contents)
		if item_obj.ExplicitPriority.ValueBoolPointer() != nil {
			data_item.ExplicitPriority = models.ToPointer(item_obj.ExplicitPriority.ValueBool())
		}
		if item_obj.Facility.ValueStringPointer() != nil {
			data_item.Facility = models.ToPointer(models.RemoteSyslogFacilityEnum(item_obj.Facility.ValueString()))
		}
		if item_obj.Host.ValueStringPointer() != nil {
			data_item.Host = models.ToPointer(item_obj.Host.ValueString())
		}
		if item_obj.Match.ValueStringPointer() != nil {
			data_item.Match = models.ToPointer(item_obj.Match.ValueString())
		}
		if item_obj.Port.ValueInt64Pointer() != nil {
			data_item.Port = models.ToPointer(int(item_obj.Port.ValueInt64()))
		}
		if item_obj.Protocol.ValueStringPointer() != nil {
			data_item.Protocol = models.ToPointer(models.RemoteSyslogServerProtocolEnum(item_obj.Protocol.ValueString()))
		}
		if item_obj.RoutingInstance.ValueStringPointer() != nil {
			data_item.RoutingInstance = models.ToPointer(item_obj.RoutingInstance.ValueString())
		}
		if item_obj.Severity.ValueStringPointer() != nil {
			data_item.Severity = models.ToPointer(models.RemoteSyslogSeverityEnum(item_obj.Severity.ValueString()))
		}
		if item_obj.SourceAddress.ValueStringPointer() != nil {
			data_item.SourceAddress = models.ToPointer(item_obj.SourceAddress.ValueString())
		}
		if item_obj.StructuredData.ValueBoolPointer() != nil {
			data_item.StructuredData = models.ToPointer(item_obj.StructuredData.ValueBool())
		}
		if item_obj.Tag.ValueStringPointer() != nil {
			data_item.Tag = models.ToPointer(item_obj.Tag.ValueString())
		}
		data = append(data, data_item)
	}

	return data
}
func remoteSyslogUsersTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.RemoteSyslogUser {

	var data []models.RemoteSyslogUser
	for _, item := range d.Elements() {
		var item_interface interface{} = item
		item_obj := item_interface.(UsersValue)

		var content_list = []models.RemoteSyslogContent{}
		for _, content := range item_obj.Contents.Elements() {
			var content_interface interface{} = content
			content_obj := content_interface.(ContentsValue)
			content_out := models.RemoteSyslogContent{}

			if content_obj.Facility.ValueStringPointer() != nil {
				content_out.Facility = models.ToPointer(models.RemoteSyslogFacilityEnum(content_obj.Facility.ValueString()))
			}
			if content_obj.Severity.ValueStringPointer() != nil {
				content_out.Severity = models.ToPointer(models.RemoteSyslogSeverityEnum(content_obj.Severity.ValueString()))
			}
			content_list = append(content_list, content_out)
		}

		data_item := models.RemoteSyslogUser{}
		if item_obj.Match.ValueStringPointer() != nil {
			data_item.Match = models.ToPointer(item_obj.Match.ValueString())
		}
		if item_obj.User.ValueStringPointer() != nil {
			data_item.User = models.ToPointer(item_obj.User.ValueString())
		}
		data_item.Contents = content_list

		data = append(data, data_item)
	}

	return data
}

func remoteSyslogTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d RemoteSyslogValue) *models.RemoteSyslog {

	data := models.RemoteSyslog{}
	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = models.ToPointer(d.Enabled.ValueBool())
	}
	if d.Network.ValueStringPointer() != nil {
		data.Network = models.ToPointer(d.Network.ValueString())
	}
	if !d.Archive.IsNull() && !d.Archive.IsUnknown() {
		data.Archive = remoteSyslogArchiveTerraformToSdk(ctx, diags, d.Archive)
	}
	if !d.Console.IsNull() && !d.Console.IsUnknown() {
		data.Console = remoteSyslogConsoleTerraformToSdk(ctx, diags, d.Console)
	}
	if !d.Files.IsNull() && !d.Files.IsUnknown() {
		data.Files = remoteSyslogFilesTerraformToSdk(ctx, diags, d.Files)
	}
	if d.Network.ValueStringPointer() != nil {
		data.Network = models.ToPointer(d.Network.ValueString())
	}
	if d.SendToAllServers.ValueBoolPointer() != nil {
		data.SendToAllServers = models.ToPointer(d.SendToAllServers.ValueBool())
	}
	if !d.Servers.IsNull() && !d.Servers.IsUnknown() {
		data.Servers = remoteSyslogServersTerraformToSdk(ctx, diags, d.Servers)
	}
	if d.TimeFormat.ValueStringPointer() != nil {
		data.TimeFormat = models.ToPointer(models.RemoteSyslogTimeFormatEnum(d.TimeFormat.ValueString()))
	}
	if !d.Users.IsNull() && !d.Users.IsUnknown() {
		data.Users = remoteSyslogUsersTerraformToSdk(ctx, diags, d.Users)
	}

	return &data
}
