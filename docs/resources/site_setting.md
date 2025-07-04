---
page_title: "mist_site_setting Resource - terraform-provider-mist"
subcategory: "Site"
description: |-
  This resource manages the Site Settings.
  The Site Settings can be used to customize the Site configuration and assign Site Variables (Sites Variables can be reused in configuration templates)
  ~> When using the Mist APIs, all the switch settings defined at the site level are stored under the site settings with all the rest of the site configuration (/api/v1/sites/{site_id}/setting Mist API Endpoint). To simplify this resource, all the site level switches related settings are moved into the mist_site_networktemplate resource
  !> Only ONE mist_site_setting resource can be configured per site. If multiple ones are configured, only the last one defined we be successfully deployed to Mist
---

# mist_site_setting (Resource)

This resource manages the Site Settings.

The Site Settings can be used to customize the Site configuration and assign Site Variables (Sites Variables can be reused in configuration templates)

~> When using the Mist APIs, all the switch settings defined at the site level are stored under the site settings with all the rest of the site configuration (`/api/v1/sites/{site_id}/setting` Mist API Endpoint). To simplify this resource, all the site level switches related settings are moved into the `mist_site_networktemplate` resource

!> Only ONE `mist_site_setting` resource can be configured per site. If multiple ones are configured, only the last one defined we be successfully deployed to Mist


## Example Usage

```terraform
resource "mist_site_setting" "site_one" {
  site_id                 = mist_site.terraform_site.id
  ap_updown_threshold     = 5
  device_updown_threshold = 5
  auto_upgrade = {
    enabled     = true
    day_of_week = "tue"
    time_of_day = "02:00"
    version     = "beta"
  }
  config_auto_revert       = true
  persist_config_on_device = true
  proxy = {
    url = "http://myproxy:3128"
  }
  rogue = {
    enabled          = true
    honeypot_enabled = true
    min_duration     = 5
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `site_id` (String)

### Optional

- `analytic` (Attributes) (see [below for nested schema](#nestedatt--analytic))
- `ap_updown_threshold` (Number) Enable threshold-based device down delivery for AP devices only. When configured it takes effect for AP devices and `device_updown_threshold` is ignored.
- `auto_upgrade` (Attributes) Auto Upgrade Settings (see [below for nested schema](#nestedatt--auto_upgrade))
- `ble_config` (Attributes) BLE AP settings (see [below for nested schema](#nestedatt--ble_config))
- `config_auto_revert` (Boolean) Whether to enable ap auto config revert
- `config_push_policy` (Attributes) Mist also uses some heuristic rules to prevent destructive configs from being pushed (see [below for nested schema](#nestedatt--config_push_policy))
- `critical_url_monitoring` (Attributes) You can define some URLs that's critical to site operations the latency will be captured and considered for site health (see [below for nested schema](#nestedatt--critical_url_monitoring))
- `default_port_usage` (String) Port usage to assign to switch ports without any port usage assigned. Default: `default` to preserve default behavior
- `device_updown_threshold` (Number) By default, device_updown_threshold, if set, will apply to all devices types if different values for specific device type is desired, use the following
- `enable_unii_4` (Boolean)
- `engagement` (Attributes) **Note**: if hours does not exist, it's treated as everyday of the week, 00:00-23:59. Currently, we don't allow multiple ranges for the same day (see [below for nested schema](#nestedatt--engagement))
- `gateway_mgmt` (Attributes) Gateway Site settings (see [below for nested schema](#nestedatt--gateway_mgmt))
- `gateway_updown_threshold` (Number) Enable threshold-based device down delivery for Gateway devices only. When configured it takes effect for GW devices and `device_updown_threshold` is ignored.
- `juniper_srx` (Attributes) (see [below for nested schema](#nestedatt--juniper_srx))
- `led` (Attributes) LED AP settings (see [below for nested schema](#nestedatt--led))
- `marvis` (Attributes) (see [below for nested schema](#nestedatt--marvis))
- `occupancy` (Attributes) Occupancy Analytics settings (see [below for nested schema](#nestedatt--occupancy))
- `persist_config_on_device` (Boolean) Whether to store the config on AP
- `proxy` (Attributes) Proxy Configuration to talk to Mist (see [below for nested schema](#nestedatt--proxy))
- `remove_existing_configs` (Boolean) By default, only the configuration generated by Mist is cleaned up during the configuration process. If `true`, all the existing configuration will be removed.
- `report_gatt` (Boolean) Whether AP should periodically connect to BLE devices and report GATT device info (device name, manufacturer name, serial number, battery %, temperature, humidity)
- `rogue` (Attributes) Rogue site settings (see [below for nested schema](#nestedatt--rogue))
- `rtsa` (Attributes) Managed mobility (see [below for nested schema](#nestedatt--rtsa))
- `simple_alert` (Attributes) Set of heuristic rules will be enabled when marvis subscription is not available. It triggers when, in a Z minute window, there are more than Y distinct client encountering over X failures (see [below for nested schema](#nestedatt--simple_alert))
- `skyatp` (Attributes) (see [below for nested schema](#nestedatt--skyatp))
- `sle_thresholds` (Attributes) (see [below for nested schema](#nestedatt--sle_thresholds))
- `srx_app` (Attributes) (see [below for nested schema](#nestedatt--srx_app))
- `ssh_keys` (List of String) When limit_ssh_access = true in Org Setting, list of SSH public keys provided by Mist Support to install onto APs (see Org:Setting)
- `ssr` (Attributes) (see [below for nested schema](#nestedatt--ssr))
- `switch_updown_threshold` (Number) Enable threshold-based device down delivery for Switch devices only. When configured it takes effect for SW devices and `device_updown_threshold` is ignored.
- `synthetic_test` (Attributes) (see [below for nested schema](#nestedatt--synthetic_test))
- `track_anonymous_devices` (Boolean) Whether to track anonymous BLE assets (requires ‘track_asset’  enabled)
- `uplink_port_config` (Attributes) AP Uplink port configuration (see [below for nested schema](#nestedatt--uplink_port_config))
- `vars` (Map of String) Dictionary of name->value, the vars can then be used in Wlans. This can overwrite those from Site Vars
- `vna` (Attributes) (see [below for nested schema](#nestedatt--vna))
- `vs_instance` (Attributes Map) Optional, for EX9200 only to segregate virtual-switches. Property key is the instance name (see [below for nested schema](#nestedatt--vs_instance))
- `wan_vna` (Attributes) (see [below for nested schema](#nestedatt--wan_vna))
- `wids` (Attributes) WIDS site settings (see [below for nested schema](#nestedatt--wids))
- `wifi` (Attributes) Wi-Fi site settings (see [below for nested schema](#nestedatt--wifi))
- `wired_vna` (Attributes) (see [below for nested schema](#nestedatt--wired_vna))
- `zone_occupancy_alert` (Attributes) Zone Occupancy alert site settings (see [below for nested schema](#nestedatt--zone_occupancy_alert))

### Read-Only

- `blacklist_url` (String)
- `watched_station_url` (String)
- `whitelist_url` (String)

<a id="nestedatt--analytic"></a>
### Nested Schema for `analytic`

Optional:

- `enabled` (Boolean) Enable Advanced Analytic feature (using SUB-ANA license)


<a id="nestedatt--auto_upgrade"></a>
### Nested Schema for `auto_upgrade`

Optional:

- `custom_versions` (Map of String) Custom versions for different models. Property key is the model name (e.g. "AP41")
- `day_of_week` (String) enum: `any`, `fri`, `mon`, `sat`, `sun`, `thu`, `tue`, `wed`
- `enabled` (Boolean) Whether auto upgrade should happen (Note that Mist may auto-upgrade if the version is not supported)
- `time_of_day` (String) `any` / HH:MM (24-hour format), upgrade will happen within up to 1-hour from this time
- `version` (String) desired version. enum: `beta`, `custom`, `stable`


<a id="nestedatt--ble_config"></a>
### Nested Schema for `ble_config`

Optional:

- `beacon_enabled` (Boolean) Whether Mist beacons is enabled
- `beacon_rate` (Number) Required if `beacon_rate_mode`==`custom`, 1-10, in number-beacons-per-second
- `beacon_rate_mode` (String) enum: `custom`, `default`
- `beam_disabled` (List of Number) List of AP BLE location beam numbers (1-8) which should be disabled at the AP and not transmit location information (where beam 1 is oriented at the top the AP, growing counter-clock-wise, with 9 being the omni BLE beam)
- `custom_ble_packet_enabled` (Boolean) Can be enabled if `beacon_enabled`==`true`, whether to send custom packet
- `custom_ble_packet_frame` (String) The custom frame to be sent out in this beacon. The frame must be a hexstring
- `custom_ble_packet_freq_msec` (Number) Frequency (msec) of data emitted by custom ble beacon
- `eddystone_uid_adv_power` (Number) Advertised TX Power, -100 to 20 (dBm), omit this attribute to use default
- `eddystone_uid_beams` (String)
- `eddystone_uid_enabled` (Boolean) Only if `beacon_enabled`==`false`, Whether Eddystone-UID beacon is enabled
- `eddystone_uid_freq_msec` (Number) Frequency (msec) of data emit by Eddystone-UID beacon
- `eddystone_uid_instance` (String) Eddystone-UID instance for the device
- `eddystone_uid_namespace` (String) Eddystone-UID namespace
- `eddystone_url_adv_power` (Number) Advertised TX Power, -100 to 20 (dBm), omit this attribute to use default
- `eddystone_url_beams` (String)
- `eddystone_url_enabled` (Boolean) Only if `beacon_enabled`==`false`, Whether Eddystone-URL beacon is enabled
- `eddystone_url_freq_msec` (Number) Frequency (msec) of data emit by Eddystone-UID beacon
- `eddystone_url_url` (String) URL pointed by Eddystone-URL beacon
- `ibeacon_adv_power` (Number) Advertised TX Power, -100 to 20 (dBm), omit this attribute to use default
- `ibeacon_beams` (String)
- `ibeacon_enabled` (Boolean) Can be enabled if `beacon_enabled`==`true`, whether to send iBeacon
- `ibeacon_freq_msec` (Number) Frequency (msec) of data emit for iBeacon
- `ibeacon_major` (Number) Major number for iBeacon
- `ibeacon_minor` (Number) Minor number for iBeacon
- `ibeacon_uuid` (String) Optional, if not specified, the same UUID as the beacon will be used
- `power` (Number) Required if `power_mode`==`custom`; else use `power_mode` as default
- `power_mode` (String) enum: `custom`, `default`


<a id="nestedatt--config_push_policy"></a>
### Nested Schema for `config_push_policy`

Optional:

- `no_push` (Boolean) Stop any new config from being pushed to the device
- `push_window` (Attributes) If enabled, new config will only be pushed to device within the specified time window (see [below for nested schema](#nestedatt--config_push_policy--push_window))

<a id="nestedatt--config_push_policy--push_window"></a>
### Nested Schema for `config_push_policy.push_window`

Optional:

- `enabled` (Boolean)
- `hours` (Attributes) Days/Hours of operation filter, the available days (mon, tue, wed, thu, fri, sat, sun) (see [below for nested schema](#nestedatt--config_push_policy--push_window--hours))

<a id="nestedatt--config_push_policy--push_window--hours"></a>
### Nested Schema for `config_push_policy.push_window.hours`

Optional:

- `fri` (String) Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59.
- `mon` (String) Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59.
- `sat` (String) Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59.
- `sun` (String) Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59.
- `thu` (String) Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59.
- `tue` (String) Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59.
- `wed` (String) Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59.




<a id="nestedatt--critical_url_monitoring"></a>
### Nested Schema for `critical_url_monitoring`

Optional:

- `enabled` (Boolean)
- `monitors` (Attributes List) (see [below for nested schema](#nestedatt--critical_url_monitoring--monitors))

<a id="nestedatt--critical_url_monitoring--monitors"></a>
### Nested Schema for `critical_url_monitoring.monitors`

Optional:

- `url` (String)
- `vlan_id` (String)



<a id="nestedatt--engagement"></a>
### Nested Schema for `engagement`

Optional:

- `dwell_tag_names` (Attributes) Name associated to each tag (see [below for nested schema](#nestedatt--engagement--dwell_tag_names))
- `dwell_tags` (Attributes) add tags to visits within the duration (in seconds) (see [below for nested schema](#nestedatt--engagement--dwell_tags))
- `hours` (Attributes) Days/Hours of operation filter, the available days (mon, tue, wed, thu, fri, sat, sun) (see [below for nested schema](#nestedatt--engagement--hours))
- `max_dwell` (Number) Max time, default is 43200(12h), max is 68400 (18h)
- `min_dwell` (Number) min time

<a id="nestedatt--engagement--dwell_tag_names"></a>
### Nested Schema for `engagement.dwell_tag_names`

Optional:

- `bounce` (String) Default to `Visitor`
- `engaged` (String) Default to `Associates`
- `passerby` (String) Default to `Passerby`
- `stationed` (String) Default to `Assets`


<a id="nestedatt--engagement--dwell_tags"></a>
### Nested Schema for `engagement.dwell_tags`

Optional:

- `bounce` (String) Default to `301-14400`
- `engaged` (String) Default to `14401-28800`
- `passerby` (String) Default to `1-300`
- `stationed` (String) Default to `28801-42000`


<a id="nestedatt--engagement--hours"></a>
### Nested Schema for `engagement.hours`

Optional:

- `fri` (String) Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59.
- `mon` (String) Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59.
- `sat` (String) Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59.
- `sun` (String) Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59.
- `thu` (String) Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59.
- `tue` (String) Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59.
- `wed` (String) Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59.



<a id="nestedatt--gateway_mgmt"></a>
### Nested Schema for `gateway_mgmt`

Optional:

- `admin_sshkeys` (List of String) For SSR only, as direct root access is not allowed
- `app_probing` (Attributes) (see [below for nested schema](#nestedatt--gateway_mgmt--app_probing))
- `app_usage` (Boolean) Consumes uplink bandwidth, requires WA license
- `auto_signature_update` (Attributes) (see [below for nested schema](#nestedatt--gateway_mgmt--auto_signature_update))
- `config_revert_timer` (Number) Rollback timer for commit confirmed
- `disable_console` (Boolean) For SSR and SRX, disable console port
- `disable_oob` (Boolean) For SSR and SRX, disable management interface
- `disable_usb` (Boolean) For SSR and SRX, disable usb interface
- `fips_enabled` (Boolean)
- `probe_hosts` (List of String)
- `protect_re` (Attributes) Restrict inbound-traffic to host
when enabled, all traffic that is not essential to our operation will be dropped 
e.g. ntp / dns / traffic to mist will be allowed by default, if dhcpd is enabled, we'll make sure it works (see [below for nested schema](#nestedatt--gateway_mgmt--protect_re))
- `root_password` (String, Sensitive) For SRX only
- `security_log_source_address` (String)
- `security_log_source_interface` (String)

<a id="nestedatt--gateway_mgmt--app_probing"></a>
### Nested Schema for `gateway_mgmt.app_probing`

Optional:

- `apps` (List of String) APp-keys from [List Applications]($e/Constants%20Definitions/listApplications)
- `custom_apps` (Attributes List) (see [below for nested schema](#nestedatt--gateway_mgmt--app_probing--custom_apps))
- `enabled` (Boolean)

<a id="nestedatt--gateway_mgmt--app_probing--custom_apps"></a>
### Nested Schema for `gateway_mgmt.app_probing.custom_apps`

Required:

- `hostnames` (List of String) Only 1 entry is allowed:
    * if `protocol`==`http`: URL (e.g. `http://test.com` or `https://test.com`)
    * if `protocol`==`icmp`: IP Address (e.g. `1.2.3.4`)
- `name` (String)
- `protocol` (String) enum: `http`, `icmp`

Optional:

- `app_type` (String)
- `network` (String)
- `packet_size` (Number) If `protocol`==`icmp`
- `vrf` (String)

Read-Only:

- `address` (String)
- `key` (String)
- `url` (String)



<a id="nestedatt--gateway_mgmt--auto_signature_update"></a>
### Nested Schema for `gateway_mgmt.auto_signature_update`

Optional:

- `day_of_week` (String) enum: `any`, `fri`, `mon`, `sat`, `sun`, `thu`, `tue`, `wed`
- `enable` (Boolean)
- `time_of_day` (String) Optional, Mist will decide the timing


<a id="nestedatt--gateway_mgmt--protect_re"></a>
### Nested Schema for `gateway_mgmt.protect_re`

Optional:

- `allowed_services` (List of String) optionally, services we'll allow. enum: `icmp`, `ssh`
- `custom` (Attributes List) (see [below for nested schema](#nestedatt--gateway_mgmt--protect_re--custom))
- `enabled` (Boolean) When enabled, all traffic that is not essential to our operation will be dropped
e.g. ntp / dns / traffic to mist will be allowed by default
     if dhcpd is enabled, we'll make sure it works
- `hit_count` (Boolean) Whether to enable hit count for Protect_RE policy
- `trusted_hosts` (List of String) host/subnets we'll allow traffic to/from

<a id="nestedatt--gateway_mgmt--protect_re--custom"></a>
### Nested Schema for `gateway_mgmt.protect_re.custom`

Required:

- `subnets` (List of String)

Optional:

- `port_range` (String) matched dst port, "0" means any. Note: For `protocol`==`any` and  `port_range`==`any`, configure `trusted_hosts` instead
- `protocol` (String) enum: `any`, `icmp`, `tcp`, `udp`. Note: For `protocol`==`any` and  `port_range`==`any`, configure `trusted_hosts` instead




<a id="nestedatt--juniper_srx"></a>
### Nested Schema for `juniper_srx`

Optional:

- `gateways` (Attributes List) (see [below for nested schema](#nestedatt--juniper_srx--gateways))
- `send_mist_nac_user_info` (Boolean)

<a id="nestedatt--juniper_srx--gateways"></a>
### Nested Schema for `juniper_srx.gateways`

Optional:

- `api_key` (String)
- `api_password` (String)
- `api_url` (String)



<a id="nestedatt--led"></a>
### Nested Schema for `led`

Optional:

- `brightness` (Number)
- `enabled` (Boolean)


<a id="nestedatt--marvis"></a>
### Nested Schema for `marvis`

Optional:

- `auto_operations` (Attributes) (see [below for nested schema](#nestedatt--marvis--auto_operations))

<a id="nestedatt--marvis--auto_operations"></a>
### Nested Schema for `marvis.auto_operations`

Optional:

- `bounce_port_for_abnormal_poe_client` (Boolean)
- `disable_port_when_ddos_protocol_violation` (Boolean)
- `disable_port_when_rogue_dhcp_server_detected` (Boolean)



<a id="nestedatt--occupancy"></a>
### Nested Schema for `occupancy`

Optional:

- `assets_enabled` (Boolean) Indicate whether named BLE assets should be included in the zone occupancy calculation
- `clients_enabled` (Boolean) Indicate whether connected Wi-Fi clients should be included in the zone occupancy calculation
- `min_duration` (Number) Minimum duration
- `sdkclients_enabled` (Boolean) Indicate whether SDK clients should be included in the zone occupancy calculation
- `unconnected_clients_enabled` (Boolean) Indicate whether unconnected Wi-Fi clients should be included in the zone occupancy calculation


<a id="nestedatt--proxy"></a>
### Nested Schema for `proxy`

Optional:

- `url` (String)


<a id="nestedatt--rogue"></a>
### Nested Schema for `rogue`

Optional:

- `allowed_vlan_ids` (List of Number) list of VLAN IDs on which rogue APs are ignored
- `enabled` (Boolean) Whether rogue detection is enabled
- `honeypot_enabled` (Boolean) Whether honeypot detection is enabled
- `min_duration` (Number) Minimum duration for a bssid to be considered neighbor
- `min_rogue_duration` (Number) Minimum duration for a bssid to be considered rogue
- `min_rogue_rssi` (Number) Minimum RSSI for an AP to be considered rogue
- `min_rssi` (Number) Minimum RSSI for an AP to be considered neighbor (ignoring APs that’s far away)
- `whitelisted_bssids` (List of String) list of BSSIDs to whitelist. Ex: "cc-:8e-:6f-:d4-:bf-:16", "cc-8e-6f-d4-bf-16", "cc-73-*", "cc:82:*"
- `whitelisted_ssids` (List of String) List of SSIDs to whitelist


<a id="nestedatt--rtsa"></a>
### Nested Schema for `rtsa`

Optional:

- `app_waking` (Boolean)
- `disable_dead_reckoning` (Boolean)
- `disable_pressure_sensor` (Boolean)
- `enabled` (Boolean)
- `track_asset` (Boolean) Asset tracking related


<a id="nestedatt--simple_alert"></a>
### Nested Schema for `simple_alert`

Optional:

- `arp_failure` (Attributes) (see [below for nested schema](#nestedatt--simple_alert--arp_failure))
- `dhcp_failure` (Attributes) (see [below for nested schema](#nestedatt--simple_alert--dhcp_failure))
- `dns_failure` (Attributes) (see [below for nested schema](#nestedatt--simple_alert--dns_failure))

<a id="nestedatt--simple_alert--arp_failure"></a>
### Nested Schema for `simple_alert.arp_failure`

Optional:

- `client_count` (Number)
- `duration` (Number) failing within minutes
- `incident_count` (Number)


<a id="nestedatt--simple_alert--dhcp_failure"></a>
### Nested Schema for `simple_alert.dhcp_failure`

Optional:

- `client_count` (Number)
- `duration` (Number) failing within minutes
- `incident_count` (Number)


<a id="nestedatt--simple_alert--dns_failure"></a>
### Nested Schema for `simple_alert.dns_failure`

Optional:

- `client_count` (Number)
- `duration` (Number) failing within minutes
- `incident_count` (Number)



<a id="nestedatt--skyatp"></a>
### Nested Schema for `skyatp`

Optional:

- `enabled` (Boolean)
- `send_ip_mac_mapping` (Boolean) Whether to send IP-MAC mapping to SkyATP


<a id="nestedatt--sle_thresholds"></a>
### Nested Schema for `sle_thresholds`

Optional:

- `capacity` (Number) Capacity, in %
- `coverage` (Number) Coverage, in dBm
- `throughput` (Number) Throughput, in Mbps
- `timetoconnect` (Number) Time to connect, in seconds


<a id="nestedatt--srx_app"></a>
### Nested Schema for `srx_app`

Optional:

- `enabled` (Boolean)


<a id="nestedatt--ssr"></a>
### Nested Schema for `ssr`

Optional:

- `conductor_hosts` (List of String) List of Conductor IP Addresses or Hosts to be used by the SSR Devices
- `conductor_token` (String, Sensitive) Token to be used by the SSR Devices to connect to the Conductor
- `disable_stats` (Boolean) Disable stats collection on SSR devices


<a id="nestedatt--synthetic_test"></a>
### Nested Schema for `synthetic_test`

Optional:

- `aggressiveness` (String) enum: `auto`, `high`, `low`
- `custom_probes` (Attributes Map) Custom probes to be used for synthetic tests (see [below for nested schema](#nestedatt--synthetic_test--custom_probes))
- `disabled` (Boolean)
- `lan_networks` (Attributes List) List of networks to be used for synthetic tests (see [below for nested schema](#nestedatt--synthetic_test--lan_networks))
- `vlans` (Attributes List) (see [below for nested schema](#nestedatt--synthetic_test--vlans))
- `wan_speedtest` (Attributes) (see [below for nested schema](#nestedatt--synthetic_test--wan_speedtest))

<a id="nestedatt--synthetic_test--custom_probes"></a>
### Nested Schema for `synthetic_test.custom_probes`

Optional:

- `aggressiveness` (String) enum: `auto`, `high`, `low`
- `host` (String) If `type`==`icmp` or `type`==`tcp`, Host to be used for the custom probe
- `port` (Number) If `type`==`tcp`, Port to be used for the custom probe
- `threshold` (Number) In milliseconds
- `type` (String) enum: `curl`, `icmp`, `tcp`
- `url` (String) If `type`==`curl`, URL to be used for the custom probe, can be url or IP


<a id="nestedatt--synthetic_test--lan_networks"></a>
### Nested Schema for `synthetic_test.lan_networks`

Optional:

- `networks` (List of String) List of networks to be used for synthetic tests
- `probes` (List of String) app name comes from `custom_probes` above or /const/synthetic_test_probes


<a id="nestedatt--synthetic_test--vlans"></a>
### Nested Schema for `synthetic_test.vlans`

Optional:

- `custom_test_urls` (List of String, Deprecated)
- `disabled` (Boolean) For some vlans where we don't want this to run
- `probes` (List of String) app name comes from `custom_probes` above or /const/synthetic_test_probes
- `vlan_ids` (List of String)


<a id="nestedatt--synthetic_test--wan_speedtest"></a>
### Nested Schema for `synthetic_test.wan_speedtest`

Optional:

- `enabled` (Boolean)
- `time_of_day` (String) `any` / HH:MM (24-hour format)



<a id="nestedatt--uplink_port_config"></a>
### Nested Schema for `uplink_port_config`

Optional:

- `dot1x` (Boolean) Whether to do 802.1x against uplink switch. When enabled, AP cert will be used to do EAP-TLS and the Org's CA Cert has to be provisioned at the switch
- `keep_wlans_up_if_down` (Boolean) By default, WLANs are disabled when uplink is down. In some scenario, like SiteSurvey, one would want the AP to keep sending beacons.


<a id="nestedatt--vna"></a>
### Nested Schema for `vna`

Optional:

- `enabled` (Boolean) Enable Virtual Network Assistant (using SUB-VNA license). This applied to AP / Switch / Gateway


<a id="nestedatt--vs_instance"></a>
### Nested Schema for `vs_instance`

Optional:

- `networks` (List of String)


<a id="nestedatt--wan_vna"></a>
### Nested Schema for `wan_vna`

Optional:

- `enabled` (Boolean)


<a id="nestedatt--wids"></a>
### Nested Schema for `wids`

Optional:

- `repeated_auth_failures` (Attributes) (see [below for nested schema](#nestedatt--wids--repeated_auth_failures))

<a id="nestedatt--wids--repeated_auth_failures"></a>
### Nested Schema for `wids.repeated_auth_failures`

Optional:

- `duration` (Number) Window where a trigger will be detected and action to be taken (in seconds)
- `threshold` (Number) Count of events to trigger



<a id="nestedatt--wifi"></a>
### Nested Schema for `wifi`

Optional:

- `cisco_enabled` (Boolean)
- `disable_11k` (Boolean) Whether to disable 11k
- `disable_radios_when_power_constrained` (Boolean)
- `enable_arp_spoof_check` (Boolean) When proxy_arp is enabled, check for arp spoofing.
- `enable_shared_radio_scanning` (Boolean)
- `enabled` (Boolean) Enable Wi-Fi feature (using SUB-MAN license)
- `locate_connected` (Boolean) Whether to locate connected clients
- `locate_unconnected` (Boolean) Whether to locate unconnected clients
- `mesh_allow_dfs` (Boolean) Whether to allow Mesh to use DFS channels. For DFS channels, Remote Mesh AP would have to do CAC when scanning for new Base AP, which is slow and will disrupt the connection. If roaming is desired, keep it disabled.
- `mesh_enable_crm` (Boolean) Used to enable/disable CRM
- `mesh_enabled` (Boolean) Whether to enable Mesh feature for the site
- `mesh_psk` (String, Sensitive) Optional passphrase of mesh networking, default is generated randomly
- `mesh_ssid` (String) Optional ssid of mesh networking, default is based on site_id
- `proxy_arp` (String) enum: `default`, `disabled`, `enabled`


<a id="nestedatt--wired_vna"></a>
### Nested Schema for `wired_vna`

Optional:

- `enabled` (Boolean)


<a id="nestedatt--zone_occupancy_alert"></a>
### Nested Schema for `zone_occupancy_alert`

Optional:

- `email_notifiers` (List of String) List of email addresses to send email notifications when the alert threshold is reached
- `enabled` (Boolean) Indicate whether zone occupancy alert is enabled for the site
- `threshold` (Number) Sending zone-occupancy-alert webhook message only if a zone stays non-compliant (i.e. actual occupancy > occupancy_limit) for a minimum duration specified in the threshold, in minutes



## Import
Using `terraform import`, import `mist_site_setting` with:
```shell
# Site Setting can be imported by specifying the site_id
terraform import mist_site_setting.site_setting_one 17b46405-3a6d-4715-8bb4-6bb6d06f316a
```