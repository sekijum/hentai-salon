data "aws_ssm_parameter" "DB_NAME" {
  name = "${local.ssm_parameter_store}/DB_NAME"
}
data "aws_ssm_parameter" "DB_USER" {
  name = "${local.ssm_parameter_store}/DB_USER"
}
data "aws_ssm_parameter" "DB_PASS" {
  name = "${local.ssm_parameter_store}/DB_PASS"
}
