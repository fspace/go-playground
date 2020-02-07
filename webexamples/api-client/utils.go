package main

import (
	"net"
	"net/http"
	"time"
)

// GetClient return a customized Client based on default httpClient
// https://medium.com/@jake0malay3/7-tips-on-how-to-write-kick-ass-high-performance-golang-microservices-9f71d4c67a0a
func GetClient() *http.Client {
	panic("not specified the <n>")
	tr := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   n * time.Second,
			KeepAlive: n * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: n * time.Second,

		ExpectContinueTimeout: n * time.Second,
		ResponseHeaderTimeout: n * time.Second, MaxIdleConns: n,
		MaxConnsPerHost: n,
	}
	cli := &http.Client{
		Transport: tr,
		Timeout:   n * time.Second,
	}
	return cli
}
