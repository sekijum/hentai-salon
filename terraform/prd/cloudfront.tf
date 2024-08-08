resource "aws_cloudfront_origin_access_control" "this" {
  name                              = local.static_domain
  origin_access_control_origin_type = "s3"
  signing_behavior                  = "always"
  signing_protocol                  = "sigv4"
}

resource "aws_cloudfront_distribution" "static" {
  aliases = [local.static_domain]
  default_cache_behavior {
    allowed_methods        = ["HEAD", "OPTIONS", "GET", "PUT", "POST", "DELETE", "PATCH"]
    cached_methods         = ["HEAD", "OPTIONS", "GET"]
    compress               = true
    default_ttl            = 86400
    max_ttl                = 31536000
    min_ttl                = 0
    smooth_streaming       = false
    target_origin_id       = aws_s3_bucket.this.id
    trusted_signers        = []
    viewer_protocol_policy = "redirect-to-https"

    cache_policy_id            = data.aws_cloudfront_cache_policy.asset.id
    origin_request_policy_id   = data.aws_cloudfront_origin_request_policy.asset.id
    response_headers_policy_id = aws_cloudfront_response_headers_policy.asset.id
  }

  default_root_object = "index.html"

  enabled         = true
  is_ipv6_enabled = false
  http_version    = "http2"

  origin {
    domain_name              = aws_s3_bucket.this.bucket_domain_name
    origin_id                = aws_s3_bucket.this.id
    origin_access_control_id = aws_cloudfront_origin_access_control.this.id
  }
  price_class = "PriceClass_All"
  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }
  viewer_certificate {
    # cloudfront_default_certificate = false
    # 今後ssl証明書がバージニアリージョンから取得できないので一旦ハードコード
    # たぶんこの辺り
    # `https://limitusus.hatenablog.com/entry/2017/07/12/110343`
    # `https://qiita.com/tos-miyake/items/f0e5f28f2a69e4d39422`
    # acm_certificate_arn            = var.data_aws_acm_certificate_this_arn
    # acm_certificate_arn      = "arn:aws:acm:us-east-1:728047905319:certificate/9fed689c-9c33-476a-8ec0-858b7305e438"
    acm_certificate_arn      = aws_acm_certificate.host_domain_us_east_1.arn
    minimum_protocol_version = "TLSv1.2_2019"
    ssl_support_method       = "sni-only"
  }
  retain_on_delete    = false
  wait_for_deployment = true
}

data "aws_cloudfront_cache_policy" "asset" {
  name = "Managed-Elemental-MediaPackage"
}
data "aws_cloudfront_origin_request_policy" "asset" {
  name = "Managed-CORS-S3Origin"
}
resource "aws_cloudfront_response_headers_policy" "asset" {
  name    = "${local.app_name}-response-headers-policy"
  comment = "Allow from limited origins"

  cors_config {
    access_control_allow_credentials = false

    access_control_allow_headers {
      items = ["*"]
    }

    access_control_allow_methods {
      items = ["GET", "HEAD"]
    }

    access_control_allow_origins {
      items = ["*"]
    }

    origin_override = true
  }
}
