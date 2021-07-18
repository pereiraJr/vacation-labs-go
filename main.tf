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
  region = "eu"
}