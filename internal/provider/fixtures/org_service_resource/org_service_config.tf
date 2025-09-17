name        = "test-service-custom"
description = "Test custom service"
type        = "custom"
traffic_type                     = "data"
traffic_class                    = "best_effort"
dscp                             = "21"
failover_policy                  = "revertible"
sle_enabled                      = true
ssr_relaxed_tcp_state_enforcement = true
client_limit_down                = 1000
client_limit_up                  = 500
service_limit_down               = 2000
service_limit_up                 = 1000
max_jitter                       = "10"
max_latency                      = "100"
max_loss                         = "1"

addresses = [
  "192.168.1.0/24",
  "10.0.0.0/8"
]

hostnames = [
  "example.com",
  "test.example.org"
]

specs = [
  {
    protocol = "tcp"
    port_range = "80-443"
  },
  {
    protocol = "udp"  
    port_range = "53-53"
  },
  {
    protocol = "https"
  }
]

␞

name = "test-service-apps"
description = "Test apps service"
type = "apps"
sle_enabled = true

apps = [
  "slack",
  "zoom",
  "teams"
]

␞

name = "test-service-app-categories"
description = "Test app categories service"
type = "app_categories"
sle_enabled = true

app_categories = [
  "business",
  "productivity"
]

app_subcategories = [
  "collaboration",
  "file_sharing"
]

␞

name = "test-service-urls"
description = "Test urls service"
type = "urls"
sle_enabled = true

urls = [
  "https://api.example.com",
  "https://service.test.com"
]