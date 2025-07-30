package config

import (
	"bytes"
	"io/ioutil"
	"log"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// สร้าง ResponseRecorder เพื่อจับข้อมูลที่เขียนลงใน ResponseWriter
type ResponseRecorder struct {
	gin.ResponseWriter
	StatusCode int
	Body       string
}

// LogRequestMiddleware สำหรับ log ข้อมูลของ request และ response
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		state := path.Base(c.Request.URL.Path)

		// Log Request Method, URL, Headers, และ Body ในบรรทัดเดียว
		logMessage := "[Req][" + state + "] method:" + c.Request.Method + " url:" + c.Request.URL.String()
		if c.Request.Method == "POST" || c.Request.Method == "PUT" {
			bodyBytes, err := c.GetRawData()
			if err != nil {
				log.Printf("Error reading body: %v\n", err)
			}
			logMessage += " body:" + strings.ReplaceAll(string(bodyBytes), "\n", "") // เปลี่ยน \n เป็นช่องว่าง
			logMessage = strings.ReplaceAll(logMessage, "\r", "")                    // ลบ \r หากมี
			// ต้องใส่คืนข้อมูลกลับไปใน c.Request.Body เพื่อให้ handler อื่นใช้ข้อมูลได้
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		log.Printf("%s\n", logMessage)

		// สร้าง ResponseRecorder สำหรับ log response
		responseRecorder := &ResponseRecorder{ResponseWriter: c.Writer, StatusCode: 200}
		c.Writer = responseRecorder

		// เรียก handler ต่อไป
		c.Next()

		// Log Response Status และ เวลาใช้ในการประมวลผล
		duration := time.Since(startTime)
		log.Printf("[Res][%s] status:%d duration:%v json:%s\n", state, responseRecorder.StatusCode, duration, responseRecorder.Body)

	}
}

// บันทึกสถานะและเนื้อหาของ Response
func (rec *ResponseRecorder) Write(b []byte) (int, error) {
	rec.Body = string(b)               // เก็บข้อมูลใน Body
	return rec.ResponseWriter.Write(b) // ส่งข้อมูลไปยัง ResponseWriter
}

// เปลี่ยนสถานะการตอบกลับ
func (rec *ResponseRecorder) WriteHeader(code int) {
	rec.StatusCode = code
	rec.ResponseWriter.WriteHeader(code)
}
