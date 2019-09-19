package main

import (
	"github.com/charz/goInfo"
)

func main() {
	gi := goInfo.GetInfo()
	gi.VarDump()
}
