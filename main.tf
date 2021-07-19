variable "HEROKU_API_KEY" {
  type = string
}

variable "HEROKU_EMAIL" {
  type = string
}

provider "heroku" {
  email   = var.HEROKU_EMAIL
  api_key = var.HEROKU_API_KEY
}

terraform {
  required_providers {
    heroku = {
      source  = "heroku/heroku"
      version = "~> 4.0"
    }
  }
}

resource "heroku_app" "default" {
  name   = "vacation-labs-go-terraform"
  region = "us"
}