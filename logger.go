package fen

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		status := fmt.Sprintf(NoticeColor, strconv.Itoa(c.StatusCode))
		if c.StatusCode >= 400 {
			status = fmt.Sprintf(WarningColor, strconv.Itoa(c.StatusCode))
		}
		if c.StatusCode >= 500 {
			status = fmt.Sprintf(ErrorColor, strconv.Itoa(c.StatusCode))
		}
		uri := fmt.Sprintf(InfoColor, c.Req.RequestURI)
		latency := fmt.Sprintf(WarningColor, time.Since(t))
		log.Printf("[%s] %s %s in %s", status, c.Method, uri, latency)
	}
}
