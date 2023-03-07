terraform {
  required_providers {
    idcloudhost = {
      version = "0.2.0"
      source  = "bonestealer.xyz/muhammad-asn/idcloudhost"
    }
  }
}

provider "idcloudhost" {}

resource "idcloudhost_s3" "test-bucket-terraform-66" {
  name               = "test-bucket-terraform-23"
  billing_account_id = 1200190928
}

resource "idcloudhost_s3" "test-bucket-terraform-67" {
  name               = "test-bucket-terraform-24"
  billing_account_id = 1200190928
}