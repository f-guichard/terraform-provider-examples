package main

type Flow struct {
	Cidr        string `json:"cidr"`
	Port        string `json:"port"`
	Description string `json:"description,omitempty"`
}
