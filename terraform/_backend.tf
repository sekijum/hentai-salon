terraform {
  required_version = "1.5.3"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.13.1"
    }
  }

  backend "s3" {
    bucket                  = "hentai-salon-terraform-tfstate"
    key                     = "state/terraform.tfstate"
    region                  = "ap-northeast-1"
    encrypt                 = true
    shared_credentials_file = "~/.aws/credentials"
    profile                 = "hentai-salon"
  }
}
