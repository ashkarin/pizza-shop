package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/ashkarin/ashkarin-pizza-shop/internal/config"
	"github.com/ashkarin/ashkarin-pizza-shop/internal/services/orderdesk"
	orderGW "github.com/ashkarin/ashkarin-pizza-shop/pkg/order/gateways"
)

// Server is the app container
type Server struct {
	orderdeskService *orderdesk.Service
	Router           *mux.Router
	server           *http.Server
}

// Initialize init the server
func (s *Server) Initialize(cfg *config.Config) {
	// Log
	log.Infof("Initialize server with: %v", cfg)

	// Open a gateway to the MongoDB storage
	ordersStorage, err := orderGW.NewMongoDbGateway(
		cfg.DB.Server, cfg.DB.Port, cfg.DB.Username, cfg.DB.Password,
		cfg.DB.DBName, "orders")

	if err != nil {
		log.Fatalf("Connection to the recipes storage: %v", err)
	}

	// Create route and service
	s.Router = mux.NewRouter()
	s.orderdeskService = orderdesk.NewService(ordersStorage, s.Router)

	// Create the server
	s.server = &http.Server{
		Handler:      s.Router,
		Addr:         fmt.Sprintf("%s:%s", cfg.Address, cfg.Port),
		WriteTimeout: time.Duration(cfg.Timeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.Timeout) * time.Second,
	}
}

// ListenAndServe start the server
func (s *Server) ListenAndServe() {
	log.Fatal(s.server.ListenAndServe())
}
