# Gateway cluster can be imported by specifying the site_id
terraform import mist_site_setting.site_setting_one 17b46405-3a6d-4715-8bb4-6bb6d06f316a
```


In Terraform v1.5.0 and later, use an import block to import `mist_site_setting` with `id={site_id}`:

```tf
import {
  to = mist_site_setting.site_setting_one
  id = "17b46405-3a6d-4715-8bb4-6bb6d06f316a
}
