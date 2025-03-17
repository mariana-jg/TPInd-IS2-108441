// Package repositories contains the methods that interact with the database
package repositories

import (
	"apirest-is2/internal/database"
	"apirest-is2/internal/models"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// CoursesRepositoryInterface is the interface that defines the methods that interact with the database
type CoursesRepositoryInterface interface {
	GetCourses() ([]models.Course, error)
	GetCourse(id int) (models.Course, error)
	CreateCourse(course models.Course) (models.Course, error)
	DeleteCourse(id int) error
}

// CourseRepository is the struct that contains the database connection
type CourseRepository struct {
	db *pgxpool.Pool
}

// NewCourseRepository creates a new CourseRepository with the database connection
func NewCourseRepository() (*CourseRepository, error) {
	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
	)

	pool, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return &CourseRepository{db: pool}, nil
}

// DB returns the database connection
func (r *CourseRepository) DB() *pgxpool.Pool {
	return r.db
}

// GetCourses returns all the courses from the database
func (r *CourseRepository) GetCourses() ([]models.Course, error) {
	rows, err := database.DB.Query(context.Background(), "SELECT id, title, description FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []models.Course
	for rows.Next() {
		var course models.Course
		err := rows.Scan(&course.ID, &course.Title, &course.Description)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, nil

}

// GetCourse returns a course from the database
func (r *CourseRepository) GetCourse(id int) (models.Course, error) {
	var course models.Course
	err := database.DB.QueryRow(context.Background(),
		"SELECT id, title, description FROM courses WHERE id = $1", id).
		Scan(&course.ID, &course.Title, &course.Description)

	if err != nil {
		return models.Course{}, &CourseNotFoundError{ID: id}
	}
	return course, nil
}

// CreateCourse creates a course in the database
func (r *CourseRepository) CreateCourse(course models.Course) (models.Course, error) {
	err := database.DB.QueryRow(context.Background(),
		"INSERT INTO courses (title, description) VALUES ($1, $2) RETURNING id",
		course.Title, course.Description).
		Scan(&course.ID)

	if err != nil {
		return models.Course{}, err
	}
	return course, nil
}

// DeleteCourse deletes a course from the database
func (r *CourseRepository) DeleteCourse(id int) error {
	result, err := database.DB.Exec(context.Background(),
		"DELETE FROM courses WHERE id = $1", id)

	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return &CourseNotFoundError{ID: id}
	}
	return nil
}
