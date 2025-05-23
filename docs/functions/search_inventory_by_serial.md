---
page_title: "search_inventory_by_serial function - terraform-provider-mist"
subcategory: "Devices"
description: |-
  Retrieve a device in the mist_org_inventory resource based on its Serial Number
---

# search_inventory_by_serial (function)

Given `mist_org_inventory` resource and a Serial Number string, will return the Device object having the provided Serial Number. The response object will contain all the information from the Mist Inventory:
* `claim_code`: Claim Code of the device 
* `deviceprofile_id`: deviceprofile id if assigned
* `hostname`: hostname reported by the device
* `id`: ID of the device
* `mac`: MAC Address of the device
* `model`: Model of the device
* `org_id`: Org ID of the device
* `serial`: Serial of the device
* `site_id`: Site ID of the device
* `type`: Type of device
* `unclaim_when_destroyed`: If the device will be unclaimed when removed from the `mist_org_inventory` resource
* `vc_mac`: only if `type`==`switch` of `type`==`gateway`, MAC Address of the Virtual Chassis Primary switch or the Gateway Cluster Master

-> The search function is case-insensitive

## Example Usage

```terraform
# result is a device object
resource "mist_device_ap" "test_ap" {
  device_id = provider::mist::search_inventory_by_serial(resource.mist_org_inventory.inventory, "A153420000000").id
  site_id   = provider::mist::search_inventory_by_serial(resource.mist_org_inventory.inventory, "A153420000000").site_id
  name      = "test_ap"
}
```

## Signature

<!-- signature generated by tfplugindocs -->
```text
search_inventory_by_serial(inventory object, serial string) object
```

## Arguments

<!-- arguments generated by tfplugindocs -->
1. `inventory` (Object) `mist_org_inventory` resource
1. `serial` (String) Device Serial

