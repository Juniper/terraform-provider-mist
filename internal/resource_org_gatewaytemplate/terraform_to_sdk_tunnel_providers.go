package resource_org_gatewaytemplate

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func tunnelProviderOptionsJseTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.TunnelProviderOptionsJse {
	data := models.TunnelProviderOptionsJse{}
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewJseValueMust(d.AttributeTypes(ctx), d.Attributes())
		if plan.Name.ValueStringPointer() != nil {
			data.Name = plan.Name.ValueStringPointer()
		}
		if plan.NumUsers.ValueInt64Pointer() != nil {
			data.NumUsers = models.ToPointer(int(plan.NumUsers.ValueInt64()))
		}
		return data
	}
}

func tunnelProviderOptionsZscalerSubLocationTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.TunnelProviderOptionsZscalerSubLocation {
	var data_list []models.TunnelProviderOptionsZscalerSubLocation
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(SubLocationsValue)
		data := models.TunnelProviderOptionsZscalerSubLocation{}
		if plan.AupAcceptanceRequired.ValueBoolPointer() != nil {
			data.AupAcceptanceRequired = plan.AupAcceptanceRequired.ValueBoolPointer()
		}
		if plan.AupExpire.ValueInt64Pointer() != nil {
			data.AupExpire = models.ToPointer(int(plan.AupExpire.ValueInt64()))
		}
		if plan.AupSslProxy.ValueBoolPointer() != nil {
			data.AupSslProxy = plan.AupSslProxy.ValueBoolPointer()
		}
		if plan.DownloadMbps.ValueInt64Pointer() != nil {
			data.DownloadMbps = models.ToPointer(int(plan.DownloadMbps.ValueInt64()))
		}
		if plan.EnableAup.ValueBoolPointer() != nil {
			data.EnableAup = plan.EnableAup.ValueBoolPointer()
		}
		if plan.EnableCaution.ValueBoolPointer() != nil {
			data.EnableCaution = plan.EnableCaution.ValueBoolPointer()
		}
		if plan.EnforceAuthentication.ValueBoolPointer() != nil {
			data.EnforceAuthentication = plan.EnforceAuthentication.ValueBoolPointer()
		}
		if !plan.Subnets.IsNull() && !plan.Subnets.IsUnknown() {
			data.Subnets = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Subnets)
		}
		if plan.UploadMbps.ValueInt64Pointer() != nil {
			data.UploadMbps = models.ToPointer(int(plan.UploadMbps.ValueInt64()))
		}

		data_list = append(data_list, data)
	}
	return data_list
}

func tunnelProviderOptionsZscalerTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.TunnelProviderOptionsZscaler {
	data := models.TunnelProviderOptionsZscaler{}
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewZscalerValueMust(d.AttributeTypes(ctx), d.Attributes())
		if plan.AupAcceptanceRequired.ValueBoolPointer() != nil {
			data.AupAcceptanceRequired = plan.AupAcceptanceRequired.ValueBoolPointer()
		}
		if plan.AupExpire.ValueInt64Pointer() != nil {
			data.AupExpire = models.ToPointer(int(plan.AupExpire.ValueInt64()))
		}
		if plan.AupSslProxy.ValueBoolPointer() != nil {
			data.AupSslProxy = plan.AupSslProxy.ValueBoolPointer()
		}
		if plan.DownloadMbps.ValueInt64Pointer() != nil {
			data.DownloadMbps = models.ToPointer(int(plan.DownloadMbps.ValueInt64()))
		}
		if plan.EnableAup.ValueBoolPointer() != nil {
			data.EnableAup = plan.EnableAup.ValueBoolPointer()
		}
		if plan.EnableCaution.ValueBoolPointer() != nil {
			data.EnableCaution = plan.EnableCaution.ValueBoolPointer()
		}
		if plan.EnforceAuthentication.ValueBoolPointer() != nil {
			data.EnforceAuthentication = plan.EnforceAuthentication.ValueBoolPointer()
		}
		if plan.Name.ValueStringPointer() != nil {
			data.Name = plan.Name.ValueStringPointer()
		}

		sub_locations := tunnelProviderOptionsZscalerSubLocationTerraformToSdk(ctx, diags, plan.SubLocations)
		if !plan.SubLocations.IsNull() && !plan.SubLocations.IsUnknown() {
			data.SubLocations = sub_locations
		}

		if plan.UploadMbps.ValueInt64Pointer() != nil {
			data.UploadMbps = models.ToPointer(int(plan.UploadMbps.ValueInt64()))
		}
		if plan.UseXff.ValueBoolPointer() != nil {
			data.UseXff = plan.UseXff.ValueBoolPointer()
		}

		return data
	}
}

func tunnelProviderOptionsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d TunnelProviderOptionsValue) models.TunnelProviderOptions {

	data := models.TunnelProviderOptions{}

	jse := tunnelProviderOptionsJseTerraformToSdk(ctx, diags, d.Jse)
	if !d.Jse.IsNull() && !d.Jse.IsUnknown() {
		data.Jse = &jse
	}

	zscaler := tunnelProviderOptionsZscalerTerraformToSdk(ctx, diags, d.Zscaler)
	if !d.Zscaler.IsNull() && !d.Zscaler.IsUnknown() {
		data.Zscaler = &zscaler
	}

	return data
}
