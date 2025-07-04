---
page_title: "mist_org_setting Resource - terraform-provider-mist"
subcategory: "Org"
description: |-
  This resource manages the Org Settings.
  The Org Settings can be used to customize the Org configuration
---

# mist_org_setting (Resource)

This resource manages the Org Settings.

The Org Settings can be used to customize the Org configuration


## Example Usage

```terraform
resource "mist_org_setting" "terraform_test" {
  org_id              = mist_org.terraform_test.id
  ap_updown_threshold = 10
  cradlepoint = {
    cp_api_id   = "cp_api_id_test"
    cp_api_key  = "secret"
    ecm_api_id  = "ecm_api_id_test"
    ecm_api_key = "secret"
  }
  device_updown_threshold  = 10
  disable_pcap             = false
  disable_remote_shell     = true
  gateway_updown_threshold = 10
  mxedge_mgmt = {
    mist_password = "root_secret_password"
    root_password = "root_secret_password"
    oob_ip_type   = "dhcp"
    oob_ip_type6  = "disabled"
  }
  password_policy = {
    enabled                  = true
    freshness                = 180
    min_length               = 12
    requires_special_char    = true
    requires_two_factor_auth = false
  }
  security = {
    disable_local_ssh = true
  }
  switch_updown_threshold = 10
  synthetic_test = {
    disabled = false
    vlans = [{
      vlan_ids         = ["10", "30"]
      custom_test_urls = ["http://www.abc.com/", "https://10.3.5.1:8080/about"]
      }, {
      vlan_ids = ["20"]
      disabled = true
    }]
  }
  ui_idle_timeout = 120
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `org_id` (String)

### Optional

- `ap_updown_threshold` (Number) Enable threshold-based device down delivery for AP devices only. When configured it takes effect for AP devices and `device_updown_threshold` is ignored.
- `api_policy` (Attributes) (see [below for nested schema](#nestedatt--api_policy))
- `cacerts` (List of String) RADSec certificates for AP
- `celona` (Attributes) (see [below for nested schema](#nestedatt--celona))
- `cloudshark` (Attributes) (see [below for nested schema](#nestedatt--cloudshark))
- `device_cert` (Attributes) common device cert, optional (see [below for nested schema](#nestedatt--device_cert))
- `device_updown_threshold` (Number) Enable threshold-based device down delivery via
  * device-updowns webhooks topic, 
  * Mist Alert Framework; e.g. send AP/SW/GW down event only if AP/SW/GW Up is not seen within the threshold in minutes; 0 - 240, default is 0 (trigger immediate)
- `disable_pcap` (Boolean) Whether to disallow Mist to analyze pcap files (this is required for marvis pcap)
- `disable_remote_shell` (Boolean) Whether to disable remote shell access for an entire org
- `gateway_updown_threshold` (Number) Enable threshold-based device down delivery for Gateway devices only. When configured it takes effect for GW devices and `device_updown_threshold` is ignored.
- `installer` (Attributes) (see [below for nested schema](#nestedatt--installer))
- `jcloud` (Attributes) (see [below for nested schema](#nestedatt--jcloud))
- `jcloud_ra` (Attributes) JCloud Routing Assurance connexion (see [below for nested schema](#nestedatt--jcloud_ra))
- `junos_shell_access` (Attributes) by default, webshell access is only enabled for Admin user (see [below for nested schema](#nestedatt--junos_shell_access))
- `marvis` (Attributes) (see [below for nested schema](#nestedatt--marvis))
- `mgmt` (Attributes) management-related properties (see [below for nested schema](#nestedatt--mgmt))
- `mist_nac` (Attributes) (see [below for nested schema](#nestedatt--mist_nac))
- `mxedge_mgmt` (Attributes) (see [below for nested schema](#nestedatt--mxedge_mgmt))
- `optic_port_config` (Attributes Map) Property key is the interface name or range (e.g. `et-0/0/47`, `et-0/0/48-49`) (see [below for nested schema](#nestedatt--optic_port_config))
- `password_policy` (Attributes) password policy (see [below for nested schema](#nestedatt--password_policy))
- `pcap` (Attributes) (see [below for nested schema](#nestedatt--pcap))
- `security` (Attributes) (see [below for nested schema](#nestedatt--security))
- `ssr` (Attributes) (see [below for nested schema](#nestedatt--ssr))
- `switch` (Attributes) (see [below for nested schema](#nestedatt--switch))
- `switch_mgmt` (Attributes) (see [below for nested schema](#nestedatt--switch_mgmt))
- `switch_updown_threshold` (Number) Enable threshold-based device down delivery for Switch devices only. When configured it takes effect for SW devices and `device_updown_threshold` is ignored.
- `synthetic_test` (Attributes) (see [below for nested schema](#nestedatt--synthetic_test))
- `ui_idle_timeout` (Number) Automatically logout the user when UI session is inactive. `0` means disabled
- `vpn_options` (Attributes) (see [below for nested schema](#nestedatt--vpn_options))
- `wan_pma` (Attributes) (see [below for nested schema](#nestedatt--wan_pma))
- `wired_pma` (Attributes) (see [below for nested schema](#nestedatt--wired_pma))
- `wireless_pma` (Attributes) (see [below for nested schema](#nestedatt--wireless_pma))

### Read-Only

- `cradlepoint` (Attributes) (see [below for nested schema](#nestedatt--cradlepoint))
- `juniper` (Attributes) (see [below for nested schema](#nestedatt--juniper))

<a id="nestedatt--api_policy"></a>
### Nested Schema for `api_policy`

Optional:

- `no_reveal` (Boolean) By default, API hides password/secrets when the user doesn't have write access
  * `true`: API will hide passwords/secrets for all users
  * `false`: API will hide passwords/secrets for read-only users


<a id="nestedatt--celona"></a>
### Nested Schema for `celona`

Required:

- `api_key` (String)
- `api_prefix` (String)


<a id="nestedatt--cloudshark"></a>
### Nested Schema for `cloudshark`

Optional:

- `apitoken` (String, Sensitive)
- `url` (String) If using CS Enterprise


<a id="nestedatt--device_cert"></a>
### Nested Schema for `device_cert`

Required:

- `cert` (String)
- `key` (String, Sensitive)


<a id="nestedatt--installer"></a>
### Nested Schema for `installer`

Optional:

- `allow_all_devices` (Boolean)
- `allow_all_sites` (Boolean)
- `extra_site_ids` (List of String)
- `grace_period` (Number)


<a id="nestedatt--jcloud"></a>
### Nested Schema for `jcloud`

Required:

- `org_apitoken` (String) JCloud Org Token
- `org_apitoken_name` (String) JCloud Org Token Name
- `org_id` (String) JCloud Org ID


<a id="nestedatt--jcloud_ra"></a>
### Nested Schema for `jcloud_ra`

Optional:

- `org_apitoken` (String) JCloud Routing Assurance Org Token
- `org_apitoken_name` (String) JCloud Routing Assurance Org Token Name
- `org_id` (String) JCloud Routing Assurance Org ID


<a id="nestedatt--junos_shell_access"></a>
### Nested Schema for `junos_shell_access`

Optional:

- `admin` (String) enum: `admin`, `viewer`, `none`
- `helpdesk` (String) enum: `admin`, `viewer`, `none`
- `read` (String) enum: `admin`, `viewer`, `none`
- `write` (String) enum: `admin`, `viewer`, `none`


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



<a id="nestedatt--mgmt"></a>
### Nested Schema for `mgmt`

Optional:

- `mxtunnel_ids` (List of String) List of Mist Tunnels
- `use_mxtunnel` (Boolean) Whether to use Mist Tunnel for mgmt connectivity, this takes precedence over use_wxtunnel
- `use_wxtunnel` (Boolean) Whether to use wxtunnel for mgmt connectivity


<a id="nestedatt--mist_nac"></a>
### Nested Schema for `mist_nac`

Optional:

- `cacerts` (List of String) List of PEM-encoded ca certs
- `default_idp_id` (String) use this IDP when no explicit realm present in the incoming username/CN OR when no IDP is explicitly mapped to the incoming realm.
- `disable_rsae_algorithms` (Boolean) to disable RSAE_PSS_SHA256, RSAE_PSS_SHA384, RSAE_PSS_SHA512 from server side. see https://www.openssl.org/docs/man3.0/man1/openssl-ciphers.html
- `eap_ssl_security_level` (Number) eap ssl security level, see https://www.openssl.org/docs/man1.1.1/man3/SSL_CTX_set_security_level.html#DEFAULT-CALLBACK-BEHAVIOUR
- `eu_only` (Boolean) By default, NAC POD failover considers all NAC pods available around the globe, i.e. EU, US, or APAC based, failover happens based on geo IP of the originating site. For strict GDPR compliance NAC POD failover would only happen between the PODs located within the EU environment, and no authentication would take place outside of EU. This is an org setting that is applicable to WLANs, switch templates, mxedge clusters that have mist_nac enabled
- `idp_machine_cert_lookup_field` (String) allow customer to choose the EAP-TLS client certificate's field to use for IDP Machine Groups lookup. enum: `automatic`, `cn`, `dns`
- `idp_user_cert_lookup_field` (String) allow customer to choose the EAP-TLS client certificate's field. To use for IDP User Groups lookup. enum: `automatic`, `cn`, `email`, `upn`
- `idps` (Attributes List) (see [below for nested schema](#nestedatt--mist_nac--idps))
- `server_cert` (Attributes) radius server cert to be presented in EAP TLS (see [below for nested schema](#nestedatt--mist_nac--server_cert))
- `use_ip_version` (String) by default, NAS devices(switches/aps) and proxies(mxedge) are configured to reach mist-nac via IPv4. enum: `v4`, `v6`
- `use_ssl_port` (Boolean) By default, NAS devices (switches/aps) and proxies(mxedge) are configured to use port TCP2083(RadSec) to reach mist-nac. Set `use_ssl_port`==`true` to override that port with TCP43 (ssl), This is an org level setting that is applicable to wlans, switch_templates, and mxedge_clusters that have mist-nac enabled

<a id="nestedatt--mist_nac--idps"></a>
### Nested Schema for `mist_nac.idps`

Required:

- `id` (String) ID of the `mist_nacidp`
- `user_realms` (List of String) Which realm should trigger this IDP. User Realm is extracted from:
  * Username-AVP (`mist.com` from john@mist.com)
  * Cert CN

Optional:

- `exclude_realms` (List of String) When the IDP of mxedge_proxy type, exclude the following realms from proxying in addition to other valid home realms in this org


<a id="nestedatt--mist_nac--server_cert"></a>
### Nested Schema for `mist_nac.server_cert`

Optional:

- `cert` (String)
- `key` (String, Sensitive)
- `password` (String, Sensitive) private key password (optional)



<a id="nestedatt--mxedge_mgmt"></a>
### Nested Schema for `mxedge_mgmt`

Optional:

- `config_auto_revert` (Boolean)
- `fips_enabled` (Boolean)
- `mist_password` (String, Sensitive)
- `oob_ip_type` (String) enum: `dhcp`, `disabled`, `static`
- `oob_ip_type6` (String) enum: `autoconf`, `dhcp`, `disabled`, `static`
- `root_password` (String, Sensitive)


<a id="nestedatt--optic_port_config"></a>
### Nested Schema for `optic_port_config`

Optional:

- `channelized` (Boolean) Enable channelization
- `speed` (String) Interface speed (e.g. `25g`, `50g`), use the chassis speed by default


<a id="nestedatt--password_policy"></a>
### Nested Schema for `password_policy`

Optional:

- `enabled` (Boolean) Whether the policy is enabled
- `expiry_in_days` (Number) password expiry in days
- `min_length` (Number) Required password length
- `requires_special_char` (Boolean) Whether to require special character
- `requires_two_factor_auth` (Boolean) Whether to require two-factor auth


<a id="nestedatt--pcap"></a>
### Nested Schema for `pcap`

Optional:

- `bucket` (String)
- `max_pkt_len` (Number) Max_len of non-management packets to capture


<a id="nestedatt--security"></a>
### Nested Schema for `security`

Optional:

- `disable_local_ssh` (Boolean) Whether to disable local SSH (by default, local SSH is enabled with allow_mist in Org is enabled
- `fips_zeroize_password` (String, Sensitive) password required to zeroize devices (FIPS) on site level
- `limit_ssh_access` (Boolean) Whether to allow certain SSH keys to SSH into the AP (see Site:Setting)


<a id="nestedatt--ssr"></a>
### Nested Schema for `ssr`

Optional:

- `conductor_hosts` (List of String) List of Conductor IP Addresses or Hosts to be used by the SSR Devices
- `conductor_token` (String, Sensitive) Token to be used by the SSR Devices to connect to the Conductor
- `disable_stats` (Boolean) Disable stats collection on SSR devices


<a id="nestedatt--switch"></a>
### Nested Schema for `switch`

Optional:

- `auto_upgrade` (Attributes) (see [below for nested schema](#nestedatt--switch--auto_upgrade))

<a id="nestedatt--switch--auto_upgrade"></a>
### Nested Schema for `switch.auto_upgrade`

Optional:

- `custom_versions` (Map of String) Custom version to be used. The Property Key is the switch hardware and the property value is the firmware version
- `enabled` (Boolean) Enable auto upgrade for the switch
- `snapshot` (Boolean) Enable snapshot during the upgrade process



<a id="nestedatt--switch_mgmt"></a>
### Nested Schema for `switch_mgmt`

Optional:

- `ap_affinity_threshold` (Number) If the field is set in both site/setting and org/setting, the value from site/setting will be used.


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



<a id="nestedatt--vpn_options"></a>
### Nested Schema for `vpn_options`

Optional:

- `as_base` (Number)
- `st_subnet` (String) requiring /12 or bigger to support 16 private IPs for 65535 gateways


<a id="nestedatt--wan_pma"></a>
### Nested Schema for `wan_pma`

Optional:

- `enabled` (Boolean)


<a id="nestedatt--wired_pma"></a>
### Nested Schema for `wired_pma`

Optional:

- `enabled` (Boolean)


<a id="nestedatt--wireless_pma"></a>
### Nested Schema for `wireless_pma`

Optional:

- `enabled` (Boolean)


<a id="nestedatt--cradlepoint"></a>
### Nested Schema for `cradlepoint`

Read-Only:

- `cp_api_id` (String)
- `cp_api_key` (String, Sensitive)
- `ecm_api_id` (String)
- `ecm_api_key` (String, Sensitive)
- `enable_lldp` (Boolean)


<a id="nestedatt--juniper"></a>
### Nested Schema for `juniper`

Read-Only:

- `accounts` (Attributes List) (see [below for nested schema](#nestedatt--juniper--accounts))

<a id="nestedatt--juniper--accounts"></a>
### Nested Schema for `juniper.accounts`

Read-Only:

- `linked_by` (String)
- `name` (String)



## Import
Using `terraform import`, import `mist_org_setting` with:
```shell
# Org Setting can be imported by specifying the org_id
terraform import mist_org_setting.setting_one 17b46405-3a6d-4715-8bb4-6bb6d06f316a
```