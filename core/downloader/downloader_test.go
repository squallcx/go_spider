// Copyright 2014 Hu Cong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//
package downloader_test

import (
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "github.com/hu17889/go_spider/core/common/page"
    "github.com/hu17889/go_spider/core/common/request"
    "github.com/hu17889/go_spider/core/downloader"
    "testing"
)

func TestDownloadHtml(t *testing.T) {
    //request := request.NewRequest("http://live.sina.com.cn/zt/api/l/get/finance/globalnews1/index.htm?format=json&callback=t13975294&id=23521&pagesize=45&dire=f&dpc=1")
    var req *request.Request
    req = request.NewRequest("http://live.sina.com.cn/zt/l/v/finance/globalnews1/", "html")

    var dl downloader.Downloader
    dl = downloader.NewHttpDownloader()

    var p *page.Page
    p = dl.Download(req)

    var doc *goquery.Document
    doc = p.GetHtmlParser()
    //fmt.Println(doc)
    //body := p.GetBodyStr()
    //fmt.Println(body)

    var s *goquery.Selection
    s = doc.Find("body")
    if s.Length() < 1 {
        t.Error("html parse failed!")
    }

    /*
       doc, err := goquery.NewDocument("http://live.sina.com.cn/zt/l/v/finance/globalnews1/")
       if err != nil {
           fmt.Printf("%v",err)
       }
       s := doc.Find("meta");
       fmt.Println(s.Length())

       resp, err := http.Get("http://live.sina.com.cn/zt/l/v/finance/globalnews1/")
       if err != nil {
           fmt.Printf("%v",err)
       }
       defer resp.Body.Close()
       doc, err = goquery.NewDocumentFromReader(resp.Body)
       s = doc.Find("meta");
       fmt.Println(s.Length())
    */
}

func TestDownloadJson(t *testing.T) {
    var req *request.Request
    req = request.NewRequest("http://live.sina.com.cn/zt/api/l/get/finance/globalnews1/index.htm?format=json&id=23521&pagesize=4&dire=f&dpc=1", "json")

    var dl downloader.Downloader
    dl = downloader.NewHttpDownloader()

    var p *page.Page
    p = dl.Download(req)

    var jsonMap interface{}
    jsonMap = p.GetJsonMap()
    fmt.Printf("%v", jsonMap)

    //fmt.Println(doc)
    //body := p.GetBodyStr()
    //fmt.Println(body)

}
