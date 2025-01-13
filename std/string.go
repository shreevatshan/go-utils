package std

import (
	"strconv"
	"strings"

	"github.com/shreevatshan/go-utils/std/maps"
)

func RemoveString(str *string, remove string) {
	*str = strings.ReplaceAll(*str, remove, "")
}

func ReplaceString(str *string, from string, to string) {
	*str = strings.ReplaceAll(*str, from, to)
}

func RemoveWhiteSpace(str *string) {
	*str = strings.ReplaceAll(*str, Space, "")
}

// Converts comma separated string to set.
// Set t as 0, if you want to store mixed type values in set.
// Set t as 1, if you want to store string values in set.
// Set t as 2, if you want to store int values in set. If conversion fails, it will be ignored.
func ConvertCommaSeparatedStringToSet(commaSeparatedString string, t ...interface{}) *maps.Set {

	const (
		mixedType  = 0
		stringType = 1
		intType    = 2
	)

	set := maps.NewSet()
	var saveas = 0

	if len(t) > 0 {
		saveas = t[0].(int)
	}

	if commaSeparatedString != "" {
		values := strings.Split(commaSeparatedString, Comma)
		for i := range values {
			switch saveas {
			case stringType:
				set.Insert(values[i])
			case intType:
				value, err := strconv.Atoi(values[i])
				if err == nil {
					set.Insert(value)
				}
			default:
				value, err := strconv.Atoi(values[i])
				if err == nil {
					set.Insert(value)
				} else {
					set.Insert(values[i])
				}
			}
		}
	}
	return set
}

func ReturnKeyAndValueFromString(keyvalueString string) (string, string) {

	var key string
	var value string

	lastIndex := strings.LastIndex(keyvalueString, EqualTo)
	if lastIndex != -1 {
		key = keyvalueString[:lastIndex]
		value = keyvalueString[lastIndex+1:]
	}
	return key, value
}

func ConvertNewlineSeparatedStringToKeyValuePairBasedOnEqual(newlineSeparatedString string) map[string]string {

	resultMap := make(map[string]string)

	if newlineSeparatedString != "" {
		individualLines := strings.Split(newlineSeparatedString, NewLineAsString)
		for i := range individualLines {
			individualLine := individualLines[i]
			key, value := ReturnKeyAndValueFromString(individualLine)
			RemoveWhiteSpace(&key)
			RemoveWhiteSpace(&value)
			resultMap[key] = value
		}
	}
	return resultMap
}

func ConvertNewlineSeparatedStringToKeyValuePairBasedOnEqualAndComma(newlineSeparatedString string) map[string]string {

	resultMap := make(map[string]string)

	if newlineSeparatedString != "" {
		individualLines := strings.Split(newlineSeparatedString, NewLineAsString)
		for i := range individualLines {
			individualLine := individualLines[i]
			keys, value := ReturnKeyAndValueFromString(individualLine)
			RemoveWhiteSpace(&keys)
			RemoveWhiteSpace(&value)
			individualKeys := strings.Split(keys, Comma)
			for j := range individualKeys {
				key := individualKeys[j]
				resultMap[key] = value
			}
		}
	}
	return resultMap
}

func HasPrefixCaseInsensitive(stringToCheck string, stringToCompare string) bool {
	return strings.HasPrefix(strings.ToLower(stringToCheck), strings.ToLower(stringToCompare))
}

// NonEmptyStrings filters out empty strings from the provided variadic string arguments.
// It returns a slice containing only the non-empty strings.
//
// Parameters:
//
//	str - variadic string arguments to be filtered.
//
// Returns:
//
//	[]string - a slice containing only the non-empty strings from the input.
func NonEmptyStrings(str ...string) []string {
	var result []string
	for i := range str {
		if str[i] != "" {
			result = append(result, str[i])
		}
	}
	return result
}
