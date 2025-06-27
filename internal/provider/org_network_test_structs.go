package provider

type OrgNetworkModel struct {
	DisallowMistServices *bool                     `hcl:"disallow_mist_services"`
	Gateway              *string                   `hcl:"gateway"`
	Gateway6             *string                   `hcl:"gateway6"`
	InternalAccess       *InternalAccessValue      `hcl:"internal_access"`
	InternetAccess       *InternetAccessValue      `hcl:"internet_access"`
	Isolation            *bool                     `hcl:"isolation"`
	Multicast            *MulticastValue           `hcl:"multicast"`
	Name                 string                    `hcl:"name"`
	OrgId                string                    `hcl:"org_id"`
	RoutedForNetworks    []string                  `hcl:"routed_for_networks"`
	Subnet               string                    `hcl:"subnet"`
	Subnet6              *string                   `hcl:"subnet6"`
	Tenants              map[string]TenantsValue   `hcl:"tenants"`
	VlanId               *string                   `hcl:"vlan_id"`
	VpnAccess            map[string]VpnAccessValue `hcl:"vpn_access"`
}
