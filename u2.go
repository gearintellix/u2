package u2

import (
	"fmt"
	"strings"
)

// Binding to apply u2 with binders
func Binding(q string, binders map[string]string) string {
	for k, v := range binders {
		q = strings.ReplaceAll(q, fmt.Sprintf("__%s__", k), v)
	}
	return q
}
