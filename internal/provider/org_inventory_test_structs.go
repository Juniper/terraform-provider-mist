package provider

type OrgInventoryModel struct {
	Inventory map[string]InventoryValue `hcl:"inventory"`
	OrgId     string                    `hcl:"org_id"`
}

type InventoryValue struct {
	SiteId               *string `cty:"site_id" hcl:"site_id"`
	UnclaimWhenDestroyed *bool   `cty:"unclaim_when_destroyed" hcl:"unclaim_when_destroyed"`
}
