package schema

type TodoRequest struct {
	Title     string `json:"title" validate:"required,min=3"`
	Completed bool   `json:"completed"`
}

/* type FindTodo struct {
	ID int `json:"id"`
} */

type TodoResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
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
