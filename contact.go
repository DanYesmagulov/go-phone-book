package phonebook

import "errors"

type Contact struct {
	Id    int     `json:"id" db:"id"`
	Phone *string `json:"phone" db:"phone" binding:"required"`
	Name  *string `json:"name" db:"name" binding:"required"`
}

type UpdateContact struct {
	Phone *string `json:"phone"`
	Name  *string `json:"name"`
}

func (i UpdateContact) Validate() error {
	if i.Name == nil && i.Phone == nil {
		return errors.New("поля измения пустые")
	}

	return nil
}
