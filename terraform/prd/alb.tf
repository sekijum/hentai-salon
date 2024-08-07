resource "aws_lb" "this" {
  name               = "${local.app_name}-alb"
  load_balancer_type = "application"
  security_groups = [
    aws_security_group.alb.id
  ]
  subnets = [
    aws_subnet.public_1a.id,
    aws_subnet.public_1c.id,
    aws_subnet.public_1d.id,
  ]
}

resource "aws_lb_listener" "http" {
  load_balancer_arn = aws_lb.this.arn
  port              = "80"
  protocol          = "HTTP"
  default_action {
    type = "redirect"
    redirect {
      port        = "443"
      protocol    = "HTTPS"
      status_code = "HTTP_301"
    }
  }
}

resource "aws_lb_listener" "https" {
  load_balancer_arn = aws_lb.this.arn
  port              = 443
  protocol          = "HTTPS"
  certificate_arn   = aws_acm_certificate.hentai_salon.arn
  default_action {
    type = "fixed-response"
    fixed_response {
      content_type = "text/plain"
      message_body = "Fixed response content"
      status_code  = "200"
    }
  }
}

resource "aws_lb_target_group" "server" {
  name                 = "${local.app_name}-server"
  vpc_id               = aws_vpc.this.id
  target_type          = "ip"
  port                 = 80
  protocol             = "HTTP"
  deregistration_delay = 60
  health_check {
    path = "/health-check"
  }
}

resource "aws_lb_target_group" "client" {
  name                 = "${local.app_name}-client"
  vpc_id               = aws_vpc.this.id
  target_type          = "ip"
  port                 = 80
  protocol             = "HTTP"
  deregistration_delay = 60
  health_check {
    path = "/"
  }
}

resource "aws_lb_target_group" "adminer" {
  name                 = "${local.app_name}-adminer"
  vpc_id               = aws_vpc.this.id
  target_type          = "ip"
  port                 = 80
  protocol             = "HTTP"
  deregistration_delay = 60
  health_check {
    path = "/"
  }
}

resource "aws_lb_listener_rule" "forward_client" {
  listener_arn = aws_lb_listener.https.arn
  priority     = 10
  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.client.arn
  }
  condition {
    host_header {
      values = [
        local.client_domain
      ]
    }
  }
}

resource "aws_lb_listener_rule" "forward_server" {
  listener_arn = aws_lb_listener.https.arn
  priority     = 20
  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.server.arn
  }
  condition {
    host_header {
      values = [
        local.server_domain
      ]
    }
  }
}

resource "aws_lb_listener_rule" "forward_adminer" {
  listener_arn = aws_lb_listener.https.arn
  priority     = 50
  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.adminer.arn
  }
  condition {
    host_header {
      values = [
        local.adminer_domain
      ]
    }
  }
}

resource "aws_lb_listener_rule" "maintenance" {
  listener_arn = aws_lb_listener.https.arn
  priority     = 80
  action {
    type = "fixed-response"
    fixed_response {
      content_type = "text/plain"
      message_body = "503 サービスは一時的に利用できません"
      status_code  = "503"
    }
  }
  condition {
    path_pattern {
      values = ["*"]
    }
  }
}

resource "aws_lb_listener_rule" "not_found" {
  listener_arn = aws_lb_listener.https.arn
  priority     = 90
  action {
    type = "fixed-response"
    fixed_response {
      content_type = "text/plain"
      message_body = "404 ページが見つかりません"
      status_code  = "404"
    }
  }
  condition {
    path_pattern {
      values = ["*"]
    }
  }
}

resource "aws_lb_listener_rule" "server_error" {
  listener_arn = aws_lb_listener.https.arn
  priority     = 100
  action {
    type = "fixed-response"
    fixed_response {
      content_type = "text/plain"
      message_body = "500 サーバー内部エラーが発生しました"
      status_code  = "500"
    }
  }
  condition {
    path_pattern {
      values = ["*"]
    }
  }
}
