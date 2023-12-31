package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/h3th-IV/nest/database"
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

	//serve static files
	//now the server
	//create new ServeMultiplexer(or better put router)
	// router := http.NewServeMux()

	//create a new router
	router := mux.NewRouter()
	//middleware
	router.Use(loggingMiddleware)

	//create file sever with mux
	router.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	router.HandleFunc("/form", formHandler)
	router.HandleFunc("/hello", helloHandler)

	// route parameters used to capture dynamic values from incomin request
	// this is done with mux.Vars
	// subrouters are used to organise and group related routes under a common url

	userRouter := router.PathPrefix("/login").Subrouter()
	userRouter.HandleFunc("/", userHandler).Methods("GET")
	userRouter.HandleFunc("/account/{password}", loginHAndler).Name("userAccount")
	userRouter.HandleFunc("/profile", userProfileHandler).Methods("GET").Name("userProfile")
	userRouter.HandleFunc("/{username:[a-zA-z0-9!@#$%^&*()-_+]+}", singleUserHandler).Name("username")

	//some kinda
	userRouter.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		routeName := route.GetName()
		fmt.Println(routeName)

		pathRegxp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("Path Regexp: ", pathRegxp)
		}
		hosTemplate, err := route.GetHostTemplate()
		if err == nil {
			fmt.Println(hosTemplate)
		}
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Println("route Methods include: ", methods)
		}
		variables, err := route.GetVarNames()
		if err == nil {
			fmt.Println("route variable include: ", variables)
		}

		fmt.Print("\n\n")
		return nil
	})

	domainRouter := router.PathPrefix("/domains").Subrouter()
	domainRouter.HandleFunc("/", domainHandler).Methods("GET")
	domainRouter.HandleFunc("/{category:[a-zA-z0-9]+}/{topic:[a-zA-z0-9]+}", domainTopicHandler).Methods("GET").Name("domains") //this uses named route
	// url, err := domainRouter.Get("domains").URL("category", "cyber Security", "topic", "Malwares")

	//we could also match with funtions
	accountRouter := router.PathPrefix("/accounts").Subrouter()
	accountRouter.HandleFunc("/logger", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is just a logger page")
	}).Methods("GET")

	accountRouter.MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		return r.ProtoMinor == 0
	})

	fmt.Println("Listening on :5500")
	server := &http.Server{
		Handler:           router,
		Addr:              "127.0.0.1:5500",
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome pls go the Login page -> /login/password \ncreate an account -> /form")
}

func loginHAndler(w http.ResponseWriter, r *http.Request) {
	var username string
	Vars := mux.Vars(r)
	passowrd := Vars["password"]

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dB, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dB.Close()

	query := `SELECT username FROM users WHERE password = ?`
	stmt, err := dB.Prepare(query)
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()
	useRow, err := stmt.Query(passowrd)
	if err != nil {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
	}
	defer useRow.Close()
	//check if user with password exist
	if useRow.Next() {
		useRow.Scan(&username)
		fmt.Fprintf(w, "Welcome %v, have successfully login", username)
	} else {
		http.Error(w, "User Account not Found", http.StatusUnauthorized)
	}
}

// middleware
func loggingMiddleware(next http.Handler) http.Handler {
	logger := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("url visited: ", r.URL, r.Method)
		next.ServeHTTP(w, r)
	})
	return logger
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
	dB, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dB.Close()
	//use prepatred statement
	query := `INSERT INTO users(username, password) VALUES(?, ?)`
	_, err = dB.Query(query, name, password)
	if err != nil {
		log.Fatal(err)
	}
}
