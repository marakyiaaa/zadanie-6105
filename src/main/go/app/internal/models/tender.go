package models

import "time"

type TenderStatus string

const (
	StatusCreated   TenderStatus = "CREATED"
	StatusPublished TenderStatus = "PUBLISHED"
	StatusClosed    TenderStatus = "CLOSED"
)

type Tender struct {
	ID             int          `json:"id"`
	OrganizationID int          `json:"organizationID"`
	Status         TenderStatus `json:"status"`
	Version        int          `json:"version"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
}
