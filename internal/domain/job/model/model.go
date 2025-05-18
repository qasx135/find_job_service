package job

type Job struct {
	Id           string `json:"id"`
	Header       string `json:"header"`
	Experience   string `json:"experience"`
	Employment   string `json:"employment"`
	Schedule     string `json:"schedule"`
	WorkFormat   string `json:"work_format"`
	WorkingHours string `json:"working_hours"`
	Description  string `json:"description"`
	EmployerId   string `json:"employer_id"`
}
