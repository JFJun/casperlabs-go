package keys

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/JFJun/casperlabs-go/keys/blake2b"
	"strings"
)

type KeyHolder interface {
	PrivateToPubKey() ([]byte, error)
	AccountHex() (string, error)
	Sign(message []byte) (sig []byte, err error)
	Verify(message, sig []byte) (bool, error)
}

//根据不同算法构造keyHolder，公钥和私钥入参不一定都需要
//private：私钥
//pub：公钥
//algorithm：具体算法
func NewKeyHolder(private []byte, pub []byte, algorithm SignatureAlgorithm) KeyHolder {
	if algorithm == Secp256K1 {
		return NewSECP256K1(private, pub)
	} else {
		return NewED25519(private, pub)
	}
}

func IsAccount(addr string) bool {
	if has0xPrefix(addr) {
		addr = addr[2:]
	}
	if len(addr) != 66 {
		return false
	}
	if !isHex(addr) {
		return false
	}
	prefix := addr[:2]
	return prefix == "01" || prefix == "02"
}

func CheckPubKey(pub []byte, l int) error {
	if pub == nil {
		return errors.New("CheckPubKey:pubKey require")
	}
	if len(pub) != l {
		return errors.New(fmt.Sprintf("CheckPubKey:invalid pubkey len"))
	}
	return nil
}

func CheckPrivKey(priv []byte, l int) error {
	if priv == nil {
		return errors.New("CheckPrivKey:privKey require")
	}
	if len(priv) != l {
		return errors.New(fmt.Sprintf("CheckPrivKey:invalid privKey len"))
	}
	return nil
}

func has0xPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}

func isHex(str string) bool {
	if len(str)%2 != 0 {
		return false
	}
	for _, c := range []byte(str) {
		if !isHexCharacter(c) {
			return false
		}
	}
	return true
}

func isHexCharacter(c byte) bool {
	return ('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F')
}

//根据公钥数据生成账号
//这里的组合格式参考casperlabs.client-py客户端程序：
/*def account_hash(self) -> bytes:
""" Generate hash of public key and key algorithm for use as primary identifier in the system as bytes """
# account hash is the one place where algorithm is used in upper case.
return crypto.blake2b_hash(
self.algorithm.upper().encode("UTF-8") + b"\x00" + self.public_key
)*/
func AccountHex(pub []byte, prefix string, sa SignatureAlgorithm) (string, error) {
	s := strings.ToUpper(string(sa))
	msg := bytes.Join([][]byte{
		[]byte(s),
		{0},
		pub,
	}, []byte{})
	return prefix + hex.EncodeToString(blake2b.Hash(msg)), nil
}
