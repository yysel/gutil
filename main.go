package main

import (
	"github.com/yysel/gutil/buffer"
)

func main() {
	b := buffer.Buffer{}
	buf := []byte{1, 2, 3, 4}
	//fmt.Println(byte(uint8(2222)))
	//b.Println()
	b.WriteBytes(buf).
		//Println().
		ReWrite(3, []byte{0, 0, 0, 1}).
		Println()
	//a := [5]byte{}
	//b.Write(a[:])
	//b.Println()
	//b.Write(buffer.Int2LittleBytes(2888, 4))
	//b.Println()
	//a := make([]byte, 4)
	//binary.LittleEndian.PutUint32(a, 2888)

}
