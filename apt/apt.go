package apt

import (
	"fmt"
	"github.com/arduino/go-apt-client"
)

func GetInfo() {
	apt.CheckForUpdates()
	pkgs, err := apt.ListUpgradable()
	if err != nil {
		panic(err)
	}
	for _, pkg := range pkgs {
		fmt.Println(pkg.Name + ":" + pkg.Version)
		fmt.Println("	" + pkg.Status)
	}
}
