package packages

import (
	"encoding/json"
	"fmt"
	"net/http"

	goremote "github.com/IRA-Package-Manager/goremote/util"
	"github.com/andev58/ira-pkgsearch/backend/util"
)

func (s *Server) CreatePackageHandler(w http.ResponseWriter, r *http.Request) {
	type RequestPackage struct {
		Name         string           `json:"name"`
		Version      string           `json:"version"`
		Dependencies []RequestPackage `json:"deps"`
		File         json.RawMessage  `json:"ipkg"`
	}

	if util.EnforceJSON(w, r) {
		return
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var reqPkg RequestPackage
	if err := decoder.Decode(&reqPkg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pkg := goremote.Package{
		Name:     reqPkg.Name,
		Version:  reqPkg.Version,
		FileName: s.store.CreateFileName(reqPkg.Name, reqPkg.Version),
	}
	if err := s.store.CreatePackage(pkg, []byte(reqPkg.File)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		fmt.Fprintln(w, "OK")
	}
}
