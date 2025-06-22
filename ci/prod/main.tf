terraform {
  required_version = ">= 1.12.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "6.0.0"
    }

    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "5.6.0"
    }

    tls = {
      source  = "hashicorp/tls"
      version = "4.1.0"
    }

    null = {
      source  = "hashicorp/null"
      version = "3.2.4"
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

  default_tags {
    tags = { app = "polimane" }
  }
}

variable "cloudflare_api_token" {
  type      = string
  nullable  = false
  sensitive = true
}

provider "cloudflare" {
  api_token = var.cloudflare_api_token
}

locals {
  domain = "polimane.com"
  app_name = "polimane-prod"

  aws_region           = "eu-central-1"
  aws_credentials_file = "${path.module}/.aws-credentials"
  aws_common_tags = aws_servicecatalogappregistry_application.app.application_tag
}

provider "null" {}
provider "tls" {}
