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
	router.Handle("/", fileServer).Methods("GET") //we can use the Methods method to add HTTP method consraints.
	router.HandleFunc("/form", formHandler).Methods("")
	router.HandleFunc("/hello", helloHandler)

	//route parameters used to capture dynamic values from incomin request
	//this is done with mux.Vars
	// router.HandleFunc("/user/{id:[0-9]+}", userHandler)
	// router.HandleFunc("/topics/{category}/{id:[0-9]+}", domainHandler)

	//subrouters are used to organise and group related routes under a common url
	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/", userHandler).Methods("GET")
	userRouter.HandleFunc("/profile", userProfileHandler).Methods("GET")
	userRouter.HandleFunc("/{username:[a-zA-z0-9]+}", singleUserHandler).Methods("GET")

	domainRouter := router.PathPrefix("/domains").Subrouter()
	domainRouter.HandleFunc("/", domainHandler).Methods("GET")
	domainRouter.HandleFunc("/{category:[a-zA-z0-9]+}/{topic:[a-zA-z0-9]+}", domainTopicHandler).Methods("GET")

	fmt.Println("Listening on :8090")
	err := http.ListenAndServe(":8090", router)
	if err != nil {
		panic(err)
	}
}

func userProfileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to your profile page")
}

func singleUserHandler(w http.ResponseWriter, r *http.Request) {
	Var := mux.Vars(r)
	userID := Var["username"]
	fmt.Fprintf(w, "Welcome user: %v", userID)
}

func domainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to our domain page\n- Cyber Security\n- Systems Backend Engineering")
}

func domainTopicHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category"]
	topic := vars["topic"]

	fmt.Fprintf(w, "Welcome to the %v page, you will learn the concept of %v ", category, topic)
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
	fmt.Fprintf(w, "welcome new user")
}
