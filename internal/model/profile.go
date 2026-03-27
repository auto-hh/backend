package model

type Profile struct {
	Experience string `json:"experience"`
	JobTitle   string `json:"jobTitle"`
	Grade      string `json:"grade"`
	WorkFormat string `json:"workFormat"`
	Salary     int `json:"salary"`
	City       string `json:"city"`
	AboutMe    string `json:"aboutMe"`
	RecentJobs string `json:"recentJobs"`
}
