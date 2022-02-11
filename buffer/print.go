package buffer

import "fmt"

func Println(buf []byte) {
	printer := func(b []byte) { fmt.Print(formatStingSlice(b)) }
	n := len(buf)
	if n > 0 {
		for i := 0; i < n; i += 8 {

			l := i + 8
			can := l
			if can >= n {
				can = n
			}
			if len(buf[i:can]) <= 0 {
				continue
			}
			if l%16 == 8 && i != 0 {
				fmt.Println("")
			}
			if l%16 == 8 {
				fmt.Printf("%c[1;40;32m%04d:  %c[0m", 0x1B, l/16+1, 0x1B)
			}
			if l >= n {
				l = n
			}
			printer(buf[i:can])
			if l%16 == 8 {
				fmt.Print("  ")
			}
		}

	} else {
		printer(buf)
	}

	fmt.Printf("\n%c[1;40;32m报文总长度：%d  %c[0m\n", 0x1B, n, 0x1B)

}
func formatStingSlice(data []byte) []string {
	s := fmt.Sprintf("%x", data)
	n := len(s)
	bufferString := make([]string, 0)
	for i := 0; i <= n-2; i = i + 2 {
		bufferString = append(bufferString, s[i:i+2])
	}
	return bufferString
}
