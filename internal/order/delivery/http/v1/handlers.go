package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type OrderHandlers struct {
	group *gin.RouterGroup
	v     *validator.Validate
}
