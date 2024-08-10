resource "aws_ecs_task_definition" "server" {
  family                   = "${local.app_name}-server"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = 256
  memory                   = 512
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn
  task_role_arn            = aws_iam_role.ecs_task_execution_role.arn
  container_definitions = jsonencode(
    [
      {
        name = "app"
        essential : true
        image        = "${aws_ecr_repository.server_app.repository_url}:latest"
        portMappings = [{ containerPort = 8080 }]
        secrets = [
          { name : "APP_ENV", valueFrom : aws_ssm_parameter.app_env.arn },
          { name : "JWT_SECRET_KEY", valueFrom : aws_ssm_parameter.server_jwt_secret_key.arn },
          { name : "CLIENT_URL", valueFrom : aws_ssm_parameter.server_client_url.arn },
          { name : "AWS_ACCESS_KEY_ID", valueFrom : aws_ssm_parameter.server_aws_access_key_id.arn },
          { name : "AWS_SECRET_ACCESS_KEY", valueFrom : aws_ssm_parameter.server_aws_secret_access_key.arn },
          { name : "AWS_DEFAULT_REGION", valueFrom : aws_ssm_parameter.server_aws_default_region.arn },
          { name : "AWS_BUCKET_NAME", valueFrom : aws_ssm_parameter.server_aws_bucket_name.arn },
          { name : "DB_PORT", valueFrom : aws_ssm_parameter.server_db_port.arn },
          { name : "DB_HOST", valueFrom : aws_ssm_parameter.server_db_host.arn },
          { name : "DB_USER", valueFrom : aws_ssm_parameter.server_db_user.arn },
          { name : "DB_NAME", valueFrom : aws_ssm_parameter.server_db_name.arn },
          { name : "DB_PASS", valueFrom : aws_ssm_parameter.server_db_pass.arn },
        ]
        logConfiguration = {
          logDriver = "awslogs"
          options = {
            awslogs-region : "ap-northeast-1"
            awslogs-group : aws_cloudwatch_log_group.server_app.name
            awslogs-stream-prefix : "ecs"
          }
        }
      },
      {
        name = "proxy"
        essential : true
        image        = "${aws_ecr_repository.server_proxy.repository_url}:latest"
        portMappings = [{ containerPort = 80 }]
        dependsOn = [
          {
            containerName = "app"
            condition     = "START"
          },
        ]
        logConfiguration = {
          logDriver = "awslogs"
          options = {
            awslogs-region : "ap-northeast-1"
            awslogs-group : aws_cloudwatch_log_group.server_proxy.name
            awslogs-stream-prefix : "ecs"
          }
        }
      },
    ]
  )
}

resource "aws_ecs_task_definition" "client" {
  family                   = "${local.app_name}-client"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = 256
  memory                   = 512
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn
  task_role_arn            = aws_iam_role.ecs_task_execution_role.arn
  container_definitions = jsonencode(
    [
      {
        name = "app"
        essential : true
        image        = "${aws_ecr_repository.client_app.repository_url}:latest"
        portMappings = [{ containerPort = 3000 }]
        logConfiguration = {
          logDriver = "awslogs"
          options = {
            awslogs-region : "ap-northeast-1"
            awslogs-group : aws_cloudwatch_log_group.client_app.name
            awslogs-stream-prefix : "ecs"
          }
        }
      },
      {
        name = "proxy"
        essential : true
        image        = "${aws_ecr_repository.client_proxy.repository_url}:latest"
        portMappings = [{ containerPort = 80 }]
        dependsOn = [
          {
            containerName = "app"
            condition     = "START"
          },
        ]
        logConfiguration = {
          logDriver = "awslogs"
          options = {
            awslogs-region : "ap-northeast-1"
            awslogs-group : aws_cloudwatch_log_group.client_proxy.name
            awslogs-stream-prefix : "ecs"
          }
        }
      },
    ]
  )
}

resource "aws_ecs_task_definition" "adminer" {
  family                   = "${local.app_name}-adminer"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = 256
  memory                   = 512
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn
  task_role_arn            = aws_iam_role.ecs_task_execution_role.arn
  container_definitions = jsonencode(
    [
      {
        name = "adminer"
        essential : true
        image        = "adminer"
        portMappings = [{ containerPort = 8080 }]
        secrets = [
          {
            name : "ADMINER_DEFAULT_SERVER",
            valueFrom : aws_ssm_parameter.server_db_host.arn
          },
        ]
      },
    ]
  )
}
