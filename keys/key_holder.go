package keys

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
)

type KeyHolder interface {
	PrivateToPubKey() ([]byte, error)
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

//根据公钥数据生成账号
func AccountHex(pub []byte, prefix string) (string, error) {
	return prefix + hex.EncodeToString(pub), nil
}

func parsePrivateKey(priv []byte) (string, error) {
	pkBytes, err := parseKey(priv[:32], 0, 32)
	if err != nil {
		return "", err
	}
	content := base64.StdEncoding.EncodeToString(bytes.Join([][]byte{
		{48, 46, 2, 1, 0, 48, 5, 6, 3, 43, 101, 112, 4, 34, 4, 32},
		pkBytes,
	}, []byte{}))

	return "-----BEGIN PRIVATE KEY-----\n" + content + "\n" + "-----END PRIVATE KEY-----\n", nil

}

func parsePublicKey(pub []byte) (string, error) {
	pkBytes, err := parseKey(pub, 32, 64)
	if err != nil {
		return "", err
	}
	content := base64.StdEncoding.EncodeToString(bytes.Join([][]byte{
		{48, 42, 48, 5, 6, 3, 43, 101, 112, 3, 33, 0},
		pkBytes,
	}, []byte{}))
	return "-----BEGIN PUBLIC KEY-----\n" + content + "\n" + "-----END PUBLIC KEY-----\n", nil
}

func parseKey(byteData []byte, from int, to int) ([]byte, error) {
	dataLen := len(byteData)
	var key []byte
	if dataLen == 32 {
		key = byteData
	} else {
		if dataLen == 64 {
			key = byteData[from:to]
		} else {
			if dataLen >= 32 && dataLen < 64 {
				key = byteData[dataLen%32:]
			} else {
				key = nil
			}
		}
	}
	if key == nil || len(key) != 32 {
		return nil, errors.New("Unexpected key len")
	}
	return key, nil
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
