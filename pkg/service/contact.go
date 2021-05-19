package service

import (
	phonebook "github.com/DanYesmagulov/go-phone-book"
	"github.com/DanYesmagulov/go-phone-book/pkg/repository"
)

type ContactService struct {
	repo repository.Contacts
}

func NewContactService(repo repository.Contacts) *ContactService {
	return &ContactService{repo: repo}
}

func (s *ContactService) Create(contact phonebook.Contact) (int, error) {
	return s.repo.Create(contact)
}

func (s *ContactService) GetAll() ([]phonebook.Contact, error) {
	return s.repo.GetAll()
}

func (s *ContactService) GetById(id int) (phonebook.Contact, error) {
	return s.repo.GetById(id)
}

func (s *ContactService) DeleteById(id int) error {
	return s.repo.DeleteById(id)
}

func (s *ContactService) UpdateById(id int, input phonebook.UpdateContact) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.UpdateById(id, input)
}
