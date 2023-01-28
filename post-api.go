package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// curl -d key1=value1&key2=value2&key3=value3 http://localhost:18888
/*
POST / HTTP/1.1
Host: localhost:18888
Accept-Encoding: gzip
Content-Length: 35
Content-Type: application/x-www-form-urlencoded
User-Agent: Go-http-client/1.1

key1=value1&key2=value2&key3=value3
*/
func main() {
	values := url.Values{
		"key1": {"value1"},
		"key2": {"value2"},
		"key3": {"value3"},
	}
	resp, _ := http.PostForm("http://localhost:18888", values) //오류 처리 생략
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
}
