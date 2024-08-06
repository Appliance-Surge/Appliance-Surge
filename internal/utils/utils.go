package utils

import (
	"errors"
	"text/template"
)

// Dict is a function to create a map from a list of key-value pairs
func Dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("invalid dict call")
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New("dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}

// FuncMap returns a map of functions to be used in templates
func FuncMap() template.FuncMap {
	return template.FuncMap{
		"dict": Dict,
	}
}
