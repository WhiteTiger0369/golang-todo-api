package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/notes", GetNotes)
	router.GET("/notes/:id", GetNoteById)
	router.POST("/notes", PostNotes)
	router.PUT("/notes/:id", UpdateNote)
	router.DELETE("notes/:id", DeleteNote)

	router.Run("localhost:8080")
}
