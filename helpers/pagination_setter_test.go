//
//  Simple REST API
//
//  Copyright © 2020. All rights reserved.
//

package helpers_test

import (
	"simple-go-clean-arch/helpers"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaginationSetterFailPerPage(t *testing.T) {
	_, _, _, _, msg := helpers.PaginationSetter("a", "")
	expectedMsg := "Invalid parameter per_page: not an int"

	assert.Equal(t, expectedMsg, msg)
}

func TestPaginationSetterFailPage(t *testing.T) {
	_, _, _, _, msg := helpers.PaginationSetter("", "b")
	expectedMsg := "Invalid parameter page: not an int"

	assert.Equal(t, expectedMsg, msg)
}

func TestPaginationSetterPage0(t *testing.T) {
	_, _, _, _, msg := helpers.PaginationSetter("10", "0")
	expectedMsg := ""

	assert.Equal(t, expectedMsg, msg)
}

func TestPaginationSetter(t *testing.T) {
	_, _, _, _, msg := helpers.PaginationSetter("", "")
	expectedMsg := ""

	assert.Equal(t, expectedMsg, msg)
}
