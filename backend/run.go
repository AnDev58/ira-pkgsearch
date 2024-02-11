// Package backend is an entry point for a full server
package backend

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"strings"

	goremote "github.com/IRA-Package-Manager/goremote/util"
	"github.com/andev58/ira-pkgsearch/backend/constants"
	"github.com/andev58/ira-pkgsearch/backend/controllers/auth"
	controllers "github.com/andev58/ira-pkgsearch/backend/controllers/other"
	"github.com/andev58/ira-pkgsearch/backend/controllers/packages"
	"github.com/andev58/ira-pkgsearch/backend/db"
	"github.com/andev58/ira-pkgsearch/backend/util"
	"github.com/gorilla/mux"
)

// Run starts server whih stage defaultStage (if nothing said in environmental variables) on socket address with a specified
// TLS certFile and keyFile. For storing packages it use directory pkgDir
func Run(defaultStage int, http bool, httpAddr string, tls bool, tlsAddr string, certFile, keyFile string, pkgDir string, dbInfo db.DBInfo) {
	router := mux.NewRouter()
	router.StrictSlash(true)

	db, err := db.Connect(util.GetDBInfo(dbInfo))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to DB")

	stage := util.GetServerStage(defaultStage)
	if stage == constants.StageDev {
		router.HandleFunc("/dev/version", controllers.VersionPageHandler).Methods("GET")
	}

	// Routing packages
	pkgServer, err := packages.NewServer(pkgDir, stage)
	if err != nil {
		log.Fatal(err)
	}
	pkgServer.Route(router.PathPrefix("/pkg").Subrouter())

	// Routing authentification
	authServer := auth.NewServer(db)
	authServer.Route(router.PathPrefix("/usr").Subrouter())

	// IRA API
	router.PathPrefix("/").Handler(goremote.NewRemoteMux(pkgDir))

	if tls {
		go listenAndServe(setupServer(tlsAddr, router), certFile, keyFile, stage, true)
	}
	if http {
		go listenAndServe(setupServer(httpAddr, router), certFile, keyFile, stage, false)
	}

	for {

	}
}

func listenAndServe(srv *http.Server, certFile, keyFile string, stage int, tls bool) {
	listener, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		log.Fatal(err)
	}

	protocolName := "HTTP"
	if tls {
		protocolName += "S"
	}
	if stage == constants.StageProd {
		log.Println("Succesifully started the " + protocolName + " production server")
	} else {
		realAddr := srv.Addr
		if realAddr[0] == ':' {
			realAddr = strings.ToLower(protocolName) + "://localhost" + srv.Addr
		}

		log.Printf("%s server running on %s", protocolName, realAddr)
	}
	if tls {
		if err := srv.ServeTLS(listener, certFile, keyFile); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := srv.Serve(listener); err != nil {
			log.Fatal(err)
		}
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
