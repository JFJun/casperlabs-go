package keys

import (
	"encoding/hex"
	"fmt"
	"github.com/JFJun/casperlabs-go/keys/blake2b"
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
	fmt.Println(hex.EncodeToString(pub))
	fmt.Println(hex.EncodeToString(priv))
}

func TestSECP256K1_GenerateKeyBySeed(t *testing.T) {
	holder := NewKeyGenerator(testSecp256k1)
	priv, pub, err := holder.GenerateKeyBySeed(blake2b.Hash([]byte("abcqwer!")))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(hex.EncodeToString(priv))
	fmt.Println(hex.EncodeToString(pub))
}

func TestSECP256K1_AccountHex(t *testing.T) {
	_, pub := getSECP256K1Key()
	holder := NewKeyHolder(nil, pub, testSecp256k1)

	accountHex, err := holder.AccountHex()
	if err != nil {
		t.Fatal(err)
	}
	if 68 != len(accountHex) {
		t.Errorf("account len error,actual:%d", len(accountHex))
	}
	if accountHex[:2] != "02" {
		t.Fatal("account prefix[:2] error")
	}
	if "0203447239548b66bdfe334131392dd9db386c054989e2b815fe68fd634c9e4703a1" != accountHex {
		t.Fatal("accountHex error")
	}
}

func TestSECP256K1_ParsePublicKeyToPem(t *testing.T) {
	_, pub := getSECP256K1Key()
	holder := NewKeyHolder(nil, pub, testSecp256k1)
	fmt.Println(holder.ParsePublicKeyToPem())
}

func TestSECP256K1_Sign(t *testing.T) {

	priv, pub := getSECP256K1Key()

	msg := blake2b.Hash([]byte("abcde!!"))

	holder := NewKeyHolder(priv, pub, testSecp256k1)
	sig, err := holder.Sign(msg)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(hex.EncodeToString(sig))

	verify, err := holder.Verify(msg, sig)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(verify)

}

func getSECP256K1Key() ([]byte, []byte) {
	priv, _ := hex.DecodeString("be798eee9bb3fa267e0525a7633260c5d2a9512dd2f96b8d621f560dd233d99a")
	pub, _ := hex.DecodeString("03447239548b66bdfe334131392dd9db386c054989e2b815fe68fd634c9e4703a1")
	return priv, pub
}
