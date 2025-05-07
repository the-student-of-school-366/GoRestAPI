package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

type Message struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user=postgres password=ananda2005 dbname=postgres port=5433 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	db.AutoMigrate(&Message{})
}
func GetHandler(c echo.Context) error {
	var messages []Message

	if err := db.Find(&messages).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Could not find the message",
		})
	}

	return c.JSON(http.StatusOK, &messages)
}

func PostHandler(c echo.Context) error {
	var message Message
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Could not add the message",
		})
	}

	if err := db.Create(&message).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Could not create the message",
		})
	}

	return c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "Added the message",
	})

}
func PatchHandler(c echo.Context) error {
	IdParam := c.Param("id")
	id, err := strconv.Atoi(IdParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Wrong ID",
		})
	}
	var updatedMessage Message

	if err := c.Bind(&updatedMessage); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Invalid input",
		})
	}

	if err := db.Model(&Message{}).Where("id = ?", id).Update("text", updatedMessage.Text).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Could not update the message",
		})
	}
	return c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "Updated the message",
	})
}

func DeleteHandler(c echo.Context) error {
	IdParam := c.Param("id")
	id, err := strconv.Atoi(IdParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Wrong ID",
		})
	}
	if err := db.Delete(&Message{}, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Could not delete the message",
		})
	}
	return c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "The message was deleted successfully",
	})
}
func main() {
	initDB()
	e := echo.New()

	e.GET("/messages", GetHandler)
	e.POST("/messages", PostHandler)
	e.PATCH("/messages/:id", PatchHandler)
	e.DELETE("/messages/:id", DeleteHandler)
	e.Start(":8082")
}
