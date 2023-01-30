package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

/*
POST / HTTP/1.1
Host: localhost:18888
Accept-Encoding: gzip
Content-Length: 1943
Content-Type: multipart/form-data; boundary=de571fe4b9797dac19eedf922cd4c03e6afff0e4888421d650377381d643
User-Agent: Go-http-client/1.1

--de571fe4b9797dac19eedf922cd4c03e6afff0e4888421d650377381d643
Content-Disposition: form-data; name="id"
*/
func main() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("id", "deneb")
	fileWriter, err := writer.CreateFormFile("image", "images/search.png")
	if err != nil {
		panic(err)
	}
	readFile, err := os.Open("images/search.png")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()
	io.Copy(fileWriter, readFile)
	writer.Close()
	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	log.Println("Status: ", resp.Status)
}
