package packages

import (
	packageModel "github.com/andev58/ira-pkgsearch/backend/models/packages"
	"github.com/gorilla/mux"
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

// Route routes handlers with given subrouter
func (s *Server) Route(subrouter *mux.Router) {
	subrouter.HandleFunc("/", s.CreatePackageHandler).Methods("POST")
}
