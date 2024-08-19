resource "oci_core_vcn" "internal" {
  dns_label      = "internal"
  cidr_block     = "172.16.0.0/20"
  compartment_id = var.compartment_id
  display_name   = "Private"
}

resource "oci_core_subnet" "web" {
  compartment_id = var.compartment_id
  vcn_id         = oci_core_vcn.internal.id
  cidr_block     = "172.16.1.0/24"
}

resource "oci_core_internet_gateway" "web" {
  compartment_id = var.compartment_id
  vcn_id         = oci_core_vcn.internal.id
}

resource "oci_core_route_table" "web" {
  compartment_id = var.compartment_id
  vcn_id         = oci_core_vcn.internal.id
  route_rules {
    network_entity_id = oci_core_internet_gateway.web.id
    destination_type  = "CIDR_BLOCK"
    destination       = "0.0.0.0/0"
  }
}

resource "oci_core_route_table_attachment" "web" {
  route_table_id = oci_core_route_table.web.id
  subnet_id      = oci_core_subnet.web.id
}
