// 中遣互联版权所有
// author: huminfo <153336400@qq.com>
// date:   2016-3-22

package alidayu

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

func postAlidayu(m map[string]string) (bool, []byte, error) {
	if Appkey == "" || AppSecret == "" {
		panic("Appkey或者AppSecret为空")
	}

	posturl := httpsURL
	if IsDebug {
		posturl = sandboxHTTPSURL
	}

	body, size := getRequestBody(m)
	client := &http.Client{} // 3600秒后超时
	req, err := http.NewRequest("POST", posturl, body)
	if err != nil {
		return false, nil, err
	}
	defer req.Body.Close()
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.ContentLength = size

	resp, err := client.Do(req)
	if err != nil {
		return false, nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, nil, err
	}
	if !bytes.Contains(data, []byte("success")) {
		return false, data, nil
	}
	return true, data, nil
}

func getRequestBody(m map[string]string) (reader io.Reader, size int64) {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	v := url.Values{}
	signString := AppSecret
	for _, k := range keys {
		v.Set(k, m[k])
		signString += k + m[k]
	}
	signString += AppSecret

	signByte := md5.Sum([]byte(signString))
	sign := strings.ToUpper(fmt.Sprintf("%x", signByte))
	v.Set("sign", sign)

	return ioutil.NopCloser(strings.NewReader(v.Encode())), int64(len(v.Encode()))
}

func responseToResult(m map[string]string) (*Result, error) {
	success, response, err := postAlidayu(m)
	if err != nil {
		return nil, err
	}

	// 解析json转化为结构体
	resultmod := &Result{}
	if err := json.Unmarshal(response, resultmod); err != nil {
		return nil, err
	}
	resultmod.Success = success
	return resultmod, nil
}
