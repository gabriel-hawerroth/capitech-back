package entity

import (
	"time"
)

type Purchase struct {
	Id            int       `json:"id"`
	UserId        int       `json:"user_id"`
	PurchaseDate  time.Time `json:"purchase_date"`
	Status        string    `json:"status"`
	PaymentMethod string    `json:"payment_method"`
	AddressId     int       `json:"address_id"`
	ShippingValue float64   `json:"shipping_value"`
}
