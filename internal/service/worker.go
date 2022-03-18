package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func get_page(url string) string {
	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		return ""
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	defer resp.Body.Close()
	return string(body)
}

func Worker(in <-chan string) {
	for url := range in {
		d := NewDownloader(get_page)
		md5 := d.GetUrlMd5(url)
		fmt.Print(" " + url + " " + md5)
	}
}
