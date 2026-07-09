package provider

type OrgMxtunnelModel struct {
	AnchorMxtunnelIds []string                        `hcl:"anchor_mxtunnel_ids"`
	AutoPreemption    *OrgMxtunnelAutoPreemptionValue `hcl:"auto_preemption"`
	HelloInterval     *int64                          `hcl:"hello_interval"`
	HelloRetries      *int64                          `hcl:"hello_retries"`
	Ipsec             *OrgMxtunnelIpsecValue          `hcl:"ipsec"`
	Mtu               *int64                          `hcl:"mtu"`
	MxclusterIds      []string                        `hcl:"mxcluster_ids"`
	Name              string                          `hcl:"name"`
	OrgId             string                          `hcl:"org_id"`
	Protocol          *string                         `hcl:"protocol"`
	VlanIds           []int64                         `hcl:"vlan_ids"`
}

type OrgMxtunnelAutoPreemptionValue struct {
	DayOfWeek *string `cty:"day_of_week" hcl:"day_of_week"`
	Enabled   *bool   `cty:"enabled" hcl:"enabled"`
	TimeOfDay *string `cty:"time_of_day" hcl:"time_of_day"`
}

type OrgMxtunnelIpsecValue struct {
	DnsServers  []string                      `cty:"dns_servers" hcl:"dns_servers"`
	DnsSuffix   []string                      `cty:"dns_suffix" hcl:"dns_suffix"`
	Enabled     *bool                         `cty:"enabled" hcl:"enabled"`
	ExtraRoutes []OrgMxtunnelExtraRoutesValue `cty:"extra_routes" hcl:"extra_routes"`
	SplitTunnel *bool                         `cty:"split_tunnel" hcl:"split_tunnel"`
	UseMxedge   *bool                         `cty:"use_mxedge" hcl:"use_mxedge"`
}

type OrgMxtunnelExtraRoutesValue struct {
	Dest    *string `cty:"dest" hcl:"dest"`
	NextHop *string `cty:"next_hop" hcl:"next_hop"`
}
