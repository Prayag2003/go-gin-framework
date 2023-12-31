package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Course struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}

// Fake DB
var courses []Course

func getCourse(c *gin.Context) {
	// Query Params
	params := c.Query("name")
	fmt.Println(params)

	for _, course := range courses {
		if params == course.Name {
			c.JSON(200, gin.H{
				"name":   course.Name,
				"author": course.Author,
			})
			return
		}
	}
	c.JSON(400, gin.H{
		params: "",
	})
}

func getAllCourses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courses)
}

func createCourse(c *gin.Context) {
	var newCourse Course

	// binding the JSON into
	if err := c.BindJSON(&newCourse); err != nil {
		return
	}

	// add new course to the Course Slice/ DB
	courses = append(courses, newCourse)
	c.IndentedJSON(http.StatusOK, newCourse)
}

func getAllCoursesByAuthor(c *gin.Context) {
	var f bool
	f = false

	author := c.Param("author")

	for _, course := range courses {
		if course.Author == author {
			c.IndentedJSON(http.StatusOK, course)
			f = true
		}
	}
	if !f {
		c.IndentedJSON(http.StatusNotFound, gin.H{"errMessage": "No courses for this author is found."})
	}
}

func main() {

	// Logging to a file.
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	// Router
	r := gin.Default()

	// r.Use(gin.Logger())

	// Populating / seeding data
	courses = append(courses, Course{"Golang", "Prayag Bhatt"})
	courses = append(courses, Course{"Typescript", "Thapa Technical"})
	courses = append(courses, Course{"Amazon Web Services", "Hitesh Chaudhary"})

	// API Endpoints
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "You are at your home, i.e, http://127.0.0.1 Hello World",
		})
	})

	r.GET("/getcoursebyquery", getCourse)
	r.GET("/courses", getAllCourses)
	r.POST("/createcourse", createCourse)
	r.GET("/getbyauth/{author}", getAllCoursesByAuthor)
	r.Run(":8000")
}
