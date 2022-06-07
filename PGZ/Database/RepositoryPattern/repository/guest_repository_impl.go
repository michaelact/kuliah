package repository

import (
	"database/sql"
	"context"
	"strconv"
	"errors"

	"repository_pattern/entity"
)

type guestRepositoryImpl struct {
	DB *sql.DB
}

func NewGuestRepository(db *sql.DB) GuestRepository {
	return &guestRepositoryImpl{ DB: db }
}

func (repository *guestRepositoryImpl) Insert(ctx context.Context, guest entity.Guest) (entity.Guest, error) {
	script := "INSERT INTO MyGuests(firstname, lastname, email) VALUES (?, ?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, guest.Firstname, guest.Lastname, guest.Email)
	if err != nil {
		return guest, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return guest, err
	}

	guest.Id = int(id)
	return guest, nil
}

func (repository *guestRepositoryImpl) FindById(ctx context.Context, id int) (entity.Guest, error) {
	script := "SELECT id, firstname, lastname, email, reg_date FROM MyGuests WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)

	var guest entity.Guest
	if err != nil {
		return guest, err
	}

	defer rows.Close()
	if rows.Next() {
		rows.Scan(&guest.Id, &guest.Firstname, &guest.Lastname, &guest.Email, &guest.RegDate)
		return guest, nil
	} else {
		return guest, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}

func (repository *guestRepositoryImpl) FindAll(ctx context.Context) ([]entity.Guest, error) {
	script := "SELECT id, firstname, lastname, email, reg_date FROM MyGuests"
	rows, err := repository.DB.QueryContext(ctx, script)

	var guests []entity.Guest
	if err != nil {
		return guests, err
	}

	defer rows.Close()
	for rows.Next() {
		var guest entity.Guest
		rows.Scan(&guest.Id, &guest.Firstname, &guest.Lastname, &guest.Email, &guest.RegDate)
		guests = append(guests, guest)
	}

	return guests, nil
}
