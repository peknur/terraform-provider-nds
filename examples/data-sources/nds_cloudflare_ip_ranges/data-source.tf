terraform {
  required_providers {
    nds = {
      source  = "peknur/nds"
      version = ">= 0.1.0"
    }
  }
}

data "nds_cloudflare_ip_ranges" "ipv4" {
}

output "cf_ipv4" {
  value = data.nds_cloudflare_ip_ranges.ipv4.data
}

# terraform output 
# cf_ipv4 = tolist([
#   "103.21.244.0/22",
#   "103.22.200.0/22",
#   "103.31.4.0/22",
#   "104.16.0.0/13",
#   "104.24.0.0/14",
#   "108.162.192.0/18",
#   "131.0.72.0/22",
#   "141.101.64.0/18",
#   "162.158.0.0/15",
#   "172.64.0.0/13",
#   "173.245.48.0/20",
#   "188.114.96.0/20",
#   "190.93.240.0/20",
#   "197.234.240.0/22",
#   "198.41.128.0/17",
# ])
