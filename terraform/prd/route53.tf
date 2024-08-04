resource "aws_route53_zone" "this" {
  name = local.host_domain
}

resource "aws_route53_record" "server_subdomain" {
  name    = local.server_domain
  type    = "A"
  zone_id = aws_route53_zone.this.zone_id
  alias {
    evaluate_target_health = true
    name                   = aws_lb.this.dns_name
    zone_id                = aws_lb.this.zone_id
  }
}

resource "aws_route53_record" "client_subdomain" {
  name    = local.client_domain
  type    = "A"
  zone_id = aws_route53_zone.this.zone_id
  alias {
    evaluate_target_health = true
    name                   = aws_lb.this.dns_name
    zone_id                = aws_lb.this.zone_id
  }
}

resource "aws_route53_record" "adminer_subdomain" {
  name    = local.adminer_domain
  type    = "A"
  zone_id = aws_route53_zone.this.zone_id
  alias {
    evaluate_target_health = true
    name                   = aws_lb.this.dns_name
    zone_id                = aws_lb.this.zone_id
  }
}
