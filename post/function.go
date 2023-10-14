package gcf

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("PetaPedia", petaPediaPost)
}

func petaPediaPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://jscroot.github.io")
	var d struct {
		Username string `json:"username" bson:"username"`
		Password string `json:"password" bson:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		fmt.Fprintf(w, "error parsing application/json: "+err.Error())
	} else {
		fmt.Fprintf(w, "Hai "+err.Error())
	}

}
