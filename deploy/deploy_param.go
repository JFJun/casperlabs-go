package deploy

type param struct {
	accountPublicKey []byte
	chainName        string
	gasPrice         uint64
	ttl              uint64
	dependencies     [][]byte
	timestamp        uint64
}

type transaction struct {
	amount      string
	target      []byte
	sourcePurse []byte
	id          uint64
}
