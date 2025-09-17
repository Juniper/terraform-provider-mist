action = "allow"
apply_tags = ["test-apply-tag-1", "test-apply-tag-2", "security-tag"]
enabled = true
guest_auth_state = "authorized"
name = "test-nacrule"
order = 100

matching = {
    auth_type = "cert"
    family = ["juniper", "aruba"]
    mfg = ["Juniper Networks", "Aruba Networks"]
    model = ["EX4300", "2930F"]
    nactags = ["test-nactag-1", "test-nactag-2", "network-tag"]
    os_type = ["junos", "arubaos"]
    port_types = ["trunk", "access"]
    vendor = ["juniper", "aruba"]
    site_ids = ["{site_id}"]
    sitegroup_ids = ["{sitegroup_id}"]
}

not_matching = {
    auth_type = "mab"
    family = ["cisco", "hp"]
    mfg = ["Cisco Systems", "Hewlett Packard"]
    model = ["C9300", "ProCurve"]
    nactags = ["excluded-tag"]
    os_type = ["ios", "procurve"]
    vendor = ["cisco", "hp"]
    port_types = ["access"]
    site_ids = ["{site_id}"]
    sitegroup_ids = ["{sitegroup_id}"]
}