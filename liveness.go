
// will reports unhealthy after 15 sec
// will reports ready     after 20 sec

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	started := time.Now()
	http.HandleFunc("/readz", func(w http.ResponseWriter, r *http.Request) {
/*
		w.WriteHeader(200)
		data := (time.Since(started)).String()
		w.Write([]byte(data))
*/
		duration := time.Since(started)
		if duration.Seconds() > 20 {
			w.WriteHeader(200)
			w.Write([]byte("ok"))
		} else {
			w.WriteHeader(500)
			w.Write([]byte("error: getting ready"))
		}
	})
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		duration := time.Since(started)
		if duration.Seconds() > 15 {
			w.WriteHeader(500)
			w.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds())))
		} else {
			w.WriteHeader(200)
			w.Write([]byte("ok"))
		}
	})
	log.Fatal(http.ListenAndServe(":8765", nil))
}
