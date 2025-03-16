package repositories

import (
	"apirest-is2/internal/database"
	"apirest-is2/internal/models"
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CourseRepository struct {
	//courses []models.Course
	db *pgxpool.Pool
}

var ErrCourseNotFound = errors.New("course not found")

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

func (r *CourseRepository) DB() *pgxpool.Pool {
	return r.db
}

func (r *CourseRepository) GetCourses() ([]models.Course, error) {
	//return r.courses
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

func (r *CourseRepository) GetCourse(id int) (models.Course, error) {
	/*for _, course := range r.courses {
		if course.ID == id {
			return course, nil
		}
	}
	return models.Course{}, ErrCourseNotFound*/
	var course models.Course
	err := database.DB.QueryRow(context.Background(),
		"SELECT id, title, description FROM courses WHERE id = $1", id).
		Scan(&course.ID, &course.Title, &course.Description)

	if err != nil {
		return models.Course{}, ErrCourseNotFound
	}
	return course, nil
}

func (r *CourseRepository) CreateCourse(course models.Course) (models.Course, error) {
	/*course.ID = len(r.courses) + 1
	r.courses = append(r.courses, course)
	return course, nil*/
	err := database.DB.QueryRow(context.Background(),
		"INSERT INTO courses (title, description) VALUES ($1, $2) RETURNING id",
		course.Title, course.Description).
		Scan(&course.ID)

	if err != nil {
		return models.Course{}, err
	}
	return course, nil
}

func (r *CourseRepository) DeleteCourse(id int) error {
	/*for i, course := range r.courses {
		if course.ID == id {
			r.courses = append(r.courses[:i], r.courses[i+1:]...)
			return nil
		}
	}
	return ErrCourseNotFound*/
	result, err := database.DB.Exec(context.Background(),
		"DELETE FROM courses WHERE id = $1", id)

	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return ErrCourseNotFound
	}
	return nil
}
