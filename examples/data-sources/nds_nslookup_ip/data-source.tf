terraform {
  required_providers {
    nds = {
      source  = "peknur/nds"
      version = ">= 0.1.0"
    }
  }
}

# Basic host query
data "nds_nslookup_ip" "example" {
  name = "example.com"
}

output "example" {
  value = data.nds_nslookup_ip.example.data
}

# Use custom resolver to get public IP address
data "nds_nslookup_ip" "my_ip" {
  name = "myip.opendns.com"
  resolver {
    addr = "208.67.222.222" # resolver1.opendns.com
  }
}

output "my_ip" {
  value = data.nds_nslookup_ip.my_ip.data
}

# terraform output 
# my_ip = tolist([
#   "xxx.xxx.xxx.xxx",
# ])
# example2 = tolist([
#   "93.184.216.34",
#   "2606:2800:220:1:248:1893:25c8:1946",
# ])

# Retry lookup 5 times with 10 sec interval 
data "nds_nslookup_ip" "retry_example" {
  name           = "example.com"
  retry          = 5
  retry_interval = 10
}
