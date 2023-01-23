package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

const (
	requestIdKey = "reqID"
)

var (
	requestCounter = 0
	processingTime = time.Second * 10
)

type testHandler struct {
	router *gin.Engine
}

func (h *testHandler) Register() {
	h.router.GET("test", h.testEndpoint)
}

func (h *testHandler) testEndpoint(c *gin.Context) {
	reqID, _ := c.Get(requestIdKey)

	select {
	case <-time.After(processingTime):
		log.Printf("processed request ID: %v\n", reqID)
		c.JSON(http.StatusOK, gin.H{"data": "processed"})
	}
}

func requestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(requestIdKey, requestCounter)
		requestCounter++
	}
}
func registerHandler(router *gin.Engine) {
	h := &testHandler{
		router: router,
	}
	h.Register()
}

func main() {
	server := gin.Default()
	server.Use(requestIdMiddleware())
	registerHandler(server)

	err := server.Run(":8080")
	if err != nil {
		log.Fatal("unable to run http server")
	}
}
