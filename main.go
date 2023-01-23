package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

const (
	requestIdKey             = "reqID"
	requestsPerSecond        = 3
	maxTokensConsumedPerCall = 3
)

var (
	requestCounter = 0
	processingTime = time.Second * 2
	rpsLimiter     = rate.NewLimiter(requestsPerSecond, maxTokensConsumedPerCall)
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

func rateLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !rpsLimiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "max requests per second (RPS) reached."})
			c.Abort()
		}
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
	server.Use(rateLimiterMiddleware())
	server.Use(requestIdMiddleware())

	registerHandler(server)

	err := server.Run(":8080")
	if err != nil {
		log.Fatal("unable to run http server")
	}
}
