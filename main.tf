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
  backend "remote" {
    organization = "vacation-labs"

    workspaces {
      name = "vacation-labs"
    }
  }
}

resource "heroku_app" "default" {
  name   = "vacation-labs-go-terraform"
  region = "us"

  buildpacks = [
    "heroku/go"
  ]
}

resource "heroku_addon" "database" {
  app  = heroku_app.default.name
  plan = "heroku-postgresql:hobby-dev"
}