variable "ssm_app_env" {
  type = string
}

# server app
variable "ssm_jwt_secret_key" {
  type = string
}
variable "ssm_client_url" {
  type = string
}
variable "ssm_aws_access_key_id" {
  type = string
}
variable "ssm_aws_secret_access_key" {
  type = string
}
variable "ssm_aws_default_region" {
  type = string
}
variable "ssm_aws_bucket_name" {
  type = string
}
variable "ssm_mail_from_address" {
  type = string
}
variable "ssm_db_port" {
  type = string
}
variable "ssm_db_user" {
  type = string
}
variable "ssm_db_name" {
  type = string
}
variable "ssm_db_pass" {
  type = string
}

# client app
variable "ssm_ga_measurement_id" {
  type = string
}
