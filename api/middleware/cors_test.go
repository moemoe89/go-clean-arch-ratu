//
//  Simple REST API
//
//  Copyright © 2020. All rights reserved.
//

package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"simple-go-clean-arch/routers"
)

func TestPingRoute(t *testing.T) {
	router := routers.GetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("OPTIONS", "/", nil)
	req.Header.Add("Access-Control-Request-Headers", "*")
	router.ServeHTTP(w, req)

	if http.StatusOK != w.Code {
		t.Errorf("Should return %v, got %v", http.StatusOK, w.Code)
	}
}