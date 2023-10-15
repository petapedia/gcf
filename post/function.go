package gcf

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/petapedia/peda"
)

func init() {
	functions.HTTP("PetaPedia", petaPediaPost)
}

func petaPediaPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://jscroot.github.io")
	fmt.Fprintf(w, peda.GCFPostHandler("PASETOPRIVATEKEY", "MONGOULBI", "petapedia", "user", r))

}
