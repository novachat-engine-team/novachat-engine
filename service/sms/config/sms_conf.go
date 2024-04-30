package config

// 短信模板设置
type SmsTemplate struct {
	Title string // 标题
	Code  string // 模板
	InternationalCode string //国际短信
}

// 短信签名配置
type SubmailSignConfig struct {
	SubmailAppID             string //国内短信应用ID
	SubmailAppKey            string //国内短信应用签名
	SubmailHost              string //国内短信发送接口
	InternationalAppID       string //国际短信应用ID
	InternationalAppKey      string //国际短信应用签名
	InternationalSubmailHost string //国际短信发送接口
}

type SubmailConfig struct {
	SmsTemplate SmsTemplate
	SmsConfig SubmailSignConfig
}

type HuaweiTemplate struct {
	AppID             string //国内短信应用ID
	AppKey            string //国内短信应用签名
	AppSecret		  string
	TemplateID		  string //模板ID
	Host              string //国内短信发送接口
	InternationalHost string //国际短信发送接口
	Signature		  string //国内签名
}

type HuaweiConfig struct {
	SmsConfig HuaweiTemplate
}

type YunPianTemplate struct {
	AppKey  string
	Host 	 			string
	InternationalHost 	string //国际短信发送接口
	Template			string
	InternationalTemplate string
}

type YunPianConfig struct {
	SmsConfig YunPianTemplate
}