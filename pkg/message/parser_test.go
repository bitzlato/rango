package message

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	bb := []byte(`{"event":"order","data":{"market":"bbbuuu","side":"sell","volume":"0.001","ord_type":"limit","price":"5000"}}`)
	req, err := Parse(bb)
	if err != nil {
		t.Fatalf("Should not return error: %s", err)
	}

	t.Run("Test method", func(t *testing.T) {
		assert.Equal(t, "order", req.Method)
	})

	t.Run("Test message", func(t *testing.T) {
		vv := map[string]interface{}{}

		err := json.Unmarshal(req.Message, &vv)
		if err != nil {
			t.Fatalf("Should not return error: %s", err)
		}

		v, ok := vv["market"]
		assert.Equal(t, true, ok)
		assert.Equal(t, "bbbuuu", v)
	})
}
