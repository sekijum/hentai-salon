resource "aws_ses_domain_identity" "this" {
  domain = data.aws_route53_zone.this.name
}

# SESのドメイン認証確認をしようとするとDKIMの認証に最長72時間かかるため無効化する
# Ref https://docs.aws.amazon.com/ja_jp/ses/latest/DeveloperGuide/troubleshoot-dkim.html
# resource "aws_ses_domain_identity_verification" "this" {
#   domain = aws_ses_domain_identity.this.id
# 
#   depends_on = [aws_route53_record.txt_this]
# }

## For DKIM
resource "aws_ses_domain_dkim" "this" {
  domain = data.aws_route53_zone.this.name
}

## For SPF
resource "aws_ses_domain_mail_from" "this" {
  domain           = data.aws_route53_zone.this.name
  mail_from_domain = "mail.${data.aws_route53_zone.this.name}"
}

## For SES
resource "aws_route53_record" "txt_this" {
  zone_id = data.aws_route53_zone.this.zone_id
  name    = "_amazonses.${data.aws_route53_zone.this.name}"
  type    = "TXT"
  ttl     = "600"
  records = [aws_ses_domain_identity.this.verification_token]
}

## For DKIM
resource "aws_route53_record" "cname_dkim_this" {
  count   = 3
  zone_id = data.aws_route53_zone.this.zone_id
  name    = "${element(aws_ses_domain_dkim.this.dkim_tokens, count.index)}._domainkey.${data.aws_route53_zone.this.name}"
  type    = "CNAME"
  ttl     = "600"
  records = ["${element(aws_ses_domain_dkim.this.dkim_tokens, count.index)}.dkim.amazonses.com"]
}

## For SPF
resource "aws_route53_record" "mx_mail_this" {
  zone_id = data.aws_route53_zone.this.zone_id
  name    = aws_ses_domain_mail_from.this.mail_from_domain
  type    = "MX"
  ttl     = "600"
  records = ["10 feedback-smtp.ap-northeast-1.amazonses.com"]
}

resource "aws_route53_record" "txt_mail_this" {
  zone_id = data.aws_route53_zone.this.zone_id
  name    = aws_ses_domain_mail_from.this.mail_from_domain
  type    = "TXT"
  ttl     = "600"
  records = ["v=spf1 include:amazonses.com ~all"]
}

## For DMARC
resource "aws_route53_record" "txt_dmarc_this" {
  zone_id = data.aws_route53_zone.this.zone_id
  name    = "_dmarc.example.com"
  type    = "TXT"
  ttl     = "600"
  records = ["v=DMARC1;p=quarantine;pct=25;rua=mailto:dmarcreports@example.com"]
}
