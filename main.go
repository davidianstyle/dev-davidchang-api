package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var db *gorm.DB

// Resume struct represents data about a resumé
type Resume struct {
	gorm.Model        // Handles ID
	Year       int    `json:"year"`
	URL        string `json:"URL"`
}

func init() {
	// Load environment variables from a file if needed
	// For simplicity, you can set them directly in your Cloud Run service configuration.
	// LoadEnvFromFile(".env") // Uncomment this line if you are using a .env file
}

func main() {
	// Initialize the database
	if err := setupDB(); err != nil {
		log.Fatalf("Error setting up the database: %v", err)
	}

	router := gin.Default()

	// Redirect to documentation if requesting https://api.davidchang.dev/ directly
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://docs.davidchang.dev")
	})

	// API to interact with resumés
	router.GET("/resumes", getResumes)
	router.GET("/resumes/:id", getResumeByID)
	router.PATCH("/resumes/:id", updateResume)
	router.POST("/resumes", postResumes)

	router.Run(":8080")
}

func setupDB() error {
	connectionString := os.Getenv("DB_CONNECTION_STRING")

	if connectionString == "" {
		// Get Cloud SQL connection details from environment variables if connectionString is not set
		dbName := os.Getenv("DB_NAME")
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbConnectionName := os.Getenv("DB_CONNECTION_NAME")

		connectionString = fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s?charset=utf8&parseTime=True&loc=Local",
			dbUser, dbPassword, dbConnectionName, dbName)
	}
	// Open a database connection
	database, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// Auto Migrate the resumé model, Resume
	if err := database.AutoMigrate(&Resume{}); err != nil {
		return fmt.Errorf("failed to migrate database: %v", err)
	}

	// Assign the database instance to the global variable
	db = database

	return nil
}

func getResumes(c *gin.Context) {
	var resumes []Resume
	db.Find(&resumes)

	c.IndentedJSON(http.StatusOK, resumes)
}

func getResumeByID(c *gin.Context) {
	var r Resume
	if err := db.First(&r, c.Param("id")).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Resumé not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, r)
}

func updateResumeByID(c *gin.Context) {
	var existingResume Resume
	if err := db.First(&existingResume, c.Param("id")).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Resume not found"})
		return
	}

	var updatedResume Resume
	if err := c.ShouldBindJSON(&updatedResume); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the existing resume with the new data
	db.Model(&existingResume).Updates(updatedResume)

	c.IndentedJSON(http.StatusOK, existingResume)
}

func postResumes(c *gin.Context) {
	var newResume Resume

	if err := c.BindJSON(&newResume); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resumeURL := strings.TrimSpace(newResume.URL)
	if !strings.HasPrefix(resumeURL, "https://") {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid resume URL"})
		return
	}

	// Add the new resume to the database
	if err := db.Create(&newResume).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create resume"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newResume)
}
