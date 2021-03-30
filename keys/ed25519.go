package keys

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"golang.org/x/crypto/ed25519"
)

type ED25519 struct {
	//使用此算法生成的密钥对应的账号前缀
	prefix string
	//使用的算法
	algorithm SignatureAlgorithm
	//生成的公钥字节长度
	pubByteLen int
	//生成的私钥字节长度，注意这里是原始的私钥长度
	privByteLen int

	//保存的私钥数据
	privateKey []byte
	//保存的公钥数据
	pubKey []byte
}

func NewED25519(private []byte, public []byte) *ED25519 {
	return &ED25519{
		prefix:      "01",
		algorithm:   Ed25519,
		pubByteLen:  32,
		privByteLen: 64,
		privateKey:  private,
		pubKey:      public,
	}
}

//注意：这里返回的私钥是64字节长度
func (e *ED25519) GenerateKey() ([]byte, []byte, error) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	if len(pub) != e.pubByteLen || len(priv) != e.privByteLen {
		return nil, nil, errors.New(fmt.Sprintf("%s GenerateKey:invalid key len", e.algorithm))
	}
	if !bytes.Equal(priv[32:], pub[:]) {
		return nil, nil, errors.New(fmt.Sprintf("%s GenerateKey:invalid private key", e.algorithm))
	}
	return priv[:], pub[:], nil
}

//注意：这里返回的私钥是64字节长度
func (e *ED25519) GenerateKeyBySeed(seed []byte) ([]byte, []byte, error) {
	priv := ed25519.NewKeyFromSeed(seed)
	if len(priv) != e.privByteLen {
		return nil, nil, errors.New(fmt.Sprintf("%s GenerateKeyBySeed:invalid key len", e.algorithm))
	}
	fmt.Println(hex.EncodeToString(priv))
	return priv[:], priv[32:], nil
}

func (e *ED25519) PrivateToPubKey() ([]byte, error) {
	if err := CheckPrivKey(e.privateKey, e.privByteLen); err != nil {
		return nil, err
	}
	if len(e.privateKey) != e.privByteLen {
		return nil, errors.New(fmt.Sprintf("%s PrivateToPubKey:invalid key len", e.algorithm))
	}
	return e.privateKey[32:], nil
}

func (e *ED25519) AccountHex() (string, error) {
	if err := CheckPubKey(e.pubKey, e.pubByteLen); err != nil {
		return "", err
	}
	return AccountHex(e.pubKey, e.prefix)
}

func (e *ED25519) Sign(message []byte) (sig []byte, err error) {
	if err := CheckPrivKey(e.privateKey, e.privByteLen); err != nil {
		return nil, err
	}
	priv := ed25519.PrivateKey(e.privateKey)
	return ed25519.Sign(priv, message), nil
}

func (e *ED25519) Verify(message, sig []byte) (bool, error) {
	if err := CheckPubKey(e.pubKey, e.pubByteLen); err != nil {
		return false, err
	}
	pub := ed25519.PublicKey(e.pubKey)
	return ed25519.Verify(pub, message, sig), nil
}

func (e *ED25519) ParsePrivateKeyToPem() (string, error) {
	pkBytes, err := parseKey(e.privateKey[:32], 0, 32)
	if err != nil {
		return "", err
	}
	content := base64.StdEncoding.EncodeToString(bytes.Join([][]byte{
		{48, 46, 2, 1, 0, 48, 5, 6, 3, 43, 101, 112, 4, 34, 4, 32},
		pkBytes,
	}, []byte{}))

	return "-----BEGIN PRIVATE KEY-----\n" + content + "\n" + "-----END PRIVATE KEY-----\n", nil

}

func (e *ED25519) ParsePublicKeyToPem() (string, error) {
	pkBytes, err := parseKey(e.pubKey, 32, 64)
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
