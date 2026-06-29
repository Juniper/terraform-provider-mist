package provider

type OrgInventoryModel struct {
	DisconnectedBefore *int64                                `hcl:"disconnected_before"`
	Inventory          map[string]OrgInventoryInventoryValue `hcl:"inventory"`
	OrgId              string                                `hcl:"org_id"`
}

type OrgInventoryInventoryValue struct {
	SiteId               *string `cty:"site_id" hcl:"site_id"`
	UnclaimWhenDestroyed *bool   `cty:"unclaim_when_destroyed" hcl:"unclaim_when_destroyed"`
}
