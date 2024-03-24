package main

import (
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"unklogger.com/src/application/services"
	infr "unklogger.com/src/infrastrucutre"
	rest "unklogger.com/src/interface/api"
)

func main() {
	var cfg infr.Config
	infr.ReadFile(&cfg)
	infr.ReadEnv(&cfg)
	fmt.Printf("%+v", cfg)
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Server.Host, cfg.Database.Username, cfg.Database.Password, cfg.Database.Name, cfg.Server.Port)
	gormDB, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	os.Setenv("JWT_SECRET", cfg.Server.Secret)

	activityRepo := infr.NewGormActivityRepository(gormDB)
	authRepo := infr.NewGormAuthRepository(gormDB)

	gormDB.AutoMigrate(&infr.User{})
	gormDB.AutoMigrate(&infr.UserAuth{})
	gormDB.AutoMigrate(&infr.ActivityLog{})
	activityService := services.NewActivityService(activityRepo)
	authService := services.NewAuthService(authRepo)

	e := echo.New()
	rest.NewActivityController(e, activityService, infr.AuthMiddleware)
	rest.NewAuthController(e, authService)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
