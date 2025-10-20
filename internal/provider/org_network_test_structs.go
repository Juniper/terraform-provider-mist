package provider

type OrgNetworkModel struct {
	DisallowMistServices *bool                               `hcl:"disallow_mist_services"`
	Gateway              *string                             `hcl:"gateway"`
	Gateway6             *string                             `hcl:"gateway6"`
	InternalAccess       *OrgNetworkInternalAccessValue      `hcl:"internal_access"`
	InternetAccess       *OrgNetworkInternetAccessValue      `hcl:"internet_access"`
	Isolation            *bool                               `hcl:"isolation"`
	Multicast            *OrgNetworkMulticastValue           `hcl:"multicast"`
	Name                 string                              `hcl:"name"`
	OrgId                string                              `hcl:"org_id"`
	RoutedForNetworks    []string                            `hcl:"routed_for_networks"`
	Subnet               string                              `hcl:"subnet"`
	Subnet6              *string                             `hcl:"subnet6"`
	Tenants              map[string]OrgNetworkTenantsValue   `hcl:"tenants"`
	VlanId               *string                             `hcl:"vlan_id"`
	VpnAccess            map[string]OrgNetworkVpnAccessValue `hcl:"vpn_access"`
}

type OrgNetworkInternalAccessValue struct {
	Enabled *bool `cty:"enabled" hcl:"enabled"`
}

type OrgNetworkInternetAccessValue struct {
	CreateSimpleServicePolicy    *bool                                                  `cty:"create_simple_service_policy" hcl:"create_simple_service_policy"`
	Enabled                      *bool                                                  `cty:"enabled" hcl:"enabled"`
	InternetAccessDestinationNat map[string]OrgNetworkInternetAccessDestinationNatValue `cty:"destination_nat" hcl:"destination_nat"`
	InternetAccessStaticNat      map[string]OrgNetworkInternetAccessStaticNatValue      `cty:"static_nat" hcl:"static_nat"`
	Restricted                   *bool                                                  `cty:"restricted" hcl:"restricted"`
}

type OrgNetworkInternetAccessDestinationNatValue struct {
	InternalIp *string `cty:"internal_ip" hcl:"internal_ip"`
	Name       *string `cty:"name" hcl:"name"`
	Port       *string `cty:"port" hcl:"port"`
	WanName    *string `cty:"wan_name" hcl:"wan_name"`
}

type OrgNetworkInternetAccessStaticNatValue struct {
	InternalIp string  `cty:"internal_ip" hcl:"internal_ip"`
	Name       string  `cty:"name" hcl:"name"`
	WanName    *string `cty:"wan_name" hcl:"wan_name"`
}

type OrgNetworkMulticastValue struct {
	DisableIgmp *bool                            `cty:"disable_igmp" hcl:"disable_igmp"`
	Enabled     *bool                            `cty:"enabled" hcl:"enabled"`
	Groups      map[string]OrgNetworkGroupsValue `cty:"groups" hcl:"groups"`
}

type OrgNetworkGroupsValue struct {
	RpIp *string `cty:"rp_ip" hcl:"rp_ip"`
}

type OrgNetworkTenantsValue struct {
	Addresses []string `cty:"addresses" hcl:"addresses"`
}

type OrgNetworkVpnAccessValue struct {
	AdvertisedSubnet          *string                                           `cty:"advertised_subnet" hcl:"advertised_subnet"`
	AllowPing                 *bool                                             `cty:"allow_ping" hcl:"allow_ping"`
	NatPool                   *string                                           `cty:"nat_pool" hcl:"nat_pool"`
	NoReadvertiseToLanBgp     *bool                                             `cty:"no_readvertise_to_lan_bgp" hcl:"no_readvertise_to_lan_bgp"`
	NoReadvertiseToLanOspf    *bool                                             `cty:"no_readvertise_to_lan_ospf" hcl:"no_readvertise_to_lan_ospf"`
	NoReadvertiseToOverlay    *bool                                             `cty:"no_readvertise_to_overlay" hcl:"no_readvertise_to_overlay"`
	OtherVrfs                 []string                                          `cty:"other_vrfs" hcl:"other_vrfs"`
	Routed                    *bool                                             `cty:"routed" hcl:"routed"`
	SourceNat                 *OrgNetworkSourceNatValue                         `cty:"source_nat" hcl:"source_nat"`
	SummarizedSubnet          *string                                           `cty:"summarized_subnet" hcl:"summarized_subnet"`
	SummarizedSubnetToLanBgp  *string                                           `cty:"summarized_subnet_to_lan_bgp" hcl:"summarized_subnet_to_lan_bgp"`
	SummarizedSubnetToLanOspf *string                                           `cty:"summarized_subnet_to_lan_ospf" hcl:"summarized_subnet_to_lan_ospf"`
	VpnAccessDestinationNat   map[string]OrgNetworkVpnAccessDestinationNatValue `cty:"destination_nat" hcl:"destination_nat"`
	VpnAccessStaticNat        map[string]OrgNetworkVpnAccessStaticNatValue      `cty:"static_nat" hcl:"static_nat"`
}

type OrgNetworkSourceNatValue struct {
	ExternalIp *string `cty:"external_ip" hcl:"external_ip"`
}

type OrgNetworkVpnAccessDestinationNatValue struct {
	InternalIp *string `cty:"internal_ip" hcl:"internal_ip"`
	Name       *string `cty:"name" hcl:"name"`
	Port       *string `cty:"port" hcl:"port"`
}

type OrgNetworkVpnAccessStaticNatValue struct {
	InternalIp string `cty:"internal_ip" hcl:"internal_ip"`
	Name       string `cty:"name" hcl:"name"`
}
