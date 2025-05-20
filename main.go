package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rizqyep/myportfolio/handlers"
	"github.com/rizqyep/myportfolio/seeders"
)

var db *sql.DB
var tmpl *template.Template

func init() {
	// Initialize SQLite database
	var err error
	db, err = sql.Open("sqlite3", "./portfolio.db")
	if err != nil {
		log.Fatal(err)
	}

	// Run migrations
	if err := runMigrations(); err != nil {
		log.Fatal(err)
	}

	seed := seeders.NewSeeder(db)
	if err := seed.SeedAll(); err != nil {
		log.Fatal(err)
	}

	// Initialize templates
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func runMigrations() error {
	// Read migration file
	migration, err := os.ReadFile("migrations/schema.sql")
	if err != nil {
		return err
	}

	// Execute migration
	_, err = db.Exec(string(migration))
	if err != nil {
		return err
	}

	log.Println("Migrations completed successfully")
	return nil
}

func main() {
	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	handler := handlers.NewHandler(db, tmpl)
	// Route handlers
	http.HandleFunc("/", handler.ServeIndex)
	http.HandleFunc("/works", handler.ServeWork)
	http.HandleFunc("/projects", handler.ServeProjects)

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
