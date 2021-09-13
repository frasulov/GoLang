package main

import (
	"configFileTest/config"
	"fmt"
)


func main() {

	config.Init()
	fmt.Println(config.Configuration)
	fmt.Println(config.Configuration.Server)
}
