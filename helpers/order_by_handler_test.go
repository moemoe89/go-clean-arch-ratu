//
//  Simple REST API
//
//  Copyright © 2020. All rights reserved.
//

package helpers_test

import (
	"simple-go-clean-arch/api/v1/api_struct/model"
	"simple-go-clean-arch/helpers"

	"testing"
)

func TestOrderByHandler(t *testing.T) {
	selectField := "id"
	orderBy := helpers.OrderByHandler(selectField, model.UserModel{})
	expectedOrderBy := "id ASC"

	if expectedOrderBy != orderBy {
		t.Errorf("Should return %s, got %s", expectedOrderBy, orderBy)
	}
}

func TestOrderByHandlerDesc(t *testing.T) {
	selectField := "-id"
	orderBy := helpers.OrderByHandler(selectField, model.UserModel{})
	expectedOrderBy := "id DESC"

	if expectedOrderBy != orderBy {
		t.Errorf("Should return %s, got %s", expectedOrderBy, orderBy)
	}
}

func TestOrderByHandleEmptyField(t *testing.T) {
	selectField := ""
	orderBy := helpers.OrderByHandler(selectField, model.UserModel{})
	expectedOrderBy := ""

	if expectedOrderBy != orderBy {
		t.Errorf("Should return %s, got %s", expectedOrderBy, orderBy)
	}
}

func TestOrderByHandleNotFoundField(t *testing.T) {
	selectField := "x"
	orderBy := helpers.OrderByHandler(selectField, model.UserModel{})
	expectedOrderBy := ""

	if expectedOrderBy != orderBy {
		t.Errorf("Should return %s, got %s", expectedOrderBy, orderBy)
	}
}
