package gbf

func (b *Buffer) ReadWordWithLength(n int, data *[]byte) *Buffer {
	l := make([]byte, n)
	b.Read(l)
	var length uint64
	if b.endianness {
		length = Be.ToUint64(l)
	} else {
		length = Le.ToUint64(l)
	}
	d := make([]byte, length)
	b.Read(d)
	data = &d
	return b
}
