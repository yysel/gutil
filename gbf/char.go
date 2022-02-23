package gbf

import (
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// Utf16BytesToUtf8 utf16 Bytes 转为 utf8 bytes
func Utf16BytesToUtf8(b []byte, order Endianness) []byte {
	e := unicode.BigEndian
	if order {
		e = unicode.LittleEndian
	}
	bs_UTF8LE, _, _ := transform.Bytes(unicode.UTF16(e, unicode.IgnoreBOM).NewDecoder(), b)
	return bs_UTF8LE
}

// Utf16BytesToUtf8String utf16 Bytes 转为 utf8 string
func Utf16BytesToUtf8String(b []byte, order Endianness) string {
	e := unicode.BigEndian
	if order {
		e = unicode.LittleEndian
	}
	bs_UTF8LE, _, _ := transform.Bytes(unicode.UTF16(e, unicode.IgnoreBOM).NewDecoder(), b)
	return string(bs_UTF8LE)
}

// Utf8StingToUtf16Bytes utf8 string 转 utf16 bytes
func Utf8StingToUtf16Bytes(s string, order Endianness) []byte {
	e := unicode.BigEndian
	if order {
		e = unicode.LittleEndian
	}
	bs_UTF16LE, _, _ := transform.Bytes(unicode.UTF16(e, unicode.IgnoreBOM).NewEncoder(), []byte(s))
	return bs_UTF16LE
}

// Utf8ToUtf16Bytes utf8 bytes 转 utf16 bytes
func Utf8ToUtf16Bytes(b []byte, order Endianness) []byte {
	e := unicode.BigEndian
	if order {
		e = unicode.LittleEndian
	}
	bs_UTF16LE, _, _ := transform.Bytes(unicode.UTF16(e, unicode.IgnoreBOM).NewEncoder(), b)
	return bs_UTF16LE
}
