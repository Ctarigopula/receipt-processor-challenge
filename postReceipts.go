package routes

import (
	"encoding/json"
	"net/http"
	"receipt-processor-challenge/types"
	"receipt-processor-challenge/variables"

	"github.com/google/uuid"
)

// Post receipt
func CreateReceipt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	var newReceipt types.Receipt
	if err := json.NewDecoder(r.Body).Decode(&newReceipt); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	id := uuid.New().String()
	variables.ReceiptStore[id] = newReceipt
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{"id": id}); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
