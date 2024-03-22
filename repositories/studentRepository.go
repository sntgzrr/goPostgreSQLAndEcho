package repositories

import (
	"database/sql"
	"echoApiRestSQL/database"
	"echoApiRestSQL/models"
	"errors"
)

// CreateStudent insert values into students table
func CreateStudent(s models.Student) error {
	// Query to execute
	query := `INSERT INTO students(name, age, active) VALUES ($1, $2, $3)`
	// Variables for Nulls
	intNull := sql.NullInt64{}
	// Get DB connection
	db := database.GetConnection()
	// Close DB connection at the end
	defer db.Close()
	// Prepare a statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	// Close the statement at the end
	defer stmt.Close()
	// Validating if the Age is null
	if s.Age == 0 {
		intNull.Valid = false
	} else {
		intNull.Valid = true
		intNull.Int64 = int64(s.Age)
	}
	// Execute the statement (Exec is used for insert, delete and update)
	r, err := stmt.Exec(s.Name, intNull, s.Active)
	if err != nil {
		return err
	}
	// Returns how many rows were affected
	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("more than 1 row was affected")
	}
	return nil
}

// ReadStudents get values from students table
func ReadStudents() (models.Students, error) {
	var students models.Students
	// Null variables
	updatedAtNull := sql.NullTime{}
	ageNull := sql.NullInt64{}
	activeNull := sql.NullBool{}
	// Query to execute
	query := `SELECT id, name, age, active, created_at, updated_at 
				FROM students`
	// Get DB connection
	db := database.GetConnection()
	// Close DB connection
	defer db.Close()
	// Get Rows from the Query executed
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	// Close Rows interaction
	defer rows.Close()
	// Lopping Rows
	for rows.Next() {
		// Creating a new student
		s := models.Student{}
		// Scanning Rows and adding them to the Student struct
		err := rows.Scan(
			&s.ID,
			&s.Name,
			&ageNull,
			&activeNull,
			&s.CreatedAt,
			&updatedAtNull,
		)
		if err != nil {
			return nil, err
		}
		s.Age = int16(ageNull.Int64)
		s.UpdatedAt = updatedAtNull.Time
		s.Active = activeNull.Bool
		// Adding student into students
		students = append(students, &s)
	}
	return students, nil
}

// UpdateStudent updates students
func UpdateStudent(s models.Student, userID int) error {
	// Query to execute
	query := `UPDATE students 
				SET name = $1, age = $2, active = $3, updated_at = now() 
					WHERE id = $4`
	// Get DB connection
	db := database.GetConnection()
	// Close DB connection
	defer db.Close()
	// Prepare a statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	// Close statement
	defer stmt.Close()
	// Execute the statement
	rows, err := stmt.Exec(s.Name, s.Age, s.Active, userID)
	if err != nil {
		return err
	}
	// Returns how many rows were affected
	i, _ := rows.RowsAffected()
	if i != 1 {
		return errors.New("more than 1 row was affected")
	}
	return nil
}

// DeleteStudent deletes a student
func DeleteStudent(userID int) error {
	// Query to execute
	query := `DELETE FROM students WHERE id = $1`
	// Get DB connection
	db := database.GetConnection()
	// Close DB Connection
	defer db.Close()
	// Create a statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	// Close statement
	defer stmt.Close()
	// Execute statement
	rows, err := stmt.Exec(userID)
	if err != nil {
		return err
	}
	// Returns how many rows were affected
	i, _ := rows.RowsAffected()
	if i != 1 {
		return errors.New("more than 1 row was affected")
	}
	return nil
}
