package walletoperation

type processRequest struct {
	FromID string `json:"from_ID"`
	ToID   string `json:"to_ID"`
	Amount int    `json:"amount"`
}

type processResponse struct {
	Status string `json:"status"`
}

type errorResponse struct {
	Error string `json:"error"`
}
