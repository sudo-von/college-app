package main

import (
	"freelancer/college-app/go/api"
	"freelancer/college-app/go/config"
	"freelancer/college-app/go/infrastructure/repository/mongo"
	"freelancer/college-app/go/usecase/university"
	"freelancer/college-app/go/usecase/user"
	"log"
)

func main() {

	// Handles MongoDB connection.
	log.Println("[main]: Starting a new connection to MongoDB...")
	db, err := mongo.NewStorage(config.DB_URL, config.DB_NAME, config.DB_USER, config.DB_PASSWORD)
	if err != nil {
		log.Println("[main] error:", err)
		return
	}
	log.Println("[main]: MongoDB connection established")

	// Repositories.
	userRepository := mongo.NewUserRepository(db)
	universityRepository := mongo.NewUniversityRepository(db)

	// Services.
	userService := user.NewService(userRepository, universityRepository)
	universityService := university.NewService(universityRepository)

	services := api.Services{
		UserService:       *userService,
		UniversityService: *universityService,
	}

	// Start http server.
	api.ListenAndServe(services)

}
