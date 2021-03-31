package numutil

import (
	"errors"
	"math/big"
)

func StrToBigInt(value string) (*big.Int, error) {
	if bi, failed := new(big.Int).SetString(value, 10); failed {
		return nil, errors.New("failed to conv str to bigint")
	} else {
		return bi, nil
	}
}
