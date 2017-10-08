package tool

import (
	"strings"

	"github.com/x-tool/console"
	"github.com/x-tool/uniqueId"
)

var Console = console.New()

func NewUniqueId() string {
	return uniqueId.NewUniqueId()
}

func Panic(errType string, err error) {
	Console.Panic(errType, err)
}

func ReplaceStrings(raw string, m []string) (s string) {
	if len(m)%2 != 0 {
		return ""
	}
	s = raw
	for i, _ := range m {
		if i%2 != 0 {
			continue
		}
		s = strings.Replace(s, m[i], m[i+1], -1)
	}
	return s
}
