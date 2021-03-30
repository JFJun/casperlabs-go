package keys

import (
	"encoding/hex"
	"fmt"
	"github.com/JFJun/casperlabs-go/keys/blake2b"
	"testing"
)

type validedKey struct {
	acccountHex   string
	privatekeyHex string
	pubKeyHex     string
	publicKeyPem  string
	privateKeyPem string
}

// 已被验证的有效的key数据
var ed25519VK = validedKey{
	acccountHex:   "0178a128a04d0c869867ce761505c47eb254f6da67828811c7e96608b3a28a1e3c",
	privatekeyHex: "18b6a78612679f090a366e29274a6c507c8b176b248647a8d32d68922e84e17978a128a04d0c869867ce761505c47eb254f6da67828811c7e96608b3a28a1e3c",
	pubKeyHex:     "78a128a04d0c869867ce761505c47eb254f6da67828811c7e96608b3a28a1e3c",
	//PEM私钥文件，完整格式：
	//-----BEGIN PRIVATE KEY-----
	//MC4CAQAwBQYDK2VwBCIEIBi2p4YSZ58JCjZuKSdKbFB8ixdrJIZHqNMtaJIuhOF5
	//-----END PRIVATE KEY-----
	privateKeyPem: "MC4CAQAwBQYDK2VwBCIEIBi2p4YSZ58JCjZuKSdKbFB8ixdrJIZHqNMtaJIuhOF5",
	//PEM公钥文件，完整格式：
	//-----BEGIN PUBLIC KEY-----
	//MCowBQYDK2VwAyEAeKEooE0MhphnznYVBcR+slT22meCiBHH6WYIs6KKHjw=
	//-----END PUBLIC KEY-----
	publicKeyPem: "MCowBQYDK2VwAyEAeKEooE0MhphnznYVBcR+slT22meCiBHH6WYIs6KKHjw=",
}

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

func TestED25519_ParsePublicKeyToPem(t *testing.T) {
	pub, _ := hex.DecodeString(ed25519VK.pubKeyHex)
	holder := NewKeyHolder(nil, pub, testEd25519)
	pem, err := holder.ParsePublicKeyToPem()
	if err != nil {
		t.Fatal("failed parse to public PEM")
	}
	if pem != ed25519VK.publicKeyPem {
		t.Fatal("public PEM error")
	}
}

func TestED25519_ParsePrivateKeyToPem(t *testing.T) {
	priv, _ := hex.DecodeString(ed25519VK.privatekeyHex)
	holder := NewKeyHolder(priv, nil, testEd25519)
	pem, err := holder.ParsePrivateKeyToPem()
	if err != nil {
		t.Fatal("failed parse to private PEM")
	}
	if pem != ed25519VK.privateKeyPem {
		t.Fatal("private PEM error")
	}
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
	pub, _ := hex.DecodeString(ed25519VK.pubKeyHex)
	holder := NewKeyHolder(nil, pub, testEd25519)

	addr, err := holder.AccountHex()
	if err != nil {
		t.Fatal(err)
	}
	if 66 != len(addr) {
		t.Fatal("accountHex len error")
	}
	if addr[:2] != "01" {
		t.Fatal("accountHex prefix[:2] error")
	}
	if addr != ed25519VK.acccountHex {
		t.Fatal("accountHex error")
	}
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
