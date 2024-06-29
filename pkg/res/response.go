package res

import (
	"github.com/yaqia77/memorandum/pkg/e"

	"github.com/gin-gonic/gin"
)

// Response 基础序列化器
type Response struct {
	Status uint        `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// DataList 带有总数的Data结构
type DataList struct {
	Item  interface{} `json:"item"`
	Total int64       `json:"total"`
}

// TokenData 带有token的Data结构
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

func ginH(msgCode int, data interface{}) gin.H {
	return gin.H{
		"code": msgCode,
		"msg":  e.GetMsg(uint(msgCode)),
		"data": data,
	}
}

// func SuccessResponse(c *gin.Context, data interface{}) {
// 	resp := Response{
// 		Code:    http.StatusOK,
// 		Message: "success",
// 		Data:    data,
// 	}
// 	c.JSON(http.StatusOK, resp)
// }

// func ErrorResponse(c *gin.Context, code int, message string) {
// 	resp := Response{
// 		Code:    code,
// 		Message: message,
// 		Data:    nil,
// 	}
// 	c.JSON(code, resp)
// }
