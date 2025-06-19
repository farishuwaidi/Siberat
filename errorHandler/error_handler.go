package errorhandler

import (
	"Siberat/dto"
	"Siberat/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerError(c *gin.Context, err error) {
	var statusCode int

	switch err.(type) {
	case *NotFoundError:
		statusCode = http.StatusNotFound
	case *BadRequestError:
		statusCode = http.StatusBadRequest
	case *InternalServerError:
		statusCode = http.StatusInternalServerError
	case *UnAuthorizedError:
		statusCode = http.StatusUnauthorized
	default:
		statusCode = http.StatusInternalServerError
	}

	response := helper.Response(dto.ResponseParams{
		StatusCode: statusCode,
		Message: err.Error(),
	})

	c.JSON(statusCode, response)
}