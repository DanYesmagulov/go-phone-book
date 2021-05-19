package service

import (
	phonebook "github.com/DanYesmagulov/go-phone-book"
	"github.com/DanYesmagulov/go-phone-book/pkg/repository"
)

type Contacts interface {
	Create(contact phonebook.Contact) (int, error)
	GetAll() ([]phonebook.Contact, error)
	GetById(id int) (phonebook.Contact, error)
	DeleteById(id int) error
	UpdateById(id int, input phonebook.UpdateContact) error
}

type Service struct {
	Contacts
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Contacts: NewContactService(repos.Contacts),
	}
}
