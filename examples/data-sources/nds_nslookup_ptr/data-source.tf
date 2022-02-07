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
