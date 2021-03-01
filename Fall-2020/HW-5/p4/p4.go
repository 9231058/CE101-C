/*
 * +===============================================
 * | Author:        Mohammad Fatemi <mr.smf8@gmail.com>
 * |
 * | Creation Date: 04-01-2021
 * |
 * | File Name:     p3.go
 * +===============================================
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func readline() string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func main() {
	line := readline()
	decoded := ""
	for i := 0; i < 26; i++ {
		decoded = Decode(line, i)
		if strings.HasPrefix(decoded, "Salam") {
			fmt.Println(decoded)
			return
		}
	}
}

func rotate(text string, shift int) string {
	shift = (shift%26 + 26) % 26 // [0, 25]
	b := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		t := text[i]
		var a int
		switch {
		case 'a' <= t && t <= 'z':
			a = 'a'
		case 'A' <= t && t <= 'Z':
			a = 'A'
		default:
			b[i] = t
			continue
		}
		b[i] = byte(a + ((int(t)-a)+shift)%26)
	}
	return string(b)
}

// Decode using Caesar Cipher.
func Decode(cipher string, shift int) (plain string) {
	return rotate(cipher, -shift)
}
