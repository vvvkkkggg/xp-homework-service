package rest_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/driftprogramming/pgxpoolmock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/vvvkkkggg/xp-homework-service/internal/application"
	"github.com/vvvkkkggg/xp-homework-service/internal/rest"
)

func TestHomeIntegration(t *testing.T) {
	app := setupTestApplication(t)
	handler := rest.NewHandler(app)
	router := gin.Default()
	router.GET("/home/:user", handler.Home)

	userID := "123"
	req, err := http.NewRequest("GET", "/home/"+userID, nil)
	require.NoError(t, err, "Error creating request")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Response status should be OK")
}

func TestTaskIntegration(t *testing.T) {
	app := setupTestApplication(t)
	handler := rest.NewHandler(app)
	router := gin.Default()
	router.GET("/task/:user/:task", handler.Task)

	userID := "123"
	taskID := "456"
	req, err := http.NewRequest("GET", "/task/"+userID+"/"+taskID, nil)
	require.NoError(t, err, "Error creating request")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Response status should be OK")
}

func setupTestApplication(t *testing.T) *application.Application {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockPool := pgxpoolmock.NewMockPgxPool(ctrl)

	return application.New(mockPool)
}
