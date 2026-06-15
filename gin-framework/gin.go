package ginframework

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func HealthCheck(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{"message":"server is running"})
}