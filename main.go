package main

import (
	"fmt"
	"github.com/wasuken/package_list_watcher_client/apt"
	"github.com/wasuken/package_list_watcher_client/pacman"
	"os"
)

func main() {
	pacMan := os.Args[1]
	if pacMan == "pacman" {
		pacman.GetInfo()
	} else if pacMan == "apt" {
		apt.GetInfo()
	} else {
		fmt.Println("Ha?")
	}
}
