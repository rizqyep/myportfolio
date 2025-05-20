package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rizqyep/myportfolio/types"
)

type Handler struct {
	db   *sql.DB
	tmpl *template.Template
}

func NewHandler(db *sql.DB, tmpl *template.Template) *Handler {
	return &Handler{
		db:   db,
		tmpl: tmpl,
	}
}

func (h *Handler) ServeIndex(w http.ResponseWriter, r *http.Request) {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
	}

	userAgent := r.Header.Get("User-Agent")
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	accessInfo := map[string]string{
		"IP":        ip,
		"UserAgent": userAgent,
		"Timestamp": timestamp,
	}
	h.tmpl.ExecuteTemplate(w, "index.html", accessInfo)
}

func (h *Handler) ServeWork(w http.ResponseWriter, r *http.Request) {
	// ... existing code ...

	var experiences []types.WorkExperience
	rows, err := h.db.Query("SELECT id, company_name, position, description, start_date, end_date FROM work_experiences ORDER BY id")
	if err != nil {
		log.Printf("Database query error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var exp types.WorkExperience
		err := rows.Scan(&exp.ID, &exp.CompanyName, &exp.Position, &exp.Description, &exp.StartDate, &exp.EndDate)
		if err != nil {
			log.Printf("Row scan error: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		experiences = append(experiences, exp)
	}

	log.Printf("Found %d work experiences", len(experiences))
	for _, exp := range experiences {
		log.Printf("Experience: %s at %s", exp.Position, exp.CompanyName)
	}

	err = h.tmpl.ExecuteTemplate(w, "work.html", experiences)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func (h *Handler) ServeProjects(w http.ResponseWriter, r *http.Request) {
	var projects []types.ProjectExperience
	rows, err := h.db.Query("SELECT id, title, description, technologies, github_link, live_link FROM project_experiences")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var proj types.ProjectExperience
		err := rows.Scan(&proj.ID, &proj.Title, &proj.Description, &proj.Technologies, &proj.GithubLink, &proj.LiveLink)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		projects = append(projects, proj)
	}

	for _, project := range projects {
		fmt.Println("project description", project.Description)
	}

	err = h.tmpl.ExecuteTemplate(w, "projects.html", projects)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
}
