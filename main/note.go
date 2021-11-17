package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type note struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var notes = []note{
	{
		Id:      "1",
		Title:   "Exercise1",
		Content: "To do app",
	},
	{
		Id:      "2",
		Title:   "Exercise2",
		Content: "To do app",
	},
}

func GetNotes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, notes)
}

func GetNoteById(c *gin.Context) {
	id := c.Param("id")

	for _, n := range notes {
		if id == n.Id {
			c.IndentedJSON(http.StatusOK, n)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found with id: " + id})
}

func PostNotes(c *gin.Context) {
	var newNote note
	if err := c.BindJSON(&newNote); err != nil {
		return
	}

	notes = append(notes, newNote)
	c.IndentedJSON(http.StatusCreated, newNote)
}

func UpdateNote(c *gin.Context) {
	id := c.Param("id")
	var noteReq note
	if err := c.BindJSON(&noteReq); err != nil {
		return
	}
	for i, n := range notes {
		if n.Id == id {
			notes[i].Title = noteReq.Title
			notes[i].Content = noteReq.Content
			c.IndentedJSON(http.StatusAccepted, notes[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found with id: " + id})
}

func DeleteNote(c *gin.Context) {
	id := c.Param("id")
	for i, n := range notes {
		if id == n.Id {
			notes = remove(notes, i)
			c.Status(http.StatusNoContent)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found with id: " + id})
}

func remove(s []note, i int) []note {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
