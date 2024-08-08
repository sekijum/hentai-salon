resource "aws_ssm_parameter" "app_env" {
  name  = local.ssm_parameter_store
  type  = "String"
  value = var.ssm_app_env
}

# server-app
resource "aws_ssm_parameter" "jwt_secret_key" {
  name  = local.ssm_parameter_store
  type  = "String"
  value = var.ssm_jwt_secret_key
}
resource "aws_ssm_parameter" "client_url" {
  name  = local.ssm_parameter_store
  type  = "String"
  value = var.ssm_client_url
}
resource "aws_ssm_parameter" "aws_access_key_id" {
  name  = local.ssm_parameter_store
  type  = "String"
  value = var.ssm_aws_access_key_id
}
resource "aws_ssm_parameter" "aws_secret_access_key" {
  name  = local.ssm_parameter_store
  type  = "String"
  value = var.ssm_aws_secret_access_key
}
resource "aws_ssm_parameter" "aws_default_region" {
  name  = local.ssm_parameter_store
  type  = "String"
  value = var.ssm_aws_default_region
}
resource "aws_ssm_parameter" "aws_bucket_name" {
  name  = local.ssm_parameter_store
  type  = "String"
  value = var.ssm_aws_bucket_name
}
resource "aws_ssm_parameter" "mail_from_address" {
  name  = local.ssm_parameter_store
  type  = "String"
  value = var.ssm_mail_from_address
}
resource "aws_ssm_parameter" "db_port" {
  name  = local.ssm_parameter_store
  type  = "String"
  value = var.ssm_db_port
}
resource "aws_ssm_parameter" "db_user" {
  name  = local.ssm_parameter_store
  type  = "String"
  value = var.ssm_db_user
}
resource "aws_ssm_parameter" "db_name" {
  name  = local.ssm_parameter_store
  type  = "String"
  value = var.ssm_db_name
}
resource "aws_ssm_parameter" "db_pass" {
  name  = local.ssm_parameter_store
  type  = "SecureString"
  value = var.ssm_db_pass
}

# client-app
resource "aws_ssm_parameter" "nuxt_public_api_base_url" {
  name  = local.ssm_parameter_store
  type  = "SecureString"
  value = var.ssm_nuxt_public_api_base_url
}
