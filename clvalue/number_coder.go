package clvalue

import (
	"bytes"
	"errors"
	"math/big"
	"strconv"
)

const (
	hexPrefix = "0x"
)

//type BigNumber struct {
//	hex         string
//	isBigNumber bool
//}
//
//func BigNumberFrom(value interface{}) (*big.Int, error) {
//	switch v := value.(type) {
//	case big.Int:
//	case *big.Int:
//		return NewBigNumber(hex.EncodeToString(v.Bytes())), nil
//	case string:
//		if hexutil.Has0xPrefix(v) {
//			v = v[:2]
//		}
//		if hexutil.IsHex(v) {
//			return NewBigNumber(v), nil
//		}
//		if numutil.IsNum(v) {
//
//			n, ok := new(big.Int).SetString(v, 10)
//			if !ok {
//				return nil, errors.New("failed to new BigNumber from string")
//			}
//			return NewBigNumber(hex.EncodeToString(n.Bytes())), nil
//		}
//		return nil, errors.New("invalid BigNumber string")
//	case []byte:
//		return NewBigNumber(hex.EncodeToString(v)), nil
//	case *[]byte:
//		return NewBigNumber(hex.EncodeToString(*v)), nil
//
//	case uint:
//	case uint8:
//	case uint16:
//	case uint32:
//	case uint64:
//	case int:
//	case int8:
//	case int16:
//	case int32:
//	case int64:
//		return NewBigNumber(strconv.Itoa(int(v))), nil
//	}
//	return nil, errors.New("invalid BigNumber value")
//}
//
//func NewBigNumber(hex string) *big.Int {
//	return &BigNumber{
//		hex:         hexPrefix + hex,
//		isBigNumber: true,
//	}
//}

type NumberCoder struct {
	clType  int
	bitSize uint32
	signed  bool
	val     *big.Int
	name    string
}

func NewNumberCoder(clType int, bitSize uint32, signed bool, val interface{}) (*NumberCoder, error) {
	bigNum, err := numberFrom(val)
	if err != nil {
		return nil, err
	}
	n := "u"
	if signed {
		n = "i"
	}
	return &NumberCoder{
		clType:  clType,
		bitSize: bitSize,
		signed:  signed,
		val:     bigNum,
		name:    n + strconv.Itoa(int(bitSize)),
	}, nil
}
func (nc *NumberCoder) GetCLType() int {
	return 0
}

func (nc *NumberCoder) ToBytes() []byte {
	vb := nc.val.Bytes()
	if nc.val.Cmp(big.NewInt(0)) >= 0 {
		// for positive number, we had to deal with paddings
		if nc.bitSize > 64 {
			// for u128, u256, u512, we have to and append extra byte for length
			r := bytes.Join([][]byte{
				vb,
				{byte(len(vb))},
			}, []byte{})
			byteReverse(&r)
			return r
		} else {
			// for other types, we have to add padding 0s
			byteLen := nc.bitSize / 8
			var b []byte
			for i := 0; i < int(byteLen)-len(vb); i++ {
				b = append(b, 0)
			}
			byteReverse(&vb)
			r := bytes.Join([][]byte{
				vb,
				b[:],
			}, []byte{})
			return r
		}
	}
	byteReverse(&vb)
	return vb
}

func byteReverse(s *[]byte) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func numberFrom(value interface{}) (*big.Int, error) {
	switch v := value.(type) {
	case big.Int:
		return &v, nil
	case *big.Int:
		return v, nil
	//case uint, uint8, uint16, uint32, int:
	//	return big.NewInt(int64(v)), nil
	case int8:
		return big.NewInt(int64(v)), nil
	case int16:
		return big.NewInt(int64(v)), nil
	case int32:
		return big.NewInt(int64(v)), nil
	case uint32:
		return big.NewInt(int64(v)), nil
	case uint64:
		i := int64(v)
		return big.NewInt(i), nil
	case int64:
		return big.NewInt(v), nil
	}
	return nil, errors.New("invalid Number value")
}
