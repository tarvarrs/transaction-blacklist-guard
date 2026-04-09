package walletoperation

import (
	"context"

	domain "github.com/tarvarrs/transaction-blacklist-guard/internal/domain/walletoperation"
)

type ProcessCommand struct {
	FromID string
	ToID   string
	Amount int
}

type ProcessResult struct {
	Status string `json:"status"`
}

type Service struct {
	DecisionService domain.DecisionService
}

func NewService(decisionService domain.DecisionService) Service {
	return Service{DecisionService: decisionService}
}

func (s Service) Process(ctx context.Context, cmd ProcessCommand) (ProcessResult, error) {
	op, err := domain.NewOperation(cmd.FromID, cmd.ToID, cmd.Amount)
	if err != nil {
		return ProcessResult{}, err
	}

	status, err := s.DecisionService.Decide(ctx, op)
	if err != nil {
		return ProcessResult{}, err
	}

	return ProcessResult{
		Status: string(status),
	}, nil
}
