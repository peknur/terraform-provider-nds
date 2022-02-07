terraform {
  required_providers {
    nds = {
      source  = "peknur/nds"
      version = ">= 0.1.0"
    }
  }
}

data "nds_nslookup_srv" "example" {
  name    = "example.debsu.fi"
  proto   = "tcp"
  service = "xmpp-client"
}

output "example" {
  value = data.nds_nslookup_srv.example.data
}

# terraform output
# example = tolist([
#   {
#     "port" = 5222
#     "priority" = 5
#     "target" = "example.com."
#     "weight" = 0
#   },
#   {
#     "port" = 5222
#     "priority" = 10
#     "target" = "example.com."
#     "weight" = 5
#   },
# ])
