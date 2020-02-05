//
//  Simple REST API
//
//  Copyright © 2020. All rights reserved.
//

package helpers_test

import (
	"simple-go-clean-arch/api/v1/api_struct/model"
	"simple-go-clean-arch/helpers"

	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckInDBAtt(t *testing.T) {
	selectField := model.UserSelectField
	filterField := model.UserSelectField
	res := helpers.CheckInDBAtt(model.UserModel{}, filterField)
	if len(res) > 0 {
		selectField = strings.Join(res, ",")
	}

	assert.Equal(t, selectField, filterField)
}

func TestCheckMatchDBAtt(t *testing.T) {
	selectField := "id"
	field := helpers.CheckMatchDBAtt(model.UserModel{}, selectField)

	assert.Equal(t, selectField, field)
}
