variable "localstack_endpoint" {
  default = "http://localhost:4566"
}

provider "aws" {
  access_key                  = "-"
  secret_key                  = "-"
  region                      = "eu-central-1"
  s3_use_path_style           = true
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true

  endpoints {
    dynamodb = var.localstack_endpoint
  }
}

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "= 5.24.0"
    }
  }
}