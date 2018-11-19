package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Variable
	listeningPort := 9099

	// MAp the function handler to the route in the default server multiplexer
	http.HandleFunc("/hellogo", helloGoHandler)

	//Just a log
	log.Printf("HTTP server starting at %v", listeningPort)

	// Start the server on the listening port. Eventully rise an error if something occour.
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", listeningPort), nil))

}

// Function handler implementation.
func helloGoHandler(w http.ResponseWriter, r *http.Request) {
	response := helloResponse{Message: "I am a string inside a Json"}
	data, err := json.Marshal(response)

	if err != nil {
		panic("Oh no!")
	}

	fmt.Fprint(w, string(data))
}

type helloResponse struct {
	Message string
}
