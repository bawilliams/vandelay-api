package handler

import (
	"encoding/json"
	"net/http"

	"github.com/bawilliams/vandelay-api/model"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

// List returns a list of all items in a warehouse's inventory
func ListInventory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// log := logging.GetLogEntry(ctx)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	warehouseID := chi.URLParam(r, "warehouseId")
	if warehouseID == "" {
		log.Error("Could not retrieve warehouse id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	inventory, err := model.FetchInventory(warehouseID)
	if err != nil {
		log.WithError(err).Error("could not retrieve inventory")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(inventory)
}
