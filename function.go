package gcf

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/petapedia/peda"
)

func init() {
	functions.HTTP("PetaPedia", petaPedia)
}

func petaPedia(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, peda.GCFHandler("MONGOULBI", "petapedia", "petapedia"))
}
