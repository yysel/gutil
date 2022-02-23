package gbf

import (
	"bytes"
	"fmt"
)

const (
	BigEndian    bool = true
	LittleEndian bool = false
)

type Buffer struct {
	order bool //true 大端法 false 小端法
	error []error
	bytes.Buffer
}

func New(order bool) *Buffer {
	return &Buffer{
		order: order,
	}
}
func (b *Buffer) Print() *Buffer {
	bufferString := formatStingSlice(b.Bytes())
	fmt.Println(bufferString)
	return b
}

func (b *Buffer) SetOrder(order bool) *Buffer {
	b.order = order
	return b
}

func (b *Buffer) GetError() []error {
	return b.error
}
func (b *Buffer) HasError() bool {
	return len(b.error) != 0
}

// Println 格式化输出字节数组
func (b *Buffer) Println() *Buffer {
	Println(b.Bytes())

	if len(b.error) > 0 {
		fmt.Printf("%c[1;40;31m错误内容：%v  %c[0m\n", 0x1B, b.error, 0x1B)
	}
	return b
}

func (b *Buffer) pushError(e error) {
	if e != nil {
		b.error = append(b.error, e)
	}
}

// WritePlaceholder 当前位置添加指定位的占位字节
func (b *Buffer) WritePlaceholder(num int) *Buffer {
	s := make([]byte, num)
	_, e := b.Write(s)
	if e != nil {
		b.pushError(e)
	}
	return b
}

// WriteInt 写入一个数值整型，并根据整型的大小调整字节长度
func (b *Buffer) WriteInt(num interface{}) *Buffer {
	buf, e := intToBytes(num, b.order)
	b.pushError(e)
	_, e = b.Write(buf)
	b.pushError(e)
	return b
}

// WriteString 写入一个字符串
func (b *Buffer) WriteString(s string) *Buffer {
	_, e := b.Write([]byte(s))
	b.pushError(e)
	return b
}

// WriteStringWithLen 写入一个字符串，并在其前部写入字符串长度字段，l指定长度字段的字节数
func (b *Buffer) WriteStringWithLen(s string, l int) *Buffer {
	b.WriteIntFixedLength(len(s), l).WriteString(s)
	return b
}

// WriteLength 向包头写入当前buffer的长度字段，并指定长度为的字节长度
func (b *Buffer) WriteLength(l int) *Buffer {
	buf := make([]byte, b.Len())
	copy(buf, b.Bytes())
	b.Reset()
	b.WriteIntFixedLength(len(buf), l)
	_, e := b.Write(buf)
	b.pushError(e)
	return b
}

// WriteIntFixedLength 写入一个数值型，并指定存放的字节长度
func (b *Buffer) WriteIntFixedLength(num interface{}, length int) *Buffer {
	buf, e := anyType2Bytes(num, length, b.order)
	b.pushError(e)
	_, e = b.Write(buf)
	b.pushError(e)
	return b
}

// WriteBytes 写入一个字节数组，等价与Write
func (b *Buffer) WriteBytes(p []byte) *Buffer {
	_, e := b.Write(p)
	b.pushError(e)
	return b
}

// ReWrite 从n位开始重写接下来的几个字节
func (b *Buffer) ReWrite(n int, bf []byte) *Buffer {
	if n > b.Len() {
		return b.WriteBytes(bf)
	}
	buf := make([]byte, b.Len())
	copy(buf, b.Bytes())
	start := buf[0:n]
	end := make([]byte, 0)
	if n+len(bf) < b.Len() {
		end = buf[n+len(bf) : b.Len()]
	}
	b.Reset()
	b.WriteMany(start, bf, end)
	return b
}

// InsertWrite 从n字节开始插入一段字节
func (b *Buffer) InsertWrite(n int, bf []byte) *Buffer {
	buf := make([]byte, b.Len())
	copy(buf, b.Bytes())
	start := buf[0:n]
	end := buf[n:b.Len()]
	b.Reset()
	b.WriteMany(start, bf, end)
	return b
}

// WriteMany 一次写入多组字节数组
func (b *Buffer) WriteMany(buf ...[]byte) *Buffer {
	for _, v := range buf {
		_, e := b.Write(v)
		b.pushError(e)
	}
	return b
}

// WriteManyByte 一次写入多个byte
func (b *Buffer) WriteManyByte(buf ...byte) *Buffer {
	for _, v := range buf {
		e := b.WriteByte(v)
		b.pushError(e)
	}
	return b
}

func (b *Buffer) WriteTop(bf []byte) *Buffer {
	return b.InsertWrite(0, bf)
}

func (b *Buffer) WriteBytesWithLength(c []byte, l int) *Buffer {
	return b.WriteIntFixedLength(len(c), l).WriteBytes(c)
}
