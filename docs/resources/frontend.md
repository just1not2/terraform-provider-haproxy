---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "haproxy_frontend Resource - terraform-provider-haproxy"
subcategory: ""
description: |-
  Manages a HAProxy frontend.
---

# haproxy_frontend (Resource)

Manages a HAProxy frontend.

## Example Usage

```terraform
resource "haproxy_frontend" "my_frontend" {
  name    = "myhaproxyfrontend"
  backend = "myhaproxybackend"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String) The name of the frontend.

### Optional

- **backend** (String) The name of the backend associated to the frontend.
- **id** (String) The ID of this resource.


