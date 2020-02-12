//
//  Simple REST API
//
//  Copyright © 2020. All rights reserved.
//

package user_test

import (
	"github.com/moemoe89/simple-go-clean-arch/api/v1/api_struct/form"
	"github.com/moemoe89/simple-go-clean-arch/api/v1/api_struct/model"
	"github.com/moemoe89/simple-go-clean-arch/api/v1/user/mocks"
	"github.com/moemoe89/simple-go-clean-arch/config"
	"github.com/moemoe89/simple-go-clean-arch/routers"

	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
)

func TestDeliveryUpadte(t *testing.T) {
	id := xid.New().String()

	lang, _ := config.InitLang()
	log := config.InitLog()

	userForm := &form.UserForm{
		ID:      xid.New().String(),
		Name:    "Momo",
		Email:   "momo@mail.com",
		Phone:   "085640",
		Address: "Indonesia",
	}
	user := &model.UserModel{
		ID:      userForm.ID,
		Name:    userForm.Name,
		Email:   userForm.Email,
		Phone:   userForm.Phone,
		Address: userForm.Address,
	}

	j, err := json.Marshal(userForm)
	assert.NoError(t, err)

	mockService := new(mocks.Service)
	mockService.On("Detail", id, "id").Return(user, 0, nil)
	mockService.On("Update", userForm, id).Return(user, 0, nil)

	router:= routers.GetRouter(lang, log, mockService)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/api/v1/user/"+id, strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)
}

func TestDeliveryList(t *testing.T) {
	lang, _ := config.InitLang()
	log := config.InitLog()

	user := &model.UserModel{
		ID:      xid.New().String(),
		Name:    "Momo",
		Email:   "momo@mail.com",
		Phone:   "085640",
		Address: "Indonesia",
	}
	users := []*model.UserModel{}
	users = append(users, user)

	filter := map[string]interface{}{}
	filter["name"] = "%a%"
	filter["email"] = "%b%"
	filter["phone"] = "%0856%"
	filterCount := filter
	filter["limit"] = 10
	filter["offset"] = 0

	mockService := new(mocks.Service)
	mockService.On("List", filter, filterCount, "WHERE deleted_at IS NULL AND name LIKE :name AND email LIKE :email AND phone LIKE :phone", "created_at DESC", model.UserSelectField).Return(users, 1, 0, nil)

	router:= routers.GetRouter(lang, log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/user?per_page=10&name=a&email=b&phone=0856", strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)
}

func TestDeliveryDetail(t *testing.T) {
	id := xid.New().String()

	lang, _ := config.InitLang()
	log := config.InitLog()

	user := &model.UserModel{
		ID:      xid.New().String(),
		Name:    "Momo",
		Email:   "momo@mail.com",
		Phone:   "085640",
		Address: "Indonesia",
	}

	mockService := new(mocks.Service)
	mockService.On("Detail", id, model.UserSelectField).Return(user, 0, nil)

	router:= routers.GetRouter(lang, log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/user/"+id, strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)
}

func TestDeliveryDetailFail(t *testing.T) {
	id := xid.New().String()

	lang, _ := config.InitLang()
	log := config.InitLog()
	mockService := new(mocks.Service)
	mockService.On("Detail", id, model.UserSelectField).Return(nil, http.StatusInternalServerError, errors.New("Unexpected database error"))

	router:= routers.GetRouter(lang, log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/user/"+id, strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestDeliveryDelete(t *testing.T) {
	id := xid.New().String()

	lang, _ := config.InitLang()
	log := config.InitLog()
	mockService := new(mocks.Service)
	mockService.On("Delete", id).Return(0, nil)

	router:= routers.GetRouter(lang, log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/user/"+id, strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeliveryDeleteFail(t *testing.T) {
	id := xid.New().String()

	lang, _ := config.InitLang()
	log := config.InitLog()
	mockService := new(mocks.Service)
	mockService.On("Delete", id).Return(http.StatusInternalServerError, errors.New("Unexpected database error"))

	router:= routers.GetRouter(lang, log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/user/"+id, strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
