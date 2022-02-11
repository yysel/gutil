package main

import (
	"bytes"
	"fmt"
)

func main() {
	//buffer.New(buffer.LittleEndian).
	//	//WriteSting("大声道").
	//	//WriteIntFixedLength(1, 1).
	//	//WriteManyByte(1, 1, 1).
	//	//WriteManyByte(1, 1, 1).
	//	//WriteMany([]byte{0, 2, 3, 4, 's'}).
	//	//WriteStingWithLen("22222", 2).
	//	WriteSting("大声道").
	//	WriteLength(2).
	//	Println()
	fmt.Println(bytes.Compare([]byte{1}, []byte{1}))
	//buffer.New(buffer.LittleEndian).
	//	WriteManyByte(1, 1, 1, 1).
	//	InsertWrite(1, []byte{0, 0}).Println()

}
