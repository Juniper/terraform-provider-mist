package provider

type SiteEvpnTopologyModel struct {
	EvpnOptions *EvpnOptionsValue        `hcl:"evpn_options"`
	Name        string                   `hcl:"name"`
	PodNames    map[string]string        `hcl:"pod_names"`
	SiteId      string                   `hcl:"site_id"`
	Switches    map[string]SwitchesValue `hcl:"switches"`
}
