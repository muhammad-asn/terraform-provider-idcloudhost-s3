// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package idcloudhost

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	s3API "github.com/muhammad-asn/idcloudhost-go-lib/api"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"auth_token": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("IDCLOUDHOST_AUTH_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"idcloudhost_s3": resourceS3(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	authToken := d.Get("auth_token").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if authToken != "" {
		c, err := s3API.NewClient(authToken)
		if err != nil {
			return nil, diag.FromErr(err)
		}
		return c, diags
	}
	c, err := s3API.NewClient("")
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return c, diags
}
