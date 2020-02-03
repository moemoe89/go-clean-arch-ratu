//
//  Simple REST API
//
//  Copyright © 2020. All rights reserved.
//

package model_test

import (
	"simple-go-clean-arch/api/v1/api_struct/model"

	"testing"
	"net/http"
)

func TestNewGenericResponse(t *testing.T) {
	status   := http.StatusOK
	success  := 0
	messages := []string{"OK"}

	resp := model.NewGenericResponse(status, success, messages)

	if status != resp.Status {
		t.Errorf("Should return %b, got %b", status, resp.Status)
	}

	if true != resp.Success {
		t.Errorf("Should return %t, got %t", true, resp.Success)
	}

	if messages[0] != resp.Messages[0] {
		t.Errorf("Should return %s, got %s", messages[0], resp.Messages[0])
	}
}

func TestNewPaginationResponse(t *testing.T) {
	page      := 1
	perPage   := 2
	totalPage := 3
	totalData := 4

	resp := model.NewPaginationResponse(page, perPage, totalPage, totalData)

	if page != resp.Page {
		t.Errorf("Should return %b, got %b", page, resp.Page)
	}

	if perPage != resp.PerPage {
		t.Errorf("Should return %b, got %b", perPage, resp.PerPage)
	}

	if totalPage != resp.TotalPage {
		t.Errorf("Should return %b, got %b", totalPage, resp.TotalPage)
	}

	if totalData != resp.TotalData {
		t.Errorf("Should return %b, got %b", totalData, resp.TotalData)
	}
}
