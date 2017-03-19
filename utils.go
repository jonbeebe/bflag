package bflag

import (
	"regexp"
	"strings"
)

func parseOption(arg string) (string, string) {
	var nameVal string

	if arg[1] == '-' {
		nameVal = arg[2:]
	} else {
		nameVal = arg[1:]
	}

	equalIdx := strings.Index(nameVal, "=")
	name := nameVal
	value := ""

	if equalIdx != -1 {
		name = nameVal[:equalIdx]
		if equalIdx != len(nameVal)-1 {
			value = nameVal[equalIdx+1:]
		}
	}

	return name, value
}

func isValidOption(name string) bool {
	patterns := []string{
		"\\A-[[:alpha:]]\\z",                                 // -x
		"\\A-[[:alpha:]]=.",                                  // -x=foo
		"\\A--[[:alpha:]]([A-Za-z0-9][A-Za-z0-9_-]*){1,}\\z", // --foo-bar
		"\\A--[[:alpha:]]([A-Za-z0-9][A-Za-z0-9_-]*){1,}=.",  // --foo-bar=baz
	}
	regPattern := strings.Join(patterns, "|")
	match, _ := regexp.MatchString(regPattern, name)
	return match
}

func isValidFlag(name string) bool {
	patterns := []string{
		"\\A-[[:alpha:]]\\z",                                           // -x
		"\\A-[[:alpha:]]=(true|false)",                                 // -x=true|false
		"\\A--[[:alpha:]]([A-Za-z0-9][A-Za-z0-9_-]*){1,}\\z",           // --foo-bar
		"\\A--[[:alpha:]]([A-Za-z0-9][A-Za-z0-9_-]*){1,}=(true|false)", // --foo-bar=true|false
	}
	regPattern := strings.Join(patterns, "|")
	match, _ := regexp.MatchString(regPattern, name)
	return match
}
