---
page_title: "mist_device_gateway_stats Data Source - terraform-provider-mist"
subcategory: "Devices"
description: |-
  This data source provides the list of Gateways with their statistics.
---

# mist_device_gateway_stats (Data Source)

This data source provides the list of Gateways with their statistics.


## Example Usage

```terraform
data "mist_device_gateway_stats" "gateway_stats" {
  org_id  = "15fca2ac-b1a6-47cc-9953-cc6906281550"

  // Filtering options
  mac     = "e8a245000000"
  site_id = "4a422ae5-7ca0-4599-87a3-8e49aa63685f"
  status  = "connected"

  // Stats time range - option #1
  // cannot be used with the `start`/`end` attribute
  // when using the `duration` attribute, `end`==`now` 
  duration = "1d"

  // Stats time range - option #2
  // cannot be used with the `duration` attribute
  start = 1736031600
  end   = 1736175934
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `org_id` (String)

### Optional

- `duration` (String) duration like 7d, 2w
- `end` (Number) end datetime, can be epoch or relative time like -1d, -2h; now if not specified
- `mac` (String)
- `site_id` (String)
- `start` (Number) start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified
- `status` (String)

### Read-Only

- `device_gateway_stats` (Attributes Set) (see [below for nested schema](#nestedatt--device_gateway_stats))

<a id="nestedatt--device_gateway_stats"></a>
### Nested Schema for `device_gateway_stats`

Read-Only:

- `ap_redundancy` (Attributes) (see [below for nested schema](#nestedatt--device_gateway_stats--ap_redundancy))
- `arp_table_stats` (Attributes) (see [below for nested schema](#nestedatt--device_gateway_stats--arp_table_stats))
- `cert_expiry` (Number)
- `cluster_config` (Attributes) (see [below for nested schema](#nestedatt--device_gateway_stats--cluster_config))
- `cluster_stat` (Attributes) (see [below for nested schema](#nestedatt--device_gateway_stats--cluster_stat))
- `conductor_name` (String)
- `config_status` (String)
- `cpu2_stat` (Attributes) (see [below for nested schema](#nestedatt--device_gateway_stats--cpu2_stat))
- `cpu_stat` (Attributes) (see [below for nested schema](#nestedatt--device_gateway_stats--cpu_stat))
- `created_time` (Number)
- `deviceprofile_id` (String)
- `dhcpd2_stat` (Attributes Map) Property key is the network name (see [below for nested schema](#nestedatt--device_gateway_stats--dhcpd2_stat))
- `dhcpd_stat` (Attributes Map) Property key is the network name (see [below for nested schema](#nestedatt--device_gateway_stats--dhcpd_stat))
- `ext_ip` (String) IP address
- `fwupdate` (Attributes) (see [below for nested schema](#nestedatt--device_gateway_stats--fwupdate))
- `has_pcap` (Boolean)
- `hostname` (String) hostname reported by the device
- `id` (String) serial
- `if2_stat` (Attributes Map) Property key is the interface name (see [below for nested schema](#nestedatt--device_gateway_stats--if2_stat))
- `if_stat` (Attributes Map) Property key is the interface name (see [below for nested schema](#nestedatt--device_gateway_stats--if_stat))
- `ip` (String) IP address
- `ip2_stat` (Attributes) (see [below for nested schema](#nestedatt--device_gateway_stats--ip2_stat))
- `ip_stat` (Attributes) (see [below for nested schema](#nestedatt--device_gateway_stats--ip_stat))
- `is_ha` (Boolean)
- `last_seen` (Number) last seen timestamp
- `mac` (String) device mac
- `map_id` (String) serial
- `memory2_stat` (Attributes) memory usage stat (for virtual chassis, memory usage of master RE) (see [below for nested schema](#nestedatt--device_gateway_stats--memory2_stat))
- `memory_stat` (Attributes) memory usage stat (for virtual chassis, memory usage of master RE) (see [below for nested schema](#nestedatt--device_gateway_stats--memory_stat))
- `model` (String) device model
- `modified_time` (Number)
- `module2_stat` (Attributes List) (see [below for nested schema](#nestedatt--device_gateway_stats--module2_stat))
- `module_stat` (Attributes List) (see [below for nested schema](#nestedatt--device_gateway_stats--module_stat))
- `name` (String) device name if configured
- `node_name` (String)
- `org_id` (String) serial
- `route_summary_stats` (Attributes) (see [below for nested schema](#nestedatt--device_gateway_stats--route_summary_stats))
- `router_name` (String) device name if configured
- `serial` (String) serial
- `service2_stat` (Attributes Map) (see [below for nested schema](#nestedatt--device_gateway_stats--service2_stat))
- `service_stat` (Attributes Map) (see [below for nested schema](#nestedatt--device_gateway_stats--service_stat))
- `service_status` (Attributes) (see [below for nested schema](#nestedatt--device_gateway_stats--service_status))
- `site_id` (String) serial
- `spu2_stat` (Attributes List) (see [below for nested schema](#nestedatt--device_gateway_stats--spu2_stat))
- `spu_stat` (Attributes List) (see [below for nested schema](#nestedatt--device_gateway_stats--spu_stat))
- `status` (String)
- `uptime` (Number)
- `version` (String)

<a id="nestedatt--device_gateway_stats--ap_redundancy"></a>
### Nested Schema for `device_gateway_stats.ap_redundancy`

Read-Only:

- `modules` (Attributes Map) Property key is the node id (see [below for nested schema](#nestedatt--device_gateway_stats--ap_redundancy--modules))
- `num_aps` (Number)
- `num_aps_with_switch_redundancy` (Number)

<a id="nestedatt--device_gateway_stats--ap_redundancy--modules"></a>
### Nested Schema for `device_gateway_stats.ap_redundancy.modules`

Read-Only:

- `num_aps` (Number)
- `num_aps_with_switch_redundancy` (Number)



<a id="nestedatt--device_gateway_stats--arp_table_stats"></a>
### Nested Schema for `device_gateway_stats.arp_table_stats`

Read-Only:

- `arp_table_count` (Number)
- `max_entries_supported` (Number)


<a id="nestedatt--device_gateway_stats--cluster_config"></a>
### Nested Schema for `device_gateway_stats.cluster_config`

Read-Only:

- `configuration` (String)
- `control_link_info` (Attributes) (see [below for nested schema](#nestedatt--device_gateway_stats--cluster_config--control_link_info))
- `ethernet_connection` (Attributes List) (see [below for nested schema](#nestedatt--device_gateway_stats--cluster_config--ethernet_connection))
- `fabric_link_info` (Attributes) (see [below for nested schema](#nestedatt--device_gateway_stats--cluster_config--fabric_link_info))
- `last_status_change_reason` (String)
- `operational` (String)
- `primary_node_health` (String)
- `redundancy_group_information` (Attributes List) (see [below for nested schema](#nestedatt--device_gateway_stats--cluster_config--redundancy_group_information))
- `secondary_node_health` (String)
- `status` (String)

<a id="nestedatt--device_gateway_stats--cluster_config--control_link_info"></a>
### Nested Schema for `device_gateway_stats.cluster_config.control_link_info`

Read-Only:

- `name` (String)
- `status` (String)


<a id="nestedatt--device_gateway_stats--cluster_config--ethernet_connection"></a>
### Nested Schema for `device_gateway_stats.cluster_config.ethernet_connection`

Read-Only:

- `name` (String)
- `status` (String)


<a id="nestedatt--device_gateway_stats--cluster_config--fabric_link_info"></a>
### Nested Schema for `device_gateway_stats.cluster_config.fabric_link_info`

Read-Only:

- `data_plane_notified_status` (String)
- `interface` (List of String)
- `internal_status` (String)
- `state` (String)
- `status` (String)


<a id="nestedatt--device_gateway_stats--cluster_config--redundancy_group_information"></a>
### Nested Schema for `device_gateway_stats.cluster_config.redundancy_group_information`

Read-Only:

- `id` (Number)
- `monitoring_failure` (String)
- `threshold` (Number)



<a id="nestedatt--device_gateway_stats--cluster_stat"></a>
### Nested Schema for `device_gateway_stats.cluster_stat`

Read-Only:

- `state` (String)


<a id="nestedatt--device_gateway_stats--cpu2_stat"></a>
### Nested Schema for `device_gateway_stats.cpu2_stat`

Read-Only:

- `idle` (Number) Percentage of CPU time that is idle
- `interrupt` (Number) Percentage of CPU time being used by interrupts
- `load_avg` (List of Number) Load averages for the last 1, 5, and 15 minutes
- `system` (Number) Percentage of CPU time being used by system processes
- `user` (Number) Percentage of CPU time being used by user processe


<a id="nestedatt--device_gateway_stats--cpu_stat"></a>
### Nested Schema for `device_gateway_stats.cpu_stat`

Read-Only:

- `idle` (Number) Percentage of CPU time that is idle
- `interrupt` (Number) Percentage of CPU time being used by interrupts
- `load_avg` (List of Number) Load averages for the last 1, 5, and 15 minutes
- `system` (Number) Percentage of CPU time being used by system processes
- `user` (Number) Percentage of CPU time being used by user processe


<a id="nestedatt--device_gateway_stats--dhcpd2_stat"></a>
### Nested Schema for `device_gateway_stats.dhcpd2_stat`

Read-Only:

- `num_ips` (Number)
- `num_leased` (Number)


<a id="nestedatt--device_gateway_stats--dhcpd_stat"></a>
### Nested Schema for `device_gateway_stats.dhcpd_stat`

Read-Only:

- `num_ips` (Number)
- `num_leased` (Number)


<a id="nestedatt--device_gateway_stats--fwupdate"></a>
### Nested Schema for `device_gateway_stats.fwupdate`

Read-Only:

- `progress` (Number)
- `status` (String)
- `status_id` (Number)
- `timestamp` (Number)
- `will_retry` (Boolean)


<a id="nestedatt--device_gateway_stats--if2_stat"></a>
### Nested Schema for `device_gateway_stats.if2_stat`

Read-Only:

- `address_mode` (String)
- `ips` (List of String)
- `nat_addresses` (List of String)
- `network_name` (String)
- `port_id` (String)
- `port_usage` (String)
- `redundancy_state` (String)
- `rx_bytes` (Number)
- `rx_pkts` (Number)
- `servp_info` (Attributes) (see [below for nested schema](#nestedatt--device_gateway_stats--if2_stat--servp_info))
- `tx_bytes` (Number)
- `tx_pkts` (Number)
- `up` (Boolean)
- `vlan` (Number)
- `wan_name` (String)
- `wan_type` (String)

<a id="nestedatt--device_gateway_stats--if2_stat--servp_info"></a>
### Nested Schema for `device_gateway_stats.if2_stat.servp_info`

Read-Only:

- `asn` (String)
- `city` (String)
- `country_code` (String)
- `latitude` (Number)
- `longitude` (Number)
- `org` (String)
- `region_code` (String)



<a id="nestedatt--device_gateway_stats--if_stat"></a>
### Nested Schema for `device_gateway_stats.if_stat`

Read-Only:

- `address_mode` (String)
- `ips` (List of String)
- `nat_addresses` (List of String)
- `network_name` (String)
- `port_id` (String)
- `port_usage` (String)
- `redundancy_state` (String)
- `rx_bytes` (Number)
- `rx_pkts` (Number)
- `servp_info` (Attributes) (see [below for nested schema](#nestedatt--device_gateway_stats--if_stat--servp_info))
- `tx_bytes` (Number)
- `tx_pkts` (Number)
- `up` (Boolean)
- `vlan` (Number)
- `wan_name` (String)
- `wan_type` (String)

<a id="nestedatt--device_gateway_stats--if_stat--servp_info"></a>
### Nested Schema for `device_gateway_stats.if_stat.servp_info`

Read-Only:

- `asn` (String)
- `city` (String)
- `country_code` (String)
- `latitude` (Number)
- `longitude` (Number)
- `org` (String)
- `region_code` (String)



<a id="nestedatt--device_gateway_stats--ip2_stat"></a>
### Nested Schema for `device_gateway_stats.ip2_stat`

Read-Only:

- `dhcp_server` (String)
- `dns` (List of String)
- `dns_suffix` (List of String)
- `gateway` (String)
- `gateway6` (String)
- `ip` (String)
- `ip6` (String)
- `ips` (Map of String)
- `netmask` (String)
- `netmask6` (String)


<a id="nestedatt--device_gateway_stats--ip_stat"></a>
### Nested Schema for `device_gateway_stats.ip_stat`

Read-Only:

- `dhcp_server` (String)
- `dns` (List of String)
- `dns_suffix` (List of String)
- `gateway` (String)
- `gateway6` (String)
- `ip` (String)
- `ip6` (String)
- `ips` (Map of String)
- `netmask` (String)
- `netmask6` (String)


<a id="nestedatt--device_gateway_stats--memory2_stat"></a>
### Nested Schema for `device_gateway_stats.memory2_stat`

Read-Only:

- `usage` (Number)


<a id="nestedatt--device_gateway_stats--memory_stat"></a>
### Nested Schema for `device_gateway_stats.memory_stat`

Read-Only:

- `usage` (Number)


<a id="nestedatt--device_gateway_stats--module2_stat"></a>
### Nested Schema for `device_gateway_stats.module2_stat`

Read-Only:

- `backup_version` (String)
- `bios_version` (String)
- `cpld_version` (String)
- `errors` (Attributes List) used to report all error states the device node is running into.
An error should always have `type` and `since` fields, and could have some other fields specific to that type. (see [below for nested schema](#nestedatt--device_gateway_stats--module2_stat--errors))
- `fans` (Attributes List) (see [below for nested schema](#nestedatt--device_gateway_stats--module2_stat--fans))
- `fpga_version` (String)
- `last_seen` (Number)
- `model` (String)
- `optics_cpld_version` (String)
- `pending_version` (String)
- `pics` (Attributes List) (see [below for nested schema](#nestedatt--device_gateway_stats--module2_stat--pics))
- `poe` (Attributes) (see [below for nested schema](#nestedatt--device_gateway_stats--module2_stat--poe))
- `poe_version` (String)
- `power_cpld_version` (String)
- `psus` (Attributes List) (see [below for nested schema](#nestedatt--device_gateway_stats--module2_stat--psus))
- `re_fpga_version` (String)
- `recovery_version` (String)
- `serial` (String)
- `status` (String)
- `temperatures` (Attributes List) (see [below for nested schema](#nestedatt--device_gateway_stats--module2_stat--temperatures))
- `tmc_fpga_version` (String)
- `uboot_version` (String)
- `uptime` (Number)
- `vc_links` (Attributes List) (see [below for nested schema](#nestedatt--device_gateway_stats--module2_stat--vc_links))
- `vc_mode` (String)
- `vc_role` (String) master / backup / linecard
- `vc_state` (String)
- `version` (String)

<a id="nestedatt--device_gateway_stats--module2_stat--errors"></a>
### Nested Schema for `device_gateway_stats.module2_stat.errors`

Read-Only:

- `feature` (String)
- `minimum_version` (String)
- `reason` (String)
- `since` (Number)
- `type` (String)


<a id="nestedatt--device_gateway_stats--module2_stat--fans"></a>
### Nested Schema for `device_gateway_stats.module2_stat.fans`

Read-Only:

- `airflow` (String)
- `name` (String)
- `status` (String)


<a id="nestedatt--device_gateway_stats--module2_stat--pics"></a>
### Nested Schema for `device_gateway_stats.module2_stat.pics`

Read-Only:

- `index` (Number)
- `model_number` (String)
- `port_groups` (Attributes List) (see [below for nested schema](#nestedatt--device_gateway_stats--module2_stat--pics--port_groups))

<a id="nestedatt--device_gateway_stats--module2_stat--pics--port_groups"></a>
### Nested Schema for `device_gateway_stats.module2_stat.pics.port_groups`

Read-Only:

- `count` (Number)
- `type` (String)



<a id="nestedatt--device_gateway_stats--module2_stat--poe"></a>
### Nested Schema for `device_gateway_stats.module2_stat.poe`

Read-Only:

- `max_power` (Number)
- `power_draw` (Number)


<a id="nestedatt--device_gateway_stats--module2_stat--psus"></a>
### Nested Schema for `device_gateway_stats.module2_stat.psus`

Read-Only:

- `name` (String)
- `status` (String)


<a id="nestedatt--device_gateway_stats--module2_stat--temperatures"></a>
### Nested Schema for `device_gateway_stats.module2_stat.temperatures`

Read-Only:

- `celsius` (Number)
- `name` (String)
- `status` (String)


<a id="nestedatt--device_gateway_stats--module2_stat--vc_links"></a>
### Nested Schema for `device_gateway_stats.module2_stat.vc_links`

Read-Only:

- `neighbor_module_idx` (Number)
- `neighbor_port_id` (String)
- `port_id` (String)



<a id="nestedatt--device_gateway_stats--module_stat"></a>
### Nested Schema for `device_gateway_stats.module_stat`

Read-Only:

- `backup_version` (String)
- `bios_version` (String)
- `cpld_version` (String)
- `errors` (Attributes List) used to report all error states the device node is running into.
An error should always have `type` and `since` fields, and could have some other fields specific to that type. (see [below for nested schema](#nestedatt--device_gateway_stats--module_stat--errors))
- `fans` (Attributes List) (see [below for nested schema](#nestedatt--device_gateway_stats--module_stat--fans))
- `fpga_version` (String)
- `last_seen` (Number)
- `model` (String)
- `optics_cpld_version` (String)
- `pending_version` (String)
- `pics` (Attributes List) (see [below for nested schema](#nestedatt--device_gateway_stats--module_stat--pics))
- `poe` (Attributes) (see [below for nested schema](#nestedatt--device_gateway_stats--module_stat--poe))
- `poe_version` (String)
- `power_cpld_version` (String)
- `psus` (Attributes List) (see [below for nested schema](#nestedatt--device_gateway_stats--module_stat--psus))
- `re_fpga_version` (String)
- `recovery_version` (String)
- `serial` (String)
- `status` (String)
- `temperatures` (Attributes List) (see [below for nested schema](#nestedatt--device_gateway_stats--module_stat--temperatures))
- `tmc_fpga_version` (String)
- `uboot_version` (String)
- `uptime` (Number)
- `vc_links` (Attributes List) (see [below for nested schema](#nestedatt--device_gateway_stats--module_stat--vc_links))
- `vc_mode` (String)
- `vc_role` (String) master / backup / linecard
- `vc_state` (String)
- `version` (String)

<a id="nestedatt--device_gateway_stats--module_stat--errors"></a>
### Nested Schema for `device_gateway_stats.module_stat.errors`

Read-Only:

- `feature` (String)
- `minimum_version` (String)
- `reason` (String)
- `since` (Number)
- `type` (String)


<a id="nestedatt--device_gateway_stats--module_stat--fans"></a>
### Nested Schema for `device_gateway_stats.module_stat.fans`

Read-Only:

- `airflow` (String)
- `name` (String)
- `status` (String)


<a id="nestedatt--device_gateway_stats--module_stat--pics"></a>
### Nested Schema for `device_gateway_stats.module_stat.pics`

Read-Only:

- `index` (Number)
- `model_number` (String)
- `port_groups` (Attributes List) (see [below for nested schema](#nestedatt--device_gateway_stats--module_stat--pics--port_groups))

<a id="nestedatt--device_gateway_stats--module_stat--pics--port_groups"></a>
### Nested Schema for `device_gateway_stats.module_stat.pics.port_groups`

Read-Only:

- `count` (Number)
- `type` (String)



<a id="nestedatt--device_gateway_stats--module_stat--poe"></a>
### Nested Schema for `device_gateway_stats.module_stat.poe`

Read-Only:

- `max_power` (Number)
- `power_draw` (Number)


<a id="nestedatt--device_gateway_stats--module_stat--psus"></a>
### Nested Schema for `device_gateway_stats.module_stat.psus`

Read-Only:

- `name` (String)
- `status` (String)


<a id="nestedatt--device_gateway_stats--module_stat--temperatures"></a>
### Nested Schema for `device_gateway_stats.module_stat.temperatures`

Read-Only:

- `celsius` (Number)
- `name` (String)
- `status` (String)


<a id="nestedatt--device_gateway_stats--module_stat--vc_links"></a>
### Nested Schema for `device_gateway_stats.module_stat.vc_links`

Read-Only:

- `neighbor_module_idx` (Number)
- `neighbor_port_id` (String)
- `port_id` (String)



<a id="nestedatt--device_gateway_stats--route_summary_stats"></a>
### Nested Schema for `device_gateway_stats.route_summary_stats`

Read-Only:

- `fib_routes` (Number)
- `max_unicast_routes_supported` (Number)
- `rib_routes` (Number)
- `total_routes` (Number)


<a id="nestedatt--device_gateway_stats--service2_stat"></a>
### Nested Schema for `device_gateway_stats.service2_stat`

Read-Only:

- `ash_version` (String)
- `cia_version` (String)
- `ember_version` (String)
- `ipsec_client_version` (String)
- `mist_agent_version` (String)
- `package_version` (String)
- `testing_tools_version` (String)
- `wheeljack_version` (String)


<a id="nestedatt--device_gateway_stats--service_stat"></a>
### Nested Schema for `device_gateway_stats.service_stat`

Read-Only:

- `ash_version` (String)
- `cia_version` (String)
- `ember_version` (String)
- `ipsec_client_version` (String)
- `mist_agent_version` (String)
- `package_version` (String)
- `testing_tools_version` (String)
- `wheeljack_version` (String)


<a id="nestedatt--device_gateway_stats--service_status"></a>
### Nested Schema for `device_gateway_stats.service_status`

Read-Only:

- `appid_install_result` (String)
- `appid_install_timestamp` (String)
- `appid_status` (String)
- `appid_version` (Number)
- `ewf_status` (String)
- `idp_install_result` (String)
- `idp_install_timestamp` (String)
- `idp_policy` (String)
- `idp_status` (String)
- `idp_update_timestamp` (String)


<a id="nestedatt--device_gateway_stats--spu2_stat"></a>
### Nested Schema for `device_gateway_stats.spu2_stat`

Read-Only:

- `spu_cpu` (Number)
- `spu_current_session` (Number)
- `spu_max_session` (Number)
- `spu_memory` (Number)
- `spu_pending_session` (Number)
- `spu_valid_session` (Number)


<a id="nestedatt--device_gateway_stats--spu_stat"></a>
### Nested Schema for `device_gateway_stats.spu_stat`

Read-Only:

- `spu_cpu` (Number)
- `spu_current_session` (Number)
- `spu_max_session` (Number)
- `spu_memory` (Number)
- `spu_pending_session` (Number)
- `spu_valid_session` (Number)