package walletoperation

import "context"

type BlacklistChecker interface {
	HasAny(ctx context.Context, ids ...ParticipantID) (bool, error)
}
