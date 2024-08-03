data "aws_acm_certificate" "host_domain" {
  domain      = local.host_domain
  most_recent = true
}
