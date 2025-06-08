terraform {
  required_version = ">= 1.12.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "6.0.0-beta3"
    }
  }

  backend "s3" {
    bucket = "polimane-prod"
    key    = "ci/terraform"
    region = "eu-central-1"
    shared_credentials_files = [".aws-credentials"]
  }
}

provider "aws" {
  shared_credentials_files = ["${path.module}/.aws-credentials"]
}
