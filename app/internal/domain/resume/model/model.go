package job

type Resume struct {
	Id         string `json:"id"`
	WorkerId   string `json:"worker_id"`
	About      string `json:"about"`
	Experience string `json:"experience"`
}
