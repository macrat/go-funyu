package funyu

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	modeInitial Mode = iota
	modeSection
	modePostScript
	modeCodeBlock
	modeEmbeddedHTML
	modeParagraph
)

type Mode int

type BlockElement interface {
	ParentElement
}

type BlockElementBase struct {
	*ParentElementBase

	mode Mode
}

func NewBlockElementBase() *BlockElementBase {
	return &BlockElementBase{
		NewParentElementBase(),
		modeInitial,
	}
}

func (self *BlockElementBase) Feed(s string) error {
	if self.mode == modeCodeBlock {
		if strings.HasPrefix(s, "```") {
			self.mode = modeInitial
			return nil
		} else if s != "" && !strings.HasPrefix(s, "\t") {
			return fmt.Errorf("Code Block must be close.")
		}
	}
	if self.mode == modeEmbeddedHTML {
		if strings.HasPrefix(s, ")))") {
			self.mode = modeInitial
			return nil
		} else if s != "" && !strings.HasPrefix(s, "\t") {
			return fmt.Errorf("Embedded HTML must be close.")
		}
	}

	if strings.HasPrefix(s, "-- ") {
		self.mode = modeSection
		self.append(NewSection(strings.TrimSpace(s[3:])))
	} else if strings.HasPrefix(s, "p.s. ") {
		self.mode = modePostScript
		self.append(NewPostScript(strings.TrimSpace(s[5:])))
	} else if strings.HasPrefix(s, "```") {
		self.mode = modeCodeBlock
		self.append(NewCodeBlock(strings.TrimSpace(s[3:])))
	} else if s == "(((" {
		self.mode = modeEmbeddedHTML
		self.append(NewEmbeddedHTML())
	} else if strings.TrimSpace(s) != ""  && (self.mode == modeInitial || self.mode != modeParagraph && !strings.HasPrefix(s, "\t")) {
		self.mode = modeParagraph
		self.append(NewParagraph())
		return self.children[len(self.children)-1].Feed(s)
	} else if self.mode == modeParagraph {
		if strings.TrimSpace(s) == "" {
			self.mode = modeInitial
		} else {
			return self.children[len(self.children)-1].Feed(s)
		}
	} else if len(s) > 0 {
		return self.children[len(self.children)-1].Feed(s[1:])
	} else if self.mode != modeInitial {
		return self.children[len(self.children)-1].Feed(s)
	}

	return nil
}

type Funyu struct {
	*BlockElementBase
}

func NewFunyu() *Funyu {
	return &Funyu{NewBlockElementBase()}
}

func (self *Funyu) String() string {
	return strings.Join(self.StringList(), "")
}

func (self *Funyu) HTML(level int) string {
	return "<article>\n" + strings.Join(self.HTMLList(level), "") + "</article>"
}

type Section struct {
	*BlockElementBase

	title string
}

func NewSection(title string) *Section {
	return &Section{
		NewBlockElementBase(),
		title,
	}
}

func (self *Section) String() string {
	if len(self.children) > 0 {
		return strings.TrimRight("-- "+self.title+"\n\t"+strings.Join(strings.Split(strings.Join(self.StringList(), ""), "\n"), "\n\t"), " \t\n") + "\n\n"
	} else {
		return strings.TrimRight("-- "+self.title, " \t\n") + "\n\n"
	}
}

func (self *Section) HTML(level int) string {
	var ls = "6"
	if 0 < level && level < 6 {
		ls = strconv.Itoa(level)
	}
	return "<section>\n<h" + ls + ">" + self.title + "</h" + ls + ">\n" + strings.Join(self.HTMLList(level), "") + "</section>\n"
}

type PostScript struct {
	*BlockElementBase

	date string
}

func NewPostScript(date string) *PostScript {
	return &PostScript{
		NewBlockElementBase(),
		date,
	}
}

func (self *PostScript) String() string {
	if len(self.children) > 0 {
		return strings.TrimRight("p.s. "+self.date+"\n\t"+strings.Join(strings.Split(strings.Join(self.StringList(), ""), "\n"), "\n\t"), " \t\n") + "\n\n"
	} else {
		return strings.TrimRight("p.s. "+self.date, " \t\n") + "\n\n"
	}
}

func (self *PostScript) HTML(level int) string {
	return "<ins>\n<span>p.s. <date>" + self.date + "</date></span>\n" + strings.Join(self.HTMLList(level), "") + "</ins>\n"
}

type PlainBlock struct {
	*BlockElementBase
}

func NewPlainBlock() *PlainBlock {
	return &PlainBlock{NewBlockElementBase()}
}

func (self *PlainBlock) Feed(s string) error {
	self.append(NewString(s))
	return nil
}

type CodeBlock struct {
	*PlainBlock

	lang string
}

func NewCodeBlock(lang string) *CodeBlock {
	return &CodeBlock{
		NewPlainBlock(),
		lang,
	}
}

func (self *CodeBlock) String() string {
	if len(self.children) > 0 {
		return strings.TrimRight("``` "+self.lang+"\n\t"+strings.Join(strings.Split(strings.Join(self.StringList(), "\n"), "\n"), "\n\t"), " \t\n") + "\n```\n"
	} else {
		return "```\n```\n"
	}
}

func (self *CodeBlock) HTML(level int) string {
	return "<pre class=\"code\" data-language=" + self.lang + ">" + strings.Replace(strings.Replace(strings.Join(self.StringList(), "\n"), "<", "&lt;", -1), ">", "&gt;", -1) + "</pre>\n"
}

type EmbeddedHTML struct {
	*PlainBlock
}

func NewEmbeddedHTML() *EmbeddedHTML {
	return &EmbeddedHTML{NewPlainBlock()}
}

func (self *EmbeddedHTML) String() string {
	if len(self.children) > 0 {
		return strings.TrimRight("(((\n\t"+strings.Join(strings.Split(strings.Join(self.StringList(), "\n"), "\n"), "\n\t"), " \t\n") + "\n)))\n"
	} else {
		return "(((\n)))\n"
	}
}

func (self *EmbeddedHTML) HTML(level int) string {
	return strings.Join(self.StringList(), "\n") + "\n"
}

type Paragraph struct {
	*BlockElementBase
}

func NewParagraph() *Paragraph {
	return &Paragraph{NewBlockElementBase()}
}

func (self *Paragraph) Feed(s string) error {
	l := NewLine()
	self.append(l)
	return l.Feed(s)
}

func (self *Paragraph) String() string {
	return strings.Join(self.StringList(), "") + "\n"
}

func (self *Paragraph) HTML(level int) string {
	return "<p>\n" + strings.Join(self.HTMLList(level), "") + "</p>\n"
}
