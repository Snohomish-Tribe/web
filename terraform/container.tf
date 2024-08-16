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
  container_image_tag = "ocir.${var.region}.oci.oraclecloud.com/${oci_artifacts_container_repository.web.namespace}/${oci_artifacts_container_repository.web.display_name}"
}

resource "oci_container_instances_container_instance" "web" {
  availability_domain = "${var.region}-ad-1"
  compartment_id      = var.compartment_id
  containers {
    image_url = "${container_image_tag}:latest"
  }
  shape = "CI.Standard.A1.Flex"
  shape_config {
    ocpus         = 1
    memory_in_gbs = 5
  }
  vnics {
    subnet_id = oci_core_subnet.web.id
  }
}

resource "oci_core_subnet" "web" {
  compartment_id = var.compartment_id
  vcn_id         = oci_core_vcn.internal.id
  cidr_block     = "172.16.1.0/24"
}
