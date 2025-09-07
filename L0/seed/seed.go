package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/tmozzze/order_checker/internal/models"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 101; i <= 110; i++ {
		Price := rand.Intn(5000) + 100
		totalPrice := Price * (90) / 100
		order := models.Order{
			OrderUID:    fmt.Sprintf("order-%d", i),
			TrackNumber: fmt.Sprintf("TN-%03d", i),
			Entry:       "WBIL",
			Locale:      "en",
			CustomerID:  fmt.Sprintf("user-%d", i),
			Delivery: models.Delivery{
				Name:    fmt.Sprintf("User %d", i),
				Phone:   fmt.Sprintf("+7972%07d", rand.Intn(999999)),
				City:    "Tel Aviv",
				Address: fmt.Sprintf("Street %d", i),
				Email:   fmt.Sprintf("user%d@example.com", i),
			},
			Payment: models.Payment{
				Transaction:  fmt.Sprintf("txn-%d", i),
				Currency:     "USD",
				Provider:     "visa",
				Amount:       rand.Intn(5000) + 100,
				PaymentDt:    time.Now().Unix(),
				Bank:         "Hapoalim",
				DeliveryCost: 500,
				GoodsTotal:   rand.Intn(4000) + 500,
				CustomFee:    0,
			},
			Items: []models.Item{
				{
					ChrtID:      rand.Intn(1000) * i,
					TrackNumber: fmt.Sprintf("TN-%03d", i),
					Price:       Price,
					RID:         fmt.Sprintf("RID-%d", (rand.Intn(1000) * i)),
					Name:        "MacBook Air",
					Sale:        10,
					Size:        "0",
					TotalPrice:  totalPrice,
					NmID:        rand.Intn(1000) * i,
					Brand:       "Apple",
					Status:      202,
				},
			},
		}

		body, _ := json.Marshal(order)

		resp, err := http.Post("http://localhost:8080/orders", "application/json", bytes.NewReader(body))
		if err != nil {
			fmt.Println("Response error:", err)
			continue
		}
		defer resp.Body.Close()

		fmt.Printf("Order: %s, status %s\n", order.OrderUID, resp.Status)
	}
}
