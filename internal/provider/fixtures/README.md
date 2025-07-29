Fixtures

The fixtures are resource specific terraform configuration that are currently used exclusively in the resource tests. In the future there may be fixtures added for datasources as well.
The `_test.go` file corresponding to a specific resource will Decode the configuration stored in the <resource>_config.tf file into go structs and create the specific config in mist to test the Provider code associated with the resource.
Multiple configurations can be added in the file, but need to be separated with ␞ delimeter.

Copilot agents can be leverage to generate valid configuration for tests by refering the agent to the Schema definition in the <resource>_resource_gen.go file. Example of an effective prompt:
```
Refer to the OrgApitokenResourceSchema in org_apitoken_resource_gen.go to create test terraform configuration for an org_apitoken resource and place it in a file under internal/fixtures/org_apitoken_resource. Do not include any comments in the config. Separate config with a ␞ character.
```

Since the test go structs are not aware of the tf boilerplate and the tests add the boilerplate for configuration, this needs to be exclude from the config ie.
```
resource "mist_org_apitoken" "apitoken_one" {
  org_id = mist_org.terraform_test.id
  name   = "apitoken_one"
  privileges = [
    {
      scope   = "site"
      role    = "admin"
      site_id = "d7c8364e-xxxx-xxxx-xxxx-37eff0475b03"
    },
    {
      scope   = "site"
      role    = "read"
      site_id = "08f8851b-xxxx-xxxx-xxxx-9ebb5aa62de4"
    }
  ]
  src_ips = [ "1.2.3.4/32" ]
}
```
should be added as:
```
  org_id = mist_org.terraform_test.id
  name   = "apitoken_one"
  privileges = [
    {
      scope   = "site"
      role    = "admin"
      site_id = "d7c8364e-xxxx-xxxx-xxxx-37eff0475b03"
    },
    {
      scope   = "site"
      role    = "read"
      site_id = "08f8851b-xxxx-xxxx-xxxx-9ebb5aa62de4"
    }
  ]
  src_ips = [ "1.2.3.4/32" ]
```