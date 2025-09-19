package middleware

import (
	"bytes"
	"io"
	"time"

	"log"

	"github.com/gin-gonic/gin"
)

// responseWriter custom เพื่อดัก response body
type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseWriter) Write(b []byte) (int, error) {
	// copy ไปเก็บ
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// --- Request log ---
		reqBody, _ := io.ReadAll(c.Request.Body)
		// restore body
		c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))

		log.Printf("[REQ] %s %s Headers: %v Body: %s", c.Request.Method, c.Request.URL.Path, c.Request.Header, string(reqBody))

		// --- Response log ---
		w := &responseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = w
		// call handler
		c.Next()

		latency := time.Since(start)

		log.Printf("[RES] %s %s | Status: %d | Latency: %v | Response: %s", c.Request.Method, c.Request.URL.Path, c.Writer.Status(), latency, w.body.String())
	}
}
