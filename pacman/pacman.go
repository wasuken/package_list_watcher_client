package pacman

import (
	"fmt"
	"github.com/goulash/pacman"
	"github.com/goulash/pacman/pkgutil"
	// "github.com/wasuken/package_list_watcher_client/send"
	"strings"
)

func GetInfo() {
	allPkgs, err := pacman.ReadAllSyncDatabases()
	if err != nil {
		panic(err)
	}
	localPkgs, err := pacman.ReadLocalDatabase(func(er error) error {
		panic(er)
	})
	if err != nil {
		panic(err)
	}
	allPkgMap := pkgutil.MapPkg(allPkgs, func(pkg pacman.AnyPackage) string {
		return pkg.Pkg().PkgName()
	})
	localPkgMap := pkgutil.MapPkg(localPkgs[0:20], func(pkg pacman.AnyPackage) string {
		return pkg.Pkg().PkgName()
	})

	for k, pkg := range localPkgMap {
		fmt.Println("# " + pkg.PkgName())
		fmt.Println("current: " + pkg.Name + " : " + pkg.Version)
		fmt.Println("		depends: " + strings.Join(pkg.Depends, ","))
		fmt.Println("		build date: " + pkg.BuildDate.Format("2006-01-02 15:04:05"))
		if _, ok := allPkgMap[k]; ok {
			originPkg := allPkgMap[k]
			opBuildTimeFmt := originPkg.BuildDate.Format("2006-01-02 15:04:05")
			fmt.Println("origin: " + originPkg.Name + " : " + originPkg.Version)
			fmt.Println("		depends: " + strings.Join(originPkg.Depends, ","))
			fmt.Println("		build date: " + opBuildTimeFmt)
			if pkg.Older(originPkg) {
				fmt.Println(pkg.PkgName() + " is old.")
			}
		}
	}
}
