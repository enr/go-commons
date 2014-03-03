package lang

import (
	"fmt"
)

// Returns true if the given string is present in a strings slice.
// Case sensitive.
func SliceContainsString(list []string, a string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Copyright Am Laher
// https://github.com/laher/goxc/blob/master/typeutils/mapstringinterfaceutils.go
// Coerce a JSON array to slice of strings.
func JsonArrayToStringSlice(v interface{}, k string) ([]string, error) {
	ret := []string{}
	switch typedV := v.(type) {
	case []interface{}:
		for _, i := range typedV {
			ret = append(ret, i.(string))
		}
		return ret, nil
	}
	return ret, fmt.Errorf("%s should be a `json array`, got a %T", k, v)
}
