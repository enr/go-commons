package lang

import (
	"fmt"
)

// Returns true if the given string is present in a strings slice.
func SliceContainsString(list []string, a string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// https://github.com/laher/goxc/blob/master/typeutils/mapstringinterfaceutils.go
// Coerce interface{} to slice of strings.
func InterfaceToStringSlice(v interface{}, k string) ([]string, error) {
	ret := []string{}
	switch typedV := v.(type) {
	case []interface{}:
		for _, i := range typedV {
			ret = append(ret, i.(string))
		}
		return ret, nil
	}
	return ret, fmt.Errorf("%s should be a `[]interface{}`, got a %T", k, v)
}
