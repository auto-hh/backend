package model

type Profile struct {
	Experience string `json:"experience"`
	JobTitle   string `json:"job_title"`
	Grade      string `json:"grade"`
	WorkFormat string `json:"work_format"`
	Salary     int    `json:"salary"`
	City       string `json:"city"`
	AboutMe    string `json:"about_me"`
	RecentJobs string `json:"recent_jobs"`
}
