# --------------------------------------------------
# Provison the service using reosource_instance
# --------------------------------------------------
// Supported Regions are us-south and us-east
resource ibm_hpcs hpcs {
  location             = var.location
  name                 = var.hpcs_instance_name
  plan                 = var.plan
  units                = var.units
  signature_threshold  = var.signature_threshold
  revocation_threshold = var.revocation_threshold
  dynamic admins {
    for_each = var.admins
    content {
      name  = admins.value.name
      key   = admins.value.key
      token = admins.value.token
    }
  }
  lifecycle {
    ignore_changes = [signature_threshold,revocation_threshold,admins]
  }
}

# --------------------------------
# Creating Keystores for HPCS Instance
# --------------------------------
resource "ibm_kms_key_rings" "key_ring" {
  instance_id = ibm_hpcs.hpcs.guid
  for_each = toset(var.key_ring_id_list)
  key_ring_id = each.key
}

# --------------------------------
# Creating Keys for HPCS Instance
# --------------------------------

resource "ibm_kms_key" "key" {
  instance_id  = ibm_hpcs.hpcs.guid
  key_name     = var.key_name
  standard_key = false
  force_delete = true
  for_each = ibm_kms_key_rings.key_ring
  key_ring_id = each.value.key_ring_id
}
