package buffer

import (
	"bytes"
	"encoding/binary"
	"math"
)

type littleEndian struct{}

type bigEndian struct{}

func autoParseUintBytes(num uint, order binary.ByteOrder) ([]byte, error) {
	switch {
	case num <= math.MaxUint8:
		return Uint2Bytes(num, 1, order)
	case num <= math.MaxUint16:
		return Uint2Bytes(num, 2, order)
	case num <= math.MaxUint32:
		return Uint2Bytes(num, 4, order)
	default:
		return Uint2Bytes(num, 8, order)
	}
}

// IntToLBytes 按int长度按小端法自动转换byte数组
func IntToLBytes(num interface{}) ([]byte, error) {
	return intToBytes(num, binary.LittleEndian)
}

// IntToBBytes 按int长度按大端法自动转换byte数组
func IntToBBytes(num interface{}) ([]byte, error) {
	return intToBytes(num, binary.BigEndian)
}

func intToBytes(num interface{}, order binary.ByteOrder) ([]byte, error) {
	p := func(n uint, l int) ([]byte, error) {
		return Uint2Bytes(n, l, order)
	}
	switch v := num.(type) {
	case int8:
		return p(uint(v), 1)
	case uint8:
		return p(uint(v), 1)
	case int16:
		return p(uint(v), 2)
	case uint16:
		return p(uint(v), 2)
	case int32:
		return p(uint(v), 2)
	case uint32:
		return p(uint(v), 2)
	case uint:
		return autoParseUintBytes(uint(v), order)
	case int:
		return autoParseUintBytes(uint(v), order)
	case int64:
		return p(uint(v), 2)
	case uint64:
		return p(uint(v), 2)
	}
	return p(num.(uint), 8)
}

// IntToLittleBytes  int类型按小端法转指定长度的bytes,
func IntToLittleBytes(num interface{}, length int) ([]byte, error) {
	v := num.(uint)
	return Uint2Bytes(v, length, binary.LittleEndian)
}
func IntToBigBytes(num interface{}, length int) ([]byte, error) {
	v := num.(uint)
	return Uint2Bytes(v, length, binary.BigEndian)
}

// Uint2Bytes  uint类型,转为指定排序端法，转指定长度的bytes数组,
func Uint2Bytes(num uint, length int, order binary.ByteOrder) ([]byte, error) {
	a := make([]byte, length)
	var e error
	switch length {
	case 1:
		tmp := int8(num)
		bytesBuffer := bytes.NewBuffer([]byte{})
		e = binary.Write(bytesBuffer, order, &tmp)
		a = bytesBuffer.Bytes()
	case 2:
		order.PutUint16(a, uint16(num))
	case 3:
		order.PutUint32(a, uint32(num))
		a = a[:3]
	case 4:
		order.PutUint32(a, uint32(num))
	case 8:
		order.PutUint64(a, uint64(num))
	default:
		order.PutUint64(a, uint64(num))
	}
	return a, e
}
