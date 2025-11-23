package responses

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvaildParam
	CodeUserExist
	CodeUserNotExist
	CodeInvaildPassword
	CodeServerBusy

	CodeNeedAuth
	CodeInvalidAuth
	CodeNeedLogin
	CodeInvalidID
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "请求成功",
	CodeInvaildParam:    "请求参数错误",
	CodeUserExist:       "用户已存在",
	CodeUserNotExist:    "用户名不存在",
	CodeInvaildPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",
	CodeNeedAuth:        "需要token",
	CodeInvalidAuth:     "无效的token",
	CodeNeedLogin:       "需要登陆",
	CodeInvalidID:       "ID错误",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
