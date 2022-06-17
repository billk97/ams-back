package employee

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ApiError struct {
	Code          string
	Err           string
	ClientIp      string
	RequestMethod string
	TimeStamp     time.Time
	ErroUuid      string
	Details       string
}

func (e *ApiError) init() {
	if e.ErroUuid == "" {
		e.ErroUuid = uuid.New().String()
	}
	e.TimeStamp = time.Now().UTC()
}

func (e *ApiError) Enhance(c *gin.Context) {
	e.ClientIp = c.ClientIP()
	e.RequestMethod = c.Request.Method
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("status: %v", e.Err)
}

// func GinError(e *ApiError) *gin.H {
// 	return fmt.Sprintf(
// 		"code":          e.code,
// 		"message":       e.err,
// 		"clientIp":      e.clientIp,
// 		"requestMethod": e.requestMethod,
// 		"timeStamp":     e.timeStamp,
// 		"erroUuid":      e.erroUuid,
// 		"details":       e.details,
// 	)
// }
