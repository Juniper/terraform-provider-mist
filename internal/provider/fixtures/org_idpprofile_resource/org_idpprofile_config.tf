base_profile = "critical"
name = "test_idp_profile_comprehensive"
overwrites = [
  {
    name = "overwrite_alert_rule"
    action = "alert"
    matching = {
      attack_name = ["SQL_Injection", "Cross_Site_Scripting"]
      dst_subnet = ["192.168.1.0/24", "10.0.0.0/8"]
      severity = ["high", "critical"]
    }
  },
  {
    name = "overwrite_drop_rule"
    action = "drop"
    matching = {
      attack_name = ["Malware_Detection"]
      dst_subnet = ["172.16.0.0/12"]
      severity = ["critical"]
    }
  },
  {
    name = "overwrite_close_rule"
    action = "close"
    matching = {
      attack_name = ["Brute_Force_Attack", "DDoS_Attack"]
      dst_subnet = ["192.168.0.0/16"]
      severity = ["medium", "high"]
    }
  }
]
␞
base_profile = "strict"
name = "test_idp_profile_strict"
org_id = "901c5705-ca11-4bf1-9158-31f7195618ef"
overwrites = [
  {
    name = "single_overwrite_minimal"
    matching = {
      attack_name = ["Buffer_Overflow"]
    }
  }
]
␞
base_profile = "standard"
name = "test_idp_profile_partial"
org_id = "901c5705-ca11-4bf1-9158-31f7195618ef"
overwrites = [
  {
    name = "subnet_only_rule"
    action = "drop"
    matching = {
      dst_subnet = ["203.0.113.0/24"]
    }
  },
  {
    name = "severity_only_rule"
    action = "close"
    matching = {
      severity = ["low"]
    }
  }
]
