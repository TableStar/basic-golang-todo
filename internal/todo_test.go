package todo

import (
	"reflect"
	"testing"
)

func TestAddTodo(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var testList TodoList

		got := AddTodo("bangun", &testList) //sending memory address to function
		want := TodoList{Todo{Id: 1, Task: "bangun", IsDone: false}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("slice has one todo", func(t *testing.T) {
		testList := TodoList{{Id: 2, Task: "belajar struct", IsDone: true}}

		got := AddTodo("belajar phyton", &testList)
		want := TodoList{
			{2, "belajar struct", true},
			{3, "belajar phyton", false},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}

func TestFindTodo(t *testing.T) {
	testList := TodoList{
		{2, "belajar struct", true},
		{3, "belajar phyton", false},
	}
	t.Run("find index of todo by Id", func(t *testing.T) {
		todoId := 3

		got := FindTodoIndexById(todoId, testList)
		want := 1
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("find index of todo by task", func(t *testing.T) {
		todoTask := "belajar struct"

		got := FindTodoIndexByTask(todoTask, testList)
		want := 0
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
