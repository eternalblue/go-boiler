package tests

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func Server() *gin.Engine {
	return gin.Default()
}

func DoRequest(method, path string, params []gin.Param, body interface{}, handlerFunc gin.HandlerFunc) (res *httptest.ResponseRecorder) {
	var req *http.Request

	if body != nil {
		req, _ = http.NewRequest(method, path, bytes.NewReader([]byte(body.(string))))
	} else {
		req, _ = http.NewRequest(method, path, nil)
	}

	res = httptest.NewRecorder()

	context, _ := gin.CreateTestContext(res)
	context.Params = params
	context.Request = req

	handlerFunc(context)

	return
}
