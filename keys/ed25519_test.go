package keys

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/JFJun/casperlabs-go/keys/blake2b"
	"testing"

	"crypto/ed25519"
)

const (
	testEd25519 = "ed25519"
)

func TestED25519_GenerateKey(t *testing.T) {
	holder := NewKeyGenerator(testEd25519)

	priv, pub, err := holder.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(hex.EncodeToString(priv))
	fmt.Println(hex.EncodeToString(pub))
}

func TestED25519_GenerateKey2(t *testing.T) {

	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(hex.EncodeToString(priv))
	fmt.Println(hex.EncodeToString(pub))

	pr := bytes.Join([][]byte{
		{48, 46, 2, 1, 0, 48, 5, 6, 3, 43, 101, 112, 4, 34, 4, 32},
		parseKey(priv[:32], 0, 32),
	}, []byte{})

	pu := bytes.Join([][]byte{
		{48, 42, 48, 5, 6, 3, 43, 101, 112, 3, 33, 0},
		parseKey(pub, 32, 64),
	}, []byte{})

	privBase64 := base64.StdEncoding.EncodeToString(pr)
	pubBase64 := base64.StdEncoding.EncodeToString(pu)

	fmt.Println(privBase64)
	fmt.Println(pubBase64)

}

func parseKey(byteArr []byte, from int, to int) []byte {
	l := len(byteArr)
	// prettier-ignore
	var key []byte
	if l == 32 {
		key = byteArr
	} else {
		if l == 64 {
			key = byteArr[from:to]
		} else {
			if l >= 32 && l < 64 {
				key = byteArr[l%32:]
			}else {
				key = nil
			}
		}
	}
	if key == nil || len(key) != 32 {
		panic("Unexpected key lengt")
	}
	return byteArr
}

func TestED25519_GenerateKeyBySeed(t *testing.T) {
	holder := NewKeyGenerator(testEd25519)
	priv, pub, err := holder.GenerateKeyBySeed([]byte("e1917caa6ef037c0ae2116cab90391aa"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(hex.EncodeToString(priv))
	fmt.Println(hex.EncodeToString(pub))
}

func TestED25519_AccountHex(t *testing.T) {
	_, pub := getED25519Key()
	holder := NewKeyHolder(nil, pub, testEd25519)

	addr, err := holder.AccountHex()
	if err != nil {
		t.Fatal(err)
	}
	if 66 != len(addr) {
		t.Fatal("account len error")
	}
	if addr[:2] != "01" {
		t.Fatal("account prefix[:2] error")
	}
	fmt.Println(addr)
}

func TestED25519_Sign(t *testing.T) {
	priv, pub := getED25519Key()
	msg := blake2b.Hash([]byte("abcde!!"))
	holder := NewKeyHolder(priv, nil, testEd25519)
	sig, err := holder.Sign(msg)
	if err != nil {
		t.Fatal(err)
	}

	//use public key to new a verifyHolder
	holderVerify := NewKeyHolder(nil, pub, testEd25519)
	verify, err := holderVerify.Verify(msg, sig)
	if err != nil {
		t.Fatal(err)
	}
	if !verify {
		t.Fatal("failed to sign message")
	}
}

func getED25519Key() ([]byte, []byte) {
	priv, _ := hex.DecodeString("b98e274c47887ff4a72a8921bbaa045ea12894cebb7ed6d99e76dbdfc784df5b66065ad33dc8adaeb8677690696918aed102be664718434316aca52d51ae3922")
	pub, _ := hex.DecodeString("66065ad33dc8adaeb8677690696918aed102be664718434316aca52d51ae3922")
	return priv, pub
}
