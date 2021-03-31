package deploy

import (
	cl "github.com/JFJun/casperlabs-go/clvalue"
	"math/big"
)

func StandardPayment(paymentAmount *big.Int) (*ModuleBytes, error) {
	u512, err := cl.NewU512(paymentAmount)
	if err != nil {
		return nil, err
	}

	argMap := map[string]cl.CLTypedAndToBytes{}
	argMap["amount"] = u512
	ra := RuntimeArgs{
		args: argMap,
	}
	return NewModuleBytes([]byte{}, ra), nil
}
