package main

import (
	"admin-ais/model"
	"admin-ais/route"
	"log"
)

func main() {

	db, err := model.DBConnection()
	if err != nil {
		log.Fatalf("Error connecting to database",err)

	}

	route.SetupRoutes(db)

	//go http.ListenAndServe(":"+os.Getenv("PORT"), nil)

}
