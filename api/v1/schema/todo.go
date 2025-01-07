package schema

type TodoRequest struct {
	Title     string `json:"title" validate:"required,min=3"`
	Completed bool   `json:"completed" validate:"required"`
}

/* type FindTodo struct {
	ID int `json:"id"`
} */

type TodoResponse struct {
	ID        int    `json:"id" validate:"required"`
	Title     string `json:"title" validate:"required,min=3"`
	Completed bool   `json:"completed" validate:"required"`
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
