package parser

import (
	"regexp"
	"ycProject/crawl/config"
	"ycProject/crawl/engine"
)

const cityErShouRe = `<li><a  class="" href="(https://[0-9a-z]+\.lianjia.com/ershoufang)/" >二手房`

var cityListRe = regexp.MustCompile(`<li><a href="(https://[a-z]+\.lianjia.com/)">([^<]+)</a></li>`)

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	matches := cityListRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	//this is the number of city limit
	//limit := 5
	for _, m := range matches {

		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCityList, config.ParseCityList),})
		//limit--
		//if limit == 0 {
		//break
		//}
	}
	re1 := regexp.MustCompile(cityErShouRe)
	matches1 := re1.FindAllSubmatch(contents, -1)

	for _, m := range matches1 {
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			//require a new parser
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),})
	}
	return result
}
