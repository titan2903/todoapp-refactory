package todo

type TodoFormatter struct {
	ID               	int    	`json:"id"`
	UserID           	int    	`json:"user_id"`
	Task 				string 	`json:"task"`
	Title 				string 	`json:"title"`
	IsCompleted         int 	`json:"is_completed"`
}

func FormatTodo(todo Todo) TodoFormatter {
	formatter := TodoFormatter{
		ID:               	todo.ID,
		UserID:           	todo.UserID,
		Task:             	todo.Task,
		IsCompleted: 		todo.IsCompleted,
		Title: 				todo.Title,
	}

	return formatter
}


func FormatTodos(todos []Todo) []TodoFormatter {
	todosFormatter := []TodoFormatter{}

	for _, todo := range todos {
		todoFormatter := FormatTodo(todo)
		todosFormatter = append(todosFormatter, todoFormatter)
	}

	return todosFormatter
}

type TodoDetailFormatter struct {
	ID               	int    	`json:"id"`
	UserID           	int    	`json:"user_id"`
	Task 				string 	`json:"task"`
	Title 				string 	`json:"title"`
	IsCompleted         int 	`json:"is_completed"`
	User 				TodoUserFormatter `json:"user"`
}

type TodoUserFormatter struct {
	Name 		string `json:"name"`
}

func FormatTodoDetail(todo Todo) TodoDetailFormatter {
	todoDetailFormatter := TodoDetailFormatter{}
	todoDetailFormatter.ID = todo.ID
	todoDetailFormatter.Task = todo.Task
	todoDetailFormatter.Title = todo.Title
	todoDetailFormatter.IsCompleted = todo.IsCompleted
	todoDetailFormatter.UserID = todo.UserID

	user := todo.User
	todoUserFormatter := TodoUserFormatter{}
	todoUserFormatter.Name = user.Name
	todoDetailFormatter.User = todoUserFormatter

	return todoDetailFormatter
}