package dto

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKopeksUnmarshalJSON(t *testing.T) {
	var kops Kopeks
	assert.NoError(t, json.Unmarshal([]byte(`"420.69"`), &kops))
	assert.Equal(t, Kopeks(42069), kops)

	bytes, err := json.Marshal(Kopeks(314))
	assert.NoError(t, err)
	assert.Equal(t, `"3.14"`, string(bytes))
}

func TestKopeksString(t *testing.T) {
	assert.Equal(t, "0.00", Kopeks(0).String())
	assert.Equal(t, "0.01", Kopeks(1).String())
	assert.Equal(t, "0.10", Kopeks(10).String())
	assert.Equal(t, "1.10", Kopeks(110).String())
	assert.Equal(t, "100.00", Kopeks(10000).String())
}
