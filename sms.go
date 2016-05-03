// 中遣互联版权所有
// author: huminfo <153336400@qq.com>
// date:   2013-3-21

package alidayu

import "encoding/json"

func messasg(cm *commonModel, sm *smsModel) (*Result, error) {
	// 发送数据到接口
	m := make(map[string]string, 11)
	m["app_key"] = Appkey
	m["format"] = cm.Format
	m["method"] = methodSendSMS
	m["sign_method"] = cm.SignMethod
	m["timestamp"] = cm.Timestamp
	m["v"] = cm.V
	m["sms_type"] = sm.SmsType
	m["sms_free_sign_name"] = sm.SmsFreeSignName
	m["rec_num"] = sm.RecNum
	m["sms_template_code"] = sm.SmsTemplateCode
	m["sms_param"] = sm.SmsParam

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

func lecall(cm *commonModel, lm *lecallModel) (*Result, error) {
	m := make(map[string]string, 11)
	m["app_key"] = Appkey
	m["format"] = cm.Format
	m["method"] = methodCallTTS
	m["sign_method"] = cm.SignMethod
	m["timestamp"] = cm.Timestamp
	m["v"] = cm.V
	m["called_num"] = lm.CalledNum
	m["called_show_num"] = lm.CalledShowNum
	m["tts_code"] = lm.TtsCode
	m["tts_param"] = lm.TtsParam

	return responseToResult(m)
}
