package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

var bold func(a ...interface{}) string = color.New(color.Bold).SprintFunc()
var yellow func(a ...interface{}) string = color.New(color.FgYellow).SprintFunc()
var green func(a ...interface{}) string = color.New(color.FgGreen).SprintFunc()
var red func(a ...interface{}) string = color.New(color.FgRed).SprintFunc()
var white func(a ...interface{}) string = color.New(color.FgWhite).SprintFunc()

func logRequest(r *http.Request) {

	header := strings.Repeat("-", 60)
	log.Println(header)

	log.Printf("%v %v %v", bold(r.Method), bold(r.URL), r.Proto)

	for name, val := range r.Header {
		log.Printf("%v: %v", name, val)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading body")
	}
	log.Println(bold(string(body)))
}

func handler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)

	responseCode := http.StatusOK
	regex := regexp.MustCompile(`/code/(\d{3})`)
	matches := regex.FindStringSubmatch(r.URL.Path)
	if len(matches) > 1 {
		responseCode, _ = strconv.Atoi(matches[1])
	}

	if responseCode > 399 {
		log.Printf("Responding with %v", red(strconv.Itoa(responseCode)))
	} else if responseCode > 299 {
		log.Printf("Responding with %v", yellow(strconv.Itoa(responseCode)))
	} else {
		log.Printf("Responding with %v", green(strconv.Itoa(responseCode)))
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
	w.WriteHeader(responseCode)
	io.WriteString(w, "OK\n")
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
