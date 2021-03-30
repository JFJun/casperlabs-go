package keys

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/JFJun/casperlabs-go/keys/blake2b"
	eth256k1 "github.com/ethereum/go-ethereum/crypto/secp256k1"
	"testing"
)

const (
	testSecp256k1 = "secp256k1"
)

func TestSECP256K1_GenerateKey(t *testing.T) {
	holder := NewKeyGenerator(testSecp256k1)

	priv, pub, err := holder.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(hex.EncodeToString(priv))
	fmt.Println(hex.EncodeToString(pub))
}

func TestSECP256K1_AccountHex(t *testing.T) {
	_, pub := getSECP256K1Key()
	holder := NewKeyHolder(nil, pub, testSecp256k1)

	addr, err := holder.AccountHex()
	if err != nil {
		t.Fatal(err)
	}
	if 66 != len(addr) {
		t.Fatal("account len error")
	}
	if addr[:2] != "02" {
		t.Fatal("account prefix[:2] error")
	}
	fmt.Println(addr)
}

func TestSECP256K1_Sign(t *testing.T) {

	priv, pub := getSECP256K1Key()

	sig := blake2b.Hash([]byte("abcde!!"))

	//sig2 := []byte("abcde!!2")

	holder := NewKeyHolder(priv, pub, testSecp256k1)
	msg, err := holder.Sign(sig)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(hex.EncodeToString(msg))

	verify, err := holder.Verify(msg, sig)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(verify)

}

func TestSECP256K1_Sign2(t *testing.T) {
	//priv, pub ,_:= GenerateKey()
	//digestHash := Blake2bHash([]byte("abcde!!"))
	hashed := []byte("testing")

	cruve := eth256k1.S256()
	priv, err := ecdsa.GenerateKey(cruve, rand.Reader)

	r, s, _ := ecdsa.Sign(rand.Reader, priv, hashed)

	verify := ecdsa.Verify(&priv.PublicKey, hashed, r, s)

	//sig, err := ethcrypto.Sign(digestHash, key)
	if err != nil {
		t.Fatal(err)
	}
	//fmt.Println(hex.EncodeToString(sig))
	//signature := eth256k1.VerifySignature(pub, digestHash, sig)
	//signature := ethcrypto.VerifySignature(ethcrypto.FromECDSAPub(&pub), digestHash, sig)
	fmt.Println(verify)

}

func getSECP256K1Key() ([]byte, []byte) {
	priv, _ := hex.DecodeString("7465d722f903cd362db4420da89b9c465289368f295be93e03cb76ba25dedce3")
	pub, _ := hex.DecodeString("04add9a85748a019b12acd8212a4447f9de2f85d14e73838e0b86fb4c36467018f5f4eddb0b5b68b2397eab74b07876b80d6ee9ca1f1a01971d6e9701986df2c66")
	return priv, pub
}
