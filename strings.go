package slackbot

import "strings"

func addBrackets(s string) string {
	if s == "" {
		return s
	} else {
		return "(" + s + ")"
	}
}

func encloseString(s, character string) string {
	if s == "" {
		return s
	} else {
		return character + s + character
	}
}

func encloseSubstring(s, target, character string) string {
	if s == "" || target == "" {
		return s
	} else {
		return strings.Replace(s, target, character+target+character, 1)
	}
}

func boldString(s string) string {
	return encloseString(s, "*")
}

func italicString(s string) string {
	return encloseString(s, "_")
}

func boldSubstring(s, target string) string {
	return encloseSubstring(s, target, "*")
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
