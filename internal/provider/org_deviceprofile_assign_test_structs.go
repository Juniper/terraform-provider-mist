package provider

type OrgDeviceprofileAssignModel struct {
	DeviceprofileId string   `hcl:"deviceprofile_id"`
	Macs            []string `hcl:"macs"` // Using []string for types.Set results in the same hcl configuration
	OrgId           string   `hcl:"org_id"`
}
