package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ddliu/go-httpclient"
)

var url string

func worker() {
	for {
		res, err := httpclient.Begin().Get(url, nil)
		if err == nil && res != nil {
			res.ReadAll()
		} else {
			time.Sleep(time.Second)
		}
	}
}

func updateUrl() error {
	res, err := httpclient.Begin().Get(os.Getenv("URL"), nil)
	if err != nil {
		return err
	}
	u, err := res.ToString()
	if err != nil {
		return err
	}
	u = strings.TrimRight(u, "\n")
	url = u
	log.Println(url)
	return nil
}

func main() {
	updateUrl()
	httpclient.Defaults(httpclient.Map{
		httpclient.OPT_TIMEOUT:        30,
		httpclient.OPT_CONNECTTIMEOUT: 5,
		httpclient.OPT_USERAGENT:      "Mozilla/5.0 (Windows NT 6.1; rv:42.0) Gecko/20100101 Firefox/42.0",
	})
	for i := 0; i < 10; i++ {
		go worker()
		time.Sleep(time.Second)
	}

	go func() {
		for {
			updateUrl()
			time.Sleep(time.Second * 30)
		}
	}()

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte("Hello World!"))
	})
	http.ListenAndServe(":80", nil)
}
