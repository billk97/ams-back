package amserr

import (
	"github.com/gin-gonic/gin"
)

func ErrorWrapper(a appHandler) gin.HandlerFunc {
	return a.HandlerFunc
}

type appHandler func(*gin.Context) *ApiError

func (a appHandler) HandlerFunc(c *gin.Context) {
	err := a(c)
	if err != nil {
		err.Enhance(c)
		// todo handle 400, 500 error type differently
		c.AbortWithStatusJSON(400, err)
	}
}
