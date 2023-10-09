package gcf

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/petapedia/peda"
)

func init() {
	functions.HTTP("PetaPedia", petaPedia)
}

func petaPedia(w http.ResponseWriter, r *http.Request) {
	mconn := peda.SetConnection("MONGOULBI", "petapedia")           //dbname : petapedia
	datagedung := peda.GetAllBangunanLineString(mconn, "petapedia") //collection name : petapedia
	jsondatagedung, _ := json.Marshal(datagedung)
	fmt.Fprintf(w, string(jsondatagedung))

	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(datagedung)
}
