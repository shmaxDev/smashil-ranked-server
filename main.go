package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"smashil-ranked/config"
	"smashil-ranked/handlers"
	internalHttp "smashil-ranked/http"
	"smashil-ranked/queueLoop"
	"smashil-ranked/repositories"
	"smashil-ranked/services"

	_ "github.com/lib/pq"
)

func main() {
	vars := config.GetEnv()
	db, err := sql.Open(vars.DbType, vars.DbConnectionString)

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	queueLoop.StartLoop()

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	matchesRepo := repositories.NewMatchRepository(db)
	matchesService := services.NewMatchService(matchesRepo)
	matchesHandler := handlers.NewMatchHandler(matchesService)

	mux := http.NewServeMux()

	
	internalHttp.SetupRouter(mux, userHandler, matchesHandler)

	log.Printf("Server started on port %s", vars.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", vars.Port), mux))

}
