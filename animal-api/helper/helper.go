package helper

import (
	"strings"

	"github.com/mitchellh/mapstructure"
)

func toCamelCase(input string) string {
	isToUpper := false
	var result string
	for i, v := range input {
		if i == 0 {
			result += strings.ToLower(string(v))
		} else if v == '_' {
			isToUpper = true
		} else {
			if isToUpper {
				result += strings.ToUpper(string(v))
				isToUpper = false
			} else {
				result += string(v)
			}
		}
	}
	return result
}

// ApplyChanges function to decode map into struct
func ApplyChanges(changes map[string]interface{}, to interface{}) error {
	camelCaseKeys := make(map[string]interface{})
	for k, v := range changes {
		camelCaseKeys[toCamelCase(k)] = v
	}

	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ErrorUnused: true,
		TagName:     "json",
		Result:      to,
		ZeroFields:  true,
	})

	if err != nil {
		return err
	}

	return dec.Decode(camelCaseKeys)
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
