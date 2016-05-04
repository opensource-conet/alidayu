package alidayu

import (
	"regexp"
	"strings"
)

var reg = regexp.MustCompile("^[0-9]*$")

// SplitCode 将数字用,隔开
//--{66666} to {6,6,6,6,6}
// 适用于文本转语音输出时验证码的输出问题
func SplitCode(code string) string {
	if reg.MatchString(code) {
		return strings.Join(strings.Split(code, ""), ",")
	}

	return ""
}
