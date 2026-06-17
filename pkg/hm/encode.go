package hm

import (
	"fmt"
)

// Public Func
func EncodeValue(v Value) (EncodedValue, error) {
	switch v.Kind {

	case Bool:
		return encodeBool(v.Raw)

	case String:
		return encodeString(v.Raw)

	case Byte:
		return encodeByte(v.Raw)

	case Int4:
		return encodeSignedInt(v.Raw, 4, Int4)

	case Int8:
		return encodeSignedInt(v.Raw, 8, Int8)

	case Int16:
		return encodeSignedInt(v.Raw, 16, Int16)

	case Int32:
		return encodeSignedInt(v.Raw, 32, Int32)

	case Int64:
		return encodeSignedInt(v.Raw, 64, Int64)

	default:
		return EncodedValue{}, fmt.Errorf("Unsupported kind: %s", v.Kind)

	}
}

// Private Exposed to Public function
func encodeBool(raw any) (EncodedValue, error) {
	b, ok := raw.(bool)

	if !ok {
		return EncodedValue{}, fmt.Errorf("Expected Bool Type but received something else")
	}

	bit := uint8(0)

	if b {
		bit = 1
	}

	return EncodedValue{
		Kind:  Bool,
		Bytes: []byte{bit},
		Bits:  []uint8{bit},
	}, nil
}

func encodeSignedInt(raw any, bitSize int, kind Kind) (EncodedValue, error) {

	n, ok := raw.(int64)

	if !ok {
		return EncodedValue{}, fmt.Errorf("Expected int64 type but received %s", kind)
	}

	min, max := signedRange(bitSize)

	if n < min || n > max {
		return EncodedValue{}, fmt.Errorf("%s value %d is out of range [%d, %d]", kind, n, min, max)
	}

	bits, err := signedIntToBits(n, bitSize)

	if err != nil {
		return EncodedValue{}, err
	}

	return EncodedValue{
		Kind:  kind,
		Bytes: nil,
		Bits:  bits,
	}, nil
}

func encodeByte(raw any) (EncodedValue, error) {
	b, ok := raw.(byte)
	if !ok {
		return EncodedValue{}, fmt.Errorf("Expected Byte Type but received: %s", raw)
	}

	return EncodedValue{
		Kind:  Byte,
		Bytes: []byte{b},
		Bits:  uintToBits(uint64(b), 8),
	}, nil
}

func encodeString(raw any) (EncodedValue, error) {
	str, ok := raw.(string)
	if !ok {
		return EncodedValue{}, fmt.Errorf("Expected String Type but received: %s", raw)
	}

	return EncodedValue{
		Kind:  String,
		Bytes: []byte(str),
		Bits:  nil,
	}, nil
}

// Private exposed to Private function
func uintToBits(value uint64, size int) []uint8 {
	bits := make([]uint8, size)

	for i := 0; i < size; i++ {
		bitPosition := size - 1 - i
		bit := (value >> bitPosition) & 1
		bits[i] = uint8(bit)
	}
	return bits
}

func signedRange(bitSize int) (int64, int64) {
	switch bitSize {
	case 4:
		return -8, 7
	case 8:
		return -128, 127
	case 16:
		return -32768, 32767
	case 32:
		return -2147483648, 2147483647
	case 64:
		return -9223372036854775808, 9223372036854775807

	default:
		return 0, 0
	}
}

func signedIntToBits(n int64, bitSize int) ([]uint8, error) {
	switch bitSize {
	case 4:
		return uintToBits(uint64(uint8(n)&15), 4), nil

	case 8:
		return uintToBits(uint64(uint8(n)), 8), nil

	case 16:
		return uintToBits(uint64(uint16(n)), 16), nil

	case 32:
		return uintToBits(uint64(uint32(n)), 32), nil

	case 64:
		return uintToBits(uint64(n), 64), nil

	default:
		return nil, fmt.Errorf("Unexpected int size, expected int4, int8, int16, int32, int64, received bitsize: %d", bitSize)
	}
}
