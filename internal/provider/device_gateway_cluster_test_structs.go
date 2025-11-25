package provider

import ()

type DeviceGatewayClusterModel struct {
	Nodes  []NodesValue `hcl:"nodes"`
	SiteId string       `hcl:"site_id"`
}

type NodesValue struct {
	Mac string `cty:"mac"`
}
