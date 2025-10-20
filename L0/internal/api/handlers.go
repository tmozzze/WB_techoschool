package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tmozzze/order_checker/internal/models"
	"github.com/tmozzze/order_checker/internal/service"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(svc *service.OrderService) *OrderHandler {
	return &OrderHandler{service: svc}
}

func (h *OrderHandler) RegisterRoutes(r chi.Router) {
	r.Post("/orders", h.SaveOrder)
	r.Get("/orders/{id}", h.GetOrder)
}

// GetOrder godoc
// @Summary Получить заказ по ID
// @Description Возвращает заказ по идентификатору
// @Tags orders
// @Param id path string true "Order ID"
// @Produce json
// @Success 200 {object} models.Order
// @Failure 400 {object} string "Missing order id"
// @Failure 404 {object} string "Order not found"
// @Router /orders/{id} [get]
func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "missing order id", http.StatusBadRequest)
		return
	}

	order, err := h.service.GetOrder(r.Context(), id)
	if err != nil {
		http.Error(w, "order not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)

}

// SaveOrder godoc
// @Summary Сохранить заказ
// @Description Сохраняет заказ в базе данных
// @Tags orders
// @Accept json
// @Produce json
// @Param order body models.Order true "Order data"
// @Success 202 {string} string "Accepted"
// @Failure 400 {object} string "Invalid request body"
// @Failure 409 {string} string "Order already exists"
// @Failure 500 {object} string "Failed to save order"
// @Router /orders [post]
func (h *OrderHandler) SaveOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order

	// Decode json
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := order.Validate(); err != nil {
		http.Error(w, "invalid order: "+err.Error(), http.StatusBadRequest)
		log.Printf("Validation failed: %v", err)
		return
	}

	// Service layer
	ctx := r.Context()
	if err := h.service.SaveOrder(ctx, &order); err != nil {
		http.Error(w, "failed to save order", http.StatusInternalServerError)
		log.Printf("Postgres error: %v", err)
		return
	}

	// Success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	log.Printf("Order %s saved", order.OrderUID)
}
