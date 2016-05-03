// 中遣互联版权所有
// author: huminfo <153336400@qq.com>
// date:   2016-3-21

package alidayu

// Result 用于描述发送短信的结果
type Result struct {
	Success     bool         `json:"success,omitempty"`                       // 成功失败信息
	ResultError *resultError `json:"error_response"`                          // 返回错误
	ResultModel *resultModel `json:"alibaba_aliqin_fc_sms_num_send_response"` // 返回正确
}

// 成功返回的结构体
type resultModel struct {
	Result    *resultSms `json:"result"`
	RequestID string     `json:"request_id"` // 请求ID
}

// 成功返回结构
type resultSms struct {
	ErrCode string `json:"err_code"` // 错误码
	Model   string `json:"model"`    // 返回结果
	Success bool   `json:"success"`  //true表示成功，false表示失败
}

// 返回错误
type resultError struct {
	Code    int    `json:"code"`     // 错误代码
	Msg     string `json:"msg"`      // 错误信息
	SubCode string `json:"sub_code"` // 错误代码
	SubMsg  string `json:"sub_msg"`  // 错误码
}

// 公共参数
type commonModel struct {
	// Method         string `json:"method"`             // 不填 API接口名称 写死在方法里
	// AppKey       string `json:"sms_type"`           // 不填 TOP分配给应用的AppKey
	Session      string `json:"sms_free_sign_name"` // 可选 用户登录授权成功后，TOP颁发给应用的授权信息
	Timestamp    string `json:"sms_param"`          // 必选 时间戳，格式为yyyy-MM-dd HH:mm:ss
	Format       string `json:"rec_num"`            // 可选 响应格式。默认为xml格式，可选值：xml，json
	V            string `json:"extend"`             // 必选 API协议版本，可选值：2.0
	PartnerID    string `json:"sms_type"`           // 可选 合作伙伴身份标识
	TargetAppKey string `json:"sms_free_sign_name"` // 可选 被调用的目标AppKey，仅当被调用的API为第三方ISV提供时有效
	Simplify     bool   `json:"sms_param"`          // 可选 是否采用精简JSON返回格式，仅当format=json时有效，默认值为：false。
	SignMethod   string `json:"rec_num"`            // 必选 签名的摘要算法，可选值为：hmac，md5。
	// Sign           string `json:"extend"`             // 不填 API输入参数签名结果 加密算法 post时自动加密
}

// 短信请求参数
type smsModel struct {
	Extend          string `json:"extend"`             // 可选 公共回传参数
	SmsType         string `json:"sms_type"`           // 必选 类型.normal：短信
	SmsFreeSignName string `json:"sms_free_sign_name"` // 必须 短信签名
	SmsParam        string `json:"sms_param"`          // 可选 短信模板变量，AckNum是变量参数
	RecNum          string `json:"rec_num"`            // 必选 接收号码
	SmsTemplateCode string `json:"sms_template_code"`  // 必选 短信模板CODE
	ExtendCode      string `json:"extend_code"`        // 可选 商家自定义扩展码
	ExtendName      string `json:"extend_name"`        //可选 商家自定义扩展名,例如店铺nick
}

// 文本转语音请求参数
type lecallModel struct {
	Extend        string `json:"extend"`          // 可选 公共回传参数
	TtsParam      string `json:"tts_param"`       //文本转语音（TTS）模板变量，传参规则{"key"："value"}
	CalledNum     string `json:"called_num"`      //被叫号码，支持国内手机号与固话号码
	CalledShowNum string `json:"called_show_num"` //被叫号显，传入的显示号码必须是阿里大鱼“管理中心-号码管理”中申请或购买的号码
	TtsCode       string `json:"tts_code"`        //TTS模板ID，传入的模板必须是在阿里大鱼“管理中心-语音TTS模板管理”中的可用模板
}
