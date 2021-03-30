package blake2b

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func Test_Blake2bHash(t *testing.T) {
	h := Hash([]byte("123456789"))
	ds, _ := hex.DecodeString("16e0bf1f85594a11e75030981c0b670370b3ad83a43f49ae58a2fd6f6513cde9")
	if !bytes.Equal(h, ds) {
		t.Fatal("blake2b hash calc error")
	}
}
