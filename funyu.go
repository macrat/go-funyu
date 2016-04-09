package funyu

import (
	"bufio"
	"fmt"
	"io"
)

func Parse(r io.Reader) (Metadata, *Funyu, error) {
	funyu := NewFunyu()
	md := NewMetadata()
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if err := md.Feed(scanner.Text()); err == EndOfMetadata {
			break
		} else if err != nil {
			return nil, nil, err
		}
	}
	for scanner.Scan() {
		if err := funyu.Feed(scanner.Text()); err != nil {
			return nil, nil, err
		}
	}
	return md, funyu, nil
}

type Element interface {
	fmt.Stringer

	HTML(level int) string
	Feed(s string) error
}

type ElementBase struct{}

func NewElementBase() *ElementBase {
	return &ElementBase{}
}

type ParentElement interface {
	Element

	Children() []Element
	append(elm Element)
}

type ParentElementBase struct {
	*ElementBase

	children []Element
}

func NewParentElementBase() *ParentElementBase {
	return &ParentElementBase{
		NewElementBase(),
		make([]Element, 0),
	}
}

func (self *ParentElementBase) Children() []Element {
	return self.children
}

func (self *ParentElementBase) append(elm Element) {
	self.children = append(self.children, elm)
}

func (self *ParentElementBase) StringList() []string {
	s := make([]string, len(self.children))
	for i, x := range self.Children() {
		s[i] = x.String()
	}
	return s
}

func (self *ParentElementBase) HTMLList(level int) []string {
	s := make([]string, len(self.children))
	for i, x := range self.Children() {
		s[i] = x.HTML(level + 1)
	}
	return s
}
