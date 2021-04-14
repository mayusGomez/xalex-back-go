package main

import (
	"github.com/mayusGomez/xalex/customers"
)

type Result struct {
	Code        string             `json:"code,omitempty"`
	Description string             `json:"description,omitempty"`
	Data        customers.Customer `json:"data,omitempty"`
}
