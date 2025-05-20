package types

type WorkExperience struct {
	ID          int
	CompanyName string
	Position    string
	Description string
	StartDate   string
	EndDate     string
}

type ProjectExperience struct {
	ID           int
	Title        string
	Description  string
	Technologies string
	GithubLink   string
	LiveLink     string
}
