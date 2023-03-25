package middleware

import (
	"Skyriders/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeserializeFlight() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		flight := &model.Flight{}
		if err := ctx.ShouldBindJSON(&flight); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to decode json"})
			return
		}

		ctx.Set("flight", flight)
		ctx.Next()
	}
}
