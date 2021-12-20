package services

import (
	"ex1/todo-api/common"
	"ex1/todo-api/dtos"
	"ex1/todo-api/entities"
	"ex1/todo-api/repositories/mocks"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestTodoService_FindAll(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockTodoRepo := mocks.NewMockTodoRepository(mockCtrl)
	mockTodoRepo.EXPECT().FindAll().Return([]entities.Todo{
		{
			Title:   "abc",
			Content: "abc",
		},
	}, common.DatabaseError{Type: "abc", Code: 1})

	todoService := NewTodoService(mockTodoRepo)
	res, _ := todoService.FindAll()
	if len(res) != 1 {
		t.Errorf("Todo not matching")
	}
}

func TestTodoService_FindByID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockTodoRepo := mocks.NewMockTodoRepository(mockCtrl)
	mockTodoRepo.EXPECT().FindByID(gomock.Any()).Return(entities.Todo{Title: "abc", Content: "abc"}, common.DatabaseError{Type: "abc", Code: 1})

	todoService := NewTodoService(mockTodoRepo)
	res, _ := todoService.FindByID(1)
	if res.Title != "abc" {
		t.Errorf("Todo not matching")
	}
}

func TestTodoService_Save(t *testing.T) {
	var todo dtos.TodoDTO
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockTodoRepo := mocks.NewMockTodoRepository(mockCtrl)
	mockTodoRepo.EXPECT().Save(gomock.Any()).Return(entities.Todo{Title: "abc", Content: "abc"}, common.DatabaseError{Type: "abc", Code: 1})

	todoService := NewTodoService(mockTodoRepo)
	res, _ := todoService.Save(todo)
	if res.Title != "abc" {
		t.Errorf("Todo not matching")
	}
}

func TestTodoService_FindByUserId(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockTodoRepo := mocks.NewMockTodoRepository(mockCtrl)
	mockTodoRepo.EXPECT().FindByUserId(gomock.Any()).Return([]entities.Todo{{Title: "abc", Content: "abc"}}, common.DatabaseError{Type: "abc", Code: 1})

	todoService := NewTodoService(mockTodoRepo)
	res, _ := todoService.FindByUserId(1)
	if len(res) != 1 {
		t.Errorf("Todo not matching")
	}
}
