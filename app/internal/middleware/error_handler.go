package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	appErrors "github.com/nithishbijadirajeev-droid/forgeflow/app/pkg/errors"
	"github.com/nithishbijadirajeev-droid/forgeflow/app/pkg/response"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err

		if appErr, ok := err.(*appErrors.AppError); ok {
			response.Error(
				c,
				appErr.StatusCode,
				appErr.Message,
				appErr.Error(),
			)
			return
		}

		response.Error(
			c,
			http.StatusInternalServerError,
			"Internal server error",
			err.Error(),
		)
	}
}
