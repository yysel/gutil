package gbf

import (
	"bytes"
	"testing"
)

func TestBuffer_InsertWrite(t *testing.T) {
	b := New(BigEndian).
		WriteManyByte(1, 1, 1, 1).
		InsertWrite(1, []byte{0, 0})
	if bytes.Compare(b.Bytes(), []byte{1, 0, 0, 1, 1, 1}) != 0 {
		t.Error()
	}
	b.InsertWrite(0, []byte{255, 255})
	if bytes.Compare(b.Bytes(), []byte{255, 255, 1, 0, 0, 1, 1, 1}) != 0 {
		t.Error()
	}
	b.InsertWrite(b.Len(), []byte{255, 255})
	if bytes.Compare(b.Bytes(), []byte{255, 255, 1, 0, 0, 1, 1, 1, 255, 255}) != 0 {
		t.Error()
	}
}

func TestBuffer_ReWrite(t *testing.T) {
	b := New(BigEndian).
		WriteManyByte(1, 1, 1, 1).
		ReWrite(1, []byte{0, 0})
	//测试正常重写
	if bytes.Compare(b.Bytes(), []byte{1, 0, 0, 1}) != 0 {
		t.Error()
	}
	//测试超过本长度重写
	b.ReWrite(1, []byte{238, 62, 33})
	if bytes.Compare(b.Bytes(), []byte{1, 238, 62, 33}) != 0 {
		t.Error()
	}
	//测试超长度重写
	b.ReWrite(1, []byte{8, 9, 10, 11, 12, 13})
	if bytes.Compare(b.Bytes(), []byte{1, 8, 9, 10, 11, 12, 13}) != 0 {
		t.Error()
	}
	//测试0开始写
	b.ReWrite(0, []byte{88, 99})
	if bytes.Compare(b.Bytes(), []byte{88, 99, 9, 10, 11, 12, 13}) != 0 {
		t.Error()
	}
	//测试末尾开始写
	b.ReWrite(b.Len(), []byte{3, 3})
	if bytes.Compare(b.Bytes(), []byte{88, 99, 9, 10, 11, 12, 13, 3, 3}) != 0 {
		t.Error()
	}
	//测试末尾开始写
	b.ReWrite(b.Len()+20, []byte{3, 3})
	if bytes.Compare(b.Bytes(), []byte{88, 99, 9, 10, 11, 12, 13, 3, 3, 3, 3}) != 0 {
		t.Error()
	}
}

func TestBuffer_WriteInt(t *testing.T) {
	b := New(BigEndian).
		WriteInt(0)
	if bytes.Compare(b.Bytes(), []byte{0}) != 0 {
		t.Error()
	}
	b.WriteInt(255)
	if bytes.Compare(b.Bytes(), []byte{0, 255}) != 0 {
		t.Error()
	}
	b.WriteInt(256)
	if bytes.Compare(b.Bytes(), []byte{0, 255, 1, 0}) != 0 {
		t.Error()
	}
	b.WriteInt(uint16(256))
	if bytes.Compare(b.Bytes(), []byte{0, 255, 1, 0, 1, 0}) != 0 {
		t.Error()
	}
	b.WriteInt(uint32(256))
	if bytes.Compare(b.Bytes(), []byte{0, 255, 1, 0, 1, 0, 0, 0, 1, 0}) != 0 {
		t.Error()
	}
	b.WriteInt(uint64(256))
	if bytes.Compare(b.Bytes(), []byte{0, 255, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0}) != 0 {
		t.Error()
	}
	b.WriteInt(uint8(11))
	if bytes.Compare(b.Bytes(), []byte{0, 255, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 11}) != 0 {
		t.Error()
	}

}

func TestBuffer_WriteBytes(t *testing.T) {
	b := New(BigEndian).
		WriteBytes([]byte{22, 33})
	if bytes.Compare(b.Bytes(), []byte{22, 33}) != 0 {
		t.Error()
	}
	b = New(LittleEndian).
		WriteBytes([]byte{0, 1})
	if bytes.Compare(b.Bytes(), []byte{0, 1}) != 0 {
		t.Error()
	}
	b.WriteBytes([]byte{99, 63})
	if bytes.Compare(b.Bytes(), []byte{0, 1, 99, 63}) != 0 {
		t.Error()
	}
}

func TestBuffer_WriteLength(t *testing.T) {
	b := New(BigEndian).
		WriteBytes([]byte{22, 33}).WriteLength(0)
	if bytes.Compare(b.Bytes(), []byte{22, 33}) != 0 {
		t.Error()
	}
	b.WriteLength(1)
	if bytes.Compare(b.Bytes(), []byte{2, 22, 33}) != 0 {
		b.Println()
		t.Error()
	}
	b.WriteLength(2)
	if bytes.Compare(b.Bytes(), []byte{0, 3, 2, 22, 33}) != 0 {
		b.Println()
		t.Error()
	}
	b.WriteLength(3)
	if bytes.Compare(b.Bytes(), []byte{0, 0, 5, 0, 3, 2, 22, 33}) != 0 {
		b.Println()
		t.Error()
	}

}

func TestBuffer_SetOrder(t *testing.T) {
	b := New(LittleEndian)
	b.WriteInt(256)
	if bytes.Compare(b.Bytes(), []byte{0, 1}) != 0 {
		b.Println()
		t.Error()
	}
	b.SetOrder(BigEndian).WriteInt(256)
	if bytes.Compare(b.Bytes(), []byte{0, 1, 1, 0}) != 0 {
		b.Println()
		t.Error()
	}
}

func TestBuffer_WriteIntFixedLength(t *testing.T) {
	b := New(BigEndian)
	b.WriteIntFixedLength(256, 2)
	if bytes.Compare(b.Bytes(), []byte{1, 0}) != 0 {
		t.Error()
	}
	b.WriteIntFixedLength(255, 2)
	if bytes.Compare(b.Bytes(), []byte{1, 0, 0, 255}) != 0 {
		t.Error()
	}
	b.WriteIntFixedLength(33, 1)
	if bytes.Compare(b.Bytes(), []byte{1, 0, 0, 255, 33}) != 0 {
		t.Error()
	}
	b.WriteIntFixedLength(33, 3)
	if bytes.Compare(b.Bytes(), []byte{1, 0, 0, 255, 33, 0, 0, 33}) != 0 {
		t.Error()
	}
	//限制长度比原值长度小的时候，原值会被截断
	b.WriteIntFixedLength(256, 1)
	if bytes.Compare(b.Bytes(), []byte{1, 0, 0, 255, 33, 0, 0, 33, 0}) != 0 {
		b.Println()
		t.Error()
	}

}

func TestBuffer_WriteMany(t *testing.T) {
	b := New(BigEndian).WriteMany([]byte{0, 1, 1}, []byte{1, 3, 4, 9}, []byte{2, 5, 6})
	if bytes.Compare(b.Bytes(), []byte{0, 1, 1, 1, 3, 4, 9, 2, 5, 6}) != 0 {
		t.Error()
	}
	b.WriteMany([]byte{0, 1, 1})
	if bytes.Compare(b.Bytes(), []byte{0, 1, 1, 1, 3, 4, 9, 2, 5, 6, 0, 1, 1}) != 0 {
		t.Error()
	}
}

func TestBuffer_WriteManyByte(t *testing.T) {

	b := New(BigEndian).WriteManyByte(1, 2, 3, 4)
	if bytes.Compare(b.Bytes(), []byte{1, 2, 3, 4}) != 0 {
		t.Error()
	}
	b.WriteManyByte(0)
	if bytes.Compare(b.Bytes(), []byte{1, 2, 3, 4, 0}) != 0 {
		t.Error()
	}
}

func TestBuffer_WritePlaceholder(t *testing.T) {
	b := New(BigEndian).WritePlaceholder(1)
	if bytes.Compare(b.Bytes(), []byte{0}) != 0 {
		t.Error()
	}
	b.WritePlaceholder(0)
	if bytes.Compare(b.Bytes(), []byte{0}) != 0 {
		t.Error()
	}
	b.WritePlaceholder(5)
	if bytes.Compare(b.Bytes(), []byte{0, 0, 0, 0, 0, 0}) != 0 {
		t.Error()
	}
}

func TestBuffer_WriteString(t *testing.T) {

	b := New(BigEndian).WriteString("123abc")
	if bytes.Compare(b.Bytes(), []byte{49, 50, 51, 97, 98, 99}) != 0 {
		t.Error()
	}
	b.WriteString("ABC")
	if bytes.Compare(b.Bytes(), []byte{49, 50, 51, 97, 98, 99, 65, 66, 67}) != 0 {
		t.Error()
	}
}

func TestBuffer_WriteStringWithLen(t *testing.T) {
	b := New(BigEndian).WriteStringWithLen("123abc", 1)
	if bytes.Compare(b.Bytes(), []byte{6, 49, 50, 51, 97, 98, 99}) != 0 {
		t.Error()
	}
	b.WriteStringWithLen("ABC", 3)
	if bytes.Compare(b.Bytes(), []byte{6, 49, 50, 51, 97, 98, 99, 0, 0, 3, 65, 66, 67}) != 0 {
		t.Error()
	}
}

func TestBuffer_GetError(t *testing.T) {
	b := New(BigEndian).WriteInt("123abc")
	if bytes.Compare(b.Bytes(), []byte{}) != 0 {
		t.Error()
	}
	if len(b.GetError()) == 0 {
		t.Error()
	}
}
func TestBuffer_HasError(t *testing.T) {
	b := New(BigEndian).WriteInt("123abc")
	if bytes.Compare(b.Bytes(), []byte{}) != 0 {
		t.Error()
	}
	if !b.HasError() {
		t.Error()
	}
}

func TestBuffer_WriteTop(t *testing.T) {
	b := New(BigEndian).WriteManyByte(1, 2, 3, 4)
	b.WriteTop([]byte{1, 1})
	if bytes.Compare(b.Bytes(), []byte{1, 1, 1, 2, 3, 4}) != 0 {
		t.Error()
	}
}
