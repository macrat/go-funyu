package funyu

import (
	"bufio"
	"fmt"
	"io"
)

func Parse(r io.Reader) (*Funyu, error) {
	funyu := NewFunyu()
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if err := funyu.Feed(scanner.Text()); err != nil {
			return nil, err
		}
	}
	return funyu, nil
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
