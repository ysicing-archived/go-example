// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by WTFPL LICENSE
// license that can be found in the LICENSE file.

package middleware

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	ns = "goexample"
)

var (
	reqCnt = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name:      "http_request_count_total",
		Namespace: ns,
		Help:      "Total number of HTTP requests made.",
	}, []string{"code", "method", "url", "client_ip"})
	reqDur = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:      "http_request_duration_seconds",
		Namespace: ns,
		Help:      "HTTP request latencies in seconds.",
	}, []string{"code", "method", "url"})
	resSz = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:      "response_size_bytes",
		Namespace: ns,
		Help:      "The HTTP response sizes in bytes.",
	}, []string{"code", "method", "url"})
	reqSz = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:      "request_size_bytes",
		Help:      "request_size_bytes",
		Namespace: ns,
	}, []string{"code", "method", "url"})
)

func init() {
	prometheus.MustRegister(reqCnt, reqDur, resSz, reqSz)
	registerWithWeight("prom", 100, func() gin.HandlerFunc {
		return prom()
	})
}

func prom() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		reqsz := computeApproximateRequestSize(c.Request)
		c.Next()
		url := c.Request.URL.Path
		if !strings.HasPrefix(url, "/api") {
			return
		}
		status := strconv.Itoa(c.Writer.Status())
		elapsed := float64(time.Since(start)) / float64(time.Second)
		ressz := float64(c.Writer.Size())

		reqCnt.WithLabelValues(status, c.Request.Method, url, c.ClientIP()).Inc()
		reqDur.WithLabelValues(status, c.Request.Method, url).Observe(elapsed)
		resSz.WithLabelValues(status, c.Request.Method, url).Observe(ressz)
		reqSz.WithLabelValues(status, c.Request.Method, url).Observe(float64(reqsz))
	}
}

func computeApproximateRequestSize(r *http.Request) int {
	s := 0
	if r.URL != nil {
		s = len(r.URL.Path)
	}

	s += len(r.Method)
	s += len(r.Proto)
	for name, values := range r.Header {
		s += len(name)
		for _, value := range values {
			s += len(value)
		}
	}
	s += len(r.Host)

	// N.B. r.Form and r.MultipartForm are assumed to be included in r.URL.

	if r.ContentLength != -1 {
		s += int(r.ContentLength)
	}
	return s
}
