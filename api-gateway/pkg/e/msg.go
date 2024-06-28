package e

var MsgFlags = map[uint]string{
	Success:       "success",
	Error:         "fail",
	InvalidParams: "请求的参数错误",
}

// GetMsg 根据code获取对应的msg
func GetMsg(code uint) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Error]
}
