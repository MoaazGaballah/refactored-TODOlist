package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "fmt"
	_ "log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {

	connectionString :=
		//	fmt.Sprintf("user=%s password=%s dbname=%s ?sslmode=disable", user, password, dbname)
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	//connStr := fmt.Sprintf("host=localhost port=5432 user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	//psqlInfo := fmt.Sprintf("host=localhost port=5432 user=%s "+
	//	"password=%s dbname=%s sslmode=disable", user, password, dbname)

	//jdbc:postgresql://localhost:5432/tododb

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	//a.DB, err = sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", user, password, dbname))
	//a.DB, err = sql.Open("postgres", "user=todouser password=1 dbname=todos sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}

func (a *App) Run(addr string) {}
