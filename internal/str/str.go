package str

import "strings"

func IsEmptyOrWhitespace(str string) bool {
	Trimmed := strings.TrimSpace(str)

	return len(Trimmed) == 0
}