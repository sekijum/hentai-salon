resource "aws_ecr_repository" "server" {
  name                 = "${local.app_name}-server"
  image_tag_mutability = "MUTABLE"
  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_ecr_repository" "nginx_server" {
  name                 = "${local.app_name}-nginx-server"
  image_tag_mutability = "MUTABLE"
  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_ecr_repository" "client" {
  name                 = "${local.app_name}-client"
  image_tag_mutability = "MUTABLE"
  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_ecr_repository" "nginx_client" {
  name                 = "${local.app_name}-nginx-client"
  image_tag_mutability = "MUTABLE"
  image_scanning_configuration {
    scan_on_push = true
  }
}
