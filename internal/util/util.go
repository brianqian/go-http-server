package util

import (
	"fmt"
	"strconv"
)

func WrapInSingleQuotes(v interface{}) string {
	switch v := v.(type) {
	case string:
		return fmt.Sprintf("'%v'", v)
	case int:
		return "'" + strconv.Itoa(v) + "'"
	default:
		return fmt.Sprintf("'%v'", v)
	}
}
