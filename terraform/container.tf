resource "oci_artifacts_container_repository" "web" {
  display_name   = "web"
  compartment_id = var.compartment_id
  is_immutable   = false
  is_public      = false
}

output "container_repo_image_tag" {
  value = "ocir.${var.region}.oci.oraclecloud.com/${oci_artifacts_container_repository.web.namespace}/${oci_artifacts_container_repository.web.display_name}"
}
