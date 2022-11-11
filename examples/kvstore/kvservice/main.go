package main

import (
	"singleinst"
	_ "singleinst/examples/kvstore/kvservice/mvc/ctrl"
)

func main() {
	nanogo.Poll()
}
