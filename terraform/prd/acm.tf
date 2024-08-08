resource "aws_acm_certificate" "host_domain" {
  domain_name       = aws_route53_registered_domain.this.domain_name
  validation_method = "DNS"
}
