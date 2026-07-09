
  name     = "test-mxtunnel"

  # Optional basic fields
  hello_interval = 60
  hello_retries  = 7
  mtu            = 0
  protocol       = "udp"
  vlan_ids       = [100, 200, 300]

  # Auto preemption scheduling
  auto_preemption = {
    day_of_week = "mon"
    enabled     = true
    time_of_day = "02:00"
  }

  # IPsec VPN configuration
  ipsec = {
    dns_servers  = ["8.8.8.8", "8.8.4.4"]
    dns_suffix   = ["corp.example.com"]
    enabled      = true
    split_tunnel = true
    use_mxedge   = true
    extra_routes = [
      {
        dest     = "10.0.0.0/8"
        next_hop = "192.168.1.1"
      },
      {
        dest     = "172.16.0.0/12"
        next_hop = "192.168.1.254"
      }
    ]
  }
