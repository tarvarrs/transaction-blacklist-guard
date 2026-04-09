package walletoperation

import "context"

type DecisionService struct {
	checker BlacklistChecker
}

func NewDecisionService(checker BlacklistChecker) DecisionService {
	return DecisionService{checker: checker}
}

func (s DecisionService) Decide(ctx context.Context, op Operation) (Status, error) {
	blocked, err := s.checker.HasAny(ctx, op.FromID, op.ToID)
	if err != nil {
		return "", err
	}
	if blocked {
		return StatusCancel, nil
	}
	return StatusSuccess, nil
}
