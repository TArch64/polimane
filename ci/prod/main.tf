terraform {
  required_version = ">= 1.12.0"

  required_providers {
    bitwarden = {
      source  = "maxlaverse/bitwarden"
      version = "0.16.0"
    }

    aws = {
      source  = "hashicorp/aws"
      version = "6.16.0"
    }

    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "5.11.0"
    }

    cockroach = {
      source  = "cockroachdb/cockroach"
      version = "1.15.1"
    }

    tls = {
      source  = "hashicorp/tls"
      version = "4.1.0"
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

variable "bitwarden_token" {
  type      = string
  sensitive = true
  nullable  = false
}

provider "bitwarden" {
  access_token = var.bitwarden_token
  server       = "https://vault.bitwarden.eu"

  experimental {
    embedded_client = true
  }
}

provider "aws" {
  shared_credentials_files = [".aws-credentials"]

  default_tags {
    tags = { app = "polimane" }
  }
}

provider "aws" {
  alias  = "us_east_1"
  region = "us-east-1"
  shared_credentials_files = [".aws-credentials"]

  default_tags {
    tags = { app = "polimane" }
  }
}


provider "cloudflare" {
  api_token = data.bitwarden_secret.cloudflare_api_token.value
}

provider "cockroach" {
  apikey = data.bitwarden_secret.cockroach_api_key.value
}

locals {
  domain   = "polimane.com"
  app_name = "polimane-prod"
  aws_common_tags = aws_servicecatalogappregistry_application.app.application_tag
}

provider "null" {}
provider "tls" {}
provider "external" {}
