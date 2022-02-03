terraform {
  required_version = ">= 0.13.0"
  required_providers {
    hashicups = {
      version = "~> 1.0.0"
      source  = "just1not2/haproxy"
    }
  }
}

provider "haproxy" {
  url      = "http://haproxy.example.com:5555/v2"
  username = "admin"
  password = "adminpwd"
}
