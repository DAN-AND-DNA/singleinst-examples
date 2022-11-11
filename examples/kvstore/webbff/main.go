package main

import (
	"singleinst"
	_ "singleinst/examples/kvstore/webbff/mvc/ctrl"
)

func main() {
	nanogo.Poll()
}
