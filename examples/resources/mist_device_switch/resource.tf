
resource "mist_device_switch" "switch_one" {
  name      = "switch_one"
  device_id = mist_org_inventory.inventory.devices[1].id
  site_id   = mist_org_inventory.inventory.devices[1].site_id
  managed = true
  role    = "test"
  networks = {
    "prx" = {
      vlan_id = "18"
    }
  }
  port_usages = {
    "prx" = {
      mode         = "trunk"
      disabled     = false
      port_network = "default"
      stp_edge     = false
      all_networks = false
      networks = [
        "default",
        "prx"
      ],
      speed         = "auto"
      duplex        = "auto"
      mac_limit     = 0
      persist_mac   = false
      poe_disabled  = false
      enable_qos    = false
      storm_control = {}
      description   = ""
    }
  }
  ip_config = {
    type    = "static"
    ip      = "10.3.18.99"
    netmask = "255.255.255.0"
    network = "prx"
    gateway = "10.3.18.11"
  }
  port_config = {
    "ge-0/0/0" = {
      usage                = "prx"
      critical             = false
      no_local_overwrite = true
    },
    "ge-0/0/11" = {
      usage                = "default"
      port_network         = "prx"
      critical             = false
    }
  }
  port_mirroring = {
    "test" = {
      output_port_id = "ge-0/0/10"
      input_port_ids_ingress = [
        "ge-0/0/2"
      ],
      input_port_ids_egress = [
        "ge-0/0/2"
      ],
      input_networks_ingress = [
        "default"
      ]
    }
  }
  mist_nac = {
    enabled = true
  }
  dhcpd_config = {
    enabled = true
    "prx" = {
      type = "relay"
      servers = [
        "1.2.3.4"
      ]
    }
  }
}
