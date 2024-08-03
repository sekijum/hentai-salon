locals {
  aws_profile = "hentai-salon"
}

locals {
  app_name      = "hentai-salon"
  
  host_domain   = "hentai-salon.com"
  server_domain = "server.hentai-salon.com"
  adminer_domain = "adminer.hentai-salon.com"
  client_domain = "hentai-salon.com"
  static_domain = "static.hentai-salon.com"

  ssm_parameter_store = "hentai-salon"

  bucket_name = "hentai-salon"
}
