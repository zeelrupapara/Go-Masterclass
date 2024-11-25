package reverse

import (
	"strings"
)

func Reverse(s string) string {
	chars := strings.Split(s, "")
	reverse_char := []string{}
	for i := len(chars) - 1; i >= 0; i-- {
		reverse_char = append(reverse_char, chars[i])
	}
	return strings.Join(reverse_char, "")
}
