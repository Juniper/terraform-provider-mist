package provider

type SiteEvpnTopologyModel struct {
	EvpnOptions *SiteEvpnTopologyEvpnOptionsValue        `hcl:"evpn_options"`
	Name        string                                   `hcl:"name"`
	PodNames    map[string]string                        `hcl:"pod_names"`
	SiteId      string                                   `hcl:"site_id"`
	Switches    map[string]SiteEvpnTopologySwitchesValue `hcl:"switches"`
}

type SiteEvpnTopologyEvpnOptionsValue struct {
	AutoLoopbackSubnet  *string                                     `cty:"auto_loopback_subnet" hcl:"auto_loopback_subnet"`
	AutoLoopbackSubnet6 *string                                     `cty:"auto_loopback_subnet6" hcl:"auto_loopback_subnet6"`
	AutoRouterIdSubnet  *string                                     `cty:"auto_router_id_subnet" hcl:"auto_router_id_subnet"`
	AutoRouterIdSubnet6 *string                                     `cty:"auto_router_id_subnet6" hcl:"auto_router_id_subnet6"`
	CoreAsBorder        *bool                                       `cty:"core_as_border" hcl:"core_as_border"`
	EnableInbandZtp     *bool                                       `cty:"enable_inband_ztp" hcl:"enable_inband_ztp"`
	Overlay             *SiteEvpnTopologyOverlayValue               `cty:"overlay" hcl:"overlay"`
	PerVlanVgaV4Mac     *bool                                       `cty:"per_vlan_vga_v4_mac" hcl:"per_vlan_vga_v4_mac"`
	PerVlanVgaV6Mac     *bool                                       `cty:"per_vlan_vga_v6_mac" hcl:"per_vlan_vga_v6_mac"`
	RoutedAt            *string                                     `cty:"routed_at" hcl:"routed_at"`
	Underlay            *SiteEvpnTopologyUnderlayValue              `cty:"underlay" hcl:"underlay"`
	VsInstances         map[string]SiteEvpnTopologyVsInstancesValue `cty:"vs_instances" hcl:"vs_instances"`
}

type SiteEvpnTopologyOverlayValue struct {
	As *int64 `cty:"as" hcl:"as"`
}

type SiteEvpnTopologyUnderlayValue struct {
	AsBase         *int64  `cty:"as_base" hcl:"as_base"`
	RoutedIdPrefix *string `cty:"routed_id_prefix" hcl:"routed_id_prefix"`
	Subnet         *string `cty:"subnet" hcl:"subnet"`
	UseIpv6        *bool   `cty:"use_ipv6" hcl:"use_ipv6"`
}

type SiteEvpnTopologyVsInstancesValue struct {
	Networks []string `cty:"networks" hcl:"networks"`
}

type SiteEvpnTopologySwitchesValue struct {
	Pod  *int64  `cty:"pod" hcl:"pod"`
	Pods []int64 `cty:"pods" hcl:"pods"`
	Role string  `cty:"role" hcl:"role"`
}
