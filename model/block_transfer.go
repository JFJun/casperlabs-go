package model

type BlockTransfer struct {
	ApiVersion string     `json:"api_version"`
	BlockHash  string     `json:"block_hash"`
	Transfers  []Transfer `json:"transfers"`
}

type Transfer struct {
	Amount     string      `json:"amount"`
	DeployHash string      `json:"deploy_hash"`
	From       string      `json:"from"`
	Gas        string      `json:"gas"`
	Id         interface{} `json:"id"`
	Source     string      `json:"source"`
	Target     string      `json:"target"`
	To         string      `json:"to"`
}
