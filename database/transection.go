package database

type User struct {
	ID int   `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type Messages struct{
	Message string `json:"message"`
}