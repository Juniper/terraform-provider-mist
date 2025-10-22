package provider

type OrgVpnModel struct {
	Name          string                      `hcl:"name"`
	OrgId         *string                     `hcl:"org_id"`
	PathSelection *PathSelectionValue         `hcl:"path_selection"`
	Paths         map[string]OrgVpnPathsValue `hcl:"paths"`
	Type          *string                     `hcl:"type"`
}

type PathSelectionValue struct {
	Strategy *string `cty:"strategy" hcl:"strategy"`
}

type OrgVpnPathsValue struct {
	BfdProfile       *string                    `cty:"bfd_profile" hcl:"bfd_profile"`
	BfdUseTunnelMode *bool                      `cty:"bfd_use_tunnel_mode" hcl:"bfd_use_tunnel_mode"`
	Ip               *string                    `cty:"ip" hcl:"ip"`
	PeerPaths        map[string]PeerPathsValue  `cty:"peer_paths" hcl:"peer_paths"`
	Pod              *int64                     `cty:"pod" hcl:"pod"`
	TrafficShaping   *OrgVpnTrafficShapingValue `cty:"traffic_shaping" hcl:"traffic_shaping"`
}

type PeerPathsValue struct {
	Preference *int64 `cty:"preference" hcl:"preference"`
}

type OrgVpnTrafficShapingValue struct {
	ClassPercentage []int64 `cty:"class_percentage" hcl:"class_percentage"`
	Enabled         *bool   `cty:"enabled" hcl:"enabled"`
	MaxTxKbps       *int64  `cty:"max_tx_kbps" hcl:"max_tx_kbps"`
}
