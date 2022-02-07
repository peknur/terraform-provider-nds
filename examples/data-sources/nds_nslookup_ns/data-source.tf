terraform {
  required_providers {
    nds = {
      source  = "peknur/nds"
      version = ">= 0.1.0"
    }
  }
}

data "nds_nslookup_ns" "example" {
  name = "example.com"
}

output "example" {
  value = data.nds_nslookup_ns.example.data
}

# terraform output 
# example = tolist([
#   "a.iana-servers.net.",
#   "b.iana-servers.net.",
# ])
