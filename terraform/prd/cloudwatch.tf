resource "aws_cloudwatch_log_group" "server" {
  name              = "/${local.app_name}/ecs/server"
  retention_in_days = 7
}
resource "aws_cloudwatch_log_group" "nginx_server" {
  name              = "/${local.app_name}/ecs/nginx-server"
  retention_in_days = 7
}
resource "aws_cloudwatch_log_group" "client" {
  name              = "/${local.app_name}/ecs/client"
  retention_in_days = 7
}
resource "aws_cloudwatch_log_group" "nginx_client" {
  name              = "/${local.app_name}/ecs/nginx-client"
  retention_in_days = 7
}
