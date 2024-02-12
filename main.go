package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// resume struct represents data about a resume
type resume struct {
	ID   string `json:"id"`
	Year int    `json:"year"`
	URL  string `json:"URL"`
}

// resume slice to seed resume data
var resumes = []resume{
	{ID: "1", Year: 2019, URL: "https://docs.google.com/document/d/1tiFMVTbrzlbHnxUusnT8veVlejbiZpts0dHW1GlJGf8/edit?usp=drive_link"},
	{ID: "2", Year: 2021, URL: "https://docs.google.com/document/d/1RQd595CTc8MBs6GSLChokZXpfXnN4eS-P4KCB0qpazo/edit?usp=drive_link"},
	{ID: "3", Year: 2023, URL: "https://docs.google.com/document/d/1RQd595CTc8MBs6GSLChokZXpfXnN4eS-P4KCB0qpazo/edit?usp=drive_link"},
}

// getResumes responds with the list of all resumes as JSON
func getResumes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, resumes)
}

// getResumeByID locates the resume whose ID value matches the id
// parameter sent by the client, then returns that resume as a response.
func getResumeByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of resumes, looking for
	// an resume whose ID value matches the parameter.
	for _, a := range resumes {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "resume not found"})
}

// postResumes adds an resume from JSON received in the request body
func postResumes(c *gin.Context) {
	var newResume resume

	// Call BindJSON to bind the received JSON to newResume
	if err := c.BindJSON(&newResume); err != nil {
		return
	}

	// Add the new resume to the slice
	resumes = append(resumes, newResume)
	c.IndentedJSON(http.StatusCreated, newResume)
}

func main() {
	router := gin.Default()
	// For anyone trying to access https://api.davidchang.dev directly, it's likely it's a person looking for documentation
	// Redirect to https://docs.davidchang.dev for documentation
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://docs.davidchang.dev")
	})

	// API to interact with resumes
	router.GET("/resumes", getResumes)
	router.GET("/resumes/:id", getResumeByID)
	router.POST("/resumes", postResumes)

	router.Run(":8080")
}
