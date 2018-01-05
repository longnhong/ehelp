package middleware

import (
	"ehelp/x/rest"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var errResponse = map[string]interface{}{
					"error":  err.(error).Error(),
					"status": "error",
				}
				if httpError, ok := err.(rest.IHttpError); ok {
					c.JSON(httpError.StatusCode(), errResponse)
				} else {
					c.JSON(500, errResponse)
				}
			}
		}()
		c.Next()
	}
}

func AddHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", c.GetHeader("Origin"))
		c.Header("Access-Control-Allow-Methods", "OPTIONS, HEAD, GET, POST, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Range, Content-Disposition")
		c.Header("Access-Control-Allow-Credentials", "true")
	}
}
