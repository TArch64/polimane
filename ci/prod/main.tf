terraform {
  required_version = ">= 1.12.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "6.0.0-beta3"
    }

    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "5.5.0"
    }

    tls = {
      source  = "hashicorp/tls"
      version = "4.0.6"
    }

    null = {
      source  = "hashicorp/null"
      version = "3.2.4"
    }

    external = {
      source  = "hashicorp/external"
      version = "2.3.5"
    }
  }

  backend "s3" {
    bucket  = "polimane-prod"
    key     = "ci/terraform"
    region  = "eu-central-1"
    encrypt = true
    shared_credentials_files = [".aws-credentials"]
  }
}

provider "aws" {
  shared_credentials_files = [local.aws_credentials_file]
}

provider "cloudflare" {
  api_token = local.cloudflare_api_token
}

locals {
  domain = "polimane.com"

  aws_region           = "eu-central-1"
  aws_credentials_file = "${path.module}/.aws-credentials"
  aws_common_tags = { app = "polimane" }
}

provider "external" {}
provider "tls" {}
