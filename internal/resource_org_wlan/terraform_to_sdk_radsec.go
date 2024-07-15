package resource_org_wlan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func radsecServersTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.RadsecServer {
	var data_list []models.RadsecServer
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(ServersValue)
		data := models.RadsecServer{}
		data.Host = plan.Host.ValueStringPointer()
		data.Port = models.ToPointer(int(plan.Port.ValueInt64()))

		data_list = append(data_list, data)
	}
	return data_list
}

func radsecTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d RadsecValue) *models.Radsec {
	data := models.Radsec{}

	data.CoaEnabled = d.CoaEnabled.ValueBoolPointer()
	data.Enabled = d.Enabled.ValueBoolPointer()
	data.IdleTimeout = models.ToPointer(int(d.IdleTimeout.ValueInt64()))
	data.MxclusterIds = mist_transform.ListOfUuidTerraformToSdk(ctx, d.MxclusterIds)
	data.ProxyHosts = mist_transform.ListOfStringTerraformToSdk(ctx, d.ProxyHosts)
	data.ServerName = d.ServerName.ValueStringPointer()

	servers := radsecServersTerraformToSdk(ctx, diags, d.Servers)
	data.Servers = servers

	data.UseMxedge = d.UseMxedge.ValueBoolPointer()
	data.UseSiteMxedge = d.UseSiteMxedge.ValueBoolPointer()

	return &data
}
