package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func TestEndToEndOK(t *testing.T) {

	url := "http://localhost:" + getPort() + "/courses"

	var createdCourseID int

	t.Run("Create Course", func(t *testing.T) {
		data := map[string]any{
			"title":       "Software Engineer II",
			"description": "Advanced software engineering concepts and good practices in software development",
		}
		payload, _ := json.Marshal(data)

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		var responseBody map[string]any
		err = json.NewDecoder(resp.Body).Decode(&responseBody)
		assert.NoError(t, err)

		data, ok := responseBody["data"].(map[string]interface{})
		assert.True(t, ok, "Response 'data' field is missing or has wrong type")

		idFloat, ok := data["id"].(float64)
		assert.True(t, ok, "Course ID missing or incorrect format")

		createdCourseID = int(idFloat)
		assert.NotZero(t, createdCourseID, "Created course ID should not be zero")
	})

	t.Run("Get Specific Course", func(t *testing.T) {
		url := fmt.Sprintf("%s/%d", url, createdCourseID)
		resp, err := http.Get(url)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var result struct {
			Data struct {
				ID int `json:"id"`
			} `json:"data"`
		}

		err = json.NewDecoder(resp.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Equal(t, createdCourseID, result.Data.ID, "Returned course ID should match the created one")
	})

	t.Run("Delete Course", func(t *testing.T) {
		url := fmt.Sprintf("%s/%d", url, createdCourseID)
		req, _ := http.NewRequest("DELETE", url, nil)
		client := &http.Client{}
		resp, err := client.Do(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	})

	t.Run("Get Deleted Course", func(t *testing.T) {
		url := fmt.Sprintf("%s/%d", url, createdCourseID)
		resp, err := http.Get(url)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})
}

func TestEndToEndErrors(t *testing.T) {

	url := "http://localhost:" + getPort() + "/courses"

	t.Run("Create Course Without Data", func(t *testing.T) {
		payload := []byte(`{}`)
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Get a Course Inexistent", func(t *testing.T) {
		resp, err := http.Get(fmt.Sprintf("%s/9999", url))
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})

	t.Run("Delete Inexistent Course", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/9999", url), nil)
		client := &http.Client{}
		resp, err := client.Do(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})

	t.Run("Delete Course With Invalid ID", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/invalid", url), nil)
		client := &http.Client{}
		resp, err := client.Do(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Create Course With Malformed JSON", func(t *testing.T) {
		malformedJSON := []byte(`{"title": "Bad JSON"`)
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(malformedJSON))

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
}
