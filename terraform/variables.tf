variable "region" {
    description = "OCI tenancy"
    type = string
    default = "us-sanjose-1"
}

variable "compartment_id" {
    description = "OCID from the tenancy page"
    type = string
    sensitive = true
}

variable "config_profile" {
  description = "Profile name"
  type = string
  sensitive = true
}

variable "security_token" {
  description = "Security Token"
  type = string
  sensitive = true
}