package main

import (
	"net/http"
	"net/http/httptest"

	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestMainHandlerWhenRequestIsFormedCorrectly(t *testing.T) {

    req := httptest.NewRequest("Get", "/cafe?count=2&city=moscow", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body.String())
    assert.NotEqual(t, "", responseRecorder.Body.String())
}

func TestMainHandlerWhenСityIsNotSupported(t *testing.T) {

    req := httptest.NewRequest("Get", "/cafe?count=2&city=Kazan", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body.String())
    assert.Equal(t, "wrong city value", responseRecorder.Body.String())

}


func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
    totalCount := 4
    req := httptest.NewRequest("Get", "/cafe?count=10&city=moscow", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body.String())
    list := strings.Split(responseRecorder.Body.String(), ",")

    assert.Equal(t, len(list), totalCount)
    

}