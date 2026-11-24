package query

type Task struct {
	ID string `json:"id"`
	Title string `json:"title"`
	IsDone bool `json:"is_done"`
}
