package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/h3th-IV/nest/databses"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Our Home Page")
}

var conStr string = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disabale", os.Getenv("SQL_HOST"), os.Getenv("SQL_PORT"), os.Getenv("SQL_NAME"), os.Getenv("SQL_PASSWORD"), os.Getenv("SQL_DBNAME"))

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("parse form error: %v", err)
	}

	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	password := r.FormValue("password")

	fmt.Fprintf(w, "Hello %v,\n Welcome to hello page", name)
	dB, err := databses.InitDB(conStr)
	if err != nil {
		log.Fatal(err)
	}
	dB.Query(`INSERT INTO users(name, password) VALUES($1, $2)`, name, password)
}

func main() {
	//intialize the get request to host
	// resp, err := http.Get("https://github.com/jim-nnamdi?tab=repositories")
	// if err != nil {
	// 	panic(err)
	// }
	// //close the body of response
	// defer resp.Body.Close()

	// //scan the resp and print output
	// scanner := bufio.NewScanner(resp.Body)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	fileServer := http.FileServer(http.Dir("./static"))
	//now the server
	//create new ServeMultiplexer
	mux := http.NewServeMux()

	//the new mux is used to route requst
	mux.Handle("/", fileServer)
	mux.HandleFunc("/form", formHandler)
	mux.HandleFunc("/hello", helloHandler)

	fmt.Println("Listeniing on :8090")
	err := http.ListenAndServe(":8090", mux)
	if err != nil {
		panic(err)
	}
}
