// Copyright: (c) 2022, Justin Béra (@just1not2) <me@just1not2.org>
// Mozilla Public License Version 2.0 (see LICENSE or https://www.mozilla.org/en-US/MPL/2.0/)

package haproxy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"password": {
				Description: "The password to login with. If omitted, the `HAPROXY_PASSWORD` environment variable is used.",
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("HAPROXY_PASSWORD", nil),
			},
			"url": {
				Description: "The URL of the HAProxy controller. If omitted, the `HAPROXY_URL` environment variable is used.",
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HAPROXY_URL", nil),
			},
			"username": {
				Description: "The username to login with. If omitted, the `HAPROXY_USERNAME` environment variable is used.",
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("HAPROXY_USERNAME", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"haproxy_frontend": resourceFrontend(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"haproxy_frontend": dataSourceFrontend(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Gets the HAProxy API configuration
	username := d.Get("username")
	password := d.Get("password")
	url := d.Get("url")

	// Creates the HAProxy client
	client := NewHAProxyClient(url, username, password)

	return client, diags
}
