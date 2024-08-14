# ecs server task
resource "aws_appautoscaling_target" "server" {
  service_namespace  = "ecs"
  scalable_dimension = "ecs:service:DesiredCount"
  resource_id        = "service/${aws_ecs_cluster.this.name}/${aws_ecs_service.server.name}"
  min_capacity       = 1
  max_capacity       = 3
}
resource "aws_appautoscaling_scheduled_action" "ecs_server_task_scale_out" {
  name               = "${local.app_name}-server-task-scale-out"
  service_namespace  = aws_appautoscaling_target.server.service_namespace
  resource_id        = aws_appautoscaling_target.server.resource_id
  scalable_dimension = aws_appautoscaling_target.server.scalable_dimension
  timezone           = "Asia/Tokyo"
  schedule           = "cron(15 10 ? * 6L 2002-2005)"
  scalable_target_action {
    min_capacity = 1
    max_capacity = 3
  }
}
resource "aws_appautoscaling_scheduled_action" "ecs_server_task_scale_in" {
  name               = "${local.app_name}-server-task-scale-in"
  service_namespace  = aws_appautoscaling_target.server.service_namespace
  resource_id        = aws_appautoscaling_target.server.resource_id
  scalable_dimension = aws_appautoscaling_target.server.scalable_dimension
  timezone           = "Asia/Tokyo"
  schedule           = "cron(15 10 ? * 6L 2002-2005)"
  scalable_target_action {
    min_capacity = 0
    max_capacity = 0
  }
}
# cpu使用率を60%に維持
resource "aws_appautoscaling_policy" "ecs_server_task_cpu_scale" {
  name               = "${local.app_name}-server-task-cpu-scale"
  policy_type        = "TargetTrackingScaling"
  resource_id        = aws_appautoscaling_target.server.resource_id
  scalable_dimension = aws_appautoscaling_target.server.scalable_dimension
  service_namespace  = aws_appautoscaling_target.server.service_namespace
  target_tracking_scaling_policy_configuration {
    target_value       = 60
    scale_in_cooldown  = 60
    scale_out_cooldown = 30

    predefined_metric_specification {
      predefined_metric_type = "ECSServiceAverageCPUUtilization"
    }
  }
}
# メモリ使用率を60%に維持
resource "aws_appautoscaling_policy" "ecs_server_task_memory_scale" {
  name               = "${local.app_name}-server-task-memory-scale"
  policy_type        = "TargetTrackingScaling"
  resource_id        = aws_appautoscaling_target.server.resource_id
  scalable_dimension = aws_appautoscaling_target.server.scalable_dimension
  service_namespace  = aws_appautoscaling_target.server.service_namespace
  target_tracking_scaling_policy_configuration {
    target_value       = 60
    scale_in_cooldown  = 60
    scale_out_cooldown = 30

    predefined_metric_specification {
      predefined_metric_type = "ECSServiceAverageMemoryUtilization"
    }
  }
}

# ecs client task
resource "aws_appautoscaling_target" "client" {
  service_namespace  = "ecs"
  scalable_dimension = "ecs:service:DesiredCount"
  resource_id        = "service/${aws_ecs_cluster.this.name}/${aws_ecs_service.client.name}"
  min_capacity       = 1
  max_capacity       = 3
}
resource "aws_appautoscaling_scheduled_action" "ecs_client_task_scale_out" {
  name               = "${local.app_name}-client-task-scale-out"
  service_namespace  = aws_appautoscaling_target.client.service_namespace
  resource_id        = aws_appautoscaling_target.client.resource_id
  scalable_dimension = aws_appautoscaling_target.client.scalable_dimension
  timezone           = "Asia/Tokyo"
  schedule           = "cron(15 10 ? * 6L 2002-2005)"
  scalable_target_action {
    min_capacity = 1
    max_capacity = 3
  }
}
resource "aws_appautoscaling_scheduled_action" "ecs_client_task_scale_in" {
  name               = "${local.app_name}-client-task-scale-in"
  service_namespace  = aws_appautoscaling_target.client.service_namespace
  resource_id        = aws_appautoscaling_target.client.resource_id
  scalable_dimension = aws_appautoscaling_target.client.scalable_dimension
  timezone           = "Asia/Tokyo"
  schedule           = "cron(15 10 ? * 6L 2002-2005)"
  scalable_target_action {
    min_capacity = 0
    max_capacity = 0
  }
}
# cpu使用率を60%に維持
resource "aws_appautoscaling_policy" "ecs_client_task_cpu_scale" {
  name               = "${local.app_name}-server-task-cpu-scale"
  policy_type        = "TargetTrackingScaling"
  resource_id        = aws_appautoscaling_target.client.resource_id
  scalable_dimension = aws_appautoscaling_target.client.scalable_dimension
  service_namespace  = aws_appautoscaling_target.client.service_namespace
  target_tracking_scaling_policy_configuration {
    target_value       = 60
    scale_in_cooldown  = 60
    scale_out_cooldown = 30

    predefined_metric_specification {
      predefined_metric_type = "ECSServiceAverageCPUUtilization"
    }
  }
}
# メモリ使用率を60%に維持
resource "aws_appautoscaling_policy" "ecs_client_task_memory_scale" {
  name               = "${local.app_name}-client-task-memory-scale"
  policy_type        = "TargetTrackingScaling"
  resource_id        = aws_appautoscaling_target.client.resource_id
  scalable_dimension = aws_appautoscaling_target.client.scalable_dimension
  service_namespace  = aws_appautoscaling_target.client.service_namespace
  target_tracking_scaling_policy_configuration {
    target_value       = 60
    scale_in_cooldown  = 60
    scale_out_cooldown = 30

    predefined_metric_specification {
      predefined_metric_type = "ECSServiceAverageMemoryUtilization"
    }
  }
}
