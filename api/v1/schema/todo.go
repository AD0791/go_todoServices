package schema

type TodoRequest struct {
	Title     string `json:"title" validate:"required,min=3"` //  example:"Sample Todo"
	Completed bool   `json:"completed"`                       //example:"false"
}

/* type FindTodo struct {
	ID int `json:"id"`
} */

type TodoResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title" validate:"required,min=3"`
	Completed bool   `json:"completed"`
}

type MessageResponse struct {
	ID      int
	Message string
}

/* type UpdateTodoResponse struct {
	ID      int
	Message string
}
*/

type TodoSQLResponse struct {
	ID        uint   `json:"id" example:"1"`
	Title     string `json:"title" example:"Sample Todo"`
	Completed bool   `json:"completed" example:"false"`
	CreatedAt string `json:"created_at"` // example:"2023-01-01T00:00:00Z"
	UpdatedAt string `json:"updated_at"` // example:"2023-01-01T00:00:00Z"
}

type MessageSQLResponse struct {
	MessageResponse
	DeletedAt string `json:"delete_at"`
}
