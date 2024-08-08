resource "aws_db_subnet_group" "this" {
  name = "${local.app_name}-database-subnet-group"
  subnet_ids = [
    aws_subnet.private_1a.id,
    aws_subnet.private_1c.id,
    aws_subnet.private_1d.id,
  ]
}

resource "aws_rds_cluster" "this" {
  cluster_identifier   = "${local.app_name}-database-cluster"
  db_subnet_group_name = aws_db_subnet_group.this.name
  vpc_security_group_ids = [
    aws_security_group.rds.id,
  ]
  engine                    = "aurora-mysql"
  engine_version            = "8.0.mysql_aurora.3.07.1"
  port                      = "3306"
  database_name             = aws_ssm_parameter.server_db_name.value
  master_username           = aws_ssm_parameter.server_db_user.value
  master_password           = aws_ssm_parameter.server_db_pass.value
  backup_retention_period   = 7
  preferred_backup_window   = "04:00-05:00"
  deletion_protection       = true
  skip_final_snapshot       = true
  final_snapshot_identifier = "${local.app_name}-final-snapshot"
  apply_immediately         = true

  db_cluster_parameter_group_name = aws_rds_cluster_parameter_group.this.name
}
