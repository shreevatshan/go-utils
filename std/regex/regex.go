package regex

import (
	"regexp"
)

const (
	containsNumberPattern = `.*\d.*`
)

var (
	containsNumberRegex *regexp.Regexp
)

func init() {

	containsNumberRegex, _ = regexp.Compile(containsNumberPattern)

}

func GetContainsNumberRegex() *regexp.Regexp {
	return containsNumberRegex
}
