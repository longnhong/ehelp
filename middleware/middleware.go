package middleware

import (
	"ehelp/o/auth"
	"ehelp/x/rest"
	"fmt"
	"g/x/web"
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
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		//remember
		if c.Request.Method == "OPTIONS" {
			c.Writer.WriteHeader(200)
			return
		}
		c.Next()
	}
}

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token = web.GetToken(c.Request)
		var _, err = auth.GetByID(token)
		if err != nil {
			rest.AssertNil(rest.Unauthorized(err.Error()))
		}
	}
}

func MustBeAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		MustAuthenticate(c, "admin")
	}
}
func MustBeSuperAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		MustAuthenticate(c, "super_admin")
	}
}
func MustBeStaff() gin.HandlerFunc {
	return func(c *gin.Context) {
		MustAuthenticate(c, "staff")
	}
}
func MustBeOwner() gin.HandlerFunc {
	return func(c *gin.Context) {
		MustAuthenticate(c, "owner")
	}
}
func MustAuthenticate(ctx *gin.Context, role string) {
	var errResponse = map[string]interface{}{
		"status": "error",
	}
	var token = web.GetToken(ctx.Request)
	var auth, err = auth.GetByID(token)
	if err != nil {
		errResponse["error"] = "access token not found"
		ctx.JSON(401, errResponse)
	} else {
		if auth.Role != role {
			errResponse["error"] = fmt.Sprintf("Unauthorize! you must be %s to access", role)
			ctx.JSON(401, errResponse)
		} else {
			ctx.Next()
		}
	}
	ctx.Abort()
}
