package dashboard

type Invocie struct {
	Id      int    `json:"id"`
	Amount  int    `json: "amount"`
	Account string `json:"account"`
	Message string `json:"message"`
}
