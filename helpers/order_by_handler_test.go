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

	"github.com/stretchr/testify/assert"
)

func TestOrderByHandler(t *testing.T) {
	selectField := "id"
	orderBy := helpers.OrderByHandler(selectField, model.UserModel{})
	expectedOrderBy := "id ASC"

	assert.Equal(t, expectedOrderBy, orderBy)
}

func TestOrderByHandlerDesc(t *testing.T) {
	selectField := "-id"
	orderBy := helpers.OrderByHandler(selectField, model.UserModel{})
	expectedOrderBy := "id DESC"

	assert.Equal(t, expectedOrderBy, orderBy)
}

func TestOrderByHandleEmptyField(t *testing.T) {
	selectField := ""
	orderBy := helpers.OrderByHandler(selectField, model.UserModel{})
	expectedOrderBy := ""

	assert.Equal(t, expectedOrderBy, orderBy)
}

func TestOrderByHandleNotFoundField(t *testing.T) {
	selectField := "x"
	orderBy := helpers.OrderByHandler(selectField, model.UserModel{})
	expectedOrderBy := ""

	assert.Equal(t, expectedOrderBy, orderBy)
}
