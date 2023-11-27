package application_test

import (
	"context"
	"testing"

	"github.com/driftprogramming/pgxpoolmock"
	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/vvvkkkggg/xp-homework-service/internal/application"
)

func TestHomeIntegration(t *testing.T) {
	db, cleanup := setupTestDatabase(t)
	defer cleanup()

	app := application.New(db)
	userID := 1
	tmpl, homePage, err := app.Home(context.Background(), userID)

	require.NoError(t, err, "Home function should not return an error")
	require.NotNil(t, tmpl, "Template should not be nil")
	require.NotNil(t, homePage, "HomePage should not be nil")

	assert.Equal(t, userID, homePage.UserInfo.UserID, "User ID should match")
}

func TestTaskIntegration(t *testing.T) {
	db, cleanup := setupTestDatabase(t)
	defer cleanup()

	app := application.New(db)

	userID := 1
	taskID := 123
	tmpl, taskPage, err := app.Task(context.Background(), userID, taskID)

	require.NoError(t, err, "Task function should not return an error")
	require.NotNil(t, tmpl, "Template should not be nil")
	require.NotNil(t, taskPage, "AssignmentPage should not be nil")

	assert.Equal(t, userID, taskPage.UserInfo.UserID, "User ID should match")
	assert.Equal(t, taskID, taskPage.TaskInfo.TaskID, "Task ID should match")
}

func setupTestDatabase(t *testing.T) (*pgxpoolmock.MockPgxPool, func()) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockPool := pgxpoolmock.NewMockPgxPool(ctrl)

	cleanup := func() {
		mockPool.Close()
	}
	return mockPool, cleanup
}
