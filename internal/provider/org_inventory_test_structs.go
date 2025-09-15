package provider

type OrgInventoryModel struct {
	Inventory map[string]InventoryValue `hcl:"inventory"`
	OrgId     string                    `hcl:"org_id"`
}

type DevicesValue struct {
	Mac                  *string `cty:"mac" hcl:"mac"`
	Magic                *string `cty:"claim_code" hcl:"claim_code"`
	SiteId               *string `cty:"site_id" hcl:"site_id"`
	UnclaimWhenDestroyed *bool   `cty:"unclaim_when_destroyed" hcl:"unclaim_when_destroyed"`
}

type InventoryValue struct {
	// All inventory nested fields for comprehensive testing
	SiteId               *string `cty:"site_id" hcl:"site_id"`
	UnclaimWhenDestroyed *bool   `cty:"unclaim_when_destroyed" hcl:"unclaim_when_destroyed"`
}
