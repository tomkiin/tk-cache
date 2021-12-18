package httpx

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RESTful 响应
type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func RenderOK(c *gin.Context) {
	c.JSON(http.StatusOK, &response{Code: 0, Message: "ok"})
}
