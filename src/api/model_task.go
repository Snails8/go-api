package api

type Task struct {
	Id   int32 `json:"id"`
	Title string `json"name"`
	Content string `json:content`
}


type GetTasksResponse struct {
	Tasks []Task `json:"users"`
}
