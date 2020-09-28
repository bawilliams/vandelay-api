package model

// This package is meant to mock out calls to fetch data from a database
// Currently, the data is just hard-coded, but we could swap this out for a connection

type Item struct {
	ID          int64  `json:"itemId"`
	Name        string `json:"itemName"`
	SKU         int64  `json:"itemSKU"`
	Quantity    int64  `json:"itemQuantity"`
	Description string `json:"itemDescription"`
	WarehouseID int64  `json:"warehouseId"`
}

func FetchInventory(warehouseID int64) []Item {
	return []Item{
		{
			ID:          1,
			Name:        "Waterproof seal",
			SKU:         12345,
			Quantity:    1000,
			Description: "For sealing specimens to preserve them",
			WarehouseID: 1,
		},
		{
			ID:          2,
			Name:        "Waterproof sealant",
			SKU:         12346,
			Quantity:    1050,
			Description: "Used alongside the waterproof seal in order to seal specimens to preserve them",
			WarehouseID: 1,
		},
	}
}
