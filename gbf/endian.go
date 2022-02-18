package gbf

type littleEndian struct {
	Error error
}

type bigEndian struct {
	Error error
}

var Le = littleEndian{}
var Be = bigEndian{}

func (b *bigEndian) ToBytes(num, length int) (bf []byte) {
	bf, b.Error = intToBigBytes(num, length)
	return bf
}

func (b *bigEndian) ToUint8(bt []byte) uint8 {
	var num uint64
	num, b.Error = bBytesToUint64(bt)
	return uint8(num)
}
func (b *bigEndian) ToUint16(bt []byte) uint16 {
	var num uint64
	num, b.Error = bBytesToUint64(bt)
	return uint16(num)
}
func (b *bigEndian) ToUint32(bt []byte) uint32 {
	var num uint64
	num, b.Error = bBytesToUint64(bt)
	return uint32(num)
}
func (b *bigEndian) ToUint64(bt []byte) (num uint64) {
	num, b.Error = bBytesToUint64(bt)
	return num
}
func (b *bigEndian) ToUint(bt []byte) uint {
	var num uint64
	num, b.Error = bBytesToUint64(bt)
	return uint(num)
}

func (b *littleEndian) ToBytes(num, length int) (bf []byte) {
	bf, b.Error = intToLittleBytes(num, length)
	return bf
}

func (b *littleEndian) ToUint8(bt []byte) uint8 {
	var num uint64
	num, b.Error = lBytesToUint64(bt)
	return uint8(num)
}
func (b *littleEndian) ToUint16(bt []byte) uint16 {
	var num uint64
	num, b.Error = lBytesToUint64(bt)
	return uint16(num)
}
func (b *littleEndian) ToUint32(bt []byte) uint32 {
	var num uint64
	num, b.Error = lBytesToUint64(bt)
	return uint32(num)
}
func (b *littleEndian) ToUint64(bt []byte) (num uint64) {
	num, b.Error = lBytesToUint64(bt)
	return num
}
func (b *littleEndian) ToUint(bt []byte) uint {
	var num uint64
	num, b.Error = lBytesToUint64(bt)
	return uint(num)
}
