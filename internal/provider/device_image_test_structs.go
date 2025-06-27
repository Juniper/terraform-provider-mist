package provider

import ()

type DeviceImageModel struct {
	DeviceId string `hcl:"device_id"`
	File string `hcl:"file"`
	ImageNumber int64 `hcl:"image_number"`
	SiteId string `hcl:"site_id"`
}

