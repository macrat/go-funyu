package funyu

import (
	"fmt"
	"strings"
)

var (
	EndOfMetadata = fmt.Errorf("end of metadata")
)

type Metadata map[string]string

func NewMetadata() Metadata {
	return make(Metadata)
}

func (self Metadata) Feed(s string) error {
	if len(s) == 0 {
		return EndOfMetadata
	}

	ls := strings.SplitN(s, ":", 2)
	if len(ls) != 2 {
		return fmt.Errorf("Invalid format metadata: "+s)
	}

	self[strings.TrimSpace(ls[0])] = strings.TrimSpace(ls[1])

	return nil
}
