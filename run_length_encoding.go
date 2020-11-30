package encode

import (
	"strconv"
	"strings"
)

func RunLengthEncode(input string) string {
	var sb strings.Builder
	if len(input) > 0 {
		currentByte := input[0]
		n := 1

		for i := 1; i < len(input); i++ {
			if input[i] == currentByte {
				n++
			} else {
				if n > 1 {
					sb.WriteString(strconv.Itoa(n))
				}
				sb.WriteByte(currentByte)
				currentByte = input[i]
				n = 1
			}
		}

		if n > 1 {
			sb.WriteString(strconv.Itoa(n))
		}
		sb.WriteByte(currentByte)
	}
	return sb.String()
}

func RunLengthDecode(input string) string {
	var sb strings.Builder
	n := 0
	for i := 0; i < len(input); i++ {
		if input[i] >= '0' && input[i] <= '9' {
			n *= 10
			n += int(input[i] - '0')
		} else {
			if n > 0 {
				sb.WriteString(strings.Repeat(string(input[i]), n))
				n = 0
			} else {
				sb.WriteByte(input[i])
			}
		}
	}
	return sb.String()
}
