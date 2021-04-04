package main

import (
	//	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"strconv"
	"math/rand"

	//	"io"
	//	"os/signal"
	//	"runtime"
	"strings"
	//	"sync/atomic"
	//	"time"
)

type key int

const (
	requestIDKey key     = 0
	VERSION      float64 = 0.2
)

var (
	port         = 8085
	httpListener = flag.String("listen", ":"+"8085", "listen port")
	//httpListener = flag.String("listen", ":" + port, "listen port")
	htmlDocument = flag.String("document", "./index.html", "default document")

	access_logger *log.Logger
	//accessLogfile = flag.String("access_log", "./access_log", "Pfad zur access_log")
	accessLogfile = "/access_log"
	htmlFallback  = []byte("<html><head><title>404</title></head><body><h1>404</body></html>")

	healthy int32

	cdr_csv = "/var/log/asterisk/cdr-csv/Master.csv"
	records [][]string
	tmprec  = [][]string{
		{"bla", "jkl"},
		{"bla2", "jkl2"},
	}
)

func init() {
	access_log, err := os.OpenFile(accessLogfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		os.Exit(1)
	}
	access_logger = log.New(access_log, "", log.LstdFlags)

	r := csv.NewReader(strings.NewReader(cdr_csv))
	records, err = r.ReadAll()
	fmt.Printf("bla: %v", len(records))

	//	records = records[:len(records)-1]

}

func main() {

        fmt.Println("READINESS:", os.Getenv("READINESS"))
        READINESS := os.Getenv("READINESS")

        f1, _ := strconv.ParseFloat(READINESS, 64)


        rand.Seed(time.Now().UnixNano())

        f := rand.Float64() * f1
        fmt.Println("f:", f)

        if(f < 50) {
                os.Exit(3)
        }

	http.HandleFunc("/", cdr)

	log.Fatal(http.ListenAndServe(":8765", nil))
}

func cdr(w http.ResponseWriter, r *http.Request) {

	/*
		// Open the file
		csvfile, err := os.Open("input.csv")
		if err != nil {
			log.Fatalln("Couldn't open the csv file", err)
		}

		// Parse the file
		r1 := csv.NewReader(csvfile)
		//r := csv.NewReader(bufio.NewReader(csvfile))

		// Iterate through the records
		for {
			// Read each record from csv
			record, err := r1.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Question: %s Answer %s\n", record[0], record[1])
		}
	*/

	// 1. Open the file
	recordFile, err := os.Open(cdr_csv)
	if err != nil {
		fmt.Println("An error encountered ::", err)
	}
	// 2. Initialize the reader
	reader := csv.NewReader(recordFile)
	// 3. Read all the records
	records, _ := reader.ReadAll()

	records2 := records[len(records)-10 : len(records)]
	// 4. Iterate through the records as you wish
	fmt.Println(records2)

	//fmt.Fprintf(w, "bla" + records2[2][2])

	fmt.Fprintf(w, "<html><head><title>404</title></head><body><h1>Asterisk CDRs")
	fmt.Fprintf(w, "<table>")

	fmt.Fprintf(w, " <th>")
	fmt.Fprintf(w, " <td style=\"text-align: right\">Zeit</td>")
	fmt.Fprintf(w, " <td style=\"text-align: right\">Caller ID Number</td>")
	fmt.Fprintf(w, " <td style=\"text-align: right\">nach</td>")
	fmt.Fprintf(w, " <td style=\"text-align: right\">nach</td>")
	fmt.Fprintf(w, " <td style=\"text-align: right\">Gesamtdauer</td>")
	fmt.Fprintf(w, " <td style=\"text-align: right\">Gesprächsdauer</td>")
	fmt.Fprintf(w, " </th>")
	fmt.Fprintf(w, "\n")

	for i := len(records2) - 1; i >= 0; i-- {
		fmt.Fprintf(w, " <tr>")
		fmt.Fprintf(w, records2[i][2])

		//  $field = trim($my_array[2], '"');
		fmt.Fprintf(w, " <td style=\"text-align: right\">")
		fmt.Fprintf(w, records2[i][9])
		fmt.Fprintf(w, " </td>")

		fmt.Fprintf(w, " <td style=\"text-align: right\">")
		fmt.Fprintf(w, records2[i][1])
		fmt.Fprintf(w, " </td>")

		fmt.Fprintf(w, " <td style=\"text-align: right\">")
		fmt.Fprintf(w, records2[i][2])
		fmt.Fprintf(w, " </td>")

		fmt.Fprintf(w, " <td style=\"text-align: right\">")
		fmt.Fprintf(w, records2[i][8])
		fmt.Fprintf(w, " </td>")

		fmt.Fprintf(w, " <td style=\"text-align: right\">")
		fmt.Fprintf(w, records2[i][12])
		fmt.Fprintf(w, " </td>")

		fmt.Fprintf(w, " <td style=\"text-align: right\">")
		fmt.Fprintf(w, records2[i][13])
		fmt.Fprintf(w, " </td>")

		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, " </tr>")
	}

	fmt.Fprintf(w, "</body></html>")
	//csvReader()

}

//func cdr(logger *log.Logger) http.Handler {
func cdrALT(w http.ResponseWriter, r *http.Request) {
	//f, _ := os.Open(cdr_csv)

	//rec := csv.NewReader(f)

	//records, _ := r.readAll()

	//	document := []byte("<html><head><title>404</title></head><body><h1>404 hamma neda</body></html>")

	//fmt.Fprintf(w, string(document))
	//fmt.Fprintf(w, records[0][0])
	//fmt.Fprintf(w, records[1][1])

	//fmt.Fprintf(w, records[1][1])

}
