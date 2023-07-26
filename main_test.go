package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"AsiaYo_BackendEngineer/controllers"

	"github.com/gin-gonic/gin"
)

func TestConvertCurrency(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/convert_currency", controllers.ConvertCurrency)

	tests := []struct {
		source     string
		target     string
		amount     string
		statusCode int
		response   string
	}{
		{"TWD", "USD", "$1,000", http.StatusOK, `{"amount":"$32.81","msg":"success"}`},
		{"USD", "JPY", "$1,525", http.StatusOK, `{"amount":"$170,496.53","msg":"success"}`},
		{"EUR", "JPY", "$100", http.StatusBadRequest, `{"error":"source not found"}`},
		{"TWD", "JPY", "abc", http.StatusBadRequest, `{"error":"strconv.ParseFloat: parsing \"abc\": invalid syntax"}`},
	}

	for _, test := range tests {
		url := "/convert_currency?source=" + test.source + "&target=" + test.target + "&amount=" + test.amount
		req, _ := http.NewRequest(http.MethodGet, url, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != test.statusCode {
			t.Errorf("Expected StatusCode should be %d but got %d", test.statusCode, w.Code)
		}

		if strings.TrimSpace(w.Body.String()) != test.response {
			t.Errorf("Expected response should be %s but got %s", test.response, w.Body.String())
		}

	}
}
