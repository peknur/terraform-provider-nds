data "nds_nslookup_ip" "my_ip" {
  name = "myip.opendns.com"
  resolver {
    addr = "208.67.222.222" # resolver1.opendns.com
  }
}

# terraform output my_ip
output "my_ip" {
  value       = data.nds_nslookup_ip.my_ip.data[0]
  description = "My Public IP address"
}
