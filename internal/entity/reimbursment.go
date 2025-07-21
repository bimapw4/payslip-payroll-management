package entity

type ReimbursementCreate struct {
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}

type ReimbursementUpdate struct {
	Id          string `json:"id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}
