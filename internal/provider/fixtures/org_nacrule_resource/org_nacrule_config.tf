action = "block"
name = "test-nacrule-not-matching"
order = 200

not_matching = {
    auth_type     = "mab"
    family        = ["cisco", "hp"]
    mfg           = ["Cisco Systems", "Hewlett Packard"]
    model         = ["C9300", "ProCurve"]
    nactags       = ["excluded-tag"]
    os_type       = ["ios", "procurve"]
    vendor        = ["cisco", "hp"]
    port_types    = ["access"]
    site_ids      = ["{site_id}"]
    sitegroup_ids = ["{sitegroup_id}"]
}