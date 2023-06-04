package transactiondto

type CreateTransaction struct {
	CounterQty int `json:"counter_qty" form:"counter_qty" validate:"required"`
	Total      int `json:"total" form:"total" validate:"required"`
	IdTrip     int `json:"id_trip" form:"id_trip" validate:"required"`
}

type UpdateTransaction struct {
	CounterQty int    `json:"counter_qty" form:"counter_qty"`
	Total      int    `json:"total" form:"total"`
	Status     string `json:"status" form:"status"`
	Attachment string `json:"attachment" form:"attachment"`
	IdTrip     int    `json:"id_trip" form:"id_trip"`
}
