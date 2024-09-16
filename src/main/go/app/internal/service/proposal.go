package service

import (
	"context"
	"main/internal/models"
	"main/internal/repository"
)

type ProposalService interface {
	CreateProposal(ctx context.Context, proposal *models.Proposal) error
	GetProposalByID(ctx context.Context, proposalID int64) (*models.Proposal, error)
	UpdateProposal(ctx context.Context, proposal *models.Proposal) error
	DeleteProposal(ctx context.Context, proposalID int64) error
	ListProposalsByTenderID(ctx context.Context, tenderID int64) ([]*models.Proposal, error)
}

type proposalService struct {
	repo repository.ProposalRepository
}

func NewProposalService(repo repository.ProposalRepository) ProposalService {
	return &proposalService{repo: repo}
}

func (s *proposalService) CreateProposal(ctx context.Context, proposal *models.Proposal) error {
	return s.repo.CreateProposal(ctx, proposal)
}

func (s *proposalService) GetProposalByID(ctx context.Context, proposalID int64) (*models.Proposal, error) {
	return s.repo.GetProposalByID(ctx, proposalID)
}

func (s *proposalService) UpdateProposal(ctx context.Context, proposal *models.Proposal) error {
	return s.repo.UpdateProposal(ctx, proposal)
}

func (s *proposalService) DeleteProposal(ctx context.Context, proposalID int64) error {
	return s.repo.DeleteProposal(ctx, proposalID)
}

func (s *proposalService) ListProposalsByTenderID(ctx context.Context, tenderID int64) ([]*models.Proposal, error) {
	return s.repo.ListProposalsByTenderID(ctx, tenderID)
}
