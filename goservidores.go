package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	initTime := time.Now()
	channel := make(chan string)
	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}
	i := 0
	for {
		if i > 2 {
			break
		}
		for _, server := range servers {
			go chekServer(server, channel)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(<-channel)
		i++
	}

	// for i := 0; i < len(servers); i++ {
	// 	fmt.Println(<-channel)
	// }

	endTime := time.Since(initTime)
	fmt.Printf("run time: %s\n", endTime)
}

func chekServer(server string, c chan string) {
	_, err := http.Get(server)
	if err != nil {
		fmt.Println(server, "does not work")
		c <- server + " does not work"
	} else {
		fmt.Println(server, "is work")
		c <- server + " is work"
	}
}
