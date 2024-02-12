package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// resume struct represents data about a resume
type resume struct {
	ID	string	`json:"id"`
	Year	int	`json:"year"`
	URL	string	`json:"URL"`
}

// resume slice to seed resume data.
var resumes = []resume{
	{ID: "1", Year: 2019, URL: "https://docs.google.com/document/d/1tiFMVTbrzlbHnxUusnT8veVlejbiZpts0dHW1GlJGf8/edit?usp=drive_link"},
	{ID: "2", Year: 2021, URL: "https://docs.google.com/document/d/1RQd595CTc8MBs6GSLChokZXpfXnN4eS-P4KCB0qpazo/edit?usp=drive_link"},
	{ID: "3", Year: 2023, URL: "https://docs.google.com/document/d/1RQd595CTc8MBs6GSLChokZXpfXnN4eS-P4KCB0qpazo/edit?usp=drive_link"},
	{ID: "4", Year: 2024, URL: "https://davidchang.dev"},
}

func main() {
	router := gin.Default()
	// For anyone trying to access https://api.davidchang.dev directly, it's likely it's a person looking for documentation.
	// Redirect to https://docs.davidchang.dev for documentation.
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://docs.davidchang.dev")
	})
	router.GET("/resumes", getResumes)

	router.Run(":8080")
}

// getResumes responds with the list of all resumes as JSON.
func getResumes(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, resumes)
}
