package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"ycProject/crawl/config"
	"ycProject/crawl/engine"
)

var (
	cityProfileRe = regexp.MustCompile(
		`alt="([^"]+)"></a><div class="info clear"><div class="title"><a class="" href="(https://[a-z]+\.lianjia.com/ershoufang/[0-9]+\.html)" target="_blank" data-log_index="[0-9]+"`)
	cityCurPageRe  = regexp.MustCompile(`<a href="(https://[a-z]+\.lianjia.com/ershoufang/)" >二手房`)
	cityNextPageRe = regexp.MustCompile(`"page-data='{"totalPage":100,"curPage":([0-9]+)}'>`)
)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	result := engine.ParseResult{}
	matches := cityProfileRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(
			result.Requests, engine.Request{
				Url:    string(m[2]),
				Parser: NewProfileParser(string(m[1])),
			})
	}
	matches3 := cityCurPageRe.FindAllSubmatch(contents, -1)
	for _, m := range matches3 {
		//u := string(m[1])
		matches2 := cityNextPageRe.FindAllSubmatch(contents, -1)
		for _, u := range matches2 {
			aa := string(u[1])
			bb, err := strconv.Atoi(aa)
			if err != nil {
				panic(err)
			}
			cc := bb + 1
			dd := strconv.Itoa(cc)
			result.Requests = append(result.Requests,
				engine.Request{
					Url:    fmt.Sprintf("%v%v%v", string(m[1]), "pg", dd),
					Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
				})
		}
	}
	return result
}
