package repository

import (
	"context"
	"database/sql"
	"errors"
	"main/internal/models"
	"time"
)

type OrganizationRepository interface {
	CreateOrganization(ctx context.Context, organization *models.Organization) error
	GetOrganizationByID(ctx context.Context, organizationID int64) (*models.Organization, error)
	UpdateOrganization(ctx context.Context, organization *models.Organization) error
	DeleteOrganization(ctx context.Context, organizationID int64) error
}

type organizationRepository struct {
	DB *sql.DB
}

func NewOrganizationRepository(db *sql.DB) OrganizationRepository {
	return &organizationRepository{DB: db}
}

func (r *organizationRepository) CreateOrganization(ctx context.Context, organization *models.Organization) error {
	query := `INSERT INTO organization (name, description, type, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.DB.QueryRowContext(ctx, query, organization.Name, organization.Description, organization.Type, time.Now(), time.Now()).Scan(&organization.ID)
	return err
}

func (r *organizationRepository) GetOrganizationByID(ctx context.Context, organizationID int64) (*models.Organization, error) {
	query := `SELECT id, name, description, type, created_at, updated_at FROM organization WHERE id = $1`
	row := r.DB.QueryRowContext(ctx, query, organizationID)

	var organization models.Organization
	if err := row.Scan(&organization.ID, &organization.Name, &organization.Description, &organization.Type, &organization.CreatedAt, &organization.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Возвращаем nil, если запись не найдена
		}
		return nil, err // Возвращаем ошибку, если произошла другая ошибка
	}
	return &organization, nil
}

func (r *organizationRepository) UpdateOrganization(ctx context.Context, organization *models.Organization) error {
	query := `UPDATE organization SET name = $1, description = $2, type = $3, updated_at = $4 WHERE id = $5`
	_, err := r.DB.ExecContext(ctx, query, organization.Name, organization.Description, organization.Type, time.Now(), organization.ID)
	return err
}

func (r *organizationRepository) DeleteOrganization(ctx context.Context, organizationID int64) error {
	query := `DELETE FROM organization WHERE id = $1`
	_, err := r.DB.ExecContext(ctx, query, organizationID)
	return err
}
