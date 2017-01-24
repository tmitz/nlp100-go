package cipher

import (
	"bytes"
	"fmt"
)

func CipherPrint(s string) {
	fmt.Println(Cipher(s))
}

func Cipher(s string) string {
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		if 97 <= s[i] && s[i] <= 122 {
			buf.WriteByte(219 - s[i])
		} else {
			buf.WriteByte(s[i])
		}
	}

	return buf.String()
}

func Decipher(s string) string {
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		if 97 <= s[i] && s[i] <= 122 {
			buf.WriteByte(219 - s[i])
		} else {
			buf.WriteByte(s[i])
		}
	}
	fmt.Println(buf.String())
	return buf.String()
}
