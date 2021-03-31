package clvalue

import (
	"fmt"
	"testing"
)

func TestToBytes(t *testing.T) {

	coder, err := NewNumberCoder(TagU256, 256, false, uint32(10000000))
	if err != nil {
		t.Fatal("failed to new number coder")
	}
	byteData := coder.ToBytes()
	fmt.Println(byteData)
}
