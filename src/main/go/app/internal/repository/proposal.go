package repository

import (
	"context"
	"database/sql"
	"errors"
	"main/internal/models"
)

type ProposalRepository interface {
	CreateProposal(ctx context.Context, proposal *models.Proposal) error
	GetProposalByID(ctx context.Context, proposalID int64) (*models.Proposal, error)
	UpdateProposal(ctx context.Context, proposal *models.Proposal) error
	DeleteProposal(ctx context.Context, proposalID int64) error
	ListProposalsByTenderID(ctx context.Context, tenderID int64) ([]*models.Proposal, error)
}

type proposalRepository struct {
	DB *sql.DB
}

func NewProposalRepository(db *sql.DB) ProposalRepository {
	return &proposalRepository{}
}

func (r *proposalRepository) CreateProposal(ctx context.Context, proposal *models.Proposal) error {
	query := `INSERT INTO proposals (tenderID, userID, status, version, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.DB.ExecContext(ctx, query, proposal.TenderID, proposal.UserID, proposal.Status, proposal.Version, proposal.CreatedAt, proposal.UpdatedAt)
	return err
}

func (r *proposalRepository) GetProposalByID(ctx context.Context, proposalID int64) (*models.Proposal, error) {
	query := `SELECT id, tenderID, userID, status, version, created_at, updated_at FROM proposals WHERE id = $1`
	row := r.DB.QueryRowContext(ctx, query, proposalID)

	var proposal models.Proposal
	if err := row.Scan(&proposal.ID, &proposal.TenderID, &proposal.UserID, &proposal.Status, &proposal.Version, &proposal.CreatedAt, &proposal.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Возвращаем nil, если запись не найдена
		}
		return nil, err // Возвращаем ошибку, если произошла другая ошибка
	}
	return &proposal, nil
}

func (r *proposalRepository) UpdateProposal(ctx context.Context, proposal *models.Proposal) error {
	query := `UPDATE proposals SET tenderID = $1, userID = $2, status = $3, version = $4, updated_at = $5 WHERE id = $6`
	_, err := r.DB.ExecContext(ctx, query, proposal.TenderID, proposal.UserID, proposal.Status, proposal.Version, proposal.UpdatedAt, proposal.ID)
	return err
}
func (r *proposalRepository) DeleteProposal(ctx context.Context, proposalID int64) error {
	query := `DELETE FROM proposals WHERE id = $1`
	_, err := r.DB.ExecContext(ctx, query, proposalID)
	return err
}

func (r *proposalRepository) ListProposalsByTenderID(ctx context.Context, tenderID int64) ([]*models.Proposal, error) {
	query := `SELECT id, tenderID, userID, status, version, created_at, updated_at FROM proposals WHERE tenderID = $1`
	rows, err := r.DB.QueryContext(ctx, query, tenderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var proposals []*models.Proposal
	for rows.Next() {
		var proposal models.Proposal
		if err := rows.Scan(&proposal.ID, &proposal.TenderID, &proposal.UserID, &proposal.Status, &proposal.Version, &proposal.CreatedAt, &proposal.UpdatedAt); err != nil {
			return nil, err
		}
		proposals = append(proposals, &proposal)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return proposals, nil
}
