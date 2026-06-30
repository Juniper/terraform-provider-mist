package resource_org_mxcluster

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func mistNacedgeSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.MxclusterNacedge) MistNacedgeValue {

	var authTtl = types.Int64Null()
	var cachingSiteIds = types.ListNull(types.StringType)
	var defaultDot1xVlan = types.StringNull()
	var defaultVlan = types.StringNull()
	var enabled = types.BoolNull()
	var nacEdgeHosts = types.ListNull(types.StringType)

	if d.AuthTtl != nil {
		authTtl = types.Int64Value(int64(*d.AuthTtl))
	}
	if d.CachingSiteIds != nil {
		var items []string
		for _, id := range d.CachingSiteIds {
			items = append(items, id.String())
		}
		cachingSiteIds = mistutils.ListOfStringSdkToTerraform(items)
	}
	if d.DefaultDot1xVlan != nil {
		defaultDot1xVlan = types.StringValue(*d.DefaultDot1xVlan)
	}
	if d.DefaultVlan != nil {
		defaultVlan = types.StringValue(*d.DefaultVlan)
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.NacEdgeHosts != nil {
		nacEdgeHosts = mistutils.ListOfStringSdkToTerraform(d.NacEdgeHosts)
	}

	data_map_attr_type := MistNacedgeValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"auth_ttl":           authTtl,
		"caching_site_ids":   cachingSiteIds,
		"default_dot1x_vlan": defaultDot1xVlan,
		"default_vlan":       defaultVlan,
		"enabled":            enabled,
		"nac_edge_hosts":     nacEdgeHosts,
	}
	data, e := NewMistNacedgeValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
