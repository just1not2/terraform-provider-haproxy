# HAProxy Terraform Provider

The HAproxy Terraform provider helps the management of HAProxy resources on Terraform. It interacts with HAProxy controllers via the Data Plane API.

## Terraform version compatibility

This collection has been tested against following Terraform versions: **>=0.14**.


## Included content

### Data Sources
Name | Description
--- | ---
[haproxy_frontend](./docs/data-sources/frontend.md)|Fetches information about a HAProxy frontend

### Resources
Name | Description
--- | ---
[haproxy_frontend](./docs/resources/frontend.md)|Manages a HAProxy frontend


## Installing this provider

You can install this provider by cloning this repository and replacing the `GNUmakefile` template with your architecture information. You can then create build the provider:

    make install


## Using this provider

In order to use this provider, the Data Plane API must be installed on your controller. For more information about the installation, check [the official guide](https://www.haproxy.com/documentation/hapee/latest/api/data-plane-api/installation/).


## See Also

* [HAProxy Data Plane API official documentation](https://www.haproxy.com/documentation/dataplaneapi)
* [Terraform Using providers](https://www.terraform.io/language/providers)


## Contributing to this collection

This provider started as personal project, but I welcome community contributions to this provider. If you find problems, please open an issue or create a PR against the [HAProxy provider repository](https://github.com/just1not2/terraform-provider-haproxy).

You can also reach me by email at `me@just1not2.org`.


## Licensing

Mozilla Public License Version 2.0.

See [LICENSE](./LICENSE) to see the full text.
