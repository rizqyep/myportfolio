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
	tmpl = template.Must(template.New("").Funcs(template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"mul": func(a, b int) int {
			return a * b
		},
		"sub": func(a, b int) int {
			return a - b
		},
		"div": func(a, b int) int {
			if b == 0 {
				return 0
			}
			return a / b
		},
		"mod": func(a, b int) int {
			if b == 0 {
				return 0
			}
			return a % b
		},
	}).ParseGlob("templates/*.html"))
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
	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("User-agent: *\nAllow: /"))
	})

	log.Println("Server starting on :8110...")
	log.Fatal(http.ListenAndServe(":8110", nil))
}
