package main

import (
	"fmt"
	"io"
	"net/http"
)

type writerWeb struct{}

func (writerWeb) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func Web_main() {
	request, err := http.Get("http://sebamgb.github.io")
	if err != nil {
		fmt.Println(err)
	}
	e := writerWeb{}
	io.Copy(e, request.Body)
}
