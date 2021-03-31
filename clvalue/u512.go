package clvalue

import "math/big"

type U512 struct {
	NumberCoder
	value big.Int
}

func NewU512(value *big.Int) (*U512, error) {
	coder, err := NewNumberCoder(TagU512, 512, false, value)
	if err != nil {
		return nil, err
	}
	return &U512{
		NumberCoder: *coder,
	}, err
}
