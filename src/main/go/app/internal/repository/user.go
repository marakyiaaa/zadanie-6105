package repository

import (
	"context"
	"database/sql"
	"errors"
	"main/internal/models"
	"time"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, userID int64) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, userID int64) error
}

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO employee (username, first_name, last_name, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.DB.QueryRowContext(ctx, query, user.Username, user.FirstName, user.LastName, time.Now(), time.Now()).Scan(&user.ID)
	return err
}

func (r *userRepository) GetUserByID(ctx context.Context, userID int64) (*models.User, error) {
	query := `SELECT id, username, first_name, last_name, created_at, updated_at FROM employee WHERE id = $1`
	row := r.DB.QueryRowContext(ctx, query, userID)

	var user models.User
	if err := row.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Возвращаем nil, если запись не найдена
		}
		return nil, err // Возвращаем ошибку, если произошла другая ошибка
	}
	return &user, nil
}

func (r *userRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	query := `SELECT id, username, first_name, last_name, created_at, updated_at FROM employee WHERE username = $1`
	row := r.DB.QueryRowContext(ctx, query, username)

	var user models.User
	if err := row.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Возвращаем nil, если запись не найдена
		}
		return nil, err // Возвращаем ошибку, если произошла другая ошибка
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *models.User) error {
	query := `UPDATE employee SET username = $1, first_name = $2, last_name = $3, updated_at = $4 WHERE id = $5`
	_, err := r.DB.ExecContext(ctx, query, user.Username, user.FirstName, user.LastName, time.Now(), user.ID)
	return err
}

func (r *userRepository) DeleteUser(ctx context.Context, userID int64) error {
	query := `DELETE FROM employee WHERE id = $1`
	_, err := r.DB.ExecContext(ctx, query, userID)
	return err
}
