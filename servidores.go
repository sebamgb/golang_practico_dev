package main

import (
	"fmt"
	"net/http"
	"time"
)

func Servidores_main() {
	initTime := time.Now()
	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}
	for _, server := range servers {
		chekServer(server)
	}
	endTime := time.Since(initTime)
	fmt.Printf("run time: %s\n", endTime)
}

func chekServer(server string) {
	_, err := http.Get(server)
	if err != nil {
		fmt.Println(server, "does not work")
	} else {
		fmt.Println(server, "is work")
	}
}
