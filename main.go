package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Output)

	// Use an explicit http.Server with timeouts to mitigate slow-client
	// (Slowloris) denial-of-service attacks. http.ListenAndServe has no
	// timeout support (gosec G114).
	srv := &http.Server{
		Addr:              ":8000",
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	// Bind to a port and serve our router
	log.Fatal(srv.ListenAndServe())
}

// Output generates numbers randomly
func Output(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Random number is : %v", randomNumbers())
}

func randomNumbers() int {
	// Use crypto/rand instead of math/rand so the output is not predictable
	// (gosec G404 / CWE-338).
	n, err := rand.Int(rand.Reader, big.NewInt(1000))
	if err != nil {
		log.Printf("failed to generate random number: %v", err)
		return 0
	}
	return int(n.Int64())
}
