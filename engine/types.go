package engine

type ParserFunc func(contents []byte, url string) ParseResult

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}
type Request struct {
	Url    string
	Parser Parser
}
type ParseResult struct {
	Requests []Request

	Items []Item
}
type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}
type FuncParser struct {
	parser ParserFunc
	name   string
}

//implement the interface of FuncParser
func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}
func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}
func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
