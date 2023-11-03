package response

type LinesReponce struct {
	Id        int    `json:"id"`
	CreatedAt string `json:"createdAt"`
	UserId    string `json:"userId"`
	Content   string `json:"content"`
}
