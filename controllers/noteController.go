package controllers

import (
	"example/gotion/initializers"
	"example/gotion/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNotes(c *gin.Context) {
	user, _ := c.Get("user")
	var notes []models.Note
	initializers.DB.Where("user_id = ?", int(user.(models.User).ID)).Find(&notes)
	c.JSON(http.StatusOK, gin.H{"message": notes})
}

func GetNote(c *gin.Context) {
	id := c.Param("id")
	user, _ := c.Get("user")
	var note models.Note
	initializers.DB.Where("user_id = ?", int(user.(models.User).ID)).First(&note, "ID = ?", id)
	c.JSON(http.StatusOK, gin.H{"message": note})
}

func CreateNote(c *gin.Context) {
	user, _ := c.Get("user")

	var body struct {
		Title   string
		Content string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	note := models.Note{Title: body.Title, Content: body.Content, UserID: int(user.(models.User).ID), User: user.(models.User)}
	result := initializers.DB.Create(&note)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create note"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": note})
}

func DeleteNote(c *gin.Context) {
	id := c.Param("id")
	user, _ := c.Get("user")
	var note models.Note
	initializers.DB.Where("user_id = ?", int(user.(models.User).ID)).First(&note, "ID = ?", id)
	if note.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note id"})
		return
	}
	initializers.DB.Delete(&note)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
