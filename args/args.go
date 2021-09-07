package args

import "strings"

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
		handler(arg, nextArg)
	}
}

// Checks arguments, if contains any of specifiec argument.
func ContainsArg(args []string, arg ...string) bool {
	for _, v := range args {
		for _, a := range arg {
			if v == a {
				return true
			}
		}
	}
	return false
}
