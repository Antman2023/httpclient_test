package main

import (
	"log"
	"os"
	"time"

	"github.com/ddliu/go-httpclient"
)

var url string

func worker() {
	for {
		res, err := httpclient.Begin().Get(url, nil)
		if err == nil && res != nil {
			res.ReadAll()
		}
	}
}

func main() {
	url = os.Getenv("URL")
	log.Println(url)
	httpclient.Defaults(httpclient.Map{
		httpclient.OPT_TIMEOUT:        30,
		httpclient.OPT_CONNECTTIMEOUT: 5,
		httpclient.OPT_USERAGENT:      "Mozilla/5.0 (Windows NT 6.1; rv:42.0) Gecko/20100101 Firefox/42.0",
	})
	for i := 0; i < 10; i++ {
		go worker()
		time.Sleep(time.Second)
	}
	select {}
}
