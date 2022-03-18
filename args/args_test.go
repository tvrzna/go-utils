package args

import "testing"

func TestParseArgs(t *testing.T) {
	args := []string{"-a", "-b", "hello there", "-c", "--d='hello here'"}

	handledA := false
	dataA := ""
	handledB := false
	dataB := ""
	handledC := false
	dataC := ""
	handledD := false
	dataD := ""

	ParseArgs(args, func(arg, nextArg string) {
		switch arg {
		case "-a":
			handledA = true
			dataA = nextArg
		case "-b":
			handledB = true
			dataB = nextArg
		case "-c":
			handledC = true
			dataC = nextArg
		case "--d":
			handledD = true
			dataD = nextArg
		}
	})

	if !handledA || dataA != "" {
		t.Error("TestParseArgs: argument A was not handled or had unexpected next argument")
	}

	if !handledB || dataB == "" {
		t.Error("TestParseArgs: argument B was not handled or had unexpected content of next argument")
	}

	if !handledC || dataC != "" {
		t.Error("TestParseArgs: argument C was not handled or had unexpected next argument")
	}

	if !handledD || dataD == "" {
		t.Log(dataD)
		t.Error("TestParseArgs: argument D was not handled or had unexpected next argument")
	}
}

func TestContainsArg(t *testing.T) {
	args := []string{"-a", "-b", "--d='hello there'"}

	if !ContainsArg(args, "--a", "-a", "--b", "-d") {
		t.Error("TestContainsArg: expected argument was not found")
	}

	if ContainsArg(args, "--b") {
		t.Error("TestContainsArg: unexpected argument was found")
	}

	if !ContainsArg(args, "--d") {
		t.Error("TestContainsArg: expected argument was not found")
	}
}
