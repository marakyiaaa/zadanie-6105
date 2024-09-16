package service

import (
	"context"
	"main/internal/models"
	"main/internal/repository"
)

type OrganizationService interface {
	CreateOrganization(ctx context.Context, organization *models.Organization) error
	GetOrganizationByID(ctx context.Context, organizationID int64) (*models.Organization, error)
	UpdateOrganization(ctx context.Context, organization *models.Organization) error
	DeleteOrganization(ctx context.Context, organizationID int64) error
}

type organizationService struct {
	repo repository.OrganizationRepository
}

func NewOrganizationService(repo repository.OrganizationRepository) OrganizationService {
	return &organizationService{repo: repo}
}

func (s *organizationService) CreateOrganization(ctx context.Context, organization *models.Organization) error {
	return s.repo.CreateOrganization(ctx, organization)
}

func (s *organizationService) GetOrganizationByID(ctx context.Context, organizationID int64) (*models.Organization, error) {
	return s.repo.GetOrganizationByID(ctx, organizationID)
}

func (s *organizationService) UpdateOrganization(ctx context.Context, organization *models.Organization) error {
	return s.repo.UpdateOrganization(ctx, organization)
}

func (s *organizationService) DeleteOrganization(ctx context.Context, organizationID int64) error {
	return s.repo.DeleteOrganization(ctx, organizationID)
}
