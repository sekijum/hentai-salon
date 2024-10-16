resource "aws_ecs_service" "server" {
  name                               = "server"
  cluster                            = aws_ecs_cluster.this.id
  platform_version                   = "LATEST"
  task_definition                    = aws_ecs_task_definition.server.arn
  desired_count                      = 1
  deployment_minimum_healthy_percent = 50
  deployment_maximum_percent         = 200
  propagate_tags                     = "SERVICE"
  launch_type                        = "FARGATE"
  health_check_grace_period_seconds  = 60
  enable_execute_command             = true
  deployment_circuit_breaker {
    enable   = true
    rollback = true
  }
  network_configuration {
    assign_public_ip = true
    subnets = [
      aws_subnet.public_1a.id,
      aws_subnet.public_1c.id,
      aws_subnet.public_1d.id,
    ]
    security_groups = [aws_security_group.server.id]
  }
  load_balancer {
    target_group_arn = aws_lb_target_group.server.arn
    container_name   = "proxy"
    container_port   = 80
  }
}

resource "aws_ecs_service" "client" {
  name                               = "client"
  cluster                            = aws_ecs_cluster.this.id
  platform_version                   = "LATEST"
  task_definition                    = aws_ecs_task_definition.client.arn
  desired_count                      = 1
  deployment_minimum_healthy_percent = 50
  deployment_maximum_percent         = 200
  propagate_tags                     = "SERVICE"
  launch_type                        = "FARGATE"
  health_check_grace_period_seconds  = 60
  enable_execute_command             = true
  deployment_circuit_breaker {
    enable   = true
    rollback = true
  }
  network_configuration {
    assign_public_ip = true
    subnets = [
      aws_subnet.public_1a.id,
      aws_subnet.public_1c.id,
      aws_subnet.public_1d.id,
    ]
    security_groups = [aws_security_group.client.id]
  }
  load_balancer {
    target_group_arn = aws_lb_target_group.client.arn
    container_name   = "proxy"
    container_port   = 80
  }
}

resource "aws_ecs_service" "adminer" {
  name                               = "adminer"
  cluster                            = aws_ecs_cluster.this.id
  platform_version                   = "LATEST"
  task_definition                    = aws_ecs_task_definition.adminer.arn
  desired_count                      = 0
  deployment_minimum_healthy_percent = 50
  deployment_maximum_percent         = 200
  propagate_tags                     = "SERVICE"
  launch_type                        = "FARGATE"
  health_check_grace_period_seconds  = 60
  enable_execute_command             = true
  deployment_circuit_breaker {
    enable   = true
    rollback = true
  }
  network_configuration {
    assign_public_ip = true
    subnets = [
      aws_subnet.public_1a.id,
      aws_subnet.public_1c.id,
      aws_subnet.public_1d.id,
    ]
    security_groups = [aws_security_group.adminer.id]
  }
  load_balancer {
    target_group_arn = aws_lb_target_group.adminer.arn
    container_name   = "adminer"
    container_port   = 8080
  }
}
