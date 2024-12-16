// package migration

// import (
// 	"fmt"
// 	"log"

// 	"github.com/golang-migrate/migrate/v4"
// 	_ "github.com/golang-migrate/migrate/v4/source/file"
// )

// // Migrate function applies migrations to the database
// func Migrate() {
// 	// Create a new migration instance
// 	m, err := migrate.New(
// 		"file://pkg/database/migrations", // Path to the migrations directory

// 	)

// 	if err != nil {
// 		log.Fatalf("Failed to initialize migrations: %v", err)
// 	}

// 	// Apply all up migrations
// 	if err := m.Up(); err != nil {
// 		if err == migrate.ErrNoChange {
// 			fmt.Println("No new migrations to apply.")
// 		} else {
// 			log.Fatalf("Migration failed: %v", err)
// 		}
// 	} else {
// 		fmt.Println("Migrations applied successfully.")
// 	}
// }
