data "nds_nslookup_ip" "my_ip" {
  name = "myip.opendns.com"
  resolver {
    addr = "208.67.222.222" # resolver1.opendns.com
  }
}
