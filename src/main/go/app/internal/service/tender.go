package service

import (
	"context"
	"main/internal/models"
	"main/internal/repository"
)

type TenderService interface {
	CreateTender(ctx context.Context, tender *models.Tender) error
	PublishTender(ctx context.Context, tenderID int64) error
	CloseTender(ctx context.Context, tenderID int64) error
	UpdateTender(ctx context.Context, tender *models.Tender) error
	GetTenderByID(ctx context.Context, tenderID int64) (*models.Tender, error)
}

type tenderService struct {
	repo repository.TenderRepository
}

func NewTenderService(repo repository.TenderRepository) TenderService {
	return &tenderService{repo: repo}
}

func (s *tenderService) CreateTender(ctx context.Context, tender *models.Tender) error {
	tender.Status = models.StatusCreated
	return s.repo.CreateTender(ctx, tender)
}

func (s *tenderService) PublishTender(ctx context.Context, tenderID int64) error {
	return s.repo.PublishTender(ctx, tenderID)
}

func (s *tenderService) CloseTender(ctx context.Context, tenderID int64) error {
	return s.repo.CloseTender(ctx, tenderID)
}

func (s *tenderService) UpdateTender(ctx context.Context, tender *models.Tender) error {
	tender.Version++
	return s.repo.UpdateTender(ctx, tender)
}

func (s *tenderService) GetTenderByID(ctx context.Context, tenderID int64) (*models.Tender, error) {
	return s.repo.GetTenderByID(ctx, tenderID)
}
