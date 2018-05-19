provider "bosh" {
    endpoint_boshdirector = "http://192.168.50.1:30002"
    user_boshdirector = "Cosmo"
    password_boshdirector = "Pirates"
}

terraform {
  backend "http" {
    address = "http://192.168.50.1:40000/v1/states/bosh"
  }
}

resource "bosh_stemcell" "dfy_stemcell_xenial" {
    name = "ubuntu-xenial"
    version = "3541.12"
    os = "ubuntu"
}

resource "bosh_stemcell" "dfy_stemcell_trusty" {
    name = "ubuntu-trusty"
    version = "3421.43"
    os = "ubuntu"
}

resource "bosh_release" "dfy_release_cassandra" {
    name = "cassandra"
    version = "3.11"
    hash = "663bb5687ae98917ff048272f0d8fcc013ea33f4f9e34531b65d2cf2c4d77eac35b292023a78cead735d7f642acb7d5f"
    jobs = ["ctools", "notifications", "index", "ioj", "cache", "db", "audit"]
}

resource "bosh_release" "dfy_release_hbase" {
    name = "hbase"
    version = "1.26"
    hash = "88a294678b35de0b5576391ca68a07358715f8e266e6dfb612e9fbc760acb932ca9e3dfe091275ddbfb3595f665040db"
    jobs = ["metrics_api","replication","hserver","backup"]
}

resource "bosh_release" "dfy_release_spark" {
    name = "spark"
    version = "2.1.2"
    hash = "11e23ab1629517a88100dd157b16e1a2358a4d321c6d2f9f9a361382acdc615f799e0236d1e27a63c35966f3790691f8"
    jobs = ["streaming","rlauncher"]
}

resource "bosh_release" "dfy_release_zookeeper" {
    name = "zookeeper"
    version = "3.4.12"
    hash = "5278b470b4e126c9396d38bdce4045edb19cd9c6227f3d0bf6c7392ec55494981a5b26965ade08c94f14cdb805199893"
    jobs = ["server","admin","jmx"]
}

resource "bosh_release" "dfy_release_mahout" {
    name = "mahout"
    version = "0.13.0"
    hash = "70a7987a267a30ae66a7c77a41614f4c90d966c68d87eb9a5fe3361e892a1399aa5383f7ebf1687f6e95bcaf6383126a"
    jobs = ["clustering", "classifier", "taste"]
}

resource "bosh_deployment" "dfy_ServiceProcessing" {
    name = "ml_back_streamer"
    releases = [ "${bosh_release.dfy_release_zookeeper.id}", "${bosh_release.dfy_release_mahout.id}", "${bosh_release.dfy_release_spark.id}" ]
    stemcell = [ "${bosh_stemcell.dfy_stemcell_trusty.id}" ]

    depends_on = [ "bosh_stemcell.dfy_stemcell_trusty", "bosh_release.dfy_release_zookeeper", "bosh_release.dfy_release_mahout", "bosh_release.dfy_release_spark" ]
}


