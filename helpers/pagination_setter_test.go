//
//  Simple REST API
//
//  Copyright © 2020. All rights reserved.
//

package helpers_test

import (
	"simple-go-clean-arch/helpers"

	"testing"
)

func TestPaginationSetterFailPerPage(t *testing.T) {
	_, _, _, _, msg := helpers.PaginationSetter("a", "")
	expectedMsg := "Invalid parameter per_page: not an int"

	if msg != expectedMsg {
		t.Errorf("Should return %s, got %s", expectedMsg, msg)
	}
}

func TestPaginationSetterFailPage(t *testing.T) {
	_, _, _, _, msg := helpers.PaginationSetter("", "b")
	expectedMsg := "Invalid parameter page: not an int"

	if msg != expectedMsg {
		t.Errorf("Should return %s, got %s", expectedMsg, msg)
	}
}

func TestPaginationSetterPage0(t *testing.T) {
	_, _, _, _, msg := helpers.PaginationSetter("10", "0")
	expectedMsg := ""

	if msg != expectedMsg {
		t.Errorf("Should return %s, got %s", expectedMsg, msg)
	}
}

func TestPaginationSetter(t *testing.T) {
	_, _, _, _, msg := helpers.PaginationSetter("", "")
	expectedMsg := ""

	if msg != expectedMsg {
		t.Errorf("Should return %s, got %s", expectedMsg, msg)
	}
}
