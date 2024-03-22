package main

import (
	"echoApiRestSQL/models"
	"echoApiRestSQL/services"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "8080"
	}
	e := echo.New()
	g := e.Group("/api")
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[method=${method}, host=${host}, uri=${uri}, status=${status}]\n",
	}))
	g.Use(middleware.BasicAuth(func(name string, password string, context echo.Context) (bool, error) {
		if name == "Santiago" && password == "root" {
			return true, nil
		}
		return false, nil
	}))
	g.GET("/student", getStudents)
	g.POST("/student", postStudent)
	g.PUT("/student/:id", putStudent)
	g.DELETE("/student", deleteStudent)
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}

func getStudents(c echo.Context) error {
	students, err := services.ReadStudentsService()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, students)
}

func postStudent(c echo.Context) error {
	var student models.Student
	if err := c.Bind(&student); err != nil {
		return err
	}
	if err := services.CreateStudentService(student); err != nil {
		return err
	}
	return c.String(http.StatusCreated, "Student created successfully")
}

func putStudent(c echo.Context) error {
	var student models.Student
	if err := c.Bind(&student); err != nil {
		return err
	}
	param := c.Param("id")
	userID, _ := strconv.Atoi(param)
	if err := services.UpdateStudentService(student, userID); err != nil {
		return err
	}
	return c.String(http.StatusOK, "Student updated successfully")
}

func deleteStudent(c echo.Context) error {
	param := c.QueryParam("id")
	userID, _ := strconv.Atoi(param)
	if err := services.DeleteStudentService(userID); err != nil {
		return err
	}
	return c.String(http.StatusOK, "Student deleted successfully")
}
