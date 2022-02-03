// Copyright: (c) 2022, Justin Béra (@just1not2) <me@just1not2.org>
// Mozilla Public License Version 2.0 (see LICENSE or https://www.mozilla.org/en-US/MPL/2.0/)

package haproxy

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFrontend() *schema.Resource {
	return &schema.Resource{
		Description: "Fetches information about a HAProxy frontend.",
		ReadContext: dataSourceFrontendRead,
		Schema: map[string]*schema.Schema{
			"backend": {
				Description: "The name of the backend associated to the frontend.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": {
				Description: "The name of the frontend.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func dataSourceFrontendRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	name := d.Get("name").(string)

	// Initializes the return body
	returnBody := make(map[string]interface{}, 0)

	// Sends the request
	client := m.(*HAProxyClient)
	if err := client.Request("GET", fmt.Sprintf("/services/haproxy/configuration/frontends/%s", name), nil, &returnBody); err != nil {
		return diag.FromErr(err)
	}

	// Gets the return data
	data := returnBody["data"].(map[string]interface{})
	if data["default_backend"] != nil {
		if err := d.Set("backend", data["default_backend"].(string)); err != nil {
			return diag.FromErr(err)
		}
	}

	// Sets the resource name as Terraform identifier
	d.SetId(name)

	return diags
}
