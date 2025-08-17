package database

type User struct {
	ID int   `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type Messages struct{
	Message string `json:"message"`
	Options []string `json:"option"`
}
type Resposes struct{
	Response string `json:"response"`
}

type BalanceResponse struct{
	UserName string `json:"username"`
	Balance float64 `json:"balance"`
}


type TransactionRequst struct {
	Amount float64 `json:"amount"`
}

type TransectionHistory struct{
	TransectionInfo string `json:"transection"`
	Amount          float64 `json:"amount"` 
	CreatedAt       string  `json:"created_at,omitempty"`
}