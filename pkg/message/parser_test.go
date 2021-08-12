package message

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	const (
		isAuthClient    = true
		isNotAuthClient = false
	)

	bb := []byte(`{"event":"order","data":{"market":"bbbuuu","side":"sell","volume":"0.001","ord_type":"limit","price":"5000"}}`)

	t.Run("Test auth client", func(t *testing.T) {
		req, err := Parse(bb, isAuthClient)
		if err != nil {
			t.Fatalf("Should not return error: %s", err)
		}

		assert.Equal(t, "order", req.Method)

		vv := map[string]interface{}{}

		err = json.Unmarshal(req.Message, &vv)
		if err != nil {
			t.Fatalf("Should not return error: %s", err)
		}

		v, ok := vv["market"]
		assert.Equal(t, true, ok)
		assert.Equal(t, "bbbuuu", v)
	})

	t.Run("Test unauth client", func(t *testing.T) {
		req, err := Parse(bb, isNotAuthClient)
		if err == nil {
			t.Fatalf("Should return error")
		}

		assert.Equal(t, "order", req.Method)
	})
}
