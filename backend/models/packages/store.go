// Package packages describes package storage which is a modification of [github.com/IRA-Package-Manager/goremote/util].Package
// It also specifies rules of storaging and naming packages
package packages // import models/packages

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"slices"
	"sync"

	goremote "github.com/IRA-Package-Manager/goremote/util"
)

// PackageStore present a local package storage
type PackageStore struct {
	sync.Mutex

	pkgsDb map[string][]goremote.Package
	pkgDir string
}

// NewPackageStore creates a new package storage
// pkgDir is a directory where PackageStore exists or need to be created
// Returns created PackageStore or error
func NewPackageStore(pkgDir string) (*PackageStore, error) {
	ps := &PackageStore{pkgDir: pkgDir, pkgsDb: make(map[string][]goremote.Package)}

	if _, err := os.Stat(pkgDir); os.IsNotExist(err) {
		err = os.MkdirAll(pkgDir, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	files, err := filepath.Glob(filepath.Join(pkgDir, "*.json"))
	if err != nil {
		return nil, err
	}

	if files == nil {
		return ps, nil
	}

	for _, file := range files {
		var pkg goremote.Package

		data, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(data, &pkg)
		if err != nil {
			return nil, err
		}

		ps.pkgsDb[pkg.Name] = append(ps.pkgsDb[pkg.Name], pkg)
	}
	return ps, nil
}

// CreatePackage add IPKG in PackageStore
// pkg is a goremote/util.Package containing configuration
// ipkg is an IPKG file need to be saved
func (ps *PackageStore) CreatePackage(pkg goremote.Package, ipkg []byte) error {
	ps.Lock()
	defer ps.Unlock()

	if ps.Exists(pkg.Name, pkg.Version) {
		return fmt.Errorf("package %s (version %s) already exists", pkg.Name, pkg.Version)
	}
	for _, val := range pkg.Dependencies {
		if !ps.Exists(val.Name, val.Version) {
			return fmt.Errorf("dependency %s (version %s) do not exists", pkg.Name, pkg.Version)
		}
	}

	pathToIpkg := filepath.Join(ps.pkgDir, pkg.Name)
	err := os.Mkdir(pathToIpkg, os.ModePerm)
	if os.IsExist(err) {
		if _, err := os.Stat(filepath.Join(pathToIpkg, pkg.FileName)); !os.IsNotExist(err) {
			return fmt.Errorf("IPKG %s already exists", pkg.FileName)
		}
	} else if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(pathToIpkg, pkg.FileName), ipkg, 0666)
	if err != nil {
		return err
	}

	jsonPkg, err := json.Marshal(pkg)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(ps.pkgDir, fmt.Sprintf("%s-v%s.json", pkg.Name, pkg.Version)), jsonPkg, os.ModePerm)
	if err != nil {
		return err
	}

	ps.pkgsDb[pkg.Name] = append(ps.pkgsDb[pkg.Name], pkg)

	return nil
}

// Exists check if package (name, version) exists in PackageStore
func (ps *PackageStore) Exists(name, version string) bool {
	if pkgs, ok := ps.pkgsDb[name]; ok {
		return slices.IndexFunc(
			pkgs,
			func(p goremote.Package) bool {
				return p.Version == version
			}) != -1
	}

	return false
}

// CreateFileName generates unique server file name for package (name, version)
func (ps *PackageStore) CreateFileName(name, version string) string {
	return fmt.Sprintf("%s-v%s-%d.ipkg", name, version, (len(name)+len(version))*rand.Int())
}
