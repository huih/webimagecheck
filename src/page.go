package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gotools/logs"
	"github.com/gotools/lists"
)

var PageUrlList = fifolist.New()
var PageStop = false

func GetDetailUrl(HostUrl string) string {
	query,err := goquery.NewDocument(HostUrl)
	if err != nil {
		logs.Debug("err: %s", err.Error())
		return ""
	}
	titles := query.Find("#primary div.primary-site h2.entry-title")

	for index := 0; index < titles.Length(); index++ {
		t := titles.Eq(index)
		ta := t.Find("a[href]")
		tsrc,_ := ta.Attr("href")
		if len(tsrc) > 0 {
			PageUrlList.Add(tsrc)
		}
	}
	
	naviDom := query.Find("#pagenavi a.page-numbers")
	if naviDom.Length() <= 0 {
		return ""
	}
	
	nextPage := naviDom.Eq(naviDom.Length() - 1)
	title,_ := nextPage.Attr("title")
	if title == "下页" {
		href,_ := nextPage.Attr("href")
		return href
	}
	return ""	
}

func GetPages(hostUrl string) {
	url := hostUrl
	
	for {
		url = GetDetailUrl(url)
		if len(url) <= 0 {
			break
		}
	}
	PageStop = true
}