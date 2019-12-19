package u2

import (
	"reflect"
	"strings"
)

// index to find string with offset
func index(s string, value string, offset int) int {
	if offset > 0 {
		s = subStr(s, offset, 0)
	}

	if offset < 0 {
		offset = 0
	}

	i := strings.Index(s, value)
	if i >= 0 {
		return i + offset
	}
	return i
}

// subStr function
func subStr(s string, f int, l int) string {
	if l == 0 {
		l = len(s) - f
	} else if l < 0 {
		l = len(s) + l
	}

	return s[f : f+l]
}

// matrixDynamic to create matrix array dynamicly
func matrixDynamic(a []interface{}, b []interface{}, fn func(interface{}, interface{}) interface{}) []interface{} {
	res := []interface{}{}
	for _, v := range a {
		for _, v2 := range b {
			c := fn(v, v2)
			res = append(res, c)
		}
	}
	return res
}

// matrixStr to create matrix array string
func matrixStr(a []string, b []string) []string {
	ai := []interface{}{}
	for _, v := range a {
		ai = append(ai, v)
	}

	bi := []interface{}{}
	for _, v := range b {
		bi = append(bi, v)
	}

	resi := matrixDynamic(ai, bi, func(ax interface{}, bx interface{}) interface{} {
		return ax.(string) + bx.(string)
	})

	res := []string{}
	for _, v := range resi {
		res = append(res, v.(string))
	}

	return res
}

// trim whitespace
func trim(value string) string {
	ws := []string{
		"\n", "\t", " ",
		"\t", "\n", " ",
	}
	ws = matrixStr(ws, ws)
	for _, v := range ws {
		value = strings.Trim(value, v)
	}
	return value
}

// isArrExists to check existance of value in array
func isArrExists(value interface{}, array interface{}) (exist bool) {
	exist = false
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(value, s.Index(i).Interface()) == true {
				exist = true
				return exist
			}
		}
	}

	return exist
}
