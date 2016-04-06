package funyu

import (
	"testing"
)

func TestMetadata(t *testing.T) {
	md := NewMetadata()

	if err := md.Feed("hoge: fuga"); err != nil {
		t.Error(err.Error())
	}

	if err := md.Feed("abc: def"); err != nil {
		t.Error(err.Error())
	}

	if err := md.Feed(""); err == nil {
		t.Error("Failed detect end of metadata")
	} else if err != EndOfMetadata {
		t.Error("Failed detect end of metadata: "+err.Error())
	}

	for k, v := range md {
		if k == "hoge" {
			if v != "fuga" {
				t.Error("unexpected value: "+k+"->"+v)
			}
		} else if k == "abc" {
			if v != "def" {
				t.Error("unexpected value: "+k+"->"+v)
			}
		} else {
			t.Error("unexpected key: "+k+"->"+v)
		}
	}
}
