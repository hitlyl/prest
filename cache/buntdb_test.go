package cache

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/hitlyl/prest/adapters/postgres"
	"github.com/hitlyl/prest/config"
)

func init() {
	config.Load()
	postgres.Load()
}

func TestBuntGetDoesntExist(t *testing.T) {
	os.Setenv("PREST_CACHE", "true")
	config.Load()
	w := httptest.NewRecorder()

	cache := BuntGet("test", w)
	if cache {
		t.Errorf("expected cache non-existent, but got %t", cache)
	}
}
