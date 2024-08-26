# Gateway cluster can be imported by specifying the org_id
terraform import mist_org.org_one 17b46405-3a6d-4715-8bb4-6bb6d06f316a
```


In Terraform v1.5.0 and later, use an import block to import `mist_org` with `id`=`{org_id}` :

```tf
import {
  to = mist_org.org_one
  id = "17b46405-3a6d-4715-8bb4-6bb6d06f316a"
}
