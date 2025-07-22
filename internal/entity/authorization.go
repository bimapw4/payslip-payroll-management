package entity

type Claim struct {
	UserID   string
	Username string
	IsAdmin  bool
	Exp      int
}

type Authorization struct {
	Username string `json:"username"`
	Password string `json:"password"`
	// Type         string `json:"type"`
	// RefreshToken string `json:"refresh_token"`
}
