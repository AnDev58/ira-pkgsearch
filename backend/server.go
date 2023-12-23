package backend

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	goremote "github.com/IRA-Package-Manager/goremote/util"
	"github.com/andev58/ira-pkgsearch/backend/constants"
	"github.com/andev58/ira-pkgsearch/backend/routes"
	"github.com/andev58/ira-pkgsearch/backend/util"
)

func Run(defaultPort, defaultStage int) {
	stage := util.GetServerStage(defaultStage)

	if stage == constants.STAGE_DEV {
		http.HandleFunc("/dev/version", routes.VersionPageHandler)
	}

	pkgDir, _ := filepath.Abs("./pkgs")

	http.Handle("/", goremote.NewRemoteMux(pkgDir))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", util.GetPort(defaultPort)), nil))
}
