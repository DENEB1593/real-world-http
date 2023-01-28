package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// curl -G --data-urlencode "query=hello word" http://localhost:18888
func main() {
	values := url.Values{
		"query": {"hello world"},
	}
	resp, _ := http.Get("http://localhost:18888" + "?" + values.Encode()) //오류 처리 생략
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
}
