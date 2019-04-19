package orderdesk

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/ashkarin/ashkarin-pizza-shop/internal/utils"
	"github.com/ashkarin/ashkarin-pizza-shop/pkg/order"
	"github.com/ashkarin/ashkarin-pizza-shop/pkg/order/usecases"
)

// Service provides a set of HTTP handlers for work with orders
type Service struct {
	storage order.StorageGateway
	router  *mux.Router
}

// NewService creates a service to work with orders
func NewService(s order.StorageGateway, router *mux.Router) *Service {
	service := &Service{
		storage: s,
		router:  router,
	}
	service.initializeRoutes()

	return service
}

func (s *Service) initializeRoutes() {
	s.router.HandleFunc("/", s.IsAlive).Methods("GET")

	// POST [place order] ?/orders
	s.router.HandleFunc("/orders", s.PlaceOrder).Methods("POST")

	// GET [get order] ?/orders/{id}
	s.router.HandleFunc("/orders/{id}", s.GetOrder).Methods("GET")

	// PUT [update order] ?/orders/{id}
	s.router.HandleFunc("/orders/{id}", s.UpdateOrder).Methods("PUT")
}

// IsAlive is the HTTP handler to check whether the service is alive or not
func (s *Service) IsAlive(w http.ResponseWriter, r *http.Request) {
	utils.ResponseWithJSON(w, http.StatusOK, "alive")
}

// PlaceOrder is the HTTP handler to place the order
func (s *Service) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var entry order.Order
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&entry); err != nil {
		log.Errorf("PlaceOrder: %v", err)
		utils.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload JSON format")
		return
	}
	defer r.Body.Close()

	// Create the order in the storage
	var err error
	entry.ID, err = usecases.PlaceOrder(s.storage, &entry)
	if err != nil {
		log.Errorf("PlaceOrder usecase: %v", err)
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseWithJSON(w, http.StatusCreated, entry)
}

// GetOrder is the HTTP handler to get the recipe from storage by ID
func (s *Service) GetOrder(w http.ResponseWriter, r *http.Request) {
	// Get the order ID
	vars := mux.Vars(r)
	id := vars["id"]

	// Get the order
	entry, err := usecases.GetOrderByID(s.storage, id)
	if err != nil {
		log.Errorf("GetOrder: %v", err)
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.ResponseWithJSON(w, http.StatusOK, entry)
}

// UpdateOrder is the HTTP handler to update the order entry in the storage
func (s *Service) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	// Get the order ID
	vars := mux.Vars(r)
	entry := order.Order{ID: vars["id"]}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&entry); err != nil {
		log.Errorf("UpdateOrder: %v", err)
		utils.ResponseWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()

	// Update the order in the storage
	if err := usecases.UpdateOrder(s.storage, &entry); err != nil {
		log.Errorf("UpdateOrder usecase: %v", err)
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseWithJSON(w, http.StatusOK, entry)
}

// DeleteOrder is the HTTP handler to delete the order from the storage
func (s *Service) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	// Get the order ID
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete the recipe
	if err := usecases.DeleteOrderByID(s.storage, id); err != nil {
		log.Errorf("DeleteOrder: %v", err)
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
