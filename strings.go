package slackbot

import "strings"

func addBrackets(s string) string {
	if s == "" {
		return s
	} else {
		return "(" + s + ")"
	}
}

func boldSubstring(s, target string) string {
	if s == "" || target == "" {
		return s
	} else {
		return strings.Replace(s, target, "*"+target+"*", 1)
	}
}

func selectString(condition bool, trueValue, falseValue string) string {
	if condition {
		return trueValue
	} else {
		return falseValue
	}
}

func containsString(strs []string, target string) bool {
	if target == "" {
		return false
	}
	for _, str := range strs {
		if str == target {
			return true
		}
	}
	return false
}
