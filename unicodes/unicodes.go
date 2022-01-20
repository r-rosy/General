package unicodes

import (
	_ "fmt"
	"strconv"
	"strings"
)

func zhToUnicode(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}
func UnicodeDecodeToStrings(textUnquoted string) string {
	v, _ := zhToUnicode([]byte(textUnquoted))
	return string(v)
}
