package services

type Service interface {
	List() (string, error) // JSON representation of the issues
}
