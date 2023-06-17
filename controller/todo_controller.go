package controller

import (
	"finalproject1/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var Todos = []model.Todo{
	{
		Id:          1,
		Name:        "Assignment 1",
		Deadline:    time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		Description: "Something To Do",
		Status:      "New",
	},
	{
		Id:          2,
		Name:        "Assignment 2",
		Deadline:    time.Date(2021, 8, 25, 14, 30, 45, 100, time.Local),
		Description: "Something To Do now",
		Status:      "New",
	},
}

var maxId = 2

// @Summary get all items in the todo list
// @ID get-all-todos
// @Produce json
// @Success 200 {object} model.Todo
// @Router /todos [get]
func GetTodosHandler(ctx *gin.Context) {
	if Todos == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"Data": "There is nothing to do",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Data": Todos,
		})
	}
}

// @Summary get a todo item by ID
// @ID get-todo-by-id
// @Produce json
// @Param id path string true "todo ID"
// @Success 200 {object} model.Todo
// @Failure 404 {object} model.Message
// @Router /todos/{id} [get]
func GetTodoHandler(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	for _, todo := range Todos {
		if todo.Id == id {
			ctx.JSON(http.StatusOK, gin.H{
				"Data": todo,
			})
			return
		}
	}

	r := model.Message{Message: "Todo not found"}
	ctx.JSON(http.StatusNotFound, r)

}

// @Summary add a new item to the todo list
// @ID create-todo
// @Produce json
// @Param data body model.TodoRequest true "Todo Request"
// @Success 200 {object} model.Todo
// @Failure 400 {object} model.Message
// @Router /todos [post]
func CreateTodoHandler(ctx *gin.Context) {
	var createData model.TodoRequest
	var todo model.Todo
	err := ctx.ShouldBindJSON(&createData)
	if err != nil {
		r := model.Message{Message: "an error occurred while creating todo"}
		ctx.JSON(http.StatusBadRequest, r)
		return
	}

	maxId++

	todo = convertRequest(maxId, createData)
	Todos = append(Todos, todo)

	ctx.JSON(http.StatusOK, gin.H{
		"Data": todo,
	})

}

// @Summary update a todo item by ID
// @ID update-todo-by-id
// @Produce json
// @Param id path string true "todo ID"
// @Param data body model.TodoRequest true "Todo Request"
// @Success 200 {object} model.Todo
// @Failure 404 {object} model.Message
// @Router /todos/{id} [put]
func UpdateTodoHandler(ctx *gin.Context) {

	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	var updateData model.TodoRequest
	err := ctx.ShouldBindJSON(&updateData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	for index, todo := range Todos {
		if todo.Id == id {

			todo.Name = updateData.Name
			todo.Deadline = updateData.Deadline
			todo.Description = updateData.Description
			todo.Status = updateData.Status
			Todos[index] = todo

			ctx.JSON(http.StatusOK, gin.H{
				"Updated Data": todo,
			})
			return
		}
	}

	r := model.Message{Message: "Todo not found"}
	ctx.JSON(http.StatusNotFound, r)

}

// @Summary delete a todo item by ID
// @ID delete-todo-by-id
// @Produce json
// @Param id path string true "todo ID"
// @Success 200 {object} model.Todo
// @Failure 404 {object} model.Message
// @Router /todos/{id} [delete]
func DeleteTodoHandler(ctx *gin.Context) {

	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	for index, todo := range Todos {
		if todo.Id == id {

			Todos = append(Todos[:index], Todos[index+1:]...)

			ctx.JSON(http.StatusOK, gin.H{
				"Deleted Data": todo,
			})
			return
		}
	}

	r := model.Message{Message: "Todo not found"}
	ctx.JSON(http.StatusNotFound, r)
}

func convertRequest(id int, todoRequest model.TodoRequest) model.Todo {
	var todo model.Todo
	todo.Id = maxId
	todo.Name = todoRequest.Name
	todo.Deadline = todoRequest.Deadline
	todo.Description = todoRequest.Description
	todo.Status = todoRequest.Status

	return todo
}
