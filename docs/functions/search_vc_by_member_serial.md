---
page_title: "search_vc_by_member_serial function - terraform-provider-mist"
subcategory: "Devices"
description: |-
  Retrieve a Virtual Chassis in the mist_org_inventory resource based on one of its member Serial Number
---

# search_vc_by_member_serial (function)

Given `mist_org_inventory` resource and a Serial Number string, will return the Device object of the Virtual Chassis having one of its member with the provided Serial Number.  
If the provided Serial Number belongs to a device that is not part of a Virtual Chassis, the function will return the device itself.

The response object will contain all the information from the Mist Inventory:
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
resource "mist_device_switch" "virtual_chassis_one" {
  device_id = provider::mist::search_vc_by_member_serial(resource.mist_org_inventory.inventory, "FJ0424000001").id
  site_id   = provider::mist::search_vc_by_member_serial(resource.mist_org_inventory.inventory, "FJ0424000001").site_id
  name      = "virtual_chassis_one"
}
```

## Signature

<!-- signature generated by tfplugindocs -->
```text
search_vc_by_member_serial(inventory object, claim_code string) object
```

## Arguments

<!-- arguments generated by tfplugindocs -->
1. `inventory` (Object) `mist_org_inventory` resource
1. `claim_code` (String) Device Claim Code

