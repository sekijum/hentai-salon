resource "aws_acm_certificate" "hentai_salon" {
  domain_name               = local.host_domain
  subject_alternative_names = ["*.hentai-salon.com"]
  validation_method         = "DNS"

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_route53_record" "acm_certificate_validation" {
  for_each = {
    for dvo in aws_acm_certificate.hentai_salon.domain_validation_options : dvo.domain_name => {
      name   = dvo.resource_record_name
      type   = dvo.resource_record_type
      record = dvo.resource_record_value
    }
  }

  name            = each.value.name
  type            = each.value.type
  zone_id         = data.aws_route53_zone.this.zone_id
  records         = [each.value.record]
  ttl             = 300
  allow_overwrite = true
}

resource "aws_acm_certificate_validation" "hentai_salon" {
  certificate_arn         = aws_acm_certificate.hentai_salon.arn
  validation_record_fqdns = [for record in aws_route53_record.acm_certificate_validation : record.fqdn]
}

# us-east-1
resource "aws_acm_certificate" "hentai_salon_us_east_1" {
  provider                  = aws.us_east_1
  domain_name               = local.host_domain
  subject_alternative_names = ["*.hentai-salon.com"]
  validation_method         = "DNS"

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_acm_certificate_validation" "hentai_salon_us_east_1" {
  provider                = aws.us_east_1
  certificate_arn         = aws_acm_certificate.hentai_salon_us_east_1.arn
  validation_record_fqdns = [for record in aws_route53_record.acm_certificate_validation : record.fqdn]
}
