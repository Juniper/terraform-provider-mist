package provider

type OrgEvpnTopologyModel struct {
	EvpnOptions *EvpnOptionsValue        `hcl:"evpn_options"`
	Name        string                   `hcl:"name"`
	OrgId       string                   `hcl:"org_id"`
	PodNames    map[string]string        `hcl:"pod_names"`
	Switches    map[string]SwitchesValue `hcl:"switches"`
}

type EvpnOptionsValue struct {
	AutoLoopbackSubnet  *string                     `cty:"auto_loopback_subnet"`
	AutoLoopbackSubnet6 *string                     `cty:"auto_loopback_subnet6"`
	AutoRouterIdSubnet  *string                     `cty:"auto_router_id_subnet"`
	AutoRouterIdSubnet6 *string                     `cty:"auto_router_id_subnet6"`
	CoreAsBorder        *bool                       `cty:"core_as_border"`
	Overlay             *OverlayValue               `cty:"overlay"`
	PerVlanVgaV4Mac     *bool                       `cty:"per_vlan_vga_v4_mac"`
	RoutedAt            *string                     `cty:"routed_at"`
	Underlay            *UnderlayValue              `cty:"underlay"`
	VsInstances         map[string]VsInstancesValue `cty:"vs_instances"`
}

type OverlayValue struct {
	As *int64 `cty:"as"`
}

type UnderlayValue struct {
	AsBase         *int64  `cty:"as_base"`
	RoutedIdPrefix *string `cty:"routed_id_prefix"`
	Subnet         *string `cty:"subnet"`
	UseIpv6        *bool   `cty:"use_ipv6"`
}

type VsInstancesValue struct {
	Networks []string `cty:"networks"`
}

type SwitchesValue struct {
	Pod  *int64  `cty:"pod"`
	Pods []int64 `cty:"pods"`
	Role string  `cty:"role"`
}
