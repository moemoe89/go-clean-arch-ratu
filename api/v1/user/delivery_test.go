//
//  Simple REST API
//
//  Copyright © 2020. All rights reserved.
//

package user_test

import (
	"github.com/moemoe89/simple-go-clean-arch/api/v1/user/mocks"
	"github.com/moemoe89/simple-go-clean-arch/config"
	"github.com/moemoe89/simple-go-clean-arch/api/v1/api_struct/model"
	"github.com/moemoe89/simple-go-clean-arch/routers"

	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
)

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
