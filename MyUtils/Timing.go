package MyUtils

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

//计算请求用时 利用洋葱模型 不过gin自带这个功能
func CountTiming(c *gin.Context) {

	startTime := time.Now()
	c.Next()
	endTime := time.Now()

	// Sub返回持续时间t-u。如果结果超过了可存储在Duration中的最大(或最小)值，则将返回最大(或最小)持续时间。要计算持续时间d的t-d，使用t.Add(-d)。
	latency := endTime.Sub(startTime)
	log.Printf("%v请求用时: %v", c.Request.URL.Path, latency)
}

func CountTiming2(c *gin.Context) {
	startTime := time.Now()
	c.Next()
	// 统计时间
	since := time.Since(startTime)
	log.Printf("%v请求用时: %v", c.Request.URL.Path, since)
}
