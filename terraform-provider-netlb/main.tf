provider "networklb" {
    endpoint_networklb = "http://192.168.50.1:30001"
    user_networklb = "Cosmo"
    password_networklb = "Pirates"
}

terraform {
  backend "http" {
    address = "http://192.168.50.1:40000/v1/states/networklb"
  }
}

resource "networklb_vip" "pamelatfp" {
    ipadrress = "193.252.14.36"
    port = 21
    tls_activated = true
}

resource "networklb_vip" "pamelaimaps" {
    ipadrress = "193.252.14.36"
    port = 993
    tls_activated = true
}

resource "networklb_vip" "pamelahttp" {
    ipadrress = "193.252.14.36"
    port = 80
    tls_activated = false
}
