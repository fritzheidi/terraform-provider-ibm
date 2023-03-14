# HPCS Instance Inputs
variable "hpcs_instance_name" {
  type        = string
  description = "Name of HPCS Instance"
}
variable "hpcs_service_name" {
  default     = "hs-crypto"
  type        = string
  description = "HPCS Service ID"
}
variable "resource_group" {
  default     = "Default"
  type        = string
  description = "Resource group name"
}
variable "location" {
  default     = "us-south"
  type        = string
  description = "Location of HPCS Instance"
}
variable "plan" {
  default     = "standard"
  type        = string
  description = "Plan of HPCS Instance"
}
variable "units" {
  type        = number
  description = "No of crypto units that has to be attached to the instance."
  default     = 2
}
variable "network_access" {
  type        = string
  description = "Network access to your service instance"
  default     = "public-and-private"
}
variable "key_ring_id_list" {
  type = list(string)
  description = "List of HPCS Key Ring IDs"
}
variable "key_name" {
  type        = string
  description = "HPCS Key Name"
}
