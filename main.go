package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/h3th-IV/nest/databses"
	"github.com/joho/godotenv"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Our Hello Page")
}

// constring for mySQL looks like this
//user:password@tcp(host:port)/dbname

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("parse form error: %v", err)
	}

	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	password := r.FormValue("password")

	fmt.Fprintf(w, "Hello %v,\n Welcome to form page", name)

	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	dB, err := databses.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	//use prepatred statement
	query := `INSERT INTO users(username, password) VALUES(?, ?)`
	_, err = dB.Query(query, name, password)
	if err != nil {
		log.Fatal(err)
	}
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

	fmt.Println("Listening on :8090")
	err := http.ListenAndServe(":8090", mux)
	if err != nil {
		panic(err)
	}
}
