package provider

import ()

type OrgInventoryModel struct {
	Devices []DevicesValue `hcl:"devices"`
	Inventory map[string]InventoryValue `hcl:"inventory"`
	OrgId string `hcl:"org_id"`
}

type DevicesValue struct {
	Mac *string `cty:"mac"`
	Magic *string `cty:"claim_code"`
	SiteId *string `cty:"site_id"`
	UnclaimWhenDestroyed *bool `cty:"unclaim_when_destroyed"`
}

type InventoryValue struct {
	SiteId *string `cty:"site_id"`
	UnclaimWhenDestroyed *bool `cty:"unclaim_when_destroyed"`
}

