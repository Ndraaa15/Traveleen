package message

import "github.com/gin-gonic/gin"

type HTTPResponse struct {
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
}

func ErrorResponse(ctx *gin.Context, StatusCode int64, Message string, Data interface{}) {
	ctx.JSON(int(StatusCode), HTTPResponse{
		Message: Message,
		Status:  false,
		Data:    Data,
	})
}

func SuccessResponse(ctx *gin.Context, StatusCode int64, Message string, Data interface{}) {
	ctx.JSON(int(StatusCode), HTTPResponse{
		Message: Message,
		Status:  true,
		Data:    Data,
	})
}
