package repository

import (
	"context"
	"database/sql"
	"errors"
	"main/internal/models"
)

type TenderRepository interface {
	CreateTender(ctx context.Context, tender *models.Tender) error
	PublishTender(ctx context.Context, tenderID int64) error
	CloseTender(ctx context.Context, tenderID int64) error
	UpdateTender(ctx context.Context, tender *models.Tender) error
	GetTenderByID(ctx context.Context, tenderID int64) (*models.Tender, error)
}

type tenderRepository struct {
	DB *sql.DB
}

func NewTenderRepository(db *sql.DB) TenderRepository {
	return &tenderRepository{DB: db}
}

func (r *tenderRepository) CreateTender(ctx context.Context, tender *models.Tender) error {
	query := `INSERT INTO tenders (organization_id, status, version, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5)`
	_, err := r.DB.ExecContext(ctx, query, tender.OrganizationID, tender.Status, tender.Version, tender.CreatedAt, tender.UpdatedAt)
	return err
}

func (r *tenderRepository) PublishTender(ctx context.Context, tenderID int64) error {
	query := `UPDATE tenders SET status = $1 WHERE id = $2`
	_, err := r.DB.ExecContext(ctx, query, models.StatusPublished, tenderID)
	return err
}

func (r *tenderRepository) CloseTender(ctx context.Context, tenderID int64) error {
	query := `UPDATE tenders SET status = $1 WHERE id = $2`
	_, err := r.DB.ExecContext(ctx, query, models.StatusClosed, tenderID)
	return err
}

func (r *tenderRepository) UpdateTender(ctx context.Context, tender *models.Tender) error {
	query := `UPDATE tenders SET organization_id = $1, status = $2, version = $3, updated_at = $4 WHERE id = $5`
	_, err := r.DB.ExecContext(ctx, query, tender.OrganizationID, tender.Status, tender.Version, tender.UpdatedAt, tender.ID)
	return err
}

func (r *tenderRepository) GetTenderByID(ctx context.Context, tenderID int64) (*models.Tender, error) {
	query := `SELECT id, organization_id, status, version, created_at, updated_at FROM tenders WHERE id = $1`
	row := r.DB.QueryRowContext(ctx, query, tenderID)

	var tender models.Tender
	if err := row.Scan(&tender.ID, &tender.OrganizationID, &tender.Status, &tender.Version, &tender.CreatedAt, &tender.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &tender, nil
}
