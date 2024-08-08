provider "aws" {
  shared_credentials_files = ["$HOME/.aws/credentials"]
  profile                  = "hentai-salon"
  region                   = "ap-northeast-1"
  default_tags {
    tags = {
      application = local.app_name
    }
  }
}

provider "aws" {
  alias                    = "us_east_1"
  shared_credentials_files = ["$HOME/.aws/credentials"]
  profile                  = "hentai-salon"
  region                   = "us-east-1"
  default_tags {
    tags = {
      application = local.app_name
    }
  }
}
