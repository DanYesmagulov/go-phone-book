package repository

import (
	"fmt"
	"strings"

	phonebook "github.com/DanYesmagulov/go-phone-book"
	"github.com/jmoiron/sqlx"
)

type ContactPostgres struct {
	db *sqlx.DB
}

func NewContactPostgres(db *sqlx.DB) *ContactPostgres {
	return &ContactPostgres{db: db}
}

func (r *ContactPostgres) Create(contact phonebook.Contact) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createContactQuery := fmt.Sprintf("INSERT INTO %s (phone, name) VALUES ($1, $2) RETURNING id", contactTable)
	row := tx.QueryRow(createContactQuery, contact.Phone, contact.Name)

	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()

}

func (r *ContactPostgres) GetAll() ([]phonebook.Contact, error) {
	var contacts []phonebook.Contact

	query := fmt.Sprintf("SELECT * FROM %s", contactTable)

	err := r.db.Select(&contacts, query)

	return contacts, err
}

func (r *ContactPostgres) GetById(id int) (phonebook.Contact, error) {
	var contacts phonebook.Contact

	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, contactTable)

	err := r.db.Get(&contacts, query, id)

	return contacts, err
}

func (r *ContactPostgres) DeleteById(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", contactTable)

	_, err := r.db.Exec(query, id)

	return err
}

func (r *ContactPostgres) UpdateById(id int, input phonebook.UpdateContact) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Phone != nil {
		setValues = append(setValues, fmt.Sprintf("phone=$%d", argId))
		args = append(args, *input.Phone)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d",
		contactTable, setQuery, argId)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}
