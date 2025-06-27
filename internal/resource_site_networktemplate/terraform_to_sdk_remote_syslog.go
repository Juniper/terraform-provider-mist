package resource_site_networktemplate

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
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
		var itemInterface interface{} = item
		itemObj := itemInterface.(ArchiveValue)
		if itemObj.Files.ValueStringPointer() != nil {
			data.Files = models.ToPointer(models.RemoteSyslogArchiveFilesContainer.FromString(itemObj.Files.ValueString()))
		}
		if itemObj.Size.ValueStringPointer() != nil {
			data.Size = models.ToPointer(itemObj.Size.ValueString())
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
		var itemInterface interface{} = item
		itemObj := itemInterface.(ArchiveValue)
		if itemObj.Files.ValueStringPointer() != nil {
			data.Files = models.ToPointer(models.RemoteSyslogArchiveFilesContainer.FromString(itemObj.Files.ValueString()))
		}
		if itemObj.Size.ValueStringPointer() != nil {
			data.Size = models.ToPointer(itemObj.Size.ValueString())
		}
		return &data
	}
}
func remoteSyslogContentTerraformToSdk(d basetypes.ListValue) []models.RemoteSyslogContent {
	var data []models.RemoteSyslogContent
	for _, v := range d.Elements() {
		var itemInterface interface{} = v
		itemIn := itemInterface.(ContentsValue)
		itemOut := models.RemoteSyslogContent{}

		facility := models.ToPointer(models.RemoteSyslogFacilityEnum(itemIn.Facility.ValueString()))
		severity := models.ToPointer(models.RemoteSyslogSeverityEnum(itemIn.Severity.ValueString()))
		itemOut.Facility = models.ToPointer(*facility)
		itemOut.Severity = models.ToPointer(*severity)
		data = append(data, itemOut)
	}
	return data
}
func remoteSyslogConsoleTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.RemoteSyslogConsole {
	data := models.RemoteSyslogConsole{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		itemObj, e := NewConsoleValue(d.AttributeTypes(ctx), d.Attributes())
		diags.Append(e...)
		// var item_interface interface{} = d
		// item_obj := item_interface.(ConsoleValue)
		if !itemObj.Contents.IsNull() && !itemObj.Contents.IsUnknown() {
			data.Contents = remoteSyslogContentTerraformToSdk(itemObj.Contents)
		}
		return &data
	}
}

func remoteSyslogFilesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.RemoteSyslogFileConfig {

	var data []models.RemoteSyslogFileConfig
	for _, item := range d.Elements() {
		var itemInterface interface{} = item
		itemObj := itemInterface.(FilesValue)

		dataItem := models.RemoteSyslogFileConfig{}
		if !itemObj.Archive.IsNull() && !itemObj.Archive.IsUnknown() {
			dataItem.Archive = remoteSyslogConfigArchiveTerraformToSdk(ctx, diags, itemObj.Archive)
		}
		if !itemObj.Contents.IsNull() && !itemObj.Contents.IsUnknown() {
			dataItem.Contents = remoteSyslogContentTerraformToSdk(itemObj.Contents)
		}
		if itemObj.EnableTls.ValueBoolPointer() != nil {
			dataItem.EnableTls = models.ToPointer(itemObj.EnableTls.ValueBool())
		}
		if itemObj.ExplicitPriority.ValueBoolPointer() != nil {
			dataItem.ExplicitPriority = models.ToPointer(itemObj.ExplicitPriority.ValueBool())
		}
		if itemObj.File.ValueStringPointer() != nil {
			dataItem.File = models.ToPointer(itemObj.File.ValueString())
		}
		if itemObj.Match.ValueStringPointer() != nil {
			dataItem.Match = models.ToPointer(itemObj.Match.ValueString())
		}
		if itemObj.StructuredData.ValueBoolPointer() != nil {
			dataItem.StructuredData = models.ToPointer(itemObj.StructuredData.ValueBool())
		}
		data = append(data, dataItem)
	}

	return data
}

func remoteSyslogServersTerraformToSdk(d basetypes.ListValue) []models.RemoteSyslogServer {

	var data []models.RemoteSyslogServer
	for _, item := range d.Elements() {
		var itemInterface interface{} = item
		itemObj := itemInterface.(ServersValue)

		dataItem := models.RemoteSyslogServer{}
		if !itemObj.Contents.IsNull() && !itemObj.Contents.IsUnknown() {
			dataItem.Contents = remoteSyslogContentTerraformToSdk(itemObj.Contents)
		}
		if itemObj.ExplicitPriority.ValueBoolPointer() != nil {
			dataItem.ExplicitPriority = models.ToPointer(itemObj.ExplicitPriority.ValueBool())
		}
		if itemObj.Facility.ValueStringPointer() != nil {
			dataItem.Facility = models.ToPointer(models.RemoteSyslogFacilityEnum(itemObj.Facility.ValueString()))
		}
		if itemObj.Host.ValueStringPointer() != nil {
			dataItem.Host = models.ToPointer(itemObj.Host.ValueString())
		}
		if itemObj.Match.ValueStringPointer() != nil {
			dataItem.Match = models.ToPointer(itemObj.Match.ValueString())
		}
		if itemObj.Port.ValueStringPointer() != nil {
			dataItem.Port = models.ToPointer(models.RemoteSyslogServerPortContainer.FromString(itemObj.Port.ValueString()))
		}
		if itemObj.Protocol.ValueStringPointer() != nil {
			dataItem.Protocol = models.ToPointer(models.RemoteSyslogServerProtocolEnum(itemObj.Protocol.ValueString()))
		}
		if itemObj.RoutingInstance.ValueStringPointer() != nil {
			dataItem.RoutingInstance = models.ToPointer(itemObj.RoutingInstance.ValueString())
		}
		if itemObj.Severity.ValueStringPointer() != nil {
			dataItem.Severity = models.ToPointer(models.RemoteSyslogSeverityEnum(itemObj.Severity.ValueString()))
		}
		if itemObj.SourceAddress.ValueStringPointer() != nil {
			dataItem.SourceAddress = models.ToPointer(itemObj.SourceAddress.ValueString())
		}
		if itemObj.StructuredData.ValueBoolPointer() != nil {
			dataItem.StructuredData = models.ToPointer(itemObj.StructuredData.ValueBool())
		}
		if itemObj.Tag.ValueStringPointer() != nil {
			dataItem.Tag = models.ToPointer(itemObj.Tag.ValueString())
		}
		data = append(data, dataItem)
	}

	return data
}
func remoteSyslogUsersTerraformToSdk(d basetypes.ListValue) []models.RemoteSyslogUser {

	var data []models.RemoteSyslogUser
	for _, item := range d.Elements() {
		var itemInterface interface{} = item
		itemObj := itemInterface.(UsersValue)

		var contentList []models.RemoteSyslogContent
		for _, content := range itemObj.Contents.Elements() {
			var contentInterface interface{} = content
			contentObj := contentInterface.(ContentsValue)
			contentOut := models.RemoteSyslogContent{}

			if contentObj.Facility.ValueStringPointer() != nil {
				contentOut.Facility = models.ToPointer(models.RemoteSyslogFacilityEnum(contentObj.Facility.ValueString()))
			}
			if contentObj.Severity.ValueStringPointer() != nil {
				contentOut.Severity = models.ToPointer(models.RemoteSyslogSeverityEnum(contentObj.Severity.ValueString()))
			}
			contentList = append(contentList, contentOut)
		}

		dataItem := models.RemoteSyslogUser{}
		if itemObj.Match.ValueStringPointer() != nil {
			dataItem.Match = models.ToPointer(itemObj.Match.ValueString())
		}
		if itemObj.User.ValueStringPointer() != nil {
			dataItem.User = models.ToPointer(itemObj.User.ValueString())
		}
		dataItem.Contents = contentList

		data = append(data, dataItem)
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
	if !d.Cacerts.IsNull() && !d.Cacerts.IsUnknown() {
		data.Cacerts = mistutils.ListOfStringTerraformToSdk(d.Cacerts)
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
		data.Servers = remoteSyslogServersTerraformToSdk(d.Servers)
	}
	if d.TimeFormat.ValueStringPointer() != nil {
		data.TimeFormat = models.ToPointer(models.RemoteSyslogTimeFormatEnum(d.TimeFormat.ValueString()))
	}
	if !d.Users.IsNull() && !d.Users.IsUnknown() {
		data.Users = remoteSyslogUsersTerraformToSdk(d.Users)
	}

	return &data
}
