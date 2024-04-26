package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	// "github.com/google/uuid"
)

type FoodOrder struct {
	ID         string `json:"id"`
	FoodName   string `json:"foodname"`
	WaiterName string `jaon:"waitername"`
}

var Orders = []FoodOrder{
	// {},
}

var lastId = 0

func GetAllOrders(f *gin.Context) {
	f.IndentedJSON(http.StatusOK, Orders)
}

func GetOrderById(f *gin.Context) {
	id := f.Param("id")

	for _, g := range Orders {
		if g.ID == id {
			f.JSON(http.StatusOK, g)
			return
		}
	}

	f.JSON(http.StatusNotFound, gin.H{
		"error": "Order Not Found",
	})
}

func MakeOrder(f *gin.Context) {
	var newOrder FoodOrder
	if err := f.BindJSON(&newOrder); err != nil {
		f.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	lastId++
	newOrder.ID = strconv.Itoa(lastId)
	fmt.Println("lastId:", lastId)
	// newOrder.ID = uuid.New().String() // Generating a new UUID as ID
	Orders = append(Orders, newOrder)
	f.IndentedJSON(http.StatusCreated, newOrder)

}

func ChangeOrder(f *gin.Context) {
	id := f.Param("id")
	var editOrder FoodOrder
	if err := f.BindJSON(&editOrder); err != nil {
		f.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	for i, r := range Orders {
		if r.ID == id {
			Orders[i] = editOrder
			f.JSON(http.StatusOK, gin.H{
				"message": "Order Changed Successfully", "Order": editOrder,
			})
			return
		}
	}

	f.JSON(http.StatusNotFound, gin.H{
		"error": "Id Not Found",
	})
}

func CancelOrder(f *gin.Context) {
	id := f.Param("id")

	for i, c := range Orders {
		if c.ID == id {
			Orders = append(Orders[:i], Orders[i+1:]...)
			f.JSON(http.StatusOK, c)
			return
		}
	}

	f.JSON(http.StatusNotFound, gin.H{
		"error": "Order Not be Cancelled by Anyway",
	})

}

func main() {

	navigate := gin.Default()
	navigate.GET("/allorders", GetAllOrders)
	navigate.GET("/order/:id", GetOrderById)
	navigate.POST("/order", MakeOrder)
	navigate.PUT("/order/:id", ChangeOrder)
	navigate.DELETE("/cancel/:id", CancelOrder)

	navigate.Run(":7000")

}
