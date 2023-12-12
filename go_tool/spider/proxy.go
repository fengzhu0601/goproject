package spider

import (
	"log"
	"net/http"
	"net/url"
)

func CreateProxyClient() *http.Client {
	proxyUrl, err := url.Parse("http://58.220.95.55:9400")
	if err != nil {
		log.Fatal(err)
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	return &http.Client{
		Transport: transport,
	}
}
