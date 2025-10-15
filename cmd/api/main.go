package main

import (
	_ "fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/Arroziqi/car-rental-technical-test-pharos.git/common/database"
	"github.com/gin-gonic/gin"

	// customer
	customerRepo "github.com/Arroziqi/car-rental-technical-test-pharos.git/features/customer/adapter/persistence/sql"
	customerUC "github.com/Arroziqi/car-rental-technical-test-pharos.git/features/customer/application/usecase"
	customerCtrl "github.com/Arroziqi/car-rental-technical-test-pharos.git/features/customer/presentation/controller"

	// car
	carRepo "github.com/Arroziqi/car-rental-technical-test-pharos.git/features/car/adapter/persistence/sql"
	carUC "github.com/Arroziqi/car-rental-technical-test-pharos.git/features/car/application/usecase"
	carCtrl "github.com/Arroziqi/car-rental-technical-test-pharos.git/features/car/presentation/controller"

	// booking
	bookingRepo "github.com/Arroziqi/car-rental-technical-test-pharos.git/features/booking/adapter/persistence/sql"
	bookingUC "github.com/Arroziqi/car-rental-technical-test-pharos.git/features/booking/application/usecase"
	bookingCtrl "github.com/Arroziqi/car-rental-technical-test-pharos.git/features/booking/presentation/controller"
)

func main() {
	// load .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL env is required")
	}

	db, err := database.NewPostgres(dsn)
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	// run migrations (AutoMigrate)
	if err := database.AutoMigrateAll(db); err != nil {
		log.Fatalf("migrate error: %v", err)
	}

	// Repos
	cr := customerRepo.NewCustomerSQLRepository(db)
	carR := carRepo.NewCarRepository(db)
	br := bookingRepo.NewBookingRepository(db)

	// Usecases
	customerUsecase := customerUC.NewCustomerUsecase(cr)
	carUsecase := carUC.NewCarUsecase(carR)
	bookingUsecase := bookingUC.NewBookingUsecase(br, carR, cr) // booking needs car & customer

	// Controllers
	customerController := customerCtrl.NewCustomerController(customerUsecase)
	carController := carCtrl.NewCarController(carUsecase)
	bookingController := bookingCtrl.NewBookingController(bookingUsecase)

	// Router
	r := gin.Default()
	api := r.Group("/api/v1")

	// Customer routes
	cust := api.Group("/customers")
	{
		cust.POST("", customerController.Create)
		cust.GET("", customerController.List)
		cust.GET("/:id", customerController.GetByID)
		cust.PUT("/:id", customerController.Update)
		cust.DELETE("/:id", customerController.Delete)
	}

	// Car routes
	car := api.Group("/cars")
	{
		car.POST("", carController.Create)
		car.GET("", carController.List)
		car.GET("/:id", carController.GetByID)
		car.PUT("/:id", carController.Update)
		car.DELETE("/:id", carController.Delete)
	}

	// Booking routes
	book := api.Group("/bookings")
	{
		book.POST("", bookingController.Create)
		book.GET("", bookingController.List)
		book.GET("/:id", bookingController.GetByID)
		book.PUT("/:id", bookingController.Update)
		book.DELETE("/:id", bookingController.Delete)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("listening on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
