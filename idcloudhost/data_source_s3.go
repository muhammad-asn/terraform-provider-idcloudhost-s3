// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package idcloudhost

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceS3() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample data source in the Terraform provider scaffolding.",

		ReadContext: dataSourceS3Read,

		Schema: map[string]*schema.Schema{
			"sample_attribute": {
				// This description is used by the documentation generator and the language server.
				Description: "Sample attribute.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func dataSourceS3Read(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	idFromAPI := "my-id"
	d.SetId(idFromAPI)

	return diag.Errorf("not implemented")
}
