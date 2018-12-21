package parser

import (
	"regexp"
	"ycProject/crawl/engine"
	"ycProject/crawl/model"
)

var PriceRe = regexp.MustCompile(`totalPrice:'([^']+)'`)
var HouseModelRe = regexp.MustCompile(`<li><span class="label">房屋户型</span>([^<]+)</li>`)
var FloorRe = regexp.MustCompile(`<li><span class="label">所在楼层</span>([^<]+)</li>`)
var CoverdAreaRe = regexp.MustCompile(`<li><span class="label">建筑面积</span>([^<]+)</li>`)
var ModelStructRe = regexp.MustCompile(`<li><span class="label">户型结构</span>([^<]+)</li>`)
var TeachingAreaRe = regexp.MustCompile(`<li><span class="label">套内面积</span>([^<]+)</li>`)
var BuildingModelRe = regexp.MustCompile(`<li><span class="label">建筑类型</span>([^<]+)</li>`)
var OrientationRe = regexp.MustCompile(`<li><span class="label">房屋朝向</span>([^<]+)</li>`)
var BuildingStructRe = regexp.MustCompile(` <li><span class="label">建筑结构</span>([^<]+)</li>`)
var DecorationRe = regexp.MustCompile(`<li><span class="label">装修情况</span>([^<]+)</li>`)
var ElevatorRateRe = regexp.MustCompile(`<li><span class="label">梯户比例</span>([^<]+)</li>`)
var ElevatorRe = regexp.MustCompile(`<li><span class="label">配备电梯</span>([^<]+)</li>`)
var PropertyRe = regexp.MustCompile(`<li><span class="label">产权年限</span>([^<]+)</li>`)
var idRe = regexp.MustCompile(`链家编号</span><span class="info">([\d]+)<span class="jubao">`)

func parseProfile(contents []byte, name string, url string) engine.ParseResult {

	profile := model.Profile{}
	profile.Name = name

	profile.Price = extractString(
		contents, PriceRe)
	profile.HouseModel = extractString(
		contents, HouseModelRe)
	profile.Floor = extractString(
		contents, FloorRe)
	profile.CoverdArea = extractString(
		contents, CoverdAreaRe)
	profile.ModelStruct = extractString(
		contents, ModelStructRe)
	profile.TeachingArea = extractString(
		contents, TeachingAreaRe)
	profile.BuildingModel = extractString(
		contents, BuildingModelRe)
	profile.Orientation = extractString(
		contents, OrientationRe)
	profile.BuildingStruct = extractString(
		contents, HouseModelRe)
	profile.HouseModel = extractString(
		contents, BuildingStructRe)
	profile.Decoration = extractString(
		contents, DecorationRe)
	profile.ElevatorRate = extractString(
		contents, ElevatorRateRe)
	profile.Elevator = extractString(
		contents, ElevatorRe)
	profile.Property = extractString(
		contents, PropertyRe)
	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "lianjia",
				Id:      extractString([]byte(url), idRe),
				Payload: profile,
			},
		},
	}
	return result
}
func extractString(
	contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return parseProfile(contents, p.userName, url)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return "ProfileParser", p.userName
}
func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		userName: name,
	}
}
