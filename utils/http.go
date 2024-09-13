package utils

import (
	"net"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

func NewHTTPClient() *http.Client {
	/**
		* Recommended Timestream write client SDK configuration:
		* Set the SDK retry count to 10.
		* Use SDK DEFAULT_BACKOFF_STRATEGY.
	  * Set RequestTimeout to 20 seconds.
	  * Set the max connections to 5000 or higher.
	*/

	tr := &http.Transport{
		ResponseHeaderTimeout: 20 * time.Second,
		Proxy:                 http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			KeepAlive: 30 * time.Second,
			Timeout:   30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	if err := http2.ConfigureTransport(tr); err != nil {
		panic(err)
	}
	return &http.Client{
		Timeout:   25 * time.Second,
		Transport: tr,
	}
}
