package funyu

import (
	"fmt"
	"strings"
)

var (
	ImutableElementError = fmt.Errorf("this element is imutable")
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

func (self *Line) Feed(s string) error {
	self.append(NewString(s))
	return nil
}

type String struct {
	*NoChildElement

	content string
}

func NewString(content string) *String {
	return &String{
		NewNoChildElement(),
		content,
	}
}

func (self *String) String() string {
	return self.content
}

func (self *String) HTML(level int) string {
	return self.content
}
