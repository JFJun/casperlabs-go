package keys

import (
	"errors"
)

type SignatureAlgorithm string

const (
	Ed25519   = SignatureAlgorithm("ed25519")
	Secp256K1 = SignatureAlgorithm("secp256k1")
)

/*
随机生成公私钥
*/
func GenerateKeys(sa SignatureAlgorithm) (private, pub []byte, err error) {
	generator := NewKeyGenerator(sa)
	return generator.GenerateKey()
}
func GenerateKeysBySeed(seed []byte, sa SignatureAlgorithm) (private, pub []byte, err error) {
	generator := NewKeyGenerator(sa)
	return generator.GenerateKeyBySeed(seed)
}

func PrivateToPubKey(private []byte, sa SignatureAlgorithm) (pub []byte, err error) {
	holder := NewKeyHolder(private, nil, sa)
	return holder.PrivateToPubKey()
}

func PublicKeyToAddress(pub []byte, sa SignatureAlgorithm) (address string, err error) {
	holder := NewKeyHolder(nil, pub, sa)
	return holder.AccountHex()
}

func ValidAddress(address string) error {
	if !IsAccount(address) {
		return errors.New("invalid address")
	}
	return nil
}

func Sign(private []byte, message []byte, sa SignatureAlgorithm) (sig []byte, err error) {
	holder := NewKeyHolder(private, nil, sa)
	return holder.Sign(message)
}
