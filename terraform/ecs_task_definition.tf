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
        name = "server"
        essential : true
        image = "${aws_ecr_repository.server.repository_url}:latest"
        portMappings = [{ containerPort = 8080 }]
        secrets = [
          {
            name : "APP_SECRET",
            valueFrom : data.aws_ssm_parameter.APP_SECRET.arn
          },
        ]
        logConfiguration = {
          logDriver = "awslogs"
          options = {
            awslogs-region : "ap-northeast-1"
            awslogs-group : aws_cloudwatch_log_group.server.name
            awslogs-stream-prefix : "ecs"
          }
        }
      },
      {
        name = "adminer"
        essential : true
        image = "adminer"
        portMappings = [{ containerPort = 8080 }]
        secrets = [
          {
            name : "DB_HOST",
            valueFrom : aws_rds_cluster.this.endpoint
          },
        ]
      }
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
        name = "client"
        essential : true
        image        = "${aws_ecr_repository.client.repository_url}:latest"
        portMappings = [{ containerPort = 3000 }]
        secrets = [
          {
            name : "APP_SECRET",
            valueFrom : data.aws_ssm_parameter.APP_SECRET.arn
          },
        ]
        logConfiguration = {
          logDriver = "awslogs"
          options = {
            awslogs-region : "ap-northeast-1"
            awslogs-group : aws_cloudwatch_log_group.client.name
            awslogs-stream-prefix : "ecs"
          }
        }
      }
    ]
  )
}
