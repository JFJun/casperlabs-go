package keys

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/JFJun/casperlabs-go/keys/blake2b"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"math/big"
)

type SECP256K1 struct {
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

func NewSECP256K1(private []byte, public []byte) *SECP256K1 {
	return &SECP256K1{
		prefix:      "02",
		algorithm:   Secp256K1,
		pubByteLen:  65,
		privByteLen: 32,
		privateKey:  private,
		pubKey:      public,
	}
}

func (s *SECP256K1) GenerateKey() ([]byte, []byte, error) {
	priv, err := ethcrypto.GenerateKey()
	if err != nil {
		return nil, nil, err
	}
	l := len(priv.D.Bytes())
	if l != s.privByteLen {
		return nil, nil, errors.New(fmt.Sprintf("%s genrateKey:invalid key len", s.algorithm))
	}
	return priv.D.Bytes(), ethcrypto.FromECDSAPub(&priv.PublicKey), nil
}

func (s *SECP256K1) GenerateKeyBySeed(seed []byte) ([]byte, []byte, error) {
	cruve := secp256k1.S256()
	priv := new(ecdsa.PrivateKey)
	priv.D = new(big.Int).SetBytes(seed)
	priv.X, priv.Y = cruve.ScalarBaseMult(priv.D.Bytes())
	comPub := secp256k1.CompressPubkey(priv.X, priv.Y)
	pub := blake2b.Hash(comPub)
	return priv.D.Bytes(), pub[:], nil
}

func (s *SECP256K1) PrivateToPubKey() ([]byte, error) {
	if err := CheckPrivKey(s.privateKey, s.privByteLen); err != nil {
		return nil, err
	}
	if len(s.privateKey) != s.privByteLen {
		return nil, errors.New(fmt.Sprintf("%s PrivateToPubKey:invalid key len", s.algorithm))
	}
	priv, err := ethcrypto.HexToECDSA(hex.EncodeToString(s.privateKey))
	if err != nil {
		return nil, err
	}
	return ethcrypto.FromECDSAPub(&priv.PublicKey), nil
}

func (s *SECP256K1) AccountHex() (string, error) {
	if err := CheckPubKey(s.pubKey, s.pubByteLen); err != nil {
		return "", err
	}
	return AccountHex(s.pubKey, s.prefix, s.algorithm)
}

func (s *SECP256K1) Sign(message []byte) (sig []byte, err error) {
	if err := CheckPrivKey(s.privateKey, s.privByteLen); err != nil {
		return nil, err
	}
	priv, err := ethcrypto.ToECDSA(s.privateKey)
	if err != nil {
		return nil, err
	}
	return ethcrypto.Sign(message, priv)
}

func (s *SECP256K1) Verify(message, sig []byte) (bool, error) {
	if err := CheckPubKey(s.pubKey, s.pubByteLen); err != nil {
		return false, err
	}
	return ethcrypto.VerifySignature(s.pubKey, message, sig), nil
}
