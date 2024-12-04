package request

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type (
	contextWrapperService interface {
		Bind(data any) error
	}

	contextWrapper struct {
		Context   *gin.Context
		validator *validator.Validate
	}
)

func ContextWrapper(ctx *gin.Context) contextWrapperService {
	return &contextWrapper{
		Context:   ctx,
		validator: validator.New(),
	}
}

func (c *contextWrapper) Bind(data any) error {
	if ex := c.Context.Bind(data); ex != nil {
		log.Printf("Error: Bind data failed: %s", ex.Error())
	}

	if ex := c.validator.Struct(data); ex != nil {
		log.Printf("Error: Validate data failed: %s", ex.Error())
	}

	return nil
}
