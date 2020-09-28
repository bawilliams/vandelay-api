package router

import (
	"context"

	"github.com/bawilliams/vandelay-api/handler"

	"github.com/pressly/chi"
)

// InitAllEndPoints initializes all routing endpoints
func InitAllEndPoints(ctx context.Context) (r *chi.Mux, _ error) {
	r = chi.NewRouter()

	// r.Get("/health", api.CheckHealth())

	r.Route("/warehouses", func(r chi.Router) {
		r.Route("/{warehouseId}/inventory", func(r chi.Router) {
			r.Get("/", handler.ListInventory)
		})
	})

	return
}
