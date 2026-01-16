package resource_org_mxedge

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func tuntermMulticastConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.MxedgeTuntermMulticastConfig) TuntermMulticastConfigValue {

	var mdns = types.ObjectNull(MdnsValue{}.AttributeTypes(ctx))
	var ssdp = types.ObjectNull(SsdpValue{}.AttributeTypes(ctx))

	if d.Mdns != nil {
		mdnsValue := mdnsSdkToTerraform(ctx, diags, d.Mdns)
		mdnsObj, e := mdnsValue.ToObjectValue(ctx)
		diags.Append(e...)
		mdns = mdnsObj
	}
	if d.Ssdp != nil {
		ssdpValue := ssdpSdkToTerraform(ctx, diags, d.Ssdp)
		ssdpObj, e := ssdpValue.ToObjectValue(ctx)
		diags.Append(e...)
		ssdp = ssdpObj
	}

	data_map_attr_type := TuntermMulticastConfigValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"mdns": mdns,
		"ssdp": ssdp,
	}
	data, e := NewTuntermMulticastConfigValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func mdnsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.MxedgeTuntermMulticastMdns) MdnsValue {

	var enabled types.Bool
	var vlanIds = types.ListNull(types.StringType)

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.VlanIds != nil {
		vlanIds = mistutils.ListOfStringSdkToTerraform(d.VlanIds)
	}

	data_map_attr_type := MdnsValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"enabled":  enabled,
		"vlan_ids": vlanIds,
	}
	data, e := NewMdnsValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func ssdpSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.MxedgeTuntermMulticastSsdp) SsdpValue {

	var enabled types.Bool
	var vlanIds = types.ListNull(types.StringType)

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.VlanIds != nil {
		vlanIds = mistutils.ListOfStringSdkToTerraform(d.VlanIds)
	}

	data_map_attr_type := SsdpValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"enabled":  enabled,
		"vlan_ids": vlanIds,
	}
	data, e := NewSsdpValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
