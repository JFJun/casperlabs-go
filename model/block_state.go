package model

type BlockState struct {
	ApiVersion  string                `json:"api_version"`
	StoredValue BlockStateStoredValue `json:"stored_value"`
}

type BlockStateStoredValue struct {
	Account BlockStateAccount `json:"Account"`
}

type BlockStateAccount struct {
	AccountHash string `json:"account_hash"`
	MainPurse   string `json:"main_purse"`
}
