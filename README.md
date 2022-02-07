#  Net Data Source Terraform Provider
[![Tests](https://github.com/peknur/terraform-provider-nds/actions/workflows/test.yml/badge.svg)](https://github.com/peknur/terraform-provider-nds/actions/workflows/test.yml)
[![golangci-lint](https://github.com/peknur/terraform-provider-nds/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/peknur/terraform-provider-nds/actions/workflows/golangci-lint.yml)
[![release](https://github.com/peknur/terraform-provider-nds/actions/workflows/release.yml/badge.svg)](https://github.com/peknur/terraform-provider-nds/actions/workflows/release.yml)

Net Data Source Terraform Provider enables users to query network data sources.

## Documentation
Full documentation is available on the Terraform registry website:  
https://registry.terraform.io/providers/peknur/nds/latest/docs

## Example usage
```terraform
terraform {
  required_providers {
    nds = {
      source  = "peknur/nds"
      version = ">= 0.1.0"
    }
  }
}

# Use custom resolver to get public IP address
data "nds_nslookup_ip" "my_ip" {
  name = "myip.opendns.com"
  resolver {
    addr = "208.67.222.222" # resolver1.opendns.com
  }
}

data "nds_nslookup_ptr" "my_ptr" {
  name = data.nds_nslookup_ip.my_ip.data[0]
}

output "my_ip_reverse_name" {
  value = data.nds_nslookup_ptr.my_ptr.data
}

# terraform output 
# my_ip_reverse_name = tolist([
#  "a.b.c.example.com.",
# ])

```
