org_id = "{org_id}"
name = "test-nactag-match"
type = "match"
match = "client_mac"
match_all = false
values = ["aa:bb:cc:dd:ee:ff", "11:22:33:44:55:66"]
␞
org_id = "{org_id}"
name = "test-nactag-radius"
type = "radius_attrs"
radius_attrs = ["Tunnel-Type=13", "Tunnel-Medium-Type=6"]
␞
org_id = "{org_id}"
name = "test-nactag-vendor"
type = "radius_vendor_attrs"
radius_vendor_attrs = ["26:9:1=test"]
␞
org_id = "{org_id}"
name = "test-nactag-egress"
type = "egress_vlan_names"
egress_vlan_names = ["guest", "iot"]
␞
org_id = "{org_id}"
name = "test-nactag-gbp"
type = "gbp_tag"
gbp_tag = "100"
␞
org_id = "{org_id}"
name = "test-nactag-session"
type = "session_timeout"
session_timeout = 3600
␞
org_id = "{org_id}"
name = "test-nactag-group"
type = "radius_group"
radius_group = "guest_users"
␞
org_id = "{org_id}"
name = "test-nactag-username"
type = "username_attr"
username_attr = "email"
␞
org_id = "{org_id}"
name = "test-nactag-vlan"
type = "vlan"
vlan = "100"
␞
org_id = "{org_id}"
name = "test-nactag-portal"
type = "redirect_guest_portal"
nacportal_id = "11111111-1111-1111-1111-111111111111"