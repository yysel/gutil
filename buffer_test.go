package main

import (
	"github.com/yysel/gutil/buffer"
	"testing"
)

func TestBuffer_Println(t *testing.T) {
	b := buffer.Buffer{}
	a := [5]byte{}
	b.Write(a[:])
	b.Println()
}
