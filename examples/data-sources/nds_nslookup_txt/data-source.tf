terraform {
  required_providers {
    nds = {
      source  = "peknur/nds"
      version = ">= 0.1.0"
    }
  }
}

data "nds_nslookup_txt" "example" {
  name = "example.debsu.fi"
}

output "example" {
  value = data.nds_nslookup_txt.example.data
}

# terraform output 
# example = tolist([
#   "Lorem Ipsum has been the industry's standard dummy text ever since the 1500s.",
# ])
