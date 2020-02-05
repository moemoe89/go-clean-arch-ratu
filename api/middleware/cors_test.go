//
//  Simple REST API
//
//  Copyright © 2020. All rights reserved.
//

package middleware_test

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
	req, _ := http.NewRequest("OPTIONS", "/", nil)
	req.Header.Add("Access-Control-Request-Headers", "*")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}