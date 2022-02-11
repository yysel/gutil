package buffer

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math"
)

func autoParseUintBytes(num uint, order bool) ([]byte, error) {
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
	return intToBytes(num, LittleEndian)
}

// IntToBBytes 按int长度按大端法自动转换byte数组
func IntToBBytes(num interface{}) ([]byte, error) {
	return intToBytes(num, BigEndian)
}

func intToBytes(num interface{}, order bool) ([]byte, error) {
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
		return p(uint(v), 4)
	case uint32:
		return p(uint(v), 4)
	case uint:
		return autoParseUintBytes(uint(v), order)
	case int:
		return autoParseUintBytes(uint(v), order)
	case int64:
		return p(uint(v), 8)
	case uint64:
		return p(uint(v), 8)
	}
	return nil, errors.New("不可将非整型的值转化为[]byte")
}

// IntToLittleBytes  int类型按小端法转指定长度的bytes,
func IntToLittleBytes(num interface{}, length int) ([]byte, error) {
	return anyType2Bytes(num, length, LittleEndian)
}
func IntToBigBytes(num interface{}, length int) ([]byte, error) {
	return anyType2Bytes(num, length, BigEndian)
}

// Uint2Bytes  uint类型,转为指定排序端法，转指定长度的bytes数组,
func Uint2Bytes(num uint, length int, order bool) ([]byte, error) {
	var orderMethod binary.ByteOrder
	orderMethod = binary.LittleEndian
	if order {
		orderMethod = binary.BigEndian
	}
	a := make([]byte, length)
	var e error
	switch length {
	case 1:
		tmp := int8(num)
		bytesBuffer := bytes.NewBuffer([]byte{})
		e = binary.Write(bytesBuffer, orderMethod, &tmp)
		a = bytesBuffer.Bytes()
	case 2:
		orderMethod.PutUint16(a, uint16(num))
	case 4:
		orderMethod.PutUint32(a, uint32(num))
	case 8:
		orderMethod.PutUint64(a, uint64(num))
	default:
		l := length
		if length < 8 {
			a = make([]byte, 8)
			orderMethod.PutUint64(a, uint64(num))
			if order {
				a = a[8-length : 8]
			} else {
				a = a[0:length]
			}
		} else {
			a = make([]byte, l)
			orderMethod.PutUint64(a, uint64(num))
		}

	}
	return a, e
}

func anyType2Bytes(num interface{}, length int, order bool) ([]byte, error) {
	switch v := num.(type) {
	case int8:
		return Uint2Bytes(uint(v), length, order)
	case uint8:
		return Uint2Bytes(uint(v), length, order)
	case int16:
		return Uint2Bytes(uint(v), length, order)
	case uint16:
		return Uint2Bytes(uint(v), length, order)
	case int32:
		return Uint2Bytes(uint(v), length, order)
	case uint32:
		return Uint2Bytes(uint(v), length, order)
	case uint:
		return Uint2Bytes(uint(v), length, order)
	case int:
		return Uint2Bytes(uint(v), length, order)
	case int64:
		return Uint2Bytes(uint(v), length, order)
	case uint64:
		return Uint2Bytes(uint(v), length, order)
	default:
		return nil, errors.New("不可将非整型的值转化为[]byte")
	}
}
