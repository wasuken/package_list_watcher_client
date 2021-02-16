package main

import (
	"fmt"
	"github.com/wasuken/package_list_watcher_client/apt"
	"github.com/wasuken/package_list_watcher_client/pacman"
	"github.com/wasuken/package_list_watcher_client/send"
	"os"
)

func main() {
	pacMan := os.Args[1]
	if pacMan == "pacman" {
		err, info := pacman.GetInfo()
		if err != nil {
			panic(err)
		}
		send.SendSrv(info)
	} else if pacMan == "apt" {
		apt.GetInfo()
		// err, info := apt.GetInfo()
		// if err != nil {
		// 	panic(err)
		// }
		// send.SendSrv(info)
	} else {
		fmt.Println("Ha?")
	}
}
