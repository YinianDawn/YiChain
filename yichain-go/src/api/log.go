package api

import (
	"fmt"
	"reflect"
	"strings"
)

type Logger interface {
	Trace(content string, replaces ...interface{})
	Info(content string, replaces ...interface{})
	Warning(content string, replaces ...interface{})
	Error(content string, replaces ...interface{})
}

func Parse(content string, replaces []interface{}) string {
	for _, replace := range replaces {
		var r string
		if v, ok := replace.(fmt.Stringer); ok {
			r = v.String()
		} else {
			switch t := replace.(type) {
			case string:
				r = replace.(string)
			case int:
				r = fmt.Sprintf("%d", replace)
			case float64:
				r = fmt.Sprintf("%f", replace)
			default:
				_ = t
				panic(fmt.Sprintf("can not transfer to string: " + reflect.TypeOf(replace).String()))
			}
		}
		content = strings.Replace(content, "{}", r, 1)
	}
	return content
}
