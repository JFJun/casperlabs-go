package keys

import (
	"encoding/hex"
	"errors"
	"fmt"
	"golang.org/x/crypto/blake2b"
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

/*
write by flynn
*/
func AddressToAccountHash(address string) ([]byte, error) {
	pub, err := hex.DecodeString(address)
	if err != nil {
		return nil, err
	}
	if len(pub) != 33 {
		return nil, fmt.Errorf("address length is not equal 33,len=[%d]", len(pub))
	}
	calcAccountHash := func(prefix string, rawPublicKey []byte) []byte {
		var data []byte
		data = append(data, []byte(prefix)...)
		data = append(data, 0x00)
		data = append(data, rawPublicKey...)
		accountHash := blake2b.Sum256(data)
		return accountHash[:]
	}
	var prefix string
	if pub[0] == 0x01 {
		prefix = "ed25519"
	} else if pub[0] == 0x02 {
		prefix = "secp256k1"
	} else {
		return nil, fmt.Errorf("unkown signature algorithm")
	}
	return calcAccountHash(prefix, pub[1:]), nil

}
