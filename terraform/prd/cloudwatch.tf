resource "aws_cloudwatch_log_group" "server_app" {
  name              = "/${local.app_name}/ecs/server/app"
  retention_in_days = 7
}
resource "aws_cloudwatch_log_group" "server_proxy" {
  name              = "/${local.app_name}/ecs/server/proxy"
  retention_in_days = 7
}
resource "aws_cloudwatch_log_group" "client_app" {
  name              = "/${local.app_name}/ecs/client/app"
  retention_in_days = 7
}
resource "aws_cloudwatch_log_group" "client_proxy" {
  name              = "/${local.app_name}/ecs/client/proxy"
  retention_in_days = 7
}
