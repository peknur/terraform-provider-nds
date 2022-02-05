data "nds_nslookup_srv" "test" {
  name    = "example.com"
  proto   = "tcp"
  service = "xmpp-client"
}
