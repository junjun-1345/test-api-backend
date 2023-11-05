package request

type CreateUserRequest struct {
	UserId         string `json:"userId"`
	Name           string `json:"name"`
	WorkInWeekDay  int    `json:"workInWeekDay"`
	WorkInWeekTime int    `json:"workInWeektime"`
}
