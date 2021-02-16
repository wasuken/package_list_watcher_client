package pacman

import (
	"fmt"
	"os"
	"strings"

	"github.com/goulash/pacman"
	"github.com/goulash/pacman/pkgutil"
	"github.com/wasuken/package_list_watcher_client/send"
)

func GetInfo() (error, send.SendInfo) {
	allPkgs, err := pacman.ReadAllSyncDatabases()
	if err != nil {
		return err, send.SendInfo{}
	}
	localPkgs, err := pacman.ReadLocalDatabase(func(er error) error {
		panic(er)
	})
	if err != nil {
		return err, send.SendInfo{}
	}
	allPkgMap := pkgutil.MapPkg(allPkgs, func(pkg pacman.AnyPackage) string {
		return pkg.Pkg().PkgName()
	})
	localPkgMap := pkgutil.MapPkg(localPkgs[0:20], func(pkg pacman.AnyPackage) string {
		return pkg.Pkg().PkgName()
	})

	name, err := os.Hostname()
	if err != nil {
		return err, send.SendInfo{}
	}
	pkgInfos := []send.PackageInfo{}
	latestPkgInfos := []send.PackageInfo{}
	for k, pkg := range localPkgMap {
		pkgInfo := send.PackageInfo{
			Name:    pkg.PkgName(),
			Version: pkg.Version,
			Date:    pkg.BuildDate.Format("2006-01-02 15:04:05")}
		pkgInfos = append(pkgInfos, pkgInfo)
		if _, ok := allPkgMap[k]; ok {
			latestPkg := allPkgMap[k]
			latestPkgInfo := send.PackageInfo{
				Name:    latestPkg.PkgName(),
				Version: latestPkg.Version,
				Date:    latestPkg.BuildDate.Format("2006-01-02 15:04:05")}
			latestPkgInfos = append(latestPkgInfos, latestPkgInfo)
		}
	}
	return nil, send.SendInfo{
		Name:        name,
		PackManType: "pacman",
		Packs:       latestPkgInfos,
		CurPacks:    pkgInfos,
	}
}
