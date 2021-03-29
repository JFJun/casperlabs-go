package model

type ChainStatus struct {
	ApiVersion    string `json:"api_version"`
	BuildVersion  string `json:"build_version"`
	ChainspecName string `json:"chainspec_name"`
}
