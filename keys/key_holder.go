package keys

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/JFJun/casperlabs-go/keys/blake2b"
)

type KeyHolder interface {
	PrivateToPubKey() ([]byte, error)
	//账号十六进制格式
	//注意：这里返回的是accountHex，并非accountHash
	//accountHex代表一个账号的唯一值，从公钥哈希值派生
	//可用作查询账号相关的操作，如余额
	AccountHex() (string, error)
	Sign(message []byte) (sig []byte, err error)
	Verify(message, sig []byte) (bool, error)

	//私钥转换成PEM文件格式（加工过的base64格式）例如：
	//-----BEGIN PRIVATE KEY-----
	//MC4CAQAwBQYDK2VwBCIEIBi2p4YSZ58JCjZuKSdKbFB8ixdrJIZHqNMtaJIuhOF5
	//-----END PRIVATE KEY-----
	ParsePrivateKeyToPem() (string, error)

	//公钥转换成PEM文件格式（加工过的base64格式）例如：
	//-----BEGIN PUBLIC KEY-----
	//MCowBQYDK2VwAyEAeKEooE0MhphnznYVBcR+slT22meCiBHH6WYIs6KKHjw=
	//-----END PUBLIC KEY-----
	ParsePublicKeyToPem() (string, error)
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
	if !isHex(addr) {
		return false
	}
	prefix := addr[:2]
	addrLen := 0
	if prefix == "01" {
		addrLen = 66
	} else if prefix == "02" {
		addrLen = 68
	} else {
		return false
	}
	return addrLen == len(addr)
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

//根据公钥数据生成accountHash
func AccountHash(pub []byte, sa SignatureAlgorithm) (string, error) {
	msg := bytes.Join([][]byte{
		[]byte(sa),
		{0},
		pub,
	}, []byte{})
	return hex.EncodeToString(blake2b.Hash(msg)), nil
}

//根据公钥数据生成accountHex
func AccountHex(pub []byte, prefix string) (string, error) {
	return prefix + hex.EncodeToString(pub), nil
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
