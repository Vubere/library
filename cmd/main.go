package main

import (
	"database/sql"
	"log"
	"victorubere/library/controllers"
	"victorubere/library/repository"
	"victorubere/library/services"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// db username: admin
// db password :Du1jO9COEEmzgEKo8BzN

// host;database-1.copayogeootm.us-east-1.rds.amazonaws.com
// port:3306
type App struct {
	db     *sql.DB
	gormDB *gorm.DB
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := App{}
	err = app.OpenDB()

	if err != nil {
		panic(err)
	}
	log.Println("Database connection established!")
	defer app.db.Close()

	userRepository := repository.NewUserRpository(app.gormDB)
	userService := services.NewUserService(userRepository)

	bookRepository := repository.NewBookRepository(app.gormDB)
	bookService := services.NewBookService(bookRepository)

	visitationRepository := repository.NewVisitationRepository(app.gormDB)
	visitationService := services.NewVisitationService(visitationRepository)

	borrowedRepository := repository.NewBorrowedRepository(app.gormDB)
	borrowedService := services.NewBorrowedService(borrowedRepository)

	bookReadRepository := repository.NewBookReadsRepository(app.gormDB)
	bookReadService := services.NewBookReadsService(bookReadRepository)

	controller := controllers.NewController(userService, bookService, visitationService, borrowedService, bookReadService)

	router, err := controller.InitializeRoutes()
	if err != nil {
		panic(err)
	}
	err = router.Run(":5000")
	if err != nil {
		panic(err)
	}
}
