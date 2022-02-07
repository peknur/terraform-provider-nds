terraform {
  required_providers {
    nds = {
      source  = "peknur/nds"
      version = ">= 0.1.0"
    }
  }
}

data "nds_nslookup_mx" "test" {
  name = "traficom.fi"
}

output "example" {
  value = data.nds_nslookup_mx.test.data
}

# terraform output
# example = tolist([
#   {
#     "host" = "mail2.traficom.fi."
#     "priority" = 10
#   },
#   {
#     "host" = "mail1.traficom.fi."
#     "priority" = 10
#   },
# ])
