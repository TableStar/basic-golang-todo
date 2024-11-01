package todo

type Todo struct {
	Id     int
	Task   string
	IsDone bool
}

type TodoList []Todo

func AddTodo(task string, list TodoList) TodoList {
	newTodo := Todo{Id: len(list) + 1, Task: task, IsDone: false}

	list = append(list, newTodo)
	return list
}
