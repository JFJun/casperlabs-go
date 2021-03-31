package deploy

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"
)

func TestStandardPayment(t *testing.T) {
	payment, err := StandardPayment(big.NewInt(100000))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", hex.EncodeToString(payment.args.args["amount"].ToBytes()	))
}
