package main

import (
	"singleinst"
	_ "singleinst/examples/kvstore/userservice/mvc/ctrl"
)

func main() {
	nanogo.Poll()
}
