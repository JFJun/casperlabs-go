package clvalue

const (
	TagBool      = 0
	TagI32       = 1
	TagI64       = 2
	TagU8        = 3
	TagU32       = 4
	TagU64       = 5
	TagU128      = 6
	TagU256      = 7
	TagU512      = 8
	TagUnit      = 9
	TagString    = 10
	TagKey       = 11
	TagURef      = 12
	TagOption    = 13
	TagList      = 14
	TagByteArray = 15
	TagResult    = 16
	TagMap       = 17
	TagTuple1    = 18
	TagTuple2    = 19
	TagTuple3    = 20
	TagAny       = 21
	TagPublicKey = 22
)

type CLValue struct {
	value  CLTypedAndToBytes
	clType int
	bytes  []byte
}

func NewCLValue(value CLTypedAndToBytes, clType int) *CLValue {
	return &CLValue{
		value:  value,
		clType: clType,
		bytes:  value.ToBytes(),
	}
}

type CLTypedAndToBytes interface {
	GetCLType() int
	ToBytes() []byte
}

func clTypeEncoded() []byte {
	//TODO
	//return CLTypeHelper.toBytesHelper(c.clType())
	return nil
}
