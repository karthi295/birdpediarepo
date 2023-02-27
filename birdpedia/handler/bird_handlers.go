package handler

import (
	"encoding/json"
	"fmt"

	"net/http"
)

type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

var birds []Bird

func GetBirdHandler(w http.ResponseWriter, r *http.Request) {
	//Convert the "birds" variable to json
	birdListBytes, err := json.Marshal(birds)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(birdListBytes)
}

func CreateBirdHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new instance of Bird
	bird := Bird{}

	// We send all our data as HTML form data
	// the `ParseForm` method of the request, parses the
	// form values
	err := r.ParseForm()

	// In case of any error, we respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the bird from the form info
	bird.Species = r.Form.Get("species")
	bird.Description = r.Form.Get("description")

	// Append our existing list of birds with a new entry
	birds = append(birds, bird)

	//Finally, we redirect the user to the original HTMl page
	// (located at `/assets/`), using the http libraries `Redirect` method
	http.Redirect(w, r, "/assets/", http.StatusFound)
}
