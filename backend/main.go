package main

import (
	"log"

	"github.com/codeschris/super-list/backend/config"
	"github.com/codeschris/super-list/backend/database"
	"github.com/codeschris/super-list/backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Connect to database
	db := database.Connect(cfg.DatabaseURL)
	defer db.Close()

	// Run migrations
	database.Migrate(db)

	// Setup router
	r := gin.Default()
	routes.Setup(r, db, cfg)

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	r.Run(":" + cfg.Port)
}
