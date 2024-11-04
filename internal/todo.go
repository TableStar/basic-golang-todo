package todo

type Todo struct {
	Id     int
	Task   string
	IsDone bool
}

type TodoList []Todo

func AddTodo(task string, list *TodoList) TodoList {
	//because list is a memory address we need *list to do anything with its value
	var lastId int
	if len(*list) > 0 {
		lastId = (*list)[len(*list)-1].Id + 1
	} else {
		lastId = 1
	}

	newTodo := Todo{Id: lastId, Task: task, IsDone: false}

	*list = append(*list, newTodo)
	return *list
}

// func UpdateTodoTask(task string, list *TodoList, id int) TodoList {
// 	*list
// }

func FindTodoIndexById(id int, list TodoList) int {
	for i, v := range list {
		if id == v.Id {
			return i
		}
	}
	return -1
}

func FindTodoIndexByTask(task string, list TodoList) int {
	for i, v := range list {
		if task == v.Task {
			return i
		}
	}
	return -1
}
