package apt

import (
	"os"

	"github.com/arduino/go-apt-client"
	"github.com/wasuken/package_list_watcher_client/send"
	"runtime"
)

func GetInfo() (error, send.SendInfo) {
	apt.CheckForUpdates()
	allPkgs, err := apt.List()
	if err != nil {
		panic(err)
	}
	pkgInfos := []send.PackageInfo{}
	for _, pkg := range allPkgs {
		pack := send.PackageInfo{Name: pkg.Name, Version: pkg.Version}
		pkgInfos = append(pkgInfos, pack)
	}
	name, err := os.Hostname()
	if err != nil {
		return err, send.SendInfo{}
	}
	return nil, send.SendInfo{
		Name:        name,
		PackManType: "apt",
		Arch:        runtime.GOARCH,
		Packs:       pkgInfos,
	}
}
