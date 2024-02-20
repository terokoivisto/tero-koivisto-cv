package api

import (
	"backend/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CVRoutes(r *gin.Engine, dynamo db.TableConfig) {
	r.GET("/cv", func(context *gin.Context) {
		resp, err := dynamo.CV("Tero Koivisto")
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, resp)
	})
}
