package main

type Request struct {
	User  string `json:"user,omitempty"`
	Email string `json:"email,omitempty"`
}
