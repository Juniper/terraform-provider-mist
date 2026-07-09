package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func tuntermMulticastConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d TuntermMulticastConfigValue) *models.SiteSettingTuntermMulticastConfig {
	data := models.SiteSettingTuntermMulticastConfig{}

	if !d.MulticastAll.IsNull() && !d.MulticastAll.IsUnknown() {
		data.MulticastAll = d.MulticastAll.ValueBoolPointer()
	}

	if !d.Mdns.IsNull() && !d.Mdns.IsUnknown() {
		data.Mdns = tuntermMulticastMdnsTerraformToSdk(ctx, diags, d.Mdns)
	}

	if !d.Ssdp.IsNull() && !d.Ssdp.IsUnknown() {
		data.Ssdp = tuntermMulticastSsdpTerraformToSdk(ctx, diags, d.Ssdp)
	}

	return &data
}

func tuntermMulticastMdnsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, obj types.Object) *models.SiteSettingTuntermMulticastConfigMdns {
	data := models.SiteSettingTuntermMulticastConfigMdns{}
	if obj.IsNull() || obj.IsUnknown() {
		return &data
	}

	attrs := obj.Attributes()
	if v, ok := attrs["enabled"].(basetypes.BoolValue); ok && !v.IsNull() && !v.IsUnknown() {
		data.Enabled = v.ValueBoolPointer()
	}
	if v, ok := attrs["vlan_ids"].(basetypes.ListValue); ok && !v.IsNull() && !v.IsUnknown() {
		data.VlanIds = mistutils.ListOfIntTerraformToSdk(v)
	}

	return &data
}

func tuntermMulticastSsdpTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, obj types.Object) *models.SiteSettingTuntermMulticastConfigSsdp {
	data := models.SiteSettingTuntermMulticastConfigSsdp{}
	if obj.IsNull() || obj.IsUnknown() {
		return &data
	}

	attrs := obj.Attributes()
	if v, ok := attrs["enabled"].(basetypes.BoolValue); ok && !v.IsNull() && !v.IsUnknown() {
		data.Enabled = v.ValueBoolPointer()
	}
	if v, ok := attrs["vlan_ids"].(basetypes.ListValue); ok && !v.IsNull() && !v.IsUnknown() {
		data.VlanIds = mistutils.ListOfIntTerraformToSdk(v)
	}

	return &data
}
