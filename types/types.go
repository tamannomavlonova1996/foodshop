package types

import (
	"time"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

type Products struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Price   float64   `json:"price"`
	Created time.Time `json:"created"`
}

type CreateRequest struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}
