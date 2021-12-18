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

func RenderData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &response{Code: 0, Message: "ok", Data: data})
}

func RenderErr(c *gin.Context, err error) {
	c.JSON(http.StatusOK, &response{Code: -1, Message: err.Error()})
}
