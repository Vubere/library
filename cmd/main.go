package main

import (
	"database/sql"
	"log"
	"victorubere/library/controllers"
	"victorubere/library/repository"
	"victorubere/library/services"

	"gorm.io/gorm"
)

type App struct {
	db     *sql.DB
	gormDB *gorm.DB
}

func main() {

	app := App{}
	err := app.OpenDB()

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

	bookReadRepository := repository.NewBookReadRepository(app.gormDB)
	bookReadService := services.NewBookReadService(bookReadRepository)

	controller := controllers.NewController(userService, bookService, visitationService, borrowedService, bookReadService)

	router, err := controller.InitializeRoutes()
	if err != nil {
		panic(err)
	}
	err = router.Run(":9000")
	if err != nil {
		panic(err)
	}
}
