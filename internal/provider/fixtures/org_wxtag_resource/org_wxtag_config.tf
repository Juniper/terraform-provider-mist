# Comprehensive Org WxTag Resource Test Configuration
# This fixture tests all major fields of the OrgWxtag resource

name   = "Comprehensive_WxTag_Test"
type   = "match"
match  = "client_mac"
op     = "in"
values = ["aa:bb:cc:dd:ee:ff", "11:22:33:44:55:66", "77:88:99:aa:bb:cc"]
specs = [
  {
    protocol   = "tcp"
    port_range = "80-443"
    subnets    = ["10.0.1.0/24", "10.0.2.0/24"]
  },
  {
    protocol   = "udp"
    port_range = "53"
    subnets    = ["192.168.1.0/24"]
  }
]
