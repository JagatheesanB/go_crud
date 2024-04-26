package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	TaskName  string `json:"taskname"`
	Completed string `json:"completed"`
}

var todos = []todo{
	{ID: "1", TaskName: "Meeting", Completed: "Yes"},
	{ID: "2", TaskName: "Coding", Completed: "No"},
}

func GetTodos(j *gin.Context) {
	j.IndentedJSON(http.StatusOK, todos)
}

func GetTodoById(j *gin.Context) {
	id := j.Param("id")

	for _, r := range todos {
		if r.ID == id {
			j.JSON(http.StatusOK, r)
			return
		}
	}

	j.JSON(http.StatusNotFound, gin.H{
		"error": "ID not found",
	})
}

func PostTodo(j *gin.Context) {
	var newTodo todo
	if err := j.BindJSON(&newTodo); err != nil {
		j.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	todos = append(todos, newTodo)
	j.IndentedJSON(http.StatusCreated, newTodo)
}

func PutTodo(j *gin.Context) {
	id := j.Param("id")
	var editTodo todo
	err := j.BindJSON(&editTodo)
	if err != nil {
		j.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	for i, v := range todos {
		if v.ID == id {
			todos[i] = editTodo
			j.JSON(http.StatusOK, gin.H{
				"message": "Todo Edited Successfully", "todo": editTodo,
			})
			return
		}
	}

	j.JSON(http.StatusNotFound, gin.H{
		"error": "Todo not found",
	})
}

func DeleteTodoById(j *gin.Context) {
	id := j.Param("id")

	for i, b := range todos {
		if b.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			j.JSON(http.StatusOK, b)
			return
		}
	}
	j.JSON(http.StatusNotFound, gin.H{
		"error": "Todo Not Found",
	})
}

func main() {
	api := gin.Default()

	api.GET("/alltodos", GetTodos)
	api.GET("/todo/:id", GetTodoById)
	api.POST("/todo", PostTodo)
	api.PUT("/todo/:id", PutTodo)
	api.DELETE("/todo/:id", DeleteTodoById)
	api.Run(":8000")
}
