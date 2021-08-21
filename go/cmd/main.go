package main

import (
	"fmt"
	"freelancer/college-app/go/api"
	"freelancer/college-app/go/config"
	"freelancer/college-app/go/infrastructure/repository/mongo"
	"freelancer/college-app/go/token"
	"freelancer/college-app/go/usecase/suggestion"
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

	// Token service.
	tokenService, err := token.NewJWTMaker(config.SECRET_KEY)
	if err != nil {
		log.Panic(fmt.Errorf("cannot create token: %w", err))
	}

	// Repositories.
	userRepository := mongo.NewUserRepository(db)
	universityRepository := mongo.NewUniversityRepository(db)
	suggestionRepository := mongo.NewSuggestionRepository(db)

	// Services.
	userService := user.NewService(userRepository, universityRepository)
	universityService := university.NewService(universityRepository)
	suggestionService := suggestion.NewService(suggestionRepository)

	services := api.Services{
		UserService:       *userService,
		UniversityService: *universityService,
		SuggestionService: *suggestionService,
		TokenService:      tokenService,
	}

	// Start http server.
	api.ListenAndServe(services)

}
