

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"math/rand"
)

func getenv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return fallback
    }
    return value
}

func main() {
        //fmt.Println("READINESS:", os.Getenv("READINESS"))
        // mode 0: default liveness
        //         will reports healtyh after 15 sec
        //         will reports ready   after 20 sec
        // mode 1: dies randomly, depending on READINESS_PERCENT
        //MODE := os.Getenv("MODE")
        MODE := getenv("MODE", "DEFAULT")
        //READINESS := os.Getenv("READINESS_PERCENT")
        //READINESS := getenv("READINESS_PERCENT", "DEFAULT")

	started := time.Now()


        switch MODE {
        // wie in tag 1
	case "DEFAULT":
		fmt.Println("MODE: DEFAULT")

		http.HandleFunc("/readz", func(w http.ResponseWriter, r *http.Request) {
			duration := time.Since(started)
			if duration.Seconds() > 20 {
				w.WriteHeader(200)
				w.Write([]byte("ok"))
			} else {
				w.WriteHeader(500)
				w.Write([]byte(fmt.Sprintf("error, getting ready: %v", duration.Seconds())))
			}
		})
		http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
			duration := time.Since(started)
			if duration.Seconds() > 15 {
				w.WriteHeader(200)
				w.Write([]byte("ok"))
			} else {
				w.WriteHeader(500)
				w.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds())))
			}
		})

	case "RANDOMFAIL":
        	rand.Seed(time.Now().UnixNano())
        	FAILURE_RATE := getenv("FAILURE_RATE", "50")
        	failure_rate, _ := strconv.Atoi(FAILURE_RATE)
        	//f := rand.Float64() * f1
        	rand := rand.Intn(100)
	        fmt.Println("random value:", rand)
	        fmt.Println("failure_rate:", failure_rate)


		if rand < failure_rate {
			fmt.Println("RANDOMFAIL: failed")
			os.Exit(1)
		} else {
			fmt.Println("RANDOMFAIL: passed")
			os.Exit(0)
		}

	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", MODE)
	}



        /*
	http.HandleFunc("/readz", func(w http.ResponseWriter, r *http.Request) {
		duration := time.Since(started)
		if duration.Seconds() > 20 {
			w.WriteHeader(200)
			w.Write([]byte("ok"))
		} else {
			w.WriteHeader(500)
			w.Write([]byte(fmt.Sprintf("error, getting ready: %v", duration.Seconds())))
		}
	})
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		duration := time.Since(started)
		if duration.Seconds() > 15 {
			w.WriteHeader(200)
			w.Write([]byte("ok"))
		} else {
			w.WriteHeader(500)
			w.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds())))
		}
	})
        */

	log.Fatal(http.ListenAndServe(":8765", nil))
}
