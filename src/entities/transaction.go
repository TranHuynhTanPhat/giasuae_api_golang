package entities

type Transactionhistories struct {
	ID         int    `json:"id"`
	Amount     int    `json:"amount"`
	Content    string `json:"content"`
	ID_account string `json:"id_account"`
	Status     int    `json:"status"`
	Created_at string `json:"created_at"`
}

type TransactionhistoriesReq struct {
	ID         int    `json:"id"`
	Amount     int    `json:"amount"`
	Content    string `json:"content"`
	ID_account int    `json:"id_account"`
	Status     int    `json:"status"`
	Created_at string `json:"created_at"`
}

type Statistics struct {
	Budget_month    int    `json:"budget_month"`
	Budget_year     int    `json:"budget_year"`
	Created_at      string `json:"created_at"`
	Status_newclass int    `json:"status_newclass"`
	Count_trans     int    `json:"count_trans"`
}
