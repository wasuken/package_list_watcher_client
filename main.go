package main

import (
	"fmt"
	"github.com/wasuken/package_list_watcher_client/apt"
	"github.com/wasuken/package_list_watcher_client/config"
	"github.com/wasuken/package_list_watcher_client/pacman"
	"github.com/wasuken/package_list_watcher_client/send"
	"os"
)

func main() {
	pacMan := os.Args[1]
	err, config := config.ReadConfig()
	if err != nil {
		panic(err)
	}
	if pacMan == "pacman" {
		err, info := pacman.GetInfo()
		if err != nil {
			panic(err)
		}
		send.SendSrv(info, config.URL)
	} else if pacMan == "apt" {
		apt.GetInfo()
		err, info := apt.GetInfo()
		if err != nil {
			panic(err)
		}
		send.SendSrv(info, config.URL)
	} else {
		fmt.Println("Ha?")
	}
}
