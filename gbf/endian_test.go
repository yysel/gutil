package gbf

import (
	"bytes"
	"testing"
)

func Test_bigEndian_ToBytes(t *testing.T) {
	b := Be.ToBytes(128, 3)
	m := []byte{0, 0, 128}
	if bytes.Compare(b, m) != 0 {
		t.Error()
	}
}
func Test_bigEndian_ToUint(t *testing.T) {
	type M struct {
		m []byte
		v uint
	}
	fs := []M{
		{[]byte{0, 0, 128}, 128},
		{[]byte{1, 0, 0, 0, 0, 0, 0, 0, 128}, 128},
		{[]byte{0x6d, 0x00, 0x61, 0x00, 0x73, 0x00, 0x74, 0x00}, 7854384404691448832},
	}
	for _, f := range fs {
		if Be.ToUint(f.m) != f.v {
			t.Error()
		}
	}
}

func Test_bigEndian_ToUint16(t *testing.T) {
	type M struct {
		m []byte
		v uint16
	}
	fs := []M{
		{[]byte{128}, 128},
		{[]byte{1, 0, 0, 0, 0, 0, 0, 0, 128}, 128},
	}
	for _, f := range fs {
		if Be.ToUint16(f.m) != f.v {
			t.Error()
		}
	}
}

func Test_bigEndian_ToUint32(t *testing.T) {
	type M struct {
		m []byte
		v uint32
	}
	fs := []M{
		{[]byte{0, 0, 128}, 128},
		{[]byte{1, 0, 0, 0, 0, 0, 0, 0, 128}, 128},
	}
	for _, f := range fs {
		if Be.ToUint32(f.m) != f.v {
			t.Error()
		}
	}
}

func Test_bigEndian_ToUint64(t *testing.T) {
	type M struct {
		m []byte
		v uint64
	}
	fs := []M{
		{[]byte{0, 0, 128}, 128},
		{[]byte{1, 0, 0, 0, 0, 0, 0, 0, 128}, 128},
	}
	for _, f := range fs {
		if Be.ToUint64(f.m) != f.v {
			t.Error()
		}
	}
}

func Test_bigEndian_ToUint8(t *testing.T) {
	type M struct {
		m []byte
		v uint8
	}
	fs := []M{
		{[]byte{128}, 128},
		{[]byte{1, 0, 0, 0, 0, 0, 0, 0, 128}, 128},
	}
	for _, f := range fs {
		if Be.ToUint8(f.m) != f.v {
			t.Error()
		}
	}
}

func Test_littleEndian_ToBytes(t *testing.T) {
	b := Le.ToBytes(128, 3)
	m := []byte{128, 0, 0}
	if bytes.Compare(b, m) != 0 {
		t.Error()
	}
}

func Test_littleEndian_ToUint(t *testing.T) {
	type M struct {
		m []byte
		v uint
	}
	fs := []M{
		{[]byte{128, 0, 0}, 128},
		{[]byte{128, 0, 0, 0, 0, 0, 0, 0, 1}, 128},
	}
	for _, f := range fs {
		if Le.ToUint(f.m) != f.v {
			t.Error()
		}
	}
}

func Test_littleEndian_ToUint16(t *testing.T) {
	type M struct {
		m []byte
		v uint16
	}
	fs := []M{
		{[]byte{128}, 128},
		{[]byte{128, 0, 0, 0, 0, 0, 0, 0, 1}, 128},
	}
	for _, f := range fs {
		if Le.ToUint16(f.m) != f.v {
			t.Error()
		}
	}
}

func Test_littleEndian_ToUint32(t *testing.T) {
	type M struct {
		m []byte
		v uint32
	}
	fs := []M{
		{[]byte{128}, 128},
		{[]byte{128, 0, 0, 0, 0, 0, 0, 0, 1}, 128},
	}
	for _, f := range fs {
		if Le.ToUint32(f.m) != f.v {
			t.Error()
		}
	}
}

func Test_littleEndian_ToUint64(t *testing.T) {
	type M struct {
		m []byte
		v uint64
	}
	fs := []M{
		{[]byte{128}, 128},
		{[]byte{128, 0, 0, 0, 0, 0, 0, 0, 1}, 128},
	}
	for _, f := range fs {
		if Le.ToUint64(f.m) != f.v {
			t.Error()
		}
	}
}

func Test_littleEndian_ToUint8(t *testing.T) {
	type M struct {
		m []byte
		v uint8
	}
	fs := []M{
		{[]byte{128}, 128},
		{[]byte{128, 0, 0, 0, 0, 0, 0, 0, 1}, 128},
	}
	for _, f := range fs {
		if Le.ToUint8(f.m) != f.v {
			t.Error()
		}
	}
}
