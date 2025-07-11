package resource_org_wlan

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func radsecServersTerraformToSdk(d basetypes.ListValue) []models.RadsecServer {
	var dataList []models.RadsecServer
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(ServersValue)
		data := models.RadsecServer{}
		data.Host = plan.Host.ValueStringPointer()
		data.Port = models.ToPointer(int(plan.Port.ValueInt64()))

		dataList = append(dataList, data)
	}
	return dataList
}

func radsecTerraformToSdk(d RadsecValue) *models.Radsec {
	data := models.Radsec{}

	if d.CoaEnabled.ValueBoolPointer() != nil {
		data.CoaEnabled = d.CoaEnabled.ValueBoolPointer()
	}
	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if d.IdleTimeout.ValueStringPointer() != nil {
		data.IdleTimeout = models.ToPointer(models.RadsecIdleTimeoutContainer.FromString(d.IdleTimeout.ValueString()))
	}
	if !d.MxclusterIds.IsNull() && !d.MxclusterIds.IsUnknown() {
		data.MxclusterIds = mistutils.ListOfUuidTerraformToSdk(d.MxclusterIds)
	}
	if !d.ProxyHosts.IsNull() && !d.ProxyHosts.IsUnknown() {
		data.ProxyHosts = mistutils.ListOfStringTerraformToSdk(d.ProxyHosts)
	}
	if d.ServerName.ValueStringPointer() != nil {
		data.ServerName = d.ServerName.ValueStringPointer()
	}

	servers := radsecServersTerraformToSdk(d.Servers)
	data.Servers = servers

	data.UseMxedge = d.UseMxedge.ValueBoolPointer()
	data.UseSiteMxedge = d.UseSiteMxedge.ValueBoolPointer()

	return &data
}
