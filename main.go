// main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Article - Our struct for all articles
type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnAllPlanetas(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllPlanetas")
	json.NewEncoder(w).Encode(Planeta)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	// update our global Articles array to include
	// our new Article
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}

}

func handleRequests() {
	fmt.Println("[B2W] - API with Planets")
	myRouter := mux.NewRouter().StrictSlash(true)
	fmt.Println("[B2W] - /")
	myRouter.HandleFunc("/", homePage)
	fmt.Println("[B2W] - /articles")
	myRouter.HandleFunc("/articles", returnAllArticles)
	fmt.Println("[B2W] - /article - POST")
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	fmt.Println("[B2W] - /article/{id} - DELETE")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	fmt.Println("[B2W] - /article/{id}")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)

	fmt.Println("[B2W] - /planetas")
	myRouter.HandleFunc("/planetas", returnAllPlanetas)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	Planetas = []Planeta{
		Planeta{Id: "1", Nome: "Saturno", Clima: "Winter", Terreno: "Congelado"},
	}

	handleRequests()
}

// func main() {
// 	fmt.Println("Iniciando a API - B2W Planets")
// 	http.HandleFunc("/planet", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Planet, %q", html.EscapeString(r.URL.Path))
// 	})

// 	http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "start")
// 	})

// 	log.Fatal(http.ListenAndServe(":8081", nil))
// }
