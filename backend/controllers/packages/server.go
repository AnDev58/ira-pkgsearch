package packages

import (
	packageModel "github.com/andev58/ira-pkgsearch/backend/models/packages"
)

type Server struct {
	store *packageModel.PackageStore
	stage int
}

func NewServer(pkgDir string, stage int) (*Server, error) {
	store, err := packageModel.NewPackageStore(pkgDir)
	return &Server{store, stage}, err
}
