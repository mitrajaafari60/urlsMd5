package service

import (
	"crypto/md5"
	"fmt"
)

type PageGetter func(url string) string

type Downloader struct {
	get_page PageGetter
}

func NewDownloader(pg PageGetter) *Downloader {
	return &Downloader{get_page: pg}
}

func (d *Downloader) GetUrlMd5(url string) string {

	body := d.get_page(url)
	//calculate md5 of the response
	return fmt.Sprintf("%x", md5.Sum([]byte(body)))
}
