package std

import (
	"strconv"
	"strings"
)

func RemoveString(str *string, remove string) {
	*str = strings.ReplaceAll(*str, remove, EmptyString)
}

func ReplaceString(str *string, from string, to string) {
	*str = strings.ReplaceAll(*str, from, to)
}

func RemoveWhiteSpace(str *string) {
	*str = strings.ReplaceAll(*str, Space, EmptyString)
}

// Converts comma separated string to set.
// Set t as 0, if you want to store mixed type values in set.
// Set t as 1, if you want to store string values in set.
// Set t as 2, if you want to store int values in set. If conversion fails, it will be ignored.
func ConvertCommaSeparatedStringToSet(commaSeparatedString string, t ...interface{}) *Set {

	const (
		mixedType  = 0
		stringType = 1
		intType    = 2
	)

	set := InitSet()
	var saveas = 0

	if len(t) > 0 {
		saveas = t[0].(int)
	}

	if commaSeparatedString != EmptyString {
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

func ReturnKeyAndValueFromString(keyvalueString string) (key, value string) {

	key = EmptyString
	value = EmptyString

	lastIndex := strings.LastIndex(keyvalueString, EqualTo)
	if lastIndex != -1 {
		key = keyvalueString[:lastIndex]
		value = keyvalueString[lastIndex+1:]
	}
	return key, value
}

func ConvertNewlineSeparatedStringToKeyValuePairBasedOnEqual(newlineSeparatedString string) map[string]string {

	resultMap := make(map[string]string)

	if newlineSeparatedString != EmptyString {
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

	if newlineSeparatedString != EmptyString {
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
