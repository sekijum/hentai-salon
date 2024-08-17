resource "aws_ssm_parameter" "app_env" {
  name  = "app_env"
  type  = "String"
  value = var.ssm_app_env
}

# server-app
resource "aws_ssm_parameter" "server_jwt_secret_key" {
  name  = "${local.ssm_parameter_store}/server/jwt_secret_key"
  type  = "String"
  value = var.ssm_jwt_secret_key
}
resource "aws_ssm_parameter" "server_client_url" {
  name  = "${local.ssm_parameter_store}/server/client_url"
  type  = "String"
  value = var.ssm_client_url
}
resource "aws_ssm_parameter" "server_aws_access_key_id" {
  name  = "${local.ssm_parameter_store}/server/aws_access_key_id"
  type  = "String"
  value = var.ssm_aws_access_key_id
}
resource "aws_ssm_parameter" "server_aws_secret_access_key" {
  name  = "${local.ssm_parameter_store}/server/aws_secret_access_key"
  type  = "String"
  value = var.ssm_aws_secret_access_key
}
resource "aws_ssm_parameter" "server_aws_default_region" {
  name  = "${local.ssm_parameter_store}/server/aws_default_region"
  type  = "String"
  value = var.ssm_aws_default_region
}
resource "aws_ssm_parameter" "server_aws_bucket_name" {
  name  = "${local.ssm_parameter_store}/server/aws_bucket_name"
  type  = "String"
  value = var.ssm_aws_bucket_name
}
resource "aws_ssm_parameter" "server_mail_from_address" {
  name  = "${local.ssm_parameter_store}/server/mail_from_address"
  type  = "String"
  value = var.ssm_mail_from_address
}
resource "aws_ssm_parameter" "server_db_port" {
  name  = "${local.ssm_parameter_store}/server/db_port"
  type  = "String"
  value = var.ssm_db_port
}
resource "aws_ssm_parameter" "server_db_host" {
  name  = "${local.ssm_parameter_store}/server/db_host"
  type  = "String"
  value = aws_rds_cluster.this.endpoint
}
resource "aws_ssm_parameter" "server_db_user" {
  name  = "${local.ssm_parameter_store}/server/db_user"
  type  = "String"
  value = var.ssm_db_user
}
resource "aws_ssm_parameter" "server_db_name" {
  name  = "${local.ssm_parameter_store}/server/db_name"
  type  = "String"
  value = var.ssm_db_name
}
resource "aws_ssm_parameter" "server_db_pass" {
  name  = "${local.ssm_parameter_store}/server/db_pass"
  type  = "SecureString"
  value = var.ssm_db_pass
}
resource "aws_ssm_parameter" "client_ga_measurement_id" {
  name  = "${local.ssm_parameter_store}/client/ssm_ga_measurement_id"
  type  = "String"
  value = var.ssm_ga_measurement_id
}
