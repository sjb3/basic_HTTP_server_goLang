package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"github.com/julienschmidt/httprouter"
)

func ServeHTML(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	file := r.URL.Path //file := "/a/welcome" localhost:3000/a/welcome

	filename := file[1:] //filename :="a/*wildcard"

	if len(filename) == 0 {
		Home(w, r, nil)
		fmt.Println("requested home page..")

	} else if len(filename) >= 3 && filename[0:2] == "a" {
		http.ServeFile(w, r, filename[2:]+".html")
		fmt.Println("requested", filename[2:], "page...")

	} else {
		Home(w, r, nil)
		fmt.Println("requested home page...")
	}

};

func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//this function returns to Home-default
	t, err := template.ParseFiles("home.html")
	riperr(err)

	t.Execute(w, nil)
};

func main() {

	Server := httprouter.New()

	Server.GET("/a/*filename", ServeHTML)                              //routing to any page of /a/ prefix
	Server.GET("/", ServeHTML)                                         //default, routing to home page
	Server.ServeFiles("/resources/*filepath", http.Dir("./resources")) //css,js
	fmt.Println("---default waiting at :3000")

	http.ListenAndServe(GetPort(), Server)
};

func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the env
	if port == "" {
		port = "3000"
		fmt.Println("INFO: No PORT env var detected, defaulting to ", port)
	}
	return ":" + port
};

func riperr(err error) {
	if err != nil {
		log.Fatal(err)
	}
};
