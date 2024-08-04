resource "aws_service_discovery_private_dns_namespace" "this" {
  name = "internal"
  vpc  = aws_vpc.main.id
}

resource "aws_service_discovery_service" "this" {
  name = "server"
  dns_config {
    namespace_id = aws_service_discovery_private_dns_namespace.this.id
    dns_records {
      ttl  = 10
      type = "A"
    }
  }
  health_check_custom_config {
    failure_threshold = 1
  }
}
