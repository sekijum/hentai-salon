# alb
resource "aws_security_group" "alb" {
  name   = "${local.app_name}-alb"
  vpc_id = aws_vpc.this.id
  egress {
    from_port   = 0
    protocol    = "-1"
    to_port     = 0
    cidr_blocks = ["0.0.0.0/0"]
  }
  tags = {
    Name = "${local.app_name}-alb"
  }
}
resource "aws_security_group_rule" "alb_from_http" {
  from_port         = 80
  protocol          = "tcp"
  to_port           = 80
  type              = "ingress"
  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = aws_security_group.alb.id
}
resource "aws_security_group_rule" "alb_from_https" {
  from_port         = 443
  protocol          = "tcp"
  to_port           = 443
  type              = "ingress"
  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = aws_security_group.alb.id
}

# server
resource "aws_security_group" "server" {
  name   = "${local.app_name}-server"
  vpc_id = aws_vpc.this.id
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  tags = {
    Name = "${local.app_name}-server"
  }
}
resource "aws_security_group_rule" "server_from_this" {
  security_group_id = aws_security_group.server.id
  type              = "ingress"
  from_port         = 0
  to_port           = 0
  protocol          = "-1"
  self              = true
}
resource "aws_security_group_rule" "server_from_alb" {
  security_group_id        = aws_security_group.server.id
  type                     = "ingress"
  from_port                = 0
  to_port                  = 0
  protocol                 = "-1"
  source_security_group_id = aws_security_group.alb.id
}
resource "aws_security_group_rule" "server_from_client" {
  security_group_id        = aws_security_group.server.id
  type                     = "ingress"
  from_port                = 8080
  to_port                  = 8080
  protocol                 = "tcp"
  source_security_group_id = aws_security_group.client.id
}

# adminer
resource "aws_security_group" "adminer" {
  name   = "${local.app_name}-adminer"
  vpc_id = aws_vpc.this.id
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  tags = {
    Name = "${local.app_name}-adminer"
  }
}
resource "aws_security_group_rule" "adminer_from_this" {
  security_group_id = aws_security_group.adminer.id
  type              = "ingress"
  from_port         = 0
  to_port           = 0
  protocol          = "-1"
  self              = true
}
resource "aws_security_group_rule" "adminer_from_alb" {
  security_group_id        = aws_security_group.adminer.id
  type                     = "ingress"
  from_port                = 0
  to_port                  = 0
  protocol                 = "-1"
  source_security_group_id = aws_security_group.alb.id
}

# client
resource "aws_security_group" "client" {
  name   = "${local.app_name}-client"
  vpc_id = aws_vpc.this.id
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  tags = {
    Name = "${local.app_name}-client"
  }
}
resource "aws_security_group_rule" "client_from_this" {
  security_group_id = aws_security_group.client.id
  type              = "ingress"
  from_port         = 0
  to_port           = 0
  protocol          = "-1"
  self              = true
}
resource "aws_security_group_rule" "client_from_alb" {
  security_group_id        = aws_security_group.client.id
  type                     = "ingress"
  from_port                = 0
  to_port                  = 0
  protocol                 = "-1"
  source_security_group_id = aws_security_group.alb.id
}

# rds
resource "aws_security_group" "rds" {
  name   = "${local.app_name}-rds"
  vpc_id = aws_vpc.this.id
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  tags = {
    Name = "${local.app_name}-rds"
  }
}
resource "aws_security_group_rule" "rds_from_server" {
  security_group_id        = aws_security_group.rds.id
  type                     = "ingress"
  from_port                = 3306
  to_port                  = 3306
  protocol                 = "tcp"
  source_security_group_id = aws_security_group.server.id
}
resource "aws_security_group_rule" "rds_from_adminer" {
  security_group_id        = aws_security_group.rds.id
  type                     = "ingress"
  from_port                = 3306
  to_port                  = 3306
  protocol                 = "tcp"
  source_security_group_id = aws_security_group.adminer.id
}
