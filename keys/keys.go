package keys

type SignatureAlgorithm string

const (
	Ed25519   = SignatureAlgorithm("ed25519")
	Secp256K1 = SignatureAlgorithm("secp256k1")
)

/*
随机生成公私钥
*/
func GenerateKeys(sa SignatureAlgorithm) (private, pub []byte, err error) {
	//todo
	return nil, nil, err
}
func GenerateKeysBySeed(seed []byte, sa SignatureAlgorithm) (private, pub []byte, err error) {
	//todo
	return nil, nil, err
}

func PrivateToPubKey(private []byte, sa SignatureAlgorithm) (pub []byte, err error) {
	//todo
	return nil, err
}

func PublicKeyToAddress(pub []byte) (address string, err error) {
	//todo
	return "", err
}

func ValidAddress(address string) error {
	return nil
}

func Sign(message []byte, sa SignatureAlgorithm) (sig []byte, err error) {
	return nil, err
}
