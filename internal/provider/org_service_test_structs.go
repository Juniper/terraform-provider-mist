package provider

type OrgServiceModel struct {
	Addresses                     []string               `hcl:"addresses"`
	AppCategories                 []string               `hcl:"app_categories"`
	AppSubcategories              []string               `hcl:"app_subcategories"`
	Apps                          []string               `hcl:"apps"`
	ClientLimitDown               *int64                 `hcl:"client_limit_down"`
	ClientLimitUp                 *int64                 `hcl:"client_limit_up"`
	Description                   *string                `hcl:"description"`
	Dscp                          *string                `hcl:"dscp"`
	FailoverPolicy                *string                `hcl:"failover_policy"`
	Hostnames                     []string               `hcl:"hostnames"`
	MaxJitter                     *string                `hcl:"max_jitter"`
	MaxLatency                    *string                `hcl:"max_latency"`
	MaxLoss                       *string                `hcl:"max_loss"`
	Name                          string                 `hcl:"name"`
	OrgId                         string                 `hcl:"org_id"`
	ServiceLimitDown              *int64                 `hcl:"service_limit_down"`
	ServiceLimitUp                *int64                 `hcl:"service_limit_up"`
	SleEnabled                    *bool                  `hcl:"sle_enabled"`
	Specs                         []OrgServiceSpecsValue `hcl:"specs"`
	SsrRelaxedTcpStateEnforcement *bool                  `hcl:"ssr_relaxed_tcp_state_enforcement"`
	TrafficClass                  *string                `hcl:"traffic_class"`
	TrafficType                   *string                `hcl:"traffic_type"`
	Type                          *string                `hcl:"type"`
	Urls                          []string               `hcl:"urls"`
}

type OrgServiceSpecsValue struct {
	PortRange *string `cty:"port_range" hcl:"port_range"`
	Protocol  *string `cty:"protocol" hcl:"protocol"`
}
