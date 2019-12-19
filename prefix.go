package u2

import (
	"regexp"
	"strings"
)

// ScanPrefix to get all u2 binding with a prefix
func ScanPrefix(q string, prefixs []string) (items map[string][]string) {
	items = make(map[string][]string)

	for _, v := range prefixs {
		items[v] = []string{}
	}

	fx := regexp.MustCompile(`__([^\s]*?)__`)
	binds := fx.FindAll([]byte(q), -1)

	for _, v := range binds {
		vx := subStr(string(v), 2, -4)

		for _, v2 := range prefixs {
			if strings.HasPrefix(vx, v2) {
				items[v2] = append(items[v2], subStr(vx, len(v2), 0))
			}
		}
	}

	return items
}
