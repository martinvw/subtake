package subtake

import (
	"github.com/valyala/fasthttp"
	"time"
)

func get(url string, ssl bool, timeout int) (body []byte) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(site(url, ssl))
	req.Header.Add("Connection", "close")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.6 Safari/605.1.1")
	resp := fasthttp.AcquireResponse()

	client := &fasthttp.Client{}
	client.DoTimeout(req, resp, time.Duration(timeout)*time.Second)

	return resp.Body()
}

func https(url string, ssl bool, timeout int) (body []byte) {
	newUrl := "https://" + url
	body = get(newUrl, ssl, timeout)

	return body
}

func site(url string, ssl bool) (site string) {
	if ssl {
		return "https://" + url
	}
	return "http://" + url
}
