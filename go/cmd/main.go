package main

import (
	"freelancer/college-app/go/api"
	"freelancer/college-app/go/config"
	"freelancer/college-app/go/infrastructure/repository/mongo"
	"freelancer/college-app/go/pkg/token"
	"freelancer/college-app/go/usecase/advice"
	"freelancer/college-app/go/usecase/contact"
	"freelancer/college-app/go/usecase/department"
	"freelancer/college-app/go/usecase/installed_app"
	"freelancer/college-app/go/usecase/suggestion"
	"freelancer/college-app/go/usecase/university"
	"freelancer/college-app/go/usecase/user"
	"freelancer/college-app/go/usecase/user_mood"
	"log"
)

func main() {

	// Handles MongoDB connection.
	log.Println("[main]: Starting a new connection to MongoDB...")
	db, err := mongo.NewStorage(config.DB_URL, config.DB_NAME, config.DB_USER, config.DB_PASSWORD)
	if err != nil {
		log.Panic("[main] NewStorage error:", err)
		return
	}
	log.Println("[main]: MongoDB connection established")

	// Token service.
	tokenService, err := token.NewService(config.SECRET_KEY)
	if err != nil {
		log.Panic("[main] NewJWTMaker error: %w", err)
	}

	// Repositories.
	userRepository := mongo.NewUserRepository(db)
	universityRepository := mongo.NewUniversityRepository(db)
	suggestionRepository := mongo.NewSuggestionRepository(db)
	contactRepository := mongo.NewContactRepository(db)
	departmentRepository := mongo.NewDepartmentRepository(db)
	userMoodRepository := mongo.NewUserMoodRepository(db)
	adviceRepository := mongo.NewAdviceRepository(db)
	installedAppRepository := mongo.NewInstalledAppRepository(db)

	// Services.
	userService := user.NewService(userRepository, universityRepository)
	universityService := university.NewService(userRepository, universityRepository)
	suggestionService := suggestion.NewService(suggestionRepository)
	contactService := contact.NewService(userRepository, contactRepository)
	departmentService := department.NewService(departmentRepository, userRepository)
	userMoodService := user_mood.NewService(userRepository, userMoodRepository)
	adviceService := advice.NewService(adviceRepository, userRepository, universityRepository)
	installedAppService := installed_app.NewService(installedAppRepository, userRepository)

	services := api.Services{
		UserService:         *userService,
		UniversityService:   *universityService,
		SuggestionService:   *suggestionService,
		ContactService:      *contactService,
		UserMoodService:     *userMoodService,
		AdviceService:       *adviceService,
		DepartmentService:   *departmentService,
		InstalledAppService: *installedAppService,
		TokenService:        tokenService,
	}

	// Start http server.
	api.ListenAndServe(services)

}
