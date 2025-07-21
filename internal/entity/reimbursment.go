package entity

type ReimbursementCreate struct {
	Amount      string `json:"amount"`
	Description string `json:"description"`
	Attachment  string `json:"attachment"`
}

type ReimbursementUpdate struct {
	Id          string `json:"id"`
	Amount      string `json:"amount"`
	Description string `json:"description"`
	Attachment  string `json:"attachment"`
}
