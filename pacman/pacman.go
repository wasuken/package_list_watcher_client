package pacman

import (
	"os"
	"runtime"

	"github.com/goulash/pacman"
	"github.com/goulash/pacman/pkgutil"
	"github.com/wasuken/package_list_watcher_client/send"
)

func GetInfo() (error, send.SendInfo) {
	localPkgs, err := pacman.ReadLocalDatabase(func(er error) error {
		panic(er)
	})
	if err != nil {
		return err, send.SendInfo{}
	}
	localPkgMap := pkgutil.MapPkg(localPkgs, func(pkg pacman.AnyPackage) string {
		return pkg.Pkg().PkgName()
	})

	name, err := os.Hostname()
	if err != nil {
		return err, send.SendInfo{}
	}

	pkgInfos := []send.PackageInfo{}
	for _, pkg := range localPkgMap {
		pkgInfo := send.PackageInfo{
			Name:        pkg.PkgName(),
			Version:     pkg.Version,
			Description: pkg.Description}
		pkgInfos = append(pkgInfos, pkgInfo)
	}
	return nil, send.SendInfo{
		Name:        name,
		Arch:        runtime.GOARCH,
		PackManType: "pacman",
		Packs:       pkgInfos,
	}
}
