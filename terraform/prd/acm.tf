resource "aws_acm_certificate" "host_domain" {
  domain_name       = local.host_domain
  validation_method = "DNS"
}
