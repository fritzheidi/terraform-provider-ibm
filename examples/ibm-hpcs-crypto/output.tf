output "InstanceGUID" {
  value = ibm_resource_instance.hpcs.guid
}
output "keyIDs" {
  value = {
    for key, value in ibm_kms_key.key : key => value.id
  }
}
