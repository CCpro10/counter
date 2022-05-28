package middleware

import (
	"github.com/gin-gonic/gin"
	"main/counter"
)

func Counter(ctx *gin.Context) {
	path := ctx.Request.URL.Path
	counter.Counter.Incr(path, 1)
	ctx.Next()
}
