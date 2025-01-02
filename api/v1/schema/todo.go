package schema

type TodoRequest struct {
	Title     string `json:"title" validate:"required,min=3"`
	Completed bool   `json:"completed"`
}

type TodoResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type DeleteTodoResponse struct {
	ID int `json:"id"`
}
