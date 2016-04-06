package funyu

import (
	"fmt"
	"strings"
	"regexp"
)

var (
	ImutableElementError = fmt.Errorf("this element is imutable")

	linkReg = regexp.MustCompile(`\[(IMG:)? *(.*?)\]\((.*?)\)`)
)

type Line struct {
	*ParentElementBase
}

func NewLine() *Line {
	return &Line{NewParentElementBase()}
}

func (self *Line) String() string {
	var r string
	for _, x := range self.children {
		r += x.String()
	}
	return strings.TrimRight(r, "\n") + "\n"
}

func (self *Line) HTML(level int) string {
	var r string
	for _, x := range self.children {
		r += x.HTML(level)
	}
	return r + "<br>\n"
}

func (self *Line) parseStr(s string) error {
	for i := 0; i < len(s)-2; i++ {
		start := ""
		end := ""
		var elm ParentElement

		switch s[i:i+2] {
		case "<<":
			start = "<<"
			end = ">>"
			elm = NewEmphasis()
		case "[[":
			start = "[["
			end = "]]"
			elm = NewKeyword()
		case "{{":
			start = "{{"
			end = "}}"
			elm = NewCode()
		default:
			continue
		}

		if i > 0 {
			self.append(NewString(s[:i]))
		}
		self.append(elm)
		st := i
		c := 0
		i += 2
		for ; i < len(s); i++ {
			if strings.HasPrefix(s[i:], start) {
				c++
			} else if strings.HasPrefix(s[i:], end) {
				if c == 0 {
					if e := elm.Feed(s[st+2:i]); e != nil {
						return e
					}
					if i+2 < len(s) {
						self.parseStr(s[i+2:])
					}
					return nil
				} else {
					c --
				}
			}
		}

		return fmt.Errorf("Inline element is not closed.")
	}

	self.append(NewString(s))
	return nil
}

func (self *Line) Feed(s string) error {
	parseLink := func(s string) {
		m := linkReg.FindStringSubmatch(s)
		if m[1] == "" {
			self.append(NewLink(m[2], m[3]))
		} else {
			self.append(NewImageLink(m[2], m[3]))
		}
	}

	pos := 0
	for _, m := range linkReg.FindAllStringIndex(s, -1) {
		if e := self.parseStr(s[pos:m[0]]); e != nil {
			return e
		}
		parseLink(s[m[0]:m[1]])
		pos = m[1]
	}
	return self.parseStr(s[pos:])
}

type String struct {
	*ElementBase

	content string
}

func (self *String) Feed(s string) error {
	return ImutableElementError
}

func (self *String) Children() []Element {
	return nil
}

func NewString(content string) *String {
	return &String{
		NewElementBase(),
		content,
	}
}

func (self *String) String() string {
	return self.content
}

func (self *String) HTML(level int) string {
	return self.content
}

type Emphasis struct {
	*Line
}

func NewEmphasis() *Emphasis {
	return &Emphasis{NewLine()}
}

func (self *Emphasis) String() string {
	return "<<" + strings.Join(self.StringList(), "") + ">>"
}

func (self *Emphasis) HTML(level int) string {
	return "<em>" + strings.Join(self.HTMLList(level), "") + "</em>"
}

type Keyword struct {
	*Line
}

func NewKeyword() *Keyword {
	return &Keyword{NewLine()}
}

func (self *Keyword) String() string {
	return "[[" + strings.Join(self.StringList(), "") + "]]"
}

func (self *Keyword) HTML(level int) string {
	return "<strong>" + strings.Join(self.HTMLList(level), "") + "</strong>"
}

type Code struct {
	*ElementBase

	content string
}

func NewCode() *Code {
	return &Code{
		NewElementBase(),
		"",
	}
}

func (self *Code) Feed(s string) error {
	self.content += s
	return nil
}

func (self *Code) append(elm Element) {
}

func (self *Code) Children() []Element {
	return nil
}

func (self *Code) String() string {
	return "{{" + self.content + "}}"
}

func (self *Code) HTML(level int) string {
	return "<code>" + self.content + "</code>"
}

type Link struct {
	*ElementBase

	text, uri string
}

func NewLink(text, uri string) *Link {
	return &Link {
		NewElementBase(),
		text,
		uri,
	}
}

func (self *Link) Feed(s string) error {
	return ImutableElementError
}

func (self *Link) Children() []Element {
	return nil
}

func (self *Link) String() string {
	return "[" + self.text + "](" + self.uri + ")"
}

func (self *Link) HTML(level int) string {
	return "<a href=\"" + self.uri + "\">" + self.text + "</a>"
}

type ImageLink struct {
	*Link
}

func NewImageLink(alt, uri string) *ImageLink {
	return &ImageLink{NewLink(alt, uri)}
}

func (self *ImageLink) String() string {
	return "[IMG: " + self.text + "](" + self.uri + ")"
}

func (self *ImageLink) HTML(level int) string {
	return "<a href=\"" + self.uri + "\"><img src=\"" + self.uri + "\" alt=\"" + self.text + "\"></a>"
}
