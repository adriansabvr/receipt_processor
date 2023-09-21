package v1

import (
	"fmt"
	"net/http"

	"github.com/adriansabvr/receipt_processor/pkg/logger"
	"github.com/gin-gonic/gin"
)

type response struct {
	Error string `json:"error" example:"message"`
}

// OK returns a 200.
func OK(c *gin.Context, val interface{}) {
	c.JSON(http.StatusOK, val)
}

// BadRequest returns a 400.
func BadRequest(c *gin.Context, l logger.Interface, err error) {
	abort(c, l, http.StatusBadRequest, err)
}

// InternalServerError returns a 500.
func InternalServerError(c *gin.Context, l logger.Interface, err error) {
	abort(c, l, http.StatusInternalServerError, err)
}

// abort does some pre-processing before returning the given HTTP status code.
func abort(c *gin.Context, l logger.Interface, code int, err error) {
	errMsg := fmt.Sprintf("%#s", err)
	l.Error(errMsg, err)
	c.AbortWithStatusJSON(code, response{Error: errMsg})
}
