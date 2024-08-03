resource "aws_ssm_parameter" "db_name" {
  name  = local.ssm_parameter_store
  type  = "String"
  value = var.ssm_db_name
}

resource "aws_ssm_parameter" "db_user" {
  name  = local.ssm_parameter_store
  type  = "String"
  value = var.ssm_db_user
}

resource "aws_ssm_parameter" "db_pass" {
  name  = local.ssm_parameter_store
  type  = "SecureString"
  value = var.ssm_db_pass
}
