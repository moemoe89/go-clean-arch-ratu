//
//  Simple REST API
//
//  Copyright © 2020. All rights reserved.
//

package user_test

import (
	"simple-go-clean-arch/api/v1/api_struct/model"
	"simple-go-clean-arch/api/v1/user"

	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/rs/xid"
)

func TestGet(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db,"sqlmock")

	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "address", "created_at", "updated_at", "deleted_at"}).
		AddRow(xid.New().String(), "Momo", "momo@mail.com", "085640", "Indonesia", time.Now().UTC(), time.Now().UTC(), nil)

	query := "SELECT "+model.UserSelectField+" FROM users WHERE deleted_at IS NULL ORDER BY id ASC LIMIT \\? OFFSET \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	u := user.NewMysqlRepository(sqlxDB, sqlxDB)

	orderBy := "id ASC"
	where := "WHERE deleted_at IS NULL"
	filter := map[string]interface{}{}
	filter["limit"]  = "10"
	filter["offset"] = "0"

	_, err = u.Get(filter, where, orderBy, "")
	if err != nil {
		t.Errorf("Should return %v, got %s", nil, err.Error())
	}
}

func TestCount(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db,"sqlmock")

	rows := sqlmock.NewRows([]string{"COUNT(id)"}).AddRow(1)

	query := "SELECT COUNT\\(id\\) FROM users WHERE deleted_at IS NULL"

	mock.ExpectQuery(query).WillReturnRows(rows)
	u := user.NewMysqlRepository(sqlxDB, sqlxDB)

	where := "WHERE deleted_at IS NULL"
	filter := map[string]interface{}{}

	_, err = u.Count(filter, where)
	if err != nil {
		t.Errorf("Should return %v, got %s", nil, err.Error())
	}
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db,"sqlmock")

	req := &model.UserModel{
		ID:        xid.New().String(),
		Name:      "Momo",
		Email:     "momo@mail.com",
		Phone:     "085640",
		Address:   "Indonesia",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	query := "INSERT INTO users \\(id, name, email, phone, address, created_at, updated_at\\) VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?, \\?\\)"

	mock.ExpectExec(query).WithArgs(req.ID, req.Name, req.Email, req.Phone, req.Address, req.CreatedAt, req.UpdatedAt).WillReturnResult(sqlmock.NewResult(0, 1))

	u := user.NewMysqlRepository(sqlxDB, sqlxDB)
	_, err = u.Create(req)
	if err != nil {
		t.Errorf("Should return %v, got %s", nil, err.Error())
	}
}

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db,"sqlmock")

	req := &model.UserModel{
		ID:        xid.New().String(),
		Name:      "Momo",
		Email:     "momo@mail.com",
		Phone:     "085640",
		Address:   "Indonesia",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "address", "created_at", "updated_at"}).
		AddRow(req.ID, req.Name, req.Email, req.Phone, req.Address, req.CreatedAt, req.UpdatedAt)

	query := "SELECT "+model.UserSelectField+" FROM users WHERE deleted_at IS NULL AND id = \\?"

	mock.ExpectQuery(query).WithArgs(req.ID).WillReturnRows(rows)
	u := user.NewMysqlRepository(sqlxDB, sqlxDB)

	_, err = u.GetByID(req.ID, "")
	if err != nil {
		t.Errorf("Should return %v, got %s", nil, err.Error())
	}
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db,"sqlmock")

	req := &model.UserModel{
		ID:        xid.New().String(),
		Name:      "Momo",
		Email:     "momo@mail.com",
		Phone:     "085640",
		Address:   "Indonesia",
		UpdatedAt: time.Now().UTC(),
	}

	query := "UPDATE users SET name = \\?, email = \\?, phone = \\?, address = \\?, updated_at = \\? WHERE id = \\?"

	mock.ExpectExec(query).WithArgs(req.Name, req.Email, req.Phone, req.Address, req.UpdatedAt, req.ID).WillReturnResult(sqlmock.NewResult(0, 1))
	u := user.NewMysqlRepository(sqlxDB, sqlxDB)

	_, err = u.Update(req)
	if err != nil {
		t.Errorf("Should return %v, got %s", nil, err.Error())
	}
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db,"sqlmock")

	id := xid.New().String()

	query := "UPDATE users SET deleted_at = UTC_TIMESTAMP\\(\\) WHERE id = \\?"

	mock.ExpectExec(query).WithArgs(id).WillReturnResult(sqlmock.NewResult(0, 1))
	u := user.NewMysqlRepository(sqlxDB, sqlxDB)

	err = u.Delete(id)
	if err != nil {
		t.Errorf("Should return %v, got %s", nil, err.Error())
	}
}
