package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/akhil/go-bookstore/pkg/routes"
	"github.com/gorilla/handlers"

)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)


	//this is the CORS middleware
	 // CORS middleware
	 corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{"*"}),
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
        handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
    )



	http.Handle("/",corsHandler(r))
	// log.Fatal(http.ListenAndServe("0.0.0.0:9010", r))
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}