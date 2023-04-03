package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Order struct {
	Amount float64 `json:"amount"`
}

func main() {
	order := Order{50.0}

	orderBytes, _ := json.Marshal(order)
	resp, err := http.Post("http://localhost:1521/orders", "application/json", bytes.NewBuffer(orderBytes))
	if err != nil {
		fmt.Println("Order failed:", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Order failed. Response status:", resp.StatusCode)
		return
	}

	fmt.Println("Order success!")
}
