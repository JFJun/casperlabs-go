package keys

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"golang.org/x/crypto/ed25519"
)

type ED25519 struct {
	//使用此算法生成的秘钥对应的账号前缀
	prefix string
	//使用的算法
	algorithm SignatureAlgorithm
	//生成的公钥字节长度
	pubByteLen int
	//生成的私钥字节长度，注意这里是原始的私钥长度
	privByteLen int

	//保持的私钥数据
	privateKey []byte
	//保持的公钥数据
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
	return AccountHex(e.pubKey, e.prefix, e.algorithm)
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
