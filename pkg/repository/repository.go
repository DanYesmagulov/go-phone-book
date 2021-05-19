package repository

import (
	phonebook "github.com/DanYesmagulov/go-phone-book"
	"github.com/jmoiron/sqlx"
)

type Contacts interface {
	Create(contact phonebook.Contact) (int, error)
	GetAll() ([]phonebook.Contact, error)
	GetById(id int) (phonebook.Contact, error)
	DeleteById(id int) error
	UpdateById(id int, input phonebook.UpdateContact) error
}

type Repository struct {
	Contacts
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Contacts: NewContactPostgres(db),
	}
}
