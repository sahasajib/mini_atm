package handler

import (
	"net/http"

	"github.com/sahasajib/mini_atm/util"
)

func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, "all transactions", http.StatusOK)
}