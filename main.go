package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/h3th-IV/nest/databses"
	"github.com/joho/godotenv"
)

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
	// router := http.NewServeMux()

	//create a new router
	router := mux.NewRouter()

	//the new mux is used to route requst
	router.Handle("/", fileServer)
	router.HandleFunc("/form", formHandler)
	router.HandleFunc("/hello", helloHandler)

	//route parameters used to capture dynamic values from incomin request
	//this is done with mux.Vars
	router.HandleFunc("/user/{id:[0-9]+}", userHandler)
	router.HandleFunc("/topics/{category}/{id:[0-9]+}", topicHandler)

	fmt.Println("Listening on :8090")
	err := http.ListenAndServe(":8090", router)
	if err != nil {
		panic(err)
	}
}

func topicHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category"]
	topicID := vars["id"]

	fmt.Fprintf(w, "Welcome to the %v page, you are under %v ", category, topicID)
}

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
	name := r.FormValue("username")
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

func userHandler(w http.ResponseWriter, r *http.Request) {
	Var := mux.Vars(r)
	userID := Var["id"]
	fmt.Fprintf(w, "Welcome user: %v", userID)
}
