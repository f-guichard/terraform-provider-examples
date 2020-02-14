provider "dns4hbx" {
    endpoint = "http://192.168.50.1:30005"
    user = "Cosmo"
    password = "Pirates"
}

terraform {
  backend "http" {
    address = "http://192.168.50.1:40000/v1/states/dns"
  }
}

resource "dns4hbx_zone" "zone_s0_p_fti_net" {
  name = "s0.hbx.geo.francetelecom.fr"
}

resource "dns4hbx_record" "cf_prod_priv" {
    name = "api.run.priv.prod.cfy"
    type = "A"
    address = "10.107.11.104"
    zone = "${dns4hbx_zone.zone_s0_p_fti_net.id}"
}

resource "dns4hbx_record" "cf_prod_priv_v6" {
    name = "api.run.priv.prod.cfy"
    type = "AAAA"
    address = "2001:db8::"
    zone = "${dns4hbx_zone.zone_s0_p_fti_net.id}"
}

resource "dns4hbx_record" "cf_prod_pub_v6" {
    name = "*.cfy-app.pub.prod.cfy"
    type = "AAAA"
    address = "2001:db8::"
    zone = "${dns4hbx_zone.zone_s0_p_fti_net.id}"
}

resource "dns4hbx_record" "cf_prod_pub" {
    name = "*.cfy-app.pub.prod.cfy"
    type = "A"
    address = "10.107.11.105"
    zone = "${dns4hbx_zone.zone_s0_p_fti_net.id}"
}