package packages

import (
	packageModel "github.com/andev58/ira-pkgsearch/backend/models/packages"
)

// Server is a wrapper for PackageStore used for controllers
type Server struct {
	store *packageModel.PackageStore
	stage int
}

// NewServer creates new Server
func NewServer(pkgDir string, stage int) (*Server, error) {
	store, err := packageModel.NewPackageStore(pkgDir)
	return &Server{store, stage}, err
}
