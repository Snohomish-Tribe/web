terraform {
  required_providers {
    oci = {
      source = "oracle/oci"
    }
  }

  backend "s3" {
    bucket                      = "terraform"
    region                      = "us-sanjose-1"
    key                         = "terraform.tfstate"
    skip_region_validation      = true
    skip_credentials_validation = true
    skip_requesting_account_id  = true
    use_path_style              = true
    skip_s3_checksum            = true
    skip_metadata_api_check     = true
    endpoints = {
      s3 = "https://axihvv9biq8w.compat.objectstorage.us-sanjose-1.oraclecloud.com"
    }
  }
}

provider "oci" {
  region              = var.region
  auth                = "SecurityToken"
  config_file_profile = var.config_profile
}
