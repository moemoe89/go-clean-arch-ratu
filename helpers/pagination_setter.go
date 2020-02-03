//
//  Simple REST API
//
//  Copyright © 2020. All rights reserved.
//

package helpers

import (
	"strconv"
)

// PaginationSetter represent the helpers for pagination
func PaginationSetter(perPageString, pageString string) (bool, int, int, int, string) {

	if len(perPageString) == 0 {
		perPageString = "10"
	}
	perPage, err := strconv.Atoi(perPageString)
	if err != nil {
		return false, 0, 0, 0, "Invalid parameter per_page: not an int"
	}

	if len(pageString) == 0 {
		pageString = "1"
	}
	page, err := strconv.Atoi(pageString)
	if err != nil {
		return false, 0, 0, 0, "Invalid parameter page: not an int"
	}

	showPage := page
	if showPage < 1 {
		showPage = 1
		page = 1
	}

	offset := (page - 1) * perPage

	return true, offset, perPage, showPage, ""
}
