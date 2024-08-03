resource "aws_rds_cluster_instance" "this" {
  count                = 1
  identifier           = "${local.app_name}-database-cluster-instance"
  cluster_identifier   = aws_rds_cluster.this.id
  engine               = aws_rds_cluster.this.engine
  engine_version       = aws_rds_cluster.this.engine_version
  instance_class       = "db.t4g.medium"
  db_subnet_group_name = aws_rds_cluster.this.db_subnet_group_name
}
