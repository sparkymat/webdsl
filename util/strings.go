package util

import (
	"regexp"
	"strings"
)

func CamelToUnderscore(camel string) string {
	regex, _ := regexp.Compile("[A-Z]")
	stage1 := string(regex.ReplaceAll([]byte(camel), []byte("_$0")))
	regex, _ = regexp.Compile("^_+")
	stage2 := regex.ReplaceAllString(stage1, "")
	return strings.ToLower(stage2)
}
