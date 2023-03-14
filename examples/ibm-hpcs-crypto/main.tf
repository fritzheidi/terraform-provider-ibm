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
   units: var.units,
   // hybrid: "true",
   // ekm_id: "6b4ca546-0117-4d3e-9786-a0323c1d9b11"
 }
}
