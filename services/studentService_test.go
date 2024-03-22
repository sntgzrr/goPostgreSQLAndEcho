package services

import (
	"echoApiRestSQL/models"
	"testing"
)

func TestStudentService(t *testing.T) {
	student := models.Student{
		Name:   "Santiago",
		Age:    21,
		Active: true,
	}
	if err := CreateStudentService(student); err != nil {
		t.Error("Error creating student")
	} else {
		t.Log("Create student passed")
	}
}

func TestReadStudentsService(t *testing.T) {
	students, err := ReadStudentsService()
	if err != nil {
		t.Error("Error reading students")
	} else if len(students) < 1 {
		t.Error("Possible error, students don't have data")
	} else {
		t.Log("Read students passed")
	}
}

func TestUpdateStudentService(t *testing.T) {
	student := models.Student{
		Name:   "Isabella",
		Age:    21,
		Active: true,
	}
	err := UpdateStudentService(student, 3)
	if err != nil {
		t.Error("Error updating student")
	} else {
		t.Log("Update student passed")
	}
}

func TestDeleteStudentService(t *testing.T) {
	if err := DeleteStudentService(3); err != nil {
		t.Error("Error deleting student")
	} else {
		t.Log("Delete student passed")
	}
}
