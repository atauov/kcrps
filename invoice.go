package dashboard

type Invoice struct {
	Id      int    `json:"id" db:"id"`
	Amount  int    `json:"amount" db:"amount" binding:"required"`
	Account string `json:"account" db:"account" binding:"required"`
	Message string `json:"message" db:"message" binding:"required"`
}
