package amserr

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

func NewApiError(code string, err error, details string) *ApiError {
	e := ApiError{}
	e.Code = code
	e.Err = fmt.Sprintf("%s", err)
	e.Details = details
	e.init()
	return &e
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
	return fmt.Sprintf("Error message : %s", e.Err)
}
