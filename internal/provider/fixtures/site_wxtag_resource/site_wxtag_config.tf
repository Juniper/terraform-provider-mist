name = "wxtag_comprehensive_match"
type = "match"
match = "ip_range_subnet"
op = "in"
values = ["192.168.1.0/24", "10.0.0.0/8", "172.16.50.1"]
␞
name = "wxtag_comprehensive_spec"
type = "spec"
specs = [
  {
    port_range = "80"
    protocol = "tcp"
    subnets = ["192.168.1.0/24", "192.168.2.0/24"]
  },
  {
    port_range = "443"
    protocol = "tcp"
    subnets = ["10.0.0.0/8"]
  },
  {
    port_range = "53"
    protocol = "udp"
    subnets = ["8.8.8.8", "8.8.4.4"]
  },
  {
    port_range = "0"
    protocol = "icmp"
    subnets = ["0.0.0.0/0"]
  },
  {
    port_range = "8080-8090"
    protocol = "tcp"
    subnets = ["172.16.0.0/16"]
  }
]
␞
name = "wxtag_mac_vlan"
type = "match"
match = "wlan_id"
op = "in"
values = ["guest-wifi"]
mac = "5c5b35000001"
vlan_id = "100"
