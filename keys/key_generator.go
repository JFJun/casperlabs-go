package keys



type KeyGenerator interface {
	//生成秘钥对
	//return:私钥 公钥
	GenerateKey() ([]byte, []byte, error)

	//根据种子生成秘钥对
	//return:私钥 公钥
	GenerateKeyBySeed(seed []byte) ([]byte, []byte, error)
}

//根据不同签名算法，生成对应的keyGenerator
//目前支持ed2519/secp256k1
func NewKeyGenerator(algorithm SignatureAlgorithm) KeyGenerator {
	if algorithm == Secp256K1 {
		return NewSECP256K1(nil, nil)
	} else {
		return NewED25519(nil, nil)
	}
}
