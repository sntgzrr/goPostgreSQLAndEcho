package services

import (
	"echoApiRestSQL/models"
	"echoApiRestSQL/repositories"
)

func CreateStudentService(s models.Student) error {
	if err := repositories.CreateStudent(s); err != nil {
		return err
	}
	return nil
}

func ReadStudentsService() (models.Students, error) {
	students, err := repositories.ReadStudents()
	if err != nil {
		return nil, err
	}
	return students, nil
}

func UpdateStudentService(student models.Student, userID int) error {
	if err := repositories.UpdateStudent(student, userID); err != nil {
		return err
	}
	return nil
}

func DeleteStudentService(userID int) error {
	if err := repositories.DeleteStudent(userID); err != nil {
		return err
	}
	return nil
}
