# --------------------------------------------------
# Provison the service using ibm_resource_instance
# --------------------------------------------------
data "ibm_resource_group" "group" {
  name = "Default"
}

resource ibm_resource_instance hpcs {
  name              = var.hpcs_instance_name
  service           = var.hpcs_service_name
  plan              = var.plan
  location          = var.location
  resource_group_id = data.ibm_resource_group.group.id

 parameters = {
   units:           var.units,
   allowed_network: var.network_access
   // hybrid:       "true",
   // ekm_id:       "6b4ca546-0117-4d3e-9786-a0323c1d9b11"
 }
}

resource "ibm_kms_key_rings" "key_ring" {
  instance_id = ibm_resource_instance.hpcs.guid
  for_each = toset(var.key_ring_id_list)
  key_ring_id = each.key
}

resource "ibm_kms_key" "key" {
  instance_id  = ibm_resource_instance.hpcs.guid
  key_name     = var.key_name
  standard_key = false
  force_delete = true
  for_each = ibm_kms_key_rings.key_ring
  key_ring_id = each.value.key_ring_id
}

resource "ibm_iam_authorization_policy" "policy" {
  roles      = [
    "Reader",
  ]

  resource_attributes {
    name     = "accountId"
    value    = var.source_account_id
  }
  resource_attributes {
    name     = "serviceName"
    value    = var.source_service
  }

  subject_attributes {
    name  = "accountId"
    value = var.target_account_id
  }
  subject_attributes {
    name  = "serviceName"
    value = var.target_service
  }

  subject_attributes { 
    name  =  "keyRing"
    value =  var.key_ring_id
  }
}
