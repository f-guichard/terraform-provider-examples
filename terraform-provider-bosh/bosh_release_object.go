package main

type Release struct {
	Name     string   `json:"name"`
	Version  string   `json:"version"`
	Hash     string   `json:"hash"`
	JobNames []string `json:"jobnames"`
}
