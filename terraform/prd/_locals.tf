locals {
  env = "prd"

  app_name      = "hentai-salon-${local.env}"
  host_domain   = "hentai-salon.com"
  server_domain = "server.hentai-salon.com"
  adminer_domain = "adminer.hentai-salon.com"
  client_domain = "hentai-salon.com"
  static_domain = "static.hentai-salon.com"

  ssm_parameter_store = "/hentai-salon/${local.env}"

  bucket_name = "hentai-salon-${local.env}"
}
