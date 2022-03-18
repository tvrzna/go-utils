package args

import (
	"strings"
)

const (
	delimiter = "="
)

// ArgumentHandler defines function used for ParseArgs method.
type ArgumentHandler func(arg, nextArg string)

// Parses all arguments, runs for each the handler method with argument value and its following argument, if it's not "-" argument.
func ParseArgs(args []string, handler ArgumentHandler) {
	for i, arg := range args {
		nextArg := ""
		if len(args) > i+1 {
			val := strings.TrimSpace(args[i+1])
			if !strings.HasPrefix(val, "-") {
				nextArg = val
			}
		}
		if strings.Contains(arg, delimiter) {
			nextArg = arg[strings.Index(arg, delimiter)+1 : len(arg)]
			arg = arg[0:strings.Index(arg, delimiter)]
			if (strings.HasPrefix(nextArg, "'") && strings.HasSuffix(nextArg, "'")) || (strings.HasPrefix(nextArg, "\"") && strings.HasSuffix(nextArg, "\"")) {
				nextArg = nextArg[1 : len(nextArg)-1]
			}
		}
		handler(arg, nextArg)
	}
}

// Checks arguments, if contains any of specifiec argument.
func ContainsArg(args []string, arg ...string) bool {
	for _, v := range args {
		for _, a := range arg {
			if v == a || (strings.Contains(v, delimiter) && v[0:strings.Index(v, delimiter)] == a) {
				return true
			}
		}
	}
	return false
}
