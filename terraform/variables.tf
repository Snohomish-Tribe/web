variable "region" {
  description = "OCI tenancy"
  type        = string
  default     = "us-sanjose-1"
}

variable "compartment_id" {
  description = "OCID from the tenancy page"
  type        = string
  sensitive   = true
}

variable "config_profile" {
  description = "Profile name"
  type        = string
  sensitive   = true
}

variable "sender_email_address" {
  description = "Sender email address. Should be owned by the same user who generated the FASTMAIL_TOKEN secret."
  type        = string
}
