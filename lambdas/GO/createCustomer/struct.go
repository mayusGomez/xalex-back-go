package main

import (
	"github.com/mayusGomez/xalex/customers"
)

type Result struct {
	Code        string
	Description string
	Data        customers.Customer
}
