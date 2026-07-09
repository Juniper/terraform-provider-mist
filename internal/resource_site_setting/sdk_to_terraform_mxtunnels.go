package resource_site_setting

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func mxtunnelsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteMxtunnel) MxtunnelsValue {
	var additionalMxtunnels = types.MapNull(AdditionalMxtunnelsValue{}.Type(ctx))
	var apSubnets = types.ListNull(types.StringType)
	var autoPreemption = types.ObjectNull(AutoPreemptionValue{}.AttributeTypes(ctx))
	var clusters = types.ListNull(ClustersValue{}.Type(ctx))
	var createdTime = types.Float64Null()
	var enabled = types.BoolNull()
	var forSite = types.BoolNull()
	var helloInterval = types.Int64Null()
	var helloRetries = types.Int64Null()
	var hosts = types.ListNull(types.StringType)
	var id = types.StringNull()
	var modifiedTime = types.Float64Null()
	var mtu = types.Int64Null()
	var orgId = types.StringNull()
	var protocol = types.StringNull()
	var radsec = types.ObjectNull(RadsecValue{}.AttributeTypes(ctx))
	var siteId = types.StringNull()
	var vlanIds = types.ListNull(types.Int64Type)

	if d.AdditionalMxtunnels != nil {
		items := make(map[string]attr.Value)
		for k, v := range d.AdditionalMxtunnels {
			item := mxtunnelsAdditionalSdkToTerraform(ctx, diags, v)
			items[k] = item
		}
		r, e := types.MapValueFrom(ctx, AdditionalMxtunnelsValue{}.Type(ctx), items)
		diags.Append(e...)
		additionalMxtunnels = r
	}

	if d.ApSubnets != nil {
		apSubnets = mistutils.ListOfStringSdkToTerraform(d.ApSubnets)
	}

	if d.AutoPreemption != nil {
		autoPreemption = mxtunnelsAutoPreemptionSdkToTerraform(ctx, diags, d.AutoPreemption)
	}

	if d.Clusters != nil {
		var items []attr.Value
		for _, item := range d.Clusters {
			var name = types.StringNull()
			var tuntermHosts = types.ListNull(types.StringType)
			if item.Name != nil {
				name = types.StringValue(*item.Name)
			}
			if item.TuntermHosts != nil {
				tuntermHosts = mistutils.ListOfStringSdkToTerraform(item.TuntermHosts)
			}
			itemAttrType := ClustersValue{}.AttributeTypes(ctx)
			itemObj, e := NewClustersValue(itemAttrType, map[string]attr.Value{
				"name":          name,
				"tunterm_hosts": tuntermHosts,
			})
			diags.Append(e...)
			items = append(items, itemObj)
		}
		r, e := types.ListValueFrom(ctx, ClustersValue{}.Type(ctx), items)
		diags.Append(e...)
		clusters = r
	}

	if d.CreatedTime != nil {
		createdTime = types.Float64Value(*d.CreatedTime)
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.ForSite != nil {
		forSite = types.BoolValue(*d.ForSite)
	}
	if d.HelloInterval != nil {
		helloInterval = types.Int64Value(int64(*d.HelloInterval))
	}
	if d.HelloRetries != nil {
		helloRetries = types.Int64Value(int64(*d.HelloRetries))
	}
	if d.Hosts != nil {
		hosts = mistutils.ListOfStringSdkToTerraform(d.Hosts)
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.ModifiedTime != nil {
		modifiedTime = types.Float64Value(*d.ModifiedTime)
	}
	if d.Mtu != nil {
		mtu = types.Int64Value(int64(*d.Mtu))
	}
	if d.OrgId != nil {
		orgId = types.StringValue(d.OrgId.String())
	}
	if d.Protocol != nil {
		protocol = types.StringValue(string(*d.Protocol))
	}
	if d.Radsec != nil {
		radsec = mxtunnelsRadsecSdkToTerraform(ctx, diags, d.Radsec)
	}
	if d.SiteId != nil {
		siteId = types.StringValue(d.SiteId.String())
	}
	if d.VlanIds != nil {
		vlanIds = mistutils.ListOfIntSdkToTerraform(d.VlanIds)
	}

	dataMapAttrType := MxtunnelsValue{}.AttributeTypes(ctx)
	dataMapValue := map[string]attr.Value{
		"additional_mxtunnels": additionalMxtunnels,
		"ap_subnets":           apSubnets,
		"auto_preemption":      autoPreemption,
		"clusters":             clusters,
		"created_time":         createdTime,
		"enabled":              enabled,
		"for_site":             forSite,
		"hello_interval":       helloInterval,
		"hello_retries":        helloRetries,
		"hosts":                hosts,
		"id":                   id,
		"modified_time":        modifiedTime,
		"mtu":                  mtu,
		"org_id":               orgId,
		"protocol":             protocol,
		"radsec":               radsec,
		"site_id":              siteId,
		"vlan_ids":             vlanIds,
	}
	data, e := NewMxtunnelsValue(dataMapAttrType, dataMapValue)
	diags.Append(e...)

	return data
}

func mxtunnelsAdditionalSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.SiteMxtunnelAdditionalMxtunnel) AdditionalMxtunnelsValue {
	var helloInterval = types.Int64Null()
	var helloRetries = types.Int64Null()
	var protocol = types.StringNull()
	var tuntermClusters = types.ListNull(TuntermClustersValue{}.Type(ctx))
	var vlanIds = types.ListNull(types.Int64Type)

	if d.HelloInterval != nil {
		helloInterval = types.Int64Value(int64(*d.HelloInterval))
	}
	if d.HelloRetries != nil {
		helloRetries = types.Int64Value(int64(*d.HelloRetries))
	}
	if d.Protocol != nil {
		protocol = types.StringValue(string(*d.Protocol))
	}
	if d.Clusters != nil {
		var items []attr.Value
		for _, item := range d.Clusters {
			var name = types.StringNull()
			var tuntermHosts = types.ListNull(types.StringType)
			if item.Name != nil {
				name = types.StringValue(*item.Name)
			}
			if item.TuntermHosts != nil {
				tuntermHosts = mistutils.ListOfStringSdkToTerraform(item.TuntermHosts)
			}
			itemObj, e := NewTuntermClustersValue(
				TuntermClustersValue{}.AttributeTypes(ctx),
				map[string]attr.Value{"name": name, "tunterm_hosts": tuntermHosts},
			)
			diags.Append(e...)
			items = append(items, itemObj)
		}
		r, e := types.ListValueFrom(ctx, TuntermClustersValue{}.Type(ctx), items)
		diags.Append(e...)
		tuntermClusters = r
	}
	if d.VlanIds != nil {
		vlanIds = mistutils.ListOfIntSdkToTerraform(d.VlanIds)
	}

	item, e := NewAdditionalMxtunnelsValue(
		AdditionalMxtunnelsValue{}.AttributeTypes(ctx),
		map[string]attr.Value{
			"hello_interval":   helloInterval,
			"hello_retries":    helloRetries,
			"protocol":         protocol,
			"tunterm_clusters": tuntermClusters,
			"vlan_ids":         vlanIds,
		},
	)
	diags.Append(e...)
	return item
}

func mxtunnelsAutoPreemptionSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.AutoPreemption) types.Object {
	attrTypes := AutoPreemptionValue{}.AttributeTypes(ctx)
	if d == nil {
		return types.ObjectNull(attrTypes)
	}

	var dayOfWeek = types.StringNull()
	var enabled = types.BoolNull()
	var timeOfDay = types.StringNull()

	if d.DayOfWeek != nil {
		dayOfWeek = types.StringValue(string(*d.DayOfWeek))
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.TimeOfDay != nil {
		timeOfDay = types.StringValue(*d.TimeOfDay)
	}

	v, e := NewAutoPreemptionValue(
		AutoPreemptionValue{}.AttributeTypes(ctx),
		map[string]attr.Value{
			"day_of_week": dayOfWeek,
			"enabled":     enabled,
			"time_of_day": timeOfDay,
		},
	)
	diags.Append(e...)

	obj, e := v.ToObjectValue(ctx)
	diags.Append(e...)
	return obj
}

func mxtunnelsRadsecSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteMxtunnelRadsec) types.Object {
	attrTypes := RadsecValue{}.AttributeTypes(ctx)
	if d == nil {
		return types.ObjectNull(attrTypes)
	}

	var acctServers = types.ListNull(AcctServersValue{}.Type(ctx))
	var authServers = types.ListNull(AuthServersValue{}.Type(ctx))
	var enabled = types.BoolNull()
	var useMxedge = types.BoolNull()

	if d.AcctServers != nil {
		var items []attr.Value
		for _, item := range d.AcctServers {
			port := types.StringNull()
			if item.Port != nil {
				if n, ok := item.Port.AsNumber(); ok {
					port = types.StringValue(fmt.Sprintf("%d", *n))
				} else if s, ok2 := item.Port.AsString(); ok2 {
					port = types.StringValue(*s)
				}
			}
			itemAttrType := AcctServersValue{}.AttributeTypes(ctx)
			itemObj, e := NewAcctServersValue(itemAttrType, map[string]attr.Value{
				"host":            types.StringValue(item.Host),
				"keywrap_enabled": types.BoolNull(),
				"keywrap_format":  types.StringNull(),
				"keywrap_kek":     types.StringNull(),
				"keywrap_mack":    types.StringNull(),
				"port":            port,
				"secret":          types.StringValue(item.Secret),
			})
			diags.Append(e...)
			items = append(items, itemObj)
		}
		r, e := types.ListValueFrom(ctx, AcctServersValue{}.Type(ctx), items)
		diags.Append(e...)
		acctServers = r
	}

	if d.AuthServers != nil {
		var items []attr.Value
		for _, item := range d.AuthServers {
			port := types.StringNull()
			if item.Port != nil {
				if n, ok := item.Port.AsNumber(); ok {
					port = types.StringValue(fmt.Sprintf("%d", *n))
				} else if s, ok2 := item.Port.AsString(); ok2 {
					port = types.StringValue(*s)
				}
			}
			keywrapFormat := types.StringNull()
			if item.KeywrapFormat != nil {
				keywrapFormat = types.StringValue(string(*item.KeywrapFormat))
			}
			requireMessageAuthenticator := types.BoolNull()
			if item.RequireMessageAuthenticator != nil {
				requireMessageAuthenticator = types.BoolValue(*item.RequireMessageAuthenticator)
			}
			itemsecret := types.StringValue(item.Secret)
			itemAttrType := AuthServersValue{}.AttributeTypes(ctx)
			itemObj, e := NewAuthServersValue(itemAttrType, map[string]attr.Value{
				"host": types.StringValue(item.Host),
				"keywrap_enabled": func() types.Bool {
					if item.KeywrapEnabled != nil {
						return types.BoolValue(*item.KeywrapEnabled)
					}
					return types.BoolNull()
				}(),
				"keywrap_format": keywrapFormat,
				"keywrap_kek": func() types.String {
					if item.KeywrapKek != nil {
						return types.StringValue(*item.KeywrapKek)
					}
					return types.StringNull()
				}(),
				"keywrap_mack": func() types.String {
					if item.KeywrapMack != nil {
						return types.StringValue(*item.KeywrapMack)
					}
					return types.StringNull()
				}(),
				"port":                          port,
				"require_message_authenticator": requireMessageAuthenticator,
				"secret":                        itemsecret,
			})
			diags.Append(e...)
			items = append(items, itemObj)
		}
		r, e := types.ListValueFrom(ctx, AuthServersValue{}.Type(ctx), items)
		diags.Append(e...)
		authServers = r
	}

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.UseMxedge != nil {
		useMxedge = types.BoolValue(*d.UseMxedge)
	}

	v, e := NewRadsecValue(
		RadsecValue{}.AttributeTypes(ctx),
		map[string]attr.Value{
			"acct_servers": acctServers,
			"auth_servers": authServers,
			"enabled":      enabled,
			"use_mxedge":   useMxedge,
		},
	)
	diags.Append(e...)

	obj, e := v.ToObjectValue(ctx)
	diags.Append(e...)
	return obj
}
