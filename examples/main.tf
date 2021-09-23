terraform {
  required_providers {
    hashicups = {
      version = "~> 0.3.2"
      source  = "hashicorp.com/edu/hashicups"
    }
  }
  required_version = "~> 1.0.3"
}

provider "hashicups" {
  username = "education"
  password = "test123"
  host     = "http://localhost:19090"
}

data "hashicups_coffees" "all" {}

resource "hashicups_order" "edu" {
  items = [{
    coffee = {
      id = 3
    }
    quantity = 2
    }, {
    coffee = {
      id = 1
    }
    quantity = 2
    }
  ]
}

output "edu_order" {
  value = hashicups_order.edu
}

output "all_coffees" {
  value = data.hashicups_coffees.all.coffees
}
