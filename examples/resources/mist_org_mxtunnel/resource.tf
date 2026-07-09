resource "mist_org_mxtunnel" "mxtunnel_one" {
  org_id         = mist_org.terraform_test.id
  name           = "mxtunnel_one"
  protocol       = "udp"
  hello_interval = 60
  hello_retries  = 7
  mtu            = 0
  vlan_ids       = [10, 20, 30]
  mxcluster_ids  = [mist_org_mxcluster.mxcluster_one.id]

  auto_preemption = {
    enabled     = true
    day_of_week = "mon"
    time_of_day = "02:00"
  }

  ipsec = {
    enabled      = true
    split_tunnel = true
    use_mxedge   = true
    dns_servers  = ["8.8.8.8", "8.8.4.4"]
    dns_suffix   = ["corp.example.com"]
    extra_routes = [
      {
        dest     = "10.0.0.0/8"
        next_hop = "192.168.1.1"
      }
    ]
  }
}
