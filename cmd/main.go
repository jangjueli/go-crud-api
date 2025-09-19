package main

import (
	"go-crud-api/internal"
	"go-crud-api/internal/auth"
	"go-crud-api/internal/middleware"
	"go-crud-api/internal/user"
	"go-crud-api/pkg/config"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.Load()
	db := internal.ConnectDB(cfg.DBUrl)
	defer db.Close()

	// module
	userModule := user.InitModule(db)
	authModule := auth.InitModule(db, *userModule.Service)

	r := gin.Default()
	r.Use(middleware.Logger())

	userModule.RegisterRoutes(r)
	authModule.RegisterRoutes(r)

	r.Run(":3000")
}
