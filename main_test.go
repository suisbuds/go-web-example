package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/suisbuds/miao/internal/router"
	"github.com/suisbuds/miao/pkg/errcode"
)

func TestSetupSetting(t *testing.T) {
	err := setupSetting()
	assert.NoError(t, err, "setupSetting should not return an error")

	// base64.StdEncoding.DecodeString("xxx") // token-payload base64解码. 不要在 Payload 中明文存储敏感信息, 否则进行不可逆加密. JWT 过期时间存储在 payload 中, 一旦签发不可变更

}

// 超时中间件
func TestTimeoutMiddleware(t *testing.T) {

	router := router.NewRouter()

	router.GET("/slow", func(c *gin.Context) {
		time.Sleep(10 * time.Second)
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	// 创建请求
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/slow", nil)
	router.ServeHTTP(w, req)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Response should be valid JSON")
	assert.Equal(t, float64(errcode.RequestTimeout.Code()), response["code"], "Error code should match RequestTimeout")
	assert.Equal(t, errcode.RequestTimeout.Msg(), response["msg"], "Error message should match RequestTimeout")
}
