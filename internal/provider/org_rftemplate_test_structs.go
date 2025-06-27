package provider

type OrgRftemplateModel struct {
	AntGain24       *int64                            `hcl:"ant_gain_24"`
	AntGain5        *int64                            `hcl:"ant_gain_5"`
	AntGain6        *int64                            `hcl:"ant_gain_6"`
	Band24          *OrgRftemplateBand24Value         `hcl:"band_24"`
	Band24Usage     *string                           `hcl:"band_24_usage"`
	Band5           *OrgRftemplateBand5Value          `hcl:"band_5"`
	Band5On24Radio  *OrgRftemplateBand5On24RadioValue `hcl:"band_5_on_24_radio"`
	Band6           *OrgRftemplateBand6Value          `hcl:"band_6"`
	CountryCode     *string                           `hcl:"country_code"`
	ModelSpecific   map[string]ModelSpecificValue     `hcl:"model_specific"`
	Name            string                            `hcl:"name"`
	OrgId           string                            `hcl:"org_id"`
	ScanningEnabled *bool                             `hcl:"scanning_enabled"`
}

type OrgRftemplateBand24Value struct {
	AllowRrmDisable *bool   `cty:"allow_rrm_disable"`
	AntGain         *int64  `cty:"ant_gain"`
	AntennaMode     *string `cty:"antenna_mode"`
	Bandwidth       *int64  `cty:"bandwidth"`
	Channels        []int64 `cty:"channels"`
	Disabled        *bool   `cty:"disabled"`
	Power           *int64  `cty:"power"`
	PowerMax        *int64  `cty:"power_max"`
	PowerMin        *int64  `cty:"power_min"`
	Preamble        *string `cty:"preamble"`
}

type OrgRftemplateBand5Value struct {
	AllowRrmDisable *bool   `cty:"allow_rrm_disable"`
	AntGain         *int64  `cty:"ant_gain"`
	AntennaMode     *string `cty:"antenna_mode"`
	Bandwidth       *int64  `cty:"bandwidth"`
	Channels        []int64 `cty:"channels"`
	Disabled        *bool   `cty:"disabled"`
	Power           *int64  `cty:"power"`
	PowerMax        *int64  `cty:"power_max"`
	PowerMin        *int64  `cty:"power_min"`
	Preamble        *string `cty:"preamble"`
}

type OrgRftemplateBand5On24RadioValue struct {
	AllowRrmDisable *bool   `cty:"allow_rrm_disable"`
	AntGain         *int64  `cty:"ant_gain"`
	AntennaMode     *string `cty:"antenna_mode"`
	Bandwidth       *int64  `cty:"bandwidth"`
	Channels        []int64 `cty:"channels"`
	Disabled        *bool   `cty:"disabled"`
	Power           *int64  `cty:"power"`
	PowerMax        *int64  `cty:"power_max"`
	PowerMin        *int64  `cty:"power_min"`
	Preamble        *string `cty:"preamble"`
}

type OrgRftemplateBand6Value struct {
	AllowRrmDisable *bool   `cty:"allow_rrm_disable"`
	AntGain         *int64  `cty:"ant_gain"`
	AntennaMode     *string `cty:"antenna_mode"`
	Bandwidth       *int64  `cty:"bandwidth"`
	Channels        []int64 `cty:"channels"`
	Disabled        *bool   `cty:"disabled"`
	Power           *int64  `cty:"power"`
	PowerMax        *int64  `cty:"power_max"`
	PowerMin        *int64  `cty:"power_min"`
	Preamble        *string `cty:"preamble"`
	StandardPower   *bool   `cty:"standard_power"`
}

type ModelSpecificValue struct {
	AntGain24      *int64               `cty:"ant_gain_24"`
	AntGain5       *int64               `cty:"ant_gain_5"`
	AntGain6       *int64               `cty:"ant_gain_6"`
	Band24         *Band24Value         `cty:"band_24"`
	Band24Usage    *string              `cty:"band_24_usage"`
	Band5          *Band5Value          `cty:"band_5"`
	Band5On24Radio *Band5On24RadioValue `cty:"band_5_on_24_radio"`
	Band6          *Band6Value          `cty:"band_6"`
}
