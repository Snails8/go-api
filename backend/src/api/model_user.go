package api

type User struct {
	Id   int32 `json:"id"`
	Name string `json"name"`
}


type GetUsersResponse struct {
	Users []User `json:"users"`
}
