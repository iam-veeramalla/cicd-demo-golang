package ginframework



import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)


func TestHealthCheck(t *testing.T){
	gin.SetMode(gin.TestMode)

	r:=gin.Default()

	r.GET("/health",HealthCheck)

	
	w:=httptest.NewRecorder()

	req,_:=http.NewRequest("GET","/health",nil)

	r.ServeHTTP(w,req)

	assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "server is ok")
}