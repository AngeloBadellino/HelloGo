package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Fast initialization of the variable name
	listeningPort := 9099

	// MAp the function handler to the route in the default server multiplexer
	http.HandleFunc("/hellogo", helloGoHandler)

	log.Printf("HTTP server starting at %v", listeningPort)

	// Start the server on the listening port. Eventully rise an error if something occour.
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", listeningPort), nil))

}

// Function handler implementation.
func helloGoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World\n")
}
