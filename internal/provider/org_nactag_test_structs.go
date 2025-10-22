package provider

type OrgNactagModel struct {
	AllowUsermacOverride *bool    `hcl:"allow_usermac_override"`
	EgressVlanNames      []string `hcl:"egress_vlan_names"`
	GbpTag               *string  `hcl:"gbp_tag"`
	Match                *string  `hcl:"match"`
	MatchAll             *bool    `hcl:"match_all"`
	NacportalId          *string  `hcl:"nacportal_id"`
	Name                 string   `hcl:"name"`
	OrgId                string   `hcl:"org_id"`
	RadiusAttrs          []string `hcl:"radius_attrs"`
	RadiusGroup          *string  `hcl:"radius_group"`
	RadiusVendorAttrs    []string `hcl:"radius_vendor_attrs"`
	SessionTimeout       *int64   `hcl:"session_timeout"`
	Type                 string   `hcl:"type"`
	UsernameAttr         *string  `hcl:"username_attr"`
	Values               []string `hcl:"values"`
	Vlan                 *string  `hcl:"vlan"`
}
