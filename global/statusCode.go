package global

const (
	OK = "00000" // 一切OK，正确执行后的返回

	CLIENT_ERROR = "A0001" // 用户端错误，一级宏观错误码

	USER_REGISTER                      = "A0100" // 用户注册错误，二级宏观错误码
	USER_NOT_AGREE_PRIVACY_AGREEMENT   = "A0101" // 用户未同意隐私协议
	RESTRICTED_REGISTRATION_AERA       = "A0102" // 注册国家或地区受限
	USERNAME_CHECK_FAIL                = "A0110" // 用户名校验失败
	USERNAME_EXIST                     = "A0111" // 用户名已存在
	USERNAME_CONTAIN_SENSITIVE_WORDS   = "A0112" // 用户名包含敏感词
	USERNAME_CONTAIN_SPECIAL_CHARACTER = "A0113" // 用户名包含特殊字符
	PASSWORD_CHECK_FAIL                = "A0120" // 密码校验失败
	PASSWORD_LENGTH_ERROR              = "A0121" // 密码长度错误
	PWSSWORD_STRENGTH_ERROR            = "A0122" // 密码强度不够
	INPUT_CHECK_CODE_ERROR             = "A0130" // 校验码输入错误
	INPUT_SMS_CHECK_CODE_ERROR         = "A0131" // 短信校验码输入错误
	INPUT_EMAIL_CHECK_CODE_ERROR       = "A0132" // 邮件校验码输入错误
	INPUT_VOICE_CHECK_CODE_ERROR       = "A0133" // 语音校验码输入错误
	USER_CREDENTIALS_ABNORMAL          = "A0140" // 用户证件异常
	USER_CREDENTIALS_TYPE_NOT_SELECTED = "A0141" // 用户证件类型未选择
	ID_NUMBER_CHECK_ERROR              = "A0142" // 身份证号校验错误
	PASSPORT_CHECK_ERROR               = "A0143" // 护照校验错误
	USER_BASE_INFO_VALIDATION_FAIL     = "A0150" // 用户基本信息校验失败
	PHONE_NUMBER_FORMAT_CHECK_FAIL     = "A0151" // 手机号格式校验失败
	ADDRESS_FORMAT_CHECK_FAIL          = "A0152" // 地址格式校验失败
	EMAIL_FORMAT_CHECK_FAIL            = "A0153" // 邮箱格式校验失败

	USER_LOGIN_EXCEPTION                        = "A0200" // 用户登陆异常，二级宏观错误码
	USER_ACCOUNT_NOT_EXIST                      = "A0201" // 用户账户不存在
	USER_ACCOUNT_ARE_FROZEN                     = "A0202" // 用户账户被冻结
	USER_ACCOUNT_INVALIDATED                    = "A0203" // 用户账户已作废
	USER_ACCOUNT_PASSWORD_FAIL                  = "A0210" // 用户密码错误
	ENTER_PASSWORD_MORE_THAN_THE_LIMIT          = "A0211" // 输入密码次数超限
	USER_IDENTITY_CHECK_FAIL                    = "A0220" // 用户身份校验失败
	USER_WITHOUT_THIRD_PARTY_AUTHORIZATION      = "A0223" // 用户未获得第三方授权
	USER_LOGIN_EXPIRED                          = "A0230" // 用户登陆已过期
	USER_VERIFICATION_CDOE_ERROR                = "A0240" // 用户验证码错误
	ENTER_VERIFICATION_CODE_MORE_THAN_THE_LIMIT = "A0241" // 验证码输入次数超限

	ACCESS_EXCEPTION = "A0300" // 访问权限异常


)

var StatusMsg = map[string]string{
	OK: "一切OK，正确执行后的返回",

	CLIENT_ERROR: "用户端错误，一级宏观错误码",

	USER_REGISTER:                      "用户注册错误，二级宏观错误码",
	USER_NOT_AGREE_PRIVACY_AGREEMENT:   "用户未同意隐私协议",
	RESTRICTED_REGISTRATION_AERA:       "注册国家或地区受限",
	USERNAME_CHECK_FAIL:                "用户名校验失败",
	USERNAME_EXIST:                     "用户名已存在",
	USERNAME_CONTAIN_SENSITIVE_WORDS:   "用户名包含敏感词",
	USERNAME_CONTAIN_SPECIAL_CHARACTER: "用户名包含特殊字符",
	PASSWORD_CHECK_FAIL:                "密码校验失败",
	PASSWORD_LENGTH_ERROR:              "密码长度错误",
	PWSSWORD_STRENGTH_ERROR:            "密码强度不够",
	INPUT_CHECK_CODE_ERROR:             "校验码输入错误",
	INPUT_SMS_CHECK_CODE_ERROR:         "短信校验码输入错误",
	INPUT_EMAIL_CHECK_CODE_ERROR:       "邮件校验码输入错误",
	INPUT_VOICE_CHECK_CODE_ERROR:       "语音校验码输入错误",
	USER_CREDENTIALS_ABNORMAL:          "用户证件异常",
	USER_CREDENTIALS_TYPE_NOT_SELECTED: "用户证件类型未选择",
	ID_NUMBER_CHECK_ERROR:              "身份证号校验错误",
	PASSPORT_CHECK_ERROR:               "护照校验错误",
	USER_BASE_INFO_VALIDATION_FAIL:     "用户基本信息校验失败",
	PHONE_NUMBER_FORMAT_CHECK_FAIL:     "手机号格式校验失败",
	ADDRESS_FORMAT_CHECK_FAIL:          "地址格式校验失败",
	EMAIL_FORMAT_CHECK_FAIL:            "邮箱格式校验失败",

}
