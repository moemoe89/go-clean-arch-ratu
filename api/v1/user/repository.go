//
//  Simple REST API
//
//  Copyright © 2020. All rights reserved.
//

package user

import (
	"simple-go-clean-arch/api/v1/api_struct/form"
	"simple-go-clean-arch/api/v1/api_struct/model"

	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rs/xid"
)

// Repository represent the repositories
type Repository interface {
	Get(filter map[string]interface{}, where, orderBy, selectField string) ([]*model.UserModel, error)
	Count(filter map[string]interface{}, where string) (int, error)
	Create(req *form.UserForm) (*model.UserModel, error)
	GetByID(filter map[string]interface{}, where, id, selectField string) (*model.UserModel, error)
	Update(req *form.UserForm, id string) (*model.UserModel, error)
	Delete(id string) error
}

type mysqlRepository struct {
	DBRead  *sqlx.DB
	DBWrite *sqlx.DB
}

// NewMysqlRepository will create an object that represent the Repository interface
func NewMysqlRepository(DBRead *sqlx.DB, DBWrite *sqlx.DB) Repository {
	return &mysqlRepository{DBRead, DBWrite}
}

func (m *mysqlRepository) Get(filter map[string]interface{}, where, orderBy, selectField string) ([]*model.UserModel, error) {
	users := []*model.UserModel{}
	query := fmt.Sprintf("SELECT %s FROM users %s ORDER BY %s LIMIT :limit OFFSET :offset", selectField, where, orderBy)
	namedQuery, args, _ := m.DBRead.BindNamed(query, filter)
	err := m.DBRead.Select(&users, namedQuery, args...)
	return users, err
}

func (m *mysqlRepository) Count(filter map[string]interface{}, where string) (int, error) {
	var count int
	queryCount := fmt.Sprintf("SELECT COUNT(id) FROM users %s", where)
	namedQueryCount, argsCount, _ := m.DBRead.BindNamed(queryCount, filter)
	err := m.DBRead.Get(&count, namedQueryCount, argsCount...)
	return count, err
}

func (m *mysqlRepository) Create(req *form.UserForm) (*model.UserModel, error) {
	user := &model.UserModel{
		ID:        xid.New().String(),
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		Address:   req.Address,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	_, err := m.DBWrite.NamedExec(`INSERT INTO users (id, name, email, phone, address, created_at, updated_at) VALUES (:id, :name, :email, :phone, :address, :created_at, :updated_at)`, user)
	return user, err
}

func (m *mysqlRepository) GetByID(filter map[string]interface{}, where, id, selectField string) (*model.UserModel, error) {
	user := &model.UserModel{}
	if len(selectField) == 0 {
		selectField = "id, name, phone, email, address, created_at, updated_at"
	}
	query := fmt.Sprintf("SELECT %s FROM users  %s", selectField, where)
	namedQuery, args, _ := m.DBRead.BindNamed(query, filter)
	err := m.DBRead.Get(user, namedQuery, args...)
	return user, err
}

func (m *mysqlRepository) Update(req *form.UserForm, id string) (*model.UserModel, error) {
	user := &model.UserModel{
		ID:        id,
		Name:      req.Name,
		Phone:     req.Phone,
		Email:     req.Email,
		Address:   req.Address,
		UpdatedAt: time.Now().UTC(),
	}

	_, err := m.DBWrite.NamedExec(`UPDATE users SET name = :name, phone = :phone, email = :email, address = :address, updated_at = :updated_at WHERE id = :id`, user)
	return user, err
}

func (m *mysqlRepository) Delete(id string) error {
	now := time.Now().UTC()
	user := &model.UserModel{
		ID:        id,
		DeletedAt: &now,
	}
	_, err := m.DBWrite.NamedExec(`UPDATE users SET deleted_at = :deleted_at WHERE id = :id`, user)
	return err
}
