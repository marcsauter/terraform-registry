package registry_test

import (
	"encoding/json"
	"testing"

	"github.com/marcsauter/tfregistry/internal/registry"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOS(t *testing.T) {
	type os struct {
		OS registry.OS `json:"os"`
	}

	obj := os{
		OS: registry.OSLinux,
	}

	str := []byte(`{"os":"linux"}`)

	t.Run("marshal", func(t *testing.T) {
		act, err := json.Marshal(&obj)
		require.NoError(t, err)
		assert.Equal(t, str, act)
	})

	t.Run("unmarshal", func(t *testing.T) {
		act := os{}
		err := json.Unmarshal(str, &act)
		require.NoError(t, err)
		assert.Equal(t, obj, act)
	})
}

func TestUnmarshalOS(t *testing.T) {

}
