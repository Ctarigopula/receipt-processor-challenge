package routes

import (
	"encoding/json"
	"net/http"
	"receipt-processor-challenge/types"
	"receipt-processor-challenge/variables"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/gorilla/mux"
)

// Calculate points for a receipt
func calculatePoints(receipt types.Receipt) string {
	points := 0

	// One point for every alphanumeric character in the retailer name
	for _, char := range receipt.Retailer {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			points++
		}
	}
	// 50 points if total is a round dollar amount with no cents
	if strings.HasSuffix(receipt.Total, ".00") {
		points += 50
	}
	// 25 points if total is a multiple of 0.25
	if total, err := strconv.ParseFloat(receipt.Total, 64); err == nil {
		if int(total*100)%25 == 0 {
			points += 25
		}
	}
	// 5 points for every two items on the receipt
	points += (len(receipt.Items) / 2) * 5
	// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer.
	for _, item := range receipt.Items {
		desc := strings.TrimSpace(item.ShortDescription)
		if len(desc)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(price*0.2 + 0.9999)
		}
	}
	// 6 points if the day in the purchase date is odd
	if purchaseDate, err := time.Parse("2006-01-02", receipt.PurchaseDate); err == nil {
		if purchaseDate.Day()%2 != 0 {
			points += 6
		}
	}
	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	parts := strings.Split(receipt.PurchaseTime, ":")
	if len(parts) == 2 {
		hour, _ := strconv.Atoi(parts[0])
		if (hour == 14) || (hour == 15) {
			points += 10
		}
	}

	return strconv.Itoa(points)
}

// Get Points
func GetPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	receiptID := vars["id"]

	receipt, exists := variables.ReceiptStore[receiptID]
	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	points := calculatePoints(receipt)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"points": points})
}
