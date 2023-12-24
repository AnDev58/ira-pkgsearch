package backend

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"

	goremote "github.com/IRA-Package-Manager/goremote/util"
	"github.com/andev58/ira-pkgsearch/backend/constants"
	controllers "github.com/andev58/ira-pkgsearch/backend/controllers/other"
	"github.com/andev58/ira-pkgsearch/backend/controllers/packages"
	"github.com/andev58/ira-pkgsearch/backend/util"
	"github.com/gorilla/mux"
)

func Run(defaultStage int, address string, certFile, keyFile string, pkgDir string) {
	router := mux.NewRouter()
	router.StrictSlash(true)

	stage := util.GetServerStage(defaultStage)
	if stage == constants.STAGE_DEV {
		router.HandleFunc("/dev/version", controllers.VersionPageHandler).Methods("GET")
	}

	// Routing packages
	pkgServer, err := packages.NewServer(pkgDir, stage)
	if err != nil {
		log.Fatal(err)
	}
	pkgSubrouter := router.PathPrefix("/pkg").Subrouter()
	pkgSubrouter.HandleFunc("/", pkgServer.CreatePackageHandler).Methods("POST")

	// IRA API
	router.PathPrefix("/").Handler(goremote.NewRemoteMux(pkgDir))

	listenAndServeTLS(setupServer(address, router), certFile, keyFile, stage)
}

func listenAndServeTLS(srv *http.Server, certFile, keyFile string, stage int) {
	listener, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		log.Fatal(err)
	}

	if stage == constants.STAGE_PROD {
		log.Println("Succesifully started the production server")
	} else {
		realAddr := srv.Addr
		if realAddr[0] == ':' {
			realAddr = "https://localhost" + srv.Addr
		}

		log.Printf("Server running on %s", realAddr)
	}
	if err := srv.ServeTLS(listener, certFile, keyFile); err != nil {
		log.Fatal(err)
	}
}

func setupServer(addr string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: handler,
		TLSConfig: &tls.Config{
			MinVersion:               tls.VersionTLS13,
			PreferServerCipherSuites: true,
		},
	}
}
