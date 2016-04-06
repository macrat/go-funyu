package funyu

import (
	"testing"
)

func TestLine(t *testing.T) {
	l := NewLine()
	if e := l.Feed("[link](uri) <<[[this]] is {{test}}>>"); e != nil {
		t.Error("Failed parse: " + e.Error())
	} else if h := l.HTML(0); h != "<a href=\"uri\">link</a> <em><strong>this</strong> is <code>test</code></em><br>\n" {
		t.Error("Failed parse line: " + h)
	}
}

func TestKeyword(t *testing.T) {
	l := NewLine()
	if e := l.Feed("this is [[test]]."); e != nil {
		t.Error("Failed parse: " + e.Error())
	} else if h := l.HTML(0); h != "this is <strong>test</strong>.<br>\n" {
		t.Error("Failed parse simple keyword: " + h)
	}

	l = NewLine()
	if e := l.Feed("[[this is [[[[test]].]]]]"); e != nil {
		t.Error("Failed parse: " + e.Error())
	} else if h := l.HTML(0); h != "<strong>this is <strong><strong>test</strong>.</strong></strong><br>\n" {
		t.Error("Failed parse nested keyword: " + h)
	}
}

func TestEmphasis(t *testing.T) {
	l := NewLine()
	if e := l.Feed("this is <<emphasis>>."); e != nil {
		t.Error("Failed parse: " + e.Error())
	} else if h := l.HTML(0); h != "this is <em>emphasis</em>.<br>\n" {
		t.Error("Failed parse simple emphasis: " + h)
	}

	l = NewLine()
	if e := l.Feed("<<<<this>> is <<emphasis>>.>>"); e != nil {
		t.Error("Failed parse: " + e.Error())
	} else if h := l.HTML(0); h != "<em><em>this</em> is <em>emphasis</em>.</em><br>\n" {
		t.Error("Failed parse nested emphasis: " + h)
	}
}

func TestCode(t *testing.T) {
	l := NewLine()
	if e := l.Feed("this is {{source code}}."); e != nil {
		t.Error("Failed parse: " + e.Error())
	} else if h := l.HTML(0); h != "this is <code>source code</code>.<br>\n" {
		t.Error("Failed parse simple code: " + h)
	}

	l = NewLine()
	if e := l.Feed("{{this is {{source code}}.}}"); e != nil {
		t.Error("Failed parse: " + e.Error())
	} else if h := l.HTML(0); h != "<code>this is {{source code}}.</code><br>\n" {
		t.Error("Failed parse nested code: " + h)
	}
}

func TestLink(t *testing.T) {
	l := NewLine()
	if e := l.Feed("this is [link](/)"); e != nil {
		t.Error("Failed parse: " + e.Error())
	} else if h := l.HTML(0); h != "this is <a href=\"/\">link</a><br>\n" {
		t.Error("Failed parse simple link: " + h)
	}

	l = NewLine()
	if e := l.Feed("[this](/test/) is [link](/)"); e != nil {
		t.Error("Failed parse: " + e.Error())
	} else if h := l.HTML(0); h != "<a href=\"/test/\">this</a> is <a href=\"/\">link</a><br>\n" {
		t.Error("Failed parse multiple link: " + h)
	}
}

func TestImageLink(t *testing.T) {
	l := NewLine()
	if e := l.Feed("this is [IMG: image](/path/to/image.png)"); e != nil {
		t.Error("Failed parse: " + e.Error())
	} else if h := l.HTML(0); h != "this is <a href=\"/path/to/image.png\"><img src=\"/path/to/image.png\" alt=\"image\"></a><br>\n" {
		t.Error("Failed parse simple image link: " + h)
	}

	l = NewLine()
	if e := l.Feed("[IMG: this](a) is [IMG: image](b)"); e != nil {
		t.Error("Failed parse: " + e.Error())
	} else if h := l.HTML(0); h != "<a href=\"a\"><img src=\"a\" alt=\"this\"></a> is <a href=\"b\"><img src=\"b\" alt=\"image\"></a><br>\n" {
		t.Error("Failed parse multiple image link: " + h)
	}
}
