package main

type Deployment struct {
	Name      string   `json:"name"`
	Releases  []string `json:"releases"`
	Stemcells []string `json:"stemcells"`
}
