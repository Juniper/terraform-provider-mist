package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func tuntermMulticastConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteSettingTuntermMulticastConfig) TuntermMulticastConfigValue {
	mdnsObj := tuntermMulticastMdnsSdkToTerraform(ctx, diags, d.Mdns)
	var multicastAll = types.BoolNull()
	ssdpObj := tuntermMulticastSsdpSdkToTerraform(ctx, diags, d.Ssdp)

	if d.MulticastAll != nil {
		multicastAll = types.BoolValue(*d.MulticastAll)
	}

	dataMapAttrType := TuntermMulticastConfigValue{}.AttributeTypes(ctx)
	dataMapValue := map[string]attr.Value{
		"mdns":          mdnsObj,
		"multicast_all": multicastAll,
		"ssdp":          ssdpObj,
	}
	data, e := NewTuntermMulticastConfigValue(dataMapAttrType, dataMapValue)
	diags.Append(e...)

	return data
}

func tuntermMulticastMdnsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteSettingTuntermMulticastConfigMdns) types.Object {
	attrTypes := MdnsValue{}.AttributeTypes(ctx)

	if d == nil {
		return types.ObjectNull(attrTypes)
	}

	var enabled = types.BoolNull()
	var vlanIds = types.ListNull(types.Int64Type)

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.VlanIds != nil {
		vlanIds = mistutils.ListOfIntSdkToTerraform(d.VlanIds)
	}

	obj, e := basetypes.NewObjectValue(
		MdnsValue{}.AttributeTypes(ctx),
		map[string]attr.Value{
			"enabled":  enabled,
			"vlan_ids": vlanIds,
		},
	)
	diags.Append(e...)
	return obj
}

func tuntermMulticastSsdpSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteSettingTuntermMulticastConfigSsdp) types.Object {
	attrTypes := SsdpValue{}.AttributeTypes(ctx)

	if d == nil {
		return types.ObjectNull(attrTypes)
	}

	var enabled = types.BoolNull()
	var vlanIds = types.ListNull(types.Int64Type)

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.VlanIds != nil {
		vlanIds = mistutils.ListOfIntSdkToTerraform(d.VlanIds)
	}

	obj, e := basetypes.NewObjectValue(
		SsdpValue{}.AttributeTypes(ctx),
		map[string]attr.Value{
			"enabled":  enabled,
			"vlan_ids": vlanIds,
		},
	)
	diags.Append(e...)
	return obj
}
