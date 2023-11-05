package request

type CreateUserRequest struct {
	UserId     string `json:"userId"`
	Name       string `json:"name"`
	WorkInWeek int    `json:"workInWeek"`
}
