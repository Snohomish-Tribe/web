resource "oci_artifacts_container_repository" "web" {
  display_name   = "web"
  compartment_id = var.compartment_id
  is_immutable   = false
  is_public      = false
}

output "container_repo_image_tag" {
  value = local.container_image_tag
}

locals {
  container_image_tag   = "ocir.${var.region}.oci.oraclecloud.com/${oci_artifacts_container_repository.web.namespace}/${oci_artifacts_container_repository.web.display_name}"
  fastmail_token_secret = data.oci_vault_secrets.fastmail_token.secrets[0]
}

data "oci_vault_secrets" "fastmail_token" {
  compartment_id = var.compartment_id
  name           = "FASTMAIL_TOKEN"
}

data "oci_secrets_secretbundle" "fastmail_token" {
  secret_id      = local.fastmail_token_secret.id
  version_number = local.fastmail_token_secret.current_version_number
}

resource "oci_container_instances_container_instance" "web" {
  availability_domain = "${var.region}-ad-1"
  compartment_id      = var.compartment_id
  containers {
    image_url = "${local.container_image_tag}:latest"
    environment_variables = {
      SENDER_EMAIL   = var.sender_email_address
      FASTMAIL_TOKEN = base64decode(data.oci_secrets_secretbundle.fastmail_token.secret_bundle_content[0].content)
    }
  }
  shape = "CI.Standard.A1.Flex"
  shape_config {
    ocpus         = 1
    memory_in_gbs = 4
  }
  vnics {
    subnet_id = oci_core_subnet.web.id
  }
}

resource "oci_identity_dynamic_group" "container_instances" {
  compartment_id = var.compartment_id
  name           = "all-container-instances"
  description    = "All Container Instances"
  matching_rule  = "ALL {resource.type='computecontainerinstance'}"
}

resource "oci_identity_policy" "allow_registry_pull" {
  compartment_id = var.compartment_id
  name           = "allow-image-pull-from-container-instances"
  description    = "Allow Container Instances to pull images from Container Registry"
  statements = [
    "Allow dynamic-group ${oci_identity_dynamic_group.container_instances.name} to read repos in tenancy"
  ]
}
