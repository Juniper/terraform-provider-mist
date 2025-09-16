
  name   = "test-org-vpn2"
  type   = "hub_spoke"
  paths = {
    "vpn_path_1" = {
      bfd_profile         = "broadband"
      bfd_use_tunnel_mode = true
      ip                  = "192.168.1.1"
      pod                 = 1
      peer_paths = {
        "peer_interface_1" = {
          preference = 100
        }
        "peer_interface_2" = {
          preference = 200
        }
      }
      traffic_shaping = {
        enabled          = true
        max_tx_kbps      = 10000
        class_percentage = [25, 25, 25, 25]
      }
    }
    
    "vpn_path_2" = {
      bfd_profile         = "lte"
      bfd_use_tunnel_mode = false
      ip                  = "192.168.2.1"
      pod                 = 2
      peer_paths = {
        "peer_interface_3" = {
          preference = 150
        }
      }
      traffic_shaping = {
        enabled          = false
        max_tx_kbps      = 5000
        class_percentage = [30, 30, 20, 20]
      }
    }
  }
␞
  name   = "test-org-vpn-mesh2"
  type   = "mesh"
  path_selection = {
    strategy = "manual"
  }
  paths = {
    "interface_1" = {
      bfd_profile         = "broadband"
      bfd_use_tunnel_mode = true
      ip                  = "10.0.1.1"
      pod                 = 5
      peer_paths = {
        "mesh_peer_1" = {
          preference = 100
        }
        "mesh_peer_2" = {
          preference = 200
        }
      }
      traffic_shaping = {
        enabled          = true
        max_tx_kbps      = 20000
        class_percentage = [40, 30, 20, 10]
      }
    }
    
    "interface_2" = {
      bfd_profile         = "lte"
      bfd_use_tunnel_mode = false
      ip                  = "10.0.2.1"
      pod                 = 10
      peer_paths = {
        "mesh_peer_3" = {
          preference = 300
        }
      }
    }
  }

