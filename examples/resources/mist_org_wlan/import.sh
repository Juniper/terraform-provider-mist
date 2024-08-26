# Gateway cluster can be imported by specifying the org_id and the wlan_id
terraform import mist_org_wlan.wlan_one 17b46405-3a6d-4715-8bb4-6bb6d06f316a.d3c42998-9012-4859-9743-6b9bee475309
```


In Terraform v1.5.0 and later, use an import block to import `mist_org_wlan` with `id`=`{org_id}.{wlan_id}`:

```tf
import {
  to = mist_org_wlan.wlan_one
  id = "17b46405-3a6d-4715-8bb4-6bb6d06f316a.d3c42998-9012-4859-9743-6b9bee475309"
}
