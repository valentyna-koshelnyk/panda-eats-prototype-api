provider "aws" {
  alias                     = "localstack"
  region                    = "eu-central-1"
  access_key                = "AKIAIOSFODNN7EXAMPLE"
  secret_key                = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true
  endpoints {
    dynamodb = "http://localhost:4566"
  }
}

resource "aws_dynamodb_table" "cart_table" {
  provider      = aws.localstack
  name          = "cart"
  billing_mode  = "PAY_PER_REQUEST"
  hash_key      = "user_id"
  range_key = "product_id"

  attribute {
    name = "user_id"
    type = "S"
  }

  attribute {
    name = "product_id"
    type = "S"
  }
}

resource "aws_dynamodb_table" "order_table" {
  provider      = aws.localstack
  name          = "order"
  billing_mode  = "PAY_PER_REQUEST"
  hash_key      = "user_id"
  range_key = "order_id"

  attribute {
    name = "user_id"
    type = "S"
  }

  attribute {
    name = "order_id"
    type = "S"
  }
}