package tool

import (
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
