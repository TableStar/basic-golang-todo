package todo

import (
	"reflect"
	"testing"
)

func TestAddTodo(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var testList TodoList

		got := AddTodo("bangun", testList)
		want := TodoList{Todo{Id: 1, Task: "bangun", IsDone: false}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("slice has one todo", func(t *testing.T) {
		testList := TodoList{{Id: 1, Task: "belajar struct", IsDone: true}}

		got := AddTodo("belajar phyton", testList)
		want := TodoList{
			{1, "belajar struct", true},
			{2, "belajar phyton", false},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
