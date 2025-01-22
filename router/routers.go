package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "hello world"})
		})

		api := v1.Group("/api")
		{
			api.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "get success")
			})

		}
	}

}
