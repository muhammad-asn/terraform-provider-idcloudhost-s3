// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package idcloudhost

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/muhammad-asn/idcloudhost-go-lib/api"
	"github.com/muhammad-asn/idcloudhost-go-lib/s3"
)

func resourceS3() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceS3Create,
		ReadContext:   resourceS3Read,
		UpdateContext: resourceS3Update,
		DeleteContext: resourceS3Delete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"size_bytes": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"billing_account_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"num_objects": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_suspended": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func resourceS3Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*api.APIClient)

	var diags diag.Diagnostics

	newS3Bucket := &s3.S3Bucket{
		Name:             d.Get("name").(string),
		BillingAccountId: d.Get("billing_account_id").(int),
	}

	s3Api := client.S3

	if err := s3Api.Create(*newS3Bucket); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create new S3 Bucket",
			Detail:   "",
		})

		return diags
	}

	d.SetId(newS3Bucket.Name)

	resourceS3Read(ctx, d, m)

	return diags
}

func resourceS3Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	time.Sleep(5 * time.Second)

	client := m.(*api.APIClient)
	var diags diag.Diagnostics

	bucketID := d.Id()

	s3Api := client.S3
	if err := s3Api.Get(bucketID); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to get S3 Bucket",
			Detail:   "",
		})

		return diags
	}

	return diags
}

func resourceS3Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*api.APIClient)
	var diags diag.Diagnostics

	bucketID := d.Id()

	if d.HasChange("billing_account_id") {

		s3Api := client.S3

		updatedS3Bucket := &s3.S3Bucket{
			Name:             bucketID,
			BillingAccountId: d.Get("billing_account_id").(int),
		}

		if err := s3Api.Modify(*updatedS3Bucket); err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to change the billing account ID",
				Detail:   "",
			})

			return diags
		}

		d.Set("modified_at", time.Now().Format(time.RFC850))

	}

	return resourceS3Read(ctx, d, m)
}

func resourceS3Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*api.APIClient)
	var diags diag.Diagnostics

	bucketID := d.Id()

	deletedS3Bucket := &s3.S3Bucket{
		Name:             bucketID,
		BillingAccountId: d.Get("billing_account_id").(int),
	}

	s3Api := client.S3
	if err := s3Api.Delete(*deletedS3Bucket); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to delete S3 Bucket",
			Detail:   "",
		})

		return diags
	}

	d.SetId("")
	return diags
}
