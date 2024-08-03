resource "aws_s3_bucket" "this" {
  bucket        = local.bucket_name
  policy        = data.aws_iam_policy_document.s3_bucket_policy.json
  force_destroy = false
  versioning {
    enabled    = true
    mfa_delete = false
  }
}
resource "aws_s3_bucket_website_configuration" "this" {
  bucket = aws_s3_bucket.this.id
  redirect_all_requests_to {
    host_name = local.static_domain
  }
}
resource "aws_s3_bucket_cors_configuration" "this" {
  bucket = aws_s3_bucket.this.id
  cors_rule {
    allowed_headers = ["*"]
    allowed_methods = ["GET", "PUT", "POST"]
    allowed_origins = ["*"]
    expose_headers  = []
    max_age_seconds = 3000
  }
}

data "aws_iam_policy_document" "s3_bucket_policy" {
  statement {
    principals {
      type        = "Service"
      identifiers = ["cloudfront.amazonaws.com"]
    }
    actions = ["s3:GetObject"]
    resources = [
      "arn:aws:s3:::${local.bucket_name}",
      "arn:aws:s3:::${local.bucket_name}/*"
    ]
    condition {
      test     = "ForAnyValue:StringEquals"
      variable = "aws:SourceArn"
      values   = [aws_cloudfront_distribution.static.arn]
    }
  }
}
