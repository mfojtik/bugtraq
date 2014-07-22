package services

import "time"

type Service interface {
	GetListJSON() (string, error) // JSON representation of the issues
	LastUpdate() time.Time
	Id() string // URI mapping for service
}

type Provider interface {
	Connect(string) error
	GetList(string) (struct{}, error)
}
