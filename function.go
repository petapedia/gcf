package gcf

import (
	"encoding/json"
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(datagedung)
}
