package alidayu

import (
	"errors"
	"regexp"
	"strings"
)

// SpritCode 将数字用,隔开
//--{66666} to {6,6,6,6,6}
// 适用于文本转语音输出时验证码的输出问题
func SpritCode(code string) (string, error) {
	match, err := regexp.MatchString("^[0-9]*$", code)
	if err != nil {
		return "", err
	}
	if match == true {
		return strings.Join(strings.Split(code, ""), ","), nil
	}
	return "", errors.New("regexp is fail")
}
