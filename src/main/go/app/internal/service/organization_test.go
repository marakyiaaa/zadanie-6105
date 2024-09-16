package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"main/internal/models"
	"testing"
)

type MockOrganizationRepository struct {
	mock.Mock
}

func (m *MockOrganizationRepository) CreateOrganization(ctx context.Context, organization *models.Organization) error {
	args := m.Called(ctx, organization)
	return args.Error(0)
}

func (m *MockOrganizationRepository) GetOrganizationByID(ctx context.Context, organizationID int64) (*models.Organization, error) {
	args := m.Called(ctx, organizationID)
	return args.Get(0).(*models.Organization), args.Error(1)
}

func (m *MockOrganizationRepository) UpdateOrganization(ctx context.Context, organization *models.Organization) error {
	args := m.Called(ctx, organization)
	return args.Error(0)
}

func (m *MockOrganizationRepository) DeleteOrganization(ctx context.Context, organizationID int64) error {
	args := m.Called(ctx, organizationID)
	return args.Error(0)
}

func TestOrganizationService_CreateOrganization(t *testing.T) {
	mockRepo := new(MockOrganizationRepository)
	service := NewOrganizationService(mockRepo)

	organization := &models.Organization{
		Name:        "Test Org",
		Description: "Test Organization",
	}

	mockRepo.On("CreateOrganization", mock.Anything, organization).Return(nil)

	err := service.CreateOrganization(context.Background(), organization)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestOrganizationService_GetOrganizationByID(t *testing.T) {
	mockRepo := new(MockOrganizationRepository)
	service := NewOrganizationService(mockRepo)

	organization := &models.Organization{
		ID:          1,
		Name:        "Test Org",
		Description: "Test Organization",
	}

	mockRepo.On("GetOrganizationByID", mock.Anything, int64(1)).Return(organization, nil)

	result, err := service.GetOrganizationByID(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, organization, result)
	mockRepo.AssertExpectations(t)
}

func TestOrganizationService_UpdateOrganization(t *testing.T) {
	mockRepo := new(MockOrganizationRepository)
	service := NewOrganizationService(mockRepo)

	organization := &models.Organization{
		ID:          1,
		Name:        "Test Org",
		Description: "Test Organization",
	}

	mockRepo.On("UpdateOrganization", mock.Anything, organization).Return(nil)

	err := service.UpdateOrganization(context.Background(), organization)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestOrganizationService_DeleteOrganization(t *testing.T) {
	mockRepo := new(MockOrganizationRepository)
	service := NewOrganizationService(mockRepo)

	mockRepo.On("DeleteOrganization", mock.Anything, int64(1)).Return(nil)

	err := service.DeleteOrganization(context.Background(), 1)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
