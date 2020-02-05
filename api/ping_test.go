//
//  Simple REST API
//
//  Copyright © 2020. All rights reserved.
//

package api_test

import (
	"simple-go-clean-arch/routers"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := routers.GetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	if http.StatusOK != w.Code {
		t.Errorf("Should return %v, got %v", http.StatusOK, w.Code)
	}

	assert.Equal(t, http.StatusOK, w.Code)
}
