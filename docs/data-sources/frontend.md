---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "haproxy_frontend Data Source - terraform-provider-haproxy"
subcategory: ""
description: |-
  Fetches information about a HAProxy frontend.
---

# haproxy_frontend (Data Source)

Fetches information about a HAProxy frontend.

## Example Usage

```terraform
data "haproxy_frontend" "my_frontend" {
  name = "myhaproxyfrontend"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String) The name of the frontend.

### Optional

- **backend** (String) The name of the backend associated to the frontend.
- **id** (String) The ID of this resource.


