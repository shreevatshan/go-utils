package ip

import "regexp"

const (
	ipv4Pattern = "(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])"
	ipv6Pattern = "((([0-9a-fA-F]){1,4})\\:){7}([0-9a-fA-F]){1,4}"
)

var (
	internetProtocolRegexes map[string]*regexp.Regexp
)

func init() {

	//IPADDRESS REGEX
	internetProtocolRegexes = make(map[string]*regexp.Regexp)

	internetProtocolRegexes["ipv4"], _ = regexp.Compile(ipv4Pattern)

	internetProtocolRegexes["ipv6"], _ = regexp.Compile(ipv6Pattern)

}

func GetInternetProtocolRegexes() map[string]*regexp.Regexp {
	return internetProtocolRegexes
}

func GetIPv4Regex() *regexp.Regexp {
	return internetProtocolRegexes["ipv4"]
}

func GetIPv6Regex() *regexp.Regexp {
	return internetProtocolRegexes["ipv6"]
}

func IsValidIPAddress(addr string) bool {

	for _, regex := range internetProtocolRegexes {
		if regex.MatchString(addr) {
			return true
		}
	}

	return false
}
