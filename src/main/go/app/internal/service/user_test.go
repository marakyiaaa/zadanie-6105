package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"main/internal/models"
	"testing"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user *models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserByID(ctx context.Context, userID int64) (*models.User, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	args := m.Called(ctx, username)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(ctx context.Context, user *models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) DeleteUser(ctx context.Context, userID int64) error {
	args := m.Called(ctx, userID)
	return args.Error(0)
}

func TestUserService_CreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	user := &models.User{
		Username:  "testuser",
		FirstName: "Test",
		LastName:  "User",
	}

	mockRepo.On("CreateUser", mock.Anything, user).Return(nil)

	err := service.CreateUser(context.Background(), user)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUserByID(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	user := &models.User{
		ID:        1,
		Username:  "testuser",
		FirstName: "Test",
		LastName:  "User",
	}

	mockRepo.On("GetUserByID", mock.Anything, int64(1)).Return(user, nil)

	result, err := service.GetUserByID(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockRepo.AssertExpectations(t)
}

func TestUserService_UpdateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	user := &models.User{
		ID:        1,
		Username:  "testuser",
		FirstName: "Test",
		LastName:  "User",
	}

	mockRepo.On("UpdateUser", mock.Anything, user).Return(nil)

	err := service.UpdateUser(context.Background(), user)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUserService_DeleteUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	mockRepo.On("DeleteUser", mock.Anything, int64(1)).Return(nil)

	err := service.DeleteUser(context.Background(), 1)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
