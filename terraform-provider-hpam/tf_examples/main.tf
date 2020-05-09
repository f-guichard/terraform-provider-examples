provider "hpam" {
  endpoint = "http://127.0.0.1:30026"
  user     = "assets"
  password = "time"
}

//terraform {
//  backend "http" {
//    address = "http://127.0.0.1:40000/v1/states/assets"
//  }
//}

resource "hpam_basicat" "powerup_basicat" {
  name = "yy"
}

resource "hpam_basicat" "powerup2_basicat" {
  name = "iiiiiu"

}
resource "hpam_application" "powerup_app_name" {
  modapp = "xx"
  name   = "xx"
}

resource "hpam_soutien" "powerup_soutien" {
  exploit      = "xx"
  support      = "xx"
  resp_exploit = "xx"
  resp_support = "xx"
  etat         = "xx"
  eds_peit     = "xx"
  eds_sat      = "xx"
}

