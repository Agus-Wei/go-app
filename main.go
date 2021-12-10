package main

import (
	"net/http"

	"example/crud/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Open()
	sqlDB, _ := db.DB.DB()

	defer sqlDB.Close()

	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello world!"})
	})

	router.GET("/items", getItems)
	router.GET("/item/:itemId", getItem)
	router.POST("/item/:itemId", postItem)
	router.PATCH("/item/:itemId", patchItem)
	router.DELETE("/item/:itemId", deleteItem)

	router.Run()
}

func getItems(c *gin.Context) {
	var items []db.Item
	err := db.GetItems(&items)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, items)
}

func getItem(c *gin.Context) {
	itemId := c.Param("itemId")
	var item db.Item
	err := db.GetItem(&item, itemId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"error": err})
		return
	}

	c.IndentedJSON(http.StatusAccepted, item)
}

func postItem(c *gin.Context) {
	var item db.Item
	if err := c.BindJSON(&item); err != nil {
		return
	}

	db.CreateItem(&item)
	c.IndentedJSON(http.StatusCreated, item)
}

type ItemUpdate struct {
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

func patchItem(c *gin.Context) {
	itemId := c.Param("itemId")
	var itemUpdate ItemUpdate
	if err := c.BindJSON(&itemUpdate); err != nil {
		return
	}

	var item db.Item
	err := db.GetItem(&item, itemId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"error": err})
		return
	}

	item.Name = itemUpdate.Name
	item.Stock = itemUpdate.Stock

	db.UpdateItem(&item)
	c.IndentedJSON(http.StatusAccepted, item)
}

func deleteItem(c *gin.Context) {
	itemId := c.Param("itemId")

	var item db.Item
	err := db.GetItem(&item, itemId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"error": err})
		return
	}

	db.DeleteItem(&item)
	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Item has been removed."})
}
