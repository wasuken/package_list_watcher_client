package pacman

import (
	"fmt"
	"github.com/goulash/pacman"
	"github.com/goulash/pacman/aur"
	"github.com/goulash/pacman/pkgutil"
	"strings"
)

func GetInfo() {
	pkgs, err := pacman.ReadLocalDatabase(func(er error) error {
		panic(er)
	})
	if err != nil {
		panic(err)
	}
	pkgMap := pkgutil.MapPkg(pkgs, func(pkg pacman.AnyPackage) string {
		return pkg.Pkg().PkgName()
	})
	pkgs = pkgs[0:10]
	var originPkgs pacman.Packages
	for _, pkg := range pkgs {
		p, err := pacman.Read(pkg.PkgName())
		p2, err2 := aur.Read(pkg.PkgName())
		if err != nil && err2 != nil {
			fmt.Println("package not found: " + pkg.PkgName())
		} else if err == nil {
			pkgMap[p.PkgName()] = p
		} else {
			pkgMap[p2.PkgName()] = p2.Pkg()
		}
	}
	originPkgMap := pkgutil.MapPkg(originPkgs, func(pkg pacman.AnyPackage) string {
		return pkg.Pkg().PkgName()
	})

	for k, pkg := range originPkgMap {
		fmt.Println("current: " + pkg.Name + " : " + pkg.Version)
		fmt.Println("		depends: " + strings.Join(pkg.Depends, ","))
		fmt.Println("		build date: " + pkg.BuildDate.Format("2006-01-02 15:04:05"))
		if _, ok := pkgMap[k]; ok {
			originPkg := pkgMap[k]
			opBuildTimeFmt := originPkg.BuildDate.Format("2006-01-02 15:04:05")
			fmt.Println("origin: " + originPkg.Name + " : " + originPkg.Version)
			fmt.Println("		depends: " + strings.Join(originPkg.Depends, ","))
			fmt.Println("		build date: " + opBuildTimeFmt)
		}
	}
}
