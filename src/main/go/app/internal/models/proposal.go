package models

type ProposalStatus string

const (
	CREATED   ProposalStatus = "CREATED"
	PUBLISHED ProposalStatus = "PUBLISHED"
	CANCELED  ProposalStatus = "CANCELED"
)

type Proposal struct {
	ID        int            `json:"id"`
	TenderID  int            `json:"tenderID"`
	UserID    int            `json:"userID"`
	Status    ProposalStatus `json:"status"`
	Version   int            `json:"version"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
}
