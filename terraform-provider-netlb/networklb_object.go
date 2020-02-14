package main

type Vip struct {
	Ipadrress string `json:"ipadrress"`
	Port      int    `json:"port"`
	TLS       bool   `json:"tls_activated"`
}
