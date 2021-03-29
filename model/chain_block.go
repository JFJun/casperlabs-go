package model

type ChainBlock struct {
	ApiVersion string      `json:"api_version"`
	Block      CasperBlock `json:"block"`
}

type CasperBlock struct {
	Hash   string            `json:"hash"`
	Body   CasperBlockBody   `json:"body"`
	Header CasperBlockHeader `json:"header"`
}
type CasperBlockBody struct {
	DeployHashes   []string `json:"deploy_hashes"`
	Proposer       string   `json:"proposer"`
	TransferHashes []string `json:"transfer_hashes"`
}
type CasperBlockHeader struct {
	AccumulatedSeed string `json:"accumulated_seed"`
	BodyHash        string `json:"body_hash"`
	EraId           int64  `json:"era_id"`
	Height          int64  `json:"height"`
	ParentHash      string `json:"parent_hash"`
	StateRootHash   string `json:"state_root_hash"`
	Timestamp       string `json:"timestamp"`
}
