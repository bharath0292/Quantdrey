package util

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type contextKey string

const ginContextKey contextKey = "GinContextKey"

func GinToHttpContext(c *gin.Context) *http.Request {
	ctx := context.WithValue(context.Background(), ginContextKey, c)
	return c.Request.WithContext(ctx)
}

func HttpToGinContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(ginContextKey)
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}

	return gc, nil
}
