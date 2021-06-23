package app

import (
	"log"
	"mysql_rest_api_users/domain"
	"mysql_rest_api_users/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func Start() {
	DB := ConnectDB()
	router := mux.NewRouter()
	ch := UserHandlers{service.NewUserService(domain.NewUserRepositoryDb(DB))}
	route(router, &ch)
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
