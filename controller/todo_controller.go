package controller

import (
	"finalproject1/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary get all items in the todo list
// @ID get-all-todos
// @Produce json
// @Success 200 {object} model.Todo
// @Router /todos [get]
func (h *handler) GetTodosHandler(ctx *gin.Context) {
	var todos []model.Todo

	if result := h.DB.Find(&todos); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &todos)
}

// @Summary get a todo item by ID
// @ID get-todo-by-id
// @Produce json
// @Param id path string true "todo ID"
// @Success 200 {object} model.Todo
// @Failure 404 {object} model.Message
// @Router /todos/{id} [get]
func (h *handler) GetTodoHandler(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	var todo model.Todo

	if result := h.DB.First(&todo, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &todo)

}

// @Summary add a new item to the todo list
// @ID create-todo
// @Produce json
// @Param data body model.TodoRequest true "Todo Request"
// @Success 200 {object} model.Todo
// @Failure 400 {object} model.Message
// @Router /todos [post]
func (h *handler) CreateTodoHandler(ctx *gin.Context) {
	input := model.TodoRequest{}
	if err := ctx.BindJSON(&input); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var todos model.Todo

	todos.Name = input.Name
	todos.Deadline = input.Deadline
	todos.Description = input.Description
	todos.Status = input.Status

	if result := h.DB.Create(&todos); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &todos)
}

// @Summary update a todo item by ID
// @ID update-todo-by-id
// @Produce json
// @Param id path string true "todo ID"
// @Param data body model.TodoRequest true "Todo Request"
// @Success 200 {object} model.Todo
// @Failure 404 {object} model.Message
// @Router /todos/{id} [put]
func (h *handler) UpdateTodoHandler(ctx *gin.Context) {

	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	input := model.TodoRequest{}
	if err := ctx.BindJSON(&input); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var todos model.Todo

	if result := h.DB.First(&todos, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	todos.Name = input.Name
	todos.Deadline = input.Deadline
	todos.Description = input.Description
	todos.Status = input.Status

	h.DB.Save(&todos)

	ctx.JSON(http.StatusCreated, &todos)

}

// @Summary delete a todo item by ID
// @ID delete-todo-by-id
// @Produce json
// @Param id path string true "todo ID"
// @Success 200 {object} model.Todo
// @Failure 404 {object} model.Message
// @Router /todos/{id} [delete]
func (h *handler) DeleteTodoHandler(ctx *gin.Context) {

	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	var todo model.Todo

	if result := h.DB.First(&todo, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&todo)

	ctx.Status(http.StatusOK)
}
