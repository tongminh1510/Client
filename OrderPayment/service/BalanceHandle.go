package service

import (
	"OrderPayment/entity"
	"encoding/json"
	"fmt"
	"net/http"
)

var balance = 1000.0

func BalanceHandle(w http.ResponseWriter, r *http.Request) {

	balanceNumber, err := json.Marshal(entity.Balance{currentBalance()})
	if err != nil {
		fmt.Println("Error marshal data:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(balanceNumber)
}

func currentBalance() float64 {
	return balance
}
