provider "flownet" {
    endpoint_flownet = "http://127.0.0.1:30000"
    user_flownet = "C6PO"
    password_flownet = "R2D2" //A remplacer par l'utilisation de l'env FLOWNET_PASSWORD
}

terraform {
  backend "http" {
    address = "http://127.0.0.1:40000/v1/states/flownet"
  }
}

resource "flownet_flow" "pamela" {
    cidr = "88.10.189.156"
    port = "80;443"
    description = "Partenaire Pamela"
}

resource "flownet_flow" "charly" {
    cidr = "10.107.25.149"
    port = "1024-65535"
    description = "Charles Marketplace"
}

resource "flownet_flow" "hasseloff" {
    cidr = "82.75.129.88"
    port = "1024-65535"
    description = "I Master"
}

resource "flownet_flow" "cardif" {
    cidr = "82.90.211.239"
    port = "1024-65535"
    description = "R Judge"
}

resource "flownet_flow" "fnet" {
    cidr = "82.90.211.240"
    port = "65535"
    description = "Gloppy"
}

resource "flownet_flow" "gibbly" {
    cidr = "82.90.211.241"
    port = "4455"
}

