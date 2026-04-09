package walletoperation

import "errors"

var (
	ErrEmptyParticipantID = errors.New("participant id is required")
)

type ParticipantID string

type Operation struct {
	FromID ParticipantID `json:"from_ID"`
	ToID   ParticipantID `json:"to_ID"`
	Amount int           `json:"amount"`
}

func NewOperation(fromID, toID string, amount int) (Operation, error) {
	if fromID == "" || toID == "" {
		return Operation{}, ErrEmptyParticipantID
	}

	return Operation{
		FromID: ParticipantID(fromID),
		ToID:   ParticipantID(toID),
		Amount: amount,
	}, nil
}

type Status string

const (
	StatusSuccess Status = "Success"
	StatusCancel  Status = "Cancel"
)
