package handler

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCurrencyAPI(t *testing.T) {
	cl := http.DefaultClient
	supportedSymbols := []string{"all"}

	_, err := NewCurrencyAPI(cl, supportedSymbols)
	//assert.Equal(t, interface{}, currencyAPI)
	assert.Equal(t, nil, err)
}
