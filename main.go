package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("parse form error: %v", err)
	}

	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")

	fmt.Fprintf(w, "Hello %v,\n Welcome to hello page", name)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Home Page")
}
func main() {
	//intialize the get request to host
	resp, err := http.Get("https://github.com/jim-nnamdi?tab=repositories")
	if err != nil {
		panic(err)
	}
	//close the body of response
	defer resp.Body.Close()

	//scan the resp and print output
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//now the server
	//create new ServeMultiplexer
	mux := http.NewServeMux()

	//the new mux is used to route requst
	mux.HandleFunc("/form", formHandler)
	mux.HandleFunc("/", homeHandler)

	fmt.Println("Listeniing on :8090")
	err = http.ListenAndServe("0.0.0.0:8090", mux)
	if err != nil {
		panic(err)
	}
}
