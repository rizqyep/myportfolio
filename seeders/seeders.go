package seeders

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Seeder struct {
	db *sql.DB
}

func NewSeeder(db *sql.DB) *Seeder {
	return &Seeder{db: db}
}

func (s *Seeder) SeedWorkExperiences() error {
	// Check if data already exists
	var count int
	err := s.db.QueryRow("SELECT COUNT(*) FROM work_experiences").Scan(&count)
	if err != nil {
		return err
	}

	// If data exists, skip seeding
	if count > 0 {
		return nil
	}

	// Seed work experiences
	_, err = s.db.Exec(`INSERT INTO work_experiences (company_name, position, description, start_date, end_date) VALUES
    ('Nodewave', 'Chief Technology Officer', 'Co-founded a digital transformation company that focuses not only on building software solutions, but products that helps company navigate through the emerging digital era. My role is to gradually and efficiently scaled the Product and Engineering teams that now already served 10+ Clients and 20+ Projects, most notably led a Goverment Centralized Data Repository Project, AI Based Record Management System and also Computer Vision Powered Construction Surveillance System', 'July 2024', 'Now'),
    ('S Quantum Engine', 'Backend Engineer', 'One of the starting teams that build SQE KYC and E-Sign System, involved in initial System Design and Planning. Notable contributions are, optimizing the billing data ingestion flow and also Optimizing the E-Sign Publish and Subscribe flow, resulting in SQEKYC now able to Serve 50k+ E-Sign Request during Business Hour', 'April 2023', 'Now'),
    ('Staff Any', 'Software Engineer', 'Joined StaffAny as a Backend-heavy Full Stack Engineer. Noticed some slow page performance due to the nonexistent of proper paginations on some heavy APIs, which later on leading to be initiating and involved in implementing Key-Set Pagination for various StaffAny main API, resulting in 200% increase in page load on key areas like Timesheet, Employee Lists and also CICO Page', 'August 2022', 'April 2023'),
    ('Pandatech', 'Backend Engineer', 'Building various backend systems for various clients, using various tech stacks including Node.js, Golang, Python and Java. Highlighted contribution is involved in the System Design and Planning, also created the initial codebase structure for a shipping permit and management system for one of the emerging shipping company in Indonesia', 'March 2022', 'August 2022'),
	('LingoTalk', 'Lead Web Platform Engineer', 'Leading the whole Web Platform Team. Learning on the spot to improve engineering team performance, experimenting with various framework and methods, resulting in an increase of productivity and task delivery efficency. Key figure on integrating Zoom WebAssembly SDK to LingoTalk appointment platform, Won the Employee of the Month in November 2021, 3rd Promotion in 1 Year', 'July 2021', 'March 2022'),
	('LingoTalk', 'Lead Backend Engineer', 'Leading Backend Engineering Team, scaling the team of interns and fulltimers. Initially tasked to build a newly migrated appointment platform for LingoTalk and also handled the migration of the old SQL to FaunaDB platform.', 'March 2021', 'July 2021'),
	('LingoTalk', 'Backend Engineer', 'Joined LingoTalk as a Backend Engineer, handling backend tasks and also participated in planning the migration of the old SQL to FaunaDB platform.', 'January 2021', 'March 2021');`)
	if err != nil {
		return err
	}

	fmt.Println("Seeded work experiences")

	return nil
}

func (s *Seeder) SeedProjectExperiences() error {
	// Check if data already exists
	var count int
	err := s.db.QueryRow("SELECT COUNT(*) FROM project_experiences").Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	_, err = s.db.Exec(`	INSERT INTO project_experiences (title, description, technologies, github_link, live_link) VALUES
    ('E-commerce Platform', 'A full-stack e-commerce solution with shopping cart and payment integration.', 'React, Node.js, PostgreSQL, Stripe', 'https://github.com/username/ecommerce', 'https://demo-ecommerce.com'),
    ('Task Management App', 'Kanban-style project management tool with real-time updates.', 'Vue.js, Express, MongoDB, Socket.io', 'https://github.com/username/task-manager', 'https://task-app-demo.com'),
    ('Weather Dashboard', 'Weather forecast application with location-based services.', 'JavaScript, OpenWeather API, HTML5, CSS3', 'https://github.com/username/weather-app', 'https://weather-dashboard-demo.com');`)

	if err != nil {
		return err
	}

	fmt.Println("Seeded project experiences")

	return nil
}

func (s *Seeder) SeedAll() error {
	if err := s.SeedWorkExperiences(); err != nil {
		return err
	}

	fmt.Println("Seeded all data")

	return nil
}
