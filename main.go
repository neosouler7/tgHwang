package main

import (
	"fmt"

	"github.com/neosouler7/tgHwang/tg"
	"github.com/neosouler7/tgHwang/utils"
)

func main() {
	fmt.Println(utils.TgConfig())

	go tg.Start()
	for {

	}

}
